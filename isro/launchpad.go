package isro

import (
	"fmt"
	"sync"
	"time"
)

type LaunchPad struct {
	Name       string
	satellites []Satellite
}

var wg sync.WaitGroup

func (l *LaunchPad) AddSAT(sat Satellite) {
	l.satellites = append(l.satellites, sat)
}

func (l *LaunchPad) LaunchSATs(c chan string) {
	fmt.Printf("Launching from LaunchPad %s\n", l.Name)
	wg.Add(len(l.satellites))
	for _, sat := range l.satellites {
		go launchSAT(sat)
	}
	wg.Wait()
	fmt.Println()
	l.satellites = nil
	c <- "Success"
}

func launchSAT(a Satellite) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		fmt.Printf("%d...", i)
		time.Sleep(1 * time.Millisecond)
	}
	a.Launch()
}
