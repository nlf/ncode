package modes

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/nlf/ncode/packages/core"
	"github.com/nlf/ncode/packages/provider"
)

type jsonAuditClient struct {
	err error
}

func (c jsonAuditClient) Name() string { return "json-audit" }

func (c jsonAuditClient) Stream(context.Context, provider.Request) (<-chan provider.Event, error) {
	if c.err != nil {
		return nil, c.err
	}
	events := make(chan provider.Event, 2)
	events <- provider.EventTextDelta{Delta: "retained JSON output"}
	events <- provider.EventDone{
		Stop: provider.StopEnd,
		Message: provider.Message{Role: provider.RoleAssistant, Content: []provider.Content{
			provider.TextBlock{Text: "retained JSON output"},
		}},
	}
	close(events)
	return events, nil
}

func TestRunJSONRetainsNeutralJSONLSuccessAndErrorOutput(t *testing.T) {
	t.Run("success emits neutral event objects", func(t *testing.T) {
		agent := core.NewAgent(jsonAuditClient{}, "audit-model", "", core.Registry{})
		var output bytes.Buffer
		if err := RunJSON(context.Background(), agent, "audit prompt", nil, &output); err != nil {
			t.Fatal(err)
		}

		rows := decodeJSONAuditRows(t, output.String())
		assertJSONAuditRow(t, rows, "text_delta", "delta", "retained JSON output")
		assertJSONAuditRow(t, rows, "assistant_message", "content", []any{
			map[string]any{"type": "text", "text": "retained JSON output"},
		})
	})

	t.Run("provider failure is returned and serialized", func(t *testing.T) {
		providerErr := errors.New("audit provider failure")
		agent := core.NewAgent(jsonAuditClient{err: providerErr}, "audit-model", "", core.Registry{})
		var output bytes.Buffer
		err := RunJSON(context.Background(), agent, "audit prompt", nil, &output)
		if !errors.Is(err, providerErr) {
			t.Fatalf("RunJSON error = %v, want %v", err, providerErr)
		}

		rows := decodeJSONAuditRows(t, output.String())
		assertJSONAuditRow(t, rows, "error", "message", providerErr.Error())
	})
}

func decodeJSONAuditRows(t *testing.T, output string) []map[string]any {
	t.Helper()
	lines := strings.Split(strings.TrimSpace(output), "\n")
	rows := make([]map[string]any, 0, len(lines))
	for _, line := range lines {
		var row map[string]any
		if err := json.Unmarshal([]byte(line), &row); err != nil {
			t.Fatalf("decode JSONL row %q: %v", line, err)
		}
		rows = append(rows, row)
	}
	return rows
}

func assertJSONAuditRow(t *testing.T, rows []map[string]any, eventType, field string, want any) {
	t.Helper()
	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range rows {
		if row["type"] != eventType {
			continue
		}
		gotJSON, err := json.Marshal(row[field])
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Equal(gotJSON, wantJSON) {
			return
		}
	}
	t.Fatalf("missing %q row with %s=%s in %#v", eventType, field, wantJSON, rows)
}
