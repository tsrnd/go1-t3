package models

import (
)

type ProductImage struct {
	BaseModel
	ProductID uint    `db:"product_id"`
	Image     string  `db:"image"`
	Product   Product
}
