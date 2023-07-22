package models

//"github.com/google/uuid"
import (
	"time"
)

type DetailEscrowsShopee struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Response  struct {
		BuyerUserName string `json:"buyer_user_name"`
		OrderIncome   struct {
			ActualShippingFee                        int `json:"actual_shipping_fee"`
			BuyerPaidShippingFee                     int `json:"buyer_paid_shipping_fee"`
			BuyerTotalAmount                         int `json:"buyer_total_amount"`
			BuyerTransactionFee                      int `json:"buyer_transaction_fee"`
			CampaignFee                              int `json:"campaign_fee"`
			Coins                                    int `json:"coins"`
			CommissionFee                            int `json:"commission_fee"`
			CostOfGoodsSold                          int `json:"cost_of_goods_sold"`
			CreditCardPromotion                      int `json:"credit_card_promotion"`
			CreditCardTransactionFee                 int `json:"credit_card_transaction_fee"`
			CrossBorderTax                           int `json:"cross_border_tax"`
			DeliverySellerProtectionFeePremiumAmount int `json:"delivery_seller_protection_fee_premium_amount"`
			DrcAdjustableRefund                      int `json:"drc_adjustable_refund"`
			EscrowAmount                             int `json:"escrow_amount"`
			EscrowTax                                int `json:"escrow_tax"`
			EstimatedShippingFee                     int `json:"estimated_shipping_fee"`
			FinalEscrowProductGst                    int `json:"final_escrow_product_gst"`
			FinalEscrowShippingGst                   int `json:"final_escrow_shipping_gst"`
			FinalProductProtection                   int `json:"final_product_protection"`
			FinalProductVatTax                       int `json:"final_product_vat_tax"`
			FinalReturnToSellerShippingFee           int `json:"final_return_to_seller_shipping_fee"`
			FinalShippingFee                         int `json:"final_shipping_fee"`
			FinalShippingVatTax                      int `json:"final_shipping_vat_tax"`
			Items                                    []struct {
				ActivityID                int     `json:"activity_id"`
				ActivityType              string  `json:"activity_type"`
				DiscountFromCoin          int     `json:"discount_from_coin"`
				DiscountFromVoucherSeller int     `json:"discount_from_voucher_seller"`
				DiscountFromVoucherShopee int     `json:"discount_from_voucher_shopee"`
				DiscountedPrice           float64 `json:"discounted_price"`
				IsB2CShopItem             bool    `json:"is_b2c_shop_item"`
				IsMainItem                bool    `json:"is_main_item"`
				ItemID                    int     `json:"item_id"`
				ItemName                  string  `json:"item_name"`
				ItemSku                   string  `json:"item_sku"`
				ModelID                   int64   `json:"model_id"`
				ModelName                 string  `json:"model_name"`
				ModelSku                  string  `json:"model_sku"`
				OriginalPrice             int     `json:"original_price"`
				QuantityPurchased         int     `json:"quantity_purchased"`
				SellerDiscount            int     `json:"seller_discount"`
				ShopeeDiscount            int     `json:"shopee_discount"`
			} `json:"items"`
			OrderChargeableWeight               int      `json:"order_chargeable_weight"`
			OriginalCostOfGoodsSold             int      `json:"original_cost_of_goods_sold"`
			OriginalPrice                       int      `json:"original_price"`
			OriginalShopeeDiscount              int      `json:"original_shopee_discount"`
			PaymentPromotion                    int      `json:"payment_promotion"`
			ReverseShippingFee                  int      `json:"reverse_shipping_fee"`
			RsfSellerProtectionFeeClaimAmount   int      `json:"rsf_seller_protection_fee_claim_amount"`
			RsfSellerProtectionFeePremiumAmount int      `json:"rsf_seller_protection_fee_premium_amount"`
			SellerCoinCashBack                  int      `json:"seller_coin_cash_back"`
			SellerDiscount                      int      `json:"seller_discount"`
			SellerLostCompensation              int      `json:"seller_lost_compensation"`
			SellerReturnRefund                  int      `json:"seller_return_refund"`
			SellerShippingDiscount              int      `json:"seller_shipping_discount"`
			SellerTransactionFee                int      `json:"seller_transaction_fee"`
			SellerVoucherCode                   []string `json:"seller_voucher_code"`
			ServiceFee                          int      `json:"service_fee"`
			ShippingFeeDiscountFrom3Pl          int      `json:"shipping_fee_discount_from_3pl"`
			ShopeeDiscount                      int      `json:"shopee_discount"`
			ShopeeShippingRebate                int      `json:"shopee_shipping_rebate"`
			VoucherFromSeller                   int      `json:"voucher_from_seller"`
			VoucherFromShopee                   int      `json:"voucher_from_shopee"`
		} `json:"order_income"`
		OrderSn           string   `json:"order_sn"`
		ReturnOrderSnList []string `json:"return_order_sn_list"`
	} `json:"response"`
}

