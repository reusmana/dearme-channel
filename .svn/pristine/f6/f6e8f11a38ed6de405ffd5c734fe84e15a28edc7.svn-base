package models

import "time"

//"github.com/google/uuid"

type TokenBukalapak struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int64  `json:"created_at"`
}

type ErrorBukalapak struct {
	Errors []struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"errors"`
}

type ProductsBukalapak struct {
	Data []struct {
		Active             bool          `json:"active"`
		Assurance          bool          `json:"assurance"`
		AvailableCountries []interface{} `json:"available_countries"`
		Category           struct {
			ID        int      `json:"id"`
			Name      string   `json:"name"`
			Structure []string `json:"structure"`
			URL       string   `json:"url"`
		} `json:"category"`
		Condition string        `json:"condition"`
		Couriers  []interface{} `json:"couriers"`
		CreatedAt time.Time     `json:"created_at"`
		Deal      struct {
		} `json:"deal"`
		DefaultCatalog       interface{}   `json:"default_catalog"`
		DefaultSkuID         int64         `json:"default_sku_id"`
		Description          string        `json:"description"`
		DescriptionBb        string        `json:"description_bb"`
		DigitalProduct       bool          `json:"digital_product"`
		Dimensions           interface{}   `json:"dimensions"`
		DiscountPercentage   int           `json:"discount_percentage"`
		DiscountSubsidy      interface{}   `json:"discount_subsidy"`
		DynamicTags          []interface{} `json:"dynamic_tags"`
		ForSale              bool          `json:"for_sale"`
		FreeShippingCoverage []interface{} `json:"free_shipping_coverage"`
		ID                   string        `json:"id"`
		Images               struct {
			Ids          []int64  `json:"ids"`
			LargeUrls    []string `json:"large_urls"`
			OriginalUrls []string `json:"original_urls"`
			SmallUrls    []string `json:"small_urls"`
		} `json:"images"`
		Imported                    bool          `json:"imported"`
		Installments                []interface{} `json:"installments"`
		InternationalCouriers       []interface{} `json:"international_couriers"`
		InternationalShippingStatus string        `json:"international_shipping_status"`
		Labels                      []interface{} `json:"labels"`
		MaxQuantity                 int64         `json:"max_quantity"`
		MerchantReturnInsurance     bool          `json:"merchant_return_insurance"`
		MinQuantity                 int           `json:"min_quantity"`
		Name                        string        `json:"name"`
		OriginalPrice               int           `json:"original_price"`
		Price                       int           `json:"price"`
		ProductSin                  []interface{} `json:"product_sin"`
		PromotedDetail              struct {
			Active    bool   `json:"active"`
			BidValue  int    `json:"bid_value"`
			EndDate   string `json:"end_date"`
			StartDate string `json:"start_date"`
		} `json:"promoted_detail"`
		Rating struct {
		} `json:"rating"`
		RelistedAt   time.Time `json:"relisted_at"`
		RushDelivery bool      `json:"rush_delivery"`
		Shipping     struct {
			ForceInsurance       bool          `json:"force_insurance"`
			FreeShippingCoverage []interface{} `json:"free_shipping_coverage"`
		} `json:"shipping"`
		SkuID   int64  `json:"sku_id"`
		SkuName string `json:"sku_name"`
		SLA     struct {
			Type  interface{} `json:"type"`
			Value interface{} `json:"value"`
		} `json:"sla"`
		SpecialCampaignID int `json:"special_campaign_id"`
		Specs             struct {
			Brand string `json:"brand"`
		} `json:"specs,omitempty"`
		State            string        `json:"state"`
		StateDescription []interface{} `json:"state_description"`
		Stats            struct {
			InterestCount       int `json:"interest_count"`
			SoldCount           int `json:"sold_count"`
			ViewCount           int `json:"view_count"`
			WaitingPaymentCount int `json:"waiting_payment_count"`
		} `json:"stats"`
		Stock int `json:"stock"`
		Store struct {
			Acceptance struct {
				AcceptanceRate      int `json:"acceptance_rate"`
				AcceptedTransaction int `json:"accepted_transaction"`
			} `json:"acceptance"`
			Address struct {
				City     string `json:"city"`
				Province string `json:"province"`
			} `json:"address"`
			Alert       string   `json:"alert"`
			AvatarURL   string   `json:"avatar_url"`
			BrandSeller bool     `json:"brand_seller"`
			Carriers    []string `json:"carriers"`
			Closing     struct {
				Closed       bool        `json:"closed"`
				ClosedReason interface{} `json:"closed_reason"`
				ReopenDate   interface{} `json:"reopen_date"`
				StartDate    interface{} `json:"start_date"`
			} `json:"closing"`
			ConnectedFacebook    bool          `json:"connected_facebook"`
			ConnectedTwitter     bool          `json:"connected_twitter"`
			DeliveryTime         string        `json:"delivery_time"`
			Description          string        `json:"description"`
			FirstUploadProductAt time.Time     `json:"first_upload_product_at"`
			Flagship             bool          `json:"flagship"`
			Groups               []interface{} `json:"groups"`
			HeaderImage          struct {
				ID  int    `json:"id"`
				URL string `json:"url"`
			} `json:"header_image"`
			ID         int `json:"id"`
			Inactivity struct {
				Inactive     bool      `json:"inactive"`
				LastAppearAt time.Time `json:"last_appear_at"`
			} `json:"inactivity"`
			InternationalShipping struct {
				Status    string `json:"status"`
				TncStatus string `json:"tnc_status"`
			} `json:"international_shipping"`
			LapakPhone        string `json:"lapak_phone"`
			LastOrderSchedule struct {
			} `json:"last_order_schedule"`
			Level struct {
				ImageURL string `json:"image_url"`
				Name     string `json:"name"`
			} `json:"level"`
			MaxDealDuration        int    `json:"max_deal_duration"`
			Name                   string `json:"name"`
			PremiumLevel           string `json:"premium_level"`
			PremiumTopSeller       bool   `json:"premium_top_seller"`
			ProductUploadRemaining struct {
				Daily   int `json:"daily"`
				Monthly int `json:"monthly"`
				Weekly  int `json:"weekly"`
			} `json:"product_upload_remaining"`
			Rejection struct {
				RecentTransactions int `json:"recent_transactions"`
				Rejected           int `json:"rejected"`
			} `json:"rejection"`
			Reputation struct {
				LevelProgress        int    `json:"level_progress"`
				NextLevel            string `json:"next_level"`
				PointNeededNextLevel int    `json:"point_needed_next_level"`
			} `json:"reputation"`
			Reviews struct {
				Negative int `json:"negative"`
				Positive int `json:"positive"`
			} `json:"reviews"`
			SLA struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"sla"`
			SubscriberAmount     int    `json:"subscriber_amount"`
			TermAndCondition     string `json:"term_and_condition"`
			URL                  string `json:"url"`
			UserTermAndCondition struct {
				Description interface{} `json:"description"`
				UpdatedAt   time.Time   `json:"updated_at"`
			} `json:"user_term_and_condition"`
		} `json:"store"`
		TagPages  []interface{} `json:"tag_pages"`
		UpdatedAt time.Time     `json:"updated_at"`
		URL       string        `json:"url"`
		Variants  []struct {
			Address struct {
				ID    int    `json:"id"`
				Title string `json:"title"`
			} `json:"address"`
			Deal struct {
			} `json:"deal"`
			DefaultSku bool `json:"default_sku"`
			Details    []struct {
				Label struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					State string `json:"state"`
				} `json:"label"`
				Value struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					State string `json:"state"`
				} `json:"value"`
				VariantID int `json:"variant_id"`
			} `json:"details"`
			Discount           int         `json:"discount"`
			DiscountPercentage int         `json:"discount_percentage"`
			DiscountSubsidy    interface{} `json:"discount_subsidy"`
			ID                 int64       `json:"id"`
			Images             struct {
				Ids       []int64  `json:"ids"`
				LargeUrls []string `json:"large_urls"`
				SmallUrls []string `json:"small_urls"`
			} `json:"images"`
			Price           int           `json:"price"`
			ProductID       string        `json:"product_id"`
			SkuName         string        `json:"sku_name"`
			State           string        `json:"state"`
			Stock           int           `json:"stock"`
			VariantImageIds []int64       `json:"variant_image_ids"`
			VariantName     string        `json:"variant_name"`
			Wholesales      []interface{} `json:"wholesales"`
		} `json:"variants"`
		VideoURL interface{} `json:"video_url"`
		Warranty struct {
			Cheapest bool `json:"cheapest"`
		} `json:"warranty"`
		Weight          int           `json:"weight"`
		Wholesales      []interface{} `json:"wholesales"`
		WithoutShipping bool          `json:"without_shipping"`
		Specs0          struct {
			Brand  string        `json:"brand"`
			Ukuran []interface{} `json:"ukuran"`
		} `json:"specs,omitempty"`
	} `json:"data"`
	Meta struct {
		HTTPStatus int `json:"http_status"`
		Offset     int `json:"offset"`
		Total      int `json:"total"`
		Limit      int `json:"limit"`
	} `json:"meta"`
}

