package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Server *Server `yaml:"server"`
}

// Read YAML config file and return a Config struct
func ReadConfig(path string) (c *Config, err error) {
	yamlFile, err := os.Open(path)
	if err != nil {
		panic("error reading YAML config file")
	}
	defer yamlFile.Close()

	decode := yaml.NewDecoder(yamlFile)
	if err := decode.Decode(&c); err != nil {
		log.Printf("error Decode YAML file: %v", err)
	}
	return c, nil
}
