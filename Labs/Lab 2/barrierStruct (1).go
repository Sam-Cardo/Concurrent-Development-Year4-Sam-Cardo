//Barrier.go Template Code
//Copyright (C) 2024 Samuel Cardo

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Samuel Cardo (C00328718@setu.ie)
// Created on 28/09/26
// Modified by:
// Description:
// A simple barrier implemented using struct
// Issues:
// None I hope
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/*
Creates a barrier data type
two channels
one mutex
a total for number of go routines
counters for each barrier stage
*/
type barrier struct {
	theChan     chan bool
	secondChan  chan bool
	theLock     sync.Mutex
	total       int
	count       int
	secondCount int
}

/*
---func createBarrier---
// creates and returns a properly initialized barrier
// N== number of threads (go Routines)
*/
func createBarrier(N int) barrier {
	theBarrier := barrier{
		theChan:     make(chan bool),
		secondChan:  make(chan bool),
		total:       N,
		count:       0,
		secondCount: 0,
	}
	return theBarrier
}

/*
---func (b *barrier) wait()---
// Method belonging to barrier data type
// Blocks until everyone reaches this point then lets everyone continue
*/
func (b *barrier) wait() {

	// first barrier
	b.theLock.Lock() // locks the mutex
	b.count++        // increases the counter

	if b.count == b.total { // checks for the last goroutine to arrive
		b.theLock.Unlock()  // unlocks the mutex
		fmt.Println("here") // prints here
		for i := 0; i < b.total-1; i++ {
			<-b.theChan // acquires
		}
	} else {
		fmt.Println(b.count) // prints the count
		b.theLock.Unlock()   // unlocks the mutex
		b.theChan <- true    // releases
	} // end of first barrier

	b.theLock.Lock() // locks the mutex
	b.secondCount++  // increases second counter

	if b.secondCount == b.total { // checks for the last goroutine to arrive
		b.count = 0       // resets counter
		b.secondCount = 0 // resets the second counter

		b.theLock.Unlock()               // unlocks
		fmt.Println("here")              // prints here
		for i := 0; i < b.total-1; i++ { //
			<-b.secondChan // acquires
		}
	} else {
		fmt.Println(b.secondCount)
		b.theLock.Unlock()   // unlocks the mutex
		b.secondChan <- true // releases
	}
} //wait

/*
---func workwithredezvous---
each will wait a random amount of time, hits the first barrier, then print B,
hits the second barrier.
*/

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, theBarrier *barrier) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here
	theBarrier.wait()
	fmt.Println("PartB", Num)
	theBarrier.wait()
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrier := createBarrier(5)
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, &barrier)
	}
	wg.Wait() //wait here until everyone (5 go routines) is done

}
