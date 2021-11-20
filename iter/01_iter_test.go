package iter_test

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/narqo/test-go-generics/iter"
)

func printNext[T any](it iter.Iterator[T]) {
	v, err := it.Next()
	fmt.Printf("val %v, err %v\n", v, err)
}

type sliceIterator[T any] struct {
	s []T
}

func (it *sliceIterator[T]) Next() (T, error) {
	if len(it.s) == 0 {
		var zero T
		return zero, io.EOF
	}
	next := it.s[0]
	it.s = it.s[1:]
	return next, nil
}

func ExampleIterator_sliceIterator() {
	a := &sliceIterator[int]{
		[]int{1, 2, 3},
	}

	printNext[int](a)
	printNext[int](a)
	printNext[int](a)
	printNext[int](a)

	// Output:
	// val 1, err <nil>
	// val 2, err <nil>
	// val 3, err <nil>
	// val 0, err EOF
}

type csvIterator struct {
	r *csv.Reader
}

func newCSVIterator(r io.Reader) *csvIterator {
	return &csvIterator{
		r: csv.NewReader(r),
	}
}

func (it *csvIterator) Next() ([]string, error) {
	return it.r.Read()
}

func ExampleIterator_csvIterator() {
	data := `1,2,3
a,b,c`
	a := newCSVIterator(strings.NewReader(data))

	printNext[[]string](a)
	printNext[[]string](a)
	printNext[[]string](a)

	// Output:
	// val [1 2 3], err <nil>
	// val [a b c], err <nil>
	// val [], err EOF
}

type timeInterator struct{}

func (it timeInterator) Next() (time.Time, error) {
	return time.Now(), nil
}

func ExampleIterator_timeIterator() {
	var it timeInterator
	it.Next()
	it.Next()
}
