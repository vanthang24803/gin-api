package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/vanthang24803/api-ecommerce/internal/database"
	"gorm.io/gorm"
)

func init() {
	database.GetDb().AutoMigrate(&User{})
}

type User struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Email     string    `gorm:"column:email;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"column:password;not null" json:"-"`
	FirstName string    `gorm:"column:first_name;not null" json:"firstName"`
	LastName  string    `gorm:"column:last_name;not null" json:"lastName"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`

	Tokens []Token `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	Roles []Role `gorm:"many2many:user_roles;joinForeignKey:UserId;joinReferences:RoleId;constraint:OnCreate:CASCADE" json:"roles,omitempty"`

	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}

func (User) TableName() string {
	return "users"
}
