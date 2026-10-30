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

	majorVersion, _ := strconv.Atoi(strings.SplitN(Version, ".", 2)[0])
	if majorVersion < 19 {
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

// TestCreateConstantRangeAttributeShortSlice checks that slices too short
// for numBits panic instead of letting LLVM read past them. It also checks
// that calculating the word count does not overflow. The check is shared,
// so it applies to every LLVM version.
func TestCreateConstantRangeAttributeShortSlice(t *testing.T) {
	ctx := NewContext()
	defer ctx.Dispose()

	for _, test := range []struct {
		name    string
		numBits uint
	}{
		{name: "128 bits", numBits: 128},
		{name: "maximum uint", numBits: ^uint(0)},
	} {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Errorf("expected panic for %d bit range with one word slices", test.numBits)
				}
			}()
			ctx.CreateConstantRangeAttribute(AttributeKindID("range"), test.numBits, []uint64{0}, []uint64{100})
		})
	}
}