type ProductDetailBukalapak struct {
	Data struct {
		Active             bool          `json:"active"`
		Assurance          bool          `json:"assurance"`
		AvailableCountries []interface{} `json:"available_countries"`
		Category           struct {
			ID        int      `json:"id"`
			Name      string   `json:"name"`
			Structure []string `json:"structure"`
			URL       string   `json:"url"`
		} `json:"category"`
		Condition string        `json:"condition"`
		Couriers  []interface{} `json:"couriers"`
		CreatedAt time.Time     `json:"created_at"`
		Deal      struct {
		} `json:"deal"`
		DefaultCatalog     interface{}   `json:"default_catalog"`
		Description        string        `json:"description"`
		DigitalProduct     bool          `json:"digital_product"`
		DiscountPercentage int           `json:"discount_percentage"`
		DiscountSubsidy    interface{}   `json:"discount_subsidy"`
		DynamicTags        []interface{} `json:"dynamic_tags"`
		ForSale            bool          `json:"for_sale"`
		ID                 string        `json:"id"`
		Images             struct {
			LargeUrls []string `json:"large_urls"`
			SmallUrls []string `json:"small_urls"`
		} `json:"images"`
		Imported                bool          `json:"imported"`
		Installments            []interface{} `json:"installments"`
		InternationalCouriers   []interface{} `json:"international_couriers"`
		MaxQuantity             int64         `json:"max_quantity"`
		MerchantReturnInsurance bool          `json:"merchant_return_insurance"`
		MinQuantity             int           `json:"min_quantity"`
		Name                    string        `json:"name"`
		OriginalPrice           int           `json:"original_price"`
		Price                   int           `json:"price"`
		ProductSin              []interface{} `json:"product_sin"`
		Rating                  struct {
			AverageRate int `json:"average_rate"`
			UserCount   int `json:"user_count"`
		} `json:"rating"`
		RelistedAt   time.Time `json:"relisted_at"`
		RushDelivery bool      `json:"rush_delivery"`
		Shipping     struct {
			ForceInsurance       bool          `json:"force_insurance"`
			FreeShippingCoverage []interface{} `json:"free_shipping_coverage"`
		} `json:"shipping"`
		SkuID int64 `json:"sku_id"`
		SLA   struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"sla"`
		SpecialCampaignID int `json:"special_campaign_id"`
		Specs             struct {
			Brand string `json:"brand"`
		} `json:"specs"`
		State            string        `json:"state"`
		StateDescription []interface{} `json:"state_description"`
		Stats            struct {
			InterestCount       int `json:"interest_count"`
			SoldCount           int `json:"sold_count"`
			ViewCount           int `json:"view_count"`
			WaitingPaymentCount int `json:"waiting_payment_count"`
		} `json:"stats"`
		Stock int `json:"stock"`
		Store struct {
			Address struct {
				City     string `json:"city"`
				Province string `json:"province"`
			} `json:"address"`
			Alert       string   `json:"alert"`
			AvatarURL   string   `json:"avatar_url"`
			BrandSeller bool     `json:"brand_seller"`
			Carriers    []string `json:"carriers"`
			Closing     struct {
				Closed bool `json:"closed"`
			} `json:"closing"`
			DeliveryTime         string        `json:"delivery_time"`
			Description          string        `json:"description"`
			FirstUploadProductAt time.Time     `json:"first_upload_product_at"`
			Flagship             bool          `json:"flagship"`
			Groups               []interface{} `json:"groups"`
			HeaderImage          struct {
				URL string `json:"url"`
			} `json:"header_image"`
			ID         int `json:"id"`
			Inactivity struct {
				Inactive     bool      `json:"inactive"`
				LastAppearAt time.Time `json:"last_appear_at"`
			} `json:"inactivity"`
			LastOrderSchedule struct {
			} `json:"last_order_schedule"`
			Level struct {
				ImageURL string `json:"image_url"`
				Name     string `json:"name"`
			} `json:"level"`
			Name             string `json:"name"`
			PremiumLevel     string `json:"premium_level"`
			PremiumTopSeller bool   `json:"premium_top_seller"`
			Rejection        struct {
				RecentTransactions int `json:"recent_transactions"`
				Rejected           int `json:"rejected"`
			} `json:"rejection"`
			Reviews struct {
				Negative int `json:"negative"`
				Positive int `json:"positive"`
			} `json:"reviews"`
			SLA struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"sla"`
			SubscriberAmount int    `json:"subscriber_amount"`
			TermAndCondition string `json:"term_and_condition"`
			URL              string `json:"url"`
		} `json:"store"`
		TagPages  []interface{} `json:"tag_pages"`
		UpdatedAt time.Time     `json:"updated_at"`
		URL       string        `json:"url"`
		VideoURL  string        `json:"video_url"`
		Warranty  struct {
			Cheapest bool `json:"cheapest"`
		} `json:"warranty"`
		Weight          int           `json:"weight"`
		Wholesales      []interface{} `json:"wholesales"`
		WithoutShipping bool          `json:"without_shipping"`
	} `json:"data"`
	Meta struct {
		HTTPStatus int `json:"http_status"`
	} `json:"meta"`
}

