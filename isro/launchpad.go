package isro

import (
	"fmt"
	"sync"
	"time"
)

type LaunchPad struct {
	Name string
}

var wg sync.WaitGroup

func (l LaunchPad) LaunchSATs(a, b, c, d Satellite) bool {
	fmt.Printf("Launching from LaunchPad %s\n", l.Name)
	wg.Add(4)
	go launchSAT(a)
	go launchSAT(b)
	go launchSAT(c)
	go launchSAT(d)
	wg.Wait()
	fmt.Println()
	return true
}

func launchSAT(a Satellite) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		fmt.Printf("%d...", i)
		time.Sleep(1 * time.Millisecond)
	}
	a.Launch()
}
