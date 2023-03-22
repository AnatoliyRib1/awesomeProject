package greetings

import (
	"fmt"
)

var divByZero = fmt.Errorf("На ноль делить нельзя")

func NotDivByZero(a, b int) (int, error) {
	if b == 0 {
		return 0, divByZero
	} else {
		return a / b, nil
	}

}
