package ducks

import (
    "math"
)

func computeProbabilityDensityPercent(a float64, t float64) float64 {
    res := -a * math.Exp(-t) - (4 - 3 * a) / 2 * math.Exp(-2 * t) - (2 * a - 4) / 4 * math.Exp(-4 * t)
    return res
}

func computeProbabilityDensity(a float64, t float64) float64 {
    res := a * math.Exp(-t) + (4 - 3 * a) * math.Exp(-2 * t) + (2 * a - 4) * math.Exp(-4 * t)
    return res
}

// ComputeExpectedValue - expected value
func ComputeExpectedValue(value float64) float64 {
    var res float64 = 0.0

    for t := 0.0; t < 100; t += Increment {
        res += computeProbabilityDensity(value, t) * t / 1000
    }
    return res
}

// ComputeVariance - variance
func ComputeVariance() {
    
}

// ComputeStandardDeviation - standard deviation
func ComputeStandardDeviation(a float64, meanTime float64) float64 {
    var res float64 = 0.0

    for t := 0.0; t < 200; t += Increment {
        res += math.Pow(t - meanTime, 2) * computeProbabilityDensity(a, t) / 1000
    }
    res = math.Sqrt(res)
    return res
}

// ComputeTimeDucksBack - Time after which [percentage] of the ducks are back
func ComputeTimeDucksBack(a float64, percentage float64) float64 {
    for t := 0.0; t < 1000; t += (Increment * 2)  {
        if computeProbabilityDensityPercent(a, t / 60) - computeProbabilityDensityPercent(a, 0) >= percentage / 100 {
            return t
        }
    }
    panic("Error to compute the time after which x% of the ducks are back")
}

// ComputePercentageDucksBack - Percentage of ducks back after x minutes:
func ComputePercentageDucksBack(min int) float64 {
    res := 0.0
    
    return res
}