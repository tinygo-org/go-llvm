//go:build !byollvm && llvm23

package llvm

// Homebrew may still ship LLVM 23 under the unversioned llvm formula.
// Search both llvm@23 and llvm for headers and libraries.

// #cgo darwin,amd64 CPPFLAGS: -I/usr/local/opt/llvm@23/include -I/usr/local/opt/llvm/include   -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo darwin,amd64 CXXFLAGS: -std=c++17
// #cgo darwin,amd64 LDFLAGS: -L/usr/local/opt/llvm@23/lib -L/usr/local/opt/llvm/lib -Wl,-search_paths_first -Wl,-headerpad_max_install_names -lLLVM -lz -lm
// #cgo darwin,arm64 CPPFLAGS: -I/opt/homebrew/opt/llvm@23/include -I/opt/homebrew/opt/llvm/include   -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo darwin,arm64 CXXFLAGS: -std=c++17
// #cgo darwin,arm64 LDFLAGS: -L/opt/homebrew/opt/llvm@23/lib -L/opt/homebrew/opt/llvm/lib -Wl,-search_paths_first -Wl,-headerpad_max_install_names -lLLVM -lz -lm
// #cgo freebsd      CPPFLAGS: -I/usr/local/llvm23/include -I/usr/local/llvm23/include/llvm-c -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo freebsd      CXXFLAGS: -std=c++17
// #cgo freebsd      LDFLAGS: -L/usr/local/llvm23/lib -lLLVM
// #cgo linux        CPPFLAGS: -I/usr/include/llvm-23 -I/usr/include/llvm-c-23 -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo linux        CXXFLAGS: -std=c++17
// #cgo linux        LDFLAGS: -L/usr/lib/llvm-23/lib -lLLVM-23
import "C"

type run_build_sh int
