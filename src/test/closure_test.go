package test

import "testing"

func TestClosure(t *testing.T) {
	inner := getInnerFunc()
	t.Logf("Inner Func: %p, Local Ref: %p", inner, &inner)
	for i := 0; i < 10; i++ {
		idx := inner()
		t.Logf("Index: %d", idx)
	}
}

func TestSameClosure(t *testing.T) {
	inner1 := getInnerFunc()
	t.Logf("(1) Func Pointer: %p, Local Ref: %p", inner1, &inner1)
	inner2 := getInnerFunc()
	t.Logf("(2) Func Pointer: %p, Local Ref: %p", inner2, &inner2)
	for i := 0; i < 4; i++ {
		idx1 := inner1()
		t.Logf("(1) Index: %d", idx1)
		idx2 := inner2()
		t.Logf("(2) Index: %d", idx2)
	}
}

func TestPassClosure(t *testing.T) {
	inner := getInnerFunc()
	t.Logf("Func Pointer: %p, Local Ref: %p", inner, &inner)
	for i := 0; i < 4; i++ {
		idx := inner()
		t.Logf("Index: %d", idx)
	}
	passClosure(t, inner)
}

func getInnerFunc() func() int {
	idx := 0

	return func() int {
		idx++
		return idx
	}
}

func passClosure(t *testing.T, closure func() int) {
	t.Logf("Func Pointer: %p, Local Ref: %p", closure, &closure)
	for i := 0; i < 4; i++ {
		idx := closure()
		t.Logf("Index: %d", idx)
	}
}
