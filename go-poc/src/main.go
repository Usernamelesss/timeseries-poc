package main

import "go-poc/src/utils"

func main() {
	data := utils.GetTimeseries()

	utils.BenchDivideBy2(data)
	utils.BenchSqrt(data)
	utils.BenchDivideBy2Chunking(data)
	utils.BenchSqrtChunking(data)
}
