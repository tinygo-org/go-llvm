//go:build !llvm23

package llvm

/*
#include "llvm-c/Core.h"
*/
import "C"

const Br Opcode = C.LLVMBr

func (v Value) IsABranchInst() (rv Value) { rv.C = C.LLVMIsABranchInst(v.C); return }
