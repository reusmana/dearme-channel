package shopeeController

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rals/dearme-channel/controllers/blibliController"
	"github.com/rals/dearme-channel/controllers/bukalapakController"
	"github.com/rals/dearme-channel/controllers/jdidController"
	"github.com/rals/dearme-channel/controllers/lazadaController"
	"github.com/rals/dearme-channel/controllers/tiktokController"
	"github.com/rals/dearme-channel/controllers/tokpedController"
	"github.com/rals/dearme-channel/controllers/zaloraController"

	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/helpers"
	"github.com/rals/dearme-channel/models"
	"github.com/rals/dearme-channel/models/response"
	"github.com/rals/dearme-channel/repositories/tokenRepository"

	"github.com/rals/dearme-channel/services/tokenService"
	"github.com/rals/dearme-channel/utils"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/google/uuid"
	// "encoding/json"
)

func AuthShopee(path string, token string, shopid string) (string, string, string, string, string, string) {
	now := time.Now()
	timest := strconv.FormatInt(now.Unix(), 10)
	host := os.Getenv("URI_API_SHOPEE")
	paths := path
	partnerId := os.Getenv("PARTNER_ID_SHOPEE")
	partnerKey := os.Getenv("PARTNER_KEY_SHOPEE")
	baseString := ""
	if token != "" {
		baseString = fmt.Sprintf("%s%s%s%s%s", partnerId, paths, timest, token, shopid)
	} else {
		baseString = fmt.Sprintf("%s%s%s", partnerId, paths, timest)
	}

	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))

	return host, partnerId, timest, sign, token, shopid

}

func GetData(c *gin.Context) {
	var response response.ResponseCrud

	paths := "/api/v2/shop/auth_partner"
	redirectUrl := "https://supplier.ramayana.co.id/getauth.shopee/"
	host, partnerId, timest, sign, _, _ := AuthShopee(paths, "", "")
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&sign=%s&redirect=%s", partnerId, timest, sign, redirectUrl)

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)
	}

	c.JSON(http.StatusOK, response)
	return
}

func GetToken(c *gin.Context) {
	code := c.Query("code")
	shop_id := c.Query("shop_id")

	var response response.ResponseCrud

	paths := "/api/v2/auth/token/get"
	host, partnerId, timest, sign, _, _ := AuthShopee(paths, "", "")

	var jsonString = `{"code":"` + code + `","shop_id":` + shop_id + `,"partner_id":` + partnerId + `}`

	var body_url = []byte(jsonString)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&sign=%s", partnerId, timest, sign)

	req, err := http.NewRequest("POST", urlshopee, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)

		//insert token
		tokenService.SaveToken("shopee", datas, "GetToken")

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetTokenVw(c *gin.Context) { //03-02-2023 update token shopee
	code := c.Query("code")

	var response response.ResponseCrud

	paths := "/api/v2/public/get_token_by_resend_code"
	host, partnerId, timest, sign, _, _ := AuthShopee(paths, "", "")

	var jsonString = `{"resend_code":"` + code + `"}`

	var body_url = []byte(jsonString)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&sign=%s", partnerId, timest, sign)

	req, err := http.NewRequest("POST", urlshopee, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)

		//insert token
		tokenService.SaveToken("shopee", datas, "GetToken")

	}

	c.JSON(http.StatusOK, response)
	return
}

func RefreshTokenAuto(c *gin.Context) {

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/auth/access_token/get"
	host, partnerId, timest, sign, _, _ := AuthShopee(paths, "", "")
	refreshtoken := fmt.Sprintf("%v", ObjToken.Value2)
	var jsonString = `{"refresh_token":"` + refreshtoken + `","shop_id":` + os.Getenv("SHOP_ID_SHOPEE") + `,"partner_id":` + partnerId + `}`

	var body_url = []byte(jsonString)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&sign=%s", partnerId, timest, sign)

	req, err := http.NewRequest("POST", urlshopee, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)

		//insert token
		tokenService.SaveToken("shopee", datas, "RefreshToken")

	}

	tiktokController.RefreshTokenTiktok(c)
	jdidController.RefreshToken(c)
	tokpedController.GetTokenTokped(c)
	bukalapakController.RefreshTokenAuto()
	lazadaController.RefreshTokenAuto()
}

func RefreshTokenV2(c *gin.Context) response.ResponseCrud {

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/public/get_refresh_token_by_upgrade_code"
	host, partnerId, timest, sign, _, _ := AuthShopee(paths, "", "")
	refreshtoken := fmt.Sprintf("%v", ObjToken.Value2)
	var jsonString = `{"refresh_token":"` + refreshtoken + `","shop_id":` + os.Getenv("SHOP_ID_SHOPEE") + `,"partner_id":` + partnerId + `}`

	var body_url = []byte(jsonString)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&sign=%s", partnerId, timest, sign)

	req, err := http.NewRequest("POST", urlshopee, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)

		//insert token
		//tokenService.SaveToken("shopee", datas, "RefreshToken")

	}

	return response
}

func GetOrder(c *gin.Context) {

	ObjToken := tokenService.FindToken("shopee")
	fmt.Println("GetOrder")
	var response response.ResponseCrud
	StatusOrder := c.Param("status")
	paths := "/api/v2/order/get_order_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")
	pagesize := os.Getenv("PAGE_SIZE_ORDER")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	fmt.Println("LEWAT")
	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari lalu
	//kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-24)*time.Hour).Unix(), 10) //1 hari lalu

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&time_from=%s&time_to=%s&time_range_field=%s&page_size=%s", partnerId, timest, refreshtoken, shop_id, sign, kmrn, skrg, "create_time", pagesize)
	//urlshopee += "&order_status=READY_TO_SHIP"
	urlshopee += "&order_status=" + StatusOrder

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get order")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err
		StatusOrder = "CANCELLED"

	} else {
		fmt.Println("sukses get order")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)
		dataJson := parseJsonList(data)

		tokenService.CekError("shopee", datas, "GetOrder")

		if datas["message"] == "" {
			for _, value := range dataJson.ResponseOrder.OrderList {
				//cek jika sudah ada di wms_sales_order gak perlu cari lagi
				if StatusOrder == "CANCELLED" {
					//objCekCancel, _ := tokenRepository.FindSalesOrderParam(value.OrderSn, "3")
					//if objCekCancel.NoOrder == "" {
					objCekCancel, _ := tokenRepository.FindSalesOrder(value.OrderSn)
					if objCekCancel.StatusProcessOrder != "3" && objCekCancel.StatusProcessOrder != "" {

						fmt.Println("CEK GetOrder " + StatusOrder)
						GetOrderDetailProses(c, value.OrderSn)
					}
					//GetOrderDetailProses(c, value.OrderSn)
				} else if StatusOrder == "COMPLETED" {
					fmt.Println(StatusOrder)
					GetOrderDetailProses(c, value.OrderSn)
				} else {
					objCek, _ := tokenRepository.FindSalesOrder(value.OrderSn)
					if objCek.NoOrder == "" {
						GetOrderDetailProses(c, value.OrderSn)
					}
				}

			}

			if dataJson.ResponseOrder.More == true {
				GetOrderLoop(c, dataJson.ResponseOrder.NextCursor, StatusOrder)
			}

			// //call API PROCESSED
			// if StatusOrder == "READY_TO_SHIP" {
			// 	GetProcessedOrder("PROCESSED")
			// }

			// //call API CANCEL
			// if StatusOrder == "PROCESSED" {
			// 	GetProcessedOrder("CANCELLED")
			// }

			//GET ORDERAN ZALORA
			//zaloraController.GetProcessedOrder("canceled")
			// if StatusOrder == "CANCELLED" {
			// 	zaloraController.GetProcessedOrderAuto("pending")
			// 	tiktokController.GetProcessedOrderAutoTiktok("111") //awaiting_shipment
			// 	tiktokController.GetProcessedOrderAutoTiktok("140") //cancel
			// }

		}

	}

	//call API PROCESSED
	if StatusOrder == "READY_TO_SHIP" {
		GetProcessedOrder("PROCESSED")
	}

	//call API CANCEL
	if StatusOrder == "PROCESSED" {
		GetProcessedOrder("CANCELLED")
	}

	CheckOrderChannelList(StatusOrder)

	c.JSON(http.StatusOK, response)
	return
}

func CheckOrderChannelList(StatusOrder string) {
	if StatusOrder == "CANCELLED" {
		// zaloraController.GetProcessedOrderAuto("pending")
		zaloraController.GetProcessedOrderAutoV2("pending")
		tiktokController.GetProcessedOrderAutoTiktok("111") //awaiting_shipment
		tiktokController.GetProcessedOrderAutoTiktok("140") //cancel

		//blibliController.GetOrderBlibliAuto("FP")

		jdidController.GetProcessedOrderAutoJdId("1") //awaiting shipment
		jdidController.GetProcessedOrderAutoJdId("5") //cancel

		tokpedController.GetOrdersAuto("220") // order ready to process.
		tokpedController.GetOrdersAuto("0")   //Seller cancel order.
		tokpedController.GetOrdersAuto("3")   //Order Reject Due Empty Stock.
		tokpedController.GetOrdersAuto("5")   //Order Canceled by Fraud
		tokpedController.GetOrdersAuto("6")   //Order Rejected (Auto Cancel Out of Stock)
		tokpedController.GetOrdersAuto("10")  //Order rejected by seller.
		tokpedController.GetOrdersAuto("15")  //Instant Cancel by Buyer.

		bukalapakController.GetOrdersAuto("paid")
		bukalapakController.GetOrdersAuto("cancelled")

		lazadaController.GetOrderAuto("pending")
		lazadaController.GetOrderAuto("canceled")
	} else if StatusOrder == "COMPLETED" {
		//zaloraController.GetProcessedOrderAuto("delivered")
		zaloraController.GetProcessedOrderAutoV2("delivered")
		bukalapakController.GetOrdersAuto("remitted")
		lazadaController.GetOrderAuto("delivered")
		tiktokController.GetProcessedOrderAutoTiktok("130") //COMPLETED

	}
}

func GetProcessedOrder(status string) {
	//PROCESSED
	//urlApi := "https://api-channel.ramayana.co.id:9130/api/v1/shopee/getOrder/" + status
	urlApi := os.Getenv("URL_ORDER_SHOPEE") + status

	req, err := http.NewRequest("GET", urlApi, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("gagal GetProcessedOrder")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

	} else {
		fmt.Println("sukses GetProcessedOrder")
		defer resp.Body.Close()
		// data, _ := ioutil.ReadAll(resp.Body)
		// datas := utils.GetByteToInterface(data)
	}

}

