//===- deprecated.h - Wrappers for deprecated LLVM C API --------*- C -*-===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// This file wraps LLVMGetGlobalContext, which was deprecated in LLVM 22 but
// has no non-deprecated replacement. All other wrapped functions are
// handled by calling their context-aware counterparts from Go.
//
//===----------------------------------------------------------------------===//

#ifndef LLVM_BINDINGS_GO_LLVM_DEPRECATED_H
#define LLVM_BINDINGS_GO_LLVM_DEPRECATED_H

#include "llvm-c/Core.h"

#ifdef __cplusplus
extern "C" {
#endif

LLVMContextRef LLVMGetGlobalContext_wrap(void);

#ifdef __cplusplus
}
#endif

#endif // LLVM_BINDINGS_GO_LLVM_DEPRECATED_H
