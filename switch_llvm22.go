//go:build llvm22

package llvm

/*
#include "llvm-c/Core.h"
*/
import "C"

// GetSwitchCaseValue obtains the case value for a successor of a switch
// instruction. i corresponds to the successor index; the first successor (0)
// is the default destination, so i must be greater than zero.
//
// LLVM 22 stopped exposing switch case values as regular instruction
// operands (only the condition and destination-block operands remain), so
// this is implemented via the new LLVMGetSwitchCaseValue C API added in the
// same release.
func (v Value) GetSwitchCaseValue(i int) (rv Value) {
	rv.C = C.LLVMGetSwitchCaseValue(v.C, C.unsigned(i))
	return
}
