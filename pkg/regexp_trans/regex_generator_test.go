package regexp_trans

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {
	phone := "1(3[0-9]|4[01456879]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\\d{8}"
	telephone := "(0\\d{2,3})-?(\\d{7,8})"
	name := "[\u4e00-\u9fa5]{2,4}"
	email := "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	idCard := "[1-9]\\d{5}(19|20)\\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]"

	fmt.Println("--- phone")
	tryMatch(10, phone)
	fmt.Println("--- name")
	tryMatch(10, name)
	fmt.Println("--- telephone")
	tryMatch(10, telephone)
	fmt.Println("--- email")
	tryMatch(10, email)
	fmt.Println("--- idCard")
	tryMatch(10, idCard)
}

func tryMatch(count int, value string) {
	for count > 0 {
		r, err := regexp.Compile(value)
		if err != nil {
			fmt.Println(err)
		}

		p := NewGenerator()

		res, _ := p.Generate(value)
		fmt.Print(res, " ")

		fmt.Println(r.MatchString(res))

		count--
	}

}
