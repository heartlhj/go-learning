package persistence

import (
	"encoding/xml"
	. "github.com/heartlhj/go-learning/workflow/converter"
	"github.com/heartlhj/go-learning/workflow/mapper"
	. "github.com/heartlhj/go-learning/workflow/model"
)

func FindDeployedProcessDefinitionByKey(key string) Process {
	bytearries, e := mapper.SelectByteByKey(key, "nil")
	if e != nil {

	}
	//解析xml数据
	data := new(Definitions)
	xml.Unmarshal([]byte(bytearries[0].Bytes), &data)
	Converter(data)
	ConvertXMLToElement(data)
	return data.Process[0]

}
