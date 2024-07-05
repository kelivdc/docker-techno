package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Admin struct {
		gorm.Model `json:"-"`
		ID         uint      `json:"id" gorm:"primaryKey;"`
		Uuid       uuid.UUID `json:"uuid" gorm:"type:uuid;default:gen_random_uuid()" example:"d60adf12-58e0-4465-baf7-8fb8ef7edbe8"`
		Email      string    `json:"email" gorm:"unique"`
		Name       string    `json:"name"`
		Active     bool      `json:"active"`
		Password   string    `json:"password"`
	}
)