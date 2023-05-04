package regexp_trans

import (
	"bytes"
	"math/rand"
	"regexp/syntax"
	"time"
)

type Generator struct {
	ran *rand.Rand

	defaultMaxRepeatCount int
	defaultAnyCharNotNL   CharRangeArray
	defaultAnyChar        CharRangeArray
	defaultWordBoundary   CharRangeArray
	defaultNoWordBoundary CharRangeArray
}

func NewGenerator() *Generator {
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &Generator{
		ran:                   ran,
		defaultMaxRepeatCount: 10,
		defaultAnyCharNotNL:   CharClassRangeLettersAndNumbers,
		defaultAnyChar:        MergeCharRangeArray(CharClassRangeLettersAndNumbers, CharClassNewLineLetter),
		defaultWordBoundary:   CharClassRangeLettersAndNumbers,
		defaultNoWordBoundary: CharClassEmptyChars,
	}
}

func (p *Generator) Generate(value string) (string, error) {
	regexpAst, err := syntax.Parse(value, syntax.PerlX)
	if err != nil {
		return "", err
	}

	res, err := p.parse(regexpAst)
	return string(res), err
}

func (p *Generator) parse(regexp *syntax.Regexp) ([]byte, error) {
	if regexp == nil {
		return nil, nil
	}
	res := bytes.Buffer{}
	switch regexp.Op {
	case syntax.OpLiteral:
		for _, item := range regexp.Rune {
			res.WriteRune(item)
		}
	case syntax.OpCharClass:
		genValue, err := p.generateCharClass(regexp.Rune)
		if err != nil {
			return nil, err
		}

		res.Write(genValue)
	case syntax.OpAnyCharNotNL:
		r, err := p.generateCharRunArray(p.defaultAnyCharNotNL)
		if err != nil {
			return nil, err
		}
		res.Write(r)
	case syntax.OpAnyChar:
		r, err := p.generateCharRunArray(p.defaultAnyChar)
		if err != nil {
			return nil, err
		}
		res.Write(r)
	case syntax.OpWordBoundary:
		// do nothing
	case syntax.OpNoWordBoundary:
		// do nothing
	case syntax.OpCapture:
		for _, sub := range regexp.Sub {
			r, err := p.parse(sub)
			if err != nil {
				return nil, err
			}
			res.Write(r)
		}
	case syntax.OpStar: // *
		r, err := p.generateRepeats(regexp.Sub0[0], p.defaultMaxRepeatCount, regexp.Min)
		if err != nil {
			return nil, err
		}
		res.Write(r)
	case syntax.OpPlus: // +
		r, err := p.generateRepeats(regexp.Sub0[0], p.defaultMaxRepeatCount, 1)
		if err != nil {
			return nil, err
		}
		res.Write(r)
	case syntax.OpQuest: // ?
		r, err := p.generateRepeats(regexp.Sub0[0], 1, 0)
		if err != nil {
			return nil, err
		}
		res.Write(r)
	case syntax.OpRepeat: //{}
		r, err := p.generateRepeats(regexp.Sub0[0], regexp.Max, regexp.Min)
		if err != nil {
			return nil, err
		}
		res.Write(r)
	case syntax.OpConcat:
		for _, item := range regexp.Sub {
			r, err := p.parse(item)
			if err != nil {
				return nil, nil
			}
			res.Write(r)
		}
	case syntax.OpAlternate:
		index := p.ran.Intn(len(regexp.Sub))
		r, err := p.parse(regexp.Sub[index])
		if err != nil {
			return nil, err
		}
		res.Write(r)
	}

	return res.Bytes(), nil
}

func (p *Generator) generateRepeats(sub *syntax.Regexp, max, min int) ([]byte, error) {
	ranRepeatCount := p.ran.Intn(max-min+1) + min
	res := bytes.Buffer{}
	for ; ranRepeatCount > 0; ranRepeatCount-- {
		r, err := p.parse(sub)
		if err != nil {
			return nil, err
		}
		res.Write(r)
	}
	return res.Bytes(), nil
}

func (p *Generator) generateCharClass(runes []rune) ([]byte, error) {
	chars, err := ParseCharRangeArray(runes)
	if err != nil {
		return nil, err
	}

	return p.generateCharRunArray(chars)
}

func (p *Generator) generateCharRunArray(array CharRangeArray) ([]byte, error) {
	res := bytes.Buffer{}

	for _, item := range RandomRangeChar(p.ran, array, 1) {
		res.WriteRune(item)
	}

	return res.Bytes(), nil
}
