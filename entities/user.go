package entities

type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string //`json:"name"`
	Email string //`json:"email"`

}
