//go:build llvm15 || llvm16 || llvm17 || llvm18

package llvm

func (c Context) createConstantRangeAttribute(kind uint, numBits uint, lower, upper []uint64) (a Attribute) {
	return
}
