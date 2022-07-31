package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	port := flag.Int("port", 9009, "port")
	dbport := flag.Int("db-port", 3306, "db port")
	dbpass := flag.String("db-pass", "123456", "db password")
	dbuser := flag.String("db-user", "root", "db user")
	dbhost := flag.String("db-host", "127.0.0.1", "db host")
	dbdb := flag.String("db-db", "agent", "db database")
	flag.Parse()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		*dbuser, *dbpass, *dbhost, *dbport, *dbdb)
	open := mysql.Open(dsn)
	db, err := gorm.Open(open)
	if err != nil {
		fmt.Println(err)
		return
	}
	service = NewService(db)
	runApiServer(*port)
	runGrpcServer(*port)
}
