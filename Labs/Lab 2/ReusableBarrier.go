//Reusable Barrier.go Code
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
// A simple reusable barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope

//--------------------------------------------

package main

// imports
import (
	"fmt"
	"sync"
	"time"
)

/*
---Func doStuff---
each go routines will go through this function
and will wait until it reaches the end of each barrier.
*/

func doStuff(goNum int, inside *int, outside *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, insideChan chan bool,
	outsideChan chan bool) bool {

	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	// lock
	sharedLock.Lock()
	//increase counter
	*inside++

	if *inside == max {

		// unlock
		sharedLock.Unlock()
		// release
		insideChan <- true
		//aquire
		<-insideChan

	} else {

		//unlock
		sharedLock.Unlock()
		//aquire
		<-insideChan
		//release
		insideChan <- true

	} //end of if-else

	// End of first barrier

	// Everyone has passed the first barrier then moves to second barrier
	sharedLock.Lock()
	*outside++

	// start of second barrier
	if *outside == max {

		// unlock
		sharedLock.Unlock()
		// release
		outsideChan <- true
		//aquire
		<-outsideChan

	} else {

		// unlock
		sharedLock.Unlock()
		//aquire
		<-outsideChan
		//release
		outsideChan <- true

	}

	// end of second barrier

	// reset counters after both barriers
	sharedLock.Lock()   // lock
	*inside--           // decrease counter
	*outside--          // decrease counter
	sharedLock.Unlock() //unlock

	fmt.Println("PartB", goNum)
	wg.Done()
	return true

} //end-doStuff

func main() {
	totalRoutines := 10 // 10 goroutines
	inside := 0         //counter for first barrier
	outside := 0        // counter for second barrier

	var wg sync.WaitGroup  // creates the wait group for go routines
	wg.Add(totalRoutines)  // adds tasks to the total amount of goroutines
	var theLock sync.Mutex // creates a mutex variable

	insideChan := make(chan bool)  // channel for first barrier
	outsideChan := make(chan bool) // channel for second barrier

	for i := range totalRoutines { //starts the go Routines here
		go doStuff(i, &inside, &outside, totalRoutines, &wg, &theLock, insideChan, outsideChan)
	} //starts the go Routines here

	wg.Wait() // waits for all go routines

} //end-main
