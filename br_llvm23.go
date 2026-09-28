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
