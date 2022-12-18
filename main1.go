package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func run(waitGroup *sync.WaitGroup, baton chan int, team string) {
	id := <-baton
	timeSleeping := time.Duration(50+rand.Int31n(100)) * time.Millisecond * 10
	fmt.Printf("[Team %s] Runner No.%d took the baton\n", team, id)
	time.Sleep(timeSleeping)
	if id < 4 {
		fmt.Printf("[Team %s] Runner No.%d passed the baton. TIME = %dms\n", team, id, timeSleeping.Milliseconds())
	} else {
		fmt.Printf("[Team %s] Runner No.%d finished the race. TIME = %dms\n", team, id, timeSleeping.Milliseconds())
	}

	baton <- id + 1
	waitGroup.Done()
}

func maketeam(teamWaitGroup *sync.WaitGroup, pole chan<- string, team string) {
	const runners int = 4

	baton := make(chan int, 1)
	baton <- 1
	var waitGroup sync.WaitGroup
	waitGroup.Add(runners)
	for i := 0; i < runners; i++ {
		go run(&waitGroup, baton, team)
	}
	waitGroup.Wait()
	pole <- team

	teamWaitGroup.Done()
}

func main() {
	teams := [...]string{
		"The Water Coolers",
		"The Scanners",
		"Boss Mode",
	}
	pole := make(chan string, (len(teams)))
	var teamsWaitGroup sync.WaitGroup
	teamsWaitGroup.Add((len(teams)))
	for i := 0; i < (len(teams)); i++ {
		go maketeam(&teamsWaitGroup, pole, teams[i])
	}
	teamsWaitGroup.Wait()
	fmt.Println("Race is over.")

	fmt.Println("Podium:")
	for i := 0; i < (len(teams)); i++ {
		fmt.Printf("%dº: %s\n", i+1, <-pole)
	}
}
