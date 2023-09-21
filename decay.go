package elo

// decay adjusts the current rating of a team towards the init rating of the system according to a give decayFactor, it is directional and
// and therefore ratings will only ever be smaller or the same size in magnitude, the movement is controlled by the decay factor whereby
// it behaves like a weighting between the current rating and inital rating
// It takes the following parameters:
// - rating (float64): The current rating of the subject team.
// - decayFactor (float64): The deacy factor used to weight current rating vs initial rating.
// It returns a float64 value of the decayed rating.
func (m *Match) decay(rating float64, decayFactor float64) float64 {
	if (rating > m.initRating) && (decayFactor > 0) {
		return (rating * decayFactor) + (m.initRating * (1 - decayFactor))
	}
	return rating
}
