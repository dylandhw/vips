package main

import (
	"fmt"
	"time"
)

const green = "\033[0;32m"
const yellow = "\033[0;33m"

func main() {
	fmt.Println(yellow + "    Video Integrity And Provenanve System    " + green + "\n")
	start := time.Now()

	PartitionImage()
	ProcessImageFrames()

	elapsed := time.Since(start)
	fmt.Println("execution time: \n", elapsed)
}
