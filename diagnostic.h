//===- diagnostic.h - Diagnostic handler shim -------------------*- C -*-===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//

#ifndef LLVM_BINDINGS_GO_LLVM_DIAGNOSTIC_H
#define LLVM_BINDINGS_GO_LLVM_DIAGNOSTIC_H

#include "llvm-c/Core.h"
#include <stdint.h>

// LLVMGoSetDiagnosticHandler installs the Go diagnostic handler on c, with
// handle as its runtime/cgo.Handle.
void LLVMGoSetDiagnosticHandler(LLVMContextRef c, uintptr_t handle);

#endif
