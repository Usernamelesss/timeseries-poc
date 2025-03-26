package utils

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func BenchDivideBy2(data TimeSeries) {
	result := make([][]float64, len(data))

	start := time.Now()
	m := data.ToMatrix()
	for rowIdx := 0; rowIdx < len(m); rowIdx++ {
		result[rowIdx] = make([]float64, len(m[rowIdx]))
		for colIdx := 0; colIdx < len(m[rowIdx]); colIdx++ {
			result[rowIdx][colIdx] = m[rowIdx][colIdx] / 2
		}
	}
	fmt.Println(fmt.Sprintf("Took %v", time.Since(start)))
}

func BenchSqrt(data TimeSeries) {
	result := make([][]float64, len(data))

	start := time.Now()
	m := data.ToMatrix()
	for rowIdx := 0; rowIdx < len(m); rowIdx++ {
		result[rowIdx] = make([]float64, len(m[rowIdx]))
		for colIdx := 0; colIdx < len(m[rowIdx]); colIdx++ {
			value := m[rowIdx][colIdx]
			result[rowIdx][colIdx] = math.Sqrt(value) + value/2 + math.Pow(value, 2) + math.Cos(math.Mod(value, 360))
		}
	}
	fmt.Println(fmt.Sprintf("Took %v", time.Since(start)))
}

func BenchDivideBy2Chunking(data TimeSeries) {
	start := time.Now()

	// Determine optimal number of goroutines based on available CPU cores
	numWorkers := 4 //runtime.NumCPU()
	rowsPerWorker := len(data) / numWorkers

	fmt.Println(fmt.Sprintf("Processing data using %d workers", numWorkers))
	matrix := data.ToMatrix()
	newMatrix := make([][]float64, len(data))
	var wg sync.WaitGroup

	// Process data in parallel
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)

		startRow := w * rowsPerWorker
		endRow := startRow + rowsPerWorker

		if w == numWorkers-1 {
			endRow = len(data) // Handle remaining rows in last worker
		}

		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				newRowData := make([]float64, len(matrix[i]))
				for col := 0; col < len(matrix[i]); col++ {
					x := matrix[i][col]
					newRowData[col] = x / 2
					//newRowData[col] = math.Sqrt(x) + x/2 + math.Pow(x, 2) + math.Cos(math.Mod(x, 360))
				}
				newMatrix[i] = newRowData
			}
		}(startRow, endRow)
	}

	wg.Wait()
	fmt.Println(fmt.Sprintf("Took %v", time.Since(start)))
}

func BenchSqrtChunking(data TimeSeries) {
	start := time.Now()

	// Determine optimal number of goroutines based on available CPU cores
	numWorkers := 4 //runtime.NumCPU()
	rowsPerWorker := len(data) / numWorkers

	fmt.Println(fmt.Sprintf("Processing data using %d workers", numWorkers))
	matrix := data.ToMatrix()
	newMatrix := make([][]float64, len(data))
	var wg sync.WaitGroup

	// Process data in parallel
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)

		startRow := w * rowsPerWorker
		endRow := startRow + rowsPerWorker

		if w == numWorkers-1 {
			endRow = len(data) // Handle remaining rows in last worker
		}

		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				newRowData := make([]float64, len(matrix[i]))
				for col := 0; col < len(matrix[i]); col++ {
					x := matrix[i][col]
					newRowData[col] = math.Sqrt(x) + x/2 + math.Pow(x, 2) + math.Cos(math.Mod(x, 360))
				}
				newMatrix[i] = newRowData
			}
		}(startRow, endRow)
	}

	wg.Wait()
	fmt.Println(fmt.Sprintf("Took %v", time.Since(start)))
}
