package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type DbConfig struct {
	Driver string `yaml:"driver"`
}

type Config struct {
	DB DbConfig `yaml:"db"`
	Version string `yaml:"version"`
}

func LoadConfig(filename string) (*Config, error) {
	// := Creates a variable and infers the type
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config = &Config{}
	yaml.Unmarshal(file,config)
	if err != nil {
		return nil, err
	}
	return config, nil
}