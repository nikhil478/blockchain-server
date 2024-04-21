package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"unique"`
    Password string
    Role string
    WalletID string
    PaymailID string
}

type Ic struct {
    gorm.Model
    Name string
    IcIp string
    IcNgrokUrl string
    WalletId string
}

type Wallet struct {
    WalletID string
    PaymailID string
}
