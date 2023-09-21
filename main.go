package elo

import "math"

func ExpectedWinProbability(ratingA float64, ratingB float64, homeAdvatange float64, c float64) float64 {
	difference := (ratingA + homeAdvatange) - ratingB
	return 1 / (1 + math.Pow(10, (-difference/c)))
}
