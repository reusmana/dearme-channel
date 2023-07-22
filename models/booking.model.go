package models

import (
	"time"
)

type TableLogOrder struct {
	UuidLogOrder string    `json:"uuid_log_order"`
	ChannelCode  string    `json:"channel_code"`
	Orderid      string    `json:"orderid"`
	Response     string    `json:"response"`
	CreatedDate  time.Time `json:"created_date"`
}

type TableLogStock struct {
	UuidLog     string    `json:"uuid_log"`
	Sku         string    `json:"sku"`
	ChannelCode string    `json:"channel_code"`
	Body        string    `json:"body"`
	Response    string    `json:"response"`
	Stock       float64   `json:"stock"`
	CreatedBy   string    `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
}

type TableSalesBooking struct {
	UuidSalesBooking string    `json:"uuid_sales_booking"`
	NoOrder          string    `json:"no_order"`
	SkuNo            string    `json:"sku_no"`
	Uid              string    `json:"uid"`
	RackingNo        string    `json:"racking_no"`
	AisleShelves     string    `json:"aisle_shelves"`
	ProductName      string    `json:"product_name"`
	CreatedBy        string    `json:"created_by"`
	CreatedDate      time.Time `json:"created_date"`
	Flag             int64     `json:"flag"`
}

type TotalRacking struct {
	RackingNo string `json:"racking_no"`
	Total     string `json:"total"`
}

// type DetailOrderPicking struct {
// 	// RackingNo    string `json:"racking_no"`
// 	DetailRack     []DetailOrderPraPickingNew `json:"add_pick"`
// 	DetailLastPick []TableProductUidBasket    `json:"last_pick"`
// 	Qty            string                     `json:"qty"`
// }

type TableProductUidBasket struct {
	UuidProductUid string    `json:"uuid_product_uid" validate:"required"`
	Uid            string    `json:"uid" validate:"required"`
	SkuNo          string    `json:"sku_no" validate:"required"`
	PoNo           string    `json:"po_no" validate:"required"`
	RcvNo          string    `json:"rcv_no" validate:"required"`
	RcvDate        time.Time `json:"rcv_date" validate:"required"`
	UuidStatus     string    `json:"uuid_status" validate:"required"`
	CreatedBy      string    `json:"created_by" validate:"required"`
	CreatedDate    time.Time `json:"created_date"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedDate    time.Time `json:"updated_date"`
	BasketId       string    `json:"basket_id"`
	ProductName    string    `json:"product_name"`
}

type DetailOrderPraPickingNew struct {
	// RackingNo    string `json:"racking_no"`
	RackingNo     interface{}
	SkuNo         string                  `json:"sku_no"`
	AisleShelves  string                  `json:"aisle_shelves"`
	Qty           string                  `json:"qty"`
	ProductName   string                  `json:"product_name"`
	DetailUid     []TableProductUidBasket `json:"detail_uid"`
	PickingWebUid []DetailOrderPraPicking `json:"picking_web_uid"`
}
type UpdateStatusPickingNoOrder struct {
	NoOrder            string `json:"no_order"`
	FlagProcessPicking string `json:"flag_process_picking"`
}

type UpdatetimePick struct {
	UuidPicking    string    `json:"uuid_picking"`
	FinishDatePick time.Time `json:"finish_date_pick"`
}

type GetOrder struct {
	NoOrder string `json:"no_order"`
	Total   string `json:"total"`
}

type DetailOrderPraPicking struct {
	Uid          string    `json:"uid"`
	RackingNo    string    `json:"racking_no"`
	SkuNo        string    `json:"sku_no"`
	ProductName  string    `json:"product_name"`
	RcvNo        string    `json:"rcv_no"`
	RcvDate      time.Time `json:"rcv_date"`
	AisleShelves string    `json:"aisle_shelves"`
	BasketId     string    `json:"basket_id"`
}

type GroupSkuPicking struct {
	SkuNo string `json:"sku_no"`
	Qty   string `json:"qty"`
}

type UpdateProductUid struct {
	UuidProductUid string `json:"uuid_product_uid"`
	UuidStatus     string `json:"uuid_status"`
}

type UpdateProductByUid struct {
	Uid        string `json:"uid"`
	UuidStatus string `json:"uuid_status"`
}

type UpdateStatusOrderCancel struct {
	NoOrder            string    `json:"no_order"`
	StatusProcessOrder string    `json:"status_process_order"`
	UpdatedBy          string    `json:"updated_by"`
	UpdatedDate        time.Time `json:"updated_date"`
}
type UpdateStatusOrder struct {
	NoOrder            string `json:"no_order"`
	StatusProcessOrder string `json:"status_process_order"`
}

type TablePicking struct {
	UuidPicking        string     `json:"uuid_picking"`
	NoOrder            string     `json:"no_order"`
	UuidBasketId       string     `json:"uuid_basket_id"`
	IdMasterUsers      string     `json:"id_master_users"`
	UuidProductUid     string     `json:"uuid_product_uid"`
	StartDatePick      time.Time  `json:"start_date_pick"`
	FinishDatePick     *time.Time `json:"finish_date_pick"`
	FlagProcessPicking string     `json:"flag_process_picking"`
	CreatedBy          string     `json:"created_by"`
	CreatedDate        time.Time  `json:"created_date"`
	UpdatedBy          *string    `json:"updated_by"`
	UpdatedDate        *time.Time `json:"updated_date"`
}

type DetailUidPicking struct {
	UuidProductUid string    `json:"uuid_product_uid" validate:"required"`
	Uid            string    `json:"uid" validate:"required"`
	SkuNo          string    `json:"sku_no" validate:"required"`
	PoNo           string    `json:"po_no" validate:"required"`
	RcvNo          string    `json:"rcv_no" validate:"required"`
	RcvDate        time.Time `json:"rcv_date" validate:"required"`
	UuidStatus     string    `json:"uuid_status" validate:"required"`
	CreatedBy      string    `json:"created_by" validate:"required"`
	CreatedDate    time.Time `json:"created_date"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedDate    time.Time `json:"updated_date"`
	RackingNo      string    `json:"racking_no"`
}

type UpdateRackPick struct {
	UuidRackingNo string `json:"uuid_racking_no" `
	UuidPutway    string `json:"uuid_putway" `
	UuidStatus    string `json:"uuid_status" `
}

type ListSalesOrder struct {
	NoOrder        string `json:"no_order" `
	ChannelCode    string `json:"channel_code" `
	RecipientsName string `json:"recipients_name" `
}
