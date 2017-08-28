package main

import (
	"log"
	"net/http"
	"os"
	"topsusmoprog/tugassatu/back/handler"
	"topsusmoprog/tugassatu/back/model"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("no port error")
	}

	err := database.InitDB()
	defer database.ChatDB.Close()

	if err != nil {
		log.Println(err.Error())
		return
	}

	handler.Hubnya = handler.NewHub()
	go handler.Hubnya.Run()

	// frontend resource
	http.Handle("/front/", http.StripPrefix("/front/", http.FileServer(http.Dir("front"))))

	// ajax handler
	http.HandleFunc("/loginusername", handler.LoginAuthUser)
	http.HandleFunc("/loginpassword", handler.LoginAuthPass)
	http.HandleFunc("/getchatfriends", handler.GetChatFriends)
	http.HandleFunc("/getchatdetail", handler.GetChatDetail)
	http.HandleFunc("/insertmessage", handler.InsertMessage)

	// websocket
	http.HandleFunc("/livechat", handler.LiveChatHandler)

	// html render handler
	http.HandleFunc("/", handler.MainSimpleChatApp)

	http.ListenAndServe(":"+port, nil)
}
