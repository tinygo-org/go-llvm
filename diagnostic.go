//===- diagnostic.go - Bindings for diagnostic handlers -------------------===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// This file defines bindings for LLVMContextSetDiagnosticHandler.
//
//===----------------------------------------------------------------------===//

package llvm

/*
#include "llvm-c/Core.h"
#include "diagnostic.h"
*/
import "C"

import (
	"runtime/cgo"
	"sync"
)

type DiagnosticSeverity C.LLVMDiagnosticSeverity

const (
	DSError   DiagnosticSeverity = C.LLVMDSError
	DSWarning DiagnosticSeverity = C.LLVMDSWarning
	DSRemark  DiagnosticSeverity = C.LLVMDSRemark
	DSNote    DiagnosticSeverity = C.LLVMDSNote
)

// DiagnosticInfo is a diagnostic that LLVM passes to a diagnostic handler.
// It is only valid during the call to the handler.
type DiagnosticInfo struct {
	C C.LLVMDiagnosticInfoRef
}

// Severity returns the severity of the diagnostic.
func (d DiagnosticInfo) Severity() DiagnosticSeverity {
	return DiagnosticSeverity(C.LLVMGetDiagInfoSeverity(d.C))
}

// Description returns the text of the diagnostic.
func (d DiagnosticInfo) Description() string {
	cmsg := C.LLVMGetDiagInfoDescription(d.C)
	defer C.LLVMDisposeMessage(cmsg)
	return C.GoString(cmsg)
}

// The handle of the diagnostic handler that is installed on each context.
var (
	diagnosticHandlersLock sync.Mutex
	diagnosticHandlers     = map[C.LLVMContextRef]cgo.Handle{}
)

// SetDiagnosticHandler passes every diagnostic that LLVM reports on c, such as
// linker warnings, to fn instead of printing it. A nil fn restores the default
// handler, which prints to stderr. LLVM can call fn from any goroutine that
// uses c.
func (c Context) SetDiagnosticHandler(fn func(DiagnosticInfo)) {
	diagnosticHandlersLock.Lock()
	defer diagnosticHandlersLock.Unlock()
	if old, ok := diagnosticHandlers[c.C]; ok {
		delete(diagnosticHandlers, c.C)
		defer old.Delete()
	}
	if fn == nil {
		C.LLVMContextSetDiagnosticHandler(c.C, nil, nil)
		return
	}
	h := cgo.NewHandle(fn)
	diagnosticHandlers[c.C] = h
	C.LLVMGoSetDiagnosticHandler(c.C, C.uintptr_t(h))
}

//export llvmGoDiagnosticHandler
func llvmGoDiagnosticHandler(info C.LLVMDiagnosticInfoRef, handle C.uintptr_t) {
	fn := cgo.Handle(handle).Value().(func(DiagnosticInfo))
	fn(DiagnosticInfo{info})
}
