package main

import (
	"fmt"
	// "strings"
	// "log"
	// "example.com/greetings"
	"time"
)

func main() {
	fmt.Println("gaming")
	start := time.Now()
	// day1()
	// day2()
	// day3()
	// day4()
	// day5() // This question is so dogwater wtf 4.3335592s
	// day6() // pen n paper but did make a python sol for it
	day7()
	// day8()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
