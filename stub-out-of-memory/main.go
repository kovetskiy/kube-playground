package main

import "log"

func main() {
	log.Println("consuming all ram")

	ints := []int{}
	for {
		ints = append(ints, 1)
	}
}
