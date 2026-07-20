//===- deprecated.c - Wrappers for deprecated LLVM C API --------*- C -*-===//
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

#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "deprecated.h"

LLVMContextRef LLVMGetGlobalContext_wrap(void) {
  return LLVMGetGlobalContext();
}

#pragma GCC diagnostic pop
