package main

import (
	"encoding/json"
	"fmt"
	"gom/pkg/obj_describe"
)

func main() {
	s := Student{}

	p := obj_describe.NewParser()
	describe, err := p.ParseStruct(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	genValue, err := describe.GenerateValue()
	if err != nil {
		fmt.Println(err)
		return
	}

	genValueJsonStr, err := json.Marshal(genValue)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(genValueJsonStr))
}
