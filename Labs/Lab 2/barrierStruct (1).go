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
// A simple barrier implemented using sturct
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

// Create a barrier data type
type barrier struct {
	theChan     chan bool
	secondChan  chan bool
	theLock     sync.Mutex
	total       int
	count       int
	secondCount int
}

// creates a properly initialized barrier
// N== number of threads (go Routines)
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

// Method belonging to barrier data type
// Blocks until everyone reaches this point then lets everyone continue
func (b *barrier) wait() {
	b.theLock.Lock()
	b.count++

	if b.count == b.total {
		b.theLock.Unlock()
		fmt.Println("here")
		for _ = range b.total - 1 {
			<-b.theChan
		}
	} else {
		fmt.Println(b.count)
		b.theLock.Unlock()
		b.theChan <- true
	}

	b.theLock.Lock()
	b.secondCount++

	if b.secondCount == b.total {
		b.count = 0
		b.secondCount = 0

		b.theLock.Unlock()
		fmt.Println("here")
		for _ = range b.total - 1 {
			<-b.secondChan
		}
	} else {
		fmt.Println(b.secondCount)
		b.theLock.Unlock()
		b.secondChan <- true
	}
} //wait

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
