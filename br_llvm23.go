//go:build llvm23

package llvm

/*
#include "llvm-c/Core.h"
*/
import "C"

// LLVM 23 splits Br into CondBr and UncondBr. Br aliases UncondBr and
// normalizeOpcode folds CondBr back into it, matching pre-23 behavior.
const (
	Br     Opcode = C.LLVMUncondBr
	CondBr Opcode = C.LLVMCondBr
)

func normalizeOpcode(op Opcode) Opcode {
	if op == CondBr {
		return Br
	}
	return op
}

func (v Value) IsACondBrInst() (rv Value)   { rv.C = C.LLVMIsACondBrInst(v.C); return }
func (v Value) IsAUncondBrInst() (rv Value) { rv.C = C.LLVMIsAUncondBrInst(v.C); return }

// IsABranchInst matches both branch kinds.
func (v Value) IsABranchInst() Value {
	if rv := v.IsACondBrInst(); !rv.IsNil() {
		return rv
	}
	return v.IsAUncondBrInst()
}
