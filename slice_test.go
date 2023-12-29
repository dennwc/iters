package iters_test

import (
	"testing"

	. "github.com/dennwc/iters"
	"github.com/dennwc/iters/testit"
)

var (
	_ TIter      = Slice([]int{})
	_ TIterCtx   = Slice([]int{})
	_ TPageIter  = Slice([]int{})
	_ TPagedIter = Slice([]int{})
)

func TestSlice(t *testing.T) {
	testit.ExpectIter(t, []int(nil), nil, Slice[int](nil))
	testit.ExpectPageIter(t, []int(nil), nil, Slice[int](nil))
	testit.ExpectPageIter(t, []int(nil), nil, PageSlice[int](nil))
	testit.ExpectIter(t, []int{1, 2, 3}, nil, Slice([]int{1, 2, 3}))
	testit.ExpectPageIter(t, []int{1, 2, 3}, nil, Slice([]int{1, 2, 3}))
	testit.ExpectPageIter(t, []int{1, 2, 3, 4}, nil, PageSlice([][]int{{1}, {2, 3}, {4}}))
}
