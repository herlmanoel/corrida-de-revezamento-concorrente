package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func receiveBaton(waitGroup *sync.WaitGroup, baton chan string, id int) {
	timeSleeping := time.Duration(rand.Int31n(1000)) * time.Millisecond
	id = id + 1
	isLast := (id+1) > 4
	person := "Runner " + strconv.Itoa(id)
	baton <- person
	time.Sleep(timeSleeping)
	fmt.Println("The " + person + " with the baton. TIME = ", timeSleeping)
	if isLast {
		fmt.Println("The race is over.")
		waitGroup.Done()
		return
	}
	waitGroup.Done()
	receiveBaton(waitGroup, baton, id)
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)
	baton := make(chan string, 4)
	go receiveBaton(&waitGroup, baton, 0)
	waitGroup.Wait()
}
