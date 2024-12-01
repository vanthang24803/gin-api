package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/vanthang24803/api-ecommerce/internal/database"
	"gorm.io/gorm"
)

func init() {
	database.GetDb().AutoMigrate(&Token{})
}

type Token struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Value     string    `gorm:"column:value;type:text" json:"value"`
	UserID    uuid.UUID `gorm:"column:user_id;type:uuid;not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (t *Token) BeforeCreate(ctx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return nil
}

func (Token) TableName() string {
	return "tokens"
}
