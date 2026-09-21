//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

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
// Created on 21/09/26
//--------------------------------------------

package main

// Imports
import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

/*
*
  - This is a function that uses a mutex to be able to track the amount of go
    routines that reach the barrier, and then it uses the channel to be able to
    make each of those go routines wait until they are able to reach Part A
    before going to Part B.
*/
func doStuff(goNum int, arrived *int, m int, wg *sync.WaitGroup, sem *semaphore.Weighted, theLock *sync.Mutex) bool {

	// this helps .Acquire to block until permission is available.
	ctx := context.Background()
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	//-- Start of barrier --

	// locks down before going to the arrived counter
	theLock.Lock()
	*arrived++

	if *arrived == m {
		theLock.Unlock()
		sem.Release(int64(m - 1))

	} else {
		theLock.Unlock()

		// makes the goroutine wait for permission to continue.
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Println("Error: ", err)
		}
	}

	// -- Past the Barrier --

	theLock.Lock()
	*arrived--
	theLock.Unlock()

	// Part B Go routines
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

/*
	This is the main function in this file and this is where most
	of the variables are declared and used and also sets up the go
	routines that will be used in this file.
*/

func main() {

	// Variables
	var wg sync.WaitGroup
	var arrived int
	var theLock sync.Mutex

	// Declarations
	totalRoutines := 10 // total go routines in this file
	m := totalRoutines
	ctx := context.TODO()

	wg.Add(totalRoutines) //tracks each go routine that completed

	// Semaphores
	sem := semaphore.NewWeighted(int64(totalRoutines)) // A semaphore with a limit which is the total routines
	sem.Acquire(ctx, int64(totalRoutines))             // this will wait until theres slot in the totalRoutines are available

	for i := 0; i < totalRoutines; i++ {
		go doStuff(i, &arrived, m, &wg, sem, &theLock) //the starts of the goRoutines and giving one a random and different number
	}

	wg.Wait() // waits until all goRoutines are done
}

//Classmate that helped Seamus

// Sources used
// Little book of semaphores
// Joseph Kehoe's reusable barrier code
// https://pkg.go.dev/golang.org/x/sync/semaphore
// Goland's code finisher