func GetOrderLoop(c *gin.Context, cursor string, statusorder string) {
	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/order/get_order_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")
	pagesize := os.Getenv("PAGE_SIZE_ORDER")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari lalu
	//kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-24)*time.Hour).Unix(), 10)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&time_from=%s&time_to=%s&time_range_field=%s&page_size=%s", partnerId, timest, refreshtoken, shop_id, sign, kmrn, skrg, "create_time", pagesize)
	// urlshopee += "&order_status=READY_TO_SHIP"
	urlshopee += "&order_status=" + statusorder
	urlshopee += "&cursor=" + cursor

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get order")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses get order")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)
		dataJson := parseJsonList(data)

		tokenService.CekError("shopee", datas, "GetOrder")

		if datas["message"] == "" {
			for _, value := range dataJson.ResponseOrder.OrderList {
				if statusorder == "CANCELLED" {
					// objCekCancel, _ := tokenRepository.FindSalesOrderParam(value.OrderSn, "3")
					// if objCekCancel.NoOrder == "" {
					objCekCancel, _ := tokenRepository.FindSalesOrder(value.OrderSn)
					if objCekCancel.StatusProcessOrder != "3" && objCekCancel.StatusProcessOrder != "" {
						fmt.Println("CEK GetOrderLoop " + statusorder)
						GetOrderDetailProses(c, value.OrderSn)
					}
					//GetOrderDetailProses(c, value.OrderSn)
				} else if statusorder == "COMPLETED" {
					fmt.Println(statusorder)
					GetOrderDetailProses(c, value.OrderSn)
				} else {
					objCek, _ := tokenRepository.FindSalesOrder(value.OrderSn)
					if objCek.NoOrder == "" {
						GetOrderDetailProses(c, value.OrderSn)
					}
				}
				//GetOrderDetailProses(c, value.OrderSn) //nonaktif sementara

			}

			if dataJson.ResponseOrder.More == true {
				GetOrderLoop(c, dataJson.ResponseOrder.NextCursor, statusorder)
			}

		}

	}

}
func parseJsonList(jsonBuffer []byte) models.ListOrderShopee {

	ListOrderShopee := models.ListOrderShopee{}

	err := json.Unmarshal(jsonBuffer, &ListOrderShopee)
	if err != nil {
		return ListOrderShopee
	}

	// the array is now filled with users
	return ListOrderShopee

}

func parseJsonListEscrow(jsonBuffer []byte) models.DetailEscrowsShopee {

	DetailEscrowsShopee := models.DetailEscrowsShopee{}

	err := json.Unmarshal(jsonBuffer, &DetailEscrowsShopee)
	if err != nil {
		return DetailEscrowsShopee
	}

	// the array is now filled with users
	return DetailEscrowsShopee

}

func parseJsonProductSKU(jsonBuffer []byte) models.ListProductBYSkuShopee {

	ListProductBYSkuShopee := models.ListProductBYSkuShopee{}

	err := json.Unmarshal(jsonBuffer, &ListProductBYSkuShopee)
	if err != nil {
		return ListProductBYSkuShopee
	}

	// the array is now filled with users
	return ListProductBYSkuShopee

}

func parseJsonUploadImage(jsonBuffer []byte) models.UploadImageShopeeHeader {

	UploadImageShopeeHeader := models.UploadImageShopeeHeader{}

	err := json.Unmarshal(jsonBuffer, &UploadImageShopeeHeader)
	if err != nil {
		return UploadImageShopeeHeader
	}

	// the array is now filled with users
	return UploadImageShopeeHeader

}

func parseJson(jsonBuffer []byte) models.Header {

	header := models.Header{}

	err := json.Unmarshal(jsonBuffer, &header)
	if err != nil {
		return header
	}

	// the array is now filled with users
	return header

}

func parseJsonPaymentShopee(jsonBuffer []byte) models.PaymentDetailShopee {

	PaymentDetailShopee := models.PaymentDetailShopee{}

	err := json.Unmarshal(jsonBuffer, &PaymentDetailShopee)
	if err != nil {
		return PaymentDetailShopee
	}

	// the array is now filled with users
	return PaymentDetailShopee

}

func GetOrderDetailProses(c *gin.Context, noorder string) {
	NoOrder := noorder

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/order/get_order_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn_list=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)
	urlshopee += "&response_optional_fields=total_amount,buyer_username,item_list,recipient_address"
	urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get order detail")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses get order detail")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJson(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)

		if datasx["message"] == "" {
			//insert ke wms_sales_order
			//
			objPayment := GetDetailPaymentShopeeAuto(noorder)
			tokenService.SaveSalesOrder("shopee", datas, objPayment)
			tokenService.UpdateAmountShopee(objPayment)
		} else {
			tokenService.CekError("shopee", datasx, "GetOrderDetail")
		}

	}

}

func GetOrderDetail(c *gin.Context) { //cek dari api-wms jika orderan ready to ship
	NoOrder := c.Param("noorder")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/order/get_order_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn_list=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)
	urlshopee += "&response_optional_fields=total_amount,buyer_username,item_list,recipient_address"
	urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason,invoice_data,actual_shipping_fee,payment_method"
	//urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason,actual_shipping_fee,note,note_update_time,pay_time,dropshipper_phone,buyer_cancel_reason,actual_shipping_fee_confirmed"
	//urlshopee += ",checkout_shipping_carrier,reverse_shipping_fee,invoice_data,package_list,pickup_done_time"

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get order detail")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses get order detail")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJson(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
		//response.Result = string(data)
		//fmt.Println(time.Unix(datas.Response.OrderList[0].UpdateTime, 0))
		objPayment := GetDetailPaymentShopeeAuto(NoOrder)
		if datasx["message"] == "" {
			//insert ke wms_sales_order
			//objPayment := GetDetailPaymentShopeeAuto(NoOrder)
			tokenService.SaveSalesOrder("shopee", datas, objPayment)
			//fmt.Println("Gak Di Save")
		} else {
			tokenService.CekError("shopee", datasx, "GetOrderDetail")
		}
		//fmt.Println(string(data))

		//update amount disini
		tokenService.UpdateAmountShopee(objPayment)

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrderDetailView(c *gin.Context) { //cek dari api-wms jika orderan ready to ship
	NoOrder := c.Param("noorder")

	Viewss := c.Param("view")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/order/get_order_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn_list=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)
	urlshopee += "&response_optional_fields=total_amount,buyer_username,item_list,recipient_address"
	urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason,invoice_data,actual_shipping_fee,payment_method"
	//urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason,actual_shipping_fee,note,note_update_time,pay_time,dropshipper_phone,buyer_cancel_reason,actual_shipping_fee_confirmed"
	//urlshopee += ",checkout_shipping_carrier,reverse_shipping_fee,invoice_data,package_list,pickup_done_time"

	if Viewss != "0" {
		urlshopee += Viewss

	}
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get order detail")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses get order detail")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJson(data)
		//datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetDiscountView(c *gin.Context) {
	Idpromo := c.Param("idpromo")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/discount/get_discount"
	paths = "/api/v2/bundle_deal/get_bundle_deal_item"
	paths = "/api/v2/bundle_deal/get_bundle_deal"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&bundle_deal_id=%s", partnerId, timest, refreshtoken, shop_id, sign, Idpromo)

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get order detail")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses get order detail")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJson(data)
		//datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func CancelOrder(c *gin.Context) {
	NoOrder := c.Param("noorder")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/order/get_order_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn_list=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)
	urlshopee += "&response_optional_fields=total_amount,buyer_username,item_list,recipient_address"
	urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get order detail cancel")

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err
		c.JSON(http.StatusOK, response)
	} else {
		fmt.Println("sukses get order detail cancel")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJson(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas

		if datas.Message == "" {
			//cancel disini
			//var objCancel []models.CancelItemList

			jsonStringIsi := `"item_list":[ "`

			i := 0
			for _, element1 := range datas.Response.OrderList[0].ItemList {
				if i == 0 {
					jsonStringIsi += `{"item_id":` + strconv.Itoa(int(element1.ItemId)) + `,`
					jsonStringIsi += `"model_id":` + strconv.Itoa(int(element1.ModelId)) + `}`
				} else {
					jsonStringIsi += ","
					jsonStringIsi += `{"item_id":` + strconv.Itoa(int(element1.ItemId)) + `,`
					jsonStringIsi += `"model_id":` + strconv.Itoa(int(element1.ModelId)) + `}`
				}

				fmt.Println(element1.ItemId)
			}
			jsonStringIsi += `]`
			var jsonString = `{"order_sn":"` + datas.Response.OrderList[0].OrderSn + `","cancel_reason":"OUT_OF_STOCK",` + jsonStringIsi + `}`

			var body_url = []byte(jsonString)
			UrlCancel(c, body_url)
			// fmt.Println("disini cancel")
			// fmt.Println(string(body_url))

		} else {
			c.JSON(http.StatusOK, response)
		}

	}

	// c.JSON(http.StatusOK, response)
	return
}

func UrlCancel(c *gin.Context, body []byte) {

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/order/cancel_order"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn_list=%s", partnerId, timest, refreshtoken, shop_id, sign)

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(body))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal cancel")

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses cancel")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJson(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas

	}

	c.JSON(http.StatusOK, response)
	return

}

func GetNoResi(c *gin.Context) {
	NoOrder := c.Param("noorder")
	//NoOrder := "211109461XQ3MW"

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/logistics/get_tracking_number"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)
	// urlshopee += "&response_optional_fields=total_amount,buyer_username,item_list,recipient_address"
	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get no resi")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses get no resi")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJson(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
		fmt.Println("=================")
		fmt.Println(datasx)
		fmt.Println("=================")
		// if datasx["message"] == "" {
		// 	//insert ke wms_sales_order
		// 	tokenService.SaveSalesOrder("shopee", datas)
		// } else {
		// 	tokenService.CekError("shopee", datasx, "GetOrderDetail")
		// }

	}

	c.JSON(http.StatusOK, response)
	return
}

func parseJsonItemList(jsonBuffer []byte) models.ProductItemListHeader {

	ProductItemListHeader := models.ProductItemListHeader{}

	err := json.Unmarshal(jsonBuffer, &ProductItemListHeader)
	if err != nil {
		return ProductItemListHeader
	}

	// the array is now filled with users
	return ProductItemListHeader

}

func parseJsonBaseItem(jsonBuffer []byte) models.BaseItemInfoHeader {

	BaseItemInfoHeader := models.BaseItemInfoHeader{}

	err := json.Unmarshal(jsonBuffer, &BaseItemInfoHeader)
	if err != nil {
		return BaseItemInfoHeader
	}

	// the array is now filled with users
	return BaseItemInfoHeader

}
func parseJsonBaseItemArray(jsonBuffer []byte) models.BaseItemInfoHeader {

	BaseItemInfoHeader := models.BaseItemInfoHeader{}

	err := json.Unmarshal(jsonBuffer, &BaseItemInfoHeader)
	if err != nil {
		return BaseItemInfoHeader
	}

	// the array is now filled with users
	return BaseItemInfoHeader

}

