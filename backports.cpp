
#include "backports.h"
#include "llvm/Transforms/IPO/PassManagerBuilder.h"
#include "llvm/Transforms/Coroutines.h"

namespace llvm {

inline PassManagerBuilder *unwrap(LLVMPassManagerBuilderRef P) {
  return reinterpret_cast<PassManagerBuilder*>(P);
}

inline LLVMPassManagerBuilderRef wrap(PassManagerBuilder *P) {
  return reinterpret_cast<LLVMPassManagerBuilderRef>(P);
}

} // end namespace llvm

void LLVMPassManagerBuilderAddCoroutinePassesToExtensionPoints_backport(LLVMPassManagerBuilderRef PMB) {
  llvm::PassManagerBuilder *Builder = llvm::unwrap(PMB);
  llvm::addCoroutinePassesToExtensionPoints(*Builder);
}
