package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/brucebales/discounts-api/src/internal/config"
)

//NewMysql returns a MySQL database connection using credentials from the Config struct
func NewMysql() (*sql.DB, error) {
	conf := config.GetConfig()

	db, err := sql.Open(conf.MysqlDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/webstore", conf.MysqlUser, conf.MysqlPass, conf.MysqlHost, conf.MysqlPort))
	if err != nil {
		return nil, err
	}
	return db, nil
}
