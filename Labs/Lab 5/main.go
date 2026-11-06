//Producer / consumer code
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
//
// Issues:
// None I hope

//--------------------------------------------

package main

import (
	"fmt"
	"sync"
)

// Event is the data that will get passed between producer and consumer
type Event struct {
	Text string
}

// producer creates events and then puts them into a buffer
func producer(total int, buffer chan Event, spaces chan Event, items chan Event, lock *sync.Mutex, wg *sync.WaitGroup) {

	// repeats for the total number of events
	for range total {

		// creates a new event
		event := Event{"Hello from Producer"}

		// waits there until there is space in the buffer to continue
		<-spaces

		// locks the buffer so only one go routine is able to go through at a time
		lock.Lock()

		// add the event to the buffer
		buffer <- event

		// unlocks the buffer
		lock.Unlock()

		// tells the consumer that an event is available to be consumed
		items <- Event{}

		// prints the produced event
		fmt.Println("Produced:", event.Text)
	}

	// says the producer that it's done
	wg.Done()
}

// consumer takes the event out of the buffer
func consumer(total int, buffer chan Event, spaces chan Event, items chan Event, lock *sync.Mutex, wg *sync.WaitGroup) {

	//loops until the total number of events have gone through
	for range total {

		// waits until an event has become available
		<-items

		// locks the buffer
		lock.Lock()

		// takes a event out of the buffer
		event := <-buffer

		// unlocks the buffer
		lock.Unlock()

		// tells the producer that there's spaces available
		spaces <- Event{}

		// prints the consumed text
		fmt.Println("Consumed:", event.Text)
	}

	// says the consumer that it's done
	wg.Done()
}

func main() {

	// Variables
	var wg sync.WaitGroup
	var lock sync.Mutex

	totalGoroutines := 10
	total := 100
	size := 20

	buffer := make(chan Event, size)
	spaces := make(chan Event, size)
	items := make(chan Event, size)

	// adds an empty event for each available space in the buffer
	// makes 20 spaces so that producer and consumer can only go 20 times
	for range size {
		spaces <- Event{}
	}

	// tells the wait group how man goRoutines to wait for
	wg.Add(totalGoroutines)

	// starts the producer function
	go producer(total, buffer, spaces, items, &lock, &wg)

	// starts the consumer function
	go consumer(total, buffer, spaces, items, &lock, &wg)

	// waits until all the go routines are done
	wg.Wait()
}
