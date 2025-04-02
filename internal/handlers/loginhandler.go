package handlers

import (
	"fmt"
	"net/http"
)

func LoginHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "login")
}