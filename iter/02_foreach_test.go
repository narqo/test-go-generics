package iter_test

import (
	"fmt"
	"io"

	"github.com/narqo/test-go-generics/iter"
)

func ExampleIterator_forEach() {
	a := &sliceIterator[int]{
		[]int{1, 2, 3},
	}

	forEach[int](
		a,
		func(v int) {
			fmt.Println(v)
		},
	)

	// Output:
	// 1
	// 2
	// 3
}

func forEach[T any](it iter.Iterator[T], f func(T)) error {
	for {
		v, err := it.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		f(v)
	}
}
