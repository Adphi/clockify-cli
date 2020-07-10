package config

import (
	"encoding/json"
	"os"

	env "github.com/Netflix/go-env"
	"github.com/sirupsen/logrus"
)

// APIConfig from clockify api
type APIConfig struct {
	APIEndpoint    string `env:"API_ENDPOINT"`
	ReportEndpoint string `env:"REPORT_ENDPOINT"`
	ClockifyKey    string `env:"CLOCKIFY_ACCESS_KEY"`
}

// Config contains the complete service configuration
type Config struct {
	APIConfig APIConfig
}

// NewTestConfig return a config object with test settings
func NewTestConfig() *Config {
	return &Config{
		APIConfig: APIConfig{
			APIEndpoint:    "https://api.clockify.me/api/v1",
			ReportEndpoint: "https://reports.api.clockify.me/v1",
			ClockifyKey:    "",
		},
	}
}

// ReadConfig reads a json file and overwrite with ENV vars
func ReadConfig(file string, log *logrus.Entry) *Config {
	var config Config

	if fileExists(file) {
		fileContent, _ := os.Open(file)

		if err := json.NewDecoder(fileContent).Decode(&config); err != nil {
			log.Fatal(err)
		}
	}

	// Override ENVs
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
