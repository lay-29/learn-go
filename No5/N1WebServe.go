package No5

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func WebServeTest() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Louis!"))
	})
	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("<h1>Hello Louis!</h1>"))
	})
	http.HandleFunc("/api/hello", func(writer http.ResponseWriter, request *http.Request) {
		resp := Response{Message: "Hello Hello!"}
		writer.Header().Set("Content-Type", "application/json")
		da := json.NewEncoder(writer)
		da.Encode(resp)
	})

	http.ListenAndServe(":8080", nil)
}
