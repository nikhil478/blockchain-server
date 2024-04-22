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
    IcIp string
    IcNgrokUrl string
    WalletID string
    PaymailID string
}

type Wallet struct {
    WalletID string
    PaymailID string
}