type ProductDetailSkuBukalapak struct {
	Data struct {
		ID        int64  `json:"id"`
		ProductID string `json:"product_id"`
		SkuName   string `json:"sku_name"`
		Stock     int    `json:"stock"`
		Price     int    `json:"price"`
		Images    struct {
			Ids       []int64  `json:"ids"`
			LargeUrls []string `json:"large_urls"`
			SmallUrls []string `json:"small_urls"`
		} `json:"images"`
		VariantName string        `json:"variant_name"`
		State       string        `json:"state"`
		Discount    int           `json:"discount"`
		Wholesales  []interface{} `json:"wholesales"`
		Address     struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"address"`
		DefaultSku bool `json:"default_sku"`
		Details    []struct {
			Label struct {
				ID    int    `json:"id"`
				State string `json:"state"`
				Name  string `json:"name"`
			} `json:"label"`
			Value struct {
				ID    int    `json:"id"`
				State string `json:"state"`
				Name  string `json:"name"`
			} `json:"value"`
			VariantID int `json:"variant_id"`
		} `json:"details"`
		Deal struct {
		} `json:"deal"`
		DiscountSubsidy    interface{} `json:"discount_subsidy"`
		DiscountPercentage int         `json:"discount_percentage"`
		VariantImageIds    []int64     `json:"variant_image_ids"`
	} `json:"data"`
	Meta struct {
		HTTPStatus int `json:"http_status"`
	} `json:"meta"`
}

type UpdateStockBukalapak struct {
	Data struct {
		ID        int64  `json:"id"`
		ProductID string `json:"product_id"`
		SkuName   string `json:"sku_name"`
		Stock     int    `json:"stock"`
		Price     int    `json:"price"`
		Images    struct {
			Ids       []int64  `json:"ids"`
			LargeUrls []string `json:"large_urls"`
			SmallUrls []string `json:"small_urls"`
		} `json:"images"`
		VariantName string        `json:"variant_name"`
		State       string        `json:"state"`
		Discount    int           `json:"discount"`
		Wholesales  []interface{} `json:"wholesales"`
		Address     struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"address"`
		DefaultSku bool `json:"default_sku"`
		Details    []struct {
			Label struct {
				ID    int    `json:"id"`
				State string `json:"state"`
				Name  string `json:"name"`
			} `json:"label"`
			Value struct {
				ID    int    `json:"id"`
				State string `json:"state"`
				Name  string `json:"name"`
			} `json:"value"`
			VariantID int `json:"variant_id"`
		} `json:"details"`
		Deal struct {
		} `json:"deal"`
		DiscountSubsidy    interface{} `json:"discount_subsidy"`
		DiscountPercentage int         `json:"discount_percentage"`
		VariantImageIds    []int64     `json:"variant_image_ids"`
	} `json:"data"`
	Meta struct {
		HTTPStatus int `json:"http_status"`
	} `json:"meta"`
}

