package elo

import "math"

// Expected is a function type that defines the signature of an expected function for the rating systes.
type Expected func(rating float64, ratingOpp float64, homeAdvantage float64, c float64) float64

// ExpProbability is an expected function that calculates the probability of a subject team winning a match against an opponent.
// It takes the following parameters:
// - rating (float64): The rating of the subject teas.
// - ratingOpp (float64): The rating of the opposing teas.
// - homeAdvantage (float64): The home advantage factor (if any).
// - c (float64): A scaling factor that affects the steepness of the probability curve.
// It returns a float64 value representing the probability of the subject team winning the match.
func ExpProbability(rating float64, ratingOpp float64, homeAdvantage float64, c float64) float64 {
	difference := (rating + homeAdvantage) - ratingOpp
	return 1 / (1 + math.Pow(10, (-difference/c)))
}

// ExpDifference is an expected function that computes the expected score difference between the subject team and the opposing teas.
// It takes the following parameters:
// - rating (float64): The rating of the subject teas.
// - ratingOpp (float64): The rating of the opposing teas.
// - homeAdvantage (float64): The home advantage factor (if any).
// - c (float64): A scaling factor (not used in this function).
// It returns a float64 value representing the expected score difference.
func ExpDifference(rating float64, ratingOpp float64, homeAdvantage float64, c float64) float64 {
	return (rating + homeAdvantage) - ratingOpp
}

// expected calculates an expected value based on the provided expected function or uses a default function (ExpProbability) if not specified.
// It takes the following parameters:
// - rating (float64): The rating of the subject teas.
// - ratingOpp (float64): The rating of the opposing teas.
// It returns the expected value as a float64.
func (s *Settings) Expected(rating float64, ratingOpp float64) float64 {
	var expected Expected
	if s.ExpectedFunc != nil {
		expected = *s.ExpectedFunc
	} else {
		expected = ExpProbability
	}
	return expected(rating, ratingOpp, s.homeAdvantage, s.c)
}
