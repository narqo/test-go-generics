package optional_test

import (
	"encoding/json"
	"fmt"

	"github.com/narqo/test-go-generics/optional"
)

type testStruct struct {
	IntVal    int                  `json:",omitempty"`
	OptIntVal *optional.Value[int] `json:",omitempty"`

	BoolVal    bool                  `json:",omitempty"`
	OptBoolVal *optional.Value[bool] `json:",omitempty"`

	StringVal    string                  `json:",omitempty"`
	OptStringVal *optional.Value[string] `json:",omitempty"`
}

func ExampleOptional() {
	ts := testStruct{}
	data, _ := json.Marshal(ts)
	fmt.Println(string(data))

	ts = testStruct{
		IntVal:       0,
		OptIntVal:    optional.New(0),
		BoolVal:      false,
		OptBoolVal:   optional.New(false),
		StringVal:    "",
		OptStringVal: optional.New(""),
	}
	data, _ = json.Marshal(ts)
	fmt.Println(string(data))

	ts = testStruct{
		IntVal:       100,
		OptIntVal:    optional.New(100),
		BoolVal:      true,
		OptBoolVal:   optional.New(true),
		StringVal:    "abc",
		OptStringVal: optional.New("abc"),
	}
	data, _ = json.Marshal(ts)
	fmt.Println(string(data))

	// Output:
	// {}
	// {"OptIntVal":0,"OptBoolVal":false,"OptStringVal":""}
	// {"IntVal":100,"OptIntVal":100,"BoolVal":true,"OptBoolVal":true,"StringVal":"abc","OptStringVal":"abc"}
}