type DetailStoreBukalapak struct {
	Data struct {
		ID         int         `json:"id"`
		Username   string      `json:"username"`
		AgentID    int         `json:"agent_id"`
		O2OAgent   interface{} `json:"o2o_agent"`
		Name       string      `json:"name"`
		Registered bool        `json:"registered"`
		Official   bool        `json:"official"`
		Verified   bool        `json:"verified"`
		Avatar     struct {
			ID  interface{} `json:"id"`
			URL string      `json:"url"`
		} `json:"avatar"`
		JoinedAt    time.Time `json:"joined_at"`
		LastLoginAt time.Time `json:"last_login_at"`
		Address     struct {
			Province   string  `json:"province"`
			City       string  `json:"city"`
			District   string  `json:"district"`
			Address    string  `json:"address"`
			PostalCode string  `json:"postal_code"`
			Latitude   float64 `json:"latitude"`
			Longitude  float64 `json:"longitude"`
		} `json:"address"`
		Unfreezing struct {
			Counter         int         `json:"counter"`
			Eligible        bool        `json:"eligible"`
			FreezedUntil    interface{} `json:"freezed_until"`
			PermanentFrozen bool        `json:"permanent_frozen"`
			Frozen          bool        `json:"frozen"`
			FreezeCategory  string      `json:"freeze_category"`
		} `json:"unfreezing"`
		TopUser                     bool          `json:"top_user"`
		Email                       string        `json:"email"`
		Gender                      string        `json:"gender"`
		BirthDate                   string        `json:"birth_date"`
		Confirmed                   bool          `json:"confirmed"`
		Phone                       string        `json:"phone"`
		PhoneConfirmed              bool          `json:"phone_confirmed"`
		TfaStatus                   string        `json:"tfa_status"`
		BullionAutoInvestmentStatus string        `json:"bullion_auto_investment_status"`
		Role                        string        `json:"role"`
		Banks                       []interface{} `json:"banks"`
		FavoritePaymentTypes        []interface{} `json:"favorite_payment_types"`
		WalletState                 string        `json:"wallet_state"`
		PriorityBuyerPackageType    interface{}   `json:"priority_buyer_package_type"`
		BlacklistedPromo            bool          `json:"blacklisted_promo"`
		LastOtp                     string        `json:"last_otp"`
	} `json:"data"`
	Meta struct {
		HTTPStatus int `json:"http_status"`
	} `json:"meta"`
}

