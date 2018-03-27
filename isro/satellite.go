package isro

import "fmt"

// Satellite is a SAT class
type Satellite struct {
	Name string
}

// NewSatellite creates a new Satellite object
func NewSatellite(name string) Satellite {
	return Satellite{Name: name}
}

// Launch records the launching of the Satellite
func (s Satellite) Launch() {
	fmt.Printf("Launching %s", s.Name)
}
