package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello worlld!"))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(err)
}
