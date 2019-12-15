package controller

import (
	"net/http"
)

func SetupControllers() {
	http.Handle("/", http.FileServer(http.Dir("./client")))
}
