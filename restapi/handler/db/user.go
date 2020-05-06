package db

import "github.com/jinzhu/gorm"

const (
	AutoPreLoad = "gorm:auto_preload"
)

type User struct {
	gorm.Model
	Name string
	Role string
}