func GetItemList(c *gin.Context) {

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_item_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	item_status := "NORMAL"
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	urlshopee += "&offset=1&page_size=10&item_status=" + item_status
	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get item list")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses get item list")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonItemList(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		// response.Result = datas
		fmt.Println("=================")
		//fmt.Println(len(datasx))
		// fmt.Println(datas)
		fmt.Println("=================")
		if datasx["message"] == "" {
			//insert ke wms_sales_order
			//tokenService.SaveSalesOrder("shopee", datas)
			i := 0
			var itemidbase []models.ArrayBaseItemShopee
			fmt.Println(len(datas.ProductItemList.Item))
			for _, element1 := range datas.ProductItemList.Item {

				fmt.Println(element1.ItemId)
				//var isi []models.BaseItemInfoHeader

				itemidbaseloop := models.ArrayBaseItemShopee{
					ItemId: element1.ItemId,
				}

				if len(itemidbase) < 50 {
					itemidbase = append(itemidbase, itemidbaseloop)
				}

				if len(itemidbase) == 50 {
					isi := GetItemListBaseLoopArray(itemidbase)
					var itemidbase []models.ArrayBaseItemShopee
					fmt.Println(itemidbase)
					//fmt.Println(isi)
					response.Result = isi
				}

				//datas.ProductItemList.Item[i].DetailItem = isi.BaseItemInfoList.ItemList

				//fmt.Println(isi)

				i++
				fmt.Println(i)
				if i == len(datas.ProductItemList.Item) {
					if len(itemidbase) > 0 {
						fmt.Println("Masuk sini terakhir")
						isi := GetItemListBaseLoopArray(itemidbase)
						var itemidbase []models.ArrayBaseItemShopee
						fmt.Println(itemidbase)
						fmt.Println(len(isi.BaseItemInfoList.ItemList))
						//fmt.Println(isi)
						response.Result = isi
					}

				}

			}
		} else {
			tokenService.CekError("shopee", datasx, "GetItemList")
		}
		//response.Result = datas

	}

	c.JSON(http.StatusOK, response)
	return
}
func GetItemListLimitAuto(c *gin.Context, wg *sync.WaitGroup) {

	//ItemId := c.Param("itemid")
	fmt.Println("mulai product shopee " + time.Now().String())

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_item_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-3600)*time.Hour).Unix(), 10) //5 bulan lalu
	//kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-1440)*time.Hour).Unix(), 10) //2 bulan lalu

	offset := 0
	pagesize := os.Getenv("PAGE_SIZE_ORDER")
	ItemStatus := "NORMAL"
	urlshopee += "&update_time_from=" + kmrn
	urlshopee += "&update_time_to=" + skrg
	urlshopee += "&offset=" + strconv.Itoa(offset)
	urlshopee += "&page_size=" + pagesize
	urlshopee += "&item_status=" + ItemStatus

	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetItemListLimit")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses GetItemListLimit")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonItemList(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas

		if datasx["message"] == "" {

			var itemidbase []models.ArrayBaseItemShopee
			totalproduct := 0
			i := 0
			//fmt.Println("Total data productItem " + strconv.Itoa(int(datas.ProductItemList.TotalCount)))
			for _, element1 := range datas.ProductItemList.Item {
				totalproduct++
				//fmt.Println(element1.ItemId)
				var isi models.BaseItemInfoHeader

				//fmt.Println(element1.ItemId)

				itemidbaseloop := models.ArrayBaseItemShopee{
					ItemId: element1.ItemId,
				}
				//fmt.Println("LEN ATAS itemidbase" + strconv.Itoa(len(itemidbase)))

				if len(itemidbase) < 50 {
					itemidbase = append(itemidbase, itemidbaseloop)
				}
				//fmt.Println("LEN itemidbase" + strconv.Itoa(len(itemidbase)))
				i++

				if len(itemidbase) == 50 || i == len(datas.ProductItemList.Item) {
					if i == len(datas.ProductItemList.Item) {
						//fmt.Println("Masuk GetItemListLimit if len(datas.ProductItemList.Item)")
					}
					if len(itemidbase) == 50 {
						//fmt.Println("Masuk GetItemListLimit if len(itemidbase) ")
					}
					//GetItemListBaseLoopArrayTEST(c, itemidbase)

					isi = GetItemListBaseLoopArray(itemidbase)

					for _, elementItemList := range isi.BaseItemInfoList.ItemList {
						// fmt.Println(strconv.Itoa(totalproduct) + " : " + strconv.Itoa(len(elementItemList.StockInfo)))
						// fmt.Println("***************************")
						// fmt.Println(elementItemList.HasModel)
						// fmt.Println(strconv.Itoa(int(elementItemList.ItemId)))
						// fmt.Println("***************************")
						if elementItemList.HasModel == true {
							//fmt.Println("HASHMODEL ADA " + strconv.Itoa(int(elementItemList.ItemId)))
							isiModel := GetModelItemLoop(c, elementItemList.ItemId)
							tokenService.SaveSkuMapping(isiModel, elementItemList.ItemId, elementItemList.ItemName, elementItemList.ItemStatus)

						} else {
							tokenService.SaveSkuMappingString(elementItemList.ItemSku, elementItemList.ItemId, elementItemList.ItemName, elementItemList.ItemStatus)
							//fmt.Println(elementItemList.ItemName)
						}

						//fmt.Println("SKU " + elementItemList.ItemSku + " Product :" + elementItemList.ItemName)
					}

					//var itemidbase []models.ArrayBaseItemShopee
					itemidbase = nil
					//fmt.Println(itemidbase)
					//fmt.Println(isi)
					//response.Result = isi
					//fmt.Println("LEN AKHIR itemidbase" + strconv.Itoa(len(itemidbase)))
				}
				//fmt.Println("LEN AKHIR BAWAH itemidbase" + strconv.Itoa(len(itemidbase)))

			}

			if datas.ProductItemList.HasNextPage == true {
				GetItemListLimitLoop(c, datas.ProductItemList.NextOffset, int64(totalproduct))

			}

			//fmt.Println("finish product shopee")
			// zaloraController.GetProductsAuto(c)
			// tiktokController.GetProductsTiktokAuto(c)
		} else {

			tokenService.CekError("shopee", datasx, "GetItemListLimit")
		}

	}

	// zaloraController.GetProductsAuto(c)
	// tiktokController.GetProductsTiktokAuto(c)
	// jdidController.GetProducts(c)
	// blibliController.GetProductBlibli(c)
	//tokpedController.GetProductsAuto() //test atas

	// c.JSON(http.StatusOK, response)
	// return
	fmt.Println("selesai product shopee " + time.Now().String())
	wg.Done()
}

func GetItemListLimit(c *gin.Context) {

	var response response.ResponseCrud

	var wg sync.WaitGroup
	wg.Add(8)
	//wg.Add(9)

	go func(c *gin.Context) {

		GetItemListLimitAuto(c, &wg)

	}(c)

	go func() {

		//tokpedController.GetProductsAuto(&wg)
		tokpedController.GetProductsAutoV2(&wg)

	}()

	go func() {

		zaloraController.GetProductsAuto(&wg)

	}()
	go func() {
		zaloraController.GetProductsAutoV2("inactive-all", &wg)

	}()

	go func() {

		tiktokController.GetProductsTiktokAuto(&wg)

	}()

	// go func() {

	// 	jdidController.GetProductsAuto(&wg)

	// }()

	go func() {

		blibliController.GetProductBlibliAuto(&wg)

	}()

	go func() {

		bukalapakController.GetProductsAuto(&wg)

	}()

	go func() {

		lazadaController.GetProductsAuto(&wg)

	}()

	wg.Wait()

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "SELESAI"

	//zaloraController.GetProductsAuto(c)
	// tiktokController.GetProductsTiktokAuto(c)
	// jdidController.GetProducts(c)
	// blibliController.GetProductBlibli(c)
	//tokpedController.GetProductsAuto() //test atas

	c.JSON(http.StatusOK, response)
	return
}

func GetItemListLimitLoop(c *gin.Context, offset int64, totalproduct int64) {

	//ItemId := c.Param("itemid")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_item_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-3600)*time.Hour).Unix(), 10) //5 bulan lalu
	//kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-1440)*time.Hour).Unix(), 10) //2 bulan lalu

	//offset = 0
	pagesize := os.Getenv("PAGE_SIZE_ORDER")
	ItemStatus := "NORMAL"
	urlshopee += "&update_time_from=" + kmrn
	urlshopee += "&update_time_to=" + skrg
	urlshopee += "&offset=" + strconv.Itoa(int(offset))
	urlshopee += "&page_size=" + pagesize
	urlshopee += "&item_status=" + ItemStatus

	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetItemListLimitLoop")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		//fmt.Println("sukses GetItemListLimitLoop")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonItemList(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas

		if datasx["message"] == "" {
			var itemidbase []models.ArrayBaseItemShopee
			i := 0
			//fmt.Println("Total data productItem " + strconv.Itoa(int(datas.ProductItemList.TotalCount)))
			for _, element1 := range datas.ProductItemList.Item {
				totalproduct++
				//fmt.Println(element1.ItemId)
				var isi models.BaseItemInfoHeader
				//fmt.Println(element1.ItemId)

				itemidbaseloop := models.ArrayBaseItemShopee{
					ItemId: element1.ItemId,
				}
				//fmt.Println("LEN ATAS itemidbase" + strconv.Itoa(len(itemidbase)))
				if len(itemidbase) < 50 {
					itemidbase = append(itemidbase, itemidbaseloop)
				}

				i++
				//fmt.Println("LEN itemidbase" + strconv.Itoa(len(itemidbase)))
				if len(itemidbase) == 50 || i == len(datas.ProductItemList.Item) {
					if i == len(datas.ProductItemList.Item) {
						//fmt.Println("Masuk GetItemListLimitLoop if len(datas.ProductItemList.Item)")
					}
					if len(itemidbase) == 50 {
						//fmt.Println("Masuk GetItemListLimitLoop if len(itemidbase) ")
					}

					isi = GetItemListBaseLoopArray(itemidbase)

					for _, elementItemList := range isi.BaseItemInfoList.ItemList {
						//fmt.Println(strconv.Itoa(int(totalproduct)) + " : " + strconv.Itoa(len(elementItemList.StockInfo)))
						//fmt.Println(elementItemList.HasModel)
						if elementItemList.HasModel == true {

							isiModel := GetModelItemLoop(c, elementItemList.ItemId)
							tokenService.SaveSkuMapping(isiModel, elementItemList.ItemId, elementItemList.ItemName, elementItemList.ItemStatus)

						} else {
							tokenService.SaveSkuMappingString(elementItemList.ItemSku, elementItemList.ItemId, elementItemList.ItemName, elementItemList.ItemStatus)
						}
						//fmt.Println(elementItemList.ItemSku)
					}

					//var itemidbase []models.ArrayBaseItemShopee
					itemidbase = nil
					//fmt.Println(itemidbase)
					//fmt.Println("LEN AKHIR itemidbase" + strconv.Itoa(len(itemidbase)))
					//fmt.Println(isi)
					//response.Result = isi
				}
				//fmt.Println("LEN AKHIR BAWAH itemidbase" + strconv.Itoa(len(itemidbase)))

			}

			if datas.ProductItemList.HasNextPage == true {
				//fmt.Println("masuk loop lagi")
				GetItemListLimitLoop(c, datas.ProductItemList.NextOffset, int64(totalproduct))

			}

		} else {
			//fmt.Println("GetItemListLimitLoop ada masalah")
			tokenService.CekError("shopee", datasx, "GetItemListLimitLoop")
		}

	}
}

