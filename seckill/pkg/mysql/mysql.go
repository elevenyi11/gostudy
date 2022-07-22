package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
)

var engin *gorose.Engin

var err error

func InitMysql(hostMysql, portMysql, userMysql, pwdMysql, dbMysql string) {
	fmt.Printf(userMysql)

	DbConfig := gorose.Config{
		Driver:          "mysql",
		Dsn:             userMysql + ":" + pwdMysql + "@tcp(" + hostMysql + ":" + portMysql + ")/" + dbMysql + "?charset=utf8&parseTime=true",
		Prefix:          "",
		SetMaxOpenConns: 300,
		SetMaxIdleConns: 10,
	}

	engin, err = gorose.Open(&DbConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DB() gorose.IOrm {
	return engin.NewOrm()
}