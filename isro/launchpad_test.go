package isro

import "testing"

func TestAddSAT(t *testing.T) {
	name1 := "TestSAT1"
	name2 := "TestSAT2"
	launchPad := LaunchPad{Name: "Test"}
	launchPad.AddSAT(NewSatellite(name1))
	launchPad.AddSAT(NewSatellite(name2))
	if name1 != launchPad.satellites[0].Name {
		t.Errorf("Name didn't match for first SAT: %s", launchPad.satellites[0].Name)
	}
	if name2 != launchPad.satellites[1].Name {
		t.Errorf("Name didn't match for first SAT: %s", launchPad.satellites[1].Name)
	}
}

func TestLaunchSATS(t *testing.T) {
	name1 := "TestSAT1"
	name2 := "TestSAT2"
	launchPad := LaunchPad{Name: "Test"}
	launchPad.AddSAT(NewSatellite(name1))
	launchPad.AddSAT(NewSatellite(name2))
	var c = make(chan string, 1)
	defer close(c)
	launchPad.LaunchSATs(c)
	success := <-c
	if "Success" != success {
		t.Errorf("Expected Success got: %s", success)
	}
	if 0 != len(launchPad.satellites) {
		t.Errorf("Expected 0 satellites and got: %d", len(launchPad.satellites))
	}
}
