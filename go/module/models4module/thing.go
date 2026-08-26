package models4module

// Thing is a minimal example record used by facade4module to demonstrate
// this module's data-access pattern end to end. Copy the shape - not this
// exact type - when modeling your own module's persisted data.
type Thing struct {
	// Greeting is a required field, kept deliberately simple so the example
	// stays concrete.
	Greeting string `json:"greeting"`
}
