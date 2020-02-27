package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type configTest struct {
	config   string
	err      error
	expected Definition
}

var parseTests []*configTest = []*configTest{
	{
		config: `
window:
  width: 200
  height: 200
browsers:
  - name: "user defined name"
    browser: "firefox"
    profile: "default"
goop: "blah"
`,
		expected: Definition{
			Window: WindowDefinition{
				Width:  200,
				Height: 200,
			},
			Browsers: []BrowserDefinition{
				{
					Name:    "user defined name",
					Browser: "firefox",
					Profile: "default",
				},
			},
		},
	},
}

func TestParseConfig(t *testing.T) {
	for _, test := range parseTests {
		t.Log(test.config)
		r := strings.NewReader(test.config)
		d, err := ParseConfig(r)
		if test.err != nil {
			assert.NotNil(t, err)
			assert.EqualError(t, err, test.err.Error())
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, d)
			assert.Equal(t, test.expected, *d)
		}
	}
}

func TestValidateValid(t *testing.T) {
	valids := []Definition{
		{
			Browsers: []BrowserDefinition{
				{
					Name:    "n",
					Browser: "firefox",
					Profile: "x",
				},
			},
		},
	}
	for _, d := range valids {
		err := d.CheckValidity()
		assert.NoError(t, err)
	}
}

func TestValidateInvalid(t *testing.T) {
	invalids := []Definition{
		{
			Browsers: []BrowserDefinition{},
		},
		{
			Browsers: []BrowserDefinition{
				{
					Browser: "not a browser",
				},
			},
		},
		{
			Browsers: []BrowserDefinition{
				{
					Browser: "firefox",
					Name:    "name",
				},
			},
		},
		{
			Browsers: []BrowserDefinition{
				{
					Browser: "firefox",
					Profile: "profile",
				},
			},
		},
	}
	for _, d := range invalids {
		err := d.CheckValidity()
		assert.NotNil(t, err, d)
	}
}
