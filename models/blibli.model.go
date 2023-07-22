package models

type ListProductBlibli struct {
	RequestId    string                 `json:"requestId" `
	Success      bool                   `json:"success" `
	ErrorMessage string                 `json:"errorMessage" `
	ErrorCode    string                 `json:"errorCode" `
	Content      []ListProductObjBlibli `json:"content" `
	PageMetaData PageDataBlibli         `json:"pageMetaData" `
}

type PageDataBlibli struct {
	PageNumber   int64 `json:"pageNumber" `
	PageSize     int64 `json:"pageSize" `
	TotalRecords int64 `json:"totalRecords" `
}

type ListProductObjBlibli struct {
	Buyable                 bool    `json:"buyable" `
	CategoryCode            string  `json:"categoryCode" `
	CreatedDate             int64   `json:"createdDate" `
	Displayable             bool    `json:"displayable" `
	GdnSku                  string  `json:"gdnSku" `
	Id                      string  `json:"id" `
	Image                   string  `json:"image" `
	IsArchive               bool    `json:"isArchive" `
	MerchantSku             string  `json:"merchantSku" `
	PickupPointCode         string  `json:"pickupPointCode" `
	PickupPointName         string  `json:"pickupPointName" `
	ProductItemCode         string  `json:"productItemCode" `
	ProductName             string  `json:"productName" `
	ProductSku              string  `json:"productSku" `
	ProductType             string  `json:"productType" `
	PromoBundling           bool    `json:"promoBundling" `
	RegularPrice            float64 `json:"regularPrice" `
	Review                  bool    `json:"review" `
	SellingPrice            int64   `json:"sellingPrice" `
	StockAvailableLv1       int64   `json:"stockAvailableLv1" `
	StockAvailableLv2       int64   `json:"stockAvailableLv2" `
	StockReservedLv1        int64   `json:"stockReservedLv1" `
	StockReservedLv2        int64   `json:"stockReservedLv2" `
	StoreId                 int64   `json:"storeId" `
	SynchronizeStock        bool    `json:"synchronizeStock" `
	UpdatedBy               string  `json:"updatedBy" `
	UpdatedDate             int64   `json:"updatedDate" `
	Version                 int64   `json:"version" `
	WholesalePriceActivated bool    `json:"wholesalePriceActivated" `
}

