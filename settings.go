package elo

// Settings represents the configuration for the rating system.
type Settings struct {
	InitRating    float64   // InitRating is the initial rating value.
	C             float64   // C is a scaling factor affecting the steepness of the probability curve.
	HomeAdvantage float64   // HomeAdvantage is the home advantage factor (if any).
	KFactor       float64   // KFactor is the factor used in rating calculations.
	MaxChangePerc float64   // MaxChangePerc defines the maximum percentage change allowed for a rating update.
	MaxChangeAbs  float64   // MaxChangeAbs defines the maximum absolute change allowed for a rating update.
	UpdateFunc    *Update   // UpdateFunc is a user-defined update function, if specified.
	ObservedFunc  *Observed // ObservedFunc is a user-defined observed function, if specified.
	ExpectedFunc  *Expected // ExpectedFunc is a user-defined expected function, if specified.
}

// UpdateRating calculates a new rating based on the provided ratings and scores using the configured functions and settings.
// It takes the following parameters:
// - rating (float64): The current rating value.
// - ratingOpp (float64): The rating of the opposing team or player.
// - score (float64): The score of the subject team or player.
// - scoreOpp (float64): The score of the opposing team or player.
// It returns the updated rating as a float64.
func (s *Settings) UpdateRating(rating float64, ratingOpp float64, score float64, scoreOpp float64) float64 {
	expected := s.expected(rating, ratingOpp)
	observed := s.observed(score, scoreOpp)
	return s.update(rating, observed, expected)
}

// Setup returns a default Settings configuration for the rating system.
func Setup() Settings {
	return Settings{
		InitRating:    2000,
		KFactor:       10,
		HomeAdvantage: 0,
	}
}

// Option is a function type that defines a configuration option for customizing the Settings.
type Option func(c *Settings)

// New creates a new Settings configuration with optional customizations using functional options.
// It takes one or more Option functions to customize the Settings.
func New(opts ...Option) Settings {
	var obs Observed = ObsWinLooseDraw
	var exp Expected = ExpProbability
	var up Update = UpdateExpected
	c := Settings{
		InitRating:    2600,
		KFactor:       10,
		C:             400,
		HomeAdvantage: 0,
		ObservedFunc:  &obs,
		ExpectedFunc:  &exp,
		UpdateFunc:    &up,
	}
	for _, o := range opts {
		o(&c)
	}
	return c
}
