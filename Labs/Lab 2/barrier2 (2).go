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
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
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

	sharedLock.Lock()   // lock
	*inside--           // decrease counter
	*outside--          // decrease counter
	sharedLock.Unlock() //unlock

	fmt.Println("PartB", goNum)
	wg.Done()
	return true

} //end-doStuff

func main() {
	totalRoutines := 10
	inside := 0
	outside := 0

	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	var theLock sync.Mutex

	insideChan := make(chan bool)
	outsideChan := make(chan bool)

	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &inside, &outside, totalRoutines, &wg, &theLock, insideChan, outsideChan)
	}
	wg.Wait()
} //end-main
