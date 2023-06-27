// Package config implements all configuration aspects of KoboMail
package config

import (
	"github.com/bjw-s/lego-auto/pkg/helpers"
	"github.com/gookit/validate"
)

// ValidateFolder validates that the path is a valid folder
func (c Config) ValidateFolder(val string) bool {
	return helpers.FolderExists(val)
}

// Validate returns if the given configuration is valid and any validation errors
func (c *Config) Validate() validate.Errors {
	v := validate.Struct(c)
	v.StopOnError = false
	return v.ValidateE()
}

func (c Config) Messages() map[string]string {
	return validate.MS{
		"ValidateFolder": "{field} must point to a valid folder.",
	}
}
