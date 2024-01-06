//go:build !byollvm && linux && llvm16

package llvm

// #cgo CPPFLAGS: -I/usr/lib/llvm-16/include -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo CXXFLAGS: -std=c++17
// #cgo LDFLAGS: -L/usr/lib/llvm-16/lib  -lLLVM-16
import "C"

type run_build_sh int
