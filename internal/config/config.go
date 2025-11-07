package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Host             string `yaml:"host"`
	Port             int    `yaml:"port"`
	UnitID           int    `yaml:"unit_id"`
	HoldingRegisters struct {
		Size          int      `yaml:"size"`
		InitialValues []uint16 `yaml:"initial_values"`
	} `yaml:"holding_registers"`
	Coils struct {
		Size int `yaml:"size"`
	} `yaml:"coils"`
}

type CyclicReaderCfg struct {
	Name       string `yaml:"name"`
	IntervalMS int    `yaml:"interval_ms"`
	Registers  []struct {
		Address int `yaml:"address"`
		Length  int `yaml:"length"`
	} `yaml:"registers"`
}

type ClientCfg struct {
	Name   string `yaml:"name"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	UnitID int    `yaml:"unit_id"`
}

type Config struct {
	Server        ServerConfig      `yaml:"server"`
	CyclicReaders []CyclicReaderCfg `yaml:"cyclic_readers"`
	Clients       []ClientCfg       `yaml:"clients"`
	API           APIConfig         `yaml:"api"`
}

func Load(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

type APIConfig struct {
	Port int `yaml:"port"`
}
