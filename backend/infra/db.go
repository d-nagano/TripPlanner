package infra

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func initDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(db)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected!")
	return db, nil
}
