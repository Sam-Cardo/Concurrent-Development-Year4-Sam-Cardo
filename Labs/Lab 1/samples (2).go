//samples(2).go
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
)

/*
calculates the total factorial of the number
*/
func factorial(N int64) int64 {
	var ans int64
	ans = 1
	for i := int64(1); i < N; i++ {
		ans = ans * int64(i)
	}
	return ans * N
}

/*
this calculates the fib number
*/
func fib(N int) int {
	var ans int
	if N < 3 {
		ans = 1
	} else {
		ans = fib(N-1) + fib(N-2)
	}
	return ans
}

/*
this will apply a function every value in the array
*/
func applyMap(theFun func(int) int, theArray []int) {
	for i := range len(theArray) {
		theArray[i] = theFun(theArray[i])
	}
}

/*
this is the main function and it will get the number from the user and show factorial, fib
and then applying it to an array.
*/
func main() {
	var num int64
	num = 1
	for num != 0 {
		fmt.Print("Enter a number:")
		fmt.Scanln(&num)
		result := factorial(num)
		fmt.Println("the factorial is: ", result)
		fmt.Println("the fibonacci is:", fib(int(num)))
	}
	addOne := func(N int) int {
		return N + 1
	}
	X := addOne(9)
	fmt.Println(X)
	myArray := make([]int, 4)
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	fmt.Println(myArray)
	applyMap(addOne, myArray)
	fmt.Println(myArray)

}
