package config

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

const (
	delimeter = "."
	tagName   = "koanf"

	upTemplate     = "================ Loaded Configuration ================"
	bottomTemplate = "======================================================"
)

func Load(print bool) (*Config, error) {
	k := koanf.New(delimeter)
	// Load configuration from file
	if err := k.Load(file.Provider("./config.yaml"), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("error loading configuration from file: %v", err)
	}

	// Load configuration from environment variables
	if err := k.Load(env.Provider(tagName, ".", func(s string) string {
		return s
	}), nil); err != nil {
		return nil, fmt.Errorf("error loading configuration from environment variables: %v", err)
	}

	// Unmarshal the configuration into your Config struct
	var config Config
	if err := k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: tagName}); err != nil {
		return nil, fmt.Errorf("error unmarshalling configuration: %v", err)
	}

	if print {
		fmt.Printf("%s\n%v\n%s\n", upTemplate, spew.Sdump(config), bottomTemplate)
	}

	return &config, nil
}
