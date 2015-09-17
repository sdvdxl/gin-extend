package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sdvdxl/gin-extend/sources/util"
	"github.com/sdvdxl/gin-extend/sources/util/log"
)

var Engine *xorm.Engine

func init() {
	log.Logger.Info("init db engine ...")

	var err error
	{ //创建数据库engine
		Engine, err = xorm.NewEngine("", ":@tcp()/?charset=utf8&parseTime=true")
		util.PanicError(err)

		Engine.ShowSQL = true
		Engine.ShowDebug = true
		Engine.ShowErr = true
		Engine.ShowInfo = true
		Engine.ShowWarn = true

		Engine.SetMaxIdleConns(20)
		Engine.SetMaxOpenConns(30)

		err = Engine.Ping()
		util.PanicError(err)
	}

	log.Logger.Info("db engine inited")
}

func Close() {
	log.Logger.Info("db engine is cloing")
	Engine.Close()
	log.Logger.Info("db engine has been closed")
}
