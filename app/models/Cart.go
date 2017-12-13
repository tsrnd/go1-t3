package models

import "github.com/goweb3/app/shared/database"

// "github.com/goweb3/app/shared/database"

type Cart struct {
	BaseModel
	UserID       uint `db:"user_id"`
	CartProducts []CartProduct
}

/**
*
* Create cart
**/
func (cart *Cart) Create(userID uint) (err error) {
	_, err = database.SQL.Exec("INSERT INTO carts (user_id) values ($1) returning id", userID)
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
	err = database.SQL.QueryRow("SELECT id, user_id FROM carts WHERE user_id = $1", userID).Scan(&cart.ID, &cart.UserID)
	return err
}
