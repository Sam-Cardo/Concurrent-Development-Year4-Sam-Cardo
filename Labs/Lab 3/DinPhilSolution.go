// Dining Philosophers Code
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
---Func Think---
This simulates a philosopher thinking for a random amount of time
and will sleep for set amount of time that ranges from zero to
four seconds. Also prints that the philosopher is thinking.
*/

func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

/*
---Func Eat---
This simulates a philosopher eating for a random amount of time
and will sleep for set amount of time that ranges from zero to
four seconds. Also prints that the philosopher is eating.
*/

func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

/*
---Func getForks---
This will attempt to acquire the two forks that philosopher will need
and also will send true saying that philosopher is holding a fork.
*/

func getForks(index int, forks map[int]chan bool) {
	if index == 4 {
		forks[4] <- true // takes fourth fork
		forks[0] <- true // takes the first fork
	} else {
		forks[index] <- true       // takes the left fork
		forks[(index+1)%5] <- true // takes the right fork
	}
}

/*
---Func putForks---
This will attempt to release the two forks that philosopher that was held by
the philosopher.
*/

func putForks(index int, forks map[int]chan bool) {
	if index == 4 {
		<-forks[index] // this will release the fourth fork
		<-forks[0]     // this will release the first fork
	} else {
		<-forks[index]       // this will release the left fork
		<-forks[(index+1)%5] // this will release the right fork
	}

}

/*
---Func doPhilStuff---
This function will be in an infinite loop of the philosophers going around the
table thinking, then acquires the forks, then eats and then sets down the forks.
*/
func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	for {
		think(index)           // starts the thinking
		getForks(index, forks) // acquires the forks
		eat(index)             // starts the eating
		putForks(index, forks) // releases the forks
	}
	wg.Wait()
}

/*
---Func main---
This is the main function, and it creates a wait group for the five philosophers
and also creates the five fork channels, This function also makes five go routines
one per philosophers, then calls wg.wait until all the go routines have been
completed.
*/
func main() {

	var wg sync.WaitGroup // creates the wait group for the philosophers
	philCount := 5        // number of philosophers
	wg.Add(philCount)     // adds the five tasks to philcount

	forks := make(map[int]chan bool) // creates the map for the channels

	for k := range philCount { //set up forks
		forks[k] = make(chan bool, 1) // creates fork channel
	} //set up forks

	for N := range philCount { //start philosophers
		go doPhilStuff(N, &wg, forks) // starts the go routines
	} //start philosophers

	wg.Wait() //wait here until everyone (5 go routines) is done

} // End of file
