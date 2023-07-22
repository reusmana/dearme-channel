package models

type TokenJdId struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int64  `json:"created_at"`
}

type ErrorJdId struct {
	ErrorResponse struct {
		Code   string `json:"code"`
		ZhDesc string `json:"zh_desc"`
		EnDesc string `json:"en_desc"`
	} `json:"error_response"`
}

type HeaderProductJdId struct {
	JingdongSellerProductGetWareInfoListByVendorIDResponse struct {
		Code       string `json:"code"`
		ReturnType struct {
			Model struct {
				TotalNum      int `json:"totalNum"`
				SpuInfoVoList []struct {
					ImgUris                []string `json:"imgUris"`
					FullCategoryID         string   `json:"fullCategoryId"`
					BrandID                int      `json:"brandId"`
					Description            string   `json:"description"`
					SpuName                string   `json:"spuName"`
					WareStatus             int      `json:"wareStatus"`
					ProductArea            string   `json:"productArea"`
					MainImgURI             string   `json:"mainImgUri"`
					SpuID                  int      `json:"spuId"`
					FullCategoryName       string   `json:"fullCategoryName"`
					AfterSale              int      `json:"afterSale"`
					AppDescription         string   `json:"appDescription"`
					CommonAttributeNameMap struct {
					} `json:"commonAttributeNameMap,omitempty"`
					CommonAttributeIds string `json:"commonAttributeIds,omitempty"`
				} `json:"spuInfoVoList"`
				PageNum int `json:"pageNum"`
			} `json:"model"`
			Code      int  `json:"code"`
			IsSuccess bool `json:"isSuccess"`
		} `json:"returnType"`
	} `json:"jingdong_seller_product_getWareInfoListByVendorId_response"`
}

type DetailProductJdId struct {
	JingdongSellerProductGetSkuInfoBySpuIDAndVenderIDResponse struct {
		Code       string `json:"code"`
		ReturnType struct {
			Model []struct {
				SellerSkuID           string      `json:"sellerSkuId"`
				PackWide              string      `json:"packWide"`
				Weight                string      `json:"weight"`
				Status                int         `json:"status"`
				MainImgURI            string      `json:"mainImgUri"`
				SpuID                 int         `json:"spuId"`
				SkuName               string      `json:"skuName"`
				JdPrice               float64     `json:"jdPrice"`
				Upc                   string      `json:"upc"`
				SaleAttributeNameMap  interface{} `json:"saleAttributeNameMap,omitempty"`
				SaleAttributeIds      string      `json:"saleAttributeIds"`
				NetWeight             string      `json:"netWeight"`
				SkuID                 int         `json:"skuId"`
				PackHeight            string      `json:"packHeight"`
				PackLong              string      `json:"packLong"`
				Piece                 int         `json:"piece"`
				SaleAttributeNameMap0 interface{} `json:"saleAttributeNameMap,omitempty"`
			} `json:"model"`
			Code    int  `json:"code"`
			Success bool `json:"success"`
		} `json:"returnType"`
	} `json:"jingdong_seller_product_getSkuInfoBySpuIdAndVenderId_response"`
}

type DetailStockJdId struct {
	JingdongEpistockQueryEpiMerchantWareStockResponse struct {
		Code            string `json:"code"`
		EptRemoteResult struct {
			Code    int         `json:"code"`
			Message string      `json:"message"`
			Model   interface{} `json:"model"`
			Success bool        `json:"success"`
		} `json:"EptRemoteResult"`
	} `json:"jingdong_epistock_queryEpiMerchantWareStock_response"`
}

type OrdersDataJdId struct {
	JingdongSellerOrderGetOrderIDListByConditionResponse struct {
		Code   string `json:"code"`
		Result struct {
			Message string `json:"message"`
			Model   []int  `json:"model"`
			Code    int    `json:"code"`
			Success bool   `json:"success"`
		} `json:"result"`
	} `json:"jingdong_seller_order_getOrderIdListByCondition_response"`
}

