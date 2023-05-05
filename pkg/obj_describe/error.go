package obj_describe

import (
	"fmt"
	"reflect"
)

func ConvertError(specifyValue interface{}, expectType string) error {
	return fmt.Errorf("can't convert value [%s] to [%s]", reflect.TypeOf(specifyValue).Name(), expectType)
}
