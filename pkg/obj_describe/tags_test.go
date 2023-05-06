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
	Arrays     []InnerObj  `max:"0"`
	ArrayPoint []*InnerObj `max:"0"`
	EmptyStr   string      `expr:""`
}

type InnerObj struct {
	SmallInt int8
	TestName string `expr:"[a-z]{10}"`
	BigInt   int    `expr:"123" force:"true"`
}

func TestParseStruct(t *testing.T) {
	ps := NewParser()
	f, err := ps.ParseStruct(TestObj{})
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

	value := TestObj{}
	err = json.Unmarshal(rJson, &value)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value.Arrays)
}
