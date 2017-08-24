package handler

import (
	"fmt"
	"net/http"
)

func MainSimpleChatApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "halo")
	return
}
