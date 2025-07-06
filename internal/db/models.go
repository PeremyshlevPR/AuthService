package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username     string    `gorm:"unique;not null;size:150"`
	Email        string    `gorm:"unique;not null;size:255"`
	PasswordHash string    `gorm:"not null" json:"-"`

	Permissions []Permission `gorm:"many2many:user_permissions;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Permission struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string    `gorm:"unique;not null;size:100"`
	Description string    `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserPermission struct {
	UserID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	PermissionID uuid.UUID `gorm:"type:uuid;primaryKey"`

	CreatedAt time.Time
	CreatedBy User
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
