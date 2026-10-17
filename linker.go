//===- linker.go - Bindings for linker ------------------------------------===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// This file defines bindings for the linker component.
//
//===----------------------------------------------------------------------===//

package llvm

/*
#include "llvm-c/Core.h"
#include "llvm-c/Linker.h"
#include "backports.h"
#include <stdlib.h>
*/
import "C"
import "errors"

func LinkModules(Dest, Src Module) error {
	failed := C.LLVMLinkModules2(Dest.C, Src.C)
	if failed != 0 {
		err := errors.New("Linking failed")
		return err
	}
	return nil
}

// Linker links modules into a single destination module. Unlike LinkModules,
// it keeps its state between calls, so that the cost of each call depends on
// the module linked in rather than on everything linked so far. Use it to link
// many modules into one.
type Linker struct {
	C C.LLVMGoLinkerRef
}

// NewLinker returns a Linker that links modules into dest.
func NewLinker(dest Module) Linker {
	return Linker{C.LLVMGoCreateLinker(dest.C)}
}

// LinkModule links src into the destination module. Like LinkModules, it
// takes ownership of src, which must not be used afterwards.
func (l Linker) LinkModule(src Module) error {
	if C.LLVMGoLinkerLinkModule(l.C, src.C) != 0 {
		return errors.New("Linking failed")
	}
	return nil
}

// Dispose releases the Linker. The destination module stays valid.
func (l Linker) Dispose() {
	C.LLVMGoDisposeLinker(l.C)
}
