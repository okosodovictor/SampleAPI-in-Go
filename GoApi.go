package main

import (
	"fmt"
	"net/http"

	"./handlers"
)

func main() {
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:4601", nil)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
}
