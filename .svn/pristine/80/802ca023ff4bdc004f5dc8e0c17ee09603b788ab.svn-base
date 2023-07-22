package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	IdMasterUsers   uuid.UUID  `json:"id_master_users"`
	IdMasterRoles   string     `json:"id_master_roles"`
	Username        string     `json:"username"`
	Fullname        string     `json:"fullname"`
	Email           string     `json:"email"`
	EmailVerifiedAt *string    `json:"email_verified_at"`
	Password        string     `json:"password"`
	UrlPhoto        *string    `json:"url_photo"`
	LoginCount      *int       `json:"login_count"`
	RememberToken   *string    `json:"remember_token"`
	Sequence        *int       `json:"sequence"`
	IsActive        *int       `json:"is_active"`
	CreatedBy       *string    `json:"created_by"`
	CreatedDate     *time.Time `json:"created_date"`
	UpdatedBy       *string    `json:"updated_by"`
	UpdatedDate     *time.Time `json:"updated_date"`
}
