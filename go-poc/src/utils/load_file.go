package utils

import (
	"fmt"
	"github.com/parquet-go/parquet-go"
	"os"
	"time"
)

type TimeSeries struct {
	Index       []time.Time
	Matrix      [][]float64
	ColumnNames []string
	NumRows     int64
}

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

func ReadParquet() *TimeSeries {
	// Open the Parquet file
	filename := fmt.Sprintf("%s/fixtures/sample_001.parquet", os.Getenv("PROJECT_ROOT"))

	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("failed to open file: %w", err))
	}
	defer file.Close()

	// Get file info to estimate capacity
	fileInfo, err := file.Stat()
	if err != nil {
		panic(fmt.Errorf("failed to get file info: %w", err))
	}

	// Create parquet file reader
	pf, err := parquet.OpenFile(file, fileInfo.Size())
	if err != nil {
		panic(fmt.Errorf("failed to open parquet file: %w", err))
	}

	// Get number of rows for pre-allocation
	numRows := pf.NumRows()

	// Pre-allocate slices for better performance
	ts := &TimeSeries{
		Index:       make([]time.Time, 0, numRows),
		Matrix:      make([][]float64, 0, numRows),
		ColumnNames: []string{"param_0", "param_1", "param_2", "param_3", "param_4", "param_5", "param_6", "param_7", "param_8", "param_9"},
		NumRows:     numRows,
	}

	// Create reader for the ParquetRow type
	reader := parquet.NewGenericReader[DataRow](pf)
	defer reader.Close()

	// Read in batches for memory efficiency
	batchSize := 10000
	batch := make([]DataRow, batchSize)

	for {
		n, err := reader.Read(batch)
		if n == 0 {
			break
		}
		if err != nil && err.Error() != "EOF" {
			panic(fmt.Errorf("failed to read batch: %w", err))
		}

		// Process the batch
		for i := 0; i < n; i++ {
			row := batch[i]

			// Add timestamp to index
			ts.Index = append(ts.Index, row.Index)

			// Create row for matrix
			matrixRow := []float64{
				row.Param1, row.Param2, row.Param3, row.Param4, row.Param5,
				row.Param6, row.Param7, row.Param8, row.Param9, row.Param10,
			}
			ts.Matrix = append(ts.Matrix, matrixRow)
		}

		if n < batchSize {
			break
		}
	}

	return ts
}

func FromMatrix(matrix [][]float64, index []time.Time) []DataRow {
	if matrix == nil {
		return make([]DataRow, 0)
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

func WriteTimeseries(filename string, ts []DataRow) {
	pathToFile := fmt.Sprintf("%s/results/%s", os.Getenv("PROJECT_ROOT"), filename)
	err := parquet.WriteFile[DataRow](pathToFile, ts)
	if err != nil {
		panic(fmt.Sprintf("Failed to write data: %v", err))
	}
}
