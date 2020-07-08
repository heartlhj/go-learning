package web

import (
	"encoding/xml"
	"github.com/heartlhj/go-learning/workflow/engine"
	"github.com/heartlhj/go-learning/workflow/engine/behavior"
	"github.com/heartlhj/go-learning/workflow/model"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	MAX_UPLOAD_SIZE = 50 * 1024 * 1024 // 文件大小 50MB
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t, e := template.ParseFiles("workflow\\templates\\index.html")
	if e != nil {
		log.Printf("Parsing template index.htmlerror: %s", e)
		return
	}
	var bytearry = &model.Bytearry{Name: "你好"}
	t.Execute(w, bytearry)
	return

}

func Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("file")
	name := r.Form.Get("name")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	body, err := ioutil.ReadAll(file)
	//解析xml数据
	data := new(engine.Definitions)
	err = xml.Unmarshal(body, &data)
	dataStr, err := xml.MarshalIndent(data, "", " ")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "XML转换异常")
		return
	}
	behavior.Converter(data)
	behavior.ConvertXMLToElement(data)
	//导出xml文件
	headerBytes := []byte(xml.Header)                //加入XML头
	xmlOutPutData := append(headerBytes, dataStr...) //拼接XML头和实际XML内容

	//设置Content-Type
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+name+"\".bpmn20.xml")

	sendNormalResponse(w, string(xmlOutPutData), 201)
}
