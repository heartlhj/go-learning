package web

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	dbops "go-learning/workflow/db"
	"go-learning/workflow/errs"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t, e := template.ParseFiles("./templates/index.html")
	if e != nil {
		log.Printf("Parsing template index.htmlerror: %s", e)
		return
	}
	var bytearry = &dbops.Bytearry{Name: "nihao"}
	t.Execute(w, bytearry)
	return

}

func create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &dbops.Bytearry{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, errs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.CreateByteArry(ubody.Name, ubody.Bytes); err != nil {
		sendErrorResponse(w, errs.ErrorDBError)
		return
	}

	sendNormalResponse(w, string("ok"), 201)
}
