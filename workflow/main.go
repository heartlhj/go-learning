package main

import (
	. "github.com/heartlhj/go-learning/workflow/web"
	"github.com/julienschmidt/httprouter"

	"log"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/import", Create)
	router.POST("/select", Query)
	return router
}

func main() {
	r := RegisterHandler()
	log.Println("StartIng Http.....")
	http.ListenAndServe(":8080", r)
}
