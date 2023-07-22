package models

import (
	"time"
)

type TableBufferStock struct {
	SkuNo       string    `json:"sku_no" `
	BufferStock int64     `json:"buffer_stock" `
	CreatedBy   string    `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedDate time.Time `json:"updated_date"`
}

type TableSkuStock struct {
	UuidSkuStock string    `json:"uuid_sku_stock"`
	SkuNo        string    `json:"sku_no"`
	Stock        string    `json:"stock"`
	CreatedBy    string    `json:"created_by"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedBy    string    `json:"updated_by"`
	UpdatedDate  time.Time `json:"updated_date"`
}

type TableChangeSku struct {
	NoOrder     string `json:"no_order"`
	SkuOrigin   string `json:"sku_origin"`
	SkuReplace  string `json:"sku_replace"`
	ChannelCode string `json:"channel_code"`
	Key1        string `json:"key1"`
	Key2        string `json:"key2"`
	Key3        string `json:"key3"`
}
