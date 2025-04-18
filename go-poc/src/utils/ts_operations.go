package utils

import (
	"fmt"
	"math"
	"sync"
)

func DivideBy2(matrix [][]*float64) [][]*float64 {
	result := make([][]*float64, len(matrix))

	for rowIdx := 0; rowIdx < len(matrix); rowIdx++ {
		result[rowIdx] = make([]*float64, len(matrix[rowIdx]))
		for colIdx := 0; colIdx < len(matrix[rowIdx]); colIdx++ {
			operationResult := *matrix[rowIdx][colIdx] / 2
			result[rowIdx][colIdx] = &operationResult
		}
	}
	return result
}

func Sqrt(m [][]*float64) [][]*float64 {
	result := make([][]*float64, len(m))
	for rowIdx := 0; rowIdx < len(m); rowIdx++ {
		result[rowIdx] = make([]*float64, len(m[rowIdx]))
		for colIdx := 0; colIdx < len(m[rowIdx]); colIdx++ {
			value := m[rowIdx][colIdx]
			if value == nil {
				result[rowIdx][colIdx] = nil
				continue
			}
			operationResult := math.Sqrt(*value) + *value/2 + math.Pow(*value, 2) + math.Cos(math.Mod(*value, 360))
			result[rowIdx][colIdx] = &operationResult
		}
	}
	return result
}

func DivideBy2Chunking(m [][]*float64) [][]*float64 {
	// Determine optimal number of goroutines based on available CPU cores
	numWorkers := 4 //runtime.NumCPU()
	rowsPerWorker := len(m) / numWorkers
	fmt.Println(fmt.Sprintf("Processing data using %d workers and %d rows per worker", numWorkers, rowsPerWorker))

	newMatrix := make([][]*float64, len(m))

	var wg sync.WaitGroup

	// Process data in parallel
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)

		startRow := w * rowsPerWorker
		endRow := startRow + rowsPerWorker

		if w == numWorkers-1 {
			endRow = len(m) // Handle remaining rows in last worker
		}

		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				newRowData := make([]*float64, len(m[i]))
				for col := 0; col < len(m[i]); col++ {
					x := m[i][col]
					if x == nil {
						newRowData[col] = nil
						continue
					}
					operationResult := *x / 2
					newRowData[col] = &operationResult
				}
				newMatrix[i] = newRowData
			}
		}(startRow, endRow)
	}
	wg.Wait()
	return newMatrix
}

func SqrtChunking(m [][]*float64) [][]*float64 {
	// Determine optimal number of goroutines based on available CPU cores
	numWorkers := 4 //runtime.NumCPU()
	rowsPerWorker := len(m) / numWorkers
	fmt.Println(fmt.Sprintf("Processing data using %d workers and %d rows per worker", numWorkers, rowsPerWorker))

	newMatrix := make([][]*float64, len(m))

	var wg sync.WaitGroup

	// Process data in parallel
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)

		startRow := w * rowsPerWorker
		endRow := startRow + rowsPerWorker

		if w == numWorkers-1 {
			endRow = len(m) // Handle remaining rows in last worker
		}

		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				newRowData := make([]*float64, len(m[i]))
				for col := 0; col < len(m[i]); col++ {
					x := m[i][col]
					if x == nil {
						newRowData[col] = nil
						continue
					}
					operationResult := math.Sqrt(*x) + *x/2 + math.Pow(*x, 2) + math.Cos(math.Mod(*x, 360))
					newRowData[col] = &operationResult
				}
				newMatrix[i] = newRowData
			}
		}(startRow, endRow)
	}
	wg.Wait()
	return newMatrix
}

func ExponentialMovingAverage(m [][]*float64) [][]*float64 {
	var k float64
	windowSize := 10
	k = 2.0 / float64(windowSize+1)
	var result [][]*float64

	colsNum := len(m[0]) // On tabular data the number of columns is equal for all the rows

	// Initialize the result array
	result = make([][]*float64, len(m))
	for i := range result {
		result[i] = make([]*float64, colsNum)
	}

	// Iterate one column at time and process the EMA on each column (series) separately
	for colIdx := 0; colIdx < colsNum; colIdx++ {
		var sum float64
		for rowIdx := 0; rowIdx < len(m); rowIdx++ {
			if rowIdx < windowSize {
				// This is the "warmup" period.
				// We need to compute first the average of the "windowSize" elements.
				result[rowIdx][colIdx] = nil // we don't have a value for the EMA yet
				if m[rowIdx][colIdx] != nil {
					sum += *m[rowIdx][colIdx]
				}

				if rowIdx == windowSize-1 {
					// Set the first value as a Simple Average of the previous "windowSize" elements
					average := sum / float64(windowSize)
					result[rowIdx][colIdx] = &average
				}
				continue
			}

			var value float64 // Is it acceptable to handle nulls has 0 ?
			if m[rowIdx][colIdx] != nil {
				value = *m[rowIdx][colIdx]
			}

			// Get EMA(t-1)
			var previousEmaValue float64
			if result[rowIdx-1][colIdx] != nil {
				previousEmaValue = *result[rowIdx-1][colIdx]
			}

			// Compute EMA for the current value
			currentEma := (value * k) + (previousEmaValue * (1 - k))
			// Store the result in the result matrix
			result[rowIdx][colIdx] = &currentEma
		}
	}
	return result
}
