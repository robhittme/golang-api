package main

import (
	"fmt"
	"net/http"
	api "./webserver/api"
)

func main() {
	router := ApiRouter()

	fmt.Println("Server starting on port 8080")
	err := http.ListenAndServer("8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