type ListOrdersBlibli struct {
	RequestId    string               `json:"requestId" `
	Success      bool                 `json:"success" `
	ErrorMessage string               `json:"errorMessage" `
	ErrorCode    string               `json:"errorCode" `
	Content      []DetailOrdersBlibli `json:"content" `
	Paging       PagingBlibli         `json:"paging" `
}
type ProductDetailBlibli struct {
	RequestID    string      `json:"requestId"`
	Headers      interface{} `json:"headers"`
	ErrorMessage interface{} `json:"errorMessage"`
	ErrorCode    interface{} `json:"errorCode"`
	Success      bool        `json:"success"`
	Value        struct {
		ID                  string      `json:"id"`
		StoreID             string      `json:"storeId"`
		CreatedDate         int64       `json:"createdDate"`
		CreatedBy           string      `json:"createdBy"`
		UpdatedDate         int64       `json:"updatedDate"`
		UpdatedBy           string      `json:"updatedBy"`
		Version             int         `json:"version"`
		ProductSku          string      `json:"productSku"`
		ProductCode         string      `json:"productCode"`
		BusinessPartnerCode string      `json:"businessPartnerCode"`
		Synchronize         bool        `json:"synchronize"`
		ProductName         string      `json:"productName"`
		ProductType         int         `json:"productType"`
		CategoryCode        string      `json:"categoryCode"`
		CategoryName        string      `json:"categoryName"`
		CategoryHierarchy   string      `json:"categoryHierarchy"`
		Brand               string      `json:"brand"`
		Description         string      `json:"description"`
		SpecificationDetail string      `json:"specificationDetail"`
		UniqueSellingPoint  string      `json:"uniqueSellingPoint"`
		ProductStory        interface{} `json:"productStory"`
		Items               []struct {
			ID                   string      `json:"id"`
			StoreID              string      `json:"storeId"`
			CreatedDate          int64       `json:"createdDate"`
			CreatedBy            string      `json:"createdBy"`
			UpdatedDate          int64       `json:"updatedDate"`
			UpdatedBy            string      `json:"updatedBy"`
			Version              interface{} `json:"version"`
			ItemSku              string      `json:"itemSku"`
			SkuCode              string      `json:"skuCode"`
			MerchantSku          string      `json:"merchantSku"`
			UpcCode              string      `json:"upcCode"`
			ItemName             string      `json:"itemName"`
			Length               float64     `json:"length"`
			Width                float64     `json:"width"`
			Height               float64     `json:"height"`
			Weight               float64     `json:"weight"`
			ShippingWeight       float64     `json:"shippingWeight"`
			DangerousGoodsLevel  int         `json:"dangerousGoodsLevel"`
			LateFulfillment      bool        `json:"lateFulfillment"`
			PickupPointCode      string      `json:"pickupPointCode"`
			PickupPointName      string      `json:"pickupPointName"`
			AvailableStockLevel1 int         `json:"availableStockLevel1"`
			ReservedStockLevel1  int         `json:"reservedStockLevel1"`
			AvailableStockLevel2 int         `json:"availableStockLevel2"`
			ReservedStockLevel2  int         `json:"reservedStockLevel2"`
			MinimumStock         int         `json:"minimumStock"`
			SynchronizeStock     bool        `json:"synchronizeStock"`
			Off2OnActiveFlag     bool        `json:"off2OnActiveFlag"`
			PristineID           interface{} `json:"pristineId"`
			Prices               []struct {
				ID                 interface{} `json:"id"`
				StoreID            interface{} `json:"storeId"`
				CreatedDate        interface{} `json:"createdDate"`
				CreatedBy          interface{} `json:"createdBy"`
				UpdatedDate        interface{} `json:"updatedDate"`
				UpdatedBy          interface{} `json:"updatedBy"`
				Version            interface{} `json:"version"`
				ChannelID          string      `json:"channelId"`
				Price              float64     `json:"price"`
				SalePrice          float64     `json:"salePrice"`
				DiscountAmount     interface{} `json:"discountAmount"`
				DiscountPercentage interface{} `json:"discountPercentage"`
				DiscountStartDate  interface{} `json:"discountStartDate"`
				DiscountEndDate    interface{} `json:"discountEndDate"`
				PromotionName      interface{} `json:"promotionName"`
			} `json:"prices"`
			ViewConfigs []struct {
				ID          interface{} `json:"id"`
				StoreID     interface{} `json:"storeId"`
				CreatedDate interface{} `json:"createdDate"`
				CreatedBy   interface{} `json:"createdBy"`
				UpdatedDate interface{} `json:"updatedDate"`
				UpdatedBy   interface{} `json:"updatedBy"`
				Version     interface{} `json:"version"`
				ChannelID   string      `json:"channelId"`
				Display     bool        `json:"display"`
				Buyable     bool        `json:"buyable"`
			} `json:"viewConfigs"`
			Images []struct {
				ID           interface{} `json:"id"`
				StoreID      interface{} `json:"storeId"`
				CreatedDate  interface{} `json:"createdDate"`
				CreatedBy    interface{} `json:"createdBy"`
				UpdatedDate  interface{} `json:"updatedDate"`
				UpdatedBy    interface{} `json:"updatedBy"`
				Version      interface{} `json:"version"`
				MainImage    bool        `json:"mainImage"`
				Sequence     int         `json:"sequence"`
				LocationPath string      `json:"locationPath"`
			} `json:"images"`
			ProductItemLevel3LogisticsWebResponses interface{}   `json:"productItemLevel3LogisticsWebResponses"`
			Cogs                                   interface{}   `json:"cogs"`
			CogsErrorCode                          string        `json:"cogsErrorCode"`
			PromoBundling                          bool          `json:"promoBundling"`
			MerchantPromoDiscount                  bool          `json:"merchantPromoDiscount"`
			MerchantPromoDiscountActivated         bool          `json:"merchantPromoDiscountActivated"`
			DisableUnSync                          bool          `json:"disableUnSync"`
			PriceEditDisabled                      bool          `json:"priceEditDisabled"`
			WholesalePromoActivated                bool          `json:"wholesalePromoActivated"`
			WholesalePriceActivated                interface{}   `json:"wholesalePriceActivated"`
			ItemCampaignMapped                     bool          `json:"itemCampaignMapped"`
			Archived                               bool          `json:"archived"`
			Rejected                               bool          `json:"rejected"`
			Wholesale                              []interface{} `json:"wholesale"`
		} `json:"items"`
		Attributes []struct {
			ID              interface{} `json:"id"`
			StoreID         interface{} `json:"storeId"`
			CreatedDate     interface{} `json:"createdDate"`
			CreatedBy       interface{} `json:"createdBy"`
			UpdatedDate     interface{} `json:"updatedDate"`
			UpdatedBy       interface{} `json:"updatedBy"`
			Version         interface{} `json:"version"`
			AttributeCode   string      `json:"attributeCode"`
			AttributeType   string      `json:"attributeType"`
			Values          []string    `json:"values"`
			SkuValue        bool        `json:"skuValue"`
			AttributeName   string      `json:"attributeName"`
			ItemSku         interface{} `json:"itemSku"`
			VariantCreation bool        `json:"variantCreation"`
			Mandatory       bool        `json:"mandatory"`
			BasicView       bool        `json:"basicView"`
		} `json:"attributes"`
		Images []struct {
			ID           interface{} `json:"id"`
			StoreID      interface{} `json:"storeId"`
			CreatedDate  interface{} `json:"createdDate"`
			CreatedBy    interface{} `json:"createdBy"`
			UpdatedDate  interface{} `json:"updatedDate"`
			UpdatedBy    interface{} `json:"updatedBy"`
			Version      interface{} `json:"version"`
			MainImage    bool        `json:"mainImage"`
			Sequence     int         `json:"sequence"`
			LocationPath string      `json:"locationPath"`
		} `json:"images"`
		URL                         string `json:"url"`
		InstallationRequired        bool   `json:"installationRequired"`
		CategoryID                  string `json:"categoryId"`
		ForceReview                 bool   `json:"forceReview"`
		EnableEdit                  bool   `json:"enableEdit"`
		WholesalePriceConfigEnabled bool   `json:"wholesalePriceConfigEnabled"`
		ProductScore                struct {
			MandatoryAttributeScore   float64 `json:"mandatoryAttributeScore"`
			ProductTitleScore         float64 `json:"productTitleScore"`
			DescriptionScore          float64 `json:"descriptionScore"`
			UspScore                  float64 `json:"uspScore"`
			RecommendedAttributeScore float64 `json:"recommendedAttributeScore"`
			RemainingAttributeScore   float64 `json:"remainingAttributeScore"`
			VideoURLScore             float64 `json:"videoUrlScore"`
			ImageScore                float64 `json:"imageScore"`
			VariantCreatingScore      float64 `json:"variantCreatingScore"`
			EanUpcScore               float64 `json:"eanUpcScore"`
			TotalScore                float64 `json:"totalScore"`
		} `json:"productScore"`
		ProductEditable bool        `json:"productEditable"`
		Suspended       bool        `json:"suspended"`
		PreOrder        interface{} `json:"preOrder"`
		BlibliURL       struct {
			Product string `json:"product"`
			Variant string `json:"variant"`
		} `json:"blibliUrl"`
		DistinctPickUpPoints interface{} `json:"distinctPickUpPoints"`
	} `json:"value"`
}
type DetailOrdersBlibli struct {
	OrderItems []OrderItemsBlibli `json:"orderItems" `
	PackageId  string             `json:"packageId" `
}