type ListProductBYSkuShopee struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	Warning   string `json:"warning"`
	RequestID string `json:"request_id"`
	Response  struct {
		ItemIDList []int64 `json:"item_id_list"`
		TotalCount int     `json:"total_count"`
		NextOffset string  `json:"next_offset"`
	} `json:"response"`
}

type ListOrderShopee struct {
	Error         string        `json:"error"`
	Message       string        `json:"message"`
	ResponseOrder ResponseOrder `json:"response"`
	RequestId     string        `json:"request_id"`
}
type ResponseOrder struct {
	More       bool          `json:"more"`
	NextCursor string        `json:"next_cursor"`
	OrderList  []ListOrderSn `json:"order_list"`
}
type ListOrderSn struct {
	OrderSn string `json:"order_sn"`
}

type Header struct {
	Error     string   `json:"error"`
	Message   string   `json:"message"`
	Response  Response `json:"response"`
	RequestId string   `json:"request_id"`
}
type Response struct {
	OrderList []OrderList `json:"order_list"`
}

type OrderList struct {
	BuyerUsername    string           `json:"buyer_username"`
	Cod              bool             `json:"cod"`
	CreateTime       int64            `json:"create_time"`
	UpdateTime       int64            `json:"update_time"`
	DaysToShip       int64            `json:"days_to_ship"`
	MessageToSeller  string           `json:"message_to_seller"`
	OrderSn          string           `json:"order_sn"`
	OrderStatus      string           `json:"order_status"`
	ItemList         []ItemList       `json:"item_list"`
	RecipientAddress RecipientAddress `json:"recipient_address"`
	PayTime          int64            `json:"pay_time"`
	PickupDoneTime   int64            `json:"pickup_done_time"`

	PackageList []PackageListShopee `json:"package_list"`
	InvoiceData InvoiceDataShopee   `json:"invoice_data"`

	ActualShippingFee float64 `json:"actual_shipping_fee"`

	ActualShippingFeeConfirmed float64 `json:"actual_shipping_fee_confirmed"`
	ShippingCarrier            string  `json:"shipping_carrier"`
	TotalAmount                int64   `json:"total_amount"`
	BuyerCancelReason          string  `json:"buyer_cancel_reason"`
	CancelBy                   string  `json:"cancel_by"`
	CancelReason               string  `json:"cancel_reason"`
	ReverseShippingFee         float64 `json:"reverse_shipping_fee"`
	CheckoutShippingCarrier    string  `json:"checkout_shipping_carrier"`
}

type InvoiceDataShopee struct {
	Number             string  `json:"number"`
	SeriesNumber       string  `json:"series_number"`
	AccessKey          string  `json:"access_key"`
	IssueDate          int64   `json:"issue_date"`
	TotalValue         float64 `json:"total_value"`
	ProductsTotalValue float64 `json:"products_total_value"`
	TaxCode            string  `json:"tax_code"`
}

type PackageListShopee struct {
	PackageNumber   string `json:"package_number"`
	LogisticsStatus string `json:"logistics_status"`
	ShippingCarrier string `json:"shipping_carrier"`
}

