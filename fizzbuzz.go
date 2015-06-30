package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// FizzBuzz logic
func FizzBuzz(n int) string {
	if (n%3 == 0) && (n%5 == 0) {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", n)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err == nil {
			fmt.Println(FizzBuzz(x))
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
