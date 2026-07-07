package llvm

import (
	"os"
	"testing"
)

// TestSwitchCaseValue checks that GetSwitchCaseValue/SuccessorsCount/Successor
// correctly read a switch instruction's cases across LLVM versions. LLVM 22
// stopped exposing case values as regular instruction operands (only the
// condition and destination-block operands remain), instead requiring the
// new LLVMGetSwitchCaseValue API; code that assumed the old operand layout
// silently reads a destination block where it expects a case value.
func TestSwitchCaseValue(t *testing.T) {
	src := `
define void @foo(i64 %callback) {
entry:
  switch i64 %callback, label %default [
    i64 0, label %case0
    i64 5, label %case1
  ]
default:
  ret void
case0:
  ret void
case1:
  ret void
}
`
	f, err := os.CreateTemp("", "switchcase-*.ll")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	if _, err := f.WriteString(src); err != nil {
		t.Fatal(err)
	}
	f.Close()

	ctx := NewContext()
	defer ctx.Dispose()

	buf, err := NewMemoryBufferFromFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	m, err := ctx.ParseIR(buf)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Dispose()

	fn := m.NamedFunction("foo")
	sw := fn.EntryBasicBlock().FirstInstruction()

	if n := sw.SuccessorsCount(); n != 3 {
		t.Fatalf("expected 3 successors (default + 2 cases), got %d", n)
	}
	if got := sw.Successor(0).AsValue().Name(); got != "default" {
		t.Errorf("expected default destination %q, got %q", "default", got)
	}

	wantCaseValues := []uint64{0, 5}
	wantCaseDests := []string{"case0", "case1"}
	for i, want := range wantCaseValues {
		successor := i + 1
		val := sw.GetSwitchCaseValue(successor)
		if got := val.ZExtValue(); got != want {
			t.Errorf("case %d: expected value %d, got %d", i, want, got)
		}
		if got := sw.Successor(successor).AsValue().Name(); got != wantCaseDests[i] {
			t.Errorf("case %d: expected destination %q, got %q", i, wantCaseDests[i], got)
		}
	}
}
