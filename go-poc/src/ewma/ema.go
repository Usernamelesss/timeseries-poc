package ewma

import "math"

// EWMACalculator implements exponential weighted moving average calculation
// with parameters matching pandas ewm() function
type EWMACalculator struct {
	Alpha         float64
	Adjust        bool
	MinPeriods    int
	IsInitialized bool
	Count         int
	Mean          float64
}

// NewEWMACalculator creates a new EWMA calculator with specified parameters
func NewEWMACalculator(alpha float64, adjust bool, minPeriods int) *EWMACalculator {
	return &EWMACalculator{
		Alpha:         alpha,
		Adjust:        adjust,
		MinPeriods:    minPeriods,
		IsInitialized: false,
		Count:         0,
		Mean:          0.0,
	}
}

// NewEWMACalculatorFromSpan creates a new EWMA calculator using span instead of alpha
// This matches pandas ewm(span=x) functionality
func NewEWMACalculatorFromSpan(span float64, adjust bool, minPeriods int) *EWMACalculator {
	alpha := 2.0 / (span + 1.0)
	return NewEWMACalculator(alpha, adjust, minPeriods)
}

// Update processes a single value and returns the current EWMA
func (e *EWMACalculator) Update(value float64) float64 {
	// Skip NaN values
	if math.IsNaN(value) {
		if e.Count >= e.MinPeriods {
			return e.Mean
		}
		return math.NaN()
	}

	// Increment observation counter
	e.Count++

	// If this is the first valid observation, use it to initialize
	if !e.IsInitialized {
		e.Mean = value
		e.IsInitialized = true
	} else {
		// Calculate the decay factor
		alpha := e.Alpha
		if e.Adjust {
			// If adjust=True, the weights are normalized
			alpha = e.Alpha / (1.0 - math.Pow(1.0-e.Alpha, float64(e.Count)))
		}

		// Update the EWMA
		e.Mean = alpha*value + (1.0-alpha)*e.Mean
	}

	// Return NaN until we have min_periods observations
	if e.Count < e.MinPeriods {
		return math.NaN()
	}

	return e.Mean
}

// ProcessSeries calculates EWMA for an entire series of values
func ProcessSeries(values []float64, span float64, adjust bool, minPeriods int) []float64 {
	calculator := NewEWMACalculatorFromSpan(span, adjust, minPeriods)
	result := make([]float64, len(values))

	for i, value := range values {
		result[i] = calculator.Update(value)
	}

	return result
}

// ProcessDataFrame calculates EWMA for each column in a DataFrame-like structure
// where data is represented as [][]float64 (rows of values)
func ProcessDataFrame(data [][]float64, span float64, adjust bool, minPeriods int) [][]float64 {
	if len(data) == 0 {
		return [][]float64{}
	}

	numRows := len(data)
	numCols := len(data[0])

	// Create a result structure with the same dimensions
	result := make([][]float64, numRows)
	for i := range result {
		result[i] = make([]float64, numCols)
	}

	// Create a calculator for each column
	calculators := make([]*EWMACalculator, numCols)
	for col := 0; col < numCols; col++ {
		calculators[col] = NewEWMACalculatorFromSpan(span, adjust, minPeriods)
	}

	// Process each row
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			result[row][col] = calculators[col].Update(data[row][col])
		}
	}

	return result
}
