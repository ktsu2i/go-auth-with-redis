package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dbUser     = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbHost     = os.Getenv("MYSQL_HOST")
	dbPort     = os.Getenv("MYSQL_PORT")
	dbName     = os.Getenv("MYSQL_DATABASE")
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func Init() (*gorm.DB, error) {
	dsn := GetDSN(Config{
		Host:     dbHost,
		User:     dbUser,
		Password: dbPassword,
		Port:     dbPort,
		DBName:   dbName,
	})
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy:    schema.NamingStrategy{SingularTable: true},
		AllowGlobalUpdate: false,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDSN(c Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", c.User, c.Password, c.Host, c.Port, c.DBName)
}
