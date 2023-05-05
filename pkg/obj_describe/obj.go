package obj_describe

import (
	"fmt"
	"gom/pkg/regexp_trans"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	FieldTypeBool   = "bool"
	FieldTypeInt    = "int"
	FieldTypeFloat  = "float"
	FieldTypeStruct = "struct"
	FieldTypeArray  = "array"
	FieldTypeString = "string"
)

// FieldDescribe describe how field value generate.
//
// FieldDescribe support the JSON-Data-Structure. JSON data type: number(int, float), string, bool, array and map. All
// the structure in everywhere can be marshaled and unmarshalled by json.
//
// The field describe like this format:
//
// {
//     "name": {
//         "type": "field.type=int",
//         "max-int": "123",
//         "min-int": "0"
//     }
// }
//
// If some field is struct(in json data is map), it contains the SubFieldDescribes. And for array, has Elem.
//
// For now, the FieldDescribe can't describe the different elem structure in one array.
type FieldDescribe struct {

	// inner random instance
	ran *rand.Rand

	Name string `json:"name"`

	fieldValueDescribe `json:"describe"`
}

type fieldValueDescribe struct {
	Type string `json:"type"`

	// RandomSeed used to create a random instance. If RandomSeed is empty, use now time nanosecond as seed.
	RandomSeed *int64 `json:"random-seed,omitempty"`

	Elem *FieldDescribe `json:"elem,omitempty"`

	// If enable ForceGenerateByExpression, the Type is FieldTypeInt, FieldTypeBool, FieldTypeString, FieldTypeFloat,
	// use string generate value as result(try to convert).
	ForceGenerateByExpression bool `json:"force-expression"`

	// The FieldTypeStruct has inner field, all the settable will describe here.
	SubFieldDescribes []*FieldDescribe `json:"sub-field-describes,omitempty"`

	MinInt *int64 `json:"min-i,omitempty"` // If Type is FieldTypeInt, the min value is it(contain MinInt value).
	MaxInt *int64 `json:"max-i,omitempty"` // If Type is FieldTypeInt, the max value is it(contain MaxInt value).

	MaxFloat *float64 `json:"max-f,omitempty"` // If Type is FieldTypeFloat, the max value is it(contain MaxFloat value).
	MinFloat *float64 `json:"min-f,omitempty"` // If Type is FieldTypeFloat, the min value is it(contain MinFloat value).

	// If Type is FieldTypeBool, the BoolTrueWeight is the probability of true value. The BoolTrueWeight must in 0-1.
	// Else, if the value bigger than 1 the value must true. Same case, if the value smaller than zero, the value must
	// be false.
	BoolTrueWeight *float64 `json:"bool-true-weight,omitempty"`

	// The PointNilWeigh is the probability of nil value. The PointNilWeight must in 0-1.
	// Else, if the value bigger than 1 the value must true. Same case, if the value smaller than zero, the value must
	// have real value of Elem. If the PointNilWeight is nil, that mean the value can't be null.
	PointNilWeight *float64 `json:"point-nil-weight,omitempty"`

	// If Type is FieldTypeArray, the array max length is the MaxArrayLength(contain MaxArrayLength). And if
	// MaxArrayLength less than zero or MinArrayLength, the MaxArrayLength set to equals of MinArrayLength.
	MaxArrayLength *int `json:"max-a,omitempty"`

	// If Type is FieldTypeArray, the array min length is the MinArrayLength(contain MinArrayLength). And if
	// MinArrayLength less than zero, MinArrayLength is zero.
	MinArrayLength *int `json:"min-a,omitempty"`

	// If Type is FieldTypeString, the string will generator by StringRegexExpression. The value of
	// StringRegexExpression follow the golang-REV2 standard.
	//
	// If RandomStringSets not emtpy, use StringRegexExpression value to generate the string value(ignore the
	// RandomStringSets). Only the StringRegexExpression is emtpy, use RandomStringSets. About more info of the
	// RandomStringSets, see doc of RandomStringSets.
	StringRegexExpression *string `json:"string-regex-expression,omitempty"`

	// StringRegexExpressionRepeatMaxCount will inject to regexp_trans.Generator. And only be used when Type is
	// FieldTypeString and StringRegexExpression not emtpy.
	StringRegexExpressionRepeatMaxCount *int `json:"string-regex-expression-repeat-max-count,omitempty"`

	// RandomStringSets storage some string value. The generator will generate the string from sets. But the
	// RandomStringSets be used only at StringRegexExpression is emtpy. If both of RandomStringSets and
	// StringRegexExpression is emtpy, return empty string as result.
	RandomStringSets []string `json:"random-string-sets,omitempty"`
}

