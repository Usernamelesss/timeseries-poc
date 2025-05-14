package utils

import (
	"fmt"
	"time"
)

func Bench(name string, function func() [][]float64) [][]float64 {
	start := time.Now()
	result := function()
	fmt.Println(fmt.Sprintf("%s took %v", name, time.Since(start)))
	return result
}
