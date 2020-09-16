package DBHelper

import (
	"Go-Common/Logger"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	err error
)

//初始化数据信息
func NewDBHelper(server string) {
	db, err = sqlx.Connect("mssql", server)
	if err != nil {
		Logger.Logger.Errorln(err.Error())
	}
}

//执行 sql 语句返回 rows
func Query(sql string, params ...interface{}) *sqlx.Rows {
	rows, err := db.Queryx(sql, params...)
	if err != nil {
		Logger.Logger.Errorln(err.Error())
	}
	return rows
}

//执行 sql 语句返回受影响的行数
func Execute(sql string, params ...interface{}) int64 {
	res, err := db.Exec(sql, params...)
	if err != nil {
		Logger.Logger.Errorln(err.Error())
	}
	ra, err := res.RowsAffected()
	if err != nil {
		Logger.Logger.Errorln(err.Error())
	}

	return ra
}
