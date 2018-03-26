package isro

import "fmt"

type Satellite struct {
	Name string
}

func NewSatellite(name string) Satellite {
	return Satellite{Name: name}
}

func (s Satellite) Launch() {
	fmt.Printf("Launching %s", s.Name)
}
