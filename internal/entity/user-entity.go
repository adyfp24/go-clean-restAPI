package entity

type User struct {
	ID      int    `gorm:"primaryKey;"`
	Name    string `gorm:"not null;unique"`
	Age     int
	Address string `gorm:"not null"`
	Role    string `gorm:"not null"`
}
