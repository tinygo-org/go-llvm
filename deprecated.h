//===- deprecated.h - Wrappers for deprecated LLVM C API --------*- C -*-===//
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

#ifndef LLVM_BINDINGS_GO_LLVM_DEPRECATED_H
#define LLVM_BINDINGS_GO_LLVM_DEPRECATED_H

#include "llvm-c/Core.h"

#ifdef __cplusplus
extern "C" {
#endif

LLVMContextRef LLVMGetGlobalContext_wrap(void);
LLVMValueRef LLVMConstGEP2_wrap(LLVMTypeRef Ty, LLVMValueRef ConstantVal,
                                LLVMValueRef *ConstantIndices,
                                unsigned NumIndices);
LLVMValueRef LLVMConstInBoundsGEP2_wrap(LLVMTypeRef Ty,
                                        LLVMValueRef ConstantVal,
                                        LLVMValueRef *ConstantIndices,
                                        unsigned NumIndices);
LLVMBuilderRef LLVMCreateBuilderInContext_wrap(LLVMContextRef C);

#ifdef __cplusplus
}
#endif

#endif // LLVM_BINDINGS_GO_LLVM_DEPRECATED_H
