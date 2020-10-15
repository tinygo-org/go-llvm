
#include "backports.h"
#include "llvm/IR/Instructions.h"
#include "llvm/Transforms/IPO/PassManagerBuilder.h"
#include "llvm/Transforms/Coroutines.h"
#include "llvm-c/DebugInfo.h"

namespace llvm {

#if LLVM_VERSION_MAJOR < 11
inline PassManagerBuilder *unwrap(LLVMPassManagerBuilderRef P) {
  return reinterpret_cast<PassManagerBuilder*>(P);
}

inline LLVMPassManagerBuilderRef wrap(PassManagerBuilder *P) {
  return reinterpret_cast<LLVMPassManagerBuilderRef>(P);
}
#endif

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

LLVMMetadataRef LLVMGoDIBuilderCreateCompileUnit(
    LLVMDIBuilderRef Builder, LLVMDWARFSourceLanguage Lang,
    LLVMMetadataRef FileRef, const char *Producer, size_t ProducerLen,
    LLVMBool isOptimized, const char *Flags, size_t FlagsLen,
    unsigned RuntimeVer, const char *SplitName, size_t SplitNameLen,
    LLVMDWARFEmissionKind Kind, unsigned DWOId, LLVMBool SplitDebugInlining,
    LLVMBool DebugInfoForProfiling, const char *SysRoot, size_t SysRootLen,
    const char *SDK, size_t SDKLen) {

#if LLVM_VERSION_MAJOR >= 11
  return LLVMDIBuilderCreateCompileUnit(Builder, Lang, FileRef, Producer,
    ProducerLen, isOptimized, Flags, FlagsLen, RuntimeVer, SplitName,
    SplitNameLen, Kind, DWOId, SplitDebugInlining, DebugInfoForProfiling,
    SysRoot, SysRootLen, SDK, SDKLen);
#else
  return LLVMDIBuilderCreateCompileUnit(Builder, Lang, FileRef, Producer,
    ProducerLen, isOptimized, Flags, FlagsLen, RuntimeVer, SplitName,
    SplitNameLen, Kind, DWOId, SplitDebugInlining, DebugInfoForProfiling);
#endif
}
