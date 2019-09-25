//===- InstrumentationBindings.cpp - instrumentation bindings -------------===//
//
//                     The LLVM Compiler Infrastructure
//
// This file is distributed under the University of Illinois Open Source
// License. See LICENSE.TXT for details.
//
//===----------------------------------------------------------------------===//
//
// This file defines C bindings for the instrumentation component.
//
//===----------------------------------------------------------------------===//

#include "InstrumentationBindings.h"
#include "llvm-c/Core.h"
#include "llvm/IR/LegacyPassManager.h"
#include "llvm/IR/Module.h"
#include "llvm/Transforms/Instrumentation.h"
#if LLVM_VERSION_MAJOR >= 9
#include "llvm/Transforms/Instrumentation/AddressSanitizer.h"
#endif
#include "llvm/Transforms/Instrumentation/MemorySanitizer.h"
#include "llvm/Transforms/Instrumentation/ThreadSanitizer.h"

using namespace llvm;

void LLVMAddAddressSanitizerFunctionPass(LLVMPassManagerRef PM) {
  unwrap(PM)->add(createAddressSanitizerFunctionPass());
}

void LLVMAddAddressSanitizerModulePass(LLVMPassManagerRef PM) {
#if LLVM_VERSION_MAJOR >= 9
  unwrap(PM)->add(createModuleAddressSanitizerLegacyPassPass());
#else
  unwrap(PM)->add(createAddressSanitizerModulePass());
#endif
}

void LLVMAddThreadSanitizerPass(LLVMPassManagerRef PM) {
  unwrap(PM)->add(createThreadSanitizerLegacyPassPass());
}

void LLVMAddMemorySanitizerLegacyPassPass(LLVMPassManagerRef PM) {
  unwrap(PM)->add(createMemorySanitizerLegacyPassPass());
}

void LLVMAddDataFlowSanitizerPass(LLVMPassManagerRef PM,
                                  int ABIListFilesNum,
                                  const char **ABIListFiles) {
  std::vector<std::string> ABIListFilesVec;
  for (int i = 0; i != ABIListFilesNum; ++i) {
    ABIListFilesVec.push_back(ABIListFiles[i]);
  }
  unwrap(PM)->add(createDataFlowSanitizerPass(ABIListFilesVec));
}