func GetItemListBaseLoopArray(itemidbase []models.ArrayBaseItemShopee) models.BaseItemInfoHeader {
	var objDatas models.BaseItemInfoHeader
	ItemId := ""
	i := 0
	for _, element1 := range itemidbase {
		i++
		if len(itemidbase) == i {
			ItemId += strconv.Itoa(int(element1.ItemId))
		} else {
			ItemId += strconv.Itoa(int(element1.ItemId)) + ","
		}

	}
	//fmt.Println("itemid: " + ItemId)

	//ItemId := strconv.Itoa(int(itemid))

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_item_base_info"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	urlshopee += "&item_id_list=" + ItemId
	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get item base info")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		//fmt.Println("sukses get item base info")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		objDatas = parseJsonBaseItemArray(data)
		datasx := utils.GetByteToInterface(data)
		// response.ResponseCode = http.StatusOK
		// response.ResponseDesc = enums.SUCCESS
		// response.ResponseTime = utils.DateToStdNow()
		// response.Message = urlshopee
		// response.Result = datas
		if datasx["message"] == "" {
			//insert ke wms_sales_order
			//tokenService.SaveSalesOrder("shopee", datas)
		}
		//else {
		// 	tokenService.CekError("shopee", datasx, "GetOrderDetail")
		// }

	}

	return objDatas
}

func GetItemListBaseLoop(itemid int64) models.BaseItemInfoHeader {
	var objDatas models.BaseItemInfoHeader
	ItemId := strconv.Itoa(int(itemid))

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_item_base_info"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	urlshopee += "&item_id_list=" + ItemId
	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get item base info")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses get item base info")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		objDatas = parseJsonBaseItem(data)
		datasx := utils.GetByteToInterface(data)
		// response.ResponseCode = http.StatusOK
		// response.ResponseDesc = enums.SUCCESS
		// response.ResponseTime = utils.DateToStdNow()
		// response.Message = urlshopee
		// response.Result = datas
		if datasx["message"] == "" {
			//insert ke wms_sales_order
			//tokenService.SaveSalesOrder("shopee", datas)
		}
		//else {
		// 	tokenService.CekError("shopee", datasx, "GetOrderDetail")
		// }

	}

	return objDatas
}

func GetItemListBase(c *gin.Context) {

	ItemId := c.Param("itemid")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_item_base_info"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	urlshopee += "&item_id_list=" + ItemId
	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get item base info")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses get item base info")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonBaseItem(data)
		//datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
		response.Result = string(data)
		// fmt.Println("=================")
		// fmt.Println(datasx)
		// fmt.Println("=================")
		// if datasx["message"] == "" {
		// 	//insert ke wms_sales_order
		// 	tokenService.SaveSalesOrder("shopee", datas)
		// } else {
		// 	tokenService.CekError("shopee", datasx, "GetOrderDetail")
		// }

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetItemListLimitOLD(c *gin.Context) {

	//ItemId := c.Param("itemid")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_item_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	// kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-3600)*time.Hour).Unix(), 10) //5 bulan lalu
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-1440)*time.Hour).Unix(), 10) //2 bulan lalu

	offset := 0
	pagesize := os.Getenv("PAGE_SIZE_ORDER")
	ItemStatus := "NORMAL"
	urlshopee += "&update_time_from=" + kmrn
	urlshopee += "&update_time_to=" + skrg
	urlshopee += "&offset=" + strconv.Itoa(offset)
	urlshopee += "&page_size=" + pagesize
	urlshopee += "&item_status=" + ItemStatus

	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetItemListLimit")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses GetItemListLimit")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonItemList(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
		// fmt.Println("=================")
		// fmt.Println(datasx)
		// fmt.Println("=================")
		if datasx["message"] == "" {
			//fmt.Println(datas.ProductItemList.HasNextPage)
			//fmt.Println(len(datas.ProductItemList.Item))
			totalproduct := 0
			fmt.Println("Total data productItem " + strconv.Itoa(int(datas.ProductItemList.TotalCount)))
			for _, element1 := range datas.ProductItemList.Item {
				totalproduct++
				//fmt.Println(element1.ItemId)
				var isi models.BaseItemInfoHeader
				isi = GetItemListBaseLoop(element1.ItemId)
				//fmt.Println("***********************************")
				//fmt.Println(isi.BaseItemInfoList.ItemList.StockInfo)
				// fmt.Println(len(elementItemList.StockInfo))
				for _, elementItemList := range isi.BaseItemInfoList.ItemList {
					fmt.Println(strconv.Itoa(totalproduct) + " : " + strconv.Itoa(len(elementItemList.StockInfo)))
					fmt.Println(elementItemList.HasModel)
					if elementItemList.HasModel == true {

						isiModel := GetModelItemLoop(c, element1.ItemId)
						tokenService.SaveSkuMapping(isiModel, element1.ItemId, elementItemList.ItemName, elementItemList.ItemStatus)

						// for _, elementModel := range isiModel.ModelProductHeaderResp.ModelProductDetail {
						// 	//fmt.Println("Model SKU " + elementModel.ModelSku)
						// 	if len(elementModel.ModelSku) > 5 {
						// 		fmt.Println("Model SKU " + elementModel.ModelSku)
						// 	}
						// }

					} else {
						tokenService.SaveSkuMappingString(elementItemList.ItemSku, elementItemList.ItemId, elementItemList.ItemName, elementItemList.ItemStatus)
					}
					fmt.Println(elementItemList.ItemSku)
				}

				//fmt.Println("***********************************")
			}

			if datas.ProductItemList.HasNextPage == true {
				//fmt.Println("masuk loop awal")
				GetItemListLimitLoop(c, datas.ProductItemList.NextOffset, int64(totalproduct))

			}

			//fmt.Println("GetItemListLimit")
			//tokenService.SaveSalesOrder("shopee", datas)
		} else {
			//fmt.Println("GetItemListLimit ada masalah")
			//fmt.Println(datasx)
			tokenService.CekError("shopee", datasx, "GetItemListLimit")
		}

	}

	c.JSON(http.StatusOK, response)
	return
}
func GetItemListLimitLoopOLD(c *gin.Context, offset int64, totalproduct int64) {

	//ItemId := c.Param("itemid")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_item_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	// kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-3600)*time.Hour).Unix(), 10) //5 bulan lalu
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-1440)*time.Hour).Unix(), 10) //2 bulan lalu

	//offset = 0
	pagesize := os.Getenv("PAGE_SIZE_ORDER")
	ItemStatus := "NORMAL"
	urlshopee += "&update_time_from=" + kmrn
	urlshopee += "&update_time_to=" + skrg
	urlshopee += "&offset=" + strconv.Itoa(int(offset))
	urlshopee += "&page_size=" + pagesize
	urlshopee += "&item_status=" + ItemStatus

	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetItemListLimitLoop")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses GetItemListLimitLoop")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonItemList(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
		// fmt.Println("=================")
		// fmt.Println(datasx)
		// fmt.Println("=================")
		if datasx["message"] == "" {
			//fmt.Println(datas.ProductItemList.HasNextPage)
			//fmt.Println(len(datas.ProductItemList.Item))
			fmt.Println("Total data productItem " + strconv.Itoa(int(datas.ProductItemList.TotalCount)))
			for _, element1 := range datas.ProductItemList.Item {
				totalproduct++
				//fmt.Println(element1.ItemId)
				var isi models.BaseItemInfoHeader
				isi = GetItemListBaseLoop(element1.ItemId)
				//fmt.Println("***********************************")
				//fmt.Println(isi.BaseItemInfoList.ItemList.StockInfo)
				// fmt.Println(len(elementItemList.StockInfo))
				for _, elementItemList := range isi.BaseItemInfoList.ItemList {
					fmt.Println(strconv.Itoa(int(totalproduct)) + " : " + strconv.Itoa(len(elementItemList.StockInfo)))
					fmt.Println(elementItemList.HasModel)
					if elementItemList.HasModel == true {

						isiModel := GetModelItemLoop(c, element1.ItemId)
						tokenService.SaveSkuMapping(isiModel, element1.ItemId, elementItemList.ItemName, elementItemList.ItemStatus)

						// for _, elementModel := range isiModel.ModelProductHeaderResp.ModelProductDetail {
						// 	if len(elementModel.ModelSku) > 5 {
						// 		fmt.Println("Model SKU " + elementModel.ModelSku)
						// 		//insert ke wms_sku_mapping

						// 	}
						// }

					} else {
						tokenService.SaveSkuMappingString(elementItemList.ItemSku, elementItemList.ItemId, elementItemList.ItemName, elementItemList.ItemStatus)
					}

					//fmt.Println(elementItemList.ItemSku)
				}

				//fmt.Println("***********************************")
			}

			if datas.ProductItemList.HasNextPage == true {
				fmt.Println("masuk loop lagi")
				GetItemListLimitLoop(c, datas.ProductItemList.NextOffset, int64(totalproduct))

			}

			// fmt.Println("GetItemListLimit")
			//tokenService.SaveSalesOrder("shopee", datas)
		} else {
			fmt.Println("GetItemListLimitLoop ada masalah")
			//fmt.Println(datasx)
			tokenService.CekError("shopee", datasx, "GetItemListLimitLoop")
		}

	}
}

func parseJsonModelItem(jsonBuffer []byte) models.ModelProductHeader {

	ModelProductHeader := models.ModelProductHeader{}

	err := json.Unmarshal(jsonBuffer, &ModelProductHeader)
	if err != nil {
		return ModelProductHeader
	}

	// the array is now filled with users
	return ModelProductHeader

}
func GetModelItemLoop(c *gin.Context, ItemId int64) models.ModelProductHeader {

	var objModel models.ModelProductHeader

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_model_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)

	urlshopee += "&item_id=" + strconv.Itoa(int(ItemId))

	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetModelItemLoop")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		//fmt.Println("sukses GetModelItemLoop")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		objModel = parseJsonModelItem(data)
		//datasx := utils.GetByteToInterface(data)
		// response.ResponseCode = http.StatusOK
		// response.ResponseDesc = enums.SUCCESS
		// response.ResponseTime = utils.DateToStdNow()
		// response.Message = urlshopee
		// response.Result = datas

	}

	// c.JSON(http.StatusOK, response)
	return objModel
}

func GetModelItem(c *gin.Context) {

	ItemId := c.Param("itemid")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_model_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)

	urlshopee += "&item_id=" + ItemId

	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetModelItem")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses GetModelItemaaa")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonModelItem(data)
		//datasx := utils.GetByteToInterface(data)

		//itemidbaru, _ := strconv.Atoi(ItemId)
		//tokenService.SaveSkuMapping(datas, int64(itemidbaru), "INI")
		for _, elementModel := range datas.ModelProductHeaderResp.ModelProductDetail {
			fmt.Println("======================")
			fmt.Println(elementModel.ModelSku)
			//fmt.Println(len(elementModel.TierIndex))
			indextier := 0
			for _, elementTier := range elementModel.TierIndex {
				fmt.Println(datas.ModelProductHeaderResp.TierVariation[indextier].Name)
				tulis := (datas.ModelProductHeaderResp.TierVariation[indextier].OptionList[elementTier].Option)
				fmt.Println(tulis)
				indextier++
			}

			// fmt.Println("*****************")
			// for _, elementVariationHead := range datas.ModelProductHeaderResp.TierVariation {
			// 	fmt.Println("########################")
			// 	for _, elementVariationDetail := range elementVariationHead.OptionList {
			// 		fmt.Println(elementVariationDetail)
			// 	}

			// }

		}

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
		response.Result = string(data)
		// fmt.Println("=================")
		// fmt.Println(datasx)
		// fmt.Println("=================")
		// if datasx["message"] == "" {
		// 	//insert ke wms_sales_order
		// 	tokenService.SaveSalesOrder("shopee", datas)
		// } else {
		// 	tokenService.CekError("shopee", datasx, "GetOrderDetail")
		// }

	}

	c.JSON(http.StatusOK, response)
	return
}

