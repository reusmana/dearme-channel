package models

import "time"

type ErrorTokped struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	EventCode        string `json:"event_code"`
	LastLoginType    string `json:"last_login_type"`
}

type TokenTokped struct {
	AccessToken   string `json:"access_token"`
	EventCode     string `json:"event_code"`
	ExpiresIn     int    `json:"expires_in"`
	LastLoginType string `json:"last_login_type"`
	SqCheck       bool   `json:"sq_check"`
	TokenType     string `json:"token_type"`
}

type ProductTokped2 struct {
	Data []struct {
		ProductID  int64  `json:"product_id"`
		Name       string `json:"name"`
		Sku        string `json:"sku"`
		ShopID     int    `json:"shop_id"`
		ShopName   string `json:"shop_name"`
		CategoryID int    `json:"category_id"`
		Desc       string `json:"desc"`
		Stock      int    `json:"stock"`
		Price      int    `json:"price"`
		Status     string `json:"status"`
	} `json:"data"`
	Status       string        `json:"status"`
	ErrorMessage []interface{} `json:"error_message"`
}

type ProductTokped struct {
	Header struct {
		ProcessTime int    `json:"process_time"`
		Messages    string `json:"messages"`
		Reason      string `json:"reason"`
		ErrorCode   string `json:"error_code"`
	} `json:"header"`
	Data struct {
		TotalData int `json:"total_data"`
		Shop      struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			URI      string `json:"uri"`
			Location string `json:"location"`
		} `json:"shop"`
		Products []struct {
			ID          int     `json:"id"`
			Name        string  `json:"name"`
			Childs      []int64 `json:"childs"`
			URL         string  `json:"url"`
			ImageURL    string  `json:"image_url"`
			ImageURL700 string  `json:"image_url_700"`
			Price       string  `json:"price"`
			Shop        struct {
				ID           int    `json:"id"`
				Name         string `json:"name"`
				URL          string `json:"url"`
				IsGold       bool   `json:"is_gold"`
				Location     string `json:"location"`
				City         string `json:"city"`
				Reputation   string `json:"reputation"`
				Clover       string `json:"clover"`
				IsOfficial   bool   `json:"is_official"`
				IsPowerBadge bool   `json:"is_power_badge"`
			} `json:"shop"`
			WholesalePrice     []interface{} `json:"wholesale_price"`
			CourierCount       int           `json:"courier_count"`
			Condition          int           `json:"condition"`
			CategoryID         int           `json:"category_id"`
			CategoryName       string        `json:"category_name"`
			CategoryBreadcrumb string        `json:"category_breadcrumb"`
			DepartmentID       int           `json:"department_id"`
			Labels             []struct {
				Title string `json:"title"`
				Color string `json:"color"`
			} `json:"labels"`
			Badges []struct {
				Title    string `json:"title"`
				ImageURL string `json:"image_url"`
				Show     bool   `json:"show"`
			} `json:"badges"`
			IsFeatured         int       `json:"is_featured"`
			Rating             int       `json:"rating"`
			CountReview        int       `json:"count_review"`
			OriginalPrice      string    `json:"original_price"`
			DiscountExpired    time.Time `json:"discount_expired"`
			DiscountPercentage int       `json:"discount_percentage"`
			Sku                string    `json:"sku"`
			Stock              int       `json:"stock"`
			Status             int       `json:"status"`
			IsPreorder         bool      `json:"is_preorder"`
		} `json:"products"`
	} `json:"data"`
}

type ProductTokped3 struct {
	Header struct {
		ProcessTime int    `json:"process_time"`
		Messages    string `json:"messages"`
		ErrorCode   string `json:"error_code"`
		Reason      string `json:"reason"`
	} `json:"header"`
	Data []struct {
		Basic struct {
			ProductID       int64  `json:"productID"`
			ShopID          int    `json:"shopID"`
			Status          int    `json:"status"`
			Name            string `json:"name"`
			MustInsurance   bool   `json:"mustInsurance"`
			Condition       int    `json:"condition"`
			ChildCategoryID int    `json:"childCategoryID"`
			ShortDesc       string `json:"shortDesc"`
			CreateTimeUnix  int    `json:"createTimeUnix"`
			UpdateTimeUnix  int    `json:"updateTimeUnix"`
		} `json:"basic"`
		Price struct {
			Value          int `json:"value"`
			Currency       int `json:"currency"`
			LastUpdateUnix int `json:"LastUpdateUnix"`
			Idr            int `json:"idr"`
		} `json:"price"`
		Weight struct {
			Value int `json:"value"`
			Unit  int `json:"unit"`
		} `json:"weight"`
		Stock struct {
			UseStock bool `json:"useStock"`
		} `json:"stock"`
		Variant struct {
			IsParent   bool    `json:"isParent"`
			IsVariant  bool    `json:"isVariant"`
			ChildrenID []int64 `json:"childrenID"`
		} `json:"variant"`
		Menu struct {
		} `json:"menu"`
		Preorder struct {
		} `json:"preorder"`
		ExtraAttribute struct {
			MinOrder           int  `json:"minOrder"`
			LastUpdateCategory int  `json:"lastUpdateCategory"`
			IsEligibleCOD      bool `json:"isEligibleCOD"`
			IsOnCampaign       bool `json:"isOnCampaign"`
		} `json:"extraAttribute"`
		CategoryTree []struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Title         string `json:"title"`
			BreadcrumbURL string `json:"breadcrumbURL"`
		} `json:"categoryTree"`
		Pictures []struct {
			PicID        int64  `json:"picID"`
			FileName     string `json:"fileName"`
			FilePath     string `json:"filePath"`
			Status       int    `json:"status"`
			OriginalURL  string `json:"OriginalURL"`
			ThumbnailURL string `json:"ThumbnailURL"`
			Width        int    `json:"width"`
			Height       int    `json:"height"`
			URL300       string `json:"URL300"`
		} `json:"pictures"`
		GMStats struct {
		} `json:"GMStats,omitempty"`
		Stats struct {
			CountView int `json:"countView"`
		} `json:"stats,omitempty"`
		Other struct {
			Sku       string `json:"sku"`
			URL       string `json:"url"`
			MobileURL string `json:"mobileURL"`
		} `json:"other"`
		Campaign struct {
			StartDate time.Time `json:"StartDate"`
			EndDate   time.Time `json:"EndDate"`
		} `json:"campaign"`
		Volume struct {
			Length int `json:"length"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"volume"`
		Warehouses []struct {
			ProductID   int64 `json:"productID"`
			WarehouseID int   `json:"warehouseID"`
			Price       struct {
				Value          int `json:"value"`
				Currency       int `json:"currency"`
				LastUpdateUnix int `json:"LastUpdateUnix"`
				Idr            int `json:"idr"`
			} `json:"price"`
			Stock struct {
				UseStock bool `json:"useStock"`
			} `json:"stock"`
		} `json:"warehouses"`
		Bundle   interface{} `json:"bundle"`
		LastSort string      `json:"last_sort"`
		GMStats0 struct {
			TransactionSuccess int `json:"transactionSuccess"`
			CountSold          int `json:"countSold"`
		} `json:"GMStats,omitempty"`
		Stats0 struct {
			CountView   int `json:"countView"`
			CountReview int `json:"countReview"`
			CountTalk   int `json:"countTalk"`
			Rating      int `json:"rating"`
		} `json:"stats,omitempty"`
	} `json:"data"`
}

