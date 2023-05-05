package obj_describe

import (
	"encoding/json"
	"fmt"
	"testing"
)

var inputs = `{
"name": "test",
"describe": {
"
}
}`

func Test_Common(t *testing.T) {
	v := FieldDescribe{}
	v.Type = FieldTypeBool
	v.Name = "test"

	value, err := v.GenerateValue()
	if err != nil {
		fmt.Println(err)
	} else {
		jv, _ := json.Marshal(value)
		fmt.Println(string(jv))
	}

	jsonValue, _ := json.Marshal(v)

	fmt.Println(string(jsonValue))
}
