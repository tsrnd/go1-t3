package database

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)
var (
	// SQL wrapper
	SQL *gorm.DB
)
type congfigure interface {
	connect() (*gorm.DB, error)
}
const (
	// TypeBolt is BoltDB
	TypeBolt string = "Bolt"
	// TypeMongoDB is MongoDB
	TypeMongoDB string = "MongoDB"
	// TypeMySQL is MySQL
	TypeMySQL string = "MySQL"
	TypePosgres string = "Postgres"
)
type Type string
type Connections struct {
	Postgres PostgresInfo
}
type Info struct {
	Driver string
	Connections Connections
}
func Connect(d Info) {
	var err error
	switch d.Driver {
	case TypePosgres:
		SQL, err = connect(d.Connections.Postgres)
		if err != nil {
			log.Println("SQL Driver Error", err)
		}
	default:
		log.Println("No registered database in config")
	}
}
func connect(g congfigure) (*gorm.DB, error) { 
	return g.connect()
}