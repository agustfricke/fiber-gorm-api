package models

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    Body string `json:"body"`
}
