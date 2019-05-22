package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("Press enter to start!")
	fmt.Scanln()

	time.Sleep(time.Duration(rand.Int63n(15)) * time.Second)

	start := time.Now()
	row := "<------------->"
	fmt.Printf("			%s\n			%s\n			%s\n			%s\n", row, row, row, row)

	fmt.Scanln()
	end := time.Now()
	elapsed := end.Sub(start) / 1000000 //Convert nanoseconds to milliseconds

	fmt.Printf("Your reaction time is %d milliseconds!\n", elapsed)
}
