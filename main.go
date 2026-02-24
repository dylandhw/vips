package main

import (
	"fmt"
	"time"
)

const (
	green  = "\033[0;32m"
	yellow = "\033[0;33m"
)

func main() {
	fmt.Println(green + "    Video Integrity And Provenanve System    " + yellow + "\n")
	start := time.Now()

	PartitionImage()

	elapsed := time.Since(start)
	fmt.Println("execution time: \n", elapsed)
}
