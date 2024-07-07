package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email      string `gorm:"unique"`
    Blocked    []User `gorm:"many2many:user_blocks"`
    Friendships []User `gorm:"many2many:friendships"`
}
