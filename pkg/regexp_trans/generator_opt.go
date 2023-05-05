package regexp_trans

// GeneratorConfig define how the Generator init.
type GeneratorConfig struct {
	DefaultMaxRepeatCount int
	DefaultAnyCharNotNL   CharRangeArray
	DefaultAnyChar        CharRangeArray
	DefaultWordBoundary   CharRangeArray
	DefaultNoWordBoundary CharRangeArray
}

// GeneratorOpt can use to quick change GeneratorConfig
type GeneratorOpt func(gc *GeneratorConfig)

func AppendDefaultAnyCharNotNL(array ...CharRangeArray) GeneratorOpt {
	return func(gc *GeneratorConfig) {
		gc.DefaultAnyCharNotNL = MergeCharRangeArray(
			gc.DefaultAnyCharNotNL,
			array...)
	}
}
