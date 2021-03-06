package models

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gobuild/gobuild2/pkg/config"
)

var (
	tables []interface{}
	orm    *xorm.Engine
)

func getwith(orig, dft string) string {
	orig = strings.TrimSpace(orig)
	if orig == "" {
		return dft
	}
	return orig
}

func InitDB() (err error) {
	dbCfg := config.Config.Database
	switch dbCfg.DbType {
	case "mysql":
		orm, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			dbCfg.User, dbCfg.Password, dbCfg.Host, getwith(dbCfg.Port, "3306"), dbCfg.Name))
	case "postgres":
		cnnstr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
			dbCfg.User, dbCfg.Password,
			getwith(dbCfg.Host, "127.0.0.1"), getwith(dbCfg.Port, "5432"), dbCfg.Name, dbCfg.SslMode)
		orm, err = xorm.NewEngine("postgres", cnnstr)
	default:
		return fmt.Errorf("Unknown database type: %s\n", dbCfg.DbType)
	}
	if err != nil {
		return fmt.Errorf("models.init(fail to conntect database): %v\n", err)
	}
	return orm.Sync(tables...)
}
