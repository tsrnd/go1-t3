package models

import (
)
type Category struct {
	BaseModel
	Name string	`db:"name"`
}
