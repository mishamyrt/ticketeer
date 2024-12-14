package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/mishamyrt/ticketeer/internal/tpl"
	"gopkg.in/yaml.v3"
)

// ErrFileNotFound is returned when configuration file is not found
var ErrFileNotFound = errors.New("config file not found")

// YAMLMessageConfig represents message configuration
type YAMLMessageConfig struct {
	Location *string `yaml:"location"`
	Format   *string `yaml:"format"`
}

// YAMLConfig represents yaml configuration
type YAMLConfig struct {
	AllowEmpty *bool             `yaml:"allow_empty"`
	Message    YAMLMessageConfig `yaml:"message"`
}

// ParseYAML parses yaml configuration
func ParseYAML(raw YAMLConfig) (config Config, err error) {
	config = defaultConfig
	if raw.AllowEmpty != nil {
		config.AllowEmpty = *raw.AllowEmpty
	}
	if raw.Message.Location != nil {
		config.TicketLocation, err = ParseLocation(*raw.Message.Location)
		if err != nil {
			return
		}
	}
	if raw.Message.Format != nil {
		config.Template = tpl.Template(*raw.Message.Format)
	} else {
		config.Template = defaultTemplates[config.TicketLocation]
	}

	return
}

// FromYAML reads and parses yaml configuration
func FromYAML(path string) (Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("%w at %s", ErrFileNotFound, path)
		}
		return Config{}, err
	}
	var raw YAMLConfig
	err = yaml.Unmarshal(content, &raw)
	if err != nil {
		return Config{}, err
	}

	return ParseYAML(raw)
}
