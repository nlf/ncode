package agent

import "testing"

func TestNcodeConfigEnvironmentOverrides(t *testing.T) {
	trueValue := true
	tests := []struct {
		name string
		env  string
		set  string
		got  func() bool
		want bool
	}{
		{name: "flat tools enables", env: "NCODE_FLAT_TOOLS", set: "1", got: func() bool { return Config{}.FlatToolRender() }, want: true},
		{name: "flat tools disables config", env: "NCODE_FLAT_TOOLS", set: "0", got: func() bool { return Config{ToolRender: "flat"}.FlatToolRender() }, want: false},
		{name: "compact input enables", env: "NCODE_COMPACT_INPUT", set: "compact", got: func() bool { return Config{}.CompactUserInput() }, want: true},
		{name: "compact input disables config", env: "NCODE_COMPACT_INPUT", set: "bubble", got: func() bool { return Config{CompactInput: &trueValue}.CompactUserInput() }, want: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(tc.env, tc.set)
			if got := tc.got(); got != tc.want {
				t.Fatalf("environment override = %v, want %v", got, tc.want)
			}
		})
	}
}
