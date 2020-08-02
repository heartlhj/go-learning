package persistence

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/errs"
	. "github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

type DefineManager struct {
}

func (define DefineManager) FindDeployedProcessDefinitionByKey(key string) ([]*Bytearry, error) {
	bytearries := make([]*Bytearry, 0)
	err := db.TXDB.Where("`key`=?", key).Find(&bytearries).Error
	return bytearries, err
}

func (define DefineManager) GetBytearry(processDefineId int64) (Bytearry, error) {
	bytearries := Bytearry{}
	err := db.TXDB.Where("id=?", processDefineId).First(&bytearries).Error
	if err != nil {
		log.Infoln("create processInstance err", err)
		return bytearries, err
	}
	return bytearries, nil
}

func (define DefineManager) CreateByteArry(name string, key string, bytes string) error {
	bytearries, err := define.FindDeployedProcessDefinitionByKey(key)
	if err != nil {
		return err
	}
	var verion = 0
	if bytearries != nil && len(bytearries) > 0 {
		verion = bytearries[0].Version
		verion++
	}
	byteArry := Bytearry{Name: name, Bytes: bytes, Key: key, Version: verion}
	err = db.TXDB.Create(&byteArry).Error
	if err != nil {
		log.Infoln("新增数据异常", err)
		return err
	}
	return nil
}

func (define DefineManager) FindProcessByTask(processInstanceId int64) (Bytearry, error) {
	bytearries := make([]Bytearry, 0)
	var sql = "SELECT b.* FROM bytearry b " +
		"LEFT JOIN process_instance p on b.key = p.key " +
		"WHERE p.id = ? "
	err := db.TXDB.Raw(sql, processInstanceId).Find(&bytearries).Error
	if err != nil {
		return Bytearry{}, err
	}
	if bytearries != nil && len(bytearries) > 0 {
		return bytearries[0], nil
	}
	return Bytearry{}, errs.ProcessError{Code: "1001", Msg: "Not Find"}
}
