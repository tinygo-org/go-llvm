
#include "llvm-c/Types.h"
#include "llvm-c/Transforms/PassManagerBuilder.h"

#ifdef __cplusplus
extern "C" {
#endif

void LLVMPassManagerBuilderAddCoroutinePassesToExtensionPoints_backport(LLVMPassManagerBuilderRef PMB);

#ifdef __cplusplus
}
#endif /* defined(__cplusplus) */
