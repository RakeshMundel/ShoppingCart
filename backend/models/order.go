package models

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    UserID uint
    CartID uint
}
