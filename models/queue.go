package models

import (
	"time"
)

type Queue struct {
	Id        string  `gorm:"type:varchar(300); primaryKey" json:"id"`
	Num     int     `json:"num"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}