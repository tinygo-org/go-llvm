//go:build !byollvm && llvm16

package llvm

// #cgo darwin,amd64 CPPFLAGS: -I/usr/local/opt/llvm@16/include   -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo darwin,amd64 CXXFLAGS: -std=c++17
// #cgo darwin,amd64 LDFLAGS: -L/usr/local/opt/llvm@16/lib -Wl,-search_paths_first -Wl,-headerpad_max_install_names -lLLVM -lz -lm
// #cgo darwin,arm64 CPPFLAGS: -I/opt/homebrew/opt/llvm@16/include   -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo darwin,arm64 CXXFLAGS: -std=c++17
// #cgo darwin,arm64 LDFLAGS: -L/opt/homebrew/opt/llvm@16/lib -Wl,-search_paths_first -Wl,-headerpad_max_install_names -lLLVM -lz -lm
// #cgo linux        CPPFLAGS: -I/usr/lib/llvm-16/include -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo linux        CXXFLAGS: -std=c++17
// #cgo linux        LDFLAGS: -L/usr/lib/llvm-16/lib  -lLLVM-16
import "C"

type run_build_sh int
