
#include "llvm-c/DebugInfo.h"
#include "llvm-c/Types.h"

#ifdef __cplusplus
extern "C" {
#endif

void LLVMGlobalObjectAddMetadata(LLVMValueRef objValue, unsigned KindID, LLVMMetadataRef md);

LLVMMemoryBufferRef LLVMGoWriteThinLTOBitcodeToMemoryBuffer(LLVMModuleRef M);

LLVMMetadataRef LLVMGoDIBuilderCreateExpression(LLVMDIBuilderRef Builder,
                                                uint64_t *Addr, size_t Length);

#ifdef __cplusplus
}
#endif /* defined(__cplusplus) */
