package test

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("init Test.go")
}

func FirstFunction(text string) int {
	return len(text)
}

func SecondFunction(text string) (string, int) {
	return strings.ToUpper(text), len(text)
}
