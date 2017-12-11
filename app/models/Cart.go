package models

import (
	"github.com/goweb3/app/shared/database"
)

type Cart struct {
	BaseModel
	UserID       uint `db:"user_id"`
	CartProducts []CartProduct
}

/**
*
* Create cart
**/
func (cart *Cart) Create() (err error) {
	// err = database.SQL.Create(&cart).Error
	return
}

/**
*
* Total price cart
**/
func (cart *Cart) TotalPrice() uint {
	sum := 0
	for _, v := range cart.CartProducts {
		sum += int(v.Quantity) * v.Product.Price
	}
	return uint(sum + 20000)
}
/**
*
* Load CartProducts
**/
func (cart *Cart) LoadCartProducts() (err error) {
	rows, err := database.SQL.Query("SELECT id, cart_id, product_id, quantity FROM cart_products WHERE deleted_at is null AND cart_id = $1", cart.ID)
    if err != nil {
        return
    }
    defer rows.Close()
    for rows.Next() {
		cartProduct:= CartProduct{}
        err := rows.Scan(&cartProduct.ID, &cartProduct.CartID, &cartProduct.ProductID, &cartProduct.Quantity)
        if err != nil {
            return err
		}
		cart.CartProducts = append(cart.CartProducts, cartProduct)
    }
	return
}

/**
*
* Delete
**/
func (cart *Cart) Delete() (err error) {
	// err := database.SQL.Delete(&cart).Error
	return err
}

/**
*
* Find cart by user id
**/
func (cart *Cart) FindByUserID(userID uint) (err error) {
	err = database.SQL.QueryRow("SELECT id, user_id FROM carts WHERE deleted_at is null AND user_id = $1", userID).Scan(&cart.ID, &cart.UserID)
	return err
}
