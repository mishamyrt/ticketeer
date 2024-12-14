package config

import (
	"os"

	"github.com/mishamyrt/ticketeer/internal/tpl"
	"gopkg.in/yaml.v3"
)

type YAMLMessageConfig struct {
	Location *string `yaml:"location"`
	Format   *string `yaml:"format"`
}

type YAMLConfig struct {
	AllowEmpty *bool             `yaml:"allow_empty"`
	Message    YAMLMessageConfig `yaml:"message"`
}

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

func FromYAML(path string) (Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = newErrFileNotFound(path)
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
