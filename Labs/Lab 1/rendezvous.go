//Rendezvous.go 
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

//--------------------------------------------
// Author: Samuel Cardo (C00328718@setu.ie)
// Created on 21/09/26
//--------------------------------------------

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
