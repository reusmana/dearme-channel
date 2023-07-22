package models

//"github.com/google/uuid"

type AuhtTiktok struct {
	Message string      `json:"message"`
	Data    TokenTiktok `json:"data"`
}

type HeadRespTiktok struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}

type TokenTiktok struct {
	AccessToken          string `json:"access_token"`
	AccessTokenExpireIn  int64  `json:"access_token_expire_in"`
	RefreshToken         string `json:"refresh_token"`
	RefreshTokenExpireIn int64  `json:"refresh_token_expire_in"`
	OpenId               string `json:"open_id"`
	SellerName           string `json:"seller_name"`
}

type RespOrdersTiktok struct {
	Code      int64                   `json:"code"`
	Message   string                  `json:"message"`
	RequestId string                  `json:"request_id"`
	Data      RespOrdersDetailsTiktok `json:"data"`
}
type RespOrdersDetailsTiktok struct {
	OrderList  []OrderListsTiktok `json:"order_list"`
	More       bool               `json:"more"`
	NextCursor string             `json:"next_cursor"`
	Total      int                `json:"total"`
}

type OrderListsTiktok struct {
	OrderId     string `json:"order_id"`
	OrderStatus int64  `json:"order_status"`
	UpdateTime  int64  `json:"update_time"`
}

//ORDER DETAIL
type RespOrderDetailTiktok struct {
	Code      int64                  `json:"code"`
	Message   string                 `json:"message"`
	RequestId string                 `json:"request_id"`
	Data      RespOrderDetailsTiktok `json:"data"`
}

type RespOrderDetailsTiktok struct {
	OrderList []OrderListDetailTiktok `json:"order_list"`
}

type OrderListDetailTiktok struct {
	OrderId                string                 `json:"order_id"`
	OrderStatus            int64                  `json:"order_status"`
	UpdateTime             int64                  `json:"update_time"`
	CreateTime             string                 `json:"create_time"`
	PaymentMethod          string                 `json:"payment_method"`
	DeliveryOption         string                 `json:"delivery_option"`
	ShippingProvider       string                 `json:"shipping_provider"`
	ShippingProviderId     string                 `json:"shipping_provider_id"`
	PaidTime               int64                  `json:"paid_time"`
	BuyerMessage           string                 `json:"buyer_message"`
	PaymentInfo            PaymentInfoTiktok      `json:"payment_info"`
	RecipientAddress       RecipientAddressTiktok `json:"recipient_address"`
	TrackingNumber         string                 `json:"tracking_number"`
	ItemList               []ItemListTiktok       `json:"item_list"`
	RtsTime                int                    `json:"rts_time"`
	RtsSla                 int                    `json:"rts_sla"`
	TtsSla                 int                    `json:"tts_sla"`
	CancelOrderSla         int                    `json:"cancel_order_sla"`
	ReceiverAddressUpdated int                    `json:"receiver_address_updated"`
	PackageList            []PackageListTiktok    `json:"package_list"`
}

type PackageListTiktok struct {
	PackageId string `json:"package_id"`
}

type ItemListTiktok struct {
	SkuId        string  `json:"sku_id"`
	ProductId    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	SkuName      string  `json:"sku_name"`
	Quantity     int     `json:"quantity"`
	SellerSku    string  `json:"seller_sku"`
	SkuSalePrice float64 `json:"sku_sale_price"`
}

type RecipientAddressTiktok struct {
	FullAddress     string   `json:"full_address"`
	Region          string   `json:"region"`
	State           string   `json:"state"`
	City            string   `json:"city"`
	District        string   `json:"district"`
	Town            string   `json:"town"`
	Phone           string   `json:"phone"`
	Name            string   `json:"name"`
	Zipcode         string   `json:"zipcode"`
	AddressDetail   string   `json:"address_detail"`
	AddressLineList []string `json:"address_line_list"`
}
type PaymentInfoTiktok struct {
	Currency                    string `json:"currency"`
	SubTotal                    int64  `json:"sub_total"`
	ShippingFee                 int64  `json:"shipping_fee"`
	SellerDiscount              int64  `json:"seller_discount"`
	TotalAmount                 int64  `json:"total_amount"`
	OriginalTotalProductPrice   int64  `json:"original_total_product_price"`
	OriginalShippingFee         int64  `json:"original_shipping_fee"`
	ShippingFeeSellerDiscount   int64  `json:"shipping_fee_seller_discount"`
	ShippingFeePlatformDiscount int64  `json:"shipping_fee_platform_discount"`
}