type ItemList struct {
	ItemName               string  `json:"item_name"`
	ItemSku                string  `json:"item_sku"`
	ModelQuantityPurchased int64   `json:"model_quantity_purchased"`
	ModelSku               string  `json:"model_sku"`
	ModelName              string  `json:"model_name"`
	ItemId                 int64   `json:"item_id"`
	ModelId                int64   `json:"model_id"`
	ModelOriginalPrice     float64 `json:"model_original_price"`
	ModelDiscountedPrice   float64 `json:"model_discounted_price"`
	Wholesale              bool    `json:"wholesale"`
	Weight                 float64 `json:"weight"`
	AddOnDeal              bool    `json:"add_on_deal"`
	MainItem               bool    `json:"main_item"`
	AddOnDealId            int64   `json:"add_on_deal_id"`
	PromotionType          string  `json:"promotion_type"`
	PromotionId            int64   `json:"promotion_id"`
	OrderItemId            int64   `json:"order_item_id"`
	PromotionGroupId       int64   `json:"promotion_group_id"`
}

type RecipientAddress struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Town        string `json:"town"`
	District    string `json:"district"`
	City        string `json:"city"`
	State       string `json:"state"`
	Region      string `json:"region"`
	Zipcode     string `json:"zipcode"`
	FullAddress string `json:"full_address"`
}

type CancelItemList struct {
	ItemId  int64  `json:"item_id"`
	ModelId int64  `json:"model_id"`
	OrderSn string `json:"order_sn"`
}

// PRODUCT
type ProductItemListHeader struct {
	Error           string          `json:"error"`
	Message         string          `json:"message"`
	ProductItemList ProductItemList `json:"response"`
	RequestId       string          `json:"request_id"`
}
type ProductItemList struct {
	Item        []ProductItemListDetail `json:"item"`
	TotalCount  int64                   `json:"total_count"`
	HasNextPage bool                    `json:"has_next_page"`
	NextOffset  int64                   `json:"next_offset"`
}

type ProductItemListDetail struct {
	ItemId     int64          `json:"item_id"`
	ItemStatus string         `json:"item_status"`
	UpdateTime int64          `json:"update_time"`
	DetailItem []BaseItemList `json:"detail_item"`
}

type BaseItemInfoHeader struct {
	Error            string           `json:"error"`
	Message          string           `json:"message"`
	BaseItemInfoList BaseItemInfoList `json:"response"`
	RequestId        string           `json:"request_id"`
}
type BaseItemInfoList struct {
	ItemList []BaseItemList `json:"item_list"`
}
type ArrayBaseItemShopee struct {
	ItemId int64 `json:"item_id"`
}

type BaseItemList struct {
	ItemId        int64                   `json:"item_id"`
	CategoryId    int64                   `json:"category_id"`
	ItemName      string                  `json:"item_name"`
	Description   string                  `json:"description"`
	ItemSku       string                  `json:"item_sku"`
	ItemStatus    string                  `json:"item_status"`
	AttributeList []BaseItemListAttribute `json:"attribute_list"`
	PriceInfo     []BaseItemPrice         `json:"price_info"`
	StockInfo     []BaseItemStock         `json:"stock_info"`
	HasModel      bool                    `json:"has_model"`
}

type BaseItemListAttribute struct {
	AttributeId           int64  `json:"attribute_id"`
	OriginalAttributeName string `json:"original_attribute_name"`
	IsMandatory           bool   `json:"is_mandatory"`
}
type BaseItemPrice struct {
	Currency      string `json:"currency"`
	OriginalPrice int64  `json:"original_price"`
	CurrentPrice  int64  `json:"current_price"`
}
type BaseItemStock struct {
	StockType       int64  `json:"stock_type"`
	StockLocationId string `json:"stock_location_id"`
	CurrentStock    int64  `json:"current_stock"`
	NormalStock     int64  `json:"normal_stock"`
	ReservedStock   int64  `json:"reserved_stock"`
}

//model detail
type ModelProductHeader struct {
	Error                  string                 `json:"error"`
	Message                string                 `json:"message"`
	ModelProductHeaderResp ModelProductHeaderResp `json:"response"`
	RequestId              string                 `json:"request_id"`
}

type ModelProductHeaderResp struct {
	TierVariation      []TierVariationProductDetail `json:"tier_variation"`
	ModelProductDetail []ModelProductDetail         `json:"model"`
}