type OrderItemsBlibli struct {
	CashlessStatusUpdateSla          int64               `json:"cashlessStatusUpdateSla" `
	CreatedDate                      int64               `json:"createdDate" `
	InstantPickupAcceptOrderDeadline int64               `json:"instantPickupAcceptOrderDeadline" `
	Logistic                         LogisticOrderBlibli `json:"logistic" `
	Order                            OrderListBlibli     `json:"order" `
	PickupPoint                      PickupPointBlibli   `json:"pickupPoint" `
	PreOrder                         PreOrderBlibli      `json:"preOrder" `
	Product                          ProductOrderBlibli  `json:"product" `
	SellerDeliveryType               string              `json:"sellerDeliveryType" `
	StoreCode                        string              `json:"storeCode" `
}

type ProductOrderBlibli struct {
	BlibliSku   string  `json:"blibliSku" `
	ItemName    string  `json:"itemName" `
	Price       float64 `json:"price" `
	SellerSku   string  `json:"sellerSku" `
	Type        string  `json:"type" `
	VariantName string  `json:"variantName" `
}

type PreOrderBlibli struct {
	Active bool   `json:"active" `
	Type   string `json:"type" `
	Value  int64  `json:"value" `
}

type PickupPointBlibli struct {
	Code string `json:"code" `
	Name string `json:"name" `
}

