package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func checkArgs() int {
	size, err := strconv.Atoi(os.Args[2])
	if err == nil {
		printTree(size)
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
