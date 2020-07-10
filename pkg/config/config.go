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

// Workhours
type Workhours struct {
	Monday    int
	Tuesday   int
	Wednesday int
	Thursday  int
	Friday    int
	Saturday  int
	Sunday    int
}

// ReportConfig
type ReportConfig struct {
	Workhours Workhours
}

// Filter for report
type Filter struct {
	Name  string
	Tags  []string
	Tasks []TaskFilter
}

// TaskFilter for report
type TaskFilter struct {
	Project string
	Task    string
}

// Config contains the complete service configuration
type Config struct {
	APIConfig     APIConfig
	ReportConfig  ReportConfig
	ReportFilters map[string]Filter
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

	if config.ReportFilters == nil {
		config.ReportFilters = map[string]Filter{}
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
