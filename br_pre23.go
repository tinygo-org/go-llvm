//go:build !llvm23

package llvm

/*
#include "llvm-c/Core.h"
*/
import "C"

// Pre-23, conditional and unconditional branches share one Br opcode.
// InstructionOpcode returns it unchanged.
const Br Opcode = C.LLVMBr

func normalizeOpcode(op Opcode) Opcode { return op }
