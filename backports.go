package llvm

// This file contains functions backported from LLVM master.

// #include "llvm-c/Core.h"
// #include <stdlib.h>
// #include "backports.h"
import "C"
import "unsafe"

// Token type (used by coroutines)
// https://reviews.llvm.org/D47684

const TokenTypeKind TypeKind = C.LLVMTokenTypeKind

func (c Context) TokenType() (t Type) { t.C = C.LLVMTokenTypeInContext(c.C); return }

// Inline assembly
// https://reviews.llvm.org/D46437

type InlineAsmDialect C.LLVMInlineAsmDialect

const (
	InlineAsmDialectATT   InlineAsmDialect = C.LLVMInlineAsmDialectATT
	InlineAsmDialectIntel InlineAsmDialect = C.LLVMInlineAsmDialectIntel
)

func InlineAsm(t Type, asmString, constraints string, hasSideEffects, isAlignStack bool, dialect InlineAsmDialect) (rv Value) {
	casm := C.CString(asmString)
	defer C.free(unsafe.Pointer(casm))
	cconstraints := C.CString(constraints)
	defer C.free(unsafe.Pointer(cconstraints))
	rv.C = C.LLVMGetInlineAsm(t.C, casm, C.size_t(len(asmString)), cconstraints, C.size_t(len(constraints)), boolToLLVMBool(hasSideEffects), boolToLLVMBool(isAlignStack), C.LLVMInlineAsmDialect(dialect))
	return
}

// Coroutine optimization passes
// https://reviews.llvm.org/D51642 (in progress)

func (pmb PassManagerBuilder) AddCoroutinePassesToExtensionPoints() {
	C.LLVMPassManagerBuilderAddCoroutinePassesToExtensionPoints_backport(pmb.C);
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

// Indices
// https://reviews.llvm.org/D53883 (in progress)

func (v Value) Indices() []uint32 {
	num := C.LLVMGetNumIndices(v.C)
	indicesPtr := C.LLVMGetIndices(v.C)
	// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
	rawIndices := (*[1 << 30]C.uint)(unsafe.Pointer(indicesPtr))[:num:num]
	indices := make([]uint32, num)
	for i := range indices {
		indices[i] = uint32(rawIndices[i])
	}
	return indices
}

// Predicates
// https://reviews.llvm.org/D53884 (in progress)

func (v Value) IntPredicate() IntPredicate {
	return IntPredicate(C.LLVMGetICmpPredicate(v.C))
}

func (v Value) FloatPredicate() FloatPredicate {
	return FloatPredicate(C.LLVMGetFCmpPredicate(v.C))
}
