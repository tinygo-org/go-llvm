//go:build !byollvm && linux && llvm15

package llvm

// #cgo CPPFLAGS: -I/usr/lib/llvm-15/include -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo CXXFLAGS: -std=c++14
// #cgo LDFLAGS: -L/usr/lib/llvm-15/lib  -lLLVM-15
import "C"

type run_build_sh int
