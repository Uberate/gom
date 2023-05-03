package regexp_trans

import "sort"

type CharRange [2]rune
type CharRangeArray []CharRange

var (
	CharClassRangeNumbers      = []CharRange{[2]rune{'0', '9'}}
	CharClassRangeLowerLetters = []CharRange{[2]rune{'a', 'z'}}
	CharClassRangeUpperLetters = []CharRange{[2]rune{'A', 'Z'}}
)

func MergeCharRangeArray(a, b CharRangeArray) CharRangeArray {

	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	// may be need big memory
	mergeAB := append(a, b...)
	sort.Slice(mergeAB, func(i, j int) bool {
		if mergeAB[i][0] < mergeAB[j][0] {
			return true
		} else if mergeAB[i][0] == mergeAB[j][0] {
			return mergeAB[i][1] < mergeAB[j][1]
		}
		return false
	})

	res := CharRangeArray{}
	for index, resIndex := 0, 0; index < len(mergeAB); index++ {
		if len(res) == 0 {
			res = append(res, mergeAB[0])
			continue
		}

		if res[resIndex][1] >= mergeAB[index][1] && res[resIndex][0] <= mergeAB[index][0] {
			continue
		}

		if res[resIndex][1] >= mergeAB[index][0] {
			res[resIndex][1] = mergeAB[index][1]
			continue
		}

		res = append(res, mergeAB[index])
		resIndex++
	}

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
