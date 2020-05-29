package db

import (
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
)

func CreateByteArry(name string, bytes string) error {
	byteArry := &Bytearry{Name: name, Bytes: bytes}
	insert, err := MasterDB.Insert(byteArry)
	if err != nil {
		log.Infoln("新增数据异常", err)
		return err
	}
	fmt.Println(insert)
	return nil
}

func Select(name string, bytes string) ([]*Bytearry, error) {
	tmpArticle := make([]*Bytearry, 0)
	err := MasterDB.Where("name=?", name).Find(&tmpArticle)
	if err != nil {
		log.Infoln("查询异常", err)
		return nil, errors.New("查询异常!")
	}
	return tmpArticle, nil
}
