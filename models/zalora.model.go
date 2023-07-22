package models

import "time"

//"github.com/google/uuid"
type UpdateStockZaloraV2 []struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type DetailStockZaloraV2 []struct {
	ProductID            int           `json:"productId"`
	SellerSku            string        `json:"sellerSku"`
	ShopSku              string        `json:"shopSku"`
	Name                 string        `json:"name"`
	Quantity             int           `json:"quantity"`
	ReservedStock        int           `json:"reservedStock"`
	PreVerificationStock int           `json:"preVerificationStock"`
	Available            int           `json:"available"`
	Consignments         interface{}   `json:"consignments"`
	Warehouses           []interface{} `json:"warehouses"`
}

type ZaloraPacked struct {
	OrderItemIds []int `json:"orderItemIds"`
}

type ErrorZaloraV2 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorZaloraV3 struct {
	Errors []struct {
		OrderItemID int    `json:"orderItemId"`
		Detail      string `json:"detail"`
	} `json:"errors"`
}

type ProductV2Zalora struct {
	Items []struct {
		ID                int           `json:"id"`
		UUID              string        `json:"uuid"`
		SrcID             string        `json:"srcId"`
		Name              string        `json:"name"`
		ParentSku         string        `json:"parentSku"`
		CreatedAt         time.Time     `json:"createdAt"`
		UpdatedAt         time.Time     `json:"updatedAt"`
		Description       string        `json:"description"`
		BrandID           int           `json:"brandId"`
		PrimaryCategoryID int           `json:"primaryCategoryId"`
		AttributeSetID    int           `json:"attributeSetId"`
		SellerID          int           `json:"sellerId"`
		Categories        []interface{} `json:"categories"`
		Attributes        struct {
			Num11  string `json:"11"`
			Num12  string `json:"12"`
			Num52  string `json:"52"`
			Num53  string `json:"53"`
			Num54  string `json:"54"`
			Num55  string `json:"55"`
			Num59  string `json:"59"`
			Num149 string `json:"149"`
			Num163 string `json:"163"`
			Num201 string `json:"201"`
			Num234 string `json:"234"`
			Num268 string `json:"268"`
		} `json:"attributes"`
		SizeSystem  int           `json:"sizeSystem"`
		BrowseNodes []interface{} `json:"browseNodes"`
		SellerSku   string        `json:"sellerSku"`
	} `json:"items"`
	Pagination struct {
		Limit      int `json:"limit"`
		Offset     int `json:"offset"`
		TotalCount int `json:"totalCount"`
	} `json:"pagination"`
}

