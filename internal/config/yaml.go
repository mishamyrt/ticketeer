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
	Template *string `yaml:"template"`
}

// YAMLBranchConfig represents branch configuration
type YAMLBranchConfig struct {
	Format *string  `yaml:"format"`
	Ignore []string `yaml:"ignore"`
}

// YAMLConfig represents yaml configuration
type YAMLConfig struct {
	Message YAMLMessageConfig `yaml:"message"`
	Ticket  YAMLTicketConfig  `yaml:"ticket"`
	Branch  YAMLBranchConfig  `yaml:"branch"`
}

// ParseYAML parses yaml configuration
func ParseYAML(raw YAMLConfig) (config Config, err error) {
	config.Ticket, err = ParseYAMLTicket(raw.Ticket)
	if err != nil {
		return
	}
	config.Branch, err = ParseYAMLBranch(raw.Branch)
	if err != nil {
		return
	}
	config.Message, err = ParseYAMLMessage(raw.Message)
	return
}

// FromYAMLFile reads and parses yaml configuration
func FromYAMLFile(path string) (Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("%w at %s", ErrFileNotFound, path)
		}
		return defaultConfig, err
	}
	var raw YAMLConfig
	err = yaml.Unmarshal(content, &raw)
	if err != nil {
		return Config{}, err
	}

	return ParseYAML(raw)
}

// YAMLTicketConfig represents ticket configuration
type YAMLTicketConfig struct {
	Format     *string `yaml:"format"`
	AllowEmpty *bool   `yaml:"allow_empty"`
}

// ParseYAMLTicket parses ticket configuration
func ParseYAMLTicket(raw YAMLTicketConfig) (TicketConfig, error) {
	config := defaultConfig.Ticket
	var err error
	if raw.Format != nil {
		config.Format, err = ParseTicketFormat(*raw.Format)
		if err != nil {
			return config, err
		}
	}
	if raw.AllowEmpty != nil {
		config.AllowEmpty = *raw.AllowEmpty
	}
	return config, nil
}

// ParseYAMLBranch parses branch configuration
func ParseYAMLBranch(raw YAMLBranchConfig) (BranchConfig, error) {
	config := defaultConfig.Branch
	if raw.Format != nil {
		format, err := ParseBranchFormat(*raw.Format)
		if err != nil {
			return config, err
		}
		config.Format = format.TicketFormat()
	}
	if len(raw.Ignore) > 0 {
		config.Ignore = append(config.Ignore, raw.Ignore...)
	}
	return config, nil
}

// ParseYAMLMessage parses message configuration
func ParseYAMLMessage(raw YAMLMessageConfig) (MessageConfig, error) {
	config := defaultConfig.Message
	var err error
	if raw.Location != nil {
		config.Location, err = ParseTicketLocation(*raw.Location)
		if err != nil {
			return config, err
		}
	}
	if raw.Template != nil {
		config.Template = tpl.Template(*raw.Template)
	} else {
		config.Template = defaultTemplates[config.Location]
	}
	return config, nil
}
