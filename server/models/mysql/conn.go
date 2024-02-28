package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"week1/server/config"
	"xorm.io/xorm"
)

var XDB *xorm.Engine

// 初始化mysql连接
func XormConn(m *config.Mysql) (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8mb4", m.Username, m.Password, m.Port, m.Databases))
}

// 数据表迁移
func Migrator() error {
	return XDB.Sync2(&User{})
}
