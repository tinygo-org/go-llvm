package llvm

import (
	"strconv"
	"strings"
	"testing"
)

// TestCreateConstantRangeAttribute checks that a range parameter attribute
// can be built and attached on LLVM 19 and later, and that earlier versions
// return a nil attribute.
func TestCreateConstantRangeAttribute(t *testing.T) {
	ctx := NewContext()
	defer ctx.Dispose()
	mod := ctx.NewModule("")
	defer mod.Dispose()

	ftyp := FunctionType(ctx.VoidType(), []Type{ctx.Int32Type()}, false)
	fn := AddFunction(mod, "foo", ftyp)

	kind := AttributeKindID("range")
	attr := ctx.CreateConstantRangeAttribute(kind, 32, []uint64{0}, []uint64{100})

	if VersionMajor < 19 {
		if !attr.IsNil() {
			t.Errorf("got attribute %p, want nil before LLVM 19", attr.C)
		}
		return
	}

	if kind == 0 {
		t.Fatal("range kind id not found")
	}
	if attr.IsNil() {
		t.Fatal("got nil range attribute")
	}

	fn.AddAttributeAtIndex(1, attr)
	if got := fn.GetEnumAttributeAtIndex(1, kind); got.IsNil() {
		t.Error("expected range attribute on param 1, got nil")
	}

	text := mod.String()
	if !strings.Contains(text, "range(i32 0, 100)") {
		t.Errorf("expected 'range(i32 0, 100)' in output, got:\n%s", text)
	}
}

// TestCreateConstantRangeAttributeBadInput checks that bad arguments panic
// before they get to LLVM. The checks apply to every LLVM version.
func TestCreateConstantRangeAttributeBadInput(t *testing.T) {
	ctx := NewContext()
	defer ctx.Dispose()

	kind := AttributeKindID("range")
	one := []uint64{0}
	wide := uint64(1) << 32
	over32 := uint(wide)
	for _, test := range []struct {
		name         string
		numBits      uint
		lower, upper []uint64
		needs64      bool
	}{
		{name: "zero bits", numBits: 0, lower: one, upper: one},
		{name: "zero bits empty slices", numBits: 0},
		{name: "short slices", numBits: 128, lower: one, upper: one},
		{name: "maximum uint bits", numBits: ^uint(0), lower: one, upper: one},
		{name: "bits over 32 bits", numBits: over32, lower: one, upper: one, needs64: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.needs64 && strconv.IntSize == 32 {
				t.Skip("uint is 32 bits")
			}
			defer func() {
				if recover() == nil {
					t.Error("expected panic")
				}
			}()
			ctx.CreateConstantRangeAttribute(kind, test.numBits, test.lower, test.upper)
		})
	}
}
