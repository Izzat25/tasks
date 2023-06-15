package main

import (
	"fmt"
	"task/tasks"
	"time"
)

func main() {
	start := time.Now()
	res := tasks.Surface(5)

	fmt.Println(res)

	timeElapsed := time.Since(start)
	fmt.Printf("it took %s\n", timeElapsed)
}
