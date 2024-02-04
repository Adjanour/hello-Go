package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	a := 5
	b := 5
	argWithProg := os.Args
	argWithoutProg := os.Args[1:]

	numA, err := strconv.Atoi(argWithProg[1])
	if err != nil {
		numA = rand.Intn(10)
	}
	numB, err := strconv.Atoi(argWithProg[2])
	if err != nil {
		numB = rand.Intn(10)
	}
	result := numA + numB
	fmt.Printf("The result of %d + %d is %d\n", numA, numB, result)
	fmt.Println(argWithProg)
	fmt.Println(argWithoutProg)
	fmt.Println("Hello World")
	fmt.Println(a)
	fmt.Print(b)

	fmt.Println("Testing math functions")
	var angle float64 = 6.285
	fmt.Println(math.Sin(angle))
	fmt.Println(math.Cos(angle))
	fmt.Println(math.Tan(angle))
}
