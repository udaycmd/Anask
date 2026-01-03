// Copyright 2026 Uday Tiwari. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"os"
)

type httpPort uint16

type meta struct {
	Name string `json:"name"`
	Desc string `json:"description"`
	Ver  string `json:"version"`
}

type prioryConfig struct {
	Meta meta     `json:"meta"`
	Port httpPort `json:"http_port"`
}

type postgresConfig struct {
	Meta meta   `json:"meta"`
	Conn string `json:"connection_string"`
}

type saltConfig struct {
	Generator string `json:"generator"`
	Value     string `json:"value"`
}

type Config struct {
	Mode     string `json:"mode"`
	Priory   prioryConfig
	Postgres postgresConfig
	Salt     saltConfig `json:"secret_salt"`
}

func Init(configPath string) (*Config, error) {
	if configPath == "" {
		configPath = "./config/config.json"
	}

	return loadConfig(configPath)
}

func loadConfig(configPath string) (*Config, error) {
	jsonFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := json.Unmarshal(jsonFile, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
