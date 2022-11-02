package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func Load(filepath string) (*Config, error) {
	c := &Config{}
	configfile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configfile, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
