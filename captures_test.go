package llvm

import (
	"strconv"
	"strings"
	"testing"
)

// TestCapturesAttribute checks that the 'captures' parameter attribute
// (which replaced the boolean 'nocapture' enum attribute starting with
// LLVM 21) round-trips through the generic enum-attribute API, and that a
// value of 0 corresponds to CaptureInfo::none(), i.e. captures(none).
func TestCapturesAttribute(t *testing.T) {
	majorVersion, _ := strconv.Atoi(strings.SplitN(Version, ".", 2)[0])
	if majorVersion < 21 {
		t.Skip("not llvm 21")
	}

	ctx := NewContext()
	mod := ctx.NewModule("")
	defer mod.Dispose()

	ptrType := PointerType(ctx.Int8Type(), 0)
	ftyp := FunctionType(ctx.VoidType(), []Type{ptrType}, false)
	fn := AddFunction(mod, "foo", ftyp)

	kind := AttributeKindID("captures")
	if kind == 0 {
		t.Fatal("captures kind id not found")
	}

	attr := ctx.CreateEnumAttribute(kind, 0)
	fn.AddAttributeAtIndex(1, attr)

	got := fn.GetEnumAttributeAtIndex(1, kind)
	if got.IsNil() {
		t.Fatal("expected captures attribute on param 1, got nil")
	}
	if val := got.GetEnumValue(); val != 0 {
		t.Errorf("expected captures value 0 (none), got %d", val)
	}

	text := mod.String()
	if !strings.Contains(text, "captures(none)") {
		t.Errorf("expected 'captures(none)' in output, got:\n%s", text)
	}
}
