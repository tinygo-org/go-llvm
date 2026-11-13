package llvm

import (
	"strings"
	"testing"
)

// Linking modules with different target triples makes LLVM report a warning.
func linkMismatchedTriples(t *testing.T, ctx Context) {
	t.Helper()
	dst := ctx.NewModule("dst")
	defer dst.Dispose()
	dst.SetTarget("x86_64-unknown-linux-gnu")
	src := ctx.NewModule("src")
	src.SetTarget("aarch64-unknown-linux-gnu")
	if err := LinkModules(dst, src); err != nil {
		t.Fatal(err)
	}
}

func TestDiagnosticHandler(t *testing.T) {
	ctx := NewContext()
	defer ctx.Dispose()

	var severities []DiagnosticSeverity
	var messages []string
	ctx.SetDiagnosticHandler(func(d DiagnosticInfo) {
		severities = append(severities, d.Severity())
		messages = append(messages, d.Description())
	})
	linkMismatchedTriples(t, ctx)

	if len(messages) != 1 {
		t.Fatalf("got %d diagnostics, want 1: %q", len(messages), messages)
	}
	if !strings.Contains(messages[0], "different target triples") {
		t.Errorf("diagnostic = %q, want a target triple warning", messages[0])
	}
	if severities[0] != DSWarning {
		t.Errorf("severity = %d, want DSWarning", severities[0])
	}
}

func TestDiagnosticHandlerPerContext(t *testing.T) {
	ctx := NewContext()
	defer ctx.Dispose()
	other := NewContext()
	defer other.Dispose()

	calls := 0
	ctx.SetDiagnosticHandler(func(DiagnosticInfo) { calls++ })
	other.SetDiagnosticHandler(func(DiagnosticInfo) {})
	linkMismatchedTriples(t, other)
	if calls != 0 {
		t.Errorf("handler got %d diagnostics from another context", calls)
	}

	ctx.SetDiagnosticHandler(nil)
	if _, ok := diagnosticHandlers[ctx.C]; ok {
		t.Error("handle kept after the handler was removed")
	}
}
