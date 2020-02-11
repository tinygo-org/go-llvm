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

// Remove instruction
// https://reviews.llvm.org/D72209 (in progress)

func (v Value) RemoveFromParentAsInstruction() { C.LLVMInstructionRemoveFromParent(v.C) }
