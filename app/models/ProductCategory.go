package models

import (
)

type ProductCategory struct {
	BaseModel
	ProductID uint	`db:"product_id"`
	CategoryID uint	`db:"category_id"`
}
