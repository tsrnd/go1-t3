package models

import (
	"time"

	"github.com/goweb3/app/shared/database"
)

type Product struct {
	BaseModel
	Name          string `db:"name"`
	Description   string `db:"description"`
	Quantity      int    `db:"quantity"`
	Price         int    `db:"price"`
	ProductImages []ProductImage
}

func (this *Product) Create() (err error) {
	statement := "insert into products (name, description, quantity, price, created_at, updated_at) values ($1, $2, $3, $4, $5, $6) returning id"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(this.Name, this.Description, this.Quantity, this.Price, time.Now(), time.Now()).Scan(&this.ID)
	if err != nil {
		return
	}
	return
}

func (this *Product) GetAll(limit int) (products []Product, err error) {
	statement := "select Id, name, description, quantity, price, created_at, updated_at from products where deleted_at is null limit $1"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	rows, err := stmt.Query(limit)
	if err != nil {
		return
	}
	defer stmt.Close()
	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Quantity, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return
}

/**
*
*
**/
func (product *Product) GetTopProducts() (products []Product, err error) {
	// err = database.SQL.Limit(9).Preload("ProductImages").Find(&products).Error
	return
}

/**
*	Find product by product id
**/
func (this *Product) FindByID(id uint) (err error) {
	err = database.SQL.QueryRow("select id, name, description, quantity, price, created_at, updated_at from products where id = $1 and deleted_at is null", id).Scan(&this.ID, &this.Name, &this.Description, &this.Quantity, &this.Price, &this.CreatedAt, &this.UpdatedAt)
	return
}

/**
*	Find product by product list id
**/
func (product *Product) FindByListID(id uint) error {
	var err error
	// err = database.SQL.Where("id = ?", id).First(&product).Error
	return err
}

/**
*
* Load ProductImages
**/
func (product *Product) LoadProductImage() (err error) {
	rows, err := database.SQL.Query("select id, product_id, image form product_images where deleted_at is null AND product_id = $1", product.ID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		productImage := ProductImage{}
		err := rows.Scan(&productImage.ID, &productImage.ProductID, &productImage.Image)
		if err != nil {
			return err
		}
		product.ProductImages = append(product.ProductImages, productImage)
	}
	return
}
