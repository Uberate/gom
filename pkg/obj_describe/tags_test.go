package obj_describe

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestObj struct {
	TestBool  bool `json:"test_bool" `
	TestBool2 bool `tweight:"1"`
	InnerObj
	Arrays     []InnerObj
	ArrayPoint []*InnerObj
}

type InnerObj struct {
	SmallInt int    `max:"10" min:"0"`
	TestName string `expr:"[a-z]{10}"`
	BigInt   int    `min:"1000"`
}

func TestParseStruct(t *testing.T) {
	f, err := ParseStruct(TestObj{})
	if err != nil {
		fmt.Println(err)
	}

	fJson, _ := json.Marshal(f)
	fmt.Println(string(fJson))

	res, err := f.GenerateValue()
	if err != nil {
		fmt.Println(err)
	}
	rJson, _ := json.Marshal(res)
	fmt.Println(string(rJson))
}
