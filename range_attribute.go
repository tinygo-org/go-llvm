package llvm

// CreateConstantRangeAttribute creates an attribute such as range that holds
// the half-open range [lower, upper) of numBits wide integers. lower and upper
// must each hold at least ceil(numBits/64) words, least significant first.
//
// LLVM before version 19 has no range attribute, so there it returns a nil
// attribute.
func (c Context) CreateConstantRangeAttribute(kind uint, numBits uint, lower, upper []uint64) Attribute {
	words := int(numBits / 64)
	if numBits%64 != 0 {
		words++
	}
	if len(lower) < words || len(upper) < words {
		panic("llvm: CreateConstantRangeAttribute: lower and upper must hold ceil(numBits/64) words")
	}
	return c.createConstantRangeAttribute(kind, numBits, lower, upper)
}