func (fd *FieldDescribe) GenerateValue() (map[string]interface{}, error) {

	seed := time.Now().UnixNano()
	if fd.RandomSeed != nil {
		seed = *fd.RandomSeed
	}

	fd.ran = rand.New(rand.NewSource(seed))

	res := map[string]interface{}{}

	if fd.PointNilWeight != nil {
		probability := fd.ran.Float64()
		if probability < *fd.PointNilWeight {
			res[fd.Name] = nil
			return res, nil
		}
	}

	if fd.ForceGenerateByExpression {
		if fd.Type == FieldTypeBool ||
			fd.Type == FieldTypeInt ||
			fd.Type == FieldTypeFloat ||
			fd.Type == FieldTypeString {
			repeatCount := 10
			if fd.StringRegexExpressionRepeatMaxCount != nil {
				repeatCount = *fd.StringRegexExpressionRepeatMaxCount
			}
			expression := ""
			if fd.StringRegexExpression != nil {
				expression = *fd.StringRegexExpression
			}
			var err error
			res[fd.Name], err = fd.forceExpression(fd.Type, expression, seed, repeatCount)
			if err != nil {
				return nil, err
			}

			return res, nil
		}
	}

	switch fd.Type {
	case FieldTypeBool:
		weight := 0.5
		if fd.BoolTrueWeight != nil {
			weight = *fd.BoolTrueWeight
		}

		res[fd.Name] = fd.ran.Float64() < weight
	case FieldTypeInt:
		max := int64(math.MaxInt)
		min := int64(0)
		if fd.MaxInt != nil {
			max = *fd.MaxInt
		}
		if fd.MinInt != nil {
			min = *fd.MinInt
		}
		if min > max {
			return nil, fmt.Errorf("max value should less than(or equals of) min value, "+
				"but max: [%d], min: [%d]", max, min)
		}
		res[fd.Name] = fd.ran.Int63n(max-min) + min
	case FieldTypeFloat:
		max := float64(math.MaxFloat32)
		min := float64(0)

		if fd.MaxFloat != nil {
			max = *fd.MaxFloat
		}
		if fd.MinFloat != nil {
			min = *fd.MinFloat
		}
		if min > max {
			return nil, fmt.Errorf("max value should less than(or equals of) min value, "+
				"but max: [%f], min: [%f]", max, min)
		}
		res[fd.Name] = math.Mod(fd.ran.NormFloat64(), max-min) + max
	case FieldTypeStruct:
		sort.Slice(fd.SubFieldDescribes, func(i, j int) bool {
			return strings.Compare(fd.SubFieldDescribes[i].Name, fd.SubFieldDescribes[j].Name) < 0
		})

		structRes := map[string]interface{}{}
		for _, item := range fd.SubFieldDescribes {
			subMap, err := item.GenerateValue()
			if err != nil {
				return nil, err
			}
			for key, value := range subMap {
				structRes[key] = value
			}
			res[fd.Name] = structRes
		}
	case FieldTypeArray:
		maxLength := 10
		minLength := 0
		if fd.MaxArrayLength != nil {
			maxLength = *fd.MaxArrayLength
		}
		if fd.MinArrayLength != nil {
			minLength = *fd.MinArrayLength
		}

		if minLength <= 0 {
			minLength = 0
		}

		if maxLength > minLength {
			return nil, fmt.Errorf("max length of arrya should less than(or equals of) min length of array, "+
				"but max: [%d], min: [%d]", maxLength, minLength)
		}

		randomLength := fd.ran.Intn(maxLength-minLength) + minLength

		var arrayRes []interface{}
		for index := 0; index < randomLength; index++ {
			if fd.Elem == nil {
				return nil, fmt.Errorf("array has no elem describe")
			}

			arrGen, err := fd.Elem.GenerateValue()
			if err != nil {
				return nil, err
			}
			arrayRes = append(arrayRes, arrGen)
		}
		res[fd.Name] = randomLength

	case FieldTypeString:
		if fd.StringRegexExpression != nil {
			repeat := 10
			if fd.StringRegexExpressionRepeatMaxCount != nil {
				repeat = *fd.StringRegexExpressionRepeatMaxCount
			}
			stringGenerator := fd.getStringGenerator(seed, repeat)

			sgr, err := stringGenerator.Generate(*fd.StringRegexExpression)
			if err != nil {
				return nil, err
			}

			res[fd.Name] = sgr
		} else if fd.RandomStringSets != nil {
			res[fd.Name] = fd.RandomStringSets[rand.Intn(len(fd.RandomStringSets))]
		} else {
			res[fd.Name] = ""
		}
	}

	return res, nil
}

func (fd *FieldDescribe) getStringGenerator(seed int64, repeat int) *regexp_trans.Generator {
	return regexp_trans.NewGenerator(regexp_trans.SetSeed(seed),
		regexp_trans.SetDefaultMaxRepeatCount(repeat))
}

func (fd *FieldDescribe) forceExpression(typ, expression string, seed int64, repeat int) (interface{}, error) {
	if len(expression) == 0 {
		switch typ {
		case FieldTypeInt, FieldTypeFloat:
			return 0, nil
		case FieldTypeBool:
			return false, nil
		case FieldTypeString:
			return "", nil
		default:
			return nil, nil
		}
	}

	g := fd.getStringGenerator(seed, repeat)

	value, err := g.Generate(expression)

	if err != nil {
		return nil, err
	}

	switch typ {
	case FieldTypeInt:
		return strconv.ParseInt(value, 10, 64)
	case FieldTypeFloat:
		return strconv.ParseFloat(value, 64)
	case FieldTypeBool:
		return strconv.ParseBool(value)
	}

	return value, nil
}
