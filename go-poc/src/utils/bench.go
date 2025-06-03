package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Bench(name string, function func() [][]float64) ([][]float64, int64) {
	start := time.Now()
	result := function()
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("%s took %v", name, elapsed))
	return result, elapsed.Milliseconds()
}

func WriteTimings(elapsed ...int64) {
	pathToFile := fmt.Sprintf("%s/results/golang_timing.csv", os.Getenv("PROJECT_ROOT"))

	fileExists := false
	if _, err := os.Stat(pathToFile); err == nil {
		fileExists = true
	}

	file, err := os.OpenFile(pathToFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if !fileExists {
		if err = writer.Write([]string{"divide_by_2[ms]", "sqrt[ms]", "ewma[ms]"}); err != nil {
			fmt.Println("Error writing headers:", err)
			return
		}
	}

	err = writer.Write([]string{
		strconv.FormatInt(elapsed[0], 10),
		strconv.FormatInt(elapsed[1], 10),
		strconv.FormatInt(elapsed[2], 10),
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to write data: %v", err))
	}

	writer.Flush()
}
