package model

import "github.com/jinzhu/gorm"

type Rider struct {
	gorm.Model
	Name     string
	Assigned bool
}
