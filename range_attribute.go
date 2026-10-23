package llvm

import "math"

// CreateConstantRangeAttribute creates a constant range attribute. kind is
// normally AttributeKindID("range"). The attribute holds the half-open range
// [lower, upper) of numBits wide integers, which can wrap. lower and upper hold
// ceil(numBits/64) words each, least significant word first. lower and upper
// must not be equal, except 0 and 0 for the empty range. For example, the range
// of values that are never negative is [0, 1<<(numBits-1)).
//
// LLVM before version 19 has no range attribute, so there it returns a nil
// attribute. Check IsNil before you add the attribute to a function or call.
func (c Context) CreateConstantRangeAttribute(kind uint, numBits uint, lower, upper []uint64) Attribute {
	if numBits == 0 || numBits > math.MaxUint32 {
		panic("llvm: CreateConstantRangeAttribute: numBits must be between 1 and 2^32-1")
	}
	words := int((numBits-1)/64 + 1)
	if len(lower) < words || len(upper) < words {
		panic("llvm: CreateConstantRangeAttribute: lower and upper must hold ceil(numBits/64) words")
	}
	return c.createConstantRangeAttribute(kind, numBits, lower, upper)
}
