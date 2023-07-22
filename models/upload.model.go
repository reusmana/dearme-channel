package models

import (
	"time"
)

type TableUploadImageChannel struct {
	UuidUploadImage string    `json:"uuid_upload_image"`
	NameFile        string    `json:"name_file"`
	FileBase64      string    `json:"file_base64"`
	CreatedBy       string    `json:"created_by"`
	CreatedDate     time.Time `json:"created_date"`
	PathImage       string    `json:"path_image"`
	ImageId         string    `json:"image_id"`
	ImageUrl        string    `json:"image_url"`
	ChannelCode     string    `json:"channel_code"`
}
