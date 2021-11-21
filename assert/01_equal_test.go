package assert_test

import (
	"fmt"
	"time"

	"github.com/narqo/test-go-generics/assert"
)

func ExampleEqual() {
	// comparable (int)
	fmt.Println(assert.Equal(1, 1))
	fmt.Println(assert.Equal(1, 2))

	// compile type error: type string of "1" does not match inferred type int for T
	//fmt.Println(assert.Equal(1, "1"))

	// comparable (string)
	fmt.Println(assert.Equal("abc", "abc"))
	fmt.Println(assert.Equal("abc", "cba"))

	// equalable
	t1, _ := time.Parse(time.RFC822Z, "20 Nov 21 15:04 +0200")
	t2, _ := time.Parse(time.RFC822Z, "20 Nov 21 13:04 +0000")

	fmt.Println(assert.Equal(t1, t2))
	fmt.Println(assert.Equal(t1, time.Time{}))

	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
}
