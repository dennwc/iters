package testit

import (
	"context"
	"testing"

	"github.com/shoenig/test/must"

	"github.com/dennwc/iters"
)

func ExpectIter[T any](t testing.TB, exp []T, expErr error, it iters.Iter[T]) {
	defer it.Close()
	got, err := iters.All(it)
	if expErr != nil {
		must.Eq(t, expErr, err)
	} else {
		must.NoError(t, err)
	}
	must.Eq(t, exp, got)
}

func ExpectIterCtx[T any](t testing.TB, exp []T, expErr error, it iters.IterCtx[T]) {
	ctx := context.Background()
	defer it.Close()
	got, err := iters.AllCtx(ctx, it)
	if expErr != nil {
		must.Eq(t, expErr, err)
	} else {
		must.NoError(t, err)
	}
	must.Eq(t, exp, got)
}

func ExpectPageIter[T any](t testing.TB, exp []T, expErr error, it iters.PageIter[T]) {
	ctx := context.Background()
	defer it.Close()
	got, err := iters.AllPages(ctx, it)
	if expErr != nil {
		must.Eq(t, expErr, err)
	} else {
		must.NoError(t, err)
	}
	must.Eq(t, exp, got)
}
