// +build !byollvm

package llvm

// Automatically generated by `make config BUILDDIR=`, do not edit.

// #cgo CPPFLAGS: -I/usr/local/opt/llvm@9/include   -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L/usr/local/opt/llvm@9/lib -Wl,-search_paths_first -Wl,-headerpad_max_install_names -lLLVM -lz -lcurses -lm -lxml2 -L/usr/local/opt/libffi/lib -lffi
import "C"

type (run_build_sh int)
