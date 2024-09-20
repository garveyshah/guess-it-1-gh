package stats

import "math"

// Guesses the range of the next number
func Guesser(data []float64) (int, int) {
	mean := Mean(data)
	stdDev := StandardDev(data)

	// Define prediction interval base Mean and Standard Deviation
	lowerB := int(math.Round(mean - 2*stdDev))
	upperB := int(math.Round(mean + 2*stdDev))
	return lowerB, upperB
}

// Calculates and returns the average of the numbers
func Mean(data []float64) float64 {
	tnum := 0.0
	count := float64(len(data))

	for _, num := range data {
		tnum += num
	}
	Averag := tnum / count
	return Averag
}

// Calculates the standard deviation and variance of a set of data
func StandardDev(data []float64) float64 {
	mean := Mean(data)

	var ss, variance float64

	// calculate standard deviation and variance
	count := len(data)
	for _, value := range data {
		ss += ((value - mean) * (value - mean))
	}

	variance = ss / float64(count)

	standardDev := math.Sqrt(variance)
	return standardDev
}
