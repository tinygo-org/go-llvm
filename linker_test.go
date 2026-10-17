package llvm

import "testing"

func TestLinker(t *testing.T) {
	ctx := NewContext()
	defer ctx.Dispose()
	dest := ctx.NewModule("dest")
	defer dest.Dispose()

	linker := NewLinker(dest)
	for _, name := range []string{"a", "b", "c"} {
		src := ctx.NewModule(name)
		fn := AddFunction(src, name, FunctionType(ctx.VoidType(), nil, false))
		b := ctx.NewBuilder()
		b.SetInsertPointAtEnd(ctx.AddBasicBlock(fn, "entry"))
		b.CreateRetVoid()
		b.Dispose()
		if err := linker.LinkModule(src); err != nil {
			t.Fatalf("linking %s: %v", name, err)
		}
	}
	linker.Dispose()

	for _, name := range []string{"a", "b", "c"} {
		if fn := dest.NamedFunction(name); fn.IsNil() || fn.IsDeclaration() {
			t.Errorf("function %s not linked in", name)
		}
	}
	if err := VerifyModule(dest, ReturnStatusAction); err != nil {
		t.Errorf("module does not verify: %v", err)
	}
}