type OrderListBlibli struct {
	AutoCancelTimestamp      int64  `json:"autoCancelTimestamp" `
	CustomerFullName         string `json:"customerFullName" `
	Date                     int64  `json:"date" `
	Id                       string `json:"id" `
	ItemId                   string `json:"itemId" `
	ItemStatus               string `json:"itemStatus" `
	Quantity                 int64  `json:"quantity" `
	StatusFPUpdatedTimestamp int64  `json:"statusFPUpdatedTimestamp" `
	Type                     string `json:"type" `
}

type LogisticOrderBlibli struct {
	AwbNumber           string `json:"awbNumber" `
	AwbValidityStatus   string `json:"awbValidityStatus" `
	OptionCode          string `json:"optionCode" `
	OptionName          string `json:"optionName" `
	ProductCode         string `json:"productCode" `
	ProductName         string `json:"productName" `
	ShippingInstruction string `json:"shippingInstruction" `
}

type PagingBlibli struct {
	PageNumber  int64 `json:"pageNumber" `
	PageSize    int64 `json:"pageSize" `
	TotalPage   int64 `json:"totalPage" `
	TotalRecord int64 `json:"totalRecord" `
}

type ListDetailOrderBlibli struct {
	RequestId    string                       `json:"requestId" `
	ErrorMessage string                       `json:"errorMessage" `
	ErrorCode    string                       `json:"errorCode" `
	Content      ContentListDetailOrderBlibli `json:"content" `
}

type ContentListDetailOrderBlibli struct {
	Id                 string                   `json:"id" `
	ItemId             string                   `json:"itemId" `
	Status             string                   `json:"status" `
	Old                OldIdBlibli              `json:"old" `
	PackageId          string                   `json:"packageId" `
	PartialFulfillment PartialFulfillmentBlibli `json:"partialFulfillment" `
	PickupPoint        PickupPointBlibliDetail  `json:"pickupPoint" `
	Product            ProductOrderBlibli       `json:"product" `
	Quantity           int64                    `json:"quantity" `
	Recipient          RecipientBlibli          `json:"recipient" `
	Shipment           LogisticOrderBlibli      `json:"shipment" `
}

type LogisticShipmentDetailBlibli struct {
}

type RecipientBlibli struct {
	City               string                   `json:"city" `
	Country            string                   `json:"country" `
	District           string                   `json:"district" `
	Email              string                   `json:"email" `
	Name               string                   `json:"name" `
	PhoneNumber        string                   `json:"phoneNumber" `
	State              string                   `json:"state" `
	StreetAddress      string                   `json:"streetAddress" `
	Subdistrict        string                   `json:"subdistrict" `
	VirtualPhoneNumber VirtualPhoneNumberBlibli `json:"virtualPhoneNumber" `
	ZipCode            string                   `json:"zipCode" `
}

type VirtualPhoneNumberBlibli struct {
	Code string `json:"code" `
	Pin  string `json:"pin" `
}

type PickupPointBlibliDetail struct {
	City          string                           `json:"city" `
	Code          CodePickupPointBlibliDetail      `json:"code" `
	Country       string                           `json:"country" `
	District      string                           `json:"district" `
	Pic           PicPickupPointBlibliDetail       `json:"pic" `
	State         string                           `json:"state" `
	StreetAddress string                           `json:"streetAddress" `
	Subdistrict   string                           `json:"subdistrict" `
	Warehouse     WarehousePickupPointBlibliDetail `json:"warehouse" `
	ZipCode       string                           `json:"zipCode" `
}

