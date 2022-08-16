package main

import (
	"fmt"
	"time"
)

func timer() {
	counter := 0

	for {
		if true {
			time.Sleep(1 * time.Second)
			counter++
			if counter%5 == 0 {
				fmt.Printf("Hello World! I was born %d seconds ago. \n", counter)
			}
		}
	}
}

func main() {
	timer()
}
