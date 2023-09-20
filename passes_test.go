package llvm

import "testing"

func TestPasses(t *testing.T) {
	InitializeNativeTarget()
	InitializeNativeAsmPrinter()

	ctx := NewContext()

	mod := ctx.NewModule("fac_module")

	fac_args := []Type{ctx.Int32Type()}
	fac_type := FunctionType(ctx.Int32Type(), fac_args, false)
	fac := AddFunction(mod, "fac", fac_type)
	fac.SetFunctionCallConv(CCallConv)
	n := fac.Param(0)

	entry := AddBasicBlock(fac, "entry")
	iftrue := AddBasicBlock(fac, "iftrue")
	iffalse := AddBasicBlock(fac, "iffalse")
	end := AddBasicBlock(fac, "end")

	builder := ctx.NewBuilder()
	defer builder.Dispose()

	builder.SetInsertPointAtEnd(entry)
	If := builder.CreateICmp(IntEQ, n, ConstInt(ctx.Int32Type(), 0, false), "cmptmp")
	builder.CreateCondBr(If, iftrue, iffalse)

	builder.SetInsertPointAtEnd(iftrue)
	res_iftrue := ConstInt(ctx.Int32Type(), 1, false)
	builder.CreateBr(end)

	builder.SetInsertPointAtEnd(iffalse)
	n_minus := builder.CreateSub(n, ConstInt(ctx.Int32Type(), 1, false), "subtmp")
	call_fac_args := []Value{n_minus}
	call_fac := builder.CreateCall(fac_type, fac, call_fac_args, "calltmp")
	res_iffalse := builder.CreateMul(n, call_fac, "multmp")
	builder.CreateBr(end)

	builder.SetInsertPointAtEnd(end)
	res := builder.CreatePHI(ctx.Int32Type(), "result")
	phi_vals := []Value{res_iftrue, res_iffalse}
	phi_blocks := []BasicBlock{iftrue, iffalse}
	res.AddIncoming(phi_vals, phi_blocks)
	builder.CreateRet(res)

	err := VerifyModule(mod, ReturnStatusAction)
	if err != nil {
		t.Errorf("Error verifying module: %s", err)
		return
	}

	targ, err := GetTargetFromTriple(DefaultTargetTriple())
	if err != nil {
		t.Error(err)
	}

	mt := targ.CreateTargetMachine(DefaultTargetTriple(), "", "", CodeGenLevelDefault, RelocDefault, CodeModelDefault)

	pbo := NewPassBuilderOptions()
	defer pbo.Dispose()

	t.Run("no error running default pass", func(t *testing.T) {
		err := mod.RunPasses("default<Os>", mt, pbo)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("errors on unknown pass name", func(t *testing.T) {
		err := mod.RunPasses("badpassnamedoesnotexist", mt, pbo)
		if err == nil {
			t.Error("expecting error but got none")
		}

		if err.Error() != "unknown pass name 'badpassnamedoesnotexist'" {
			t.Errorf("expected error about unknow pass name, instead got %s", err)
		}
	})
}
