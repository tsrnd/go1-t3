package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"	
)
type PostgresInfo struct {
	Username string
	Password  string
	Name string
	Hostname string
	Port int
}
func (p PostgresInfo) connect() (*sql.DB, error) {
	psqlInfo := dnsInfo(p)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
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