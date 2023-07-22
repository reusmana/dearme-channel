package models

import (
	"time"
)

type TableChannel struct {
	UuidChannel string    `json:"uuid_channel" `
	ChannelCode string    `json:"channel_code" `
	ChannelName string    `json:"channel_name" `
	ChannelAddr string    `json:"channel_addr" `
	UuidStatus  string    `json:"uuid_status" `
	CreatedBy   string    `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedDate time.Time `json:"updated_date"`
	FlagChannel string    `json:"flag_channel"`
}
