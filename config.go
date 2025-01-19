package joute

import (
	"io"
	"os"
	"path/filepath"
)

// ConfigFileSource represents a generic source from where the config will be loaded
type ConfigFileSource interface {
	Reader() (io.Reader, error)
}

// All sources supported by Joute
type (
	ConfigFileLocation string
	WorkingDirectory   struct{}
)

// Configuration file defaults
const (
	DefaultConfigFileName     = ".jouterc"
	DefaultConfigFileLocation = "."
)

func (loc ConfigFileLocation) Reader() (io.Reader, error) {
	return os.Open(filepath.Join(string(loc), DefaultConfigFileName))
}

func (dir WorkingDirectory) Reader() (io.Reader, error) {
	return os.Open(filepath.Join(".", DefaultConfigFileName))
}