func StockUpdateModelItem(c *gin.Context) {
	nmRoute := "StockUpdateModelItem"
	ItemId := c.Param("itemid")
	Modelid := c.Param("modelid")
	Stock := c.Param("stock")
	Sku := c.Param("sku")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/update_stock"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)

	jsonString := ""
	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	var objRest models.TableLogStock

	//02-09-2022
	jsonString = `{"item_id":` + ItemId + `,
	"stock_list":[
		{"model_id" : ` + Modelid + `,"normal_stock":` + Stock + `,"seller_stock":[{"stock":` + Stock + `}]}
		] }`

	if Modelid == "00" {
		jsonString = `{"item_id":` + ItemId + `,
		"stock_list":[
			{"model_id":0,"normal_stock":` + Stock + `,"seller_stock":[{"stock":` + Stock + `}]}
			] }`
	}

	objRest.UuidLog = uuid.New().String()
	objRest.ChannelCode = os.Getenv("KODE_SHOPEE")
	objRest.Sku = Sku
	objRest.Body = jsonString

	//cek reservedstok

	if s, err := strconv.ParseFloat(Stock, 64); err == nil {
		objRest.Stock = s
	}

	// jsonString = `{"item_id":` + ItemId + `,
	// "stock_list":[
	// 	{"model_id" : ` + Modelid + `,"normal_stock":` + Stock + `,
	// 	"seller_stock": [
	// 		{
	// 			"location_id": "-",
	// 			"stock": ` + Stock + `
	// 		}
	// 	]}
	// 	] }`

	// if Modelid == "00" {
	// 	jsonString = `{"item_id":` + ItemId + `,
	// 	"stock_list":[
	// 		{"normal_stock":` + Stock + `,"seller_stock": [
	// 			{
	// 				"location_id": "-",
	// 				"stock": ` + Stock + `
	// 			}
	// 		]}
	// 		] }`
	// }

	//fmt.Println(jsonString)
	var body_url = []byte(jsonString)
	req, err := http.NewRequest("POST", urlshopee, bytes.NewBuffer(body_url))

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal StockUpdateModelItem")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err
		helpers.KirimEmail(nmRoute, nmRoute+" Gagal Koneksi", "")
		objRest.Response = nmRoute + " Gagal Koneksi"
	} else {
		//fmt.Println("sukses StockUpdateModelItem")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonModelItem(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = datasx["message"]
		response.Result = datas
		objRest.Response = string(data)
		//fmt.Println("=================")
		//fmt.Println(datasx)
		//fmt.Println("=================")
		if datasx["message"] == "" {
			//insert ke wms_sales_order
			//tokenService.SaveSalesOrder("shopee", datas)
		} else {

			tokenService.CekError("shopee", datasx, "StockUpdateModelItem")
			tokenService.SaveErrorString("shopee", ItemId+"|"+Modelid, "StockUpdateModelItem")
			helpers.KirimEmail(nmRoute, jsonString+" | "+string(data), "")
			if datasx["message"] == "Product not found" {

				if Modelid != "00" {
					tokenService.UpdateProductNuse(ItemId, Modelid, os.Getenv("KODE_SHOPEE"))
				} else {
					tokenService.UpdateProductNuse(ItemId, "", os.Getenv("KODE_SHOPEE"))
				}

			}
		}

	}

	objRest.CreatedBy = "API"
	objRest.CreatedDate = time.Now()
	tokenRepository.SaveStockAPI(objRest)

	c.JSON(http.StatusOK, response)
	return
}
func parseJsonShipParam(jsonBuffer []byte) models.ShippingParamHeader {

	ShippingParamHeader := models.ShippingParamHeader{}

	err := json.Unmarshal(jsonBuffer, &ShippingParamHeader)
	if err != nil {
		return ShippingParamHeader
	}

	// the array is now filled with users
	return ShippingParamHeader

}

func parseJsonTrackShipParam(jsonBuffer []byte) models.TrackShippingParamHeader {

	TrackShippingParamHeader := models.TrackShippingParamHeader{}

	err := json.Unmarshal(jsonBuffer, &TrackShippingParamHeader)
	if err != nil {
		return TrackShippingParamHeader
	}

	// the array is now filled with users
	return TrackShippingParamHeader

}

func GetShippingParam(c *gin.Context) {
	NoOrder := c.Param("noorder")
	ObjToken := tokenService.FindToken("shopee")
	fmt.Println(NoOrder)
	var response response.ResponseCrud

	paths := "/api/v2/logistics/get_shipping_parameter"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	urlshopee += "&order_sn=" + NoOrder
	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get shipping param")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "GAGAL REQSHIPPING"
		response.Result = err

	} else {
		fmt.Println("sukses get shipping param")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonShipParam(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = datasx["message"]
		response.Result = datas
		fmt.Println("=================")
		fmt.Println(datasx)

		//fmt.Println(datas)
		fmt.Println("=================")
		if datasx["message"] == "" {
			// param_shipping1 := strconv.Itoa(int(datas.ShippingParamInfo.Pickup.AddressList[0].AddressId))
			// param_shipping2 := datas.ShippingParamInfo.Pickup.AddressList[0].TimeSlotList[0].PickupTimeId
			// fmt.Println(param_shipping1)
			// fmt.Println(param_shipping2)
			param_shipping1 := ""
			param_shipping2 := ""
			param := ""
			// fmt.Println("========SHIP========")
			// fmt.Println(datas.ShippingParamInfo.InfoNeeded)
			if len(datas.ShippingParamInfo.InfoNeeded.Pickup) > 0 {
				// param = "picking"
				// param_shipping2 = fmt.Sprintf("%v", datas.ShippingParamInfo.InfoNeeded.Pickup[1])
				for _, elementPickup := range datas.ShippingParamInfo.Pickup.AddressList {
					if elementPickup.AddressId == 86960452 {
						param = "picking"
						id_address_env := os.Getenv("ADDRESS_ID_SHOPEES")
						froms_shipping_env := os.Getenv("FROM_SHIPPING")
						if os.Getenv("ID_ALAMAT_SHOPEE") == strconv.Itoa(int(elementPickup.AddressId)) {
							fmt.Println("ALAMAT SHOPEE ENV SAMA ")
						}
						fmt.Println("==== parameter env ====")
						fmt.Println(os.Getenv("ID_ALAMAT_SHOPEE"))
						fmt.Println(id_address_env)
						fmt.Println(froms_shipping_env)
						fmt.Println(strconv.Itoa(int(elementPickup.AddressId)))
						fmt.Println("==== parameter env ====")
						param_shipping1 = strconv.Itoa(int(elementPickup.AddressId))
						if len(elementPickup.TimeSlotList) > 0 {
							param_shipping2 = elementPickup.TimeSlotList[0].PickupTimeId
						}

					}

				}
				// param_shipping1 = strconv.Itoa(int(datas.ShippingParamInfo.Pickup.AddressList[0].AddressId))
				// param_shipping2 = datas.ShippingParamInfo.Pickup.AddressList[0].TimeSlotList[0].PickupTimeId
			}

			// if len(datas.ShippingParamInfo.InfoNeeded.Dropoff) > 0 {
			// 	param = "dropoff"
			// 	param_shipping1 = fmt.Sprintf("%v", datas.ShippingParamInfo.InfoNeeded.Dropoff[0])
			// 	param_shipping2 = fmt.Sprintf("%v", datas.ShippingParamInfo.InfoNeeded.Dropoff[1])
			// }

			if param == "" { //masuk dropoff

				param_kurir := ""
				objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
				if objCek.ExpeditionType != "" {
					param_kurir = objCek.ExpeditionType
				}
				BalikanDrop := GetLogisticChannelOrderShip(param_kurir)

				fmt.Println("======DROPOFF======")
				// fmt.Println(BalikanDrop)
				if BalikanDrop.Message == "" {
					param = "dropoff"
					param_shipping1 = BalikanDrop.ResponseDesc
					param_shipping2 = "RAMAYANA ECOMMERCE"

					fmt.Println("======DROPOFF AMAN======")
					//fmt.Println(BalikanDrop)
					fmt.Println(BalikanDrop.ResponseDesc)
					response.Result = BalikanDrop.Result
				}

			}

			fmt.Println(param)
			fmt.Println(param_shipping1)
			fmt.Println(param_shipping2)

			// fmt.Println(param_shipping1)
			// fmt.Println(param_shipping2)

			respon := ShipOrder(c, NoOrder, param, param_shipping1, param_shipping2)
			if respon != "sukses" {
				response.ResponseDesc = enums.ERROR
				response.Message = respon + " (SHOPEE)"

				//cek jika cod pakai dropoff sicepat 10-01-2022
				param_kurir := ""
				objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
				if objCek.ExpeditionType != "" {
					param_kurir = objCek.ExpeditionType
				}
				BalikanDrop := GetLogisticChannelOrderShip(param_kurir)

				fmt.Println("======DROPOFF======")
				// fmt.Println(BalikanDrop)
				if BalikanDrop.Message == "" {
					param = "dropoff"
					param_shipping1 = BalikanDrop.ResponseDesc
					param_shipping2 = "RAMAYANA ECOMMERCE"

					fmt.Println("======DROPOFF AMAN======")
					//fmt.Println(BalikanDrop)
					fmt.Println(BalikanDrop.ResponseDesc)
					response.Result = BalikanDrop.Result
				}
				respon1 := ShipOrder(c, NoOrder, param, param_shipping1, param_shipping2)
				if respon1 != "sukses" {
					response.ResponseDesc = enums.ERROR
					response.Message = respon1 + " (SHOPEE)"
				}
				//end cek jika cod pakai dropoff sicepat 10-01-2022

			}
			//insert ke wms_sales_order
			//tokenService.SaveSalesOrder("shopee", datas)

		} else {
			tokenService.CekError("shopee", datasx, "GetShippingParam")
			response.ResponseDesc = enums.ERROR
			response.Message = fmt.Sprintf("%v", datasx["message"])
		}

		//else {
		// 	tokenService.CekError("shopee", datasx, "GetOrderDetail")
		// }

	}

	c.JSON(http.StatusOK, response)
	return
}

