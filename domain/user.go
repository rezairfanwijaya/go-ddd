package domain

type User struct {
	Id       int `gorm:"primaryKey"`
	Email    string
	Password string
	Posts    []Post
}
