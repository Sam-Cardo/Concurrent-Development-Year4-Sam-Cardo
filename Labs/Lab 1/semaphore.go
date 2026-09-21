//semaphore.go
//Copyright (C) 2026 Samuel Cardo

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

// --------------------------------------------
// Author: Samuel Cardo (C00328718@setu.ie)
// Created on 21/09/26
// --------------------------------------------
package main

import (
	"fmt"
	"sync"
	"time"
)

// make struct containing channel
// add init, acquire and release

/*
a semaphore type that uses a channel to control the access of things
*/
type semaphore struct {
	theCounter chan struct{}
}

/*
gets a spot in the semaphore
*/

func Acquire(sem *semaphore) {
	sem.theCounter <- struct{}{}
}

/*
releases the spot from the semaphore
*/

func Release(sem *semaphore) {
	<-sem.theCounter
}

/*
this is the main function that starts 20  go routines while only allows 5 of them
actually run at one time
*/
func main() {

	maxGoroutines := 5
	semaphore := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Simulate a task
			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
