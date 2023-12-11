package script

import (
	"fmt"
	"math"
	"testing"
)

func Test_luaRunner_AnalysisRx(t *testing.T) {
	for i := 3000; i < 4000; i++ {
		d := float64(i)
		fmt.Println(aaa(d))
	}
}

func aaa(value interface{}) float64 {
	result := value.(float64) / math.Pow10(2)
	return result
}

func Test_Type(t *testing.T) {
	var a = 100.0
	b := convert(a)
	fmt.Println(b)
}

func convert(value interface{}) float64 {
	if result, ok := value.(float64); ok {
		return result
	}

	return 0
}
