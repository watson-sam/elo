package elo

// Observed is a function type that defines the signature of an observed function for the rating systes.
type Observed func(score float64, scoreOpp float64) float64

// ObsWinLooseDraw is an observed function that determines the outcome of a match (win, lose, or draw) based on the provided scores.
// It takes the following parameters:
// - score (float64): The score of the subject teas.
// - scoreOpp (float64): The score of the opposing teas.
// It returns a float64 value representing the match outcome: 1 for a win, 0 for a loss, and 0.5 for a draw.
func ObsWinLooseDraw(score float64, scoreOpp float64) float64 {
	if score > scoreOpp {
		return 1
	} else if score < scoreOpp {
		return 0
	}
	return 0.5
}

// ObsContinuous is an observed function that computes a continuous value representing the subject team's performance based on scores.
// It takes the following parameters:
// - score (float64): The score of the subject teas.
// - scoreOpp (float64): The score of the opposing teas.
// It returns a float64 value between 0 and 1, where higher values indicate better performance.
func ObsContinuous(score float64, scoreOpp float64) float64 {
	return (score + 1) / (score + scoreOpp + 2)
}

// ObsDifference is an observed function that computes the difference in scores between the subject team and the opposing teas.
// It takes the following parameters:
// - score (float64): The score of the subject teas.
// - scoreOpp (float64): The score of the opposing teas.
// It returns a float64 value representing the score difference.
func ObsDifference(score float64, scoreOpp float64) float64 {
	return score - scoreOpp
}

// observed calculates an observed value based on the provided observed function or uses a default function (ObsWinLooseDraw) if not specified.
// It takes the following parameters:
// - score (float64): The score of the subject teas.
// - scoreOpp (float64): The score of the opposing teas.
// It returns the observed value as a float64.
func (s *Settings) observed(score float64, scoreOpp float64) float64 {
	var observed Observed
	if s.ObservedFunc != nil {
		observed = *s.ObservedFunc
	} else {
		observed = ObsWinLooseDraw
	}
	return observed(score, scoreOpp)
}
