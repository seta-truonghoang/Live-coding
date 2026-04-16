package main

import "fmt"

func main() {
	var name string = "Gemini"
	var age = 25
	job := "Software Engineer"
	const Pi = 3.14

	fmt.Printf("Name: %s, Age: %d, Job: %s, Pi: %f\n", name, age, job, Pi)

	score := 90
	fmt.Println("Initial score:", score)

	updateValue(score)
	fmt.Println("After updateValue:", score)

	updateWithPointer(&score)
	fmt.Println("After updateWithPointer:", score)

	var p *int = &score
	fmt.Printf("Address: %v, Value at address: %d\n", p, *p)

	s1, s2 := swap("World", "Hello")
	fmt.Println(s1, s2)
}

func updateValue(s int) {
	s = s + 10
}

func updateWithPointer(s *int) {
	*s = *s + 10
}

func add(a int, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}
