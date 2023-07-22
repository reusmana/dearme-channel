package models

import (
	"github.com/google/uuid"
)

type Modulars struct {
	IdMasterModulars     uuid.UUID `json:"id_master_modulars"`
	IdMasterApplications uuid.UUID `json:"id_master_applications"`
	ModularName          string    `json:"modular_name"`
	ModularIcon          string    `json:"modular_icon"`
	Sequence             *int      `json:"sequence"`
	IsActive             *int      `json:"is_active"`
	CreatedBy            *string   `json:"created_by"`
	CreatedDate          *string   `json:"created_date"`
	UpdatedBy            *string   `json:"updated_by"`
	UpdatedDate          *string   `json:"updated_date"`
}
