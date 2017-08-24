package main

import (
	"net/http"
	"topsusmoprog/tugassatu/back/handler"
)

func main() {

	http.HandleFunc("/", handler.MainSimpleChatApp)

	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
