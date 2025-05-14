package utils

import (
	"fmt"
	"math"
	"sync"
)

func DivideBy2(matrix [][]float64) [][]float64 {
	result := make([][]float64, len(matrix))

	for rowIdx := 0; rowIdx < len(matrix); rowIdx++ {
		result[rowIdx] = make([]float64, len(matrix[rowIdx]))
		for colIdx := 0; colIdx < len(matrix[rowIdx]); colIdx++ {
			operationResult := matrix[rowIdx][colIdx] / 2
			result[rowIdx][colIdx] = operationResult
		}
	}
	return result
}

func Sqrt(m [][]float64) [][]float64 {
	result := make([][]float64, len(m))
	for rowIdx := 0; rowIdx < len(m); rowIdx++ {
		result[rowIdx] = make([]float64, len(m[rowIdx]))
		for colIdx := 0; colIdx < len(m[rowIdx]); colIdx++ {
			value := m[rowIdx][colIdx]
			if math.IsNaN(value) {
				result[rowIdx][colIdx] = math.NaN()
				continue
			}
			operationResult := math.Sqrt(value) + value/2 + math.Pow(value, 2) + math.Cos(math.Mod(value, 360))
			result[rowIdx][colIdx] = operationResult
		}
	}
	return result
}

func DivideBy2Chunking(m [][]float64) [][]float64 {
	// Determine optimal number of goroutines based on available CPU cores
	numWorkers := 4 //runtime.NumCPU()
	rowsPerWorker := len(m) / numWorkers
	fmt.Println(fmt.Sprintf("Processing data using %d workers and %d rows per worker", numWorkers, rowsPerWorker))

	newMatrix := make([][]float64, len(m))

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
				newRowData := make([]float64, len(m[i]))
				for col := 0; col < len(m[i]); col++ {
					x := m[i][col]
					if math.IsNaN(x) {
						newRowData[col] = math.NaN()
						continue
					}
					operationResult := x / 2
					newRowData[col] = operationResult
				}
				newMatrix[i] = newRowData
			}
		}(startRow, endRow)
	}
	wg.Wait()
	return newMatrix
}

func SqrtChunking(m [][]float64) [][]float64 {
	// Determine optimal number of goroutines based on available CPU cores
	numWorkers := 4 //runtime.NumCPU()
	rowsPerWorker := len(m) / numWorkers
	fmt.Println(fmt.Sprintf("Processing data using %d workers and %d rows per worker", numWorkers, rowsPerWorker))

	newMatrix := make([][]float64, len(m))

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
				newRowData := make([]float64, len(m[i]))
				for col := 0; col < len(m[i]); col++ {
					x := m[i][col]
					if math.IsNaN(x) {
						newRowData[col] = math.NaN()
						continue
					}
					operationResult := math.Sqrt(x) + x/2 + math.Pow(x, 2) + math.Cos(math.Mod(x, 360))
					newRowData[col] = operationResult
				}
				newMatrix[i] = newRowData
			}
		}(startRow, endRow)
	}
	wg.Wait()
	return newMatrix
}
