package llvm_test

import "tinygo.org/x/go-llvm"

import "testing"

// Dummy test function.
// All it does is test whether we can use LLVM at all.
func TestLLVM(t *testing.T) {
	t.Log("LLVM version:", llvm.Version)
}
