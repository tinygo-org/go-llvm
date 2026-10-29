package llvm

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTypeString(t *testing.T) {
	ctx := NewContext()
	defer ctx.Dispose()

	ptr := PointerType(ctx.Int8Type(), 0)
	namedPair := ctx.StructCreateNamed("pair")
	namedPair.StructSetBody([]Type{ctx.Int8Type(), ctx.Int32Type()}, false)
	opaque := ctx.StructCreateNamed("opaque")
	named := ctx.StructCreateNamed("node")
	named.StructSetBody([]Type{ctx.Int32Type(), PointerType(named, 0)}, false)

	// There is no metadata type constructor in the Go bindings.
	path := filepath.Join(t.TempDir(), "metadata.ll")
	if err := os.WriteFile(path, []byte("declare void @metadata_param(metadata)\n"), 0600); err != nil {
		t.Fatal(err)
	}
	buf, err := NewMemoryBufferFromFile(path)
	if err != nil {
		t.Fatal(err)
	}
	mod, err := ctx.ParseIR(buf)
	if err != nil {
		t.Fatal(err)
	}
	defer mod.Dispose()
	metadataType := mod.NamedFunction("metadata_param").GlobalValueType().ParamTypes()[0]

	tests := []struct {
		name string
		typ  Type
		want string
	}{
		{"pointer", PointerType(ctx.Int8Type(), 0), "PointerType"},
		{"pointer-address-space-1", PointerType(ctx.Int8Type(), 1), "PointerType"},
		{"pointer-address-space-255", PointerType(ctx.Int32Type(), 255), "PointerType"},
		{"pointer-to-pointer", PointerType(ptr, 0), "PointerType"},
		{"pointer-to-function", PointerType(FunctionType(ctx.VoidType(), nil, false), 0), "PointerType"},
		{"pointer-array", ArrayType(ptr, 3), "ArrayType(PointerType[3])"},
		{"pointer-struct", ctx.StructType([]Type{ptr, ctx.Int8Type(), ptr}, false), "StructType(PointerType, IntegerType(8 bits), PointerType)"},
		{"pointer-param", FunctionType(ctx.VoidType(), []Type{ptr}, false), "FunctionType(PointerType):VoidType"},
		{"pointer-return", FunctionType(ptr, nil, false), "FunctionType():PointerType"},
		{"pointer-function", FunctionType(ptr, []Type{ptr, ptr}, false), "FunctionType(PointerType, PointerType):PointerType"},
		{"pointer-vector", VectorType(ptr, 2), "VectorType"},
		{"recursive-named-struct", named, "%node: StructType(IntegerType(32 bits), PointerType)"},
		{"integer-1", ctx.Int1Type(), "IntegerType(1 bits)"},
		{"integer-8", ctx.Int8Type(), "IntegerType(8 bits)"},
		{"integer-16", ctx.Int16Type(), "IntegerType(16 bits)"},
		{"integer-32", ctx.Int32Type(), "IntegerType(32 bits)"},
		{"integer-64", ctx.Int64Type(), "IntegerType(64 bits)"},
		{"integer-7", ctx.IntType(7), "IntegerType(7 bits)"},
		{"integer-128", ctx.IntType(128), "IntegerType(128 bits)"},
		{"float", ctx.FloatType(), "FloatType"},
		{"double", ctx.DoubleType(), "DoubleType"},
		{"x86-fp80", ctx.X86FP80Type(), "X86_FP80Type"},
		{"fp128", ctx.FP128Type(), "FP128Type"},
		{"ppc-fp128", ctx.PPCFP128Type(), "PPC_FP128Type"},
		{"void", ctx.VoidType(), "VoidType"},
		{"label", ctx.LabelType(), "LabelType"},
		{"metadata", metadataType, "MetadataType"},
		{"empty-array", ArrayType(ctx.Int8Type(), 0), "ArrayType(IntegerType(8 bits)[0])"},
		{"array", ArrayType(ctx.Int16Type(), 4), "ArrayType(IntegerType(16 bits)[4])"},
		{"nested-array", ArrayType(ArrayType(ptr, 2), 3), "ArrayType(ArrayType(PointerType[2])[3])"},
		{"empty-struct", ctx.StructType(nil, false), "StructType()"},
		{"literal-struct", ctx.StructType([]Type{ctx.Int8Type(), ctx.Int32Type()}, false), "StructType(IntegerType(8 bits), IntegerType(32 bits))"},
		{"packed-struct", ctx.StructType([]Type{ptr, ctx.Int8Type()}, true), "StructType(PointerType, IntegerType(8 bits))"},
		{"named-struct", namedPair, "%pair: StructType(IntegerType(8 bits), IntegerType(32 bits))"},
		{"opaque-struct", opaque, "%opaque: StructType()"},
		{"nested-struct", ctx.StructType([]Type{ctx.StructType([]Type{ptr}, false), ArrayType(ptr, 2)}, false), "StructType(StructType(PointerType), ArrayType(PointerType[2]))"},
		{"repeated-named-struct", ctx.StructType([]Type{namedPair, namedPair}, false), "StructType(%pair: StructType(IntegerType(8 bits), IntegerType(32 bits)), %pair: StructType(IntegerType(8 bits), IntegerType(32 bits)))"},
		{"vector", VectorType(ctx.Int32Type(), 4), "VectorType"},
		{"float-vector", VectorType(ctx.FloatType(), 2), "VectorType"},
		{"function-no-params", FunctionType(ctx.VoidType(), nil, false), "FunctionType():VoidType"},
		{"function-one-param", FunctionType(ctx.Int32Type(), []Type{ctx.Int8Type()}, false), "FunctionType(IntegerType(8 bits)):IntegerType(32 bits)"},
		{"function-many-params", FunctionType(ctx.DoubleType(), []Type{ctx.Int8Type(), ptr, ctx.FloatType()}, false), "FunctionType(IntegerType(8 bits), PointerType, FloatType):DoubleType"},
		{"variadic-function", FunctionType(ctx.VoidType(), []Type{ptr}, true), "FunctionType(PointerType):VoidType"},
		{"variadic-function-no-params", FunctionType(ctx.VoidType(), nil, true), "FunctionType():VoidType"},
		{"function-struct-return", FunctionType(ctx.StructType([]Type{ptr}, false), nil, false), "FunctionType():StructType(PointerType)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.typ.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
			if got := tt.typ.String(); got != tt.want {
				t.Errorf("second String() = %q, want %q", got, tt.want)
			}
		})
	}
}
