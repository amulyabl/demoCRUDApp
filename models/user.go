package models

import (
	_ "fmt"
	_ "os/user"
	"time"
)

type User struct {
	ID        int    //`gorm:"primaryKey"`
	Name      string //`json:"name"`
	Email     string //`json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}
