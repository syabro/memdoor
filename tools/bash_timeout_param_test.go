package tools

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

// A call's timeout is its own up to the maximum; no timeout is the default.
// The config file sets both, and the model reads them in the description.
func TestABashCallAsksForItsOwnTimeoutUpToTheMaximum(t *testing.T) {
	oldLimit, oldMax, oldDesc := BashTimeoutLimit, BashTimeoutMax, BashDefinition.Description
	defer func() { BashTimeoutLimit, BashTimeoutMax, BashDefinition.Description = oldLimit, oldMax, oldDesc }()

	SetBashTimeouts(600*time.Second, 1800*time.Second)
	cases := []struct {
		input string
		want  time.Duration
	}{
		{`{"command":"go test ./..."}`, 600 * time.Second},
		{`{"command":"rejudge","timeout":1200}`, 1200 * time.Second},
		{`{"command":"rejudge","timeout":99999}`, 1800 * time.Second},
		{`{"command":"ls","timeout":-5}`, 600 * time.Second},
		{`{"command":"ls","timeout":"long"}`, 600 * time.Second},
	}
	for _, c := range cases {
		if got := ParseBashTimeout(json.RawMessage(c.input)); got != c.want {
			t.Errorf("%s: got %s, want %s", c.input, got, c.want)
		}
	}
	if d := BashDefinition.Description; !strings.Contains(d, "600 seconds") || !strings.Contains(d, "at most 1800") {
		t.Errorf("the description must name the default and the maximum: %q", d)
	}

	SetBashTimeouts(0, 60*time.Second)
	if BashTimeoutLimit != 600*time.Second || BashTimeoutMax != 600*time.Second {
		t.Errorf("0 keeps the default, and the maximum is never below it: %s, %s", BashTimeoutLimit, BashTimeoutMax)
	}
}
