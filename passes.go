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
//	default<Os>  -- run the default -Os passes, like -O2 but size concious
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
// ensuring all functions inside the module are valid.
func (pbo PassBuilderOptions) SetVerifyEach(verifyEach bool) {
	C.LLVMPassBuilderOptionsSetVerifyEach(pbo.C, boolToLLVMBool(verifyEach))
}

func (pbo PassBuilderOptions) SetDebugLogging(debugLogging bool) {
	C.LLVMPassBuilderOptionsSetDebugLogging(pbo.C, boolToLLVMBool(debugLogging))
}

func (pbo PassBuilderOptions) SetLoopInterleaving(loopInterleaving bool) {
	C.LLVMPassBuilderOptionsSetLoopInterleaving(pbo.C, boolToLLVMBool(loopInterleaving))
}

func (pbo PassBuilderOptions) SetLoopVectorization(loopVectorization bool) {
	C.LLVMPassBuilderOptionsSetLoopVectorization(pbo.C, boolToLLVMBool(loopVectorization))
}

func (pbo PassBuilderOptions) SetSLPVectorization(slpVectorization bool) {
	C.LLVMPassBuilderOptionsSetSLPVectorization(pbo.C, boolToLLVMBool(slpVectorization))
}

func (pbo PassBuilderOptions) SetLoopUnrolling(loopUnrolling bool) {
	C.LLVMPassBuilderOptionsSetLoopUnrolling(pbo.C, boolToLLVMBool(loopUnrolling))
}

func (pbo PassBuilderOptions) SetForgetAllSCEVInLoopUnroll(forgetSCEV bool) {
	C.LLVMPassBuilderOptionsSetForgetAllSCEVInLoopUnroll(pbo.C, boolToLLVMBool(forgetSCEV))
}

func (pbo PassBuilderOptions) SetLicmMssaOptCap(optCap uint) {
	C.LLVMPassBuilderOptionsSetLicmMssaOptCap(pbo.C, C.unsigned(optCap))
}

func (pbo PassBuilderOptions) SetLicmMssaNoAccForPromotionCap(promotionCap uint) {
	C.LLVMPassBuilderOptionsSetLicmMssaNoAccForPromotionCap(pbo.C, C.unsigned(promotionCap))
}

func (pbo PassBuilderOptions) SetCallGraphProfile(cgProfile bool) {
	C.LLVMPassBuilderOptionsSetCallGraphProfile(pbo.C, boolToLLVMBool(cgProfile))
}

func (pbo PassBuilderOptions) SetMergeFunctions(mergeFuncs bool) {
	C.LLVMPassBuilderOptionsSetMergeFunctions(pbo.C, boolToLLVMBool(mergeFuncs))
}

func (pbo PassBuilderOptions) Dispose() {
	C.LLVMDisposePassBuilderOptions(pbo.C)
}
