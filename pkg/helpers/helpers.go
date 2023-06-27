// Package helpers implements several useful functions
package helpers

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-acme/lego/v4/certificate"
)

// FileExists takes a string returns if it is an existing file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FolderExists takes a string returns if it is an existing folder
func FolderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// DeleteFile removes the file if it exists
func DeleteFile(filename string) (bool, error) {
	var err error
	if FileExists(filename) {
		err = os.Remove(filename)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, fmt.Errorf("file %s does not exist", filename)
}

func AtomicJSON(file string, data interface{}) error {
	if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
		return err
	}
	tempFile := file + ".tmp"
	f, err := os.Create(tempFile)
	if err != nil {
		return err
	}
	defer os.Remove(tempFile)
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	err = enc.Encode(data)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return os.Rename(tempFile, file)
}

func ParseCert(cert *certificate.Resource) (*x509.Certificate, error) {
	info, _ := pem.Decode(cert.Certificate)
	return x509.ParseCertificate(info.Bytes)
}
