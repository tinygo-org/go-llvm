package llvm

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestTypeStringToken(t *testing.T) {
	ctx := NewContext()
	defer ctx.Dispose()

	if got := ctx.TokenType().String(); got != "TokenType" {
		t.Errorf("got %q, want %q", got, "TokenType")
	}
}

func TestNonPointerTypeString(t *testing.T) {
	ctx := NewContext()
	defer ctx.Dispose()

	named := ctx.StructCreateNamed("pair")
	named.StructSetBody([]Type{ctx.Int8Type(), ctx.Int32Type()}, false)
	opaque := ctx.StructCreateNamed("opaque")
	tests := []struct {
		name string
		typ  Type
		want string
	}{
		{"void", ctx.VoidType(), "VoidType"},
		{"label", ctx.LabelType(), "LabelType"},
		{"token", ctx.TokenType(), "TokenType"},
		{"int1", ctx.Int1Type(), "IntegerType(1 bits)"},
		{"int8", ctx.Int8Type(), "IntegerType(8 bits)"},
		{"int16", ctx.Int16Type(), "IntegerType(16 bits)"},
		{"int32", ctx.Int32Type(), "IntegerType(32 bits)"},
		{"int64", ctx.Int64Type(), "IntegerType(64 bits)"},
		{"int7", ctx.IntType(7), "IntegerType(7 bits)"},
		{"int128", ctx.IntType(128), "IntegerType(128 bits)"},
		{"float", ctx.FloatType(), "FloatType"},
		{"double", ctx.DoubleType(), "DoubleType"},
		{"x86-fp80", ctx.X86FP80Type(), "X86_FP80Type"},
		{"fp128", ctx.FP128Type(), "FP128Type"},
		{"ppc-fp128", ctx.PPCFP128Type(), "PPC_FP128Type"},
		{"array-empty", ArrayType(ctx.Int8Type(), 0), "ArrayType(IntegerType(8 bits)[0])"},
		{"array", ArrayType(ctx.Int16Type(), 3), "ArrayType(IntegerType(16 bits)[3])"},
		{"array-nested", ArrayType(ArrayType(ctx.Int8Type(), 2), 3), "ArrayType(ArrayType(IntegerType(8 bits)[2])[3])"},
		{"vector", VectorType(ctx.Int32Type(), 4), "VectorType"},
		{"vector-float", VectorType(ctx.FloatType(), 2), "VectorType"},
		{"struct-empty", ctx.StructType(nil, false), "StructType()"},
		{"struct-single", ctx.StructType([]Type{ctx.Int8Type()}, false), "StructType(IntegerType(8 bits))"},
		{"struct-packed", ctx.StructType([]Type{ctx.Int8Type(), ctx.FloatType()}, true), "StructType(IntegerType(8 bits), FloatType)"},
		{"struct-named", named, "%pair: StructType(IntegerType(8 bits), IntegerType(32 bits))"},
		{"struct-opaque", opaque, "%opaque: StructType()"},
		{"struct-repeated", ctx.StructType([]Type{named, named}, false), "StructType(%pair: StructType(IntegerType(8 bits), IntegerType(32 bits)), %pair: StructType(IntegerType(8 bits), IntegerType(32 bits)))"},
		{"struct-nested", ctx.StructType([]Type{ArrayType(ctx.Int8Type(), 2), ctx.StructType(nil, false)}, false), "StructType(ArrayType(IntegerType(8 bits)[2]), StructType())"},
		{"function-empty", FunctionType(ctx.VoidType(), nil, false), "FunctionType():VoidType"},
		{"function-single", FunctionType(ctx.Int32Type(), []Type{ctx.FloatType()}, false), "FunctionType(FloatType):IntegerType(32 bits)"},
		{"function-multiple", FunctionType(ctx.DoubleType(), []Type{ctx.Int8Type(), ctx.Int16Type(), ctx.FloatType()}, false), "FunctionType(IntegerType(8 bits), IntegerType(16 bits), FloatType):DoubleType"},
		{"function-vararg-empty", FunctionType(ctx.VoidType(), nil, true), "FunctionType():VoidType"},
		{"function-vararg", FunctionType(ctx.VoidType(), []Type{ctx.Int32Type()}, true), "FunctionType(IntegerType(32 bits)):VoidType"},
		{"function-aggregate", FunctionType(named, []Type{named, named}, false), "FunctionType(%pair: StructType(IntegerType(8 bits), IntegerType(32 bits)), %pair: StructType(IntegerType(8 bits), IntegerType(32 bits))):%pair: StructType(IntegerType(8 bits), IntegerType(32 bits))"},
		{"global-context-struct", StructType([]Type{GlobalContext().Int8Type()}, false), "StructType(IntegerType(8 bits))"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 2; i++ {
				if got := tt.typ.String(); got != tt.want {
					t.Errorf("call %d: got %q, want %q", i+1, got, tt.want)
				}
			}
		})
	}
}

func TestTypeStringParsed(t *testing.T) {
	major, err := strconv.Atoi(strings.SplitN(Version, ".", 2)[0])
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name       string
		ir         string
		want       string
		minVersion int
		maxVersion int
	}{
		{"half", "half", "HalfType", 15, 0},
		{"bfloat", "bfloat", "BFloatType", 15, 0},
		{"scalable-vector", "<vscale x 4 x i32>", "ScalableVectorType", 15, 0},
		{"x86-amx", "x86_amx", "X86_AMXType", 15, 0},
		{"target-ext", `target("test")`, "TargetExtType", 16, 0},
		{"byte", "b8", "ByteType", 23, 0},
		{"x86-mmx", "x86_mmx", "X86_MMXType", 15, 19},
		{"metadata", "metadata", "MetadataType", 15, 0},
		{"half-array", "[3 x half]", "ArrayType(HalfType[3])", 15, 0},
		{"bfloat-struct", "{bfloat, i32}", "StructType(BFloatType, IntegerType(32 bits))", 15, 0},
		{"scalable-vector-struct", "{<vscale x 4 x i32>}", "StructType(ScalableVectorType)", 15, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if major < tt.minVersion {
				t.Skipf("requires LLVM %d or later", tt.minVersion)
			}
			if tt.maxVersion != 0 && major > tt.maxVersion {
				t.Skipf("removed after LLVM %d", tt.maxVersion)
			}
			path := filepath.Join(t.TempDir(), "type.ll")
			if err := os.WriteFile(path, []byte("declare void @f("+tt.ir+")\n"), 0600); err != nil {
				t.Fatal(err)
			}
			ctx := NewContext()
			defer ctx.Dispose()
			buf, err := NewMemoryBufferFromFile(path)
			if err != nil {
				t.Fatal(err)
			}
			mod, err := ctx.ParseIR(buf)
			if err != nil {
				t.Fatal(err)
			}
			defer mod.Dispose()

			typ := mod.NamedFunction("f").GlobalValueType().ParamTypes()[0]
			if got := typ.String(); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
