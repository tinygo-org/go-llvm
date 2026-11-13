//===- diagnostic.c - Diagnostic handler shim -------------------*- C -*-===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// A Go file with //export can only declare C functions, so the handler shim
// that converts the handle to a pointer lives here.
//
//===----------------------------------------------------------------------===//

#include "diagnostic.h"
#include "_cgo_export.h"

static void diagnosticHandler(LLVMDiagnosticInfoRef info, void *ctx) {
	llvmGoDiagnosticHandler(info, (uintptr_t)ctx);
}

void LLVMGoSetDiagnosticHandler(LLVMContextRef c, uintptr_t handle) {
	LLVMContextSetDiagnosticHandler(c, diagnosticHandler, (void *)handle);
}