type WarehousePickupPointBlibliDetail struct {
	Code string `json:"code" `
}

type PicPickupPointBlibliDetail struct {
	Name  string `json:"name" `
	Phone int64  `json:"phone" `
}

type CodePickupPointBlibliDetail struct {
	Blibli string `json:"blibli" `
	Seller string `json:"seller" `
}

type PartialFulfillmentBlibli struct {
	CanceledQuantity int64  `json:"canceledQuantity" `
	Issuer           string `json:"issuer" `
	Reason           string `json:"reason" `
	RefundSolution   string `json:"refundSolution" `
}

type OldIdBlibli struct {
	Id     string `json:"id" `
	ItemId string `json:"itemId" `
}

type OrderDetailAirwayBillBlibli struct {
	RequestID    string      `json:"requestId"`
	Headers      interface{} `json:"headers"`
	ErrorMessage string      `json:"errorMessage"`
	ErrorCode    string      `json:"errorCode"`
	Success      bool        `json:"success"`
	Value        struct {
		ID                               string      `json:"id"`
		StoreID                          string      `json:"storeId"`
		CreatedDate                      interface{} `json:"createdDate"`
		CreatedBy                        interface{} `json:"createdBy"`
		UpdatedDate                      interface{} `json:"updatedDate"`
		UpdatedBy                        interface{} `json:"updatedBy"`
		Version                          interface{} `json:"version"`
		OrderNo                          string      `json:"orderNo"`
		OrderItemNo                      string      `json:"orderItemNo"`
		OrderType                        string      `json:"orderType"`
		OldOrderNo                       interface{} `json:"oldOrderNo"`
		OldOrderItemNo                   interface{} `json:"oldOrderItemNo"`
		Qty                              int         `json:"qty"`
		OrderDate                        int64       `json:"orderDate"`
		AutoCancelDate                   int64       `json:"autoCancelDate"`
		ProductName                      string      `json:"productName"`
		ProductItemName                  string      `json:"productItemName"`
		ProductPrice                     float64     `json:"productPrice"`
		FinalPrice                       float64     `json:"finalPrice"`
		GdnSku                           string      `json:"gdnSku"`
		GdnItemSku                       string      `json:"gdnItemSku"`
		TotalWeight                      float64     `json:"totalWeight"`
		MerchantSku                      string      `json:"merchantSku"`
		Total                            float64     `json:"total"`
		FinalPriceTotal                  float64     `json:"finalPriceTotal"`
		FastFulfillmentTimestamp         int64       `json:"fastFulfillmentTimestamp"`
		CustName                         string      `json:"custName"`
		OrderStatus                      string      `json:"orderStatus"`
		OrderStatusString                string      `json:"orderStatusString"`
		CustomerAddress                  string      `json:"customerAddress"`
		CustomerEmail                    string      `json:"customerEmail"`
		LogisticsService                 string      `json:"logisticsService"`
		CurrentLogisticService           string      `json:"currentLogisticService"`
		PickupPoint                      string      `json:"pickupPoint"`
		PickupPointName                  string      `json:"pickupPointName"`
		PickupPointAddress               string      `json:"pickupPointAddress"`
		PickupPointCity                  string      `json:"pickupPointCity"`
		PickupPointProvince              string      `json:"pickupPointProvince"`
		PickupPointCountry               string      `json:"pickupPointCountry"`
		PickupPointZipcode               string      `json:"pickupPointZipcode"`
		SellerPickupPointCode            interface{} `json:"sellerPickupPointCode"`
		MerchantDeliveryType             string      `json:"merchantDeliveryType"`
		InstallationRequired             bool        `json:"installationRequired"`
		AwbNumber                        string      `json:"awbNumber"`
		AwbStatus                        string      `json:"awbStatus"`
		ShippingStreetAddress            string      `json:"shippingStreetAddress"`
		ShippingCity                     string      `json:"shippingCity"`
		ShippingSubDistrict              string      `json:"shippingSubDistrict"`
		ShippingDistrict                 string      `json:"shippingDistrict"`
		ShippingProvince                 string      `json:"shippingProvince"`
		ShippingZipCode                  string      `json:"shippingZipCode"`
		ShippingMobile                   string      `json:"shippingMobile"`
		ShippingCost                     float64     `json:"shippingCost"`
		ShippingInsuredAmount            float64     `json:"shippingInsuredAmount"`
		StartOperationalTime             interface{} `json:"startOperationalTime"`
		EndOperationalTime               interface{} `json:"endOperationalTime"`
		Issuer                           interface{} `json:"issuer"`
		RefundResolution                 interface{} `json:"refundResolution"`
		UnFullFillReason                 interface{} `json:"unFullFillReason"`
		UnFullFillQuantity               interface{} `json:"unFullFillQuantity"`
		ProductTypeCode                  string      `json:"productTypeCode"`
		ProductTypeName                  string      `json:"productTypeName"`
		CustNote                         string      `json:"custNote"`
		ShippingRecipientName            string      `json:"shippingRecipientName"`
		LogisticsProductCode             string      `json:"logisticsProductCode"`
		LogisticsProductName             string      `json:"logisticsProductName"`
		LogisticsOptionCode              string      `json:"logisticsOptionCode"`
		OriginLongitude                  float64     `json:"originLongitude"`
		OriginLatitude                   float64     `json:"originLatitude"`
		DestinationLongitude             float64     `json:"destinationLongitude"`
		DestinationLatitude              float64     `json:"destinationLatitude"`
		ItemWeightInKg                   float64     `json:"itemWeightInKg"`
		FulfillmentInfo                  interface{} `json:"fulfillmentInfo"`
		SettlementInfo                   interface{} `json:"settlementInfo"`
		FinanceSettlementInfo            interface{} `json:"financeSettlementInfo"`
		InstantPickup                    bool        `json:"instantPickup"`
		InstantPickupDeadline            interface{} `json:"instantPickupDeadline"`
		InstantPickupAcceptOrderDeadline interface{} `json:"instantPickupAcceptOrderDeadline"`
		SettlementCodeExpired            bool        `json:"settlementCodeExpired"`
		OnlineBookingID                  string      `json:"onlineBookingId"`
		PackageID                        string      `json:"packageId"`
		PackageCreated                   bool        `json:"packageCreated"`
		LogisticsOptionName              string      `json:"logisticsOptionName"`
		CashlessHandover                 bool        `json:"cashlessHandover"`
		CashlessStatusUpdateSLA          int         `json:"cashlessStatusUpdateSla"`
		OrderHistory                     []struct {
			ID               string      `json:"id"`
			StoreID          string      `json:"storeId"`
			CreatedDate      int64       `json:"createdDate"`
			CreatedBy        string      `json:"createdBy"`
			UpdatedDate      int64       `json:"updatedDate"`
			UpdatedBy        string      `json:"updatedBy"`
			Version          interface{} `json:"version"`
			OrderStatus      string      `json:"orderStatus"`
			OrderStatusDesc  string      `json:"orderStatusDesc"`
			CreatedTimestamp int64       `json:"createdTimestamp"`
		} `json:"orderHistory"`
		ManifestInfo          interface{} `json:"manifestInfo"`
		Manifest              interface{} `json:"manifest"`
		Combo                 interface{} `json:"combo"`
		Wholesale             interface{} `json:"wholesale"`
		HasVirtualPhoneNumber bool        `json:"hasVirtualPhoneNumber"`
		IsShipBySeller        bool        `json:"isShipBySeller"`
		Flags                 struct {
			SupportingDocumentExist bool `json:"supportingDocumentExist"`
			FulfilledByBlibli       bool `json:"fulfilledByBlibli"`
		} `json:"flags"`
		StoreCode                         string      `json:"storeCode"`
		PreOrder                          interface{} `json:"preOrder"`
		InstantPickupScheduleDateResponse interface{} `json:"instantPickupScheduleDateResponse"`
		IsCashOnDelivery                  bool        `json:"isCashOnDelivery"`
	} `json:"value"`
}

type PackageHeadBlibli struct {
	RequestId    string                `json:"requestId" `
	Success      bool                  `json:"success" `
	ErrorMessage string                `json:"errorMessage" `
	ErrorCode    string                `json:"errorCode" `
	Value        PackageListHeadBlibli `json:"value" `
}

type PackageListHeadBlibli struct {
	PackageId string `json:"packageId" `
}
