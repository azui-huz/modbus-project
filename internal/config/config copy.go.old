package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		UnitID int    `yaml:"unit_id"`

		HoldingRegisters struct {
			Size int `yaml:"size"`
		} `yaml:"holding_registers"`

		InputRegisters struct {
			Size int `yaml:"size"`
		} `yaml:"input_registers"`

		Coils struct {
			Size int `yaml:"size"`
		} `yaml:"coils"`

		DiscreteInputs struct {
			Size int `yaml:"size"`
		} `yaml:"discrete_inputs"`
	} `yaml:"server"`

	API struct {
		Port int `yaml:"port"`
	} `yaml:"api"`

	CyclicReaders []struct {
		Name       string `yaml:"name"`
		IntervalMS int    `yaml:"interval_ms"`
		Registers  []struct {
			Address int `yaml:"address"`
			Length  int `yaml:"length"`
		} `yaml:"registers"`
	} `yaml:"cyclic_readers"`

	Clients []struct {
		Name   string `yaml:"name"`
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		UnitID int    `yaml:"unit_id"`
	} `yaml:"clients"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
