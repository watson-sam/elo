package elo_test

import (
	"elo"
	"testing"
)

func TestObsWinLooseDraw(t *testing.T) {
	// Test case 1: Subject team wins
	score := 2.0
	scoreOpp := 1.0
	result := elo.ObsWinLooseDraw(score, scoreOpp)
	expectedResult := 1.0
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 2: Subject team loses
	score = 1.0
	scoreOpp = 2.0
	result = elo.ObsWinLooseDraw(score, scoreOpp)
	expectedResult = 0.0
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 3: Match ends in a draw
	score = 1.0
	scoreOpp = 1.0
	result = elo.ObsWinLooseDraw(score, scoreOpp)
	expectedResult = 0.5
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}
}

func TestObsContinuous(t *testing.T) {
	// Test case 1: Subject team's performance is better
	score := 2.0
	scoreOpp := 1.0
	result := elo.ObsContinuous(score, scoreOpp)
	expectedResult := 0.6
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 2: Subject team's performance is worse
	score = 1.0
	scoreOpp = 2.0
	result = elo.ObsContinuous(score, scoreOpp)
	expectedResult = 0.4
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 3: Subject team's performance is same
	score = 1.0
	scoreOpp = 1.0
	result = elo.ObsContinuous(score, scoreOpp)
	expectedResult = 0.5
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}
}

func TestObsDifference(t *testing.T) {
	// Test case 1: Subject team has a higher score
	score := 2.0
	scoreOpp := 1.0
	result := elo.ObsDifference(score, scoreOpp)
	expectedResult := 1.0
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 2: Subject team has a lower score
	score = 1.0
	scoreOpp = 2.0
	result = elo.ObsDifference(score, scoreOpp)
	expectedResult = -1.0
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 3: Subject team has same score
	score = 1.0
	scoreOpp = 1.0
	result = elo.ObsDifference(score, scoreOpp)
	expectedResult = 0
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}
}
