
#include "backports.h"
#include "llvm/IR/Instructions.h"
#include "llvm/Transforms/IPO/PassManagerBuilder.h"
#include "llvm/Transforms/Coroutines.h"
#include "llvm-c/DebugInfo.h"

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

void LLVMGlobalObjectAddMetadata(LLVMValueRef Global, unsigned KindID, LLVMMetadataRef MD) {
  llvm::MDNode *N = MD ? llvm::unwrap<llvm::MDNode>(MD) : nullptr;
  llvm::GlobalObject *O = llvm::unwrap<llvm::GlobalObject>(Global);
  O->addMetadata(KindID, *N);
}

LLVMMetadataRef
LLVMGoDIBuilderCreateTypedef(LLVMDIBuilderRef Builder, LLVMMetadataRef Type,
                             const char *Name, size_t NameLen,
                             LLVMMetadataRef File, unsigned LineNo,
                             LLVMMetadataRef Scope, uint32_t AlignInBits) {
#if LLVM_VERSION_MAJOR >= 10
	return LLVMDIBuilderCreateTypedef(Builder, Type, Name, NameLen, File, LineNo, Scope, AlignInBits);
#else
	return LLVMDIBuilderCreateTypedef(Builder, Type, Name, NameLen, File, LineNo, Scope);
#endif
}
