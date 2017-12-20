package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// blank import.
	_ "github.com/lib/pq"
)

// SQLHandler struct.
type SQLHandler struct {
	Database    *gorm.DB
}

// NewSQLHandler returns new SQLHandler.
// repository: https://github.com/dalu/i18n
func NewSQLHandler() *SQLHandler {
	dbms := GetConfigString("database.dbms")
	host := GetConfigString("database.host")
	user := GetConfigString("database.user")
	pass := GetConfigString("database.pass")
	name := GetConfigString("database.name")
	fmt.Printf(dbms)

	connect := "host=" + host + " user=" + user + " dbname=" + name + " sslmode=disable password=" + pass
	db, err := gorm.Open(dbms, connect)
	// Disable table name's pluralization globally
	// if set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected
	db.SingularTable(true)
	if err != nil {
		panic(err)
	}

	return &SQLHandler{db}
}
