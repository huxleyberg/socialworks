package models

type User struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}
