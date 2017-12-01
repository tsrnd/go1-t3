package models

import (
	"time"
	"github.com/goweb3/app/shared/database"
	"database/sql"
	
)
type Product struct {
	Id int	     			`db:"id" bson:"id"`
	Name string	     		`db:"name" bson:"name"`
	Description string	    `db:"description" bson:"description"`
	Quantity int	    	`db:"quantity" bson:"quantity"`
	Price int				`db:"price" bson:"price"`	
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	DeletedAt time.Time     `db:"deleted_at" bson:"deleted_at"`
}
var (
	Relations []string
)
const (
	TableName = "products"
)

/**
*
*
**/
func (product *Product) GetByCategory(idCategory int) (products []Product, err error) {
	statement := "select "+TableName+".* from "+TableName+" join product_categories on product_categories.product_id = products.id where product_categories.category_id=$1"
	rows, err := database.SQL.Query(statement, idCategory)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
        product := Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Quantity, &product.Price, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			return
		}
        products = append(products, product)
    }
	return
}

/**
*
*
**/
func (product *Product) Limit(limit int) (products []Product, err error) {

	statement := "select "+TableName+".* from "+TableName+" limit $1"
	rows, err := database.SQL.Query(statement, limit)
	if err != nil {
		return
	}
	defer rows.Close()
	var a_string sql.NullString
	for rows.Next() {
        product := Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Quantity, &product.Price, &product.CreatedAt, &product.UpdatedAt, &a_string)
		if err != nil {
			return
		}
        products = append(products, product)
    }
	return
}

func (product *Product) With(relations []string) (err error) {
	return
}
