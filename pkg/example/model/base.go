package model

import (
	"time"
)

// tablePrefix, every model struct has TableName() specified name in database
var tablePrefix = "t_"

// BaseModel has default fields for tables.
// gorm has default `gorm.Model` defines more fields
// you can use/create what you need.
type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
