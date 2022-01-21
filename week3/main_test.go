package main

import "testing"

func TestIsAtBase(t *testing.T) {
	matrix1 := [][]int{
		{1, 3, 5},
		{2, 1, 2},
		{4, 3, 1},
	}
	x := 1
	y := 1

	test1 := isAtBase(matrix1, x, y)
	if test1 {
		t.Errorf("Not at base")
		return
	}

	test2 := isAtBase(matrix1, 2, 2)
	if !test2 {
		t.Errorf("Is at base but says it isn't")
		return
	}
}

func BenchmarkStore(b *testing.B) {
	for n := 0; n < b.N; n++ {
		store(matrix4)
	}
}
