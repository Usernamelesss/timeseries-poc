package main

import "go-poc/src/utils"

func main() {
	data := utils.GetTimeseries()

	/*
		Is it fair or not including in the benchmark serialization/deserialization between TimeSeries <--> [][]*float64 ?
	*/

	matrix := data.ToMatrix()

	r1 := utils.Bench("Divide by 2", func() [][]*float64 {
		return utils.DivideBy2(matrix)
	})
	r2 := utils.Bench("Sqrt", func() [][]*float64 {
		return utils.Sqrt(matrix)
	})
	utils.Bench("Divide by 2 chunking", func() [][]*float64 {
		return utils.DivideBy2Chunking(matrix)
	})
	utils.Bench("Sqrt chunking", func() [][]*float64 {
		return utils.SqrtChunking(matrix)
	})

	utils.WriteTimeseries("golang_divide_by2.parquet", utils.FromMatrix(r1))
	utils.WriteTimeseries("golang_sqrt.parquet", utils.FromMatrix(r2))
}
