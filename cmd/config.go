package cmd 

import (
    "io/ioutil"
    "log"

    "gopkg.in/yaml.v2"
)

// Config to use
type Config struct {
    DbUser       string `yaml:"user"`
    DbPassword   string `yaml:"password"`
    DbSchema     string `yaml:"schema"`
}

// GetConfig of user
func (c *Config) GetConfig() *Config {
    yamlFile, err := ioutil.ReadFile("config.yml")
    if err != nil {
        log.Printf("Error in %v ", err)
    }   
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }   
    return c
}