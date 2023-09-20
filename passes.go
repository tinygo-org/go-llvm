package llvm

/*
#include "llvm-c/Transforms/PassBuilder.h"
#include "llvm-c/Error.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

// PassBuilderOptions allows specifying several options for the PassBuilder.
type PassBuilderOptions struct {
	C C.LLVMPassBuilderOptionsRef
}

// NewPassBuilderOptions creates a PassBuilderOptions which can be used
// to specify pass options for RunPasses.
func NewPassBuilderOptions() (pbo PassBuilderOptions) {
	pbo.C = C.LLVMCreatePassBuilderOptions()
	return
}

// RunPasses runs the specified optimization passes on the functions in the module.
// `passes` is a comma separated list of pass names in the same format as llvm's
// `opt -passes=...` command. Running `opt -print-passes` can list the available
// passes.
//
// Some notable passes include:
//
//	default<O0>  -- run the default -O0 passes
//	default<O1>  -- run the default -O1 passes
//	default<O2>  -- run the default -O2 passes
//	default<O3>  -- run the default -O3 passes
//	default<Os>  -- run the default -Os passes, like -O2 but size conscious
//	default<Oz>  -- run the default -Oz passes, optimizing for size above all else
func (mod Module) RunPasses(passes string, tm TargetMachine, options PassBuilderOptions) error {
	cpasses := C.CString(passes)
	defer C.free(unsafe.Pointer(cpasses))

	err := C.LLVMRunPasses(mod.C, cpasses, tm.C, options.C)
	if err != nil {
		cstr := C.LLVMGetErrorMessage(err)
		gstr := C.GoString(cstr)
		C.LLVMDisposeErrorMessage(cstr)

		return errors.New(gstr)
	}
	return nil
}

// SetVerifyEach toggles adding a VerifierPass to the PassBuilder,
// ensuring all functions inside the module are valid. Useful for
// debugging, but adds a significant amount of overhead.
func (pbo PassBuilderOptions) SetVerifyEach(verifyEach bool) {
	C.LLVMPassBuilderOptionsSetVerifyEach(pbo.C, boolToLLVMBool(verifyEach))
}

// SetDebugLogging toggles debug logging for the PassBuilder.
func (pbo PassBuilderOptions) SetDebugLogging(debugLogging bool) {
	C.LLVMPassBuilderOptionsSetDebugLogging(pbo.C, boolToLLVMBool(debugLogging))
}

// SetLoopInterleaving toggles loop interleaving, which is part of
// loop vectorization.
func (pbo PassBuilderOptions) SetLoopInterleaving(loopInterleaving bool) {
	C.LLVMPassBuilderOptionsSetLoopInterleaving(pbo.C, boolToLLVMBool(loopInterleaving))
}

// SetLoopVectorization toggles loop vectorization.
func (pbo PassBuilderOptions) SetLoopVectorization(loopVectorization bool) {
	C.LLVMPassBuilderOptionsSetLoopVectorization(pbo.C, boolToLLVMBool(loopVectorization))
}

// SetSLPVectorization toggles Super-Word Level Parallelism vectorization,
// whose goal is to combine multiple similar independent instructions into
// a vector instruction.
func (pbo PassBuilderOptions) SetSLPVectorization(slpVectorization bool) {
	C.LLVMPassBuilderOptionsSetSLPVectorization(pbo.C, boolToLLVMBool(slpVectorization))
}

// SetLoopUnrolling toggles loop unrolling.
func (pbo PassBuilderOptions) SetLoopUnrolling(loopUnrolling bool) {
	C.LLVMPassBuilderOptionsSetLoopUnrolling(pbo.C, boolToLLVMBool(loopUnrolling))
}

// SetForgetAllSCEVInLoopUnroll toggles forgetting all SCEV (Scalar Evolution)
// information in loop unrolling. Scalar Evolution is a pass that analyses
// the how scalars evolve over iterations of a loop in order to optimize
// the loop better. Forgetting this information can be useful in some cases.
func (pbo PassBuilderOptions) SetForgetAllSCEVInLoopUnroll(forgetSCEV bool) {
	C.LLVMPassBuilderOptionsSetForgetAllSCEVInLoopUnroll(pbo.C, boolToLLVMBool(forgetSCEV))
}

// SetLicmMssaOptCap sets a tuning option to cap the number of calls to
// retrieve clobbering accesses in MemorySSA, in Loop Invariant Code Motion
// optimization.
// See [llvm::PipelineTuningOptions::LicmMssaOptCap].
func (pbo PassBuilderOptions) SetLicmMssaOptCap(optCap uint) {
	C.LLVMPassBuilderOptionsSetLicmMssaOptCap(pbo.C, C.unsigned(optCap))
}

// SetLicmMssaNoAccForPromotionCap sets a tuning option to cap the number of
// promotions to scalars in Loop Invariant Code Motion with MemorySSA, if
// the number of accesses is too large.
// See [llvm::PipelineTuningOptions::LicmMssaNoAccForPromotionCap].
func (pbo PassBuilderOptions) SetLicmMssaNoAccForPromotionCap(promotionCap uint) {
	C.LLVMPassBuilderOptionsSetLicmMssaNoAccForPromotionCap(pbo.C, C.unsigned(promotionCap))
}

// SetCallGraphProfile toggles whether call graph profiling should be used.
func (pbo PassBuilderOptions) SetCallGraphProfile(cgProfile bool) {
	C.LLVMPassBuilderOptionsSetCallGraphProfile(pbo.C, boolToLLVMBool(cgProfile))
}

// SetMergeFunctions toggles finding functions which will generate identical
// machine code by considering all pointer types to be equivalent. Once
// identified, they will be folded by replacing a call to one with a call to a
// bitcast of the other.
func (pbo PassBuilderOptions) SetMergeFunctions(mergeFuncs bool) {
	C.LLVMPassBuilderOptionsSetMergeFunctions(pbo.C, boolToLLVMBool(mergeFuncs))
}

// Dispose of the memory allocated for the PassBuilderOptions.
func (pbo PassBuilderOptions) Dispose() {
	C.LLVMDisposePassBuilderOptions(pbo.C)
}
