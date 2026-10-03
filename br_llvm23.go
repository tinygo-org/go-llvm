//go:build llvm23

package llvm

/*
#include "llvm-c/Core.h"
*/
import "C"

// LLVM 23 replaced Br with UncondBr and CondBr. Conditional branch operands
// changed from (cond, else, then) to (cond, then, else), use Successor instead.
const (
	UncondBr Opcode = C.LLVMUncondBr
	CondBr   Opcode = C.LLVMCondBr
)

func (v Value) IsACondBrInst() (rv Value)   { rv.C = C.LLVMIsACondBrInst(v.C); return }
func (v Value) IsAUncondBrInst() (rv Value) { rv.C = C.LLVMIsAUncondBrInst(v.C); return }

// IsABranchInst matches both branch kinds.
func (v Value) IsABranchInst() Value {
	if rv := v.IsACondBrInst(); !rv.IsNil() {
		return rv
	}
	return v.IsAUncondBrInst()
}
