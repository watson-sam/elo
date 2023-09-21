package elo_test

import (
	"elo"
	"math"
	"testing"
)

func TestExpProbability(t *testing.T) {
	// Test case 1: Subject team is stronger, no home advantage
	rating := 1500.0
	ratingOpp := 1400.0
	homeAdvantage := 0.0
	c := 200.0
	result := elo.ExpProbability(rating, ratingOpp, homeAdvantage, c)
	expectedResult := 0.759747
	if math.Abs(result-expectedResult) > 0.0001 {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 2: Subject team is weaker, with home advantage
	rating = 1400.0
	ratingOpp = 1500.0
	homeAdvantage = 50.0
	c = 200.0
	result = elo.ExpProbability(rating, ratingOpp, homeAdvantage, c)
	expectedResult = 0.3599350003907806
	if math.Abs(result-expectedResult) > 0.0001 {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}
}

func TestExpDifference(t *testing.T) {
	// Test case 1: Subject team is stronger, no home advantage
	rating := 1500.0
	ratingOpp := 1400.0
	homeAdvantage := 0.0
	c := 200.0
	result := elo.ExpDifference(rating, ratingOpp, homeAdvantage, c)
	expectedResult := 100.0
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 2: Subject team is weaker, with home advantage
	rating = 1400.0
	ratingOpp = 1500.0
	homeAdvantage = 50.0
	c = 200.0
	result = elo.ExpDifference(rating, ratingOpp, homeAdvantage, c)
	expectedResult = -50.0
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}
}
