
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
#ifdef __cplusplus
}
#endif /* defined(__cplusplus) */
