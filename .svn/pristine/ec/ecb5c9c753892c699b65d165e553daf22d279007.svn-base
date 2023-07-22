package models

import "time"

type DetalTransactionLazada struct {
	Data []struct {
		OrderNo             string `json:"order_no"`
		TransactionDate     string `json:"transaction_date"`
		Amount              string `json:"amount"`
		PaidStatus          string `json:"paid_status"`
		ShippingProvider    string `json:"shipping_provider"`
		WHTIncludedInAmount string `json:"WHT_included_in_amount"`
		LazadaSku           string `json:"lazada_sku"`
		FeeType             string `json:"fee_type"`
		TransactionType     string `json:"transaction_type"`
		OrderItemNo         string `json:"orderItem_no"`
		OrderItemStatus     string `json:"orderItem_status"`
		Reference           string `json:"reference"`
		FeeName             string `json:"fee_name"`
		ShippingSpeed       string `json:"shipping_speed"`
		WHTAmount           string `json:"WHT_amount"`
		TransactionNumber   string `json:"transaction_number"`
		SellerSku           string `json:"seller_sku"`
		Statement           string `json:"statement"`
		Details             string `json:"details"`
		VATInAmount         string `json:"VAT_in_amount"`
		ShipmentType        string `json:"shipment_type"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

type ErrorUpdateStockLazada struct {
	Code    string `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Detail  []struct {
		Message   string `json:"message"`
		SellerSku string `json:"seller_sku"`
	} `json:"detail"`
	RequestID string `json:"request_id"`
}
type TokenLazada struct {
	AccessToken      string `json:"access_token"`
	Country          string `json:"country"`
	RefreshToken     string `json:"refresh_token"`
	AccountPlatform  string `json:"account_platform"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	CountryUserInfo  []struct {
		Country   string `json:"country"`
		UserID    string `json:"user_id"`
		SellerID  string `json:"seller_id"`
		ShortCode string `json:"short_code"`
	} `json:"country_user_info"`
	ExpiresIn int    `json:"expires_in"`
	Account   string `json:"account"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

type ErrorLazada struct {
	Type      string `json:"type"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

type ProductsLazada struct {
	Data struct {
		TotalProducts int `json:"total_products"`
		Products      []struct {
			TrialProduct bool     `json:"trialProduct"`
			CreatedTime  string   `json:"created_time"`
			UpdatedTime  string   `json:"updated_time"`
			Images       []string `json:"images"`
			Skus         []struct {
				Status                    string        `json:"Status"`
				Quantity                  int           `json:"quantity"`
				Images                    []interface{} `json:"Images"`
				SellerSku                 string        `json:"SellerSku"`
				ShopSku                   string        `json:"ShopSku"`
				URL                       string        `json:"Url"`
				MultiWarehouseInventories []struct {
					OccupyQuantity   int    `json:"occupyQuantity"`
					Quantity         int    `json:"quantity"`
					TotalQuantity    int    `json:"totalQuantity"`
					WithholdQuantity int    `json:"withholdQuantity"`
					WarehouseCode    string `json:"warehouseCode"`
					SellableQuantity int    `json:"sellableQuantity"`
				} `json:"multiWarehouseInventories"`
				PackageWidth            string        `json:"package_width"`
				PackageHeight           string        `json:"package_height"`
				Size                    string        `json:"size"`
				FblWarehouseInventories []interface{} `json:"fblWarehouseInventories"`
				SpecialPrice            int           `json:"special_price"`
				Price                   int           `json:"price"`
				ChannelInventories      []interface{} `json:"channelInventories"`
				SizeGroup               string        `json:"sizeGroup"`
				PackageLength           string        `json:"package_length"`
				PackageWeight           string        `json:"package_weight"`
				SkuID                   int64         `json:"SkuId"`
			} `json:"skus"`
			ItemID          int64         `json:"item_id"`
			PrimaryCategory int           `json:"primary_category"`
			MarketImages    []interface{} `json:"marketImages"`
			Attributes      struct {
				Name            string `json:"name"`
				Brand           string `json:"brand"`
				FaPattern       string `json:"fa_pattern"`
				FaGeneralStyles string `json:"fa_general_styles"`
				FaSeason        string `json:"fa_season"`
				Sleeves         string `json:"sleeves"`
				IsUnisex        string `json:"is_unisex"`
				Hazmat          string `json:"Hazmat"`
				Source          string `json:"source"`
			} `json:"attributes"`
			Status string `json:"status"`
		} `json:"products"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

type ProductDetailLazada struct {
	Data struct {
		CreatedTime string   `json:"created_time"`
		UpdatedTime string   `json:"updated_time"`
		Images      []string `json:"images"`
		Skus        []struct {
			Status                    string   `json:"Status"`
			Quantity                  int      `json:"quantity"`
			Images                    []string `json:"Images"`
			SellerSku                 string   `json:"SellerSku"`
			ShopSku                   string   `json:"ShopSku"`
			URL                       string   `json:"Url"`
			MultiWarehouseInventories []struct {
				OccupyQuantity   int    `json:"occupyQuantity"`
				Quantity         int    `json:"quantity"`
				TotalQuantity    int    `json:"totalQuantity"`
				WithholdQuantity int    `json:"withholdQuantity"`
				WarehouseCode    string `json:"warehouseCode"`
				SellableQuantity int    `json:"sellableQuantity"`
			} `json:"multiWarehouseInventories"`
			PackageWidth            string   `json:"package_width"`
			PackageHeight           string   `json:"package_height"`
			Size                    string   `json:"size"`
			FblWarehouseInventories []string `json:"fblWarehouseInventories"`
			SpecialPrice            float64  `json:"special_price"`
			Price                   float64  `json:"price"`
			ChannelInventories      []string `json:"channelInventories"`
			SizeGroup               string   `json:"sizeGroup"`
			PackageLength           string   `json:"package_length"`
			PackageWeight           string   `json:"package_weight"`
			Available               int      `json:"Available"`
			SkuID                   int64    `json:"SkuId"`
		} `json:"skus"`
		ItemID    int64 `json:"item_id"`
		Variation struct {
			Variation1 struct {
				Name      string   `json:"name"`
				Label     string   `json:"label"`
				HasImage  bool     `json:"hasImage"`
				Customize bool     `json:"customize"`
				Options   []string `json:"options"`
			} `json:"Variation1"`
		} `json:"variation"`
		TrialProduct    bool     `json:"trialProduct"`
		PrimaryCategory int      `json:"primary_category"`
		MarketImages    []string `json:"marketImages"`
		Attributes      struct {
			Name             string `json:"name"`
			Description      string `json:"description"`
			Brand            string `json:"brand"`
			FaPattern        string `json:"fa_pattern"`
			Sleeves          string `json:"sleeves"`
			ClothingMaterial string `json:"clothing_material"`
			MTopNeckline     string `json:"m_top_neckline"`
			Hazmat           string `json:"Hazmat"`
			Source           string `json:"source"`
		} `json:"attributes"`
		Status string `json:"status"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

type OrdersLazada struct {
	Data struct {
		Count      int `json:"count"`
		CountTotal int `json:"countTotal"`
		Orders     []struct {
			VoucherPlatform             float64  `json:"voucher_platform"`
			Voucher                     float64  `json:"voucher"`
			WarehouseCode               string   `json:"warehouse_code"`
			OrderNumber                 int64    `json:"order_number"`
			VoucherSeller               float64  `json:"voucher_seller"`
			CreatedAt                   string   `json:"created_at"`
			VoucherCode                 string   `json:"voucher_code"`
			GiftOption                  bool     `json:"gift_option"`
			ShippingFeeDiscountPlatform float64  `json:"shipping_fee_discount_platform"`
			CustomerLastName            string   `json:"customer_last_name"`
			PromisedShippingTimes       string   `json:"promised_shipping_times"`
			UpdatedAt                   string   `json:"updated_at"`
			Price                       string   `json:"price"`
			NationalRegistrationNumber  string   `json:"national_registration_number"`
			ShippingFeeOriginal         float64  `json:"shipping_fee_original"`
			PaymentMethod               string   `json:"payment_method"`
			CustomerFirstName           string   `json:"customer_first_name"`
			ShippingFeeDiscountSeller   float64  `json:"shipping_fee_discount_seller"`
			ShippingFee                 float64  `json:"shipping_fee"`
			BranchNumber                string   `json:"branch_number"`
			TaxCode                     string   `json:"tax_code"`
			ItemsCount                  int      `json:"items_count"`
			DeliveryInfo                string   `json:"delivery_info"`
			Statuses                    []string `json:"statuses"`
			AddressBilling              struct {
				Country   string `json:"country"`
				Address3  string `json:"address3"`
				Phone     string `json:"phone"`
				Address2  string `json:"address2"`
				City      string `json:"city"`
				Address1  string `json:"address1"`
				PostCode  string `json:"post_code"`
				Phone2    string `json:"phone2"`
				LastName  string `json:"last_name"`
				Address5  string `json:"address5"`
				Address4  string `json:"address4"`
				FirstName string `json:"first_name"`
			} `json:"address_billing"`
			ExtraAttributes string `json:"extra_attributes"`
			OrderID         int64  `json:"order_id"`
			Remarks         string `json:"remarks"`
			GiftMessage     string `json:"gift_message"`
			AddressShipping struct {
				Country   string `json:"country"`
				Address3  string `json:"address3"`
				Phone     string `json:"phone"`
				Address2  string `json:"address2"`
				City      string `json:"city"`
				Address1  string `json:"address1"`
				PostCode  string `json:"post_code"`
				Phone2    string `json:"phone2"`
				LastName  string `json:"last_name"`
				Address5  string `json:"address5"`
				Address4  string `json:"address4"`
				FirstName string `json:"first_name"`
			} `json:"address_shipping"`
		} `json:"orders"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

type OrdersDetailGenerated struct {
	Data struct {
		Voucher                     float64  `json:"voucher"`
		WarehouseCode               string   `json:"warehouse_code"`
		OrderNumber                 int64    `json:"order_number"`
		CreatedAt                   string   `json:"created_at"`
		VoucherCode                 string   `json:"voucher_code"`
		GiftOption                  bool     `json:"gift_option"`
		ShippingFeeDiscountPlatform float64  `json:"shipping_fee_discount_platform"`
		CustomerLastName            string   `json:"customer_last_name"`
		UpdatedAt                   string   `json:"updated_at"`
		PromisedShippingTimes       string   `json:"promised_shipping_times"`
		Price                       string   `json:"price"`
		NationalRegistrationNumber  string   `json:"national_registration_number"`
		ShippingFeeOriginal         float64  `json:"shipping_fee_original"`
		PaymentMethod               string   `json:"payment_method"`
		CustomerFirstName           string   `json:"customer_first_name"`
		ShippingFeeDiscountSeller   float64  `json:"shipping_fee_discount_seller"`
		ShippingFee                 float64  `json:"shipping_fee"`
		BranchNumber                string   `json:"branch_number"`
		TaxCode                     string   `json:"tax_code"`
		ItemsCount                  int      `json:"items_count"`
		DeliveryInfo                string   `json:"delivery_info"`
		Statuses                    []string `json:"statuses"`
		AddressBilling              struct {
			Country   string `json:"country"`
			Address3  string `json:"address3"`
			Address2  string `json:"address2"`
			City      string `json:"city"`
			Phone     string `json:"phone"`
			Address1  string `json:"address1"`
			PostCode  string `json:"post_code"`
			Phone2    string `json:"phone2"`
			LastName  string `json:"last_name"`
			Address5  string `json:"address5"`
			Address4  string `json:"address4"`
			FirstName string `json:"first_name"`
		} `json:"address_billing"`
		ExtraAttributes string `json:"extra_attributes"`
		OrderID         int64  `json:"order_id"`
		GiftMessage     string `json:"gift_message"`
		Remarks         string `json:"remarks"`
		AddressShipping struct {
			Country   string `json:"country"`
			Address3  string `json:"address3"`
			Address2  string `json:"address2"`
			City      string `json:"city"`
			Phone     string `json:"phone"`
			Address1  string `json:"address1"`
			PostCode  string `json:"post_code"`
			Phone2    string `json:"phone2"`
			LastName  string `json:"last_name"`
			Address5  string `json:"address5"`
			Address4  string `json:"address4"`
			FirstName string `json:"first_name"`
		} `json:"address_shipping"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

type OrderDetailItemLazadaOld struct {
	Data []struct {
		PickUpStoreInfo struct {
		} `json:"pick_up_store_info"`
		TaxAmount                   float64   `json:"tax_amount"`
		Reason                      string    `json:"reason"`
		SLATimeStamp                time.Time `json:"sla_time_stamp"`
		VoucherSeller               int       `json:"voucher_seller"`
		PurchaseOrderID             string    `json:"purchase_order_id"`
		VoucherCodeSeller           string    `json:"voucher_code_seller"`
		VoucherCode                 string    `json:"voucher_code"`
		PackageID                   string    `json:"package_id"`
		BuyerID                     int64     `json:"buyer_id"`
		Variation                   string    `json:"variation"`
		ProductID                   string    `json:"product_id"`
		VoucherCodePlatform         string    `json:"voucher_code_platform"`
		PurchaseOrderNumber         string    `json:"purchase_order_number"`
		Sku                         string    `json:"sku"`
		OrderType                   string    `json:"order_type"`
		InvoiceNumber               string    `json:"invoice_number"`
		CancelReturnInitiator       string    `json:"cancel_return_initiator"`
		ShopSku                     string    `json:"shop_sku"`
		IsReroute                   int       `json:"is_reroute"`
		StagePayStatus              string    `json:"stage_pay_status"`
		SkuID                       string    `json:"sku_id"`
		TrackingCodePre             string    `json:"tracking_code_pre"`
		OrderItemID                 int64     `json:"order_item_id"`
		ShopID                      string    `json:"shop_id"`
		OrderFlag                   string    `json:"order_flag"`
		IsFbl                       int       `json:"is_fbl"`
		Name                        string    `json:"name"`
		DeliveryOptionSof           int       `json:"delivery_option_sof"`
		OrderID                     int64     `json:"order_id"`
		Status                      string    `json:"status"`
		ProductMainImage            string    `json:"product_main_image"`
		VoucherPlatform             int       `json:"voucher_platform"`
		PaidPrice                   float64   `json:"paid_price"`
		ProductDetailURL            string    `json:"product_detail_url"`
		WarehouseCode               string    `json:"warehouse_code"`
		PromisedShippingTime        string    `json:"promised_shipping_time"`
		ShippingType                string    `json:"shipping_type"`
		CreatedAt                   string    `json:"created_at"`
		VoucherSellerLpi            int       `json:"voucher_seller_lpi"`
		ShippingFeeDiscountPlatform int       `json:"shipping_fee_discount_platform"`
		WalletCredits               int       `json:"wallet_credits"`
		UpdatedAt                   string    `json:"updated_at"`
		Currency                    string    `json:"currency"`
		ShippingProviderType        string    `json:"shipping_provider_type"`
		VoucherPlatformLpi          int       `json:"voucher_platform_lpi"`
		ShippingFeeOriginal         float64   `json:"shipping_fee_original"`
		ItemPrice                   float64   `json:"item_price"`
		IsDigital                   int       `json:"is_digital"`
		ShippingServiceCost         int       `json:"shipping_service_cost"`
		TrackingCode                string    `json:"tracking_code"`
		ShippingFeeDiscountSeller   int       `json:"shipping_fee_discount_seller"`
		ShippingAmount              float64   `json:"shipping_amount"`
		ReasonDetail                string    `json:"reason_detail"`
		ReturnStatus                string    `json:"return_status"`
		ShipmentProvider            string    `json:"shipment_provider"`
		VoucherAmount               int       `json:"voucher_amount"`
		DigitalDeliveryInfo         string    `json:"digital_delivery_info"`
		ExtraAttributes             string    `json:"extra_attributes"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

type OrderDetailItemLazada struct {
	Data []struct {
		PickUpStoreInfo struct {
		} `json:"pick_up_store_info"`
		TaxAmount                   float64 `json:"tax_amount"`
		Reason                      string  `json:"reason"`
		SLATimeStamp                string  `json:"sla_time_stamp"`
		VoucherSeller               float64 `json:"voucher_seller"`
		PurchaseOrderID             string  `json:"purchase_order_id"`
		VoucherCodeSeller           string  `json:"voucher_code_seller"`
		VoucherCode                 string  `json:"voucher_code"`
		PackageID                   string  `json:"package_id"`
		BuyerID                     int64   `json:"buyer_id"`
		Variation                   string  `json:"variation"`
		ProductID                   string  `json:"product_id"`
		VoucherCodePlatform         string  `json:"voucher_code_platform"`
		PurchaseOrderNumber         string  `json:"purchase_order_number"`
		Sku                         string  `json:"sku"`
		OrderType                   string  `json:"order_type"`
		InvoiceNumber               string  `json:"invoice_number"`
		CancelReturnInitiator       string  `json:"cancel_return_initiator"`
		ShopSku                     string  `json:"shop_sku"`
		IsReroute                   int     `json:"is_reroute"`
		StagePayStatus              string  `json:"stage_pay_status"`
		SkuID                       string  `json:"sku_id"`
		TrackingCodePre             string  `json:"tracking_code_pre"`
		OrderItemID                 int64   `json:"order_item_id"`
		ShopID                      string  `json:"shop_id"`
		OrderFlag                   string  `json:"order_flag"`
		IsFbl                       int     `json:"is_fbl"`
		Name                        string  `json:"name"`
		DeliveryOptionSof           int     `json:"delivery_option_sof"`
		OrderID                     int64   `json:"order_id"`
		Status                      string  `json:"status"`
		ProductMainImage            string  `json:"product_main_image"`
		VoucherPlatform             int     `json:"voucher_platform"`
		PaidPrice                   float64 `json:"paid_price"`
		ProductDetailURL            string  `json:"product_detail_url"`
		WarehouseCode               string  `json:"warehouse_code"`
		PromisedShippingTime        string  `json:"promised_shipping_time"`
		ShippingType                string  `json:"shipping_type"`
		CreatedAt                   string  `json:"created_at"`
		VoucherSellerLpi            int     `json:"voucher_seller_lpi"`
		ShippingFeeDiscountPlatform int     `json:"shipping_fee_discount_platform"`
		WalletCredits               int     `json:"wallet_credits"`
		UpdatedAt                   string  `json:"updated_at"`
		Currency                    string  `json:"currency"`
		ShippingProviderType        string  `json:"shipping_provider_type"`
		VoucherPlatformLpi          int     `json:"voucher_platform_lpi"`
		ShippingFeeOriginal         float64 `json:"shipping_fee_original"`
		ItemPrice                   float64 `json:"item_price"`
		IsDigital                   int     `json:"is_digital"`
		ShippingServiceCost         int     `json:"shipping_service_cost"`
		TrackingCode                string  `json:"tracking_code"`
		ShippingFeeDiscountSeller   float64 `json:"shipping_fee_discount_seller"`
		ShippingAmount              float64 `json:"shipping_amount"`
		ReasonDetail                string  `json:"reason_detail"`
		ReturnStatus                string  `json:"return_status"`
		ShipmentProvider            string  `json:"shipment_provider"`
		VoucherAmount               float64 `json:"voucher_amount"`
		DigitalDeliveryInfo         string  `json:"digital_delivery_info"`
		ExtraAttributes             string  `json:"extra_attributes"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

type SetPackedLazada struct {
	Data struct {
		OrderItems []struct {
			OrderItemID      int64  `json:"order_item_id"`
			TrackingNumber   string `json:"tracking_number"`
			ShipmentProvider string `json:"shipment_provider"`
			PackageID        string `json:"package_id"`
		} `json:"order_items"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}
