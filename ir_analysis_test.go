package llvm

import (
	"math/rand"
	"testing"
)

// reachesItself is a plain version of BasicBlock.InCycle on an explicit graph.
func reachesItself(succs [][]int, n int) bool {
	seen := make([]bool, len(succs))
	stack := append([]int(nil), succs[n]...)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur == n {
			return true
		}
		if seen[cur] {
			continue
		}
		seen[cur] = true
		stack = append(stack, succs[cur]...)
	}
	return false
}

// buildCFG creates a function whose blocks have the given successors. A block
// without successors returns.
func buildCFG(ctx Context, mod Module, succs [][]int) []BasicBlock {
	i1 := ctx.Int1Type()
	fn := AddFunction(mod, "f", FunctionType(ctx.VoidType(), []Type{i1, ctx.Int32Type()}, false))
	blocks := make([]BasicBlock, len(succs))
	for i := range blocks {
		blocks[i] = ctx.AddBasicBlock(fn, "")
	}
	b := ctx.NewBuilder()
	defer b.Dispose()
	for i, out := range succs {
		b.SetInsertPointAtEnd(blocks[i])
		switch len(out) {
		case 0:
			b.CreateRetVoid()
		case 1:
			b.CreateBr(blocks[out[0]])
		case 2:
			b.CreateCondBr(fn.Param(0), blocks[out[0]], blocks[out[1]])
		default:
			sw := b.CreateSwitch(fn.Param(1), blocks[out[0]], len(out)-1)
			for j, o := range out[1:] {
				sw.AddCase(ConstInt(ctx.Int32Type(), uint64(j), false), blocks[o])
			}
		}
	}
	return blocks
}

func checkInCycle(t *testing.T, succs [][]int) {
	t.Helper()
	ctx := NewContext()
	defer ctx.Dispose()
	mod := ctx.NewModule("")
	defer mod.Dispose()
	blocks := buildCFG(ctx, mod, succs)
	for i, bb := range blocks {
		if got, want := bb.InCycle(), reachesItself(succs, i); got != want {
			t.Errorf("graph %v block %d: InCycle() = %v, want %v", succs, i, got, want)
		}
	}
}

func TestBasicBlockInCycle(t *testing.T) {
	for _, succs := range [][][]int{
		{{}},                           // a single block
		{{0}},                          // a self loop in the entry block
		{{1}, {}},                      // a straight line
		{{1}, {1, 2}, {}},              // a self loop after the entry
		{{1}, {2}, {1, 3}, {}},         // a loop with two blocks
		{{1, 2}, {3}, {3}, {}},         // a diamond has no cycle
		{{1, 2}, {2}, {1, 3}, {}},      // an irreducible loop
		{{1}, {}, {3}, {2}},            // unreachable blocks, one of them in a cycle
		{{1}, {}, {2}},                 // an unreachable self loop
		{{1}, {2, 3}, {1}, {4, 0}, {}}, // nested cycles that go back to the entry
	} {
		checkInCycle(t, succs)
	}
}

func TestBasicBlockInCycleRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for iter := 0; iter < 300; iter++ {
		n := 1 + rng.Intn(14)
		succs := make([][]int, n)
		for i := range succs {
			// Few edges per block so that acyclic graphs are common.
			for k := rng.Intn(4); k > 0; k-- {
				succs[i] = append(succs[i], rng.Intn(n))
			}
			// The entry block cannot have predecessors.
		}
		for i := range succs {
			for j, s := range succs[i] {
				if s == 0 {
					succs[i][j] = (s + 1) % n
				}
			}
		}
		checkInCycle(t, succs)
	}
}

func TestBasicBlockInCycleLargeFunction(t *testing.T) {
	// A long acyclic chain must not be reported as a cycle, however many
	// blocks there are, and a loop at the end must be.
	const n = 500
	succs := make([][]int, n)
	for i := 0; i < n-1; i++ {
		succs[i] = []int{i + 1}
	}
	succs[n-1] = []int{n - 2}
	checkInCycle(t, succs)
}

func TestUnderlyingObject(t *testing.T) {
	ctx := NewContext()
	defer ctx.Dispose()
	mod := ctx.NewModule("")
	defer mod.Dispose()

	i8 := ctx.Int8Type()
	i32 := ctx.Int32Type()
	ptr := PointerType(i8, 0)
	global := AddGlobal(mod, i8, "g")
	fn := AddFunction(mod, "f", FunctionType(ctx.VoidType(), []Type{ptr, ctx.Int1Type()}, false))
	b := ctx.NewBuilder()
	defer b.Dispose()
	b.SetInsertPointAtEnd(ctx.AddBasicBlock(fn, "entry"))

	alloca := b.CreateAlloca(i8, "a")
	other := b.CreateAlloca(i8, "b")
	idx := ConstInt(i32, 1, false)

	// A chain of GEPs, longer than the default lookup limit of LLVM.
	chain := alloca
	for i := 0; i < 40; i++ {
		chain = b.CreateGEP(i8, chain, []Value{idx}, "")
	}
	tests := []struct {
		name  string
		value Value
		want  Value
	}{
		{"alloca", alloca, alloca},
		{"gep", b.CreateGEP(i8, alloca, []Value{idx}, ""), alloca},
		{"long gep chain", chain, alloca},
		{"argument", b.CreateGEP(i8, fn.Param(0), []Value{idx}, ""), fn.Param(0)},
		{"global", b.CreateGEP(i8, global, []Value{idx}, ""), global},
	}
	sel := b.CreateSelect(fn.Param(1), alloca, other, "")
	tests = append(tests, struct {
		name  string
		value Value
		want  Value
	}{"select stays", sel, sel})
	for _, tc := range tests {
		if got := tc.value.UnderlyingObject(); got != tc.want {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
		}
	}
}
