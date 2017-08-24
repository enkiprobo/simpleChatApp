package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
	"topsusmoprog/tugassatu/back/model"
)

func LoginAuth(w http.ResponseWriter, r *http.Request) {

	u, err := database.GetUser("enkiprobo")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(u)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	return
}

func MainSimpleChatApp(w http.ResponseWriter, r *http.Request) {

	mainLogin := path.Join("front", "templates", "main_login.html")
	tmpl, err := template.ParseFiles(mainLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