type TierVariationProductDetail struct {
	Name       string           `json:"name"`
	OptionList []OptionListTier `json:"option_list"`
}

type OptionListTier struct {
	Option string `json:"option"`
}

type ModelProductDetail struct {
	ModelId        int64            `json:"model_id"`
	TierIndex      []int64          `json:"tier_index"`
	StockInfoModel []StockInfoModel `json:"stock_info"`
	PriceInfoModel []PriceInfoModel `json:"price_info"`
	ModelSku       string           `json:"model_sku"`
}

type StockInfoModel struct {
	StockType     int64 `json:"stock_type"`
	CurrentStock  int64 `json:"current_stock"`
	NormalStock   int64 `json:"normal_stock"`
	ReservedStock int64 `json:"reserved_stock"`
}

type PriceInfoModel struct {
	CurrentPrice                 float64 `json:"current_price"`
	OriginalPrice                float64 `json:"original_price"`
	InflatedPriceOfCurrentPrice  float64 `json:"inflated_price_of_current_price"`
	InflatedPriceOfOriginalPrice float64 `json:"inflated_price_of_original_price"`
}

// PRODUCT

//shipping
type ShippingParamHeader struct {
	Error             string            `json:"error"`
	Message           string            `json:"message"`
	ShippingParamInfo ShippingParamInfo `json:"response"`
	RequestId         string            `json:"request_id"`
}
type ShippingParamInfo struct {
	// Dropoff []Dropoff `json:"dropoff"`
	InfoNeeded InfoNeeded `json:"info_needed"`
	Pickup     Pickup     `json:"pickup"`
	Dropoff    Dropoff    `json:"dropoff"`
}
type InfoNeeded struct {
	Dropoff []interface{} `json:"dropoff"`
	Pickup  []interface{} `json:"pickup"`
}

type Pickup struct {
	AddressList []AddressList `json:"address_list"`
}
type Dropoff struct {
	BranchList []BranchList `json:"branch_list"`
}

