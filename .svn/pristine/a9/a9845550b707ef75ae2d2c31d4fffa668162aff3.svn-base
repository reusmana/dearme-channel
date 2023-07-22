package models

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	Token           string      `json:"auth_bearer_token"`
	RefreshToken    string      `json:"refresh_bearer_token"`
	IdMasterUsers   uuid.UUID   `json:"id_master_users"`
	IdMasterRoles   string      `json:"id_master_roles"`
	Username        string      `json:"username"`
	Fullname        string      `json:"fullname"`
	Email           string      `json:"email"`
	EmailVerifiedAt *string     `json:"email_verified_at"`
	Password        string      `json:"password"`
	UrlPhoto        *string     `json:"url_photo"`
	LoginCount      *int        `json:"login_count"`
	RememberToken   *string     `json:"remember_token"`
	Sequence        *int        `json:"sequence"`
	IsActive        *int        `json:"is_active"`
	CreatedBy       *string     `json:"created_by"`
	CreatedDate     *time.Time  `json:"created_date"`
	MasterModulars  interface{} `json:"master_modulars"`
}
