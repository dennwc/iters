package iters

// RowsScanner is a minimal interface that sql.Rows implements.
type RowsScanner interface {
	Next() bool
	Err() error
	Close() error
}

// FromRows converts sql.Rows to Iter.
func FromRows[T any](rows RowsScanner, scan func() (T, error)) Iter[T] {
	return &rowsIter[T]{
		rows: rows,
		scan: scan,
	}
}

type rowsIter[T any] struct {
	rows RowsScanner
	scan func() (T, error)
}

func (it *rowsIter[T]) Close() {
	_ = it.rows.Close()
}
func (it *rowsIter[T]) Next() (T, error) {
	var zero T
	if !it.rows.Next() {
		return zero, it.rows.Err()
	}
	return it.scan()
}
