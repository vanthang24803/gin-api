package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/vanthang24803/api-ecommerce/internal/database"
	"gorm.io/gorm"
)

func init() {
	database.GetDb().AutoMigrate(&UserRole{}, &Role{})

	seedRoles()
}

type UserRole struct {
	UserID uuid.UUID `gorm:"column:user_id;type:uuid;primaryKey"`
	RoleID uuid.UUID `gorm:"column:role_id;type:uuid;primaryKey"`
}

type Role struct {
	ID   uuid.UUID `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"column:name;type:varchar(255);not null" json:"name"`

	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return nil
}

func (Role) TableName() string {
	return "roles"
}

func (UserRole) TableName() string {
	return "user_roles"
}

func seedRoles() {
	db := database.GetDb()

	var roles []Role
	if err := db.Find(&roles).Error; err != nil {
		log.Fatalf("Failed to fetch roles: %v", err)
	}

	if len(roles) == 0 {
		defaultRoles := []Role{
			{Name: "Admin"},
			{Name: "Customer"},
			{Name: "Manager"},
		}

		for _, role := range defaultRoles {
			if err := db.Create(&role).Error; err != nil {
				log.Fatalf("Failed to create role %v: %v", role.Name, err)
			}
		}
		log.Println("Default roles created successfully!")
	}
}
