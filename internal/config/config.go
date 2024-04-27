package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const pathToConfig = "../../config.yaml"

// Config is a struct to store configuration to test library
type Config struct {
	AntiCaptcha struct {
		Token string `yaml:"token"`
	} `yaml:"anti_captcha"`
}

// FromFile reads config from file
func FromFile() Config {
	file, err := os.Open(pathToConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var conf Config
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err = decoder.Decode(&conf); err != nil {
			log.Fatal(err.Error())
		}
	}

	return conf
}
