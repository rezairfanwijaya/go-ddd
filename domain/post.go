package domain

type Post struct {
	Id          int `gorm:"primaryKey"`
	UserId      int
	Title       string
	Description string
}
