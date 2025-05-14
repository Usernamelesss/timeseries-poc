package utils

import (
	"fmt"
	"github.com/parquet-go/parquet-go"
	"os"
	"time"
)

type DataRow struct {
	Index   time.Time `parquet:"__index_level_0__"`
	Param1  float64   `parquet:"param_0"`
	Param2  float64   `parquet:"param_1"`
	Param3  float64   `parquet:"param_2"`
	Param4  float64   `parquet:"param_3"`
	Param5  float64   `parquet:"param_4"`
	Param6  float64   `parquet:"param_5"`
	Param7  float64   `parquet:"param_6"`
	Param8  float64   `parquet:"param_7"`
	Param9  float64   `parquet:"param_8"`
	Param10 float64   `parquet:"param_9"`
}

type TimeSeries []DataRow

func (s *TimeSeries) ToMatrix() [][]float64 {
	if s == nil {
		return make([][]float64, 0)
	}

	matrix := make([][]float64, len(*s))
	for i, row := range *s {
		// Extract values from struct into a flat array
		rowValues := []float64{
			row.Param1, row.Param2, row.Param3, row.Param4, row.Param5,
			row.Param6, row.Param7, row.Param8, row.Param9, row.Param10,
		}
		matrix[i] = rowValues
	}

	return matrix
}

func (s *TimeSeries) GetIndex() []time.Time {
	if s == nil {
		return make([]time.Time, 0)
	}

	index := make([]time.Time, len(*s))
	for i, row := range *s {
		index[i] = row.Index
	}

	return index
}

func FromMatrix(matrix [][]float64, index []time.Time) TimeSeries {
	if matrix == nil {
		return TimeSeries{}
	}

	ts := make([]DataRow, len(matrix))
	for i, row := range matrix {
		ts[i] = DataRow{
			Index:   index[i],
			Param1:  row[0],
			Param2:  row[1],
			Param3:  row[2],
			Param4:  row[3],
			Param5:  row[4],
			Param6:  row[5],
			Param7:  row[6],
			Param8:  row[7],
			Param9:  row[8],
			Param10: row[9],
		}
	}

	return ts
}

func GetTimeseries() TimeSeries {
	// Open the Parquet file
	filename := fmt.Sprintf("%s/fixtures/sample_001.parquet", os.Getenv("PROJECT_ROOT"))

	data, err := parquet.ReadFile[DataRow](filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to read data: %v", err))
	}

	fmt.Println(fmt.Sprintf("Loaded %d rows from parquet file", len(data)))

	return data
}

func WriteTimeseries(filename string, ts TimeSeries) {
	pathToFile := fmt.Sprintf("%s/results/%s", os.Getenv("PROJECT_ROOT"), filename)
	err := parquet.WriteFile[DataRow](pathToFile, ts)
	if err != nil {
		panic(fmt.Sprintf("Failed to write data: %v", err))
	}
}
