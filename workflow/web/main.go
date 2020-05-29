package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", index)
	router.POST("/import", create)
	return router
}

func main() {
	r := RegisterHandler()
	log.Println("StartIng Http.....")
	http.ListenAndServe(":8080", r)
}
