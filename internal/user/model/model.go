package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id       int    `json:"id"  gorm:"primaryKey"`
	Username string `json:"username"   gorm:"type:varchar(100);unique"`
	//Username     string `sql:"not null;unique"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	Addresses []Address `json:"addresses" gorm:"foreignKey: UserId"`
	//PhotoId   string    `json:"photoId"`
	//Photo     Photo `json:"photo" gorm:"foreignKey:PhotoId"`
	Photo     string `json:"photo" gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

// type Photo struct {
// 	Id        int    `json:"id"  gorm:"primaryKey"`
// 	Photo     string `json:"photo" gorm:"type:varchar(100);"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt sql.NullTime `gorm:"index"`
// }
type Address struct {
	AddressId int    `json:"id" gorm:"primaryKey"`
	City      string `json:"city"   gorm:"type:varchar(100);unique"`
	//Username     string `sql:"not null;unique"`
	Subcity   string `json:"subcity"`
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
type ResponseUser struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
}
