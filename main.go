package main

import (
	"io"
	"log"
	"net/http"
)

// func timer(counter int) {
// 	for {
// 		if true {
// 			time.Sleep(1 * time.Second)
// 			counter++
// 			if counter%5 == 0 {
// 				s := fmt.Sprintf("Hello World! I was born %d seconds ago. \n", counter)
// 				return s
// 			}
// 		}
// 	}
// }

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