type OrderDetailJdIdOLD struct {
	JingdongSellerOrderGetOrderInfoByOrderIDResponse struct {
		Code   string `json:"code"`
		Result struct {
			Message string `json:"message"`
			Model   struct {
				Phone          string  `json:"phone"`
				InstallmentFee float64 `json:"installmentFee"`
				FullCutAmount  float64 `json:"fullCutAmount"`
				VenderID       int     `json:"venderId"`
				PayTime        int64   `json:"payTime"`
				JdSubsidy      float64 `json:"jdSubsidy"`
				SendPay        string  `json:"sendPay"`
				City           string  `json:"city"`
				UserPin        string  `json:"userPin"`
				BookTime       int64   `json:"bookTime"`
				OrderSkuNum    int     `json:"orderSkuNum"`
				O2OOrder       bool    `json:"o2oOrder"`
				ModifyTime     int64   `json:"modifyTime"`
				OrderState     int     `json:"orderState"`
				OrderType      int     `json:"orderType"`
				ServiceType    int     `json:"serviceType"`
				AddLng         string  `json:"addLng"`
				TaxationAmount float64 `json:"taxationAmount"`
				Mobile         string  `json:"mobile"`
				CreateTime     int64   `json:"createTime"`
				CouponAmount   float64 `json:"couponAmount"`
				SellerSubsidy  float64 `json:"sellerSubsidy"`
				State          string  `json:"state"`
				Area           string  `json:"area"`
				OrderSkuinfos  []struct {
					HasPromo        int     `json:"hasPromo"`
					CostPrice       float64 `json:"costPrice"`
					Weight          int     `json:"weight"`
					CouponAmount    float64 `json:"couponAmount"`
					FullCutAmount   float64 `json:"fullCutAmount"`
					SpuID           int     `json:"spuId"`
					JdPrice         float64 `json:"jdPrice"`
					SkuName         string  `json:"skuName"`
					OrderSkuUUIDID  string  `json:"orderSkuUuidId"`
					SkuNumber       int     `json:"skuNumber"`
					SkuImage        string  `json:"skuImage"`
					SkuID           int     `json:"skuId"`
					PopSkuID        string  `json:"popSkuId"`
					TaxationAmount  float64 `json:"taxationAmount"`
					Commission      float64 `json:"commission"`
					PromotionAmount float64 `json:"promotionAmount"`
					CrossType       int     `json:"crossType"`
				} `json:"orderSkuinfos"`
				FreightAmount   float64 `json:"freightAmount"`
				PromotionAmount float64 `json:"promotionAmount"`
				OrderID         int     `json:"orderId"`
				TotalPrice      float64 `json:"totalPrice"`
				CustomerName    string  `json:"customerName"`
				DeliveryType    int     `json:"deliveryType"`
				PaymentType     int     `json:"paymentType"`
				PostCode        string  `json:"postCode"`
				CarrierCode     int     `json:"carrierCode"`
				AddLat          string  `json:"addLat"`
				Address         string  `json:"address"`
				PaySubtotal     float64 `json:"paySubtotal"`
			} `json:"model"`
			Code    int  `json:"code"`
			Success bool `json:"success"`
		} `json:"result"`
	} `json:"jingdong_seller_order_getOrderInfoByOrderId_response"`
}

