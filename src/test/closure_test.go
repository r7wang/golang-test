package test

import "testing"

// Summary:
//  - As long as multiple closures reference a variable within the same scope, any modifications to
//    the underlying variable will be seen by other closures.
//  - Every new instance of a closure that references a local variable will reference a different
//    copy of that local variable.

func TestClosure(t *testing.T) {
	inner := getInnerFunc(t)
	t.Logf("Inner Func: %p, Local Ref: %p", inner, &inner)
	for i := 0; i < 10; i++ {
		idx := inner()
		t.Logf("Index: %d", idx)
	}
}

func TestSameClosure(t *testing.T) {
	inner1 := getInnerFunc(t)
	t.Logf("(1) Func Pointer: %p, Local Ref: %p", inner1, &inner1)
	inner2 := getInnerFunc(t)
	t.Logf("(2) Func Pointer: %p, Local Ref: %p", inner2, &inner2)
	for i := 0; i < 4; i++ {
		idx1 := inner1()
		t.Logf("(1) Index: %d", idx1)
		idx2 := inner2()
		t.Logf("(2) Index: %d", idx2)
	}
}

func TestPassClosure(t *testing.T) {
	inner := getInnerFunc(t)
	t.Logf("Func Pointer: %p, Local Ref: %p", inner, &inner)
	for i := 0; i < 4; i++ {
		idx := inner()
		t.Logf("Index: %d", idx)
	}
	passClosure(t, inner)
}

func TestMultipleSameScopeClosure(t *testing.T) {
	inner1, inner2 := getMultipleInnerFuncSameScope(t)
	for i := 0; i < 4; i++ {
		idx1 := inner1()
		t.Logf("(1) Index: %d", idx1)
		idx2 := inner2()
		t.Logf("(2) Index: %d", idx2)
	}
}

func getInnerFunc(t *testing.T) func() int {
	idx := 0

	closure := func() int {
		idx++
		return idx
	}
	t.Logf("Closure Pointer: %p, Local Ref: %p", closure, &closure)
	return closure
}

func getMultipleInnerFuncSameScope(t *testing.T) (func() int, func() int) {
	idx := 0

	closure1 := func() int {
		idx++
		return idx
	}
	closure2 := func() int {
		idx++
		return idx
	}
	return closure1, closure2
}

func passClosure(t *testing.T, closure func() int) {
	t.Logf("Func Pointer: %p, Local Ref: %p", closure, &closure)
	for i := 0; i < 4; i++ {
		idx := closure()
		t.Logf("Index: %d", idx)
	}
}