type OrdersZalora struct {
	Items []struct {
		UUID            string `json:"uuid"`
		InvoiceRequired bool   `json:"invoiceRequired"`
		ID              int    `json:"id"`
		SellerID        int    `json:"sellerId"`
		Number          string `json:"number"`
		Customer        struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
		} `json:"customer"`
		Address struct {
			Billing struct {
				FirstName string   `json:"firstName"`
				LastName  string   `json:"lastName"`
				Phone     []string `json:"phone"`
				Address   []string `json:"address"`
				Email     string   `json:"email"`
				City      string   `json:"city"`
				Ward      string   `json:"ward"`
				Region    string   `json:"region"`
				PostCode  string   `json:"postCode"`
				Country   string   `json:"country"`
			} `json:"billing"`
			Shipping struct {
				FirstName string   `json:"firstName"`
				LastName  string   `json:"lastName"`
				Phone     []string `json:"phone"`
				Address   []string `json:"address"`
				Email     string   `json:"email"`
				City      string   `json:"city"`
				Ward      string   `json:"ward"`
				Region    string   `json:"region"`
				PostCode  string   `json:"postCode"`
				Country   string   `json:"country"`
			} `json:"shipping"`
		} `json:"address"`
		NationalRegistrationNumber interface{} `json:"nationalRegistrationNumber"`
		PayoutPending              bool        `json:"payoutPending"`
		Gift                       struct {
			Option  bool   `json:"option"`
			Message string `json:"message"`
		} `json:"gift"`
		Voucher struct {
			Code string `json:"code"`
			Type string `json:"type"`
		} `json:"voucher"`
		DeliveryInfo       string      `json:"deliveryInfo"`
		PaymentMethod      string      `json:"paymentMethod"`
		Currency           string      `json:"currency"`
		Remarks            string      `json:"remarks"`
		CreatedAt          time.Time   `json:"createdAt"`
		ImportedAt         time.Time   `json:"importedAt"`
		UpdatedAt          time.Time   `json:"updatedAt"`
		AddressUpdatedAt   time.Time   `json:"addressUpdatedAt"`
		ExchangeByOrderID  interface{} `json:"exchangeByOrderId"`
		ExchangeForOrderID interface{} `json:"exchangeForOrderId"`
		Source             string      `json:"source"`
		ExtraAttributes    interface{} `json:"extraAttributes"`
		StatusList         struct {
			Pending int `json:"pending"`
		} `json:"statusList"`
		ItemCount                   int    `json:"itemCount"`
		UnitPriceSumWithFees        int    `json:"unitPriceSumWithFees"`
		ShipmentProviderType        string `json:"shipmentProviderType"`
		ShipmentProviderPreSelected bool   `json:"shipmentProviderPreSelected"`
		TargetToShip                string `json:"targetToShip"`
		PackedItemsCount            int    `json:"packedItemsCount"`
		OrderItemIds                []int  `json:"orderItemIds"`
		Items                       []struct {
			ID            int    `json:"id"`
			SrcID         string `json:"srcId"`
			SellerID      int    `json:"sellerId"`
			OrderID       int    `json:"orderId"`
			UUID          string `json:"uuid"`
			Status        string `json:"status"`
			IsProcessable bool   `json:"isProcessable"`
			FailureReason struct {
				Type        interface{} `json:"type"`
				Name        interface{} `json:"name"`
				Description interface{} `json:"description"`
				Details     string      `json:"details"`
			} `json:"failureReason"`
			Shipment struct {
				Type                     string      `json:"type"`
				CrossdockingDeliveryType interface{} `json:"crossdockingDeliveryType"`
				Method                   string      `json:"method"`
				PreProvider              interface{} `json:"preProvider"`
				Provider                 struct {
					UUID        string      `json:"uuid"`
					Name        string      `json:"name"`
					IsDefault   bool        `json:"isDefault"`
					DigitalType interface{} `json:"digitalType"`
				} `json:"provider"`
				ProviderPreselected bool        `json:"providerPreselected"`
				ProviderProduct     string      `json:"providerProduct"`
				ProviderType        string      `json:"providerType"`
				Weight              int         `json:"weight"`
				TrackingCode        string      `json:"trackingCode"`
				PreTrackingCode     interface{} `json:"preTrackingCode"`
			} `json:"shipment"`
			InvoiceNumber    interface{} `json:"invoiceNumber"`
			InvoiceAccesskey interface{} `json:"invoiceAccesskey"`
			InTransit        bool        `json:"inTransit"`
			Premium          bool        `json:"premium"`
			TargetToShipAt   interface{} `json:"targetToShipAt"`
			Product          struct {
				Name      string `json:"name"`
				Sku       string `json:"sku"`
				Variation string `json:"variation"`
				SellerSku string `json:"sellerSku"`
			} `json:"product"`
			UnitPrice             int         `json:"unitPrice"`
			TaxAmount             int         `json:"taxAmount"`
			TaxPercent            int         `json:"taxPercent"`
			PaidPrice             int         `json:"paidPrice"`
			PaidCommission        interface{} `json:"paidCommission"`
			ShippingFee           int         `json:"shippingFee"`
			ShippingServiceCost   interface{} `json:"shippingServiceCost"`
			WalletCredits         int         `json:"walletCredits"`
			StoreCredits          int         `json:"storeCredits"`
			ShippingVoucherAmount int         `json:"shippingVoucherAmount"`
			PriceAfterDiscount    int         `json:"priceAfterDiscount"`
			SalesDueAmount        int         `json:"salesDueAmount"`
			ItemSerialNumber      interface{} `json:"itemSerialNumber"`
			AbatementRate         interface{} `json:"abatementRate"`
			ExciseRate            interface{} `json:"exciseRate"`
			HsnCode               interface{} `json:"hsnCode"`
			CodCollectableAmount  interface{} `json:"codCollectableAmount"`
			Purchase              struct {
				OrderSrcID    interface{} `json:"orderSrcId"`
				OrderNumber   interface{} `json:"orderNumber"`
				InvoiceNumber interface{} `json:"invoiceNumber"`
			} `json:"purchase"`
			CreatedAt     time.Time     `json:"createdAt"`
			UpdatedAt     time.Time     `json:"updatedAt"`
			WarehouseName interface{}   `json:"warehouseName"`
			IsHybrid      bool          `json:"isHybrid"`
			IsOutlet      bool          `json:"isOutlet"`
			Actions       []string      `json:"actions"`
			Vouchers      []interface{} `json:"vouchers"`
		} `json:"items"`
		RegionID string `json:"regionId"`
	} `json:"items"`
	Pagination struct {
		Limit      int `json:"limit"`
		Offset     int `json:"offset"`
		TotalCount int `json:"totalCount"`
	} `json:"pagination"`
}

