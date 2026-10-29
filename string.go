//===- string.go - Stringer implementation for Type -----------------------===//
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===----------------------------------------------------------------------===//
//
// This file implements the Stringer interface for the Type type.
//
//===----------------------------------------------------------------------===//

package llvm

import (
	"fmt"
	"strings"
)

// String returns a diagnostic description of t. Pointers are printed as
// PointerType, without a pointee type or address space.
func (t Type) String() string {
	ts := typeStringer{s: make(map[Type]string)}
	return ts.typeString(t)
}

type typeStringer struct {
	s map[Type]string
}

func (ts *typeStringer) typeString(t Type) string {
	if s, ok := ts.s[t]; ok {
		return s
	}

	k := t.TypeKind()
	s := strings.TrimSuffix(k.String(), "Kind")

	switch k {
	case ArrayTypeKind:
		s += fmt.Sprintf("(%v[%v])", ts.typeString(t.ElementType()), t.ArrayLength())
	case FunctionTypeKind:
		params := t.ParamTypes()
		s += "("
		if len(params) > 0 {
			s += fmt.Sprintf("%v", ts.typeString(params[0]))
			for i := 1; i < len(params); i++ {
				s += fmt.Sprintf(", %v", ts.typeString(params[i]))
			}
		}
		s += fmt.Sprintf("):%v", ts.typeString(t.ReturnType()))
	case StructTypeKind:
		if name := t.StructName(); name != "" {
			ts.s[t] = "%" + name
			s = fmt.Sprintf("%%%s: %s", name, s)
		}
		etypes := t.StructElementTypes()
		s += "("
		if n := len(etypes); n > 0 {
			s += ts.typeString(etypes[0])
			for i := 1; i < n; i++ {
				s += fmt.Sprintf(", %v", ts.typeString(etypes[i]))
			}
		}
		s += ")"
	case IntegerTypeKind:
		s += fmt.Sprintf("(%d bits)", t.IntTypeWidth())
	}

	ts.s[t] = s
	return s
}
