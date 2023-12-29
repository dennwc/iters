package iters

import "context"

type Iter[T any] interface {
	// Next returns the next item.
	// It returns io.EOF if there are no more items left.
	Next() (T, error)
	// Close the iterator.
	Close()
}

type IterCtx[T any] interface {
	// NextCtx returns the next item.
	// It returns io.EOF if there are no more items left.
	NextCtx(ctx context.Context) (T, error)
	// Close the iterator.
	Close()
}
