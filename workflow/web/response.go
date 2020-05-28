package web

import (
	"encoding/json"
	"go-learning/workflow/errs"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp errs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
