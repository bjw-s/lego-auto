// Package config implements all configuration aspects of lego-auto
package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	yaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	flag "github.com/spf13/pflag"
)

// Config config struct
type Config struct {
	DataDir     string        `koanf:"datadir" validate:"required|ValidateFolder"`
	CacheDir    string        `koanf:"cachedir" validate:"required"`
	Domains     []string      `koanf:"domains" validate:"required"`
	DNS         []string      `koanf:"dns"`
	Email       string        `koanf:"email" validate:"required|email"`
	Provider    string        `koanf:"provider" validate:"required"`
	Directory   string        `koanf:"directory" validate:"ValidateDirectory"`
	RenewBefore time.Duration `koanf:"renewbefore"`
	Timeout     time.Duration `koanf:"timeout"`
	k           *koanf.Koanf
}

// LoadConfig instantiates a new Config
func LoadConfig(flags *flag.FlagSet) (*Config, error) {
	var err error

	const envVarPrefix = "LA_"

	cwd, _ := os.Getwd()

	var k = koanf.New(".")

	// Fetch flags
	if err = k.Load(posflag.Provider(flags, ".", k), nil); err != nil {
		return nil, err
	}

	// Defaults
	err = k.Load(confmap.Provider(map[string]interface{}{
		"datadir":     cwd,
		"cachedir":    fmt.Sprintf("%s/.cache", cwd),
		"directory":   "production",
		"dns":         []string{"8.8.8.8"},
		"renewbefore": "720h",
		"timeout":     "1m",
	}, ""), nil)
	if err != nil {
		return nil, err
	}

	// YAML Config
	yamlConfig := k.String("config")
	if yamlConfig != "" {
		err = k.Load(file.Provider(yamlConfig), yaml.Parser())
		if err != nil {
			return nil, err
		}
	}

	// Environment variables
	err = k.Load(env.ProviderWithValue(envVarPrefix, ".", func(s string, v string) (string, interface{}) {
		// Process and normalize the key
		s = strings.TrimPrefix(s, envVarPrefix)
		s = strings.ToLower(s)
		s = strings.Replace(s, "__", ".", -1)

		// If there is a space in the value, split the value into a slice by the space.
		if strings.Contains(v, " ") {
			return s, strings.Split(v, " ")
		}

		// Otherwise, return the plain string.
		return s, v
	}), nil)
	if err != nil {
		return nil, err
	}

	// Load flag overrides again to make sure they override everything
	if err = k.Load(posflag.Provider(flags, ".", k), nil); err != nil {
		return nil, err
	}

	var out Config
	err = k.Unmarshal("", &out)
	if err != nil {
		return nil, err
	}

	out.k = k
	return &out, nil
}
