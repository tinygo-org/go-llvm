//go:build !llvm15 && !llvm16 && !llvm17 && !llvm18

package llvm

/*
#include "llvm-c/Core.h"
#include <stdint.h>
*/
import "C"

import "unsafe"

func (c Context) createConstantRangeAttribute(kind uint, numBits uint, lower, upper []uint64) (a Attribute) {
	a.C = C.LLVMCreateConstantRangeAttribute(c.C, C.unsigned(kind), C.unsigned(numBits),
		(*C.uint64_t)(unsafe.Pointer(&lower[0])), (*C.uint64_t)(unsafe.Pointer(&upper[0])))
	return
}
