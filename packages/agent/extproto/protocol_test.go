package extproto

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestHelloAckPublishesExactNcodeProtocolV2Identity(t *testing.T) {
	if ProtocolVersion != 2 {
		t.Fatalf("ProtocolVersion = %d, want 2", ProtocolVersion)
	}
	encoded, err := Encode(HelloAckFromHost{
		Type:            "hello_ack",
		Product:         Product,
		ProtocolVersion: ProtocolVersion,
		NcodeVersion:    "0.4.0-test",
		Provider:        "anthropic",
		Model:           "claude-test",
		CWD:             "/work",
	})
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]any
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	want := map[string]any{
		"type":             "hello_ack",
		"product":          "ncode",
		"protocol_version": float64(2),
		"ncode_version":    "0.4.0-test",
		"provider":         "anthropic",
		"model":            "claude-test",
		"cwd":              "/work",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("hello_ack = %#v, want %#v", got, want)
	}
}

func TestNeutralHelloAndToolFramesRemainUnchanged(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want string
	}{
		{
			name: "hello",
			in: HelloFromExt{
				Type: "hello", Name: "weather", Version: "1.0.0",
				Capabilities: []string{"commands", "tools"},
			},
			want: `{"type":"hello","name":"weather","version":"1.0.0","capabilities":["commands","tools"]}` + "\n",
		},
		{
			name: "tool_call",
			in: ToolCallFromHost{
				Type: "tool_call", ID: "call-1", Name: "weather",
				Args: json.RawMessage(`{"city":"Berlin"}`),
			},
			want: `{"type":"tool_call","id":"call-1","name":"weather","args":{"city":"Berlin"}}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != tt.want {
				t.Fatalf("frame = %q, want %q", got, tt.want)
			}
		})
	}
}
