package main

// imports
import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/*
	this function goes through a go routine where it makes each go
	routine wait until all of Part to be done before it goes to
	Part B.
*/

func WorkWithRendezvous(wg *sync.WaitGroup, barrier *sync.WaitGroup, Num int) bool {

	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)

	//Rendezvous point here
	barrier.Done() // decreases count each go routine that arrived
	barrier.Wait() // waits until the counter goes to zero before continuing

	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {

	// declaration of variables
	var wg sync.WaitGroup
	var barrier sync.WaitGroup

	threadCount := 5

	wg.Add(threadCount)      // tracks each go routine completion
	barrier.Add(threadCount) // tracks the arrivals to the rendezvous point in WorkWithRendezvous
	for N := range threadCount {
		go WorkWithRendezvous(&wg, &barrier, N)
	}
	wg.Wait() //waits here until every go routine is done

}

//Classmate that helped Seamus
