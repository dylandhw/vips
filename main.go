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
	fmt.Println(green + "---------------------------------------------" + yellow + "\n")
	fmt.Println(green + "    Video Integrity And Provenance System    " + yellow + "\n")
	fmt.Println(green + "---------------------------------------------" + yellow + "\n")

	start := time.Now()

	fmt.Print(PartitionImage())

	elapsed := time.Since(start)
	fmt.Println("\n\nTime to perform image partitioning and pixel analysis per frame: ", elapsed)
}
