//go:build !llvm17
// +build !llvm17

//===- ir.go - Bindings for ir --------------------------------------------===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// This file defines bindings for the ir component.
//
//===----------------------------------------------------------------------===//

package llvm

/*
#include "llvm-c/Core.h"
#include "llvm-c/Comdat.h"
#include "IRBindings.h"
#include <stdlib.h>
*/
import "C"

//-------------------------------------------------------------------------
// llvm.Value
//-------------------------------------------------------------------------

func ConstSelect(cond, iftrue, iffalse Value) (rv Value) {
	rv.C = C.LLVMConstSelect(cond.C, iftrue.C, iffalse.C)
	return
}
