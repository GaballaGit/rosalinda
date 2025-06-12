package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	line := ""

	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println("Read line:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading stdin:", err)
	}

	var dna map[byte]int
	dna = make(map[byte]int)

	dna['A'] = 0
	dna['C'] = 0
	dna['G'] = 0
	dna['T'] = 0

	fmt.Println()
	for i := 0; i < len(line); i++ {
		dna[line[i]] += 1
	}

	fmt.Println(dna['A'], " ", dna['C'], " ", dna['G'], " ", dna['T'])
}
