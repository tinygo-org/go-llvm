package llvm

// This file contains functions backported from LLVM master.

// #include "llvm-c/Core.h"
// #include <stdlib.h>
// #include "backports.h"
import "C"

// Coroutine optimization passes
// https://reviews.llvm.org/D51642 (in progress)

func (pmb PassManagerBuilder) AddCoroutinePassesToExtensionPoints() {
	C.LLVMPassManagerBuilderAddCoroutinePassesToExtensionPoints_backport(pmb.C)
}

// Erase instruction
// https://reviews.llvm.org/D52694 (in progress)

func (v Value) EraseFromParentAsInstruction()  { C.LLVMInstructionEraseFromParent(v.C) }
func (v Value) RemoveFromParentAsInstruction() { C.LLVMInstructionRemoveFromParent(v.C) }

// Called function from a CallInst
// https://reviews.llvm.org/D52972 (in progress)

func (v Value) CalledValue() (rv Value) {
	rv.C = C.LLVMGetCalledValue(v.C)
	return
}

// Predicates
// https://reviews.llvm.org/D53884 (in progress)

func (v Value) IntPredicate() IntPredicate {
	return IntPredicate(C.LLVMGetICmpPredicate(v.C))
}

func (v Value) FloatPredicate() FloatPredicate {
	return FloatPredicate(C.LLVMGetFCmpPredicate(v.C))
}