type OrdersBukalapakOLD struct {
	Data []struct {
		ID             int64     `json:"id"`
		Type           string    `json:"type"`
		InvoiceID      int64     `json:"invoice_id"`
		TransactionID  string    `json:"transaction_id"`
		State          string    `json:"state"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		StateChangedAt struct {
			PendingAt              time.Time `json:"pending_at"`
			PaidAt                 time.Time `json:"paid_at"`
			AcceptedAt             time.Time `json:"accepted_at"`
			RejectedAt             string    `json:"rejected_at"`
			RefundedAt             string    `json:"refunded_at"`
			RemittedAt             string    `json:"remitted_at"`
			CancelledAt            string    `json:"cancelled_at"`
			DeliveredAt            string    `json:"delivered_at"`
			ReceivedAt             string    `json:"received_at"`
			ExpiredAt              time.Time `json:"expired_at"`
			RefundAt               time.Time `json:"refund_at"`
			RefundByAcceptAt       time.Time `json:"refund_by_accept_at"`
			RemitAt                string    `json:"remit_at"`
			RefundConfirmInvalidAt string    `json:"refund_confirm_invalid_at"`
		} `json:"state_changed_at"`
		StateChangedBy struct {
			CancelledBy interface{} `json:"cancelled_by"`
		} `json:"state_changed_by"`
		Actionable bool   `json:"actionable"`
		CreatedOn  string `json:"created_on"`
		Items      []struct {
			ID              int64  `json:"id"`
			Name            string `json:"name"`
			Price           int    `json:"price"`
			Quantity        int    `json:"quantity"`
			AgentCommission int    `json:"agent_commission"`
			Stuff           struct {
				ReferenceType string `json:"reference_type"`
				ID            int64  `json:"id"`
				Product       struct {
					ID          string `json:"id"`
					Price       int    `json:"price"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Condition   string `json:"condition"`
					Weight      int    `json:"weight"`
					URL         string `json:"url"`
					Shipping    struct {
						ForceInsurance       bool          `json:"force_insurance"`
						FreeShippingCoverage []interface{} `json:"free_shipping_coverage"`
					} `json:"shipping"`
					Assurance interface{} `json:"assurance"`
				} `json:"product"`
				Store struct {
					ID                        int       `json:"id"`
					Name                      string    `json:"name"`
					TermAndCondition          string    `json:"term_and_condition"`
					TermAndConditionUpdatedAt time.Time `json:"term_and_condition_updated_at"`
					Address                   struct {
						City string `json:"city"`
					} `json:"address"`
				} `json:"store"`
				Price int `json:"price"`
				Image struct {
					LargeUrls []string `json:"large_urls"`
					SmallUrls []string `json:"small_urls"`
				} `json:"image"`
				VariantName string `json:"variant_name"`
				SkuName     string `json:"sku_name"`
				Discount    int    `json:"discount"`
			} `json:"stuff"`
			TotalPrice int `json:"total_price"`
			Category   struct {
				Name string `json:"name"`
			} `json:"category"`
			FlashDealDiscount int `json:"flash_deal_discount"`
			Product           struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Price int    `json:"price"`
			} `json:"product"`
		} `json:"items"`
		Buyer struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
			Phone  string `json:"phone"`
		} `json:"buyer"`
		Store struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			Description interface{} `json:"description"`
			Avatar      string      `json:"avatar"`
			Brand       bool        `json:"brand"`
			Official    bool        `json:"official"`
			SuperSeller bool        `json:"super_seller"`
			City        string      `json:"city"`
		} `json:"store"`
		Amount struct {
			Buyer struct {
				Total         int `json:"total"`
				PaymentAmount int `json:"payment_amount"`
				CodedAmount   int `json:"coded_amount"`
				RefundAmount  int `json:"refund_amount"`
				Details       struct {
					Item                      int `json:"item"`
					Delivery                  int `json:"delivery"`
					Insurance                 int `json:"insurance"`
					AxinanInsuranceAmount     int `json:"axinan_insurance_amount"`
					GadgetInsuranceAmount     int `json:"gadget_insurance_amount"`
					GoodsInsuranceAmount      int `json:"goods_insurance_amount"`
					CosmeticInsuranceAmount   int `json:"cosmetic_insurance_amount"`
					FmcgInsuranceAmount       int `json:"fmcg_insurance_amount"`
					ReturnInsuranceAmount     int `json:"return_insurance_amount"`
					GamersInsuranceAmount     int `json:"gamers_insurance_amount"`
					Administration            int `json:"administration"`
					TippingAmount             int `json:"tipping_amount"`
					Negotiation               int `json:"negotiation"`
					Vat                       int `json:"vat"`
					FlashDealDiscount         int `json:"flash_deal_discount"`
					PriorityBuyer             int `json:"priority_buyer"`
					VoucherDiscount           int `json:"voucher_discount"`
					RetargetDiscountAmount    int `json:"retarget_discount_amount"`
					ShippingFeeDiscountAmount int `json:"shipping_fee_discount_amount"`
				} `json:"details"`
			} `json:"buyer"`
			Seller struct {
				Total            int `json:"total"`
				GrossRemitAmount int `json:"gross_remit_amount"`
				Details          struct {
					Items              int           `json:"items"`
					Delivery           int           `json:"delivery"`
					Insurance          int           `json:"insurance"`
					ShippingReductions []interface{} `json:"shipping_reductions"`
					Reductions         []interface{} `json:"reductions"`
					SuperSeller        []interface{} `json:"super_seller"`
				} `json:"details"`
			} `json:"seller"`
		} `json:"amount"`
		Cashback []interface{} `json:"cashback"`
		Delivery struct {
			ID        interface{} `json:"id"`
			Consignee struct {
				Name       string  `json:"name"`
				Phone      string  `json:"phone"`
				Country    string  `json:"country"`
				Province   string  `json:"province"`
				City       string  `json:"city"`
				District   string  `json:"district"`
				Address    string  `json:"address"`
				PostalCode string  `json:"postal_code"`
				Longitude  float64 `json:"longitude"`
				Latitude   float64 `json:"latitude"`
			} `json:"consignee"`
			RequestedCarrier          string        `json:"requested_carrier"`
			TrackingNumber            string        `json:"tracking_number"`
			Carrier                   string        `json:"carrier"`
			WhiteLabelCourier         bool          `json:"white_label_courier"`
			ForceAwb                  bool          `json:"force_awb"`
			Booking                   interface{}   `json:"booking"`
			History                   []interface{} `json:"history"`
			ShippingReceiptState      interface{}   `json:"shipping_receipt_state"`
			ForceFindDriver           bool          `json:"force_find_driver"`
			AllowDifferentCourier     bool          `json:"allow_different_courier"`
			AllowManualReceiptVoucher bool          `json:"allow_manual_receipt_voucher"`
			ManualSwitchFee           int           `json:"manual_switch_fee"`
			AllowRedeliver            bool          `json:"allow_redeliver"`
			PickupTime                struct {
				From string `json:"from"`
				To   string `json:"to"`
			} `json:"pickup_time"`
			ForceAwbVoucher           bool        `json:"force_awb_voucher"`
			EstimatedReceivedAt       time.Time   `json:"estimated_received_at"`
			EstimatedShipmentDuration interface{} `json:"estimated_shipment_duration"`
			ConvenienceStore          interface{} `json:"convenience_store"`
			AvailableShippingService  interface{} `json:"available_shipping_service"`
			BuyerLogisticChoice       string      `json:"buyer_logistic_choice"`
			SellerLogisticChoice      interface{} `json:"seller_logistic_choice"`
			ReceiptValidity           interface{} `json:"receipt_validity"`
		} `json:"delivery"`
		Dropship struct {
			Note string `json:"note"`
		} `json:"dropship"`
		Feedback struct {
		} `json:"feedback"`
		Options struct {
			AdminNotes string `json:"admin_notes"`
		} `json:"options"`
		OnHold    bool        `json:"on_hold"`
		Deal      bool        `json:"deal"`
		ClaimID   interface{} `json:"claim_id"`
		Promotion struct {
			PromotedPush bool `json:"promoted_push"`
			Push         bool `json:"push"`
			Voucher      bool `json:"voucher"`
		} `json:"promotion"`
		SLA struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"sla"`
		LastPrintedAt                  interface{} `json:"last_printed_at"`
		VirtualTransactionSerialNumber interface{} `json:"virtual_transaction_serial_number"`
		CanClaimAssurance              bool        `json:"can_claim_assurance"`
		PaymentMethod                  string      `json:"payment_method"`
		MerchantReturnInsuranceID      interface{} `json:"merchant_return_insurance_id"`
		DigitalGoodsTransaction        bool        `json:"digital_goods_transaction"`
		GamersInsuranceID              interface{} `json:"gamers_insurance_id"`
	} `json:"data"`
	Meta struct {
		Total      int `json:"total"`
		HTTPStatus int `json:"http_status"`
		Limit      int `json:"limit"`
		Offset     int `json:"offset"`
	} `json:"meta"`
}

type OrdersBukalapak struct {
	Data []struct {
		ID             int64     `json:"id"`
		Type           string    `json:"type"`
		InvoiceID      int64     `json:"invoice_id"`
		TransactionID  string    `json:"transaction_id"`
		State          string    `json:"state"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		StateChangedAt struct {
			PendingAt              time.Time `json:"pending_at"`
			PaidAt                 time.Time `json:"paid_at"`
			AcceptedAt             string    `json:"accepted_at"`
			RejectedAt             string    `json:"rejected_at"`
			RefundedAt             string    `json:"refunded_at"`
			RemittedAt             string    `json:"remitted_at"`
			CancelledAt            string    `json:"cancelled_at"`
			DeliveredAt            string    `json:"delivered_at"`
			ReceivedAt             string    `json:"received_at"`
			ExpiredAt              time.Time `json:"expired_at"`
			RefundAt               time.Time `json:"refund_at"`
			RefundByAcceptAt       time.Time `json:"refund_by_accept_at"`
			RemitAt                string    `json:"remit_at"`
			RefundConfirmInvalidAt string    `json:"refund_confirm_invalid_at"`
		} `json:"state_changed_at"`
		StateChangedBy struct {
			CancelledBy interface{} `json:"cancelled_by"`
		} `json:"state_changed_by"`
		Actionable bool   `json:"actionable"`
		CreatedOn  string `json:"created_on"`
		Items      []struct {
			ID              int64  `json:"id"`
			Name            string `json:"name"`
			Price           int    `json:"price"`
			Quantity        int    `json:"quantity"`
			AgentCommission int    `json:"agent_commission"`
			Stuff           struct {
				ReferenceType string `json:"reference_type"`
				ID            int64  `json:"id"`
				Product       struct {
					ID          string `json:"id"`
					Price       int    `json:"price"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Condition   string `json:"condition"`
					Weight      int    `json:"weight"`
					URL         string `json:"url"`
					Shipping    struct {
						ForceInsurance       bool          `json:"force_insurance"`
						FreeShippingCoverage []interface{} `json:"free_shipping_coverage"`
					} `json:"shipping"`
					Assurance interface{} `json:"assurance"`
				} `json:"product"`
				Store struct {
					ID                        int       `json:"id"`
					Name                      string    `json:"name"`
					TermAndCondition          string    `json:"term_and_condition"`
					TermAndConditionUpdatedAt time.Time `json:"term_and_condition_updated_at"`
					Address                   struct {
						City string `json:"city"`
					} `json:"address"`
				} `json:"store"`
				Price int `json:"price"`
				Image struct {
					LargeUrls []string `json:"large_urls"`
					SmallUrls []string `json:"small_urls"`
				} `json:"image"`
				VariantName string `json:"variant_name"`
				SkuName     string `json:"sku_name"`
				Discount    int    `json:"discount"`
			} `json:"stuff"`
			TotalPrice int `json:"total_price"`
			Category   struct {
				Name string `json:"name"`
			} `json:"category"`
			FlashDealDiscount int `json:"flash_deal_discount"`
			Product           struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Price int    `json:"price"`
			} `json:"product"`
		} `json:"items"`
		Buyer struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
			Phone  string `json:"phone"`
		} `json:"buyer"`
		Store struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			Description interface{} `json:"description"`
			Avatar      string      `json:"avatar"`
			Brand       bool        `json:"brand"`
			Official    bool        `json:"official"`
			SuperSeller bool        `json:"super_seller"`
			City        string      `json:"city"`
		} `json:"store"`
		Amount struct {
			Buyer struct {
				Total         int `json:"total"`
				PaymentAmount int `json:"payment_amount"`
				CodedAmount   int `json:"coded_amount"`
				RefundAmount  int `json:"refund_amount"`
				Details       struct {
					Item                      int `json:"item"`
					Delivery                  int `json:"delivery"`
					Insurance                 int `json:"insurance"`
					AxinanInsuranceAmount     int `json:"axinan_insurance_amount"`
					GadgetInsuranceAmount     int `json:"gadget_insurance_amount"`
					GoodsInsuranceAmount      int `json:"goods_insurance_amount"`
					CosmeticInsuranceAmount   int `json:"cosmetic_insurance_amount"`
					FmcgInsuranceAmount       int `json:"fmcg_insurance_amount"`
					ReturnInsuranceAmount     int `json:"return_insurance_amount"`
					GamersInsuranceAmount     int `json:"gamers_insurance_amount"`
					Administration            int `json:"administration"`
					TippingAmount             int `json:"tipping_amount"`
					Negotiation               int `json:"negotiation"`
					Vat                       int `json:"vat"`
					FlashDealDiscount         int `json:"flash_deal_discount"`
					PriorityBuyer             int `json:"priority_buyer"`
					VoucherDiscount           int `json:"voucher_discount"`
					RetargetDiscountAmount    int `json:"retarget_discount_amount"`
					ShippingFeeDiscountAmount int `json:"shipping_fee_discount_amount"`
				} `json:"details"`
			} `json:"buyer"`
			Seller struct {
				Total            int `json:"total"`
				GrossRemitAmount int `json:"gross_remit_amount"`
				Details          struct {
					Items              int           `json:"items"`
					Delivery           int           `json:"delivery"`
					Insurance          int           `json:"insurance"`
					ShippingReductions []interface{} `json:"shipping_reductions"`
					Reductions         []interface{} `json:"reductions"`
					SuperSeller        []interface{} `json:"super_seller"`
				} `json:"details"`
			} `json:"seller"`
		} `json:"amount"`
		Cashback []interface{} `json:"cashback"`
		Delivery struct {
			ID        interface{} `json:"id"`
			Consignee struct {
				Name       string  `json:"name"`
				Phone      string  `json:"phone"`
				Country    string  `json:"country"`
				Province   string  `json:"province"`
				City       string  `json:"city"`
				District   string  `json:"district"`
				Address    string  `json:"address"`
				PostalCode string  `json:"postal_code"`
				Longitude  float64 `json:"longitude"`
				Latitude   float64 `json:"latitude"`
			} `json:"consignee"`
			RequestedCarrier          string        `json:"requested_carrier"`
			TrackingNumber            string        `json:"tracking_number"`
			Carrier                   string        `json:"carrier"`
			WhiteLabelCourier         bool          `json:"white_label_courier"`
			ForceAwb                  bool          `json:"force_awb"`
			Booking                   interface{}   `json:"booking"`
			History                   []interface{} `json:"history"`
			ShippingReceiptState      interface{}   `json:"shipping_receipt_state"`
			ForceFindDriver           bool          `json:"force_find_driver"`
			AllowDifferentCourier     bool          `json:"allow_different_courier"`
			AllowManualReceiptVoucher bool          `json:"allow_manual_receipt_voucher"`
			ManualSwitchFee           int           `json:"manual_switch_fee"`
			AllowRedeliver            bool          `json:"allow_redeliver"`
			PickupTime                struct {
				From string `json:"from"`
				To   string `json:"to"`
			} `json:"pickup_time"`
			ForceAwbVoucher           bool        `json:"force_awb_voucher"`
			EstimatedReceivedAt       time.Time   `json:"estimated_received_at"`
			EstimatedShipmentDuration interface{} `json:"estimated_shipment_duration"`
			ConvenienceStore          interface{} `json:"convenience_store"`
			AvailableShippingService  interface{} `json:"available_shipping_service"`
			BuyerLogisticChoice       string      `json:"buyer_logistic_choice"`
			SellerLogisticChoice      interface{} `json:"seller_logistic_choice"`
			ReceiptValidity           interface{} `json:"receipt_validity"`
		} `json:"delivery"`
		Dropship struct {
			Note string `json:"note"`
		} `json:"dropship"`
		Feedback struct {
		} `json:"feedback"`
		Options struct {
			AdminNotes string `json:"admin_notes"`
		} `json:"options"`
		OnHold    bool        `json:"on_hold"`
		Deal      bool        `json:"deal"`
		ClaimID   interface{} `json:"claim_id"`
		Promotion struct {
			PromotedPush bool `json:"promoted_push"`
			Push         bool `json:"push"`
			Voucher      bool `json:"voucher"`
		} `json:"promotion"`
		SLA struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"sla"`
		LastPrintedAt                  interface{} `json:"last_printed_at"`
		VirtualTransactionSerialNumber interface{} `json:"virtual_transaction_serial_number"`
		CanClaimAssurance              bool        `json:"can_claim_assurance"`
		PaymentMethod                  string      `json:"payment_method"`
		MerchantReturnInsuranceID      interface{} `json:"merchant_return_insurance_id"`
		DigitalGoodsTransaction        bool        `json:"digital_goods_transaction"`
		GamersInsuranceID              interface{} `json:"gamers_insurance_id"`
	} `json:"data"`
	Meta struct {
		Total      int `json:"total"`
		HTTPStatus int `json:"http_status"`
		Limit      int `json:"limit"`
		Offset     int `json:"offset"`
	} `json:"meta"`
}

type OrderDetailBukalapak struct {
	Data struct {
		ID             int64     `json:"id"`
		Type           string    `json:"type"`
		InvoiceID      int64     `json:"invoice_id"`
		TransactionID  string    `json:"transaction_id"`
		State          string    `json:"state"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		StateChangedAt struct {
			PendingAt              time.Time `json:"pending_at"`
			PaidAt                 time.Time `json:"paid_at"`
			AcceptedAt             string    `json:"accepted_at"`
			RejectedAt             string    `json:"rejected_at"`
			RefundedAt             string    `json:"refunded_at"`
			RemittedAt             string    `json:"remitted_at"`
			CancelledAt            string    `json:"cancelled_at"`
			DeliveredAt            string    `json:"delivered_at"`
			ReceivedAt             string    `json:"received_at"`
			ExpiredAt              time.Time `json:"expired_at"`
			RefundAt               time.Time `json:"refund_at"`
			RefundByAcceptAt       time.Time `json:"refund_by_accept_at"`
			RemitAt                string    `json:"remit_at"`
			RefundConfirmInvalidAt string    `json:"refund_confirm_invalid_at"`
		} `json:"state_changed_at"`
		StateChangedBy struct {
			CancelledBy interface{} `json:"cancelled_by"`
		} `json:"state_changed_by"`
		Actionable bool   `json:"actionable"`
		CreatedOn  string `json:"created_on"`
		Items      []struct {
			ID              int64  `json:"id"`
			Name            string `json:"name"`
			Price           int    `json:"price"`
			Quantity        int    `json:"quantity"`
			AgentCommission int    `json:"agent_commission"`
			Stuff           struct {
				ReferenceType string `json:"reference_type"`
				ID            int64  `json:"id"`
				Product       struct {
					ID          string `json:"id"`
					Price       int    `json:"price"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Condition   string `json:"condition"`
					Weight      int    `json:"weight"`
					URL         string `json:"url"`
					Shipping    struct {
						ForceInsurance       bool          `json:"force_insurance"`
						FreeShippingCoverage []interface{} `json:"free_shipping_coverage"`
					} `json:"shipping"`
					Assurance interface{} `json:"assurance"`
				} `json:"product"`
				Store struct {
					ID                        int       `json:"id"`
					Name                      string    `json:"name"`
					TermAndCondition          string    `json:"term_and_condition"`
					TermAndConditionUpdatedAt time.Time `json:"term_and_condition_updated_at"`
					Address                   struct {
						City string `json:"city"`
					} `json:"address"`
				} `json:"store"`
				Price int `json:"price"`
				Image struct {
					LargeUrls []string `json:"large_urls"`
					SmallUrls []string `json:"small_urls"`
				} `json:"image"`
				VariantName string `json:"variant_name"`
				SkuName     string `json:"sku_name"`
				Discount    int    `json:"discount"`
			} `json:"stuff"`
			TotalPrice int `json:"total_price"`
			Category   struct {
				Name string `json:"name"`
			} `json:"category"`
			FlashDealDiscount int `json:"flash_deal_discount"`
			Product           struct {
				ID    string `json:"id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				Price int    `json:"price"`
			} `json:"product"`
		} `json:"items"`
		Buyer struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
			Phone  string `json:"phone"`
		} `json:"buyer"`
		Store struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			Description interface{} `json:"description"`
			Avatar      string      `json:"avatar"`
			Brand       bool        `json:"brand"`
			Official    bool        `json:"official"`
			SuperSeller bool        `json:"super_seller"`
			City        string      `json:"city"`
		} `json:"store"`
		Amount struct {
			Buyer struct {
				Total         int `json:"total"`
				PaymentAmount int `json:"payment_amount"`
				CodedAmount   int `json:"coded_amount"`
				RefundAmount  int `json:"refund_amount"`
				Details       struct {
					Item                      int `json:"item"`
					Delivery                  int `json:"delivery"`
					Insurance                 int `json:"insurance"`
					AxinanInsuranceAmount     int `json:"axinan_insurance_amount"`
					GadgetInsuranceAmount     int `json:"gadget_insurance_amount"`
					GoodsInsuranceAmount      int `json:"goods_insurance_amount"`
					CosmeticInsuranceAmount   int `json:"cosmetic_insurance_amount"`
					FmcgInsuranceAmount       int `json:"fmcg_insurance_amount"`
					ReturnInsuranceAmount     int `json:"return_insurance_amount"`
					GamersInsuranceAmount     int `json:"gamers_insurance_amount"`
					Administration            int `json:"administration"`
					TippingAmount             int `json:"tipping_amount"`
					Negotiation               int `json:"negotiation"`
					Vat                       int `json:"vat"`
					FlashDealDiscount         int `json:"flash_deal_discount"`
					PriorityBuyer             int `json:"priority_buyer"`
					VoucherDiscount           int `json:"voucher_discount"`
					RetargetDiscountAmount    int `json:"retarget_discount_amount"`
					ShippingFeeDiscountAmount int `json:"shipping_fee_discount_amount"`
				} `json:"details"`
			} `json:"buyer"`
			Seller struct {
				Total            int `json:"total"`
				GrossRemitAmount int `json:"gross_remit_amount"`
				Details          struct {
					Items              int           `json:"items"`
					Delivery           int           `json:"delivery"`
					Insurance          int           `json:"insurance"`
					ShippingReductions []interface{} `json:"shipping_reductions"`
					Reductions         []interface{} `json:"reductions"`
					SuperSeller        []interface{} `json:"super_seller"`
				} `json:"details"`
			} `json:"seller"`
		} `json:"amount"`
		Cashback []interface{} `json:"cashback"`
		Delivery struct {
			ID        interface{} `json:"id"`
			Consignee struct {
				Name       string  `json:"name"`
				Phone      string  `json:"phone"`
				Country    string  `json:"country"`
				Province   string  `json:"province"`
				City       string  `json:"city"`
				District   string  `json:"district"`
				Address    string  `json:"address"`
				PostalCode string  `json:"postal_code"`
				Longitude  float64 `json:"longitude"`
				Latitude   float64 `json:"latitude"`
			} `json:"consignee"`
			RequestedCarrier          string        `json:"requested_carrier"`
			TrackingNumber            string        `json:"tracking_number"`
			Carrier                   string        `json:"carrier"`
			WhiteLabelCourier         bool          `json:"white_label_courier"`
			ForceAwb                  bool          `json:"force_awb"`
			Booking                   interface{}   `json:"booking"`
			History                   []interface{} `json:"history"`
			ShippingReceiptState      interface{}   `json:"shipping_receipt_state"`
			ForceFindDriver           bool          `json:"force_find_driver"`
			AllowDifferentCourier     bool          `json:"allow_different_courier"`
			AllowManualReceiptVoucher bool          `json:"allow_manual_receipt_voucher"`
			ManualSwitchFee           int           `json:"manual_switch_fee"`
			AllowRedeliver            bool          `json:"allow_redeliver"`
			PickupTime                struct {
				From string `json:"from"`
				To   string `json:"to"`
			} `json:"pickup_time"`
			ForceAwbVoucher           bool        `json:"force_awb_voucher"`
			EstimatedReceivedAt       time.Time   `json:"estimated_received_at"`
			EstimatedShipmentDuration interface{} `json:"estimated_shipment_duration"`
			ConvenienceStore          interface{} `json:"convenience_store"`
			AvailableShippingService  struct {
				Type string `json:"type"`
			} `json:"available_shipping_service"`
			BuyerLogisticChoice  string      `json:"buyer_logistic_choice"`
			SellerLogisticChoice interface{} `json:"seller_logistic_choice"`
			ReceiptValidity      interface{} `json:"receipt_validity"`
		} `json:"delivery"`
		Dropship struct {
			Note string `json:"note"`
		} `json:"dropship"`
		Feedback struct {
		} `json:"feedback"`
		Options struct {
			AdminNotes string `json:"admin_notes"`
		} `json:"options"`
		OnHold    bool        `json:"on_hold"`
		Deal      bool        `json:"deal"`
		ClaimID   interface{} `json:"claim_id"`
		Promotion struct {
			PromotedPush bool `json:"promoted_push"`
			Push         bool `json:"push"`
			Voucher      bool `json:"voucher"`
		} `json:"promotion"`
		SLA struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"sla"`
		LastPrintedAt                  interface{} `json:"last_printed_at"`
		VirtualTransactionSerialNumber interface{} `json:"virtual_transaction_serial_number"`
		CanClaimAssurance              bool        `json:"can_claim_assurance"`
		PaymentMethod                  string      `json:"payment_method"`
		MerchantReturnInsuranceID      interface{} `json:"merchant_return_insurance_id"`
		DigitalGoodsTransaction        bool        `json:"digital_goods_transaction"`
		GamersInsuranceID              interface{} `json:"gamers_insurance_id"`
	} `json:"data"`
	Meta struct {
		HTTPStatus int `json:"http_status"`
	} `json:"meta"`
}

type BookingBukalapak struct {
	Data struct {
		ID                 int         `json:"id"`
		TransactionID      int64       `json:"transaction_id"`
		ShippingID         int         `json:"shipping_id"`
		Courier            string      `json:"courier"`
		CourierName        string      `json:"courier_name"`
		BookingCode        string      `json:"booking_code"`
		State              string      `json:"state"`
		ShippingFee        int         `json:"shipping_fee"`
		InsuranceFee       int         `json:"insurance_fee"`
		Weight             interface{} `json:"weight"`
		UsedAt             string      `json:"used_at"`
		Invoicing          bool        `json:"invoicing"`
		PartnerLabelURL    string      `json:"partner_label_url"`
		CourierRoutingCode string      `json:"courier_routing_code"`
		CreatedAt          time.Time   `json:"created_at"`
	} `json:"data"`
	Meta struct {
		HTTPStatus int `json:"http_status"`
	} `json:"meta"`
}

type DetailKurirBukalapak struct {
	Data []struct {
		Carrier      string `json:"carrier"`
		Express      bool   `json:"express"`
		CourierGroup string `json:"courier_group"`
	} `json:"data"`
	Meta struct {
		HTTPStatus int `json:"http_status"`
	} `json:"meta"`
}