func ShipOrder(c *gin.Context, noorder string, param string, param_shipping1 string, param_shipping2 string) string {
	balikan := "gagal"
	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/logistics/ship_order"
	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)

	var jsonString = ""

	if param == "picking" {
		jsonString = `{"order_sn":"` + noorder + `","pickup":{ "address_id":` + param_shipping1 + `,"pickup_time_id":"` + param_shipping2 + `" }}`

	} else if param == "dropoff" {

		jsonString = `{"order_sn":"` + noorder + `","dropoff":{"branch_id":` + param_shipping1 + `,"sender_real_name":"` + param_shipping2 + `"} }`
	}
	fmt.Println(jsonString)
	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", urlshopee, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Gagal !")
		//log.Println("Error on response.\n[ERRO] -", err)
		fmt.Println("gagal get ship Order")

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datasx := utils.GetByteToInterface(data)
		fmt.Println("sukses get ship Order")
		fmt.Println(datasx)
		balikan = fmt.Sprintf("%v", datasx["message"])
		if datasx["message"] == "" {
			balikan = "sukses"
		} else {
			tokenService.CekError("shopee", datasx, "ShipOrder")
		}
		// response.ResponseCode = http.StatusOK
		// response.ResponseDesc = enums.SUCCESS
		// response.ResponseTime = utils.DateToStdNow()
		// response.Message = urlshopee
		// response.Result = string(data)

		//insert token
		//tokenService.SaveToken("shopee", datas, "RefreshToken")

	}

	return balikan
}

func GetShippingResi(c *gin.Context) {
	NoOrder := c.Param("noorder")
	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/logistics/get_tracking_number"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	urlshopee += "&order_sn=" + NoOrder
	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get shipping resi param")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses get shipping resi param")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonTrackShipParam(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = datasx["message"]
		//response.Result = datas
		fmt.Println("=================")
		fmt.Println(datasx["response"])

		fmt.Println(datas)
		fmt.Println("=================")
		if datasx["message"] == "" {
			response.ResponseDesc = datas.ResTrackShippingParamHeader.TrackingNumber
		} else {
			tokenService.CekError("shopee", datasx, "GetShippingResi")
			response.ResponseDesc = enums.ERROR
			response.Message = fmt.Sprintf("%v", datasx["message"])
		}

		//else {
		// 	tokenService.CekError("shopee", datasx, "GetOrderDetail")
		// }

	}

	c.JSON(http.StatusOK, response)
	return
}

func parseJsonChannelLogisticParam(jsonBuffer []byte) models.ChannelLogisticsHeader {

	ChannelLogisticsHeader := models.ChannelLogisticsHeader{}

	err := json.Unmarshal(jsonBuffer, &ChannelLogisticsHeader)
	if err != nil {
		return ChannelLogisticsHeader
	}

	// the array is now filled with users
	return ChannelLogisticsHeader

}

func GetLogisticChannelOrderShip(kurir string) response.ResponseCrud {
	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/logistics/get_channel_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetLogisticChannelOrderShip")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses GetLogisticChannelOrderShip")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonChannelLogisticParam(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = datasx["message"]
		response.Result = datas
		fmt.Println("=====GetLogisticChannelOrderShip======")
		//fmt.Println(datasx["response"])

		//fmt.Println(datas)
		fmt.Println("=====GetLogisticChannelOrderShip======")
		if datasx["message"] == "" {
			for _, value := range datas.LogisticsChannelList.DetailLogisticsChannelList {
				if value.LogisticsChannelName == kurir {
					response.ResponseDesc = strconv.Itoa(int(value.LogisticsChannelId))
				}

			}

		}
		// else {
		// 	tokenService.CekError("shopee", datasx, "GetShippingResi")
		// 	response.ResponseDesc = enums.ERROR
		// 	response.Message = fmt.Sprintf("%v", datasx["message"])
		// }

	}

	return response
}

func GetLogisticChannel(c *gin.Context) {
	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/logistics/get_channel_list"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetLogisticChannel")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses GetLogisticChannel")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonChannelLogisticParam(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = datasx["message"]
		response.Result = datas
		fmt.Println("=====GetLogisticChannel======")
		fmt.Println(datasx["response"])

		fmt.Println(datas)
		fmt.Println("=====GetLogisticChannel======")
		// if datasx["message"] == "" {
		// 	response.ResponseDesc = datas.ResTrackShippingParamHeader.TrackingNumber
		// } else {
		// 	tokenService.CekError("shopee", datasx, "GetShippingResi")
		// 	response.ResponseDesc = enums.ERROR
		// 	response.Message = fmt.Sprintf("%v", datasx["message"])
		// }

	}

	c.JSON(http.StatusOK, response)
	return
}

func StatusCompleted(c *gin.Context) {
	var response response.ResponseCrud

	ObjComplete := tokenService.CariOrderByResi("SHOPEE") //param by channel name
	if len(ObjComplete) > 0 {
		for _, value := range ObjComplete {
			//fmt.Println(value.NoOrder)
			GetOrderDetailProses(c, value.NoOrder)
		}
	}
	//cari

	zaloraController.GetProcessedOrderAuto("delivered")

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "DISINI"

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStockOtomatis(c *gin.Context) {
	var response response.ResponseCrud

	// ObjMapping := tokenService.CariSkuMappingObjGroup(os.Getenv("KODE_SHOPEE")) //param by channel name
	// //spew.Dump(ObjMapping)
	// if len(ObjMapping) > 0 {
	// 	for _, value := range ObjMapping {
	// 		helpers.UpdateStock(value.SkuNo, "API_CHANNEL", os.Getenv("KODE_SHOPEE"))

	// 	}
	// }

	// zaloraController.UpdateOtomatisStockZalora(c) //ZALORA
	// tiktokController.UpdateStockOtomatisTiktok(c) //TIKTOK

	SkuMapping, _ := tokenRepository.CariSkuMappingGroup()
	if len(SkuMapping) > 0 {
		for _, value := range SkuMapping {
			//time.Sleep(1 * time.Second)
			//helpers.UpdateStock(value.SkuNo, "API_CHANNEL", "")
			tokenService.GroupChannelSkuMap(value.SkuNo)

		}
	}

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "DISINISS"

	c.JSON(http.StatusOK, response)
	return
}

func parseJsonTrackresi(jsonBuffer []byte) models.TrackResiHeader {

	TrackResiHeader := models.TrackResiHeader{}

	err := json.Unmarshal(jsonBuffer, &TrackResiHeader)
	if err != nil {
		return TrackResiHeader
	}

	// the array is now filled with users
	return TrackResiHeader

}
func TrackingResi(c *gin.Context) {
	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud
	NoOrder := c.Param("noorder")

	paths := "/api/v2/logistics/get_tracking_info"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	urlshopee += "&order_sn=" + NoOrder
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal TrackingResi")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		fmt.Println("sukses TrackingResi")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonTrackresi(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = datasx["message"]
		response.Result = datas
		if datasx["message"] == "" {

			for key, val := range datas.ResTrackResiHeader.TrackingInfo {
				datas.ResTrackResiHeader.TrackingInfo[key].WaktuUpdate = time.Unix(val.UpdateTime, 0)
				//req.Header[key] = val
			}

		} else {
			tokenService.CekError("shopee", datasx, "TrackingResi")
			response.ResponseDesc = enums.ERROR
			response.Message = fmt.Sprintf("%v", datasx["message"])
		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func CekStatusProductNull(c *gin.Context) {

	ObjProduct := tokenService.SearchProductNull(os.Getenv("KODE_SHOPEE"))

	var response response.ResponseCrud

	for _, values := range ObjProduct {
		fmt.Println("disini " + values.IdSkuParent)
		GetItemProductNull(values.IdSkuParent)

	}

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = ""
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "sukses"
	response.Result = len(ObjProduct)

	c.JSON(http.StatusOK, response)
	return
}

func GetItemProductNull(ItemId string) {

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/get_item_base_info"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	urlshopee += "&item_id_list=" + ItemId
	// urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		//fmt.Println("gagal get GetItemProductNull")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		//fmt.Println("sukses get GetItemProductNull")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonBaseItem(data)

		for _, elementItemList := range datas.BaseItemInfoList.ItemList {

			tokenService.UpdateStatusSkuMappingShopeeNull(elementItemList.ItemId, elementItemList.ItemStatus)

		}
	}
}

// SHOPEE
type ResultStockShopee struct {
	Result ResultShopee `json:"result"`
}
type ResultShopee struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	//Response  ResponseDetailShopee `json:"response"`
	RequestId string `json:"request_id"`
}

func StatusProductNull(parent string) {

	returnchannel := ResultStockShopee{}

	urlshopee := os.Getenv("URL_API_SHOPEE") + "GetItemListBase/" + parent

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(data, &returnchannel)
		if err != nil {
		}

	}

}

func GetDetailPaymentOrderId(c *gin.Context) { //cek dari api-wms jika orderan ready to ship
	NoOrder := c.Param("noorder")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/payment/get_escrow_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetDetailPaymentOrderId")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonPaymentShopee(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas

		if datasx["message"] == "" {

		} else {
			tokenService.CekError("shopee", datasx, "GetDetailPaymentOrderId")
		}
		fmt.Println(string(data))

	}

	c.JSON(http.StatusOK, response)
	return
}