//ORDER DETAIL

//PRODUCTS
type RespProdutcsTiktok struct {
	Code      int64              `json:"code"`
	Message   string             `json:"message"`
	RequestId string             `json:"request_id"`
	Data      RespProductsTiktok `json:"data"`
}

type RespProductsTiktok struct {
	Total    int64                `json:"total"`
	Products []ProductsListTiktok `json:"products"`
}

type ProductsListTiktok struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Status      int             `json:"status"`
	SaleRegions []string        `json:"sale_regions"`
	Skus        []ListSkuTiktok `json:"skus"`
}

type ListSkuTiktok struct {
	Id              string                  `json:"id"`
	SellerSku       string                  `json:"seller_sku"`
	Price           PriceSkuListTiktok      `json:"price"`
	StockInfos      []StockInfosTiktok      `json:"stock_infos"`
	SalesAttributes []SalesAttributesTiktok `json:"sales_attributes"`
}

type SalesAttributesTiktok struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ValueId   string `json:"value_id"`
	ValueName string `json:"value_name"`
}

type PriceSkuListTiktok struct {
	OriginalPrice string `json:"original_price"`
	Currency      string `json:"currency"`
}

type StockInfosTiktok struct {
	WarehouseId    string `json:"warehouse_id"`
	AvailableStock int    `json:"available_stock"`
}

//PRODUCTS

//PRODUCT DETAIL
type RespProdutcDetailTiktok struct {
	Code      int64                   `json:"code"`
	Message   string                  `json:"message"`
	RequestId string                  `json:"request_id"`
	Data      RespProductDetailTiktok `json:"data"`
}

type RespProductDetailTiktok struct {
	ProductId          string               `json:"product_id"`
	ProductStatus      int64                `json:"product_status"`
	ProductName        string               `json:"product_name"`
	CategoryListTiktok []CategoryListTiktok `json:"category_list"`
	Brand              BrandTiktok          `json:"brand"`
	Skus               []ListSkuTiktok      `json:"skus"`
}

type BrandTiktok struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type CategoryListTiktok struct {
	Id               string `json:"id"`
	ParentId         string `json:"parent_id"`
	LocalDisplayName string `json:"local_display_name"`
	IsLeaf           bool   `json:"is_leaf"`
}

//PRODUCT DETAIL

//TRACKING INFO

type RespTrackingTiktok struct {
	Code      int64                  `json:"code"`
	Message   string                 `json:"message"`
	RequestId string                 `json:"request_id"`
	Data      RespTrackingListTiktok `json:"data"`
}

type RespTrackingListTiktok struct {
	TrackingInfoList []TrackingInfoListDetailTiktok `json:"tracking_info_list"`
}

type TrackingInfoListDetailTiktok struct {
	TrackingInfo []TrackingInfoTiktok `json:"tracking_info"`
}

type TrackingInfoTiktok struct {
	UpdateTime  int64  `json:"update_time"`
	Description string `json:"description"`
}

//TRACKING INFO

type PackageTiktok struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Data      struct {
		PackageList []struct {
			PackageID   string   `json:"package_id"`
			OrderIDList []string `json:"order_id_list"`
		} `json:"package_list"`
	} `json:"data"`
}

