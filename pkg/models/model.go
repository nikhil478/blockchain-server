package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"unique"`
}

type Ic struct {
    gorm.Model
    Name string
    IcIp string
    IcNgrokUrl string
    WalletId string
}
