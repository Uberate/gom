package regexp_trans

import (
	"fmt"
	"math/rand"
	"sort"
)

type CharRange [2]rune
type CharRangeArray []CharRange

func (cr *CharRange) Clone() CharRange {
	return [2]rune{cr[0], cr[1]}
}

func (cra *CharRangeArray) Clone() CharRangeArray {
	value := make(CharRangeArray, len(*cra))
	for index := range *cra {
		value[index] = (*cra)[index].Clone()
	}

	return value
}

var (
	CharClassRangeNumbers           = []CharRange{[2]rune{'0', '9'}}
	CharClassRangeLowerLetters      = []CharRange{[2]rune{'a', 'z'}}
	CharClassRangeUpperLetters      = []CharRange{[2]rune{'A', 'Z'}}
	CharClassNewLineLetter          = []CharRange{[2]rune{'\n', '\n'}}
	CharClassRangeAllLetters        = MergeCharRangeArray(CharClassRangeLowerLetters, CharClassRangeUpperLetters)
	CharClassRangeLettersAndNumbers = MergeCharRangeArray(CharClassRangeAllLetters, CharClassRangeNumbers)

	CharClassEmptyChars = MergeCharRangeArray(
		CharRangeArray{[2]rune{' ', ' '}},
		CharRangeArray{[2]rune{'\t', '\t'}},
	)
)

// MergeCharRangeArray return a CharRangeArray, merged the at latest one slice.
func MergeCharRangeArray(a CharRangeArray, b ...CharRangeArray) CharRangeArray {

	if len(a) == 0 {
		if len(b) == 0 {
			return CharRangeArray{}
		} else {
			if len(b) == 1 {
				return b[0].Clone()
			}
			return MergeCharRangeArray(b[0].Clone(), b[1:]...)
		}
	}

	if len(b) == 0 {
		return a
	}

	// may be need big memory
	merges := a.Clone()
	for _, bItem := range b {
		merges = append(merges, bItem.Clone()...)
	}

	// sort slice: merges
	sort.Slice(merges, func(i, j int) bool {
		if merges[i][0] < merges[j][0] {
			return true
		} else if merges[i][0] == merges[j][0] {
			return merges[i][1] < merges[j][1]
		}
		return false
	})

	// merge A and B...
	res := CharRangeArray{}
	for index, resIndex := 0, 0; index < len(merges); index++ {
		if len(res) == 0 {
			res = append(res, merges[0])
			continue
		}

		if res[resIndex][1] >= merges[index][1] && res[resIndex][0] <= merges[index][0] {
			continue
		}

		if res[resIndex][1] >= merges[index][0] {
			res[resIndex][1] = merges[index][1]
			continue
		}

		res = append(res, merges[index])
		resIndex++
	}

	// compress res
	compressRes := CharRangeArray{}
	compressResIndex := 0
	for index := range res {
		if len(compressRes) == 0 {
			compressRes = append(compressRes, res[index])
			continue
		}
		if compressRes[compressResIndex][1] == res[index][0]-1 {
			compressRes[compressResIndex][1] = res[index][1]
			continue
		}

		compressRes = append(compressRes, res[index])
		compressResIndex++
	}

	return compressRes
}

// RandomRangeChar will generate specif count char in a range.
func RandomRangeChar(ran *rand.Rand, charClasses CharRangeArray, count int) []rune {
	res := make([]rune, count, count)
	totalCount := int32(0)
	var indexCache []int32
	for _, item := range charClasses {
		totalCount += (item[1] - item[0]) + 1
		indexCache = append(indexCache, totalCount)
	}
	for i := 0; i < count; i++ {
		randIndex := ran.Int31n(totalCount) + 1
		realIndex := 0
		lastValue := int32(0)
		for ; realIndex < len(indexCache); realIndex++ {
			if randIndex <= indexCache[realIndex] {
				break
			}
			lastValue = indexCache[realIndex]
		}

		if realIndex == len(indexCache) {
			realIndex--
		}
		res[i] = charClasses[realIndex][0] + (randIndex - lastValue - int32(1))
	}

	return res
}

// ParseCharRangeArray to CharRangeArray from a runes, the rs must have even counts elements.
func ParseCharRangeArray(rs []rune) (CharRangeArray, error) {
	if len(rs)%2 != 0 {
		return nil, fmt.Errorf("char range shoud an even number, but [%d]", len(rs))
	}

	res := CharRangeArray{}
	for index := 0; index <= len(rs)/2; index += 2 {
		res = append(res, [2]rune{rs[index], rs[index+1]})
	}
	return res, nil
}
