package models

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	BaseModel
	Username    string       `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password    string       `gorm:"type:varchar(255);not null"`
	Role        UserRole     `gorm:"type:varchar(20);not null;default:'user';check:role IN ('admin','user')"`
	Purchasings []Purchasing `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL"` // atau CASCADE sesuai kebutuhan
}
