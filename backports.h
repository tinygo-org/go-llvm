
#include "llvm-c/DebugInfo.h"
#include "llvm-c/Types.h"
#include "llvm-c/Transforms/PassManagerBuilder.h"

#ifdef __cplusplus
extern "C" {
#endif

void LLVMGlobalObjectAddMetadata(LLVMValueRef objValue, unsigned KindID, LLVMMetadataRef md);

LLVMMemoryBufferRef LLVMGoWriteThinLTOBitcodeToMemoryBuffer(LLVMModuleRef M);

LLVMMetadataRef LLVMGoDIBuilderCreateExpression(LLVMDIBuilderRef Builder,
                                                uint64_t *Addr, size_t Length);

#if LLVM_VERSION_MAJOR < 15
/**
 * Determine whether a pointer is opaque.
 *
 * True if this is an instance of an opaque PointerType.
 *
 * @see llvm::Type::isOpaquePointerTy()
 */
LLVMBool LLVMPointerTypeIsOpaque(LLVMTypeRef Ty);

#endif


#ifdef __cplusplus
}
#endif /* defined(__cplusplus) */
