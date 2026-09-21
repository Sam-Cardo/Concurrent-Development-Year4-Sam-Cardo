//main.go 
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
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func fib(N int) int {
	if N < 2 {
		return 1
	} else {
		return fib(N-1) + fib(N-2)
	}
}

func parFib(N int) int {
	var wg sync.WaitGroup
	var A, B int
	wg.Add(2)
	if N < 2 {
		return 1
	} else {
		go func(N int, Ans *int) {
			defer wg.Done()
			*Ans = parFib(N - 1)
		}(N, &A)
		go func(N int, Ans *int) {
			defer wg.Done()
			*Ans = parFib(N - 2)
		}(N, &B)
		wg.Wait()
		return A + B
	}
}

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	for i := 0; i < 10; i++ {
		Seq := fib(i * 5)
		par := parFib(i * 5)
		fmt.Println(Seq, "---", par)
	}

}

//Classmate that helped Seamus
