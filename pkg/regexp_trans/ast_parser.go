package regexp_trans

import (
	"bytes"
	"fmt"
	"regexp/syntax"
)

type Parser struct {
}

func (p *Parser) Parse(value string) (string, error) {
	regexpAst, err := syntax.Parse(value, syntax.PerlX)
	if err != nil {
		return "", err
	}

	res, err := p.parse(regexpAst)
	return string(res), err
}

func (p *Parser) parse(regexp *syntax.Regexp) ([]byte, error) {
	if regexp == nil {
		return []byte{}, nil
	}

	res := bytes.Buffer{}

	switch regexp.Op {
	case syntax.OpConcat:
		for _, item := range regexp.Sub {
			r, err := p.parse(item)
			if err != nil {
				return []byte{}, nil
			}

			res.Write(r)
		}
	case syntax.OpCharClass:

	}

	panic("TODO")
}

func (p *Parser) parseCharClass(runes []rune) ([]byte, error) {
	if len(runes)%2 != 0 {
		return nil, fmt.Errorf("error char class")
	}

	panic("TODO")
}
