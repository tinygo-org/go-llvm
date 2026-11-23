//===- deprecated.c - Wrappers for deprecated LLVM C API --------*- C -*-===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// LLVM 22 deprecates LLVMGetGlobalContext, and LLVM 24 deprecates
// LLVMConstGEP2, LLVMConstInBoundsGEP2 and LLVMCreateBuilderInContext.
//
//===----------------------------------------------------------------------===//

#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "deprecated.h"

LLVMContextRef LLVMGetGlobalContext_wrap(void) {
  return LLVMGetGlobalContext();
}

LLVMValueRef LLVMConstGEP2_wrap(LLVMTypeRef Ty, LLVMValueRef ConstantVal,
                                LLVMValueRef *ConstantIndices,
                                unsigned NumIndices) {
  return LLVMConstGEP2(Ty, ConstantVal, ConstantIndices, NumIndices);
}

LLVMValueRef LLVMConstInBoundsGEP2_wrap(LLVMTypeRef Ty,
                                        LLVMValueRef ConstantVal,
                                        LLVMValueRef *ConstantIndices,
                                        unsigned NumIndices) {
  return LLVMConstInBoundsGEP2(Ty, ConstantVal, ConstantIndices, NumIndices);
}

LLVMBuilderRef LLVMCreateBuilderInContext_wrap(LLVMContextRef C) {
  return LLVMCreateBuilderInContext(C);
}

#pragma GCC diagnostic pop