type DocumentZaloraV2 struct {
	ID               int       `json:"id"`
	UserID           int       `json:"userId"`
	SellerID         int       `json:"sellerId"`
	Status           string    `json:"status"`
	ExportAction     string    `json:"exportAction"`
	IsDeleted        bool      `json:"isDeleted"`
	CreatedAt        string    `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	ExpiresAt        string    `json:"expiresAt"`
	RequestedFormats []string  `json:"requestedFormats"`
	DownloadLinks    []string  `json:"downloadLinks"`
	ExportContent    string    `json:"exportContent"`
}

type OrderDetailZaolraV2 struct {
	UUID            string `json:"uuid"`
	InvoiceRequired bool   `json:"invoiceRequired"`
	ID              int    `json:"id"`
	SellerID        int    `json:"sellerId"`
	Number          string `json:"number"`
	Customer        struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"customer"`
	Address struct {
		Billing struct {
			FirstName string   `json:"firstName"`
			LastName  string   `json:"lastName"`
			Phone     []string `json:"phone"`
			Address   []string `json:"address"`
			Email     string   `json:"email"`
			City      string   `json:"city"`
			Ward      string   `json:"ward"`
			Region    string   `json:"region"`
			PostCode  string   `json:"postCode"`
			Country   string   `json:"country"`
		} `json:"billing"`
		Shipping struct {
			FirstName string   `json:"firstName"`
			LastName  string   `json:"lastName"`
			Phone     []string `json:"phone"`
			Address   []string `json:"address"`
			Email     string   `json:"email"`
			City      string   `json:"city"`
			Ward      string   `json:"ward"`
			Region    string   `json:"region"`
			PostCode  string   `json:"postCode"`
			Country   string   `json:"country"`
		} `json:"shipping"`
	} `json:"address"`
	NationalRegistrationNumber interface{} `json:"nationalRegistrationNumber"`
	PayoutPending              bool        `json:"payoutPending"`
	Gift                       struct {
		Option  bool   `json:"option"`
		Message string `json:"message"`
	} `json:"gift"`
	Voucher struct {
		Code string `json:"code"`
		Type string `json:"type"`
	} `json:"voucher"`
	DeliveryInfo       string      `json:"deliveryInfo"`
	PaymentMethod      string      `json:"paymentMethod"`
	Currency           string      `json:"currency"`
	Remarks            string      `json:"remarks"`
	CreatedAt          string      `json:"createdAt"`
	ImportedAt         time.Time   `json:"importedAt"`
	UpdatedAt          time.Time   `json:"updatedAt"`
	AddressUpdatedAt   time.Time   `json:"addressUpdatedAt"`
	ExchangeByOrderID  interface{} `json:"exchangeByOrderId"`
	ExchangeForOrderID interface{} `json:"exchangeForOrderId"`
	Source             string      `json:"source"`
	ExtraAttributes    interface{} `json:"extraAttributes"`
	StatusList         struct {
		Pending int `json:"pending"`
	} `json:"statusList"`
	ItemCount                   int    `json:"itemCount"`
	UnitPriceSumWithFees        int    `json:"unitPriceSumWithFees"`
	ShipmentProviderType        string `json:"shipmentProviderType"`
	ShipmentProviderPreSelected bool   `json:"shipmentProviderPreSelected"`
	TargetToShip                string `json:"targetToShip"`
	PackedItemsCount            int    `json:"packedItemsCount"`
	OrderItemIds                []int  `json:"orderItemIds"`
	Items                       []struct {
		ID            int    `json:"id"`
		SrcID         string `json:"srcId"`
		SellerID      int    `json:"sellerId"`
		OrderID       int    `json:"orderId"`
		UUID          string `json:"uuid"`
		Status        string `json:"status"`
		IsProcessable bool   `json:"isProcessable"`
		FailureReason struct {
			Type        interface{} `json:"type"`
			Name        interface{} `json:"name"`
			Description interface{} `json:"description"`
			Details     string      `json:"details"`
		} `json:"failureReason"`
		Shipment struct {
			Type                     string      `json:"type"`
			CrossdockingDeliveryType interface{} `json:"crossdockingDeliveryType"`
			Method                   string      `json:"method"`
			PreProvider              interface{} `json:"preProvider"`
			Provider                 struct {
				UUID        string      `json:"uuid"`
				Name        string      `json:"name"`
				IsDefault   bool        `json:"isDefault"`
				DigitalType interface{} `json:"digitalType"`
			} `json:"provider"`
			ProviderPreselected bool        `json:"providerPreselected"`
			ProviderProduct     string      `json:"providerProduct"`
			ProviderType        string      `json:"providerType"`
			Weight              int         `json:"weight"`
			TrackingCode        string      `json:"trackingCode"`
			PreTrackingCode     interface{} `json:"preTrackingCode"`
		} `json:"shipment"`
		InvoiceNumber    interface{} `json:"invoiceNumber"`
		InvoiceAccesskey interface{} `json:"invoiceAccesskey"`
		InTransit        bool        `json:"inTransit"`
		Premium          bool        `json:"premium"`
		TargetToShipAt   interface{} `json:"targetToShipAt"`
		Product          struct {
			Name      string `json:"name"`
			Sku       string `json:"sku"`
			Variation string `json:"variation"`
			SellerSku string `json:"sellerSku"`
		} `json:"product"`
		UnitPrice int       `json:"unitPrice"`
		PaidPrice int       `json:"paidPrice"`
		CreatedAt string    `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"items"`
	RegionID string `json:"regionId"`
}

type TokenZalora struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type Test struct {
	Detail       string       `json:"detail"`
	DetailObject []TestObject `json:"detail_object"`
}

type TestObject struct {
	DetailObjects string `json:"detail_objects"`
	DetailObjectx string `json:"detail_objectx"`
}

type ErrorResponseHeader struct {
	ErrorResponse ErrorResponseHeaderList `json:"ErrorResponse"`
}
type ErrorResponseHeaderList struct {
	Head HeadListError `json:"Head"`
}
type HeadListError struct {
	RequestAction string `json:"RequestAction"`
	ErrorType     string `json:"ErrorType"`
	ErrorCode     string `json:"ErrorCode"`
	ErrorMessage  string `json:"ErrorMessage"`
}

type SuccessResponseHeader struct {
	SuccessResponse SuccessResponseList `json:"SuccessResponse"`
}

type SuccessResponseList struct {
	Body BodyList `json:"Body"`
	Head HeadList `json:"Head"`
}
type HeadList struct {
	RequestId     string `json:"RequestId"`
	RequestAction string `json:"RequestAction"`
	ResponseType  string `json:"ResponseType"`
	Timestamp     string `json:"Timestamp"`
	TotalCount    string `json:"TotalCount"`
}

type BodyList struct {
	Orders []OrdersListZalora `json:"Orders"`
}

type OrdersListZalora struct {
	Order OrderListZalora `json:"Order"`
}
type OrderListZalora struct {
	OrderId                    string                `json:"OrderId"`
	CustomerFirstName          string                `json:"CustomerFirstName"`
	CustomerLastName           string                `json:"CustomerLastName"`
	OrderNumber                string                `json:"OrderNumber"`
	PaymentMethod              string                `json:"PaymentMethod"`
	Currency                   string                `json:"Currency"`
	Remarks                    string                `json:"Remarks"`
	DeliveryInfo               string                `json:"DeliveryInfo"`
	Price                      string                `json:"Price"`
	GiftOption                 string                `json:"GiftOption"`
	GiftMessage                string                `json:"GiftMessage"`
	VoucherCode                string                `json:"VoucherCode"`
	CreatedAt                  string                `json:"CreatedAt"`
	UpdatedAt                  string                `json:"UpdatedAt"`
	AddressUpdatedAt           string                `json:"AddressUpdatedAt"`
	AddressBilling             AddressBillingZalora  `json:"AddressBilling"`
	AddressShipping            AddressShippingZalora `json:"AddressShipping"`
	NationalRegistrationNumber string                `json:"NationalRegistrationNumber"`
	ItemsCount                 string                `json:"ItemsCount"`
	PromisedShippingTime       string                `json:"PromisedShippingTime"`
	ExtraAttributes            string                `json:"ExtraAttributes"`
	InvoiceRequired            string                `json:"InvoiceRequired"`
	ExchangeForOrderId         string                `json:"ExchangeForOrderId"`
	ExchangeByOrderId          string                `json:"ExchangeByOrderId"`
	Statuses                   []StatusesZalora      `json:"Statuses"`
}
type StatusesZalora struct {
	Status string `json:"Status"`
}

type AddressShippingZalora struct {
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	Phone         string `json:"Phone"`
	Phone2        string `json:"Phone2"`
	Address1      string `json:"Address1"`
	Address2      string `json:"Address2"`
	Address3      string `json:"Address3"`
	Address4      string `json:"Address4"`
	Address5      string `json:"Address5"`
	CustomerEmail string `json:"CustomerEmail"`
	City          string `json:"City"`
	Ward          string `json:"Ward"`
	Region        string `json:"Region"`
	PostCode      string `json:"PostCode"`
	Country       string `json:"Country"`
}

type AddressBillingZalora struct {
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	Phone         string `json:"Phone"`
	Phone2        string `json:"Phone2"`
	Address1      string `json:"Address1"`
	Address2      string `json:"Address2"`
	Address3      string `json:"Address3"`
	Address4      string `json:"Address4"`
	Address5      string `json:"Address5"`
	CustomerEmail string `json:"CustomerEmail"`
	City          string `json:"City"`
	Ward          string `json:"Ward"`
	Region        string `json:"Region"`
	PostCode      string `json:"PostCode"`
	Country       string `json:"Country"`
}

//ORDERS SINGLE
type SuccessResponseHeaderOrderSingle struct {
	SuccessResponse SuccessResponseListOrderSingle `json:"SuccessResponse"`
}

type SuccessResponseListOrderSingle struct {
	Body BodyListOrderSingle `json:"Body"`
	Head HeadList            `json:"Head"`
}

type BodyListOrderSingle struct {
	Orders OrdersListZaloraOrderSingle `json:"Orders"`
}

type OrdersListZaloraOrderSingle struct {
	Order OrderListZalora `json:"Order"`
}

//ORDER DETAIL
type DetailOrderZalora struct {
	SuccessResponse DetailOrderZaloraList `json:"SuccessResponse"`
}
type DetailOrderZaloraList struct {
	Body BodyListOrderDetail `json:"Body"`
	Head HeadList            `json:"Head"`
}

type BodyListOrderDetail struct {
	Orders OrdersListDetailZalora `json:"Orders"`
}

type OrdersListDetailZalora struct {
	Order OrderListZalora `json:"Order"`
	//Order OrderListDetailZalora `json:"Order"`
}

// type OrderListDetailZalora struct {
// 	OrderId           string `json:"OrderId"`
// 	CustomerFirstName string `json:"CustomerFirstName"`
// 	CustomerLastName  string `json:"CustomerLastName"`
// 	OrderNumber       string `json:"OrderNumber"`
// }

//END ORDER DETAIL

type ObjSkuZalora struct {
	Sku              string  `json:"sku"`
	Name             string  `json:"name"`
	Variation        string  `json:"variation"`
	ShipmentProvider string  `json:"shipment_provider"`
	CreatedAt        string  `json:"created_at"`
	Qty              int     `json:"qty"`
	OrderId          string  `json:"order_id"`
	Status           string  `json:"status"`
	Amount           float64 `json:"amount"`
}

type SuccessResponseItemZalora struct {
	SuccessResponse SuccessResponseItem `json:"SuccessResponse"`
}
type SuccessResponseItem struct {
	Body BodyListItemZalora `json:"Body"`
	Head HeadListItemZalora `json:"Head"`
}
type HeadListItemZalora struct {
	RequestId     string `json:"RequestId"`
	RequestAction string `json:"RequestAction"`
	ResponseType  string `json:"ResponseType"`
	Timestamp     string `json:"Timestamp"`
}

type BodyListItemZalora struct {
	OrderItems OrdersItemsZalora `json:"OrderItems"`
}
type OrdersItemsZalora struct {
	OrderItem []OrdersItemZalora `json:"OrderItem"`
}
type OrdersItemZalora struct {
	OrderItemId          string                `json:"OrderItemId"`
	ShopId               string                `json:"ShopId"`
	OrderId              string                `json:"OrderId"`
	Name                 string                `json:"Name"`
	Sku                  string                `json:"Sku"`
	Variation            string                `json:"Variation"`
	ShopSku              string                `json:"ShopSku"`
	ShippingType         string                `json:"ShippingType"`
	ItemPrice            string                `json:"ItemPrice"`
	PaidPrice            string                `json:"PaidPrice"`
	Currency             string                `json:"Currency"`
	WalletCredits        string                `json:"WalletCredits"`
	TaxAmount            string                `json:"TaxAmount"`
	CodCollectableAmount string                `json:"CodCollectableAmount"`
	ShippingAmount       string                `json:"ShippingAmount"`
	ShippingServiceCost  string                `json:"ShippingServiceCost"`
	VoucherAmount        string                `json:"VoucherAmount"`
	VoucherCode          string                `json:"VoucherCode"`
	Status               string                `json:"Status"`
	IsProcessable        string                `json:"IsProcessable"`
	ShipmentProvider     string                `json:"ShipmentProvider"`
	IsDigital            string                `json:"IsDigital"`
	DigitalDeliveryInfo  string                `json:"DigitalDeliveryInfo"`
	TrackingCode         string                `json:"TrackingCode"`
	TrackingCodePre      string                `json:"TrackingCodePre"`
	Reason               string                `json:"Reason"`
	ReasonDetail         string                `json:"ReasonDetail"`
	PurchaseOrderId      string                `json:"PurchaseOrderId"`
	PurchaseOrderNumber  string                `json:"PurchaseOrderNumber"`
	PackageId            string                `json:"PackageId"`
	PromisedShippingTime string                `json:"PromisedShippingTime"`
	ExtraAttributes      ExtraAttributesZalora `json:"ExtraAttributes"`
	ShippingProviderType string                `json:"ShippingProviderType"`
	CreatedAt            string                `json:"CreatedAt"`
	UpdatedAt            string                `json:"UpdatedAt"`
	Vouchers             string                `json:"Vouchers"`
	ShippingVoucher      string                `json:"ShippingVoucher"`
	WarehouseName        string                `json:"WarehouseName"`
	StoreCredits         string                `json:"StoreCredits"`
	ExchangeForOrderId   string                `json:"ExchangeForOrderId"`
	ExchangeByOrderId    string                `json:"ExchangeByOrderId"`
	IsHybrid             string                `json:"IsHybrid"`
	IsOutlet             string                `json:"IsOutlet"`
	ReturnStatus         string                `json:"ReturnStatus"`
}

//tambahan jika order 1 sku 1 qty
type SuccessResponseItemZaloras struct {
	SuccessResponse SuccessResponseItems `json:"SuccessResponse"`
}
type SuccessResponseItems struct {
	Body BodyListItemZaloras `json:"Body"`
	Head HeadListItemZalora  `json:"Head"`
}

type BodyListItemZaloras struct {
	OrderItems OrdersItemsZaloras `json:"OrderItems"`
}
type OrdersItemsZaloras struct {
	OrderItem OrdersItemZaloras `json:"OrderItem"`
}
type OrdersItemZaloras struct {
	OrderItemId          string                `json:"OrderItemId"`
	ShopId               string                `json:"ShopId"`
	OrderId              string                `json:"OrderId"`
	Name                 string                `json:"Name"`
	Sku                  string                `json:"Sku"`
	Variation            string                `json:"Variation"`
	ShopSku              string                `json:"ShopSku"`
	ShippingType         string                `json:"ShippingType"`
	ItemPrice            string                `json:"ItemPrice"`
	PaidPrice            string                `json:"PaidPrice"`
	Currency             string                `json:"Currency"`
	WalletCredits        string                `json:"WalletCredits"`
	TaxAmount            string                `json:"TaxAmount"`
	CodCollectableAmount string                `json:"CodCollectableAmount"`
	ShippingAmount       string                `json:"ShippingAmount"`
	ShippingServiceCost  string                `json:"ShippingServiceCost"`
	VoucherAmount        string                `json:"VoucherAmount"`
	VoucherCode          string                `json:"VoucherCode"`
	Status               string                `json:"Status"`
	IsProcessable        string                `json:"IsProcessable"`
	ShipmentProvider     string                `json:"ShipmentProvider"`
	IsDigital            string                `json:"IsDigital"`
	DigitalDeliveryInfo  string                `json:"DigitalDeliveryInfo"`
	TrackingCode         string                `json:"TrackingCode"`
	TrackingCodePre      string                `json:"TrackingCodePre"`
	Reason               string                `json:"Reason"`
	ReasonDetail         string                `json:"ReasonDetail"`
	PurchaseOrderId      string                `json:"PurchaseOrderId"`
	PurchaseOrderNumber  string                `json:"PurchaseOrderNumber"`
	PackageId            string                `json:"PackageId"`
	PromisedShippingTime string                `json:"PromisedShippingTime"`
	ExtraAttributes      ExtraAttributesZalora `json:"ExtraAttributes"`
	ShippingProviderType string                `json:"ShippingProviderType"`
	CreatedAt            string                `json:"CreatedAt"`
	UpdatedAt            string                `json:"UpdatedAt"`
	Vouchers             string                `json:"Vouchers"`
	ShippingVoucher      string                `json:"ShippingVoucher"`
	WarehouseName        string                `json:"WarehouseName"`
	StoreCredits         string                `json:"StoreCredits"`
	ExchangeForOrderId   string                `json:"ExchangeForOrderId"`
	ExchangeByOrderId    string                `json:"ExchangeByOrderId"`
	IsHybrid             string                `json:"IsHybrid"`
	IsOutlet             string                `json:"IsOutlet"`
	ReturnStatus         string                `json:"ReturnStatus"`
}

//end tambahan jika order 1 sku 1 qty

type SuccessResponseShipmentProvidesZalora struct {
	SuccessResponse SuccessResponseShipmentProvides `json:"SuccessResponse"`
}
type SuccessResponseShipmentProvides struct {
	Body BodyListShipmentProvideZaloras `json:"Body"`
	Head HeadListItemZalora             `json:"Head"`
}

type BodyListShipmentProvideZaloras struct {
	ShipmentProviders ShipmentProvider `json:"ShipmentProviders"`
}
type ShipmentProvider struct {
	ShipmentProvider []ShipmentProviderList `json:"ShipmentProvider"`
}
type ShipmentProviderList struct {
	Name                        string `json:"name"`
	Default                     string `json:"default"`
	ApiIntegration              string `json:"apiintegration"`
	Cod                         string `json:"cod"`
	TrackingCodeValidationRegex string `json:"trackingcodevalidationregex"`
	TrackingCodeExample         string `json:"trackingcodeexample"`
	TrackingUrl                 string `json:"trackingurl"`
	TrackingCodeSetOnStep       string `json:"trackingcodesetonstep"`
	EnabledDeliveryOptions      string `json:"enableddeliveryoptions"`
}

type ExtraAttributesZalora struct {
	SalesServiceTax  string `json:"sales_service_tax"`
	PriceExcludedTax string `json:"price_excluded_tax"`
}

//PRODUCT ZALORA
type SuccessResponseProductZalora struct {
	SuccessResponse SuccessResponseProduct `json:"SuccessResponse"`
}
type SuccessResponseProduct struct {
	Body BodyListProductZalora `json:"Body"`
	Head HeadList              `json:"Head"`
}

type BodyListProductZalora struct {
	Products ProductsListZalora `json:"Products"`
}

type ProductsListZalora struct {
	Product []ProductListZalora `json:"Product"`
}

type ProductListZalora struct {
	SellerSku string `json:"SellerSku"`
	ShopSku   string `json:"ShopSku"`
	Name      string `json:"Name"`
	Variation string `json:"Variation"`
	ParentSku string `json:"ParentSku"`
	Status    string `json:"Status"`
}

//END PRODUCT ZALORA

//PRODUCT STOCK ZALORA
type SuccessResponseProductStockZalora struct {
	SuccessResponse SuccessResponseProductStock `json:"SuccessResponse"`
}
type SuccessResponseProductStock struct {
	Body BodyListProductStockZalora `json:"Body"`
	Head HeadList                   `json:"Head"`
}
type BodyListProductStockZalora struct {
	ProductStocks ProductsStockListZalora `json:"ProductStocks"`
}
type ProductsStockListZalora struct {
	ProductStock ProductListStockZalora `json:"ProductStock"`
}
type ProductListStockZalora struct {
	SellerSku            string `json:"SellerSku"`
	ShopSku              string `json:"ShopSku"`
	Name                 string `json:"Name"`
	Quantity             string `json:"Quantity"`
	ReservedStock        string `json:"ReservedStock"`
	PreVerificationStock string `json:"PreVerificationStock"`
	Available            string `json:"Available"`
	Consignments         string `json:"Consignments"`
	Warehouses           string `json:"Warehouses"`
}

//END PRODUCT STOCK ZALORA

//PICKUP ZALORA
type SuccessResponsePickZalora struct {
	SuccessResponse SuccessPickZalora `json:"SuccessResponse"`
}

type SuccessPickZalora struct {
	Body BodyListPickZalora `json:"Body"`
	Head HeadList           `json:"Head"`
}

type BodyListPickZalora struct {
	OrderItems OrderItemsPickZalora `json:"OrderItems"`
}

type OrderItemsPickZalora struct {
	OrderItem []OrderItemPickZalora `json:"OrderItem"`
}

type OrderItemPickZalora struct {
	PurchaseOrderId     string `json:"PurchaseOrderId"`
	PurchaseOrderNumber string `json:"PurchaseOrderNumber"`
	TrackingCode        string `json:"TrackingCode"`
}

//END PICKUP ZALORA

//SKU BATCH
type UpdateStockBatchZalora struct {
	SkuList []SkuListZalora `json:"sku_list"`
}

type SkuListZalora struct {
	Sku   string `json:"sku"`
	Stock int16  `json:"stock"`
}
type DocumentShipZalora struct {
	SuccessResponse struct {
		Head struct {
			RequestID     string `json:"RequestId"`
			RequestAction string `json:"RequestAction"`
			ResponseType  string `json:"ResponseType"`
			Timestamp     string `json:"Timestamp"`
		} `json:"Head"`
		Body struct {
			Documents struct {
				Document struct {
					DocumentType string `json:"DocumentType"`
					MimeType     string `json:"MimeType"`
					File         string `json:"File"`
				} `json:"Document"`
			} `json:"Documents"`
		} `json:"Body"`
	} `json:"SuccessResponse"`
}
