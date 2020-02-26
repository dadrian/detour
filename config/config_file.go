package config

import (
	"errors"
	"io"

	"gopkg.in/yaml.v2"
)

var ErrNoBrowsers = errors.New("no browsers specified in config")
var ErrUnknownBrowser = errors.New("unknown browser")

type WindowDefinition struct {
}

type BrowserDefinition struct {
	Name    string `yaml:"name"`
	Browser string `yaml:"browser"`
	Profile string `yaml:"profile"`
}

// CheckValidity performs baseline validity checks for missing or conflicting
// values in a BrowserDefinition.
func (bd *BrowserDefinition) CheckValidity() error {
	if (bd.Browser != "") && (bd.Browser != "firefox") {
		return ErrUnknownBrowser
	}
	return nil
}

type Definition struct {
	Window   WindowDefinition    `yaml:"window"`
	Browsers []BrowserDefinition `yaml:"browsers"`
}

func (d *Definition) CheckValidity() error {
	if len(d.Browsers) == 0 {
		return ErrNoBrowsers
	}
	return nil
}

func ParseConfig(r io.Reader) (*Definition, error) {
	d := yaml.NewDecoder(r)
	definition := Definition{}
	if err := d.Decode(&definition); err != nil {
		return nil, err
	}
	return &definition, nil
}
