package models

import "gorm.io/gorm"

type Friend struct {
    gorm.Model
    RequestorID uint
    ToID        uint
    Status      string
    Requestor   User `gorm:"foreignKey:RequestorID"`
    To          User `gorm:"foreignKey:ToID"`
}
