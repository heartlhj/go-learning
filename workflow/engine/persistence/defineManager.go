package persistence

import (
	"fmt"
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

func FindDeployedProcessDefinitionByKey(key string) []*model.Bytearry {
	bytearries := make([]*model.Bytearry, 0)
	err := db.MasterDB.Where("`key`=?", key).Find(&bytearries)
	if err != nil {

	}
	return bytearries
}

func CreateByteArry(name string, key string, bytes string) error {
	bytearries := FindDeployedProcessDefinitionByKey(key)
	var verion = 0
	if bytearries != nil && len(bytearries) > 0 {
		verion = bytearries[0].Version
		verion++
	}
	byteArry := &model.Bytearry{Name: name, Bytes: bytes, Key: key, Version: verion}
	insert, err := db.MasterDB.Insert(byteArry)
	if err != nil {
		log.Infoln("新增数据异常", err)
		return err
	}
	fmt.Println(insert)
	return nil
}
