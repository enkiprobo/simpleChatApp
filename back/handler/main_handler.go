package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"topsusmoprog/tugassatu/back/model"
)

func InsertMessage(w http.ResponseWriter, r *http.Request) {
	unCookie, err := r.Cookie("username")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
		return
	}

	user1, err := database.GetUser(unCookie.Value)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	friend := r.FormValue("friend")
	user2, err := database.GetUser(friend)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chatroom := database.GetChatRoom(user1.ID, user2.ID)
	if chatroom == 0 {
		log.Println("chat room not exist")
		http.Error(w, "chat room not exist", http.StatusNotAcceptable)
		return
	}

	message := r.FormValue("message")
	chatDetailID, err := database.InsertChat(message, chatroom, user1.ID)

	mapResponse := map[string]int64{
		"chat_detail_id": chatDetailID,
	}

	response, err := json.Marshal(mapResponse)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	return
}

func GetChatDetail(w http.ResponseWriter, r *http.Request) {

	unCookie, err := r.Cookie("username")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
		return
	}

	userLogin, err := database.GetUser(unCookie.Value)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	friend := r.FormValue("friend")
	user2, err := database.GetUser(friend)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chatroom := database.GetChatRoom(userLogin.ID, user2.ID)
	if chatroom == 0 {
		log.Println("chat room not exist")
		http.Error(w, "chat room not exist", http.StatusNotAcceptable)
		return
	}

	chatDetail, err := database.GetChatDetail(chatroom)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(chatDetail)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func GetChatFriends(w http.ResponseWriter, r *http.Request) {
	unCookie, err := r.Cookie("username")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
		return
	}

	friends, err := database.GetUsers(unCookie.Value)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(friends)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	return
}

func LoginAuthUser(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")

	u, err := database.GetUser(username)
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

func LoginAuthPass(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := database.GetUser(username)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status := false
	if u.Password == string(password) {
		status = true
		cookie := http.Cookie{
			Name:  "username",
			Value: username,
		}
		http.SetCookie(w, &cookie)
	}

	mapresponse := map[string]bool{
		"status": status,
	}

	response, err := json.Marshal(mapresponse)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	return
}

func MainSimpleChatApp(w http.ResponseWriter, r *http.Request) {

	htmlfile := "main_chat.html"

	_, err := r.Cookie("username")
	if err != nil {
		htmlfile = "main_login.html"
	}

	main := path.Join("front", "templates", htmlfile)
	tmpl, err := template.ParseFiles(main)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
