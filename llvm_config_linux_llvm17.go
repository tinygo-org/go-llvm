//go:build !byollvm && linux && llvm17
// +build !byollvm,linux,llvm17

package llvm

// #cgo CPPFLAGS: -I/usr/lib/llvm-17/include -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo CXXFLAGS: -std=c++17
// #cgo LDFLAGS: -L/usr/lib/llvm-17/lib  -lLLVM-17
import "C"

type run_build_sh int
