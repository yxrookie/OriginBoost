package utils

import (
	
	"fmt"
	"testing"
)



func TestConvert(t *testing.T) {
	
	fmt.Println(Convert("你好", "zh", "en"))
}

func TestMethod(t *testing.T) {
	fmt.Println(Method2("跟大多数人相比，我的习惯性没有那么强。"))
	fmt.Println(Method7("跟大多数人相比，我的习惯性没有那么强。"))
}
