package main

import (
	"fmt"
	"math"
)

// No classes in go functions,structs and interfaces only
func sum(a int, b int) int {
	return a + b
}

func variables() {
	// variables
	var a int = 10 //a:=10 short hand
	var b float64 = 20.5
	var c string = "Hello"
	var d bool = true

	c += "Hey"
	fmt.Println(c)

	//Operations

	fmt.Println(a + int(b))
	//pow in go only accepts flaot64
	fmt.Println(math.Pow(b, 2))

	fmt.Println(a, b, c, d)

	// array
	var arr [5]int

	// slices

	var slice []int = []int{1, 2, 3, 4, 5}
	var strslice []string = []string{"A", "B", "c"}

	strslice = append(strslice, "D")

	fmt.Println(strslice)

	slice = append(slice, 6, 7, 8, 9, 10)
	fmt.Println(slice)

	//map

	var m map[string]int = make(map[string]int)
	m["one"] = 1
	m["two"] = 2

	//map with shorthand

	ma := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println("short hand map: ", ma)

	fmt.Println(m)

	// No while loop in Go, use for loop instead
	// Normal for loop
	for i := 0; i < 5; i++ {
		r := sum(a, i)
		fmt.Println(r)
		arr[i] = i + 1
	}
	//Range modernized for loop

	for i := range 5 {
		fmt.Println(i)
	}
	fmt.Println(arr)

	// structs in go

	type basic struct {
		Name  string
		Age   int
		Place string
	}

	basc := basic{Name: "github", Age: 20, Place: "go practice"}

	fmt.Println(basc.Name, basc.Place, basc.Age)

}