type OrderDetailJdId struct {
	JingdongSellerOrderGetOrderInfoByOrderIDResponse struct {
		Code   string `json:"code"`
		Result struct {
			Message string `json:"message"`
			Model   struct {
				Phone            string  `json:"phone"`
				InstallmentFee   float64 `json:"installmentFee"`
				FullCutAmount    float64 `json:"fullCutAmount"`
				VenderID         int     `json:"venderId"`
				PayTime          int64   `json:"payTime"`
				JdSubsidy        float64 `json:"jdSubsidy"`
				SendPay          string  `json:"sendPay"`
				City             string  `json:"city"`
				UserPin          string  `json:"userPin"`
				BookTime         int64   `json:"bookTime"`
				OrderSkuNum      int     `json:"orderSkuNum"`
				O2OOrder         bool    `json:"o2oOrder"`
				ModifyTime       int64   `json:"modifyTime"`
				OrderState       int     `json:"orderState"`
				OrderType        int     `json:"orderType"`
				CarrierCompany   string  `json:"carrierCompany"`
				ServiceType      int     `json:"serviceType"`
				AddLng           string  `json:"addLng"`
				TaxationAmount   float64 `json:"taxationAmount"`
				Mobile           string  `json:"mobile"`
				CreateTime       int64   `json:"createTime"`
				CouponAmount     float64 `json:"couponAmount"`
				SellerSubsidy    float64 `json:"sellerSubsidy"`
				State            string  `json:"state"`
				ExpressAttribute int     `json:"expressAttribute"`
				Area             string  `json:"area"`
				ExpressNo        string  `json:"expressNo"`
				OrderSkuinfos    []struct {
					HasPromo        int     `json:"hasPromo"`
					CostPrice       float64 `json:"costPrice"`
					Weight          int     `json:"weight"`
					CouponAmount    float64 `json:"couponAmount"`
					FullCutAmount   float64 `json:"fullCutAmount"`
					SpuID           int     `json:"spuId"`
					JdPrice         float64 `json:"jdPrice"`
					SkuName         string  `json:"skuName"`
					OrderSkuUUIDID  string  `json:"orderSkuUuidId"`
					SkuNumber       int     `json:"skuNumber"`
					SkuImage        string  `json:"skuImage"`
					SkuID           int     `json:"skuId"`
					PopSkuID        string  `json:"popSkuId"`
					TaxationAmount  float64 `json:"taxationAmount"`
					Commission      float64 `json:"commission"`
					PromotionAmount float64 `json:"promotionAmount"`
					CrossType       int     `json:"crossType"`
				} `json:"orderSkuinfos"`
				FreightAmount   float64 `json:"freightAmount"`
				PromotionAmount float64 `json:"promotionAmount"`
				OrderID         int     `json:"orderId"`
				TotalPrice      float64 `json:"totalPrice"`
				CustomerName    string  `json:"customerName"`
				DeliveryType    int     `json:"deliveryType"`
				PaymentType     int     `json:"paymentType"`
				PostCode        string  `json:"postCode"`
				CarrierCode     int     `json:"carrierCode"`
				AddLat          string  `json:"addLat"`
				Address         string  `json:"address"`
				PaySubtotal     float64 `json:"paySubtotal"`
			} `json:"model"`
			Code    int  `json:"code"`
			Success bool `json:"success"`
		} `json:"result"`
	} `json:"jingdong_seller_order_getOrderInfoByOrderId_response"`
}

type PrintDocShipJdId struct {
	JingdongSellerOrderPrintOrderResponse struct {
		Code   string `json:"code"`
		Result struct {
			Message string `json:"message"`
			Model   struct {
				Content   string `json:"content"`
				ExpressNo string `json:"expressNo"`
				OrderID   int    `json:"orderId"`
			} `json:"model"`
			Code    int  `json:"code"`
			Success bool `json:"success"`
		} `json:"result"`
	} `json:"jingdong_seller_order_printOrder_response"`
}

type ReqPickJdId struct {
	JingdongSellerOrderSendGoodsOpenAPIResponse struct {
		Code   string `json:"code"`
		Result struct {
			Message string `json:"message"`
			Model   struct {
				ExpressNo      string `json:"expressNo"`
				ExpressCompany string `json:"expressCompany"`
				ExpressID      string `json:"expressId"`
				OrderID        int    `json:"orderId"`
			} `json:"model"`
			Code    int  `json:"code"`
			Success bool `json:"success"`
		} `json:"result"`
	} `json:"jingdong_seller_order_sendGoodsOpenApi_response"`
}

type UpdateStockJdID struct {
	JingdongEpistockUpdateEpiMerchantWareStockResponse struct {
		Code            string `json:"code"`
		EptRemoteResult struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Model   struct {
				Num674177247 struct {
					Code    int    `json:"code"`
					Message string `json:"message"`
					Success bool   `json:"success"`
				} `json:"674177247"`
			} `json:"model"`
			Success bool `json:"success"`
		} `json:"EptRemoteResult"`
	} `json:"jingdong_epistock_updateEpiMerchantWareStock_response"`
}
type ErrorUpdateStockJdID struct {
	JingdongEpistockUpdateEpiMerchantWareStockResponse struct {
		Code            string `json:"code"`
		EptRemoteResult struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Success bool   `json:"success"`
		} `json:"EptRemoteResult"`
	} `json:"jingdong_epistock_updateEpiMerchantWareStock_response"`
}
