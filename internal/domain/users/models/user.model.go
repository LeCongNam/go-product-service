package models

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email" gorm:"unique"`
	Age   int    `json:"age" binding:"required,gte=0"`
}
