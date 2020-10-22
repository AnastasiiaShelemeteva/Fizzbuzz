package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Please write, how many number you would like to have")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	amountOfNumbers, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= amountOfNumbers; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

	}
}
