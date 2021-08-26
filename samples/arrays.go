package main

import "fmt"

func main() {
	println("Arrays and Slices in go")

	planets := [8]string{"Mercury", "Venus", "Mars", "Earth", "Jupiter", "Saturn", "Uranus", "Neptune"}
	//fmt.Println(planets)

	for index, value := range planets {
		fmt.Println(index)
		fmt.Println(value)
	}

	planetSlice := []string{"Mercury", "Venus", "Mars", "Earth", "Jupiter", "Saturn", "Uranus", "Neptune"}

	for index, value := range planetSlice {
		fmt.Println(index)
		fmt.Println(value)
	}

	var planetSliceVerbose []string
	planetSliceVerbose = append(planetSliceVerbose, "Mercury")

	fmt.Println(planetSliceVerbose)

}
