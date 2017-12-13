package models

import (
	"github.com/goweb3/app/shared/database"
)

type CartProduct struct {
	BaseModel
	CartID    uint `db:"cart_id"`
	ProductID uint `db:"product_id"`
	Quantity  uint `db:"quantity"`
	Product   Product
}

/**
*
* Price follow quantity
**/
func (cartProduct *CartProduct) PriceFollowQuantity() uint {
	return uint(cartProduct.Quantity) * uint(cartProduct.Product.Price)
}

/**
*
* Load Product
**/
func (cartProduct *CartProduct) LoadProducts() (err error) {
	err = database.SQL.QueryRow("SELECT id, name, description, quantity, price FROM products WHERE deleted_at is null AND id = $1", cartProduct.ProductID).Scan(&cartProduct.Product.ID, &cartProduct.Product.Name, &cartProduct.Product.Description, &cartProduct.Product.Quantity, &cartProduct.Product.Price)
	return
}

/**
*
* Delete cart product
**/
func (cartProduct *CartProduct) Delete(cartProductID uint) (err error) {
	_, err = database.SQL.Exec("DELETE FROM cart_products WHERE id = $1", cartProductID)
	return err
}

/**
*
* Create cart product
**/
func (cartProduct *CartProduct) Create() (err error) {
	statement := "INSERT INTO cart_products (cart_id, product_id, quantity) values ($1, $2, $3) returning id"
	stmt, err := database.SQL.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(cartProduct.CartID, cartProduct.ProductID, cartProduct.Quantity).Scan(&cartProduct.ID)
	return
}

/**
*
* Update cart product
**/
func (cartProduct *CartProduct) Update(quantity uint, cartID uint, productID uint) (err error) {
	_, err = database.SQL.Exec("UPDATE cart_products SET quantity = $1 WHERE cart_id = $2 AND product_id = $3", quantity, cartID, productID)
	return
}

/**
*
* Find cart product by cart id and product id
**/
func (cartProduct *CartProduct) FindByCartIDAndProductID(cartID uint, productID uint) (err error) {
	row := database.SQL.QueryRow("SELECT id, cart_id, product_id, quantity FROM cart_products WHERE deleted_at is null AND cart_id = $1 AND product_id = $2", cartID, productID)
	err = row.Scan(&cartProduct.ID, &cartProduct.CartID, &cartProduct.ProductID, &cartProduct.Quantity)
	return err
}

/**
*
* Find cart products by cart id
**/
func (cartProduct *CartProduct) FindByCartID(cartID uint) (cartProducts []CartProduct) {
	// database.SQL.Where("cart_id = ?", cartID).Find(&cartProducts)
	return
}

/**
*
* Get all cart products by cart id
**/
func (cartProduct *CartProduct) GetByCartID(cartID uint) (cartProducts []CartProduct, err error) {
	rows, err := database.SQL.Query("SELECT id, cart_id, product_id, quantity FROM cart_products WHERE deleted_at is null AND cart_id = $1", cartID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		cartProduct := CartProduct{}
		err := rows.Scan(&cartProduct.ID, &cartProduct.CartID, &cartProduct.ProductID, &cartProduct.Quantity)
		if err != nil {
			return nil, err
		}
		cartProduct.LoadProduct(cartProduct.ProductID)
		cartProducts = append(cartProducts, cartProduct)
	}
	return
}

/**
*
* Load product
**/
func (cartProduct *CartProduct) LoadProduct(productID uint) (err error) {
	product := Product{}
	rows := database.SQL.QueryRow("SELECT id, name, description, quantity, price FROM products WHERE deleted_at is null AND id = $1", productID)
	err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Quantity, &product.Price)
	if err != nil {
		return err
	}
	cartProduct.Product = product
	return
}
