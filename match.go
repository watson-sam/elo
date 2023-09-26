package elo

type PlayerTeam struct {
	RatingRaw float64
	Rating    float64
}

// decay adjusts the current rating of a team towards the init rating of the system according to a given decayFactor, it is directional and
// and therefore ratings will only ever be smaller or the same size in magnitude, the movement is controlled by the decay factor whereby
// it behaves like a weighting between the current rating and inital rating
// It takes the following parameters:
// - decayFactor (float64): The deacy factor used to weight current rating vs initial rating.
// - initRating (float64): The default rating of the subject teams
func (pt *PlayerTeam) decay(decayFactor float64, initRating float64) {
	pt.Rating = pt.RatingRaw
	if pt.RatingRaw > initRating {
		pt.Rating = (pt.RatingRaw * decayFactor) + (initRating * (1 - decayFactor))
	}
}

type Match struct {
	Pt       PlayerTeam
	PtOpp    PlayerTeam
	score    float64
	scoreOpp float64
	settings Settings
	Expected float64
}

// UpdateRating calculates a new rating based on the provided ratings and scores using the configured functions and settings.
// It takes the following parameters:
// - rating (float64): The current rating value.
// - ratingOpp (float64): The rating of the opposing team or player.
// - score (float64): The score of the subject team or player.
// - scoreOpp (float64): The score of the opposing team or player.
// It returns the updated rating as a float64.
func (m *Match) UpdateRating() float64 {
	m.Pt.decay(m.settings.DecayFactor, m.settings.InitRating)
	m.PtOpp.decay(m.settings.DecayFactor, m.settings.InitRating)

	m.Expected = m.settings.Expected(m.Pt.Rating, m.PtOpp.Rating)
	observed := m.settings.observed(m.score, m.scoreOpp)
	return m.settings.update(m.Pt.Rating, observed, m.Expected)
}
