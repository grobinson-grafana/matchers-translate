package matchers_translate

import (
	"fmt"
	"reflect"

	"github.com/grobinson-grafana/matchers"
	"github.com/prometheus/alertmanager/pkg/labels"
)

// IsEquivalent returns true if the matcher(s) can be parsed with both parsers,
// and the parsed matcher(s) are equivalent.
func IsEquivalent(s string) (bool, error) {
	a, err := labels.ParseMatchers(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse matchers using old parser: %w", err)
	}
	b, err := matchers.Parse(s)
	if err != nil {
		return false, fmt.Errorf("failed to parse matchers using new parser: %w", err)
	}
	if len(a) != len(b) {
		return false, nil
	}
	for i := 0; i < len(a); i++ {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false, nil
		}
	}
	return true, nil
}

// Translate will translate matchers that can be parsed with the old parser but not
// the new parser into a matcher that can be parsed with both. If however the matcher
// can be parsed with both parsers, and the parsed matcher(s) are equivalent,
// then no translation is required
func Translate(s string) (string, error) {
	if ok, _ := IsEquivalent(s); ok {
		return s, nil
	}
	m, err := labels.ParseMatchers(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse matchers using old parser: %s", err)
	}
	if len(m) == 0 {
		return "", nil
	}
	ts := labels.Matchers(m).String()
	if ok, err := IsEquivalent(ts); err != nil || !ok {
		return ts, fmt.Errorf("translated matcher(s) are not equivalent in both parsers: %w", err)
	}
	return ts, nil
}
