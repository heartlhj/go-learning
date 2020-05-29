package db

import (
	"database/sql"
	"errors"
	"fmt"

	. "github.com/polaris1119/config"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"
	"xorm.io/xorm"
)

var MasterDB *xorm.Engine

var dns string

func init() {
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
	if err = MasterDB.Ping(); err != nil {
		panic(err)
	}
}

var (
	ConnectDBErr = errors.New("connect db error")
	UseDBErr     = errors.New("use db error")
)

func initConf() {
	var (
		// domain       = global.App.Host + ":" + global.App.Port
		xormLogLevel = "0"
		xormShowSql  = "true"
	)

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
	ConfigFile.SetKeyComments("mysql", "max_idle", "最大空闲连接数")
	ConfigFile.SetValue("mysql", "max_idle", "2")
	ConfigFile.SetKeyComments("mysql", "max_conn", "最大打开连接数")
	ConfigFile.SetValue("mysql", "max_conn", "10")

	ConfigFile.SetSectionComments("xorm", "")
	ConfigFile.SetValue("xorm", "show_sql", xormShowSql)
	ConfigFile.SetKeyComments("xorm", "log_level", "0-debug, 1-info, 2-warning, 3-error, 4-off, 5-unknow")
	ConfigFile.SetValue("xorm", "log_level", xormLogLevel)
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
	egnine, err := xorm.NewEngine("mysql", tmpDns)
	if err != nil {
		fmt.Println("new engine error:", err)
		return err
	}
	defer egnine.Close()

	// 测试数据库连接是否 OK
	if err = egnine.Ping(); err != nil {
		fmt.Println("ping db error:", err)
		return ConnectDBErr
	}

	_, err = egnine.Exec("use " + mysqlConfig["dbname"])
	if err != nil {
		fmt.Println("use db error:", err)
		_, err = egnine.Exec("CREATE DATABASE " + mysqlConfig["dbname"] + " DEFAULT CHARACTER SET " + mysqlConfig["charset"])
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

	MasterDB, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		return err
	}

	maxIdle := ConfigFile.MustInt("mysql", "max_idle", 2)
	maxConn := ConfigFile.MustInt("mysql", "max_conn", 10)

	MasterDB.SetMaxIdleConns(maxIdle)
	MasterDB.SetMaxOpenConns(maxConn)

	showSQL := ConfigFile.MustBool("xorm", "show_sql", false)
	logLevel := ConfigFile.MustInt("xorm", "log_level", 1)

	MasterDB.ShowSQL(showSQL)
	MasterDB.Logger().SetLevel(core.LogLevel(logLevel))

	// 启用缓存
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// MasterDB.SetDefaultCacher(cacher)

	return nil
}

func StdMasterDB() *sql.DB {
	return MasterDB.DB().DB
}
