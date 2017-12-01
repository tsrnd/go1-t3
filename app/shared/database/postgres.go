package database

import (
	_ "github.com/lib/pq"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
type PostgresInfo struct {
	Username string
	Password  string
	Name string
	Hostname string
	Port int
}
func (p PostgresInfo) connect() (*gorm.DB, error) {
	psqlInfo := dnsInfo(p)
	db, err := gorm.Open("postgres", psqlInfo)
	
	if err != nil {
		fmt.Println("err connect database", err)
		panic(err)
	}
	return db, err
}
func dnsInfo(p PostgresInfo) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.Hostname, p.Port, p.Username, p.Password, p.Name)
	return psqlInfo
}