func OrderSalesShopee(c *gin.Context) {
	NoOrder := c.Param("noorder")
	var ObjOrderSales []models.TableOrderSalesShopee
	var ObjOrderSalesPayment models.TablePaymentSalesShopee
	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/order/get_order_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn_list=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)
	urlshopee += "&response_optional_fields=total_amount,buyer_username,item_list,recipient_address"
	//urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason,actual_shipping_fee,note,note_update_time,pay_time,dropshipper_phone,buyer_cancel_reason,actual_shipping_fee_confirmed"
	urlshopee += ",checkout_shipping_carrier,reverse_shipping_fee,invoice_data,package_list,pickup_done_time"

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get OrderSalesShopee")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJson(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
		//fmt.Println(time.Unix(datas.Response.OrderList[0].UpdateTime, 0))
		if datasx["message"] == "" {

			createdDate := time.Now()
			Resi := GetResiOrderSalesShopee(datas.Response.OrderList[0].OrderSn)
			objPayment := GetDetailPaymentSalesOrderShopee(datas.Response.OrderList[0].OrderSn)
			total_berat := 0
			total_item := 0
			for _, value := range datas.Response.OrderList[0].ItemList {
				var ObjOrderSalesLoop models.TableOrderSalesShopee
				ObjOrderSalesLoop.NoPesanan = datas.Response.OrderList[0].OrderSn
				ObjOrderSalesLoop.StatusPesanan = "Selesai"
				ObjOrderSalesLoop.StatusPembatalan = "Completed"
				ObjOrderSalesLoop.NoResi = Resi
				ObjOrderSalesLoop.OpsiPengiriman = datas.Response.OrderList[0].ShippingCarrier
				ObjOrderSalesLoop.AntarKe = "Pickup"
				ObjOrderSalesLoop.PesananDikirimSebelum = createdDate.AddDate(0, 0, +int(datas.Response.OrderList[0].DaysToShip))
				if datas.Response.OrderList[0].PickupDoneTime != 0 {
					ObjOrderSalesLoop.WaktuPengirimanDiatur = time.Unix(datas.Response.OrderList[0].PickupDoneTime, 0)
				}
				ObjOrderSalesLoop.WaktuPesananDibuat = time.Unix(datas.Response.OrderList[0].CreateTime, 0)
				ObjOrderSalesLoop.WaktuPembayaran = time.Unix(datas.Response.OrderList[0].PayTime, 0)
				ObjOrderSalesLoop.SkuInduk = value.ItemSku
				ObjOrderSalesLoop.NamaProduk = value.ItemName
				ObjOrderSalesLoop.NamaReferensiSku = value.ModelSku
				ObjOrderSalesLoop.NamaVariasi = value.ModelName

				for _, valuepay := range objPayment.PaymentDetailShopeeList.OrderIncome.Items {
					if valuepay.ItemId == value.ItemId && valuepay.ModelId == value.ModelId {
						// ObjOrderSalesLoop.HargaAwal = strconv.Itoa(int(value.ModelDiscountedPrice))
						// ObjOrderSalesLoop.HargaSetelahDiskon = strconv.Itoa(int(value.ModelDiscountedPrice))
						// ObjOrderSalesLoop.Jumlah = int16(value.ModelQuantityPurchased)
						// ObjOrderSalesLoop.TotalHargaProduk = strconv.Itoa(int(value.ModelDiscountedPrice))
						// ObjOrderSalesLoop.TotalDiskon = strconv.Itoa(int(value.ModelDiscountedPrice))

						totaldiskon := valuepay.OriginalPrice - valuepay.DiscountedPrice
						ObjOrderSalesLoop.HargaAwal = strconv.Itoa(int(valuepay.DiscountedPrice))
						ObjOrderSalesLoop.HargaSetelahDiskon = strconv.Itoa(int(valuepay.DiscountedPrice))
						ObjOrderSalesLoop.Jumlah = int16(valuepay.QuantityPurchased)
						ObjOrderSalesLoop.TotalHargaProduk = strconv.Itoa(int(valuepay.DiscountedPrice))
						ObjOrderSalesLoop.TotalDiskon = strconv.Itoa(int(totaldiskon))
						// ObjOrderSalesLoop.DiskonDariPenjual = strconv.Itoa(int(valuepay.DiscountedPrice))
						ObjOrderSalesLoop.DiskonDariPenjual = strconv.Itoa(int(totaldiskon))

						// ObjOrderSalesLoop.VoucherDitanggungPenjual = strconv.Itoa(int(valuepay.DiscountFromVoucherSeller))
						// ObjOrderSalesLoop.VoucherDitanggungShopee = strconv.Itoa(int(valuepay.DiscountFromVoucherShopee))
						ObjOrderSalesLoop.VoucherDitanggungPenjual = strconv.Itoa(int(objPayment.PaymentDetailShopeeList.OrderIncome.VoucherFromSeller)) //2203297A7CETRY
						ObjOrderSalesLoop.VoucherDitanggungShopee = strconv.Itoa(int(objPayment.PaymentDetailShopeeList.OrderIncome.VoucherFromShopee))  //2203297A7CETRY

					}

				}

				//ObjOrderSalesLoop.TotalDiskon = "belom"

				// ObjOrderSalesLoop.DiskonDariPenjual = fmt.Sprintf("%v", objPayment.PaymentDetailShopeeList.OrderIncome.SellerDiscount)
				ObjOrderSalesLoop.DiskonDariShopee = fmt.Sprintf("%v", objPayment.PaymentDetailShopeeList.OrderIncome.ShopeeDiscount)
				ObjOrderSalesLoop.BeratProduk = strconv.Itoa(int(value.Weight*1000)) + " gr"
				total_berat += int(value.Weight * 1000)
				total_item += int(value.ModelQuantityPurchased)
				ObjOrderSalesLoop.JmlProduk = int16(total_item)
				// ObjOrderSalesLoop.TotalBerat = strconv.Itoa(int(value.Weight*1000)) + " gr"
				ObjOrderSalesLoop.TotalBerat = strconv.Itoa(total_berat) + " gr"

				// ObjOrderSalesLoop.VoucherDitanggungPenjual = "belom"
				// ObjOrderSalesLoop.CashbackKoin = "belom"  seller_coin_cash_back
				ObjOrderSalesLoop.CashbackKoin = fmt.Sprintf("%v", objPayment.PaymentDetailShopeeList.OrderIncome.SellerCoinCashBack)
				// ObjOrderSalesLoop.VoucherDitanggungShopee = "belom"
				// ObjOrderSalesLoop.PotonganKoinShopee = 123
				ObjOrderSalesLoop.PotonganKoinShopee = int16(objPayment.PaymentDetailShopeeList.OrderIncome.Coins)

				ObjOrderSalesLoop.DiskonKartuKredit = 123
				ObjOrderSalesLoop.OngkirDibayarPembeli = int16(objPayment.PaymentDetailShopeeList.OrderIncome.EstimatedShippingFee)
				ObjOrderSalesLoop.TotalPembayaran = strconv.Itoa(int(datas.Response.OrderList[0].TotalAmount))
				ObjOrderSalesLoop.PerkiraanOngkir = strconv.Itoa(int(datas.Response.OrderList[0].ActualShippingFee))
				ObjOrderSalesLoop.UsernamePembeli = datas.Response.OrderList[0].BuyerUsername
				ObjOrderSalesLoop.NamaPenerima = "belom"
				ObjOrderSalesLoop.Tlpn = datas.Response.OrderList[0].RecipientAddress.Phone
				ObjOrderSalesLoop.AlamatPengiriman = datas.Response.OrderList[0].RecipientAddress.FullAddress
				ObjOrderSalesLoop.Kota = datas.Response.OrderList[0].RecipientAddress.City
				ObjOrderSalesLoop.Provinsi = datas.Response.OrderList[0].RecipientAddress.State
				ObjOrderSalesLoop.WaktuPesananSelesai = time.Unix(datas.Response.OrderList[0].UpdateTime, 0)

				for indexs, _ := range ObjOrderSales { // update total qty & total berat
					ObjOrderSales[indexs].JmlProduk = int16(total_item)
					ObjOrderSales[indexs].TotalBerat = strconv.Itoa(total_berat) + " gr"
				}

				ObjOrderSales = append(ObjOrderSales, ObjOrderSalesLoop)

				ObjOrderSalesPayment.IdPaymentSales = "belum"
				ObjOrderSalesPayment.OrderId = datas.Response.OrderList[0].OrderSn
				ObjOrderSalesPayment.Username = datas.Response.OrderList[0].BuyerUsername
				ObjOrderSalesPayment.OrderCrateingDate = time.Unix(datas.Response.OrderList[0].CreateTime, 0)
				ObjOrderSalesPayment.PayoutCompletedDate = time.Unix(datas.Response.OrderList[0].PayTime, 0)
				ObjOrderSalesPayment.OriginalProductPrice = int16(objPayment.PaymentDetailShopeeList.OrderIncome.OriginalPrice)

				discseller := 0
				if objPayment.PaymentDetailShopeeList.OrderIncome.SellerDiscount != 0 {
					discseller = -int(objPayment.PaymentDetailShopeeList.OrderIncome.SellerDiscount)
				}

				ObjOrderSalesPayment.SellerProductPromotion = int16(discseller)
				ObjOrderSalesPayment.ProductDiscountRebateShopee = int16(objPayment.PaymentDetailShopeeList.OrderIncome.ShopeeDiscount)
				ObjOrderSalesPayment.Voucher = 123
				ObjOrderSalesPayment.ShippingFeePaidByBuyer = int16(objPayment.PaymentDetailShopeeList.OrderIncome.BuyerPaidShippingFee)
				ObjOrderSalesPayment.ShippingFeeDiscountFrom3pl = int16(objPayment.PaymentDetailShopeeList.OrderIncome.ShippingFeeDiscountFrom3pl)
				ObjOrderSalesPayment.ShippingRebateFromShopee = int16(objPayment.PaymentDetailShopeeList.OrderIncome.ShopeeShippingRebate)
				shipbehalf := 0
				if objPayment.PaymentDetailShopeeList.OrderIncome.ActualShippingFee > 0 {
					shipbehalf = -int(objPayment.PaymentDetailShopeeList.OrderIncome.ActualShippingFee)
				}
				ObjOrderSalesPayment.ShippingFeePaidByShopeeOnYourBehalf = int16(shipbehalf)
				ObjOrderSalesPayment.RefundAmountToBuyer = int16(objPayment.PaymentDetailShopeeList.OrderIncome.SellerReturnRefund)

				commission := 0
				if objPayment.PaymentDetailShopeeList.OrderIncome.CommissionFee > 0 {
					commission = -int(objPayment.PaymentDetailShopeeList.OrderIncome.CommissionFee)
				}
				ObjOrderSalesPayment.CommissionFee = int16(commission)

				servicefee := 0
				if objPayment.PaymentDetailShopeeList.OrderIncome.ServiceFee > 0 {
					servicefee = -int(objPayment.PaymentDetailShopeeList.OrderIncome.ServiceFee)
				}
				ObjOrderSalesPayment.ServiceFee = int16(servicefee)

				ObjOrderSalesPayment.TotalReleasedAmount = int16(objPayment.PaymentDetailShopeeList.OrderIncome.EscrowAmount)
				ObjOrderSalesPayment.VoucherCode = "belum"
				ObjOrderSalesPayment.Code = "belum"
				ObjOrderSalesPayment.Refound = int16(objPayment.PaymentDetailShopeeList.OrderIncome.SellerReturnRefund)
				ObjOrderSalesPayment.ShippingFeePromotionBySeller = strconv.Itoa(int(objPayment.PaymentDetailShopeeList.OrderIncome.SellerShippingDiscount))
				ObjOrderSalesPayment.Status = 0

			}

			fmt.Println(ObjOrderSalesPayment)
			response.Result = ObjOrderSales
			response.Result = string(data)
		} else {
			tokenService.CekError("shopee", datasx, "OrderSalesShopee")
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetDetailPaymentShopee(c *gin.Context) {
	NoOrder := c.Param("noorder")

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("shopee")

	paths := "/api/v2/payment/get_escrow_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetDetailPaymentSalesOrderShopee")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonListEscrow(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
	}

	c.JSON(http.StatusOK, response)
	return
}

func GetDetailPaymentShopeeAuto(NoOrder string) models.DetailEscrowsShopee {

	var response response.ResponseCrud

	var objs models.DetailEscrowsShopee

	ObjToken := tokenService.FindToken("shopee")

	paths := "/api/v2/payment/get_escrow_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetDetailPaymentSalesOrderShopee")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonListEscrow(data)
		objs = datas

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
	}

	return objs
}

func GetDetailPaymentSalesOrderShopee(NoOrder string) models.PaymentDetailShopee {

	var objPaymentDetail models.PaymentDetailShopee

	ObjToken := tokenService.FindToken("shopee")

	paths := "/api/v2/payment/get_escrow_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetDetailPaymentSalesOrderShopee")

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonPaymentShopee(data)
		datasx := utils.GetByteToInterface(data)

		if datasx["message"] == "" {
			objPaymentDetail = datas

		} else {
			tokenService.CekError("shopee", datasx, "GetDetailPaymentSalesOrderShopee")
		}
	}

	return objPaymentDetail
}

func GetResiOrderSalesShopee(NoOrder string) string {
	balikan := ""
	ObjToken := tokenService.FindToken("shopee")

	paths := "/api/v2/logistics/get_tracking_number"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)
	urlshopee += "&order_sn=" + NoOrder

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetResiOrderSalesShopee")

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonTrackShipParam(data)
		datasx := utils.GetByteToInterface(data)

		if datasx["message"] == "" {
			balikan = datas.ResTrackShippingParamHeader.TrackingNumber
		} else {
			tokenService.CekError("shopee", datasx, "GetResiOrderSalesShopee")
		}

	}

	return balikan
}

