//mutex(1).go
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
)

/*
	this function increases the total n times
	and uses a mutex to protect the total
*/

func adds(n int, theLock *sync.Mutex, total int64, wg *sync.WaitGroup) bool {
	for i := 0; i < n; i++ {
		theLock.Lock()
		total++
		theLock.Unlock()
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

/*
	this is the main function and it starts 10 go routines then
	adds the total up and then it waits to be done.
*/

func main() {

	//theLock will be passed by reference between go routines
	var wg sync.WaitGroup
	var total int64
	var theLock sync.Mutex

	total = 0
	//the waitgroup is used as a barrier
	// init it to number of go routines
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		//starting
		fmt.Println(i)
		go adds(1000, &theLock, total, &wg)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total)
}
