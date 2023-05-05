package obj_describe

import (
	"reflect"
	"strconv"
)

const (
	NameTag               = "json"
	IgnoreTag             = "rignore"
	MaxTag                = "max"
	MinTag                = "min"
	ExpressionTag         = "expr"
	ForceExpressionTag    = "force"
	ExpressionRepeatCount = "expr-rc"
	TrueWeight            = "tweight"
	NullWeight            = "nweight"
)

func ParseStruct(obj interface{}) (*FieldDescribe, error) {
	t := reflect.TypeOf(obj)
	return parseType(t)
}

func parseType(t reflect.Type) (*FieldDescribe, error) {
	res := FieldDescribe{}
	res.Type = FieldTypeStruct
	for index := 0; index < t.NumField(); index++ {
		item, err := parseField(t.Field(index))
		if err != nil {
			return nil, err
		}
		if item != nil {
			res.SubFieldDescribes = append(res.SubFieldDescribes, item)
		}
	}

	return &res, nil
}

func parseField(f reflect.StructField) (*FieldDescribe, error) {
	fName := f.Name
	if value, ok := f.Tag.Lookup(NameTag); ok {
		fName = value
	}
	if value, ok := f.Tag.Lookup(IgnoreTag); ok {
		isIgnore, err := strconv.ParseBool(value)
		if err != nil {
			isIgnore = true
		}
		if isIgnore {
			return nil, err
		}
	}

	res := FieldDescribe{}
	res.Name = fName
	if value, ok := f.Tag.Lookup(NullWeight); ok {
		weight, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		res.PointNilWeight = &weight
	}
	if value, ok := f.Tag.Lookup(ExpressionTag); ok {
		res.StringRegexExpression = &value
	}
	if value, ok := f.Tag.Lookup(ForceExpressionTag); ok {
		isForce, err := strconv.ParseBool(value)
		if err != nil {
			isForce = false
		}
		res.ForceGenerateByExpression = isForce
	}
	if value, ok := f.Tag.Lookup(ExpressionRepeatCount); ok {
		rc, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		rcInt := int(rc)
		res.StringRegexExpressionRepeatMaxCount = &rcInt
	}
	switch f.Type.Kind() {
	case reflect.Bool:
		res.Type = FieldTypeBool
		if value, ok := f.Tag.Lookup(TrueWeight); ok {
			weight, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
			res.BoolTrueWeight = &weight
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res.Type = FieldTypeInt
		if value, ok := f.Tag.Lookup(MaxTag); ok {
			maxInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, err
			}
			res.MaxInt = &maxInt
		}
		if value, ok := f.Tag.Lookup(MinTag); ok {
			minInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, err
			}
			res.MinInt = &minInt
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res.Type = FieldTypeInt
		if value, ok := f.Tag.Lookup(MaxTag); ok {
			maxInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, err
			}
			res.MaxInt = &maxInt
			// todo: auto set max value by type
		}
		if value, ok := f.Tag.Lookup(MinTag); ok {
			minInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, err
			}
			if minInt < 0 {
				minInt = 0
			}
			res.MinInt = &minInt
		}
	case reflect.Float64, reflect.Float32:
		res.Type = FieldTypeFloat
		if value, ok := f.Tag.Lookup(MaxTag); ok {
			maxFloat, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
			res.MaxFloat = &maxFloat
		}
		if value, ok := f.Tag.Lookup(MinTag); ok {
			minFloat, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
			// todo: auto set max value by type
			res.MinFloat = &minFloat
		}
	case reflect.String:
		res.Type = FieldTypeString
	case reflect.Array, reflect.Slice:
		res.Type = FieldTypeArray
		if value, ok := f.Tag.Lookup(MaxTag); ok {
			maxLength, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, err
			}
			maxLengthInt := int(maxLength)
			res.MaxArrayLength = &maxLengthInt
		}
		if value, ok := f.Tag.Lookup(MinTag); ok {
			minLength, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, err
			}
			minLengthInt := int(minLength)
			res.MinArrayLength = &minLengthInt
		}

		t := f.Type.Elem()
		for {
			if t.Kind() != reflect.Pointer {
				break
			}
			t = t.Elem()
		}

		var err error
		res.Elem, err = parseType(t)
		if err != nil {
			return nil, err
		}
	case reflect.Struct, reflect.Pointer:
		t := f.Type
		for {
			if t.Kind() != reflect.Pointer {
				break
			}
			t = t.Elem()
		}

		r, err := parseType(t)
		if err != nil {
			return nil, err
		}
		r.Name = fName
		return r, nil
	}

	return &res, nil
}
