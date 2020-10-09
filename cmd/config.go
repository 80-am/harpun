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
    Hook         bool   `yaml:"hook"`
    LogPath      string `yaml:"log"`
    Multiplier   int64  `yaml:"multiplier"`
}

// GetConfig of user
func (c *Config) GetConfig(config string) *Config {
    if config == "" {
        config = "config.yml"
    }
    yamlFile, err := ioutil.ReadFile(config)
    if err != nil {
        log.Printf("Error in %v ", err)
    }   
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }   
    return c
}