func GetOrderDetailV2(c *gin.Context) {
	NoOrder := c.Param("noorder")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/order/get_order_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn_list=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)
	urlshopee += "&response_optional_fields=total_amount,buyer_username,item_list,recipient_address"
	//urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason,actual_shipping_fee,note,note_update_time,pay_time,dropshipper_phone,buyer_cancel_reason,actual_shipping_fee_confirmed"
	urlshopee += ",checkout_shipping_carrier,reverse_shipping_fee,invoice_data,package_list,pickup_done_time"

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetOrderDetailV2")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJson(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
		//fmt.Println(time.Unix(datas.Response.OrderList[0].UpdateTime, 0))
		if datasx["message"] == "" {
		} else {
			tokenService.CekError("shopee", datasx, "GetOrderDetailV2")
		}
		//fmt.Println(string(data))

	}

	c.JSON(http.StatusOK, response)
	return
}

func PaymentSalesShopee(c *gin.Context) {
	NoOrder := c.Param("noorder")
	var ObjOrderSalesPayment models.TablePaymentSalesShopee
	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/order/get_order_detail"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn_list=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)
	urlshopee += "&response_optional_fields=total_amount,buyer_username,item_list,recipient_address"
	//urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason"
	urlshopee += ",shipping_carrier,total_amount,cancel_by,cancel_reason,actual_shipping_fee,note,note_update_time,pay_time,dropshipper_phone,buyer_cancel_reason,actual_shipping_fee_confirmed"
	urlshopee += ",checkout_shipping_carrier,reverse_shipping_fee,invoice_data,package_list,pickup_done_time"

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal get PaymentSalesShopee")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJson(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas
		//fmt.Println(time.Unix(datas.Response.OrderList[0].UpdateTime, 0))
		if datasx["message"] == "" {

			objPayment := GetDetailPaymentSalesOrderShopee(datas.Response.OrderList[0].OrderSn)

			ObjOrderSalesPayment.IdPaymentSales = "belum"
			ObjOrderSalesPayment.OrderId = datas.Response.OrderList[0].OrderSn
			ObjOrderSalesPayment.Username = datas.Response.OrderList[0].BuyerUsername
			ObjOrderSalesPayment.OrderCrateingDate = time.Unix(datas.Response.OrderList[0].CreateTime, 0)
			ObjOrderSalesPayment.PayoutCompletedDate = time.Unix(datas.Response.OrderList[0].PayTime, 0)
			ObjOrderSalesPayment.OriginalProductPrice = int16(objPayment.PaymentDetailShopeeList.OrderIncome.OriginalPrice)

			discseller := 0
			if objPayment.PaymentDetailShopeeList.OrderIncome.SellerDiscount != 0 {
				discseller = -int(objPayment.PaymentDetailShopeeList.OrderIncome.SellerDiscount)
			}

			ObjOrderSalesPayment.SellerProductPromotion = int16(discseller)
			ObjOrderSalesPayment.ProductDiscountRebateShopee = int16(objPayment.PaymentDetailShopeeList.OrderIncome.ShopeeDiscount)
			ObjOrderSalesPayment.Voucher = 123
			ObjOrderSalesPayment.ShippingFeePaidByBuyer = int16(objPayment.PaymentDetailShopeeList.OrderIncome.BuyerPaidShippingFee)
			ObjOrderSalesPayment.ShippingFeeDiscountFrom3pl = int16(objPayment.PaymentDetailShopeeList.OrderIncome.ShippingFeeDiscountFrom3pl)
			ObjOrderSalesPayment.ShippingRebateFromShopee = int16(objPayment.PaymentDetailShopeeList.OrderIncome.ShopeeShippingRebate)
			shipbehalf := 0
			if objPayment.PaymentDetailShopeeList.OrderIncome.ActualShippingFee > 0 {
				shipbehalf = -int(objPayment.PaymentDetailShopeeList.OrderIncome.ActualShippingFee)
			}
			ObjOrderSalesPayment.ShippingFeePaidByShopeeOnYourBehalf = int16(shipbehalf)
			ObjOrderSalesPayment.RefundAmountToBuyer = int16(objPayment.PaymentDetailShopeeList.OrderIncome.SellerReturnRefund)

			commission := 0
			if objPayment.PaymentDetailShopeeList.OrderIncome.CommissionFee > 0 {
				commission = -int(objPayment.PaymentDetailShopeeList.OrderIncome.CommissionFee)
			}
			ObjOrderSalesPayment.CommissionFee = int16(commission)

			servicefee := 0
			if objPayment.PaymentDetailShopeeList.OrderIncome.ServiceFee > 0 {
				servicefee = -int(objPayment.PaymentDetailShopeeList.OrderIncome.ServiceFee)
			}
			ObjOrderSalesPayment.ServiceFee = int16(servicefee)

			ObjOrderSalesPayment.TotalReleasedAmount = int16(objPayment.PaymentDetailShopeeList.OrderIncome.EscrowAmount)
			ObjOrderSalesPayment.VoucherCode = "belum"
			ObjOrderSalesPayment.Code = "belum"
			ObjOrderSalesPayment.Refound = int16(objPayment.PaymentDetailShopeeList.OrderIncome.SellerReturnRefund)
			ObjOrderSalesPayment.ShippingFeePromotionBySeller = strconv.Itoa(int(objPayment.PaymentDetailShopeeList.OrderIncome.SellerShippingDiscount))
			ObjOrderSalesPayment.Status = 0

			response.Result = ObjOrderSalesPayment

		} else {
			tokenService.CekError("shopee", datasx, "PaymentSalesShopee")
		}

	}

	c.JSON(http.StatusOK, response)
	return

}

func UploadImageShopee(c *gin.Context) {
	//NoOrder := c.Param("path")

	var response response.ResponseCrud

	Filenya, _ := c.FormFile("path")
	NameFile := c.PostForm("nama")
	Username := c.PostForm("username")
	fmt.Println(Filenya.Filename)
	//path := enums.PATH_SERVER + filepath.Base(Filenya.Filename)
	path := enums.PATH_SERVER + NameFile + filepath.Ext(filepath.Base(Filenya.Filename))
	// path := enums.PATH_LOCAL + NameFile + filepath.Ext(filepath.Base(Filenya.Filename))
	fmt.Println(path)
	//path := enums.PATH_LOCAL
	c.SaveUploadedFile(Filenya, path)
	fmt.Println(path)
	data, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(path))
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(part, data)

	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	paths := "/api/v2/media_space/upload_image"

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, "", "")
	//RefreshTokenAuto(c)
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&sign=%s", partnerId, timest, sign)

	req, err := http.NewRequest("POST", urlshopee, body)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "application/octet-stream")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal UploadImageShopee")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonUploadImage(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		// response.Message = urlshopee
		response.Result = datas

		if datasx["message"] == "success" {
			//insert ke table
			tokenService.SaveUploadImage(datas.UploadImageShopeeDetail.ImageInfo, path, NameFile, Username)

		} else {
			response.ResponseDesc = enums.ERROR
			tokenService.CekError("shopee", datasx, "UplaodImageShopee")
		}
		//fmt.Println(string(data))

	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateSkuAllChannel(c *gin.Context) {
	var response response.ResponseCrud
	SkuNya := c.Param("sku")
	//helpers.UpdateStock(SkuNya, "API_CHANNEL", "")

	//cari group bby channel code
	tokenService.GroupChannelSkuMap(SkuNya)

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "SUKSES UPDATE SKU"
	c.JSON(http.StatusOK, response)
	return
}

func GetShipDocument(c *gin.Context) {
	var response response.ResponseCrud
	NoOrder := c.Query("noorder")
	document := tiktokController.GetDocumnetShipTiktokAuto(NoOrder)
	fmt.Println("Document Tiktok " + document)
	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "SUKSES " + NoOrder
	c.JSON(http.StatusOK, response)
	return
}

func DetailShipping(c *gin.Context) {
	NoOrder := c.Param("noorder")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/logistics/get_shipping_document_info"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s&order_sn=%s", partnerId, timest, refreshtoken, shop_id, sign, NoOrder)

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal DetailShipping")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJson(data)
		datasx := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = string(data)
		//fmt.Println(time.Unix(datas.Response.OrderList[0].UpdateTime, 0))
		if datasx["message"] == "" {
		} else {
			tokenService.CekError("shopee", datasx, "DetailShipping")
		}
		//fmt.Println(string(data))

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductBySKU(c *gin.Context) {
	Sku := c.Param("sku")

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/search_item"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)

	urlshopee += "&offset=0"
	urlshopee += "&page_size=10"
	urlshopee += "&item_sku=" + Sku

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetProductBySKU")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductSKU(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas

		for indexx, _ := range datas.Response.ItemIDList {
			var isi models.BaseItemInfoHeader

			IdItem := datas.Response.ItemIDList[indexx]

			isi = GetItemListBaseLoop(IdItem)
			for _, elementItemList := range isi.BaseItemInfoList.ItemList {
				fmt.Println(elementItemList.HasModel)
				if elementItemList.HasModel == true {

					isiModel := GetModelItemLoop(c, IdItem)
					tokenService.SaveSkuMapping(isiModel, IdItem, elementItemList.ItemName, elementItemList.ItemStatus)

				} else {
					tokenService.SaveSkuMappingString(elementItemList.ItemSku, elementItemList.ItemId, elementItemList.ItemName, elementItemList.ItemStatus)
				}

				//fmt.Println(elementItemList.ItemSku)
			}

		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductBySKUAuto(c *gin.Context, sku string) {

	ObjToken := tokenService.FindToken("shopee")

	var response response.ResponseCrud

	paths := "/api/v2/product/search_item"

	refreshtoken := fmt.Sprintf("%v", ObjToken.Value1)
	shop_id := os.Getenv("SHOP_ID_SHOPEE")

	host, partnerId, timest, sign, _, _ := AuthShopee(paths, refreshtoken, shop_id)

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&access_token=%s&shop_id=%s&sign=%s", partnerId, timest, refreshtoken, shop_id, sign)

	urlshopee += "&offset=0"
	urlshopee += "&page_size=10"
	urlshopee += "&item_sku=" + sku

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("gagal GetProductBySKUAuto")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductSKU(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		response.Result = datas

		for indexx, _ := range datas.Response.ItemIDList {
			var isi models.BaseItemInfoHeader

			IdItem := datas.Response.ItemIDList[indexx]

			isi = GetItemListBaseLoop(IdItem)
			for _, elementItemList := range isi.BaseItemInfoList.ItemList {
				fmt.Println(elementItemList.HasModel)
				if elementItemList.HasModel == true {

					isiModel := GetModelItemLoop(c, IdItem)
					tokenService.SaveSkuMapping(isiModel, IdItem, elementItemList.ItemName, elementItemList.ItemStatus)

				} else {
					tokenService.SaveSkuMappingString(elementItemList.ItemSku, elementItemList.ItemId, elementItemList.ItemName, elementItemList.ItemStatus)
				}

				//fmt.Println(elementItemList.ItemSku)
			}

		}

	}

}
