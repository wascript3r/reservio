package registry

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	ConfigENV = "RESERVIO_CONFIG"
)

var (
	ErrConfigNotProvided = errors.New("config file is not provided")
)

type Duration struct {
	time.Duration
}

func (d *Duration) MarhsalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	pd, err := time.ParseDuration(str)
	if err != nil {
		return err
	}

	d.Duration = pd
	return nil
}

type Config struct {
	Log struct {
		ShowTimestamp bool `json:"showTimestamp"`
	} `json:"log"`

	Database struct {
		Postgres struct {
			DSN          string   `json:"dsn"`
			QueryTimeout Duration `json:"queryTimeout"`
		} `json:"postgres"`
	} `json:"database"`

	Auth struct {
		PasswordCost int `json:"passwordCost"`
		JWT          struct {
			SecretKey  string   `json:"secretKey"`
			Expiration Duration `json:"expiration"`
			Issuer     string   `json:"issuer"`
		} `json:"jwt"`
	} `json:"auth"`

	HTTP struct {
		Port        string `json:"port"`
		EnablePprof bool   `json:"enablePprof"`
	} `json:"http"`
}

func getConfigPath() (string, error) {
	path := os.Getenv(ConfigENV)
	path = strings.TrimSpace(path)

	if len(path) == 0 {
		return "", ErrConfigNotProvided
	}

	return path, nil
}

func parseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = json.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadConfig() (*Config, error) {
	// Get working directory
	workDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	return parseConfig(filepath.Join(workDir, path))
}
