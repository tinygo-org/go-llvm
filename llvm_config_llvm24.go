//go:build !byollvm && llvm24

package llvm

// #cgo darwin,amd64 CPPFLAGS: -I/usr/local/opt/llvm@24/include   -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo darwin,amd64 CXXFLAGS: -std=c++17
// #cgo darwin,amd64 LDFLAGS: -L/usr/local/opt/llvm@24/lib -Wl,-search_paths_first -Wl,-headerpad_max_install_names -lLLVM -lz -lm
// #cgo darwin,arm64 CPPFLAGS: -I/opt/homebrew/opt/llvm@24/include   -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo darwin,arm64 CXXFLAGS: -std=c++17
// #cgo darwin,arm64 LDFLAGS: -L/opt/homebrew/opt/llvm@24/lib -Wl,-search_paths_first -Wl,-headerpad_max_install_names -lLLVM -lz -lm
// #cgo freebsd      CPPFLAGS: -I/usr/local/llvm24/include -I/usr/local/llvm24/include/llvm-c -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo freebsd      CXXFLAGS: -std=c++17
// #cgo freebsd      LDFLAGS: -L/usr/local/llvm24/lib -lLLVM
// #cgo linux        CPPFLAGS: -I/usr/include/llvm-24 -I/usr/include/llvm-c-24 -I/usr/lib/llvm24/include -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo linux        CXXFLAGS: -std=c++17
// #cgo linux        LDFLAGS: -L/usr/lib/llvm-24/lib -L/usr/lib/llvm24/lib -lLLVM-24
import "C"

type run_build_sh int
