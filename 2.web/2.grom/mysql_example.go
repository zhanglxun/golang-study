package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDb() error {
	var err error
	dsn := "pxndev:pxndev@2018@tcp(rm-wz97vpexguxsff7en5o.mysql.rds.aliyuncs.com:3306)/db_apiserver"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

type User struct {
	Id   int64          `db:"id"`
	Name sql.NullString `db:"string"`
	Age  int            `db:"age"`
}

func testQueryData() {
	for i := 0; i < 101; i++ {
		fmt.Printf("query %d times\n", i)
		sqlstr := "select id, name, age from user where id=?"
		row := DB.QueryRow(sqlstr, 1)
		/*if row != nil {
			continue
		}*/
		var user User
		err := row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}

		fmt.Printf("id:%d name:%v age:%d\n", user.Id, user.Name, user.Age)
	}

}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	testQueryData()
}