type ProductDetailTokped struct {
	Header struct {
		ProcessTime int    `json:"process_time"`
		Messages    string `json:"messages"`
		Reason      string `json:"reason"`
		ErrorCode   string `json:"error_code"`
	} `json:"header"`
	Data []struct {
		Basic struct {
			ProductID       int    `json:"productID"`
			ShopID          int    `json:"shopID"`
			Status          int    `json:"status"`
			Name            string `json:"name"`
			MustInsurance   bool   `json:"mustInsurance"`
			Condition       int    `json:"condition"`
			ChildCategoryID int    `json:"childCategoryID"`
			ShortDesc       string `json:"shortDesc"`
			CreateTimeUnix  int    `json:"createTimeUnix"`
			UpdateTimeUnix  int    `json:"updateTimeUnix"`
		} `json:"basic"`
		Price struct {
			Value          int `json:"value"`
			Currency       int `json:"currency"`
			LastUpdateUnix int `json:"LastUpdateUnix"`
			Idr            int `json:"idr"`
		} `json:"price"`
		Weight struct {
			Value int `json:"value"`
			Unit  int `json:"unit"`
		} `json:"weight"`
		Stock struct {
			UseStock     bool   `json:"useStock"`
			Value        int    `json:"value"`
			StockWording string `json:"stockWording"`
		} `json:"stock"`
		MainStock int `json:"main_stock"`
		Variant   struct {
		} `json:"variant"`
		Menu struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"menu"`
		Preorder struct {
		} `json:"preorder"`
		ExtraAttribute struct {
			MinOrder           int  `json:"minOrder"`
			LastUpdateCategory int  `json:"lastUpdateCategory"`
			IsEligibleCOD      bool `json:"isEligibleCOD"`
			IsOnCampaign       bool `json:"isOnCampaign"`
		} `json:"extraAttribute"`
		CategoryTree []struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Title         string `json:"title"`
			BreadcrumbURL string `json:"breadcrumbURL"`
		} `json:"categoryTree"`
		Pictures []struct {
			PicID        int64  `json:"picID"`
			FileName     string `json:"fileName"`
			FilePath     string `json:"filePath"`
			Status       int    `json:"status"`
			OriginalURL  string `json:"OriginalURL"`
			ThumbnailURL string `json:"ThumbnailURL"`
			Width        int    `json:"width"`
			Height       int    `json:"height"`
			URL300       string `json:"URL300"`
		} `json:"pictures"`
		GMStats struct {
			TransactionSuccess int `json:"transactionSuccess"`
			TransactionReject  int `json:"transactionReject"`
			CountSold          int `json:"countSold"`
		} `json:"GMStats"`
		Stats struct {
			CountView   int `json:"countView"`
			CountReview int `json:"countReview"`
			CountTalk   int `json:"countTalk"`
			Rating      int `json:"rating"`
		} `json:"stats"`
		Other struct {
			Sku       string `json:"sku"`
			URL       string `json:"url"`
			MobileURL string `json:"mobileURL"`
		} `json:"other"`
		Campaign struct {
			IsActive           bool      `json:"IsActive"`
			OriginalPrice      int       `json:"OriginalPrice"`
			DiscountPercentage int       `json:"DiscountPercentage"`
			DiscountPrice      int       `json:"DiscountPrice"`
			CampaignTypeName   string    `json:"CampaignTypeName"`
			StartDate          time.Time `json:"StartDate"`
			EndDate            time.Time `json:"EndDate"`
			Stock              int       `json:"Stock"`
			MinOrder           int       `json:"MinOrder"`
			OriginalStock      int       `json:"OriginalStock"`
		} `json:"campaign"`
		Volume struct {
		} `json:"volume"`
		Warehouses []struct {
			ProductID   int `json:"productID"`
			WarehouseID int `json:"warehouseID"`
			Price       struct {
				Value          int `json:"value"`
				Currency       int `json:"currency"`
				LastUpdateUnix int `json:"LastUpdateUnix"`
				Idr            int `json:"idr"`
			} `json:"price"`
			Stock struct {
				UseStock bool `json:"useStock"`
				Value    int  `json:"value"`
			} `json:"stock"`
		} `json:"warehouses"`
		Bundle   interface{} `json:"bundle"`
		LastSort string      `json:"last_sort"`
	} `json:"data"`
}

type StockTokped struct {
	Header struct {
		ProcessTime int    `json:"process_time"`
		Messages    string `json:"messages"`
		ErrorCode   string `json:"error_code"`
		Reason      string `json:"reason"`
	} `json:"header"`
	Data struct {
		TotalData       int `json:"total_data"`
		SucceedRows     int `json:"succeed_rows"`
		SucceedRowsData []struct {
			ProductID   int64 `json:"productID"`
			WarehouseID int   `json:"warehouseID"`
			ShopID      int   `json:"shopID"`
			Stock       int   `json:"stock"`
			Price       int   `json:"price"`
		} `json:"succeed_rows_data"`
		FailedRows     int           `json:"failed_rows"`
		FailedRowsData []interface{} `json:"failed_rows_data"`
	} `json:"data"`
}

type OrdersTokped struct {
	Header struct {
		ProcessTime int    `json:"process_time"`
		Messages    string `json:"messages"`
		ErrorCode   string `json:"error_code"`
		Reason      string `json:"reason"`
	} `json:"header"`
	Data []struct {
		FsID              string `json:"fs_id"`
		OrderID           int    `json:"order_id"`
		IsCodMitra        bool   `json:"is_cod_mitra"`
		AcceptPartial     bool   `json:"accept_partial"`
		InvoiceRefNum     string `json:"invoice_ref_num"`
		HaveProductBundle bool   `json:"have_product_bundle"`
		Products          []struct {
			ID          int64   `json:"id"`
			Name        string  `json:"name"`
			Quantity    int     `json:"quantity"`
			Notes       string  `json:"notes"`
			Weight      float64 `json:"weight"`
			TotalWeight float64 `json:"total_weight"`
			Price       int     `json:"price"`
			TotalPrice  int     `json:"total_price"`
			Currency    string  `json:"currency"`
			Sku         string  `json:"sku"`
			IsWholesale bool    `json:"is_wholesale"`
		} `json:"products"`
		ProductsFulfilled []struct {
			ProductID       int64 `json:"product_id"`
			QuantityDeliver int   `json:"quantity_deliver"`
			QuantityReject  int   `json:"quantity_reject"`
		} `json:"products_fulfilled"`
		BundleDetail struct {
			Bundle       interface{} `json:"bundle"`
			NonBundle    interface{} `json:"non_bundle"`
			TotalProduct int         `json:"total_product"`
		} `json:"bundle_detail"`
		DeviceType string `json:"device_type"`
		Buyer      struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Phone      string `json:"phone"`
			Email      string `json:"email"`
			UserStatus int    `json:"user_status"`
		} `json:"buyer"`
		ShopID      int       `json:"shop_id"`
		PaymentID   int       `json:"payment_id"`
		PaymentDate time.Time `json:"payment_date"`
		Recipient   struct {
			Name    string `json:"name"`
			Phone   string `json:"phone"`
			Address struct {
				AddressFull string `json:"address_full"`
				District    string `json:"district"`
				City        string `json:"city"`
				Province    string `json:"province"`
				Country     string `json:"country"`
				PostalCode  string `json:"postal_code"`
				DistrictID  int    `json:"district_id"`
				CityID      int    `json:"city_id"`
				ProvinceID  int    `json:"province_id"`
				Geo         string `json:"geo"`
			} `json:"address"`
		} `json:"recipient"`
		Logistics struct {
			ShippingID     int    `json:"shipping_id"`
			DistrictID     int    `json:"district_id"`
			CityID         int    `json:"city_id"`
			ProvinceID     int    `json:"province_id"`
			Geo            string `json:"geo"`
			ShippingAgency string `json:"shipping_agency"`
			ServiceType    string `json:"service_type"`
		} `json:"logistics"`
		Amt struct {
			TTLProductPrice int `json:"ttl_product_price"`
			ShippingCost    int `json:"shipping_cost"`
			InsuranceCost   int `json:"insurance_cost"`
			TTLAmount       int `json:"ttl_amount"`
			VoucherAmount   int `json:"voucher_amount"`
			ToppointsAmount int `json:"toppoints_amount"`
		} `json:"amt"`
		DropshipperInfo struct {
		} `json:"dropshipper_info"`
		VoucherInfo struct {
			VoucherCode string `json:"voucher_code"`
			VoucherType int    `json:"voucher_type"`
		} `json:"voucher_info"`
		OrderStatus  int `json:"order_status"`
		WarehouseID  int `json:"warehouse_id"`
		FulfillBy    int `json:"fulfill_by"`
		CreateTime   int `json:"create_time"`
		CustomFields struct {
			Awb string `json:"awb"`
		} `json:"custom_fields"`
		PromoOrderDetail struct {
			OrderID               int `json:"order_id"`
			TotalCashback         int `json:"total_cashback"`
			TotalDiscount         int `json:"total_discount"`
			TotalDiscountProduct  int `json:"total_discount_product"`
			TotalDiscountShipping int `json:"total_discount_shipping"`
			TotalDiscountDetails  []struct {
				Amount int    `json:"amount"`
				Type   string `json:"type"`
			} `json:"total_discount_details"`
			SummaryPromo []struct {
				Name               string      `json:"name"`
				IsCoupon           bool        `json:"is_coupon"`
				ShowCashbackAmount bool        `json:"show_cashback_amount"`
				ShowDiscountAmount bool        `json:"show_discount_amount"`
				CashbackAmount     int         `json:"cashback_amount"`
				CashbackPoints     int         `json:"cashback_points"`
				CashbackDetails    interface{} `json:"cashback_details"`
				Type               string      `json:"type"`
				DiscountAmount     int         `json:"discount_amount"`
				DiscountDetails    []struct {
					Amount        int    `json:"amount"`
					Type          string `json:"type"`
					BudgetDetails []struct {
						BudgetType          int `json:"budget_type"`
						BenefitAmount       int `json:"benefit_amount"`
						ActualBenefitAmount int `json:"actual_benefit_amount"`
					} `json:"budget_details"`
				} `json:"discount_details"`
				InvoiceDesc string `json:"invoice_desc"`
			} `json:"summary_promo"`
		} `json:"promo_order_detail"`
		Encryption struct {
			Secret  string `json:"secret"`
			Content string `json:"content"`
			Message string `json:"message"`
		} `json:"encryption"`
		AddonInfo           interface{} `json:"addon_info"`
		ShipmentFulfillment struct {
			AcceptDeadline          time.Time `json:"accept_deadline"`
			ConfirmShippingDeadline time.Time `json:"confirm_shipping_deadline"`
		} `json:"shipment_fulfillment"`
		IsPlus bool `json:"is_plus"`
	} `json:"data"`
}

type OrderDetailTokped struct {
	Header struct {
		ProcessTime int    `json:"process_time"`
		Messages    string `json:"messages"`
		ErrorCode   string `json:"error_code"`
		Reason      string `json:"reason"`
	} `json:"header"`
	Data struct {
		OrderID        int  `json:"order_id"`
		BuyerID        int  `json:"buyer_id"`
		SellerID       int  `json:"seller_id"`
		PaymentID      int  `json:"payment_id"`
		IsAffiliate    bool `json:"is_affiliate"`
		IsFulfillment  bool `json:"is_fulfillment"`
		OrderWarehouse struct {
			WarehouseID int `json:"warehouse_id"`
			FulfillBy   int `json:"fulfill_by"`
			MetaData    struct {
				WarehouseID   int    `json:"warehouse_id"`
				PartnerID     int    `json:"partner_id"`
				ShopID        int    `json:"shop_id"`
				WarehouseName string `json:"warehouse_name"`
				DistrictID    int    `json:"district_id"`
				DistrictName  string `json:"district_name"`
				CityID        int    `json:"city_id"`
				CityName      string `json:"city_name"`
				ProvinceID    int    `json:"province_id"`
				ProvinceName  string `json:"province_name"`
				Status        int    `json:"status"`
				PostalCode    string `json:"postal_code"`
				IsDefault     int    `json:"is_default"`
				Latlon        string `json:"latlon"`
				Latitude      string `json:"latitude"`
				Longitude     string `json:"longitude"`
				Email         string `json:"email"`
				AddressDetail string `json:"address_detail"`
				CountryName   string `json:"country_name"`
				IsFulfillment bool   `json:"is_fulfillment"`
			} `json:"meta_data"`
		} `json:"order_warehouse"`
		OrderStatus   int    `json:"order_status"`
		InvoiceNumber string `json:"invoice_number"`
		InvoicePdf    string `json:"invoice_pdf"`
		InvoiceURL    string `json:"invoice_url"`
		OpenAmt       int    `json:"open_amt"`
		LpAmt         int    `json:"lp_amt"`
		CashbackAmt   int    `json:"cashback_amt"`
		Info          string `json:"info"`
		Comment       string `json:"comment"`
		ItemPrice     int    `json:"item_price"`
		BuyerInfo     struct {
			BuyerID       int    `json:"buyer_id"`
			BuyerFullname string `json:"buyer_fullname"`
			BuyerEmail    string `json:"buyer_email"`
			BuyerPhone    string `json:"buyer_phone"`
		} `json:"buyer_info"`
		ShopInfo struct {
			ShopOwnerID    int    `json:"shop_owner_id"`
			ShopOwnerEmail string `json:"shop_owner_email"`
			ShopOwnerForm  string `json:"shop_owner_form"`
			ShopName       string `json:"shop_name"`
			ShopDomain     string `json:"shop_domain"`
			ShopID         int    `json:"shop_id"`
			LastLoginAt    string `json:"last_login_at"`
		} `json:"shop_info"`
		ShipmentFulfillment struct {
			ID                      int       `json:"id"`
			OrderID                 int       `json:"order_id"`
			PaymentDateTime         time.Time `json:"payment_date_time"`
			IsSameDay               bool      `json:"is_same_day"`
			AcceptDeadline          time.Time `json:"accept_deadline"`
			ConfirmShippingDeadline time.Time `json:"confirm_shipping_deadline"`
			ItemDeliveredDeadline   struct {
				Time  time.Time `json:"Time"`
				Valid bool      `json:"Valid"`
			} `json:"item_delivered_deadline"`
			IsAccepted        bool `json:"is_accepted"`
			IsConfirmShipping bool `json:"is_confirm_shipping"`
			IsItemDelivered   bool `json:"is_item_delivered"`
			FulfillmentStatus int  `json:"fulfillment_status"`
		} `json:"shipment_fulfillment"`
		Preorder struct {
			OrderID              int    `json:"order_id"`
			PreorderType         int    `json:"preorder_type"`
			PreorderProcessTime  int    `json:"preorder_process_time"`
			PreorderProcessStart string `json:"preorder_process_start"`
			PreorderDeadline     string `json:"preorder_deadline"`
			ShopID               int    `json:"shop_id"`
			CustomerID           int    `json:"customer_id"`
		} `json:"preorder"`
		OrderInfo struct {
			OrderDetail []struct {
				OrderDetailID   int64   `json:"order_detail_id"`
				ProductID       int     `json:"product_id"`
				ProductName     string  `json:"product_name"`
				ProductDescPdp  string  `json:"product_desc_pdp"`
				ProductDescAtc  string  `json:"product_desc_atc"`
				ProductPrice    float64 `json:"product_price"`
				SubtotalPrice   int     `json:"subtotal_price"`
				Weight          float64 `json:"weight"`
				TotalWeight     float64 `json:"total_weight"`
				Quantity        int     `json:"quantity"`
				QuantityDeliver int     `json:"quantity_deliver"`
				QuantityReject  int     `json:"quantity_reject"`
				IsFreeReturns   bool    `json:"is_free_returns"`
				InsurancePrice  int     `json:"insurance_price"`
				NormalPrice     int     `json:"normal_price"`
				CurrencyID      int     `json:"currency_id"`
				CurrencyRate    int     `json:"currency_rate"`
				MinOrder        int     `json:"min_order"`
				ChildCatID      int     `json:"child_cat_id"`
				CampaignID      string  `json:"campaign_id"`
				ProductPicture  string  `json:"product_picture"`
				SnapshotURL     string  `json:"snapshot_url"`
				Sku             string  `json:"sku"`
			} `json:"order_detail"`
			OrderHistory []struct {
				ActionBy       string    `json:"action_by"`
				HistStatusCode int       `json:"hist_status_code"`
				Message        string    `json:"message"`
				Timestamp      time.Time `json:"timestamp"`
				Comment        string    `json:"comment"`
				CreateBy       int       `json:"create_by"`
				UpdateBy       string    `json:"update_by"`
			} `json:"order_history"`
			OrderAgeDay     int  `json:"order_age_day"`
			ShippingAgeDay  int  `json:"shipping_age_day"`
			DeliveredAgeDay int  `json:"delivered_age_day"`
			PartialProcess  bool `json:"partial_process"`
			ShippingInfo    struct {
				SpID                   int    `json:"sp_id"`
				ShippingID             int    `json:"shipping_id"`
				LogisticName           string `json:"logistic_name"`
				LogisticService        string `json:"logistic_service"`
				ShippingPrice          int    `json:"shipping_price"`
				ShippingPriceRate      int    `json:"shipping_price_rate"`
				ShippingFee            int    `json:"shipping_fee"`
				InsurancePrice         int    `json:"insurance_price"`
				Fee                    int    `json:"fee"`
				IsChangeCourier        bool   `json:"is_change_courier"`
				SecondSpID             int    `json:"second_sp_id"`
				SecondShippingID       int    `json:"second_shipping_id"`
				SecondLogisticName     string `json:"second_logistic_name"`
				SecondLogisticService  string `json:"second_logistic_service"`
				SecondAgencyFee        int    `json:"second_agency_fee"`
				SecondInsurance        int    `json:"second_insurance"`
				SecondRate             int    `json:"second_rate"`
				Awb                    string `json:"awb"`
				AutoresiCashlessStatus int    `json:"autoresi_cashless_status"`
				AutoresiAwb            string `json:"autoresi_awb"`
				AutoresiShippingPrice  int    `json:"autoresi_shipping_price"`
				CountAwb               int    `json:"count_awb"`
				IsCashless             bool   `json:"isCashless"`
				IsFakeDelivery         bool   `json:"is_fake_delivery"`
			} `json:"shipping_info"`
			Destination struct {
				ReceiverName      string `json:"receiver_name"`
				ReceiverPhone     string `json:"receiver_phone"`
				AddressStreet     string `json:"address_street"`
				AddressDistrict   string `json:"address_district"`
				AddressCity       string `json:"address_city"`
				AddressProvince   string `json:"address_province"`
				AddressPostal     string `json:"address_postal"`
				CustomerAddressID int    `json:"customer_address_id"`
				DistrictID        int    `json:"district_id"`
				CityID            int    `json:"city_id"`
				ProvinceID        int    `json:"province_id"`
			} `json:"destination"`
			IsReplacement         bool `json:"is_replacement"`
			ReplacementMultiplier int  `json:"replacement_multiplier"`
		} `json:"order_info"`
		OriginInfo struct {
			SenderName            string `json:"sender_name"`
			OriginProvince        int    `json:"origin_province"`
			OriginProvinceName    string `json:"origin_province_name"`
			OriginCity            int    `json:"origin_city"`
			OriginCityName        string `json:"origin_city_name"`
			OriginAddress         string `json:"origin_address"`
			OriginDistrict        int    `json:"origin_district"`
			OriginDistrictName    string `json:"origin_district_name"`
			OriginPostalCode      string `json:"origin_postal_code"`
			OriginGeo             string `json:"origin_geo"`
			ReceiverName          string `json:"receiver_name"`
			DestinationAddress    string `json:"destination_address"`
			DestinationProvince   int    `json:"destination_province"`
			DestinationCity       int    `json:"destination_city"`
			DestinationDistrict   int    `json:"destination_district"`
			DestinationPostalCode string `json:"destination_postal_code"`
			DestinationGeo        string `json:"destination_geo"`
			DestinationLoc        struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"destination_loc"`
		} `json:"origin_info"`
		PaymentInfo struct {
			PaymentID       int       `json:"payment_id"`
			PaymentRefNum   string    `json:"payment_ref_num"`
			PaymentDate     time.Time `json:"payment_date"`
			PaymentMethod   int       `json:"payment_method"`
			PaymentStatus   string    `json:"payment_status"`
			PaymentStatusID int       `json:"payment_status_id"`
			CreateTime      time.Time `json:"create_time"`
			PgID            int       `json:"pg_id"`
			GatewayName     string    `json:"gateway_name"`
			DiscountAmount  int       `json:"discount_amount"`
			VoucherCode     string    `json:"voucher_code"`
			VoucherID       int       `json:"voucher_id"`
		} `json:"payment_info"`
		InsuranceInfo struct {
			InsuranceType int `json:"insurance_type"`
		} `json:"insurance_info"`
		HoldInfo          interface{} `json:"hold_info"`
		CancelRequestInfo interface{} `json:"cancel_request_info"`
		CreateTime        time.Time   `json:"create_time"`
		ShippingDate      string      `json:"shipping_date"`
		UpdateTime        time.Time   `json:"update_time"`
		PaymentDate       time.Time   `json:"payment_date"`
		DeliveredDate     string      `json:"delivered_date"`
		EstShippingDate   string      `json:"est_shipping_date"`
		EstDeliveryDate   string      `json:"est_delivery_date"`
		RelatedInvoices   interface{} `json:"related_invoices"`
		CustomFields      interface{} `json:"custom_fields"`
		PromoOrderDetail  struct {
			OrderID               int `json:"order_id"`
			TotalCashback         int `json:"total_cashback"`
			TotalDiscount         int `json:"total_discount"`
			TotalDiscountProduct  int `json:"total_discount_product"`
			TotalDiscountShipping int `json:"total_discount_shipping"`
			TotalDiscountDetails  []struct {
				Amount int    `json:"amount"`
				Type   string `json:"type"`
			} `json:"total_discount_details"`
			SummaryPromo []struct {
				Name               string      `json:"name"`
				IsCoupon           bool        `json:"is_coupon"`
				ShowCashbackAmount bool        `json:"show_cashback_amount"`
				ShowDiscountAmount bool        `json:"show_discount_amount"`
				CashbackAmount     int         `json:"cashback_amount"`
				CashbackPoints     int         `json:"cashback_points"`
				CashbackDetails    interface{} `json:"cashback_details"`
				Type               string      `json:"type"`
				DiscountAmount     int         `json:"discount_amount"`
				DiscountDetails    []struct {
					Amount        int    `json:"amount"`
					Type          string `json:"type"`
					BudgetDetails []struct {
						BudgetType          int `json:"budget_type"`
						BenefitAmount       int `json:"benefit_amount"`
						ActualBenefitAmount int `json:"actual_benefit_amount"`
					} `json:"budget_details"`
				} `json:"discount_details"`
				InvoiceDesc string `json:"invoice_desc"`
			} `json:"summary_promo"`
		} `json:"promo_order_detail"`
		Encryption struct {
			Secret  string `json:"secret"`
			Content string `json:"content"`
			Message string `json:"message"`
		} `json:"encryption"`
		IsPlus bool `json:"is_plus"`
	} `json:"data"`
}

