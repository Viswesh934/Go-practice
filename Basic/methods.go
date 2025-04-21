package main

import "fmt"

//methods related to arrays,slices, maps

func main() {

	//arrays

	var arr []int = []int{1, 2, 3, 5}

	arr[0] = 34

	//slices

	slice := []int{1, 2, 3, 4}

	slice = append(slice, 2)

	slice = slice[:len(slice)-1]

	fmt.Println(slice)

	//maps

	ma := map[string]int{
		"a": 1,
		"b": 2,
	}
	delete(ma, "b")
	fmt.Println(ma)

	//set

	se := make(map[string]struct{})

	se["a"] = struct{}{}
	se["b"] = struct{}{}

	fmt.Printf("%v", se)

}
