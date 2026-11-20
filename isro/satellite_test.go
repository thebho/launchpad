package isro

import (
	"testing"
)

func TestNewSatellite(t *testing.T) {
	name := "TestSAT"
	sat := NewSatellite(name)
	if sat.Name != name {
		t.Errorf("Expected satellite name %s, got %s", name, sat.Name)
	}
}

func TestNewSatelliteWithEmptyName(t *testing.T) {
	sat := NewSatellite("")
	if sat.Name != "" {
		t.Errorf("Expected empty satellite name, got %s", sat.Name)
	}
}

func TestNewSatelliteWithSpecialCharacters(t *testing.T) {
	name := "SAT-123_XYZ"
	sat := NewSatellite(name)
	if sat.Name != name {
		t.Errorf("Expected satellite name %s, got %s", name, sat.Name)
	}
}
