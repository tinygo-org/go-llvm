
#include "backports.h"
#include "llvm/IR/Instructions.h"
#include "llvm/Transforms/IPO/PassManagerBuilder.h"
#include "llvm/Transforms/Coroutines.h"
#include "llvm-c/DebugInfo.h"

void LLVMPassManagerBuilderAddCoroutinePassesToExtensionPoints_backport(LLVMPassManagerBuilderRef PMB) {
  llvm::PassManagerBuilder *Builder = llvm::unwrap(PMB);
  llvm::addCoroutinePassesToExtensionPoints(*Builder);
}

void LLVMGlobalObjectAddMetadata(LLVMValueRef Global, unsigned KindID, LLVMMetadataRef MD) {
  llvm::MDNode *N = MD ? llvm::unwrap<llvm::MDNode>(MD) : nullptr;
  llvm::GlobalObject *O = llvm::unwrap<llvm::GlobalObject>(Global);
  O->addMetadata(KindID, *N);
}

LLVMValueRef LLVMGoGetInlineAsm(LLVMTypeRef Ty, char *AsmString,
                               size_t AsmStringSize, char *Constraints,
                               size_t ConstraintsSize, LLVMBool HasSideEffects,
                               LLVMBool IsAlignStack,
                               LLVMInlineAsmDialect Dialect, LLVMBool CanThrow) {
#if LLVM_VERSION_MAJOR >= 13
  return LLVMGetInlineAsm(Ty, AsmString, AsmStringSize, Constraints, ConstraintsSize, HasSideEffects, IsAlignStack, Dialect, CanThrow);
#else
  return LLVMGetInlineAsm(Ty, AsmString, AsmStringSize, Constraints, ConstraintsSize, HasSideEffects, IsAlignStack, Dialect);
#endif
}
