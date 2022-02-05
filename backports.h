
#include "llvm-c/DebugInfo.h"
#include "llvm-c/Types.h"
#include "llvm-c/Transforms/PassManagerBuilder.h"

#ifdef __cplusplus
extern "C" {
#endif

void LLVMPassManagerBuilderAddCoroutinePassesToExtensionPoints_backport(LLVMPassManagerBuilderRef PMB);

void LLVMGlobalObjectAddMetadata(LLVMValueRef objValue, unsigned KindID, LLVMMetadataRef md);

LLVMMetadataRef
LLVMGoDIBuilderCreateTypedef(LLVMDIBuilderRef Builder, LLVMMetadataRef Type,
                             const char *Name, size_t NameLen,
                             LLVMMetadataRef File, unsigned LineNo,
                             LLVMMetadataRef Scope, uint32_t AlignInBits);

LLVMMetadataRef LLVMGoDIBuilderCreateCompileUnit(
    LLVMDIBuilderRef Builder, LLVMDWARFSourceLanguage Lang,
    LLVMMetadataRef FileRef, const char *Producer, size_t ProducerLen,
    LLVMBool isOptimized, const char *Flags, size_t FlagsLen,
    unsigned RuntimeVer, const char *SplitName, size_t SplitNameLen,
    LLVMDWARFEmissionKind Kind, unsigned DWOId, LLVMBool SplitDebugInlining,
    LLVMBool DebugInfoForProfiling, const char *SysRoot, size_t SysRootLen,
    const char *SDK, size_t SDKLen);

LLVMMemoryBufferRef LLVMGoWriteThinLTOBitcodeToMemoryBuffer(LLVMModuleRef M);

LLVMMetadataRef LLVMGoDIBuilderCreateExpression(LLVMDIBuilderRef Builder,
                                                uint64_t *Addr, size_t Length);

#ifdef __cplusplus
}
#endif /* defined(__cplusplus) */
