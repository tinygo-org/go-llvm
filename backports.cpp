
#include "backports.h"
#include "llvm/IR/Instructions.h"
#include "llvm/IR/LegacyPassManager.h"
#include "llvm/IR/Module.h"
#include "llvm/Pass.h"
#include "llvm/Support/MemoryBuffer.h"
#include "llvm/Transforms/IPO.h"
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

// See https://reviews.llvm.org/D119431
LLVMMemoryBufferRef LLVMGoWriteThinLTOBitcodeToMemoryBuffer(LLVMModuleRef M) {
  std::string Data;
  llvm::raw_string_ostream OS(Data);
  llvm::legacy::PassManager PM;
  PM.add(createWriteThinLTOBitcodePass(OS));
  PM.run(*llvm::unwrap(M));
  return llvm::wrap(llvm::MemoryBuffer::getMemBufferCopy(OS.str()).release());
}

LLVMMetadataRef LLVMGoDIBuilderCreateExpression(LLVMDIBuilderRef Builder,
                                                uint64_t *Addr, size_t Length) {
#if LLVM_VERSION_MAJOR >= 14
  return LLVMDIBuilderCreateExpression(Builder, Addr, Length);
#else
  return LLVMDIBuilderCreateExpression(Builder, (int64_t*)Addr, Length);
#endif
}