type OrderDetailTokpedSingle struct {
	Header struct {
		ProcessTime int    `json:"process_time"`
		Messages    string `json:"messages"`
		ErrorCode   string `json:"error_code"`
		Reason      string `json:"reason"`
	} `json:"header"`
	Data struct {
		OrderID        int  `json:"order_id"`
		BuyerID        int  `json:"buyer_id"`
		SellerID       int  `json:"seller_id"`
		PaymentID      int  `json:"payment_id"`
		IsAffiliate    bool `json:"is_affiliate"`
		IsFulfillment  bool `json:"is_fulfillment"`
		OrderWarehouse struct {
			WarehouseID int `json:"warehouse_id"`
			FulfillBy   int `json:"fulfill_by"`
			MetaData    struct {
				WarehouseID   int    `json:"warehouse_id"`
				PartnerID     int    `json:"partner_id"`
				ShopID        int    `json:"shop_id"`
				WarehouseName string `json:"warehouse_name"`
				DistrictID    int    `json:"district_id"`
				DistrictName  string `json:"district_name"`
				CityID        int    `json:"city_id"`
				CityName      string `json:"city_name"`
				ProvinceID    int    `json:"province_id"`
				ProvinceName  string `json:"province_name"`
				Status        int    `json:"status"`
				PostalCode    string `json:"postal_code"`
				IsDefault     int    `json:"is_default"`
				Latlon        string `json:"latlon"`
				Latitude      string `json:"latitude"`
				Longitude     string `json:"longitude"`
				Email         string `json:"email"`
				AddressDetail string `json:"address_detail"`
				CountryName   string `json:"country_name"`
				IsFulfillment bool   `json:"is_fulfillment"`
			} `json:"meta_data"`
		} `json:"order_warehouse"`
		OrderStatus   int    `json:"order_status"`
		InvoiceNumber string `json:"invoice_number"`
		InvoicePdf    string `json:"invoice_pdf"`
		InvoiceURL    string `json:"invoice_url"`
		OpenAmt       int    `json:"open_amt"`
		LpAmt         int    `json:"lp_amt"`
		CashbackAmt   int    `json:"cashback_amt"`
		Info          string `json:"info"`
		Comment       string `json:"comment"`
		ItemPrice     int    `json:"item_price"`
		BuyerInfo     struct {
			BuyerID       int    `json:"buyer_id"`
			BuyerFullname string `json:"buyer_fullname"`
			BuyerEmail    string `json:"buyer_email"`
			BuyerPhone    string `json:"buyer_phone"`
		} `json:"buyer_info"`
		ShopInfo struct {
			ShopOwnerID    int    `json:"shop_owner_id"`
			ShopOwnerEmail string `json:"shop_owner_email"`
			ShopOwnerForm  string `json:"shop_owner_form"`
			ShopName       string `json:"shop_name"`
			ShopDomain     string `json:"shop_domain"`
			ShopID         int    `json:"shop_id"`
			LastLoginAt    string `json:"last_login_at"`
		} `json:"shop_info"`
		ShipmentFulfillment struct {
			ID                      int       `json:"id"`
			OrderID                 int       `json:"order_id"`
			PaymentDateTime         time.Time `json:"payment_date_time"`
			IsSameDay               bool      `json:"is_same_day"`
			AcceptDeadline          time.Time `json:"accept_deadline"`
			ConfirmShippingDeadline time.Time `json:"confirm_shipping_deadline"`
			ItemDeliveredDeadline   struct {
				Time  string `json:"Time"`
				Valid bool   `json:"Valid"`
			} `json:"item_delivered_deadline"`
			IsAccepted        bool `json:"is_accepted"`
			IsConfirmShipping bool `json:"is_confirm_shipping"`
			IsItemDelivered   bool `json:"is_item_delivered"`
			FulfillmentStatus int  `json:"fulfillment_status"`
		} `json:"shipment_fulfillment"`
		Preorder struct {
			OrderID              int    `json:"order_id"`
			PreorderType         int    `json:"preorder_type"`
			PreorderProcessTime  int    `json:"preorder_process_time"`
			PreorderProcessStart string `json:"preorder_process_start"`
			PreorderDeadline     string `json:"preorder_deadline"`
			ShopID               int    `json:"shop_id"`
			CustomerID           int    `json:"customer_id"`
		} `json:"preorder"`
		OrderInfo struct {
			OrderDetail []struct {
				OrderDetailID   int64   `json:"order_detail_id"`
				ProductID       int64   `json:"product_id"`
				ProductName     string  `json:"product_name"`
				ProductDescPdp  string  `json:"product_desc_pdp"`
				ProductDescAtc  string  `json:"product_desc_atc"`
				ProductPrice    float64 `json:"product_price"`
				SubtotalPrice   int     `json:"subtotal_price"`
				Weight          float64 `json:"weight"`
				TotalWeight     float64 `json:"total_weight"`
				Quantity        int     `json:"quantity"`
				QuantityDeliver int     `json:"quantity_deliver"`
				QuantityReject  int     `json:"quantity_reject"`
				IsFreeReturns   bool    `json:"is_free_returns"`
				InsurancePrice  int     `json:"insurance_price"`
				NormalPrice     int     `json:"normal_price"`
				CurrencyID      int     `json:"currency_id"`
				CurrencyRate    int     `json:"currency_rate"`
				MinOrder        int     `json:"min_order"`
				ChildCatID      int     `json:"child_cat_id"`
				CampaignID      string  `json:"campaign_id"`
				ProductPicture  string  `json:"product_picture"`
				SnapshotURL     string  `json:"snapshot_url"`
				Sku             string  `json:"sku"`
			} `json:"order_detail"`
			OrderHistory []struct {
				ActionBy       string    `json:"action_by"`
				HistStatusCode int       `json:"hist_status_code"`
				Message        string    `json:"message"`
				Timestamp      time.Time `json:"timestamp"`
				Comment        string    `json:"comment"`
				CreateBy       int       `json:"create_by"`
				UpdateBy       string    `json:"update_by"`
			} `json:"order_history"`
			OrderAgeDay     int  `json:"order_age_day"`
			ShippingAgeDay  int  `json:"shipping_age_day"`
			DeliveredAgeDay int  `json:"delivered_age_day"`
			PartialProcess  bool `json:"partial_process"`
			ShippingInfo    struct {
				SpID                   int         `json:"sp_id"`
				ShippingID             int         `json:"shipping_id"`
				LogisticName           string      `json:"logistic_name"`
				LogisticService        string      `json:"logistic_service"`
				ShippingPrice          int         `json:"shipping_price"`
				ShippingPriceRate      int         `json:"shipping_price_rate"`
				ShippingFee            int         `json:"shipping_fee"`
				InsurancePrice         int         `json:"insurance_price"`
				Fee                    int         `json:"fee"`
				IsChangeCourier        bool        `json:"is_change_courier"`
				SecondSpID             int         `json:"second_sp_id"`
				SecondShippingID       int         `json:"second_shipping_id"`
				SecondLogisticName     string      `json:"second_logistic_name"`
				SecondLogisticService  string      `json:"second_logistic_service"`
				SecondAgencyFee        int         `json:"second_agency_fee"`
				SecondInsurance        int         `json:"second_insurance"`
				SecondRate             int         `json:"second_rate"`
				Awb                    string      `json:"awb"`
				AutoresiCashlessStatus int         `json:"autoresi_cashless_status"`
				AutoresiAwb            string      `json:"autoresi_awb"`
				AutoresiShippingPrice  int         `json:"autoresi_shipping_price"`
				CountAwb               int         `json:"count_awb"`
				IsCashless             bool        `json:"isCashless"`
				IsFakeDelivery         bool        `json:"is_fake_delivery"`
				RecommendedCourierInfo interface{} `json:"recommended_courier_info"`
			} `json:"shipping_info"`
			Destination struct {
				ReceiverName      string `json:"receiver_name"`
				ReceiverPhone     string `json:"receiver_phone"`
				AddressStreet     string `json:"address_street"`
				AddressDistrict   string `json:"address_district"`
				AddressCity       string `json:"address_city"`
				AddressProvince   string `json:"address_province"`
				AddressPostal     string `json:"address_postal"`
				CustomerAddressID int    `json:"customer_address_id"`
				DistrictID        int    `json:"district_id"`
				CityID            int    `json:"city_id"`
				ProvinceID        int    `json:"province_id"`
			} `json:"destination"`
			IsReplacement         bool `json:"is_replacement"`
			ReplacementMultiplier int  `json:"replacement_multiplier"`
			IsPlus                bool `json:"is_plus"`
		} `json:"order_info"`
		OriginInfo struct {
			SenderName            string `json:"sender_name"`
			OriginProvince        int    `json:"origin_province"`
			OriginProvinceName    string `json:"origin_province_name"`
			OriginCity            int    `json:"origin_city"`
			OriginCityName        string `json:"origin_city_name"`
			OriginAddress         string `json:"origin_address"`
			OriginDistrict        int    `json:"origin_district"`
			OriginDistrictName    string `json:"origin_district_name"`
			OriginPostalCode      string `json:"origin_postal_code"`
			OriginGeo             string `json:"origin_geo"`
			ReceiverName          string `json:"receiver_name"`
			DestinationAddress    string `json:"destination_address"`
			DestinationProvince   int    `json:"destination_province"`
			DestinationCity       int    `json:"destination_city"`
			DestinationDistrict   int    `json:"destination_district"`
			DestinationPostalCode string `json:"destination_postal_code"`
			DestinationGeo        string `json:"destination_geo"`
			DestinationLoc        struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"destination_loc"`
		} `json:"origin_info"`
		PaymentInfo struct {
			PaymentID       int       `json:"payment_id"`
			PaymentRefNum   string    `json:"payment_ref_num"`
			PaymentDate     time.Time `json:"payment_date"`
			PaymentMethod   int       `json:"payment_method"`
			PaymentStatus   string    `json:"payment_status"`
			PaymentStatusID int       `json:"payment_status_id"`
			CreateTime      time.Time `json:"create_time"`
			PgID            int       `json:"pg_id"`
			GatewayName     string    `json:"gateway_name"`
			DiscountAmount  int       `json:"discount_amount"`
			VoucherCode     string    `json:"voucher_code"`
			VoucherID       int       `json:"voucher_id"`
		} `json:"payment_info"`
		InsuranceInfo struct {
			InsuranceType int `json:"insurance_type"`
		} `json:"insurance_info"`
		HoldInfo          interface{} `json:"hold_info"`
		CancelRequestInfo interface{} `json:"cancel_request_info"`
		CreateTime        time.Time   `json:"create_time"`
		ShippingDate      string      `json:"shipping_date"`
		UpdateTime        time.Time   `json:"update_time"`
		PaymentDate       time.Time   `json:"payment_date"`
		DeliveredDate     string      `json:"delivered_date"`
		EstShippingDate   string      `json:"est_shipping_date"`
		EstDeliveryDate   string      `json:"est_delivery_date"`
		RelatedInvoices   interface{} `json:"related_invoices"`
		CustomFields      interface{} `json:"custom_fields"`
		PromoOrderDetail  struct {
			OrderID               int         `json:"order_id"`
			TotalCashback         int         `json:"total_cashback"`
			TotalDiscount         int         `json:"total_discount"`
			TotalDiscountProduct  int         `json:"total_discount_product"`
			TotalDiscountShipping int         `json:"total_discount_shipping"`
			TotalDiscountDetails  interface{} `json:"total_discount_details"`
			SummaryPromo          []struct {
				Name               string `json:"name"`
				IsCoupon           bool   `json:"is_coupon"`
				ShowCashbackAmount bool   `json:"show_cashback_amount"`
				ShowDiscountAmount bool   `json:"show_discount_amount"`
				CashbackAmount     int    `json:"cashback_amount"`
				CashbackPoints     int    `json:"cashback_points"`
				CashbackDetails    []struct {
					AmountPoints       int    `json:"amount_points"`
					AmountIdr          int    `json:"amount_idr"`
					ActualAmountIdr    int    `json:"actual_amount_idr"`
					ActualAmountPoints int    `json:"actual_amount_points"`
					CurrencyType       string `json:"currency_type"`
					CurrencyTypeStr    string `json:"currency_type_str"`
					BudgetDetails      []struct {
						BudgetType          int `json:"budget_type"`
						BenefitAmount       int `json:"benefit_amount"`
						ActualBenefitAmount int `json:"actual_benefit_amount"`
					} `json:"budget_details"`
				} `json:"cashback_details"`
				Type            string      `json:"type"`
				DiscountAmount  int         `json:"discount_amount"`
				DiscountDetails interface{} `json:"discount_details"`
				InvoiceDesc     string      `json:"invoice_desc"`
			} `json:"summary_promo"`
		} `json:"promo_order_detail"`
		Encryption struct {
			Secret  string `json:"secret"`
			Content string `json:"content"`
			Message string `json:"message"`
		} `json:"encryption"`
		HaveProductBundle bool        `json:"have_product_bundle"`
		BundleDetail      interface{} `json:"bundle_detail"`
	} `json:"data"`
}

type DetailBookingTokped struct {
	Header struct {
		ProcessTime int    `json:"process_time"`
		Messages    string `json:"messages"`
		ErrorCode   string `json:"error_code"`
		Reason      string `json:"reason"`
	} `json:"header"`
	Data struct {
		OrderData []struct {
			Order struct {
				OrderID        int    `json:"order_id"`
				BuyerID        int    `json:"buyer_id"`
				SellerID       int    `json:"seller_id"`
				PaymentID      int    `json:"payment_id"`
				OrderStatus    int    `json:"order_status"`
				InvoiceNumber  string `json:"invoice_number"`
				InvoicePdfLink string `json:"invoice_pdf_link"`
				OpenAmt        int    `json:"open_amt"`
				PaymentAmtCod  int    `json:"payment_amt_cod"`
				ShopID         int    `json:"shop_id"`
				WarehouseID    int    `json:"warehouse_id"`
			} `json:"order"`
			OrderHistory []struct {
				OrderHistID  int64       `json:"order_hist_id"`
				Status       int         `json:"status"`
				ShippingDate interface{} `json:"shipping_date"`
				CreateBy     int         `json:"create_by"`
			} `json:"order_history"`
			OrderDetail []struct {
				OrderDetailID  int64   `json:"order_detail_id"`
				ProductID      int64   `json:"product_id"`
				ProductName    string  `json:"product_name"`
				Quantity       int     `json:"quantity"`
				ProductPrice   float64 `json:"product_price"`
				InsurancePrice int     `json:"insurance_price"`
			} `json:"order_detail"`
			DropShipper struct {
				OrderID      int    `json:"order_id"`
				DropshipName string `json:"dropship_name"`
				DropshipTelp string `json:"dropship_telp"`
			} `json:"drop_shipper"`
			TypeMeta struct {
				AffiliateData struct {
					AffiliateDetail interface{} `json:"affiliate_detail"`
				} `json:"affiliate_data"`
				B2B2C struct {
					IsB2B2C         bool `json:"is_b2b2c"`
					IsMultiCheckout bool `json:"is_multi_checkout"`
				} `json:"b2b2c"`
				B2BEnterprise struct {
				} `json:"b2b_enterprise"`
				Cod struct {
					Fee int `json:"fee"`
				} `json:"cod"`
				Kelontong struct {
				} `json:"kelontong"`
				Now struct {
				} `json:"now"`
				Ppp struct {
				} `json:"ppp"`
				ReadinessInsurance struct {
				} `json:"readiness_insurance"`
				Sampai struct {
				} `json:"sampai"`
				Tokonow struct {
					Type int `json:"type"`
				} `json:"tokonow"`
				TradeIn struct {
				} `json:"trade_in"`
				VehicleLeasing struct {
				} `json:"vehicle_leasing"`
			} `json:"type_meta"`
			OrderShipmentFulfillment struct {
				EstimatedArrivalMaxDate struct {
					Time  time.Time `json:"Time"`
					Valid bool      `json:"Valid"`
				} `json:"EstimatedArrivalMaxDate"`
				EstimatedArrivalMinDate struct {
					Time  time.Time `json:"Time"`
					Valid bool      `json:"Valid"`
				} `json:"EstimatedArrivalMinDate"`
				AcceptDeadline          time.Time `json:"accept_deadline"`
				ConfirmShippingDeadline time.Time `json:"confirm_shipping_deadline"`
				CreateTime              struct {
					Time  time.Time `json:"Time"`
					Valid bool      `json:"Valid"`
				} `json:"create_time"`
				EstimatedArrivalText struct {
					String string `json:"String"`
					Valid  bool   `json:"Valid"`
				} `json:"estimated_arrival_text"`
				FulfillmentStatus     int  `json:"fulfillment_status"`
				ID                    int  `json:"id"`
				IsAccepted            bool `json:"is_accepted"`
				IsConfirmShipping     bool `json:"is_confirm_shipping"`
				IsItemDelivered       bool `json:"is_item_delivered"`
				IsSameDay             bool `json:"is_same_day"`
				ItemDeliveredDeadline struct {
					Time  time.Time `json:"Time"`
					Valid bool      `json:"Valid"`
				} `json:"item_delivered_deadline"`
				OrderID         int       `json:"order_id"`
				PaymentDateTime time.Time `json:"payment_date_time"`
				UpdateTime      struct {
					Time  time.Time `json:"Time"`
					Valid bool      `json:"Valid"`
				} `json:"update_time"`
			} `json:"order_shipment_fulfillment"`
			BookingData struct {
				OrderID       int    `json:"order_id"`
				BookingCode   string `json:"booking_code"`
				BookingStatus int    `json:"booking_status"`
			} `json:"booking_data"`
		} `json:"order_data"`
		NextOrderID  int `json:"next_order_id"`
		FirstOrderID int `json:"first_order_id"`
	} `json:"data"`
}
