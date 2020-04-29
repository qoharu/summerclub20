package main

import "fmt"

var haveCar bool = false
var distanceToOffice int = 600

func main() {
	fmt.Println("Hi Gopher!")
}

func isWalkable(distance int) bool {
	if distance > 500 {
		return false
	}
	
	return true
	
}

func walkToWork() {
	fmt.Println("Walk to Work")
}

func driveToWork() {
	fmt.Println("Drive to Work")
}

func usePublicTransport() {
	fmt.Println("Use Public Transportation")
}

func mul(a float32, b float32) float32 {
	return a * b
}
