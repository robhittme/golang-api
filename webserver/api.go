package webserver

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func ApiRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Prepare the response data
		response := Response{Message: "Hello, World!"}

		// Encode the response data to JSON and write it to the response writer
		json.NewEncoder(w).Encode(response)
	})

	return router
}
