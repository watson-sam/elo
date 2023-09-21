package elo

// Update is a function type that defines the signature of an update function for the rating system.
type Update func(observed float64, expected float64, kFactor float64) float64

// UpdateExpected is an update function that calculates the change in rating based on the expected and observed values.
// It takes the following parameters:
// - observed (float64): The actual observed value.
// - expected (float64): The expected value.
// - kFactor (float64): The K-factor used in rating calculations.
// It returns the calculated rating change as a float64 value.
func UpdateExpected(observed float64, expected float64, kFactor float64) float64 {
	return (kFactor * (observed - expected))
}

// UpdatePoints is an update function that calculates the change in rating based on the difference between observed and expected values, divided by a factor.
// It takes the following parameters:
// - observed (float64): The actual observed value.
// - expected (float64): The expected value.
// - kFactor (float64): The K-factor used in rating calculations.
// It returns the calculated rating change as a float64 value.
func UpdatePoints(observed float64, expected float64, kFactor float64) float64 {
	difference := observed - expected
	return ((difference / kFactor) / 2)
}

// applyMaxChange applies a maximum change limit to a new rating value, ensuring it stays within a specified range.
// It takes the following parameters:
// - minRating (float64): The minimum allowed rating value.
// - maxRating (float64): The maximum allowed rating value.
// - newRating (float64): The new rating value to be checked and possibly adjusted.
// It returns the adjusted rating as a float64 value.
func ApplyMaxChange(minRating, maxRating, newRating float64) float64 {
	if newRating < minRating {
		return minRating
	}
	if newRating > maxRating {
		return maxRating
	}
	return newRating
}

// applymaxChangePerc applies a maximum percentage change to a new rating value.
// It takes the following parameters:
// - oldRating (float64): The previous rating value.
// - newRating (float64): The new rating value to be checked and possibly adjusted.
// It returns the adjusted rating as a float64 value.
func (m *Match) ApplyMaxChangePerc(oldRating float64, newRating float64) float64 {
	minRating := oldRating * (1 - m.maxChangePerc)
	maxRating := oldRating * (1 + m.maxChangePerc)
	return ApplyMaxChange(minRating, maxRating, newRating)
}

// applymaxChangeAbs applies a maximum absolute change to a new rating value.
// It takes the following parameters:
// - oldRating (float64): The previous rating value.
// - newRating (float64): The new rating value to be checked and possibly adjusted.
// It returns the adjusted rating as a float64 value.
func (m *Match) ApplyMaxChangeAbs(oldRating float64, newRating float64) float64 {
	minRating := oldRating - m.maxChangeAbs
	maxRating := oldRating + m.maxChangeAbs
	return ApplyMaxChange(minRating, maxRating, newRating)
}

// update calculates a new rating based on the observed and expected values using the specified update function and applies maximum changes if configured.
// It takes the following parameters:
// - rating (float64): The current rating value.
// - observed (float64): The actual observed value.
// - expected (float64): The expected value.
// It returns the adjusted new rating as a float64 value.
func (m *Match) update(rating float64, observed float64, expected float64) float64 {
	var updateFunc Update
	if m.UpdateFunc != nil {
		updateFunc = *m.UpdateFunc
	} else {
		updateFunc = UpdateExpected
	}
	change := updateFunc(observed, expected, m.kFactor)
	newRating := rating + change
	if m.maxChangePerc != 0 {
		newRating = m.ApplyMaxChangePerc(rating, newRating)
	} else if m.maxChangeAbs != 0 {
		newRating = m.ApplyMaxChangeAbs(rating, newRating)
	}
	return newRating
}

// UpdateRating calculates a new rating based on the provided ratings and scores using the configured functions and settings.
// It takes the following parameters:
// - rating (float64): The current rating value.
// - ratingOpp (float64): The rating of the opposing team or player.
// - score (float64): The score of the subject team or player.
// - scoreOpp (float64): The score of the opposing team or player.
// It returns the updated rating as a float64.
func (m *Match) UpdateRating(rating float64, ratingOpp float64, score float64, scoreOpp float64) float64 {
	expected := m.Expected(rating, ratingOpp)
	observed := m.observed(score, scoreOpp)
	return m.update(rating, observed, expected)
}
