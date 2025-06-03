package main

import (
	"go-poc/src/ewma"
	"go-poc/src/utils"
	"os"
)

func main() {
	data := utils.ReadParquet()
	writeDf := os.Getenv("WRITE_RESULT_PARQUET")

	/*
		Is it fair or not including in the benchmark serialization/deserialization between TimeSeries <--> [][]*float64 ?
	*/

	matrix := data.Matrix

	r1, elapsed1 := utils.Bench("Divide By 2", func() [][]float64 {
		return utils.DivideBy2(matrix)
	})
	r2, elapsed2 := utils.Bench("Sqrt", func() [][]float64 {
		return utils.Sqrt(matrix)
	})
	//utils.Bench("Divide by 2 chunking", func() [][]float64 {
	//	return utils.DivideBy2Chunking(matrix)
	//})
	//utils.Bench("Sqrt chunking", func() [][]float64 {
	//	return utils.SqrtChunking(matrix)
	//})
	r3, elapsed3 := utils.Bench("EWMA", func() [][]float64 {
		return ewma.ProcessDataFrame(matrix, 10, false, 10)
	})

	if writeDf == "true" {
		utils.WriteTimeseries("golang_divide_by2.parquet", utils.FromMatrix(r1, data.Index))
		utils.WriteTimeseries("golang_sqrt.parquet", utils.FromMatrix(r2, data.Index))
		utils.WriteTimeseries("golang_ema.parquet", utils.FromMatrix(r3, data.Index))
	}

	utils.WriteTimings(elapsed1, elapsed2, elapsed3)
}
