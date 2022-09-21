
#include "backports.h"
#include "llvm/IR/Instructions.h"
#include "llvm/IR/LegacyPassManager.h"
#include "llvm/IR/Module.h"
#include "llvm/Pass.h"
#include "llvm/Support/MemoryBuffer.h"
#include "llvm/Transforms/IPO.h"
#include "llvm/Transforms/IPO/PassManagerBuilder.h"
#include "llvm-c/DebugInfo.h"

void LLVMGlobalObjectAddMetadata(LLVMValueRef Global, unsigned KindID, LLVMMetadataRef MD) {
  llvm::MDNode *N = MD ? llvm::unwrap<llvm::MDNode>(MD) : nullptr;
  llvm::GlobalObject *O = llvm::unwrap<llvm::GlobalObject>(Global);
  O->addMetadata(KindID, *N);
}

// See https://reviews.llvm.org/D119431
LLVMMemoryBufferRef LLVMGoWriteThinLTOBitcodeToMemoryBuffer(LLVMModuleRef M) {
  std::string Data;
  llvm::raw_string_ostream OS(Data);
  llvm::legacy::PassManager PM;
  PM.add(createWriteThinLTOBitcodePass(OS));
  PM.run(*llvm::unwrap(M));
  return llvm::wrap(llvm::MemoryBuffer::getMemBufferCopy(OS.str()).release());
}
