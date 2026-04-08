package core

import "testing"

func TestInitAcceptsEmptyConfig(t *testing.T) {
	resp := Init("")
	if resp.Code != CodeOK {
		t.Fatalf("expected CodeOK, got %d", resp.Code)
	}
}

func TestInitRejectsBadJSON(t *testing.T) {
	resp := Init("{bad-json")
	if resp.Code != CodeInvalidInput {
		t.Fatalf("expected CodeInvalidInput, got %d", resp.Code)
	}
}

func TestExecuteRequiresAction(t *testing.T) {
	resp := Execute(`{"trace_id":"t-1"}`)
	if resp.Code != CodeInvalidInput {
		t.Fatalf("expected CodeInvalidInput, got %d", resp.Code)
	}
}

func TestExecuteAcceptsAction(t *testing.T) {
	resp := Execute(`{"action":"ping","trace_id":"t-2"}`)
	if resp.Code != CodeOK {
		t.Fatalf("expected CodeOK, got %d", resp.Code)
	}
	if resp.TraceID != "t-2" {
		t.Fatalf("expected trace_id t-2, got %s", resp.TraceID)
	}
}

func TestShutdown(t *testing.T) {
	resp := Shutdown()
	if resp.Code != CodeOK {
		t.Fatalf("expected CodeOK, got %d", resp.Code)
	}
}
