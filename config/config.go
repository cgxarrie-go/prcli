package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

var cfg *Config

// GetInstance returns singleton instance.
func GetInstance() *Config {
	if cfg == nil {
		lock.Lock()
		defer lock.Unlock()
		if cfg == nil {
			cfg = &Config{}
		}
	}

	return cfg
}

var (
	configFileName string = ".prcliconfig"
)

// Config main configuration for CLI
type Config struct {
	Azure AzureConfig
}

// AzureConfig configuration for Azure
type AzureConfig struct {
	CompanyName string
	PAT         string
	Projects    []AzureProjectConfig
}

// NewConfig creates a new instance of Config
func NewConfig() Config {
	return Config{
		Azure: AzureConfig{},
	}
}

// AzureProjectConfig configuration for Azure Projects
type AzureProjectConfig struct {
	ID            string
	RepositoryIDs []string
}

// Load read config file into config struct
func (c *Config) Load() (err error) {

	if _, err := os.Stat(configFileName); err != nil {
		cfg := NewConfig()
		err = cfg.save()
		if err != nil {
			return err
		}
	}

	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return fmt.Errorf("Error loading config file: %s", err.Error())

	}

	err = json.Unmarshal(file, c)
	if err != nil {
		return fmt.Errorf("Error parsing config: %s", err.Error())
	}

	return nil
}

func (c *Config) save() (err error) {
	f, err := os.Create(configFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(b))

	return err
}
