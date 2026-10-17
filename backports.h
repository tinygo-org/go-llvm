
#include "llvm-c/DebugInfo.h"
#include "llvm-c/Types.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct LLVMGoOpaqueLinker *LLVMGoLinkerRef;

LLVMGoLinkerRef LLVMGoCreateLinker(LLVMModuleRef Dest);

LLVMBool LLVMGoLinkerLinkModule(LLVMGoLinkerRef L, LLVMModuleRef Src);

void LLVMGoDisposeLinker(LLVMGoLinkerRef L);

void LLVMGlobalObjectAddMetadata(LLVMValueRef objValue, unsigned KindID, LLVMMetadataRef md);

LLVMMemoryBufferRef LLVMGoWriteThinLTOBitcodeToMemoryBuffer(LLVMModuleRef M);

void LLVMGoDIBuilderInsertDbgValueRecordAtEnd(
    LLVMDIBuilderRef Builder, LLVMValueRef Val, LLVMMetadataRef VarInfo,
    LLVMMetadataRef Expr, LLVMMetadataRef DebugLoc, LLVMBasicBlockRef Block);

#ifdef __cplusplus
}
#endif /* defined(__cplusplus) */
