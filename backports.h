
#include "llvm-c/DebugInfo.h"
#include "llvm-c/Types.h"
#include "llvm-c/Transforms/PassManagerBuilder.h"

#ifdef __cplusplus
extern "C" {
#endif

void LLVMPassManagerBuilderAddCoroutinePassesToExtensionPoints_backport(LLVMPassManagerBuilderRef PMB);

void LLVMGlobalObjectAddMetadata(LLVMValueRef objValue, unsigned KindID, LLVMMetadataRef md);

LLVMValueRef LLVMGoGetInlineAsm(LLVMTypeRef Ty, char *AsmString,
                               size_t AsmStringSize, char *Constraints,
                               size_t ConstraintsSize, LLVMBool HasSideEffects,
                               LLVMBool IsAlignStack,
                               LLVMInlineAsmDialect Dialect, LLVMBool CanThrow);

#ifdef __cplusplus
}
#endif /* defined(__cplusplus) */
