package continuous

import (
    "math"
)

// ComputeProbabilityDensity - probabilty density
func ComputeProbabilityDensity(a float64, t float64) float64 {
    res := a * math.Exp(-t) + (4 - 3 * a) * math.Exp(-2 * t) + (2 * a - 4) * math.Exp(-4 * t)
    return res
}

// ComputeExpectedValue - expected value
func ComputeExpectedValue(value float64) float64 {
    var res float64 = 0.0

    for t := 0.0; t < 50; t += 0.001 {
        res += ComputeProbabilityDensity(value, t) * t / 1000
    }
    return res
}

// ComputeAverageReturnTime - average return time
func ComputeAverageReturnTime() {
	
}

// ComputeVariance - variance
func ComputeVariance() {
    
}

// ComputeStandardDeviation - standard deviation
func ComputeStandardDeviation() float64 {
    return 0.0
}

// ComputeTimeDucksBack - Time after which [percentage] of the ducks are back
func ComputeTimeDucksBack(percentage int) string {
    return ""
}

// ComputePercentageDucksBack - Percentage of ducks back after x minutes:
func ComputePercentageDucksBack(min int) float64 {
    res := 0.0
    
    return res
}