//go:build !byollvm && linux && !llvm13
// +build !byollvm,linux,!llvm13

package llvm

// // Note: disabling deprecated warnings as a transition to new function
// // signatures. We should fix this in LLVM 15 or so, when support for
// // LLVM 13 is dropped and the old function signatures aren't supported
// // anymore.
// // Examples are: LLVMConstGEP2, LLVMAddAlias2, ...
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #cgo CPPFLAGS: -I/usr/lib/llvm-14/include -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo CXXFLAGS: -std=c++14
// #cgo LDFLAGS: -L/usr/lib/llvm-14/lib  -lLLVM-14
import "C"

type run_build_sh int
