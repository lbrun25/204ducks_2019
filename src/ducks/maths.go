package ducks

import (
    "math"
)

func computeProbabilityDensityPercent(t float64) float64 {
    res := -A * math.Exp(-t) - (4 - 3 * A) / 2 * math.Exp(-2 * t) - (2 * A - 4) / 4 * math.Exp(-4 * t)
    return res
}

func computeProbabilityDensity(t float64) float64 {
    res := A * math.Exp(-t) + (4 - 3 * A) * math.Exp(-2 * t) + (2 * A - 4) * math.Exp(-4 * t)
    return res
}

// ComputeExpectedValue - expected value
func ComputeExpectedValue() float64 {
    var res float64 = 0.0

    for t := 0.0; t < 100; t += Increment {
        res += computeProbabilityDensity(t) * t / 1000
    }
    return res
}

// ComputeStandardDeviation - standard deviation
func ComputeStandardDeviation(meanTime float64) float64 {
    var res float64 = 0.0

    for t := 0.0; t < 200; t += Increment {
        res += math.Pow(t - meanTime, 2) * computeProbabilityDensity(t) / 1000
    }
    res = math.Sqrt(res)
    return res
}

// ComputeTimeDucksBack - Time after which [percentage] of the ducks are back
func ComputeTimeDucksBack(percentage float64) float64 {
    for t := 0.0; t < 1000; t += (Increment * 2)  {
        if computeProbabilityDensityPercent(t / 60) - computeProbabilityDensityPercent(0) >= percentage / 100 {
            return t
        }
    }
    panic("Error to compute the time after which x% of the ducks are back")
}

// ComputePercentageDucksBack - Percentage of ducks back after x minutes:
func ComputePercentageDucksBack(min float64) float64 {
    res := computeProbabilityDensityPercent(min) - computeProbabilityDensityPercent(0)
    res *= 100
    return res
}