package db

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/heartlhj/go-learning/workflow/runtime"
	"github.com/jinzhu/gorm"
	. "github.com/polaris1119/config"
	"sync"
)

//var MasterDB *xorm.Engine

var GORM_DB *gorm.DB

var TXDB *sync.Map

var dns string

func init() {
	TXDB = new(sync.Map)
	initConf()
	mysqlConfig, err := ConfigFile.GetSection("mysql")
	if err != nil {
		fmt.Println("get mysql config error:", err)
		return
	}

	fillDns(mysqlConfig)
	// 启动时就打开数据库连接
	if err = initEngine(); err != nil {
		panic(err)
	}
	// 测试数据库连接是否 OK
	if err = GORM_DB.DB().Ping(); err != nil {
		panic(err)
	}
}

var (
	ConnectDBErr = errors.New("connect db error")
	UseDBErr     = errors.New("use db error")
)

func initConf() {
	dbname := "go-workflow"
	uname := "root"
	pwd := "1234"
	dbhost := "localhost"
	dbport := "3306"

	ConfigFile.SetSectionComments("mysql", "")
	ConfigFile.SetValue("mysql", "host", dbhost)
	ConfigFile.SetValue("mysql", "port", dbport)
	ConfigFile.SetValue("mysql", "user", uname)
	ConfigFile.SetValue("mysql", "password", pwd)
	ConfigFile.SetValue("mysql", "dbname", dbname)
	ConfigFile.SetValue("mysql", "charset", "utf8")
}

// TestDB 测试数据库
func TestDB() error {
	initConf()
	mysqlConfig, err := ConfigFile.GetSection("mysql")
	if err != nil {
		fmt.Println("get mysql config error:", err)
		return err
	}

	tmpDns := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s&parseTime=True&loc=Local",
		mysqlConfig["user"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["charset"])
	_, err = gorm.Open("mysql", tmpDns)
	if err != nil {
		fmt.Println("new engine error:", err)
		return err
	}
	defer GORM_DB.Close()

	// 测试数据库连接是否 OK
	if err = GORM_DB.DB().Ping(); err != nil {
		fmt.Println("ping db error:", err)
		return ConnectDBErr
	}

	err = GORM_DB.Exec("use " + mysqlConfig["dbname"]).Error
	if err != nil {
		fmt.Println("use db error:", err)
		err = GORM_DB.Exec("CREATE DATABASE " + mysqlConfig["dbname"] + " DEFAULT CHARACTER SET " + mysqlConfig["charset"]).Error
		if err != nil {
			fmt.Println("create database error:", err)

			return UseDBErr
		}

		fmt.Println("create database successfully!")
	}
	// 初始化 MasterDB
	return Init()
}

func Init() error {
	initConf()
	mysqlConfig, err := ConfigFile.GetSection("mysql")
	if err != nil {
		fmt.Println("get mysql config error:", err)
		return err
	}

	fillDns(mysqlConfig)

	// 启动时就打开数据库连接
	if err = initEngine(); err != nil {
		fmt.Println("mysql is not open:", err)
		return err
	}

	return nil
}

func fillDns(mysqlConfig map[string]string) {
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConfig["user"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["dbname"],
		mysqlConfig["charset"])
}

func initEngine() error {
	var err error
	GORM_DB, err = gorm.Open("mysql", dns)
	if err != nil {
		return err
	}
	GORM_DB.LogMode(true)
	return nil
}

func InitTXDB(db *gorm.DB) {
	TXDB.Store(runtime.GoroutineId(), db)
}

func ClearTXDB() {
	TXDB.Delete(runtime.GoroutineId())
}

func DB() *gorm.DB {
	db, ok := TXDB.Load(runtime.GoroutineId())
	if !ok {
		panic("db not init")
	}
	return db.(*gorm.DB)
}
