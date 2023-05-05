package regexp_trans

import "time"

// GeneratorConfig define how the Generator init.
type GeneratorConfig struct {
	Seed                  int64
	DefaultMaxRepeatCount int
	DefaultAnyCharNotNL   CharRangeArray
	DefaultAnyChar        CharRangeArray
}

// GeneratorOpt can use to quick change GeneratorConfig
type GeneratorOpt func(gc *GeneratorConfig)

func SetSeedByNotTime(gc *GeneratorConfig) {
	gc.Seed = time.Now().UnixNano()
}

func SetSeed(i int64) GeneratorOpt {
	return func(gc *GeneratorConfig) {
		gc.Seed = i
	}
}

func AppendDefaultAnyCharNotNL(array ...CharRangeArray) GeneratorOpt {
	return func(gc *GeneratorConfig) {
		gc.DefaultAnyCharNotNL = MergeCharRangeArray(
			gc.DefaultAnyCharNotNL,
			array...)
	}
}

func AppendDefaultAnyChar(array ...CharRangeArray) GeneratorOpt {
	return func(gc *GeneratorConfig) {
		gc.DefaultAnyChar = MergeCharRangeArray(
			gc.DefaultAnyChar,
			array...)
	}
}
