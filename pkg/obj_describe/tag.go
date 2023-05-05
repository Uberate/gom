package obj_describe

const (
	MinTag                   = "min"         // for FieldTypeInt, FieldTypeFloat, FieldTypeArray
	MaxTag                   = "max"         // for FieldTypeInt, FieldTypeFloat, FieldTypeArray
	REnableTag               = "renable"     // enable the random string regex parser, and set value by it.
	StringRegexExpressionTag = "expression"  // enable the expression to generate
	ExpressionRepeatCountTag = "rm-count"    // the expression repeat max count(use {n}, n can bigger than rm-count).
	BoolWeightTag            = "true-weight" // default 0.5
	NilWeight                = "nil-weight"  // default zero
)
