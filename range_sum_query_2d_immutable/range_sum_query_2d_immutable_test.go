package range_sum_query_2d_immutable

import "testing"

func TestRangeSumQuery2dImmutable(t *testing.T) {
	matrix := Constructor(
		[][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		},
	)

	got := matrix.SumRegion(2, 1, 4, 3)
	want := 8
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = matrix.SumRegion(1, 1, 2, 2)
	want = 11
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = matrix.SumRegion(1, 2, 2, 4)
	want = 12
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
