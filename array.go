package main

import "fmt"

func main() {
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7}

	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)

	s1 := []string{"Go", "Python", "TypeScript"}
	s2 := arr2[1:3]
	s3 := make([]int, 5, 10)
	var s4 []int

	s1 = append(s1, "Rust")
	s1 = append(s1, "Java", "C++")

	fmt.Println("Slice 1:", s1)
	fmt.Println("Slice 2:", s2)
	fmt.Println("Slice 3:", s3, "Len:", len(s3), "Cap:", cap(s3))
	fmt.Println("Slice 4 is nil:", s4 == nil)
}