type RespOrderDetailTiktok2 struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Data      struct {
		OrderList []struct {
			BuyerMessage    string `json:"buyer_message"`
			BuyerUID        string `json:"buyer_uid"`
			CancelOrderSLA  int    `json:"cancel_order_sla"`
			CreateTime      string `json:"create_time"`
			DeliveryOption  string `json:"delivery_option"`
			ExtStatus       int    `json:"ext_status"`
			FulfillmentType int    `json:"fulfillment_type"`
			ItemList        []struct {
				ProductID                string  `json:"product_id"`
				ProductName              string  `json:"product_name"`
				Quantity                 int     `json:"quantity"`
				SellerSku                string  `json:"seller_sku"`
				SkuDisplayStatus         int     `json:"sku_display_status"`
				SkuExtStatus             int     `json:"sku_ext_status"`
				SkuID                    string  `json:"sku_id"`
				SkuImage                 string  `json:"sku_image"`
				SkuName                  string  `json:"sku_name"`
				SkuOriginalPrice         int     `json:"sku_original_price"`
				SkuPlatformDiscount      int     `json:"sku_platform_discount"`
				SkuPlatformDiscountTotal int     `json:"sku_platform_discount_total"`
				SkuRtsTime               int     `json:"sku_rts_time"`
				SkuSalePrice             float64 `json:"sku_sale_price"`
				SkuSellerDiscount        int     `json:"sku_seller_discount"`
				SkuType                  int     `json:"sku_type"`
			} `json:"item_list"`
			OrderID       string `json:"order_id"`
			OrderLineList []struct {
				OrderLineID string `json:"order_line_id"`
				SkuID       string `json:"sku_id"`
			} `json:"order_line_list"`
			OrderStatus int `json:"order_status"`
			PackageList []struct {
				PackageID string `json:"package_id"`
			} `json:"package_list"`
			PaymentInfo struct {
				Currency                    string `json:"currency"`
				OriginalShippingFee         int    `json:"original_shipping_fee"`
				OriginalTotalProductPrice   int    `json:"original_total_product_price"`
				PlatformDiscount            int    `json:"platform_discount"`
				SellerDiscount              int    `json:"seller_discount"`
				ShippingFee                 int    `json:"shipping_fee"`
				ShippingFeePlatformDiscount int    `json:"shipping_fee_platform_discount"`
				ShippingFeeSellerDiscount   int    `json:"shipping_fee_seller_discount"`
				SubTotal                    int    `json:"sub_total"`
				Taxes                       int    `json:"taxes"`
				TotalAmount                 int    `json:"total_amount"`
			} `json:"payment_info"`
			PaymentMethod          string `json:"payment_method"`
			ReceiverAddressUpdated int    `json:"receiver_address_updated"`
			RecipientAddress       struct {
				AddressDetail   string   `json:"address_detail"`
				AddressLineList []string `json:"address_line_list"`
				City            string   `json:"city"`
				District        string   `json:"district"`
				FullAddress     string   `json:"full_address"`
				Name            string   `json:"name"`
				Phone           string   `json:"phone"`
				Region          string   `json:"region"`
				RegionCode      string   `json:"region_code"`
				State           string   `json:"state"`
				Town            string   `json:"town"`
				Zipcode         string   `json:"zipcode"`
			} `json:"recipient_address"`
			RtsSLA             int    `json:"rts_sla"`
			RtsTime            int    `json:"rts_time"`
			ShippingProvider   string `json:"shipping_provider"`
			ShippingProviderID string `json:"shipping_provider_id"`
			TrackingNumber     string `json:"tracking_number"`
			TtsSLA             int    `json:"tts_sla"`
			UpdateTime         int    `json:"update_time"`
			WarehouseID        string `json:"warehouse_id"`
		} `json:"order_list"`
	} `json:"data"`
}

type ConfigPickupTiktok struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Data      struct {
		DropOffPointURL string `json:"drop_off_point_url"`
		IsDropOff       bool   `json:"is_drop_off"`
		IsPickUp        bool   `json:"is_pick_up"`
		PickUpTimeList  []struct {
			EndTime   string `json:"end_time"`
			StartTime string `json:"start_time"`
		} `json:"pick_up_time_list"`
	} `json:"data"`
}

type DocumentShipTiktok struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Data      struct {
		DocURL string `json:"doc_url"`
	} `json:"data"`
}
