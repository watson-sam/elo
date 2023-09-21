package elo_test

import (
	"elo"
	"testing"
)

var ERROR_MESSAGE string = "Expected %f, but got %f"

func TestUpdateExpected(t *testing.T) {
	// Test case 1: Positive change in rating
	observed := 1.5
	expected := 1.0
	kFactor := 0.5
	result := elo.UpdateExpected(observed, expected, kFactor)
	expectedResult := 0.25
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 2: Negative change in rating
	observed = 0.5
	expected = 1.0
	kFactor = 0.5
	result = elo.UpdateExpected(observed, expected, kFactor)
	expectedResult = -0.25
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}
}

func TestUpdatePoints(t *testing.T) {
	// Test case 1: Positive change in rating
	observed := 1.5
	expected := 1.0
	kFactor := 0.5
	result := elo.UpdatePoints(observed, expected, kFactor)
	expectedResult := 0.5
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 2: Negative change in rating
	observed = 0.5
	expected = 1.0
	kFactor = 0.5
	result = elo.UpdatePoints(observed, expected, kFactor)
	expectedResult = -0.5
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}
}

func TestApplyMaxChange(t *testing.T) {
	// Test case 1: New rating is below the minimum allowed
	minRating := 100.0
	maxRating := 200.0
	newRating := 50.0
	result := elo.ApplyMaxChange(minRating, maxRating, newRating)
	expectedResult := 100.0 // Adjusted to the minimum
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 2: New rating is above the maximum allowed
	minRating = 100.0
	maxRating = 200.0
	newRating = 250.0
	result = elo.ApplyMaxChange(minRating, maxRating, newRating)
	expectedResult = 200.0 // Adjusted to the maximum
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 3: New rating is within the allowed range
	minRating = 100.0
	maxRating = 200.0
	newRating = 150.0
	result = elo.ApplyMaxChange(minRating, maxRating, newRating)
	expectedResult = 150.0 // No adjustment needed
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}
}

func TestSettingsApplyMaxChangePerc(t *testing.T) {
	// Create a sample Settings configuration with maxChangePerc

	settings := elo.New(elo.WithMaxChangePerc(0.2))

	// Test case 1: New rating exceeds the maximum allowed percentage change
	oldRating := 100.0
	newRating := 140.0
	result := settings.ApplyMaxChangePerc(oldRating, newRating)
	expectedResult := 120.0 // Adjusted to stay within 20% of the old rating
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 2: New rating is within the allowed percentage change
	oldRating = 100.0
	newRating = 110.0
	result = settings.ApplyMaxChangePerc(oldRating, newRating)
	expectedResult = 110.0 // No adjustment needed
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 3: New rating exceeds the minimum allowed percentage change
	oldRating = 100.0
	newRating = 60.0
	result = settings.ApplyMaxChangePerc(oldRating, newRating)
	expectedResult = 80.0 // No adjustment needed
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}
}

func TestSettingsApplymaxChangeAbs(t *testing.T) {
	// Create a sample Settings configuration with maxChangeAbs
	settings := elo.New(elo.WithMaxChangeAbs(20))

	// Test case 1: New rating exceeds the maximum allowed absolute change
	oldRating := 100.0
	newRating := 130.0
	result := settings.ApplyMaxChangeAbs(oldRating, newRating)
	expectedResult := 120.0 // Adjusted to stay within 20 units of the old rating
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 2: New rating is within the allowed absolute change
	oldRating = 100.0
	newRating = 110.0
	result = settings.ApplyMaxChangeAbs(oldRating, newRating)
	expectedResult = 110.0 // No adjustment needed
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}

	// Test case 3: New rating exceeds the minimum allowed absolute change
	oldRating = 100.0
	newRating = 40.0
	result = settings.ApplyMaxChangeAbs(oldRating, newRating)
	expectedResult = 80.0 // Adjusted to stay within 20 units of the old rating
	if result != expectedResult {
		t.Errorf(ERROR_MESSAGE, expectedResult, result)
	}
}
