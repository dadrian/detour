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

var parseTests []*configTest

func TestParseConfig(t *testing.T) {
	for _, test := range parseTests {
		r := strings.NewReader(test.config)
		d, err := ParseConfig(r)
		if test.err != nil {
			assert.EqualError(t, err, test.err.Error())
		} else {
			assert.Equal(t, test.expected, d)
		}
	}
}
