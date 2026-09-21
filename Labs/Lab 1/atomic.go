//atomic.go
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

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
adds one to the total n times and then signals that it is done
*/
func addsAtomic(n int, total *atomic.Int64, wg *sync.WaitGroup) bool {
	for i := 0; i < n; i++ {
		total.Add(1)
	}
	wg.Done() //let wait group know we have finished
	return true
}

/*
this is the main function that starts 10 go routines
and adds 1000 to the total and then waits to for it to be done
*/
func main() {

	var total atomic.Int64
	var wg sync.WaitGroup

	//for loop using range option
	for i := range 10 {
		//the wait group is used as a barrier
		// init it to number of go routines
		wg.Add(1)
		fmt.Println("go Routine ", i)
		go addsAtomic(1000, &total, &wg)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total.Load())

}
