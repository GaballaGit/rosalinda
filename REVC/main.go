package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var ReveresedDNA []byte

	var b []byte
	for scanner.Scan() {
		b = scanner.Bytes()

		ReveresedDNA = make([]byte, len(b))
		copy(ReveresedDNA, b)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	compliment := make(map[byte]byte)

	compliment['A'] = 'T'
	compliment['T'] = 'A'
	compliment['C'] = 'G'
	compliment['G'] = 'C'

	for i := 0; i < len(ReveresedDNA); i++ {
		ReveresedDNA[i] = compliment[ReveresedDNA[i]]
	}

	slices.Reverse(ReveresedDNA)
	fmt.Printf("%s\n-----------------------------------------\n%s\n", b, ReveresedDNA)
}
