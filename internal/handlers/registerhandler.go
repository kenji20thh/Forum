package handlers

import (
	"fmt"
	"net/http"
)

func RegiserHandler(w http.ResponseWriter, r* http.Request) {
	fmt.Println(w, "register")
}
