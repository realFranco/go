/*
A tiny set of functions to use as lambda functions

> go run annon-funct.go

Ref: https://yourbasic.org/golang/anonymous-function-literal-lambda-closure
*/

package main

import (
	"fmt"
	"sort"
)

func normalSlice() {
	people := []string{"Alice", "Bob", "Dave"}
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})

	fmt.Println(people)
}

// Using an intermediate variable.
// This is a closre, it references a "people" var.
// which is declared outside the function
func intermediateSlice() {
	people := []string{"Alice", "Bob", "Dave"}
	less := func(i, j int) bool {
		return len(people[i]) < len(people[j])
	}

	sort.Slice(people, less)
	fmt.Println(people)
}

func main() {
	fmt.Println("Sorting using normalSlice")
	normalSlice()

	fmt.Println("\nSorting using intermeidateSlice")
	intermediateSlice()
}
