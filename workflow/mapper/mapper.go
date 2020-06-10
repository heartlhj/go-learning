package mapper

import (
	"errors"
	"fmt"
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/heartlhj/go-learning/workflow/model"
	"github.com/prometheus/common/log"
)

func CreateByteArry(name string, bytes string) error {
	byteArry := &model.Bytearry{Name: name, Bytes: bytes}
	insert, err := db.MasterDB.Insert(byteArry)
	if err != nil {
		log.Infoln("新增数据异常", err)
		return err
	}
	fmt.Println(insert)
	return nil
}

func Select(name string, bytes string) ([]*model.Bytearry, error) {
	tmpArticle := make([]*model.Bytearry, 0)
	err := db.MasterDB.Where("name=?", name).Find(&tmpArticle)
	if err != nil {
		log.Infoln("查询异常", err)
		return nil, errors.New("查询异常!")
	}
	return tmpArticle, nil
}
