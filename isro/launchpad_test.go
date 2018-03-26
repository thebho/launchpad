package isro

import (
	"testing"
)

func TestLaunchSATS(t *testing.T) {
	testSATS := []Satellite{
		NewSatellite("Test1"),
		NewSatellite("Test2"),
		NewSatellite("Test3"),
		NewSatellite("Test4"),
	}
	testLaunchPad := LaunchPad{}
	testLaunchPad.LaunchSATs(testSATS[0], testSATS[1], testSATS[2], testSATS[3])
}
