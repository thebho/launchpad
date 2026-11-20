package isro

import (
	"testing"
)

func TestLaunchPadCreation(t *testing.T) {
	name := "TestLaunchPad"
	launchPad := LaunchPad{Name: name}
	if launchPad.Name != name {
		t.Errorf("Expected LaunchPad name %s, got %s", name, launchPad.Name)
	}
	if len(launchPad.satellites) != 0 {
		t.Errorf("Expected 0 satellites on new LaunchPad, got %d", len(launchPad.satellites))
	}
}

func TestAddSingleSAT(t *testing.T) {
	launchPad := LaunchPad{Name: "Test"}
	sat := NewSatellite("SAT1")
	launchPad.AddSAT(sat)
	if len(launchPad.satellites) != 1 {
		t.Errorf("Expected 1 satellite, got %d", len(launchPad.satellites))
	}
	if launchPad.satellites[0].Name != "SAT1" {
		t.Errorf("Expected satellite name SAT1, got %s", launchPad.satellites[0].Name)
	}
}

func TestAddMultipleSATs(t *testing.T) {
	launchPad := LaunchPad{Name: "Test"}
	count := 10
	for i := 1; i <= count; i++ {
		launchPad.AddSAT(NewSatellite("SAT" + string(rune(i))))
	}
	if len(launchPad.satellites) != count {
		t.Errorf("Expected %d satellites, got %d", count, len(launchPad.satellites))
	}
}

func TestLaunchSATsEmptyPad(t *testing.T) {
	launchPad := LaunchPad{Name: "Test"}
	c := make(chan string, 1)
	defer close(c)
	launchPad.LaunchSATs(c)
	success := <-c
	if success != "Success" {
		t.Errorf("Expected Success, got %s", success)
	}
}

func TestLaunchSATsClearsAfterLaunch(t *testing.T) {
	launchPad := LaunchPad{Name: "Test"}
	launchPad.AddSAT(NewSatellite("SAT1"))
	launchPad.AddSAT(NewSatellite("SAT2"))
	launchPad.AddSAT(NewSatellite("SAT3"))
	c := make(chan string, 1)
	defer close(c)
	launchPad.LaunchSATs(c)
	<-c
	if len(launchPad.satellites) != 0 {
		t.Errorf("Expected 0 satellites after launch, got %d", len(launchPad.satellites))
	}
}

func TestMultipleLaunchSequences(t *testing.T) {
	launchPad := LaunchPad{Name: "Test"}
	c := make(chan string, 1)
	defer close(c)

	// First launch sequence
	launchPad.AddSAT(NewSatellite("SAT1"))
	launchPad.LaunchSATs(c)
	<-c
	if len(launchPad.satellites) != 0 {
		t.Errorf("Expected 0 satellites after first launch, got %d", len(launchPad.satellites))
	}

	// Second launch sequence
	launchPad.AddSAT(NewSatellite("SAT2"))
	launchPad.AddSAT(NewSatellite("SAT3"))
	launchPad.LaunchSATs(c)
	<-c
	if len(launchPad.satellites) != 0 {
		t.Errorf("Expected 0 satellites after second launch, got %d", len(launchPad.satellites))
	}
}

func TestLaunchPadWithManyLargeSatellites(t *testing.T) {
	launchPad := LaunchPad{Name: "Test"}
	count := 100
	for i := 1; i <= count; i++ {
		launchPad.AddSAT(NewSatellite("SAT" + string(rune(i))))
	}
	c := make(chan string, 1)
	defer close(c)
	launchPad.LaunchSATs(c)
	success := <-c
	if success != "Success" {
		t.Errorf("Expected Success, got %s", success)
	}
	if len(launchPad.satellites) != 0 {
		t.Errorf("Expected 0 satellites after launch, got %d", len(launchPad.satellites))
	}
}
