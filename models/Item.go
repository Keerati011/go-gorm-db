// models/Item.go
package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model //gorm create ID, CreatedAt,UpdatedAt, DeletedAt,
	Name  string
	Price float64
}
