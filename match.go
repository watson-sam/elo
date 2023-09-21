package elo

const (
	DefaultInitRating    float64 = 2600
	DefaultC             float64 = 400
	DefaultHomeAdvantage float64 = 0
	DefaultKFactor       float64 = 32
)

// Settings represents the configuration for the rating system.
type Match struct {
	initRating     float64   // initRating is the initial rating value.
	c              float64   // c is a scaling factor affecting the steepness of the probability curve.
	homeAdvantage  float64   // homeAdvantage is the home advantage factor (if any).
	kFactor        float64   // kFactor is the update factor used in rating calculations.
	decayFactor    float64   // decayFactor is the factor used to decay rating.
	decayFactorOpp float64   // decayFactorOpp is the factor used to decay opposition rating.
	maxChangePerc  float64   // maxChangePerc defines the maximum percentage change allowed for a rating update.
	maxChangeAbs   float64   // maxChangeAbs defines the maximum absolute change allowed for a rating update.
	UpdateFunc     *Update   // UpdateFunc is a user-defined update function, if specified.
	ObservedFunc   *Observed // ObservedFunc is a user-defined observed function, if specified.
	ExpectedFunc   *Expected // ExpectedFunc is a user-defined expected function, if specified.
}

// Option is a function type that defines a configuration option for customizing the Settingm.
type Option func(m *Match)

func WithInitRating(initRating float64) Option {
	return func(m *Match) {
		m.initRating = initRating
	}
}

func WithC(c float64) Option {
	return func(m *Match) {
		m.c = c
	}
}

func WithHomeAdvantage(homeAdvantage float64) Option {
	return func(m *Match) {
		m.homeAdvantage = homeAdvantage
	}
}

func WithKFactor(kFactor float64) Option {
	return func(m *Match) {
		m.kFactor = kFactor
	}
}

func WithDecayFactor(decayFactor float64) Option {
	return func(m *Match) {
		m.decayFactor = decayFactor
	}
}

func WithDecayFactorOpp(decayFactorOpp float64) Option {
	return func(m *Match) {
		m.decayFactorOpp = decayFactorOpp
	}
}

func WithMaxChangePerc(maxChangePerc float64) Option {
	return func(m *Match) {
		m.maxChangePerc = maxChangePerc
	}
}

func WithMaxChangeAbs(maxChangeAbs float64) Option {
	return func(m *Match) {
		m.maxChangeAbs = maxChangeAbs
	}
}

func WithObservedFunc(observed Observed) Option {
	return func(m *Match) {
		m.ObservedFunc = &observed
	}
}

func WithExpectedFunc(expeced Expected) Option {
	return func(m *Match) {
		m.ExpectedFunc = &expeced
	}
}

func WithUpdateFunc(update Update) Option {
	return func(m *Match) {
		m.UpdateFunc = &update
	}
}

// New creates a new Settings configuration with optional customizations using functional optionm.
// It takes one or more Option functions to customize the Settingm.
func New(opts ...Option) Match {
	var obs Observed = ObsWinLooseDraw
	var exp Expected = ExpProbability
	var up Update = UpdateExpected
	m := Match{
		initRating:    DefaultInitRating,
		kFactor:       DefaultKFactor,
		c:             DefaultC,
		homeAdvantage: DefaultC,
		maxChangePerc: 0,
		maxChangeAbs:  0,
		ObservedFunc:  &obs,
		ExpectedFunc:  &exp,
		UpdateFunc:    &up,
	}
	for _, o := range opts {
		o(&m)
	}
	return m
}

func (m Match) NewRating() float64 {
	return m.initRating
}
