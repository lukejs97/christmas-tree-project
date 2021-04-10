package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func checkArgs() int {
	size := 3
	if len(os.Args) < 3 {
		println("Size not specified. Defaulting to size 3.")
	} else if len(os.Args) > 3 {
		println("Error: Too many arguments")
        println("Usage: go run tree.go -height [TREE_SIZE]")
        println("Exiting...")
        os.Exit(1)
	} else {
		size, err := strconv.Atoi(os.Args[2])
		if err == nil {
			return size
		}
		println("Error: TREE_SIZE must be a integer")
		os.Exit(1)
	}
	return size
}

func printTree(size int) {
	fmt.Printf("%*s\n", size + 1, "*")
	midSpace := 2
	for i := size; i > 0; i-- {
		fmt.Printf("%*s ", i, "/")
		fmt.Printf("%*s", midSpace, "\\\n")
		midSpace += 2
	}
	println(strings.Repeat("-", ((size * 2) + 1)))
	fmt.Printf("%*s\n", size + 1, "#")
}

func main() {
	size := checkArgs()
	printTree(size)
}
