package behavior

import (
	"github.com/heartlhj/go-learning/workflow/db"
	"github.com/jinzhu/gorm"
)

type TransactionContextInterceptor struct {
	Next CommandInterceptor
}

func (transactionContextInterceptor TransactionContextInterceptor) Execute(command Command) (value interface{}, err error) {
	defer db.ClearTXDB()
	db.GORM_DB.Transaction(func(tx *gorm.DB) error {
		db.InitTXDB(tx)
		value, err = transactionContextInterceptor.Next.Execute(command)
		return err
	})
	return value, err
}

func (transactionContextInterceptor *TransactionContextInterceptor) SetNext(next CommandInterceptor) {
	transactionContextInterceptor.Next = next
}
