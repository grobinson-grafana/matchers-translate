package matchers_translate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsEquivalent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		err      string
	}{{
		name:     "no parens",
		input:    "",
		expected: true,
	}, {
		name:     "equals",
		input:    "foo=bar",
		expected: true,
	}, {
		name:     "equals unicode emoji",
		input:    "{foo=\"🙂\"}",
		expected: true,
	}, {
		name:     "equals without quotes",
		input:    "{foo=bar}",
		expected: true,
	}, {
		name:     "equals without parens",
		input:    "foo=\"bar\"",
		expected: true,
	}, {
		name:     "equals without parens or quotes",
		input:    "foo=bar",
		expected: true,
	}, {
		name:     "equals with trailing comma",
		input:    "{foo=\"bar\",}",
		expected: true,
	}, {
		name:     "equals without parens but trailing comma",
		input:    "foo=\"bar\",",
		expected: true,
	}, {
		name:     "equals with newline",
		input:    "{foo=\"bar\\n\"}",
		expected: true,
	}, {
		name:     "equals with tab",
		input:    "{foo=\"bar\\t\"}",
		expected: true,
	}, {
		name:     "equals with escaped quotes",
		input:    "{foo=\"\\\"bar\\\"\"}",
		expected: true,
	}, {
		name:     "equals with escaped backslash",
		input:    "{foo=\"bar\\\\\"}",
		expected: true,
	}, {
		name:     "not equals",
		input:    "foo!=bar",
		expected: true,
	}, {
		name:     "match regex",
		input:    "{foo=~\"[a-z]+\"}",
		expected: true,
	}, {
		name:  "match regex without quotes",
		input: "{foo=~[a-z]+}",
		err:   "failed to parse matchers using new parser: 6:7: [: invalid input: expected label value",
	}, {
		name:     "doesn't match regex",
		input:    "{foo!~\"[a-z]+\"}",
		expected: true,
	}, {
		name:  "doesn't match regex without quotes",
		input: "{foo!~[a-z]+}",
		err:   "failed to parse matchers using new parser: 6:7: [: invalid input: expected label value",
	}, {
		name:     "complex",
		input:    "{foo=\"bar\",bar!=\"baz\"}",
		expected: true,
	}, {
		name:     "complex without quotes",
		input:    "{foo=bar,bar!=baz}",
		expected: true,
	}, {
		name:     "complex without parens",
		input:    "foo=\"bar\",bar!=\"baz\"",
		expected: true,
	}, {
		name:     "complex without parens or quotes",
		input:    "foo=bar,bar!=baz",
		expected: true,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := IsEquivalent(test.input)
			if test.err == "" {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, actual)
			} else {
				assert.Equal(t, false, actual)
				assert.EqualError(t, err, test.err)
			}
		})
	}
}

func TestTranslate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		err      string
	}{{
		name:     "no parens",
		input:    "",
		expected: "",
	}, {
		name:     "equals",
		input:    "foo=bar",
		expected: "foo=bar",
	}, {
		name:     "equals unicode emoji",
		input:    "{foo=\"🙂\"}",
		expected: "{foo=\"🙂\"}",
	}, {
		name:     "equals without quotes",
		input:    "{foo=bar}",
		expected: "{foo=bar}",
	}, {
		name:     "equals without parens",
		input:    "foo=\"bar\"",
		expected: "foo=\"bar\"",
	}, {
		name:     "equals without parens or quotes",
		input:    "foo=bar",
		expected: "foo=bar",
	}, {
		name:     "equals with trailing comma",
		input:    "{foo=\"bar\",}",
		expected: "{foo=\"bar\",}",
	}, {
		name:     "equals without parens but trailing comma",
		input:    "foo=\"bar\",",
		expected: "foo=\"bar\",",
	}, {
		name:     "equals with newline",
		input:    "{foo=\"bar\\n\"}",
		expected: "{foo=\"bar\\n\"}",
	}, {
		name:     "equals with tab",
		input:    "{foo=\"bar\\t\"}",
		expected: "{foo=\"bar\\t\"}",
	}, {
		name:     "not equals",
		input:    "foo!=bar",
		expected: "foo!=bar",
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := Translate(test.input)
			if test.err == "" {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, actual)
			} else {
				assert.Equal(t, "", actual)
				assert.EqualError(t, err, test.err)
			}
		})
	}
}
