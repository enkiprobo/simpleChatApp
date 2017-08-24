package main

import (
	"log"
	"net/http"
	"topsusmoprog/tugassatu/back/handler"
	"topsusmoprog/tugassatu/back/model"
)

func main() {
	err := database.InitDB()

	if err != nil {
		log.Println(err.Error())
		return
	}

	// frontend resource
	http.Handle("/front/", http.StripPrefix("/front/", http.FileServer(http.Dir("front"))))

	// ajax handler
	http.HandleFunc("/login", handler.LoginAuth)

	// html render handler
	http.HandleFunc("/", handler.MainSimpleChatApp)

	http.ListenAndServe(":8080", nil)
}
