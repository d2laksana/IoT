package entity

import (
	"time"

	"gorm.io/gorm"
)

type Suhu struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Temp      float64        `json:"temp"`
	Hum       float64        `json:"hum"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
