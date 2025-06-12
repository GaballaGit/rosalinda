package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var BUFFER []byte
	for scanner.Scan() {
		b := scanner.Bytes()

		BUFFER = make([]byte, len(b))
		copy(BUFFER, b)

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := 0; i < len(BUFFER); i++ {
		if BUFFER[i] == 'T' {
			BUFFER[i] = 'U'
		}
	}
	fmt.Printf("%s\n", BUFFER)
}