type BranchList struct {
	BranchId int64  `json:"branch_id"`
	Region   string `json:"region"`
	State    string `json:"state"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Zipcode  string `json:"zipcode"`
	Town     string `json:"town"`
	District string `json:"district"`
	//TimeSlotList []TimeSlotList `json:"time_slot_list"`
}

type AddressList struct {
	Address      string         `json:"address"`
	AddressId    int64          `json:"address_id"`
	Region       string         `json:"region"`
	State        string         `json:"state"`
	City         string         `json:"city"`
	Town         string         `json:"town"`
	District     string         `json:"district"`
	Zipcode      string         `json:"zipcode"`
	TimeSlotList []TimeSlotList `json:"time_slot_list"`
}

type TimeSlotList struct {
	PickupTimeId string `json:"pickup_time_id"`
	Date         string `json:"date"`
	TimeText     string `json:"time_text"`
}

type TrackShippingParamHeader struct {
	Error                       string                      `json:"error"`
	Message                     string                      `json:"message"`
	ResTrackShippingParamHeader ResTrackShippingParamHeader `json:"response"`
	RequestId                   string                      `json:"request_id"`
}

type ResTrackShippingParamHeader struct {
	FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
	Hint                    string `json:"hint"`
	TrackingNumber          string `json:"tracking_number"`
}

//LOGISTICS
type ChannelLogisticsHeader struct {
	Error                string               `json:"error"`
	Message              string               `json:"message"`
	LogisticsChannelList LogisticsChannelList `json:"response"`
	RequestId            string               `json:"request_id"`
}
type LogisticsChannelList struct {
	DetailLogisticsChannelList []DetailLogisticsChannelList `json:"logistics_channel_list"`
}

type DetailLogisticsChannelList struct {
	LogisticsChannelId   int64                     `json:"logistics_channel_id"`
	LogisticsChannelName string                    `json:"logistics_channel_name"`
	CodEnabled           bool                      `json:"cod_enabled"`
	Enabled              bool                      `json:"enabled"`
	FeeType              string                    `json:"fee_type"`
	SizeList             SizeListLogistics         `json:"size_list"`
	WeightLimit          WeightLimitLogistics      `json:"weight_limit"`
	ItemMaxDimension     ItemMaxDimensionLogistics `json:"item_max_dimension"`
	VolumeLimit          VolumeLimitLogistics      `json:"volume_limit"`
	LogisticsDescription string                    `json:"logistics_description"`
	ForceEnable          bool                      `json:"force_enable"`
	MaskChannelId        int64                     `json:"mask_channel_id"`
}

type VolumeLimitLogistics struct {
	ItemMaxVolume float64 `json:"item_max_volume"`
	ItemMinVolume float64 `json:"item_min_volume"`
}

type SizeListLogistics struct {
	SizeId       string  `json:"size_id"`
	Name         string  `json:"name"`
	DefaultPrice float64 `json:"default_price"`
}

type WeightLimitLogistics struct {
	ItemMaxWeight float64 `json:"item_max_weight"`
	ItemMinWeight float64 `json:"item_min_weight"`
}

type ItemMaxDimensionLogistics struct {
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Length       float64 `json:"length"`
	Unit         string  `json:"unit"`
	DimensionSum float64 `json:"dimension_sum"`
}

//track resi
type TrackResiHeader struct {
	Error              string             `json:"error"`
	Message            string             `json:"message"`
	ResTrackResiHeader ResTrackResiHeader `json:"response"`
	RequestId          string             `json:"request_id"`
}

type ResTrackResiHeader struct {
	OrderSn         string             `json:"order_sn"`
	PackageNumber   string             `json:"package_number"`
	LogisticsStatus string             `json:"logistics_status"`
	TrackingInfo    []TrackingInfoResi `json:"tracking_info"`
}

type TrackingInfoResi struct {
	UpdateTime      int64     `json:"update_time"`
	Description     string    `json:"description"`
	LogisticsStatus string    `json:"logistics_status"`
	WaktuUpdate     time.Time `json:"waktu_pdate"`
}

//payment detail
type PaymentDetailShopee struct {
	Error                   string                  `json:"error"`
	Message                 string                  `json:"message"`
	PaymentDetailShopeeList PaymentDetailShopeeList `json:"response"`
	RequestId               string                  `json:"request_id"`
}

type PaymentDetailShopeeList struct {
	OrderSn           string                `json:"order_sn"`
	BuyerUserName     string                `json:"buyer_user_name"`
	ReturnOrderSnList []string              `json:"return_order_sn_list"`
	OrderIncome       OrderIncomeShopeeList `json:"order_income"`
}
type OrderIncomeShopeeList struct {
	EscrowAmount               float64                 `json:"escrow_amount"`
	BuyerTotalAmount           float64                 `json:"buyer_total_amount"`
	OriginalPrice              float64                 `json:"original_price"`
	SellerDiscount             float64                 `json:"seller_discount"`
	ShopeeDiscount             float64                 `json:"shopee_discount"`
	VoucherFromSeller          float64                 `json:"voucher_from_seller"`
	VoucherFromShopee          float64                 `json:"voucher_from_shopee"`
	Coins                      float64                 `json:"coins"`
	BuyerPaidShippingFee       float64                 `json:"buyer_paid_shipping_fee"`
	BuyerTransactionFee        float64                 `json:"buyer_transaction_fee"`
	CrossBorderTax             float64                 `json:"cross_border_tax"`
	PaymentPromotion           float64                 `json:"payment_promotion"`
	CommissionFee              float64                 `json:"commission_fee"`
	ServiceFee                 float64                 `json:"service_fee"`
	SellerTransactionFee       float64                 `json:"seller_transaction_fee"`
	SellerLostCompensation     float64                 `json:"seller_lost_compensation"`
	SellerCoinCashBack         float64                 `json:"seller_coin_cash_back"`
	EscrowTax                  float64                 `json:"escrow_tax"`
	FinalShippingFee           float64                 `json:"final_shipping_fee"`
	ActualShippingFee          float64                 `json:"actual_shipping_fee"`
	ShopeeShippingRebate       float64                 `json:"shopee_shipping_rebate"`
	ShippingFeeDiscountFrom3pl float64                 `json:"shipping_fee_discount_from_3pl"`
	SellerShippingDiscount     float64                 `json:"seller_shipping_discount"`
	EstimatedShippingFee       float64                 `json:"estimated_shipping_fee"`
	SellerVoucherCode          []string                `json:"seller_voucher_code"`
	DrcAdjustableRefund        float64                 `json:"drc_adjustable_refund"`
	CostOfGoodsSold            float64                 `json:"cost_of_goods_sold"`
	OriginalCostOfGoodsSold    float64                 `json:"original_cost_of_goods_sold"`
	OriginalShopeeDiscount     float64                 `json:"original_shopee_discount"`
	SellerReturnRefund         float64                 `json:"seller_return_refund"`
	ReverseShippingFee         float64                 `json:"reverse_shipping_fee"`
	Items                      []ItemListPaymentShopee `json:"items"`
	EscrowAmountPr             float64                 `json:"escrow_amount_pr"`
	BuyerTotalAmountPri        float64                 `json:"buyer_total_amount_pri"`
	OriginalPricePri           float64                 `json:"original_price_pri"`
	CommissionFeePri           float64                 `json:"commission_fee_pri"`
	ServiceFeePri              float64                 `json:"service_fee_pri"`
	DrcAdjustableRefundPri     float64                 `json:"drc_adjustable_refund_pri"`
	PriCurrency                string                  `json:"pri_currency"`
	AffCurrency                string                  `json:"aff_currency"`
	ExchangeRate               float64                 `json:"exchange_rate"`
	CreditCardPromotion        float64                 `json:"credit_card_promotion"`
	CreditCardTransactionFee   float64                 `json:"credit_card_transaction_fee"`
}

type ItemListPaymentShopee struct {
	ItemId                    int64   `json:"item_id"`
	ItemName                  string  `json:"item_name"`
	ItemSku                   string  `json:"item_sku"`
	ModelId                   int64   `json:"model_id"`
	ModelName                 string  `json:"model_name"`
	ModelSku                  string  `json:"model_sku"`
	OriginalPrice             float64 `json:"original_price"`
	DiscountedPrice           float64 `json:"discounted_price"`
	DiscountFromCoin          float64 `json:"discount_from_coin"`
	DiscountFromVoucherShopee float64 `json:"discount_from_voucher_shopee"`
	DiscountFromVoucherSeller float64 `json:"discount_from_voucher_seller"`
	ActivityType              string  `json:"activity_type"`
	IsMainItem                bool    `json:"is_main_item"`
	ActivityId                int64   `json:"activity_id"`
	QuantityPurchased         int64   `json:"quantity_purchased"`
}

type TableOrderSalesShopee struct {
	IdOrderSales             string    `json:"id_order_sales"`
	NoPesanan                string    `json:"no_pesanan"`
	StatusPesanan            string    `json:"status_pesanan"`
	StatusPembatalan         string    `json:"status_pembatalan"`
	NoResi                   string    `json:"no_resi"`
	OpsiPengiriman           string    `json:"opsi_pengiriman"`
	AntarKe                  string    `json:"antar_ke"`
	PesananDikirimSebelum    time.Time `json:"pesanan_dikirim_sebelum"`
	WaktuPengirimanDiatur    time.Time `json:"waktu_pengiriman_diatur"`
	WaktuPesananDibuat       time.Time `json:"waktu_pesanan_dibuat"`
	WaktuPembayaran          time.Time `json:"waktu_pembayaran"`
	SkuInduk                 string    `json:"sku_induk"`
	NamaProduk               string    `json:"nama_produk"`
	NamaReferensiSku         string    `json:"nama_referensi_sku"`
	NamaVariasi              string    `json:"nama_variasi"`
	HargaAwal                string    `json:"harga_awal"`
	HargaSetelahDiskon       string    `json:"harga_setelah_diskon"`
	Jumlah                   int16     `json:"jumlah"`
	TotalHargaProduk         string    `json:"total_harga_produk"`
	TotalDiskon              string    `json:"total_diskon"`
	DiskonDariPenjual        string    `json:"diskon_dari_penjual"`
	DiskonDariShopee         string    `json:"diskon_dari_shopee"`
	BeratProduk              string    `json:"berat_produk"`
	JmlProduk                int16     `json:"jml_produk"`
	TotalBerat               string    `json:"total_berat"`
	VoucherDitanggungPenjual string    `json:"voucher_ditanggung_penjual"`
	CashbackKoin             string    `json:"cashback_koin"`
	VoucherDitanggungShopee  string    `json:"voucher_ditanggung_shopee"`
	PaketDiskon              string    `json:"paket_diskon"`
	PaketDiskonShopee        string    `json:"paket_diskon_shopee"`
	PaketDiskonPenjual       string    `json:"paket_diskon_penjual"`
	PotonganKoinShopee       int16     `json:"potongan_koin_shopee"`
	DiskonKartuKredit        int16     `json:"diskon_kartu_kredit"`
	OngkirDibayarPembeli     int16     `json:"ongkir_dibayar_pembeli"`
	TotalPembayaran          string    `json:"total_pembayaran"`
	PerkiraanOngkir          string    `json:"perkiraan_ongkir"`
	CatatanDariPembeli       string    `json:"catatan_dari_pembeli"`
	Catatan                  string    `json:"catatan"`
	UsernamePembeli          string    `json:"username_pembeli"`
	NamaPenerima             string    `json:"nama_penerima"`
	Tlpn                     string    `json:"tlpn"`
	AlamatPengiriman         string    `json:"alamat_pengiriman"`
	Kota                     string    `json:"kota"`
	Provinsi                 string    `json:"provinsi"`
	WaktuPesananSelesai      time.Time `json:"waktu_pesanan_selesai"`
	Code                     string    `json:"code"`
}

type TablePaymentSalesShopee struct {
	IdPaymentSales                      string    `json:"id_payment_sales"`
	OrderId                             string    `json:"order_id"`
	Username                            string    `json:"username"`
	OrderCrateingDate                   time.Time `json:"order_crateing_date"`
	PayoutCompletedDate                 time.Time `json:"payout_completed_date"`
	OriginalProductPrice                int16     `json:"original_product_price"`
	SellerProductPromotion              int16     `json:"seller_product_promotion"`
	ProductDiscountRebateShopee         int16     `json:"product_discount_rebate_shopee"`
	Voucher                             int16     `json:"voucher"`
	ShippingFeePaidByBuyer              int16     `json:"shipping_fee_paid_by_buyer"`
	ShippingFeeDiscountFrom3pl          int16     `json:"shipping_fee_discount_from_3pl"`
	ShippingRebateFromShopee            int16     `json:"shipping_rebate_from_shopee"`
	ShippingFeePaidByShopeeOnYourBehalf int16     `json:"shipping_fee_paid_by_shopee_on_your_behalf"`
	RefundAmountToBuyer                 int16     `json:"refund_amount_to_buyer"`
	CommissionFee                       int16     `json:"commission_fee"`
	ServiceFee                          int16     `json:"service_fee"`
	TotalReleasedAmount                 int16     `json:"total_released_amount"`
	VoucherCode                         string    `json:"voucher_code"`
	Code                                string    `json:"code"`
	Refound                             int16     `json:"refound"`
	ShippingFeePromotionBySeller        string    `json:"shipping_fee_promotion_by_seller"`
	Status                              int       `json:"status"`
}

//upload image
type UploadImageShopeeHeader struct {
	Error                   string                  `json:"error"`
	Message                 string                  `json:"message"`
	Warning                 string                  `json:"warning"`
	UploadImageShopeeDetail UploadImageShopeeDetail `json:"response"`
	RequestId               string                  `json:"request_id"`
}

type UploadImageShopeeDetail struct {
	ImageInfo UploadImageShopeeInfo `json:"image_info"`
}

type UploadImageShopeeInfo struct {
	ImageId      string               `json:"image_id"`
	ImageUrlList []ImageUrlListShopee `json:"image_url_list"`
}

type ImageUrlListShopee struct {
	ImageUrlRegion string `json:"image_url_region"`
	ImageUrl       string `json:"image_url"`
}
