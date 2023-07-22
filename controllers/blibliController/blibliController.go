package blibliController

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/models"
	"github.com/rals/dearme-channel/models/response"
	"github.com/rals/dearme-channel/repositories/tokenRepository"
	"github.com/rals/dearme-channel/services/tokenService"
	"github.com/rals/dearme-channel/utils"
	// "encoding/json"
)

var ChannelName = "blibli"

func SignatureBlibli(methode string, bodynya string, typenya string, paths string) (string, string, string) {
	now := time.Now()
	timest := strconv.FormatInt(now.Unix(), 10)
	baseString := ""

	if bodynya != "" {
		hashMD5 := md5.Sum([]byte(bodynya))
		bodynya = hex.EncodeToString(hashMD5[:])
	}

	nowbaru := now.Format("2006 Jan 02 15:04:05 Mon")

	hari := ""
	jam := ""
	bln := ""
	nowObj := strings.Split(nowbaru, " ")
	if len(nowObj) > 0 {
		hari = nowObj[4]
		jam = nowObj[3]
		bln = nowObj[1]
	}

	tgl := hari + " " + bln + " " + strconv.Itoa(now.Day()) + " " + jam + " WIB " + strconv.Itoa(now.Year())

	bodynya += "\n"
	methode += "\n"
	typenya += "\n"
	tgl += "\n"

	baseString = fmt.Sprintf("%s%s%s%s%s", methode, bodynya, typenya, tgl, paths)
	param := os.Getenv("SIGNATURE_KEY_BLIBLI")

	key := []byte(param)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(baseString))

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), timest, baseString
}

func authBlibli() string {
	login := os.Getenv("CLIENT_ID_BLIBLI") + ":" + os.Getenv("CLIENT_KEY_BLIBLI")
	token := base64.StdEncoding.EncodeToString([]byte(login))
	return token
}

func parseJsonProductDetailBlibli(jsonBuffer []byte) models.ProductDetailBlibli {

	ProductDetailBlibli := models.ProductDetailBlibli{}

	err := json.Unmarshal(jsonBuffer, &ProductDetailBlibli)
	if err != nil {
		return ProductDetailBlibli
	}

	// the array is now filled with users
	return ProductDetailBlibli

}

func parseJsonPackageBlibli(jsonBuffer []byte) models.PackageHeadBlibli {

	PackageHeadBlibli := models.PackageHeadBlibli{}

	err := json.Unmarshal(jsonBuffer, &PackageHeadBlibli)
	if err != nil {
		return PackageHeadBlibli
	}

	// the array is now filled with users
	return PackageHeadBlibli

}

func parseJsonPriductBlibli(jsonBuffer []byte) models.ListProductBlibli {

	ListProductBlibli := models.ListProductBlibli{}

	err := json.Unmarshal(jsonBuffer, &ListProductBlibli)
	if err != nil {
		return ListProductBlibli
	}

	// the array is now filled with users
	return ListProductBlibli

}

func parseJsonOrdersBlibli(jsonBuffer []byte) models.ListOrdersBlibli {

	ListOrdersBlibli := models.ListOrdersBlibli{}

	err := json.Unmarshal(jsonBuffer, &ListOrdersBlibli)
	if err != nil {
		return ListOrdersBlibli
	}

	// the array is now filled with users
	return ListOrdersBlibli

}

func parseJsonOrdersDetailBlibli(jsonBuffer []byte) models.ListDetailOrderBlibli {

	ListDetailOrderBlibli := models.ListDetailOrderBlibli{}

	err := json.Unmarshal(jsonBuffer, &ListDetailOrderBlibli)
	if err != nil {
		return ListDetailOrderBlibli
	}

	// the array is now filled with users
	return ListDetailOrderBlibli

}

func parseJsonDetailAirwaybillBlibli(jsonBuffer []byte) models.OrderDetailAirwayBillBlibli {

	OrderDetailAirwayBillBlibli := models.OrderDetailAirwayBillBlibli{}

	err := json.Unmarshal(jsonBuffer, &OrderDetailAirwayBillBlibli)
	if err != nil {
		return OrderDetailAirwayBillBlibli
	}

	// the array is now filled with users
	return OrderDetailAirwayBillBlibli

}

func GetProductBlibliAuto(wg *sync.WaitGroup) {
	var response response.ResponseCrud
	fmt.Println("mulai product blibli " + time.Now().String())
	urlnya := "/proxy/mta/api/businesspartner/v2/product/getProductList"

	url := os.Getenv("URL_API_BLIBLI") + urlnya
	//max size =100

	sizePage := os.Getenv("LIMIT_PRODUCT_BLIBLI")
	numPage := 0
	bodynya := `
	{
		"size": ` + sizePage + `,
		"page": ` + strconv.Itoa(numPage) + `
	  }`

	var body_url = []byte(bodynya)

	//Sign, SignTime, BaseString := SignatureBlibli("POST", bodynya, "application/json", urlnya)

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&businessPartnerCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetProductBlibli")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		objDatas := parseJsonPriductBlibli(data)
		//datasx := utils.GetByteToInterface(data)
		index := 0
		if len(objDatas.Content) > 0 {

			for _, val := range objDatas.Content {
				index++
				//fmt.Println("cek detail produk " + val.GdnSku + " | " + val.MerchantSku)
				tokenService.SaveSkuMappingBlibli(val)

			}

			numPage++
			fmt.Println("masuk loop product blibli " + strconv.Itoa(numPage))
			GetProductLoopBlibli(index, sizePage, numPage)

		}

		response.Message = ""

		response.Result = objDatas

	}
	fmt.Println("selesai product blibli " + time.Now().String())
	wg.Done()

}

func GetProductBlibli(c *gin.Context) {
	var response response.ResponseCrud

	urlnya := "/proxy/mta/api/businesspartner/v2/product/getProductList"

	url := os.Getenv("URL_API_BLIBLI") + urlnya
	//max size =100

	sizePage := os.Getenv("LIMIT_PRODUCT_BLIBLI")
	numPage := 0
	bodynya := `
	{
		"size": ` + sizePage + `,
		"page": ` + strconv.Itoa(numPage) + `
	  }`

	var body_url = []byte(bodynya)

	//Sign, SignTime, BaseString := SignatureBlibli("POST", bodynya, "application/json", urlnya)

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&businessPartnerCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetProductBlibli")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		objDatas := parseJsonPriductBlibli(data)
		//datasx := utils.GetByteToInterface(data)
		index := 0
		if len(objDatas.Content) > 0 {

			for _, val := range objDatas.Content {
				index++
				//fmt.Println("cek detail produk " + val.GdnSku + " | " + val.MerchantSku)
				tokenService.SaveSkuMappingBlibli(val)

			}

			numPage++
			fmt.Println("masuk loop product blibli " + strconv.Itoa(numPage))
			GetProductLoopBlibli(index, sizePage, numPage)

		}

		response.Message = ""

		response.Result = objDatas

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductLoopBlibli(index int, sizePage string, numPage int) {
	var response response.ResponseCrud

	urlnya := "/proxy/mta/api/businesspartner/v2/product/getProductList"

	url := os.Getenv("URL_API_BLIBLI") + urlnya
	//max size =100

	bodynya := `
	{
		"size": ` + sizePage + `,
		"page": ` + strconv.Itoa(numPage) + `
	  }`

	var body_url = []byte(bodynya)

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&businessPartnerCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetProductLoopBlibli")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		objDatas := parseJsonPriductBlibli(data)
		//datasx := utils.GetByteToInterface(data)

		if len(objDatas.Content) > 0 {

			for _, val := range objDatas.Content {
				index++
				//fmt.Println("cek detail produk " + val.GdnSku + " | " + val.MerchantSku)
				tokenService.SaveSkuMappingBlibli(val)
			}

			numPage++
			fmt.Println("masuk loop product blibli " + strconv.Itoa(numPage))
			GetProductLoopBlibli(index, sizePage, numPage)

		}

		response.Message = ""

		response.Result = objDatas

	}

}

func GetProductBlibliDetail(c *gin.Context) {
	var response response.ResponseCrud
	SkuBlibli := c.Param("skublibli")

	urlnya := "/proxy/mta/api/businesspartner/v1/product/detailProduct"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&businessPartnerCode=" + os.Getenv("PARTNER_ID_BLIBLI")
	url += "&gdnSku=" + SkuBlibli
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetProductBlibliDetail")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		objDatas := parseJsonProductDetailBlibli(data)

		response.Message = ""

		response.Result = objDatas

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductBlibliDetailV2(c *gin.Context) {
	var response response.ResponseCrud
	nmroute := "GetProductBlibliDetailV2"
	SkuBlibli := c.Param("skublibli")

	urlnya := "/proxy/seller/v1/products/" + SkuBlibli

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&product-sku=" + SkuBlibli
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&storeId=10001"
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal " + nmroute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		//objDatas := parseJsonProductDetailBlibli(data)

		response.Message = ""

		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetBrandBlibli(c *gin.Context) {
	var response response.ResponseCrud
	urlnya := "/proxy/sas/product-approve"
	urlnya = "/proxy/mta/api/businesspartner/v2/product/getBrands"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&businessPartnerCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetProductBlibli")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		response.Message = ""
		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrderBlibli(c *gin.Context) {
	var response response.ResponseCrud
	status := c.Param("status")

	urlnya := "/proxy/seller/v1/orders/packages/filter"
	//urlnya = "/proxy/mta/api/businesspartner/v1/order/getReturnedOrderSummary"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()

	url += "&storeId=10001"
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari lalu
	PageSize := 50
	NumPage := 0

	bodynya := `
	{
		"filter": {
			"orderItemStatuses": [
				"` + status + `",
				"X",
				"D"
			  ],
			  "statusFPDateRange": {
				"end": ` + skrg + `000,
				"start": ` + kmrn + `000
			  }
		},
		"paging": {
		  "page": ` + strconv.Itoa(NumPage) + `,
		  "size": ` + strconv.Itoa(PageSize) + `
		}
	  }
	  `

	var body_url = []byte(bodynya)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetOrderBlibli")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersBlibli(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//fmt.Println(string(data))
		response.Message = url
		response.Result = string(data)
		if datas.Success == true {
			response.Message = ""
		}
		response.Result = datas

		if len(datas.Content) > 0 {
			//save ke database
			//fmt.Println(datas.Content)
			//tokenService.SaveOrderObjBlibli(datas.Content)

			for _, val := range datas.Content {
				//fmt.Println(val.PackageId)
				NoOrderTemp := ""
				for _, valOrderItems := range val.OrderItems {
					// tokenService.SaveOrderObjBlibli(datas.Content)
					//GetOrderDetailAuto(valOrderItems.Order.ItemId)
					//fmt.Println(index)
					if NoOrderTemp != valOrderItems.Order.Id {
						NoOrderTemp = valOrderItems.Order.Id
						//fmt.Println(NoOrderTemp)
						//fmt.Println(valOrderItems)
						//fmt.Println("============================")
						objCek, _ := tokenRepository.FindSalesOrder(NoOrderTemp)
						if objCek.NoOrder == "" {
							//if NoOrderTemp == "12103467292" {
							objCust := GetOrderDetailAuto(valOrderItems.Order.ItemId)

							if objCust.Content.Id != "" {
								tokenService.SaveOrderObjBlibli(val.OrderItems, objCust)
							}

							//}

						}

					}

				}

			}

			NumPage++
			GetOrderLoopBlibli(status, NumPage)
		}

		//fmt.Println(string(data))
	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrderBlibliAuto(status string) {
	var response response.ResponseCrud
	fmt.Println("GetOrderBlibliAuto " + status)
	urlnya := "/proxy/seller/v1/orders/packages/filter"
	//urlnya = "/proxy/mta/api/businesspartner/v1/order/getReturnedOrderSummary"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()

	url += "&storeId=10001"
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari lalu
	PageSize := 50
	NumPage := 0

	bodynya := `
	{
		"filter": {
			"orderItemStatuses": [
				"` + status + `",
				"X",
				"D"
			  ],
			  "statusFPDateRange": {
				"end": ` + skrg + `000,
				"start": ` + kmrn + `000
			  }
		},
		"paging": {
		  "page": ` + strconv.Itoa(NumPage) + `,
		  "size": ` + strconv.Itoa(PageSize) + `
		}
	  }
	  `

	var body_url = []byte(bodynya)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetOrderBlibli")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersBlibli(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//fmt.Println(string(data))
		response.Message = url
		response.Result = string(data)
		if datas.Success == true {
			response.Message = ""
		}
		response.Result = datas

		if len(datas.Content) > 0 {
			//save ke database
			//fmt.Println(datas.Content)
			//tokenService.SaveOrderObjBlibli(datas.Content)

			for _, val := range datas.Content {
				//fmt.Println(val.PackageId)
				NoOrderTemp := ""
				for _, valOrderItems := range val.OrderItems {
					// tokenService.SaveOrderObjBlibli(datas.Content)
					//GetOrderDetailAuto(valOrderItems.Order.ItemId)
					//fmt.Println(index)
					if NoOrderTemp != valOrderItems.Order.Id {
						NoOrderTemp = valOrderItems.Order.Id
						//fmt.Println(NoOrderTemp)
						//fmt.Println(valOrderItems)
						//fmt.Println("============================")
						objCek, _ := tokenRepository.FindSalesOrder(NoOrderTemp)
						if objCek.NoOrder == "" {
							//if NoOrderTemp == "12103467292" {
							objCust := GetOrderDetailAuto(valOrderItems.Order.ItemId)

							if objCust.Content.Id != "" {
								tokenService.SaveOrderObjBlibli(val.OrderItems, objCust)
							}

							//}

						} else if objCek.StatusProcessOrder == "0" && valOrderItems.Order.ItemStatus == "X" {
							var objCust models.ListDetailOrderBlibli
							tokenService.SaveOrderObjBlibli(val.OrderItems, objCust)
						}

					}

				}

			}

			NumPage++
			GetOrderLoopBlibli(status, NumPage)
		}

		//fmt.Println(string(data))
	}

}

func GetOrderLoopBlibli(status string, NumPage int) {
	var response response.ResponseCrud

	urlnya := "/proxy/seller/v1/orders/packages/filter"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()

	url += "&storeId=10001"
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari lalu
	PageSize := 50

	bodynya := `
	{
		"filter": {
			"orderItemStatuses": [
				"` + status + `",
				"X",
				"D"
			  ],
			  "statusFPDateRange": {
				"end": ` + skrg + `000,
				"start": ` + kmrn + `000
			  }
		},
		"paging": {
		  "page": ` + strconv.Itoa(NumPage) + `,
		  "size": ` + strconv.Itoa(PageSize) + `
		}
	  }
	  `

	var body_url = []byte(bodynya)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetOrderBlibli")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersBlibli(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//fmt.Println(string(data))
		response.Message = url
		response.Result = string(data)
		if datas.Success == true {
			response.Message = ""
		}
		response.Result = datas

		if len(datas.Content) > 0 {
			//save ke database
			//fmt.Println(datas.Content)
			//tokenService.SaveOrderObjBlibli(datas.Content)

			for _, val := range datas.Content {
				//fmt.Println(val.PackageId)
				NoOrderTemp := ""
				for _, valOrderItems := range val.OrderItems {
					// tokenService.SaveOrderObjBlibli(datas.Content)
					//GetOrderDetailAuto(valOrderItems.Order.ItemId)
					//fmt.Println(index)
					if NoOrderTemp != valOrderItems.Order.Id {
						NoOrderTemp = valOrderItems.Order.Id
						//fmt.Println(NoOrderTemp)
						//fmt.Println(valOrderItems)
						//fmt.Println("============================")
						objCek, _ := tokenRepository.FindSalesOrder(NoOrderTemp)
						if objCek.NoOrder == "" {
							//if NoOrderTemp == "12103467292" {
							objCust := GetOrderDetailAuto(valOrderItems.Order.ItemId)

							if objCust.Content.Id != "" {
								tokenService.SaveOrderObjBlibli(val.OrderItems, objCust)
							}

							//}

						} else if objCek.StatusProcessOrder == "0" && valOrderItems.Order.ItemStatus == "X" {
							var objCust models.ListDetailOrderBlibli
							tokenService.SaveOrderObjBlibli(val.OrderItems, objCust)
						}

					}

				}

				NumPage++
				GetOrderLoopBlibli(status, NumPage)
			}

		}

	}

}

func GetOrderDetailAuto(id string) models.ListDetailOrderBlibli {
	var response response.ResponseCrud

	var obj models.ListDetailOrderBlibli

	NoOrder := id
	urlnya := "/proxy/seller/v1/orders/items/" + NoOrder

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()
	url += "&order-item-id=" + NoOrder
	url += "&storeId=" + os.Getenv("STORE_ID_BLIBLI")
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetOrderDetailAuto")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersDetailBlibli(data)
		obj = datas
	}

	return obj
}

func GetOrderDetail(c *gin.Context) {
	var response response.ResponseCrud
	NoOrder := c.Param("noorder")
	nmRoute := "GetOrderDetail"
	NoOrderObj := strings.Split(NoOrder, "-")

	if len(NoOrderObj) > 0 {
		NoOrder = NoOrderObj[1]
	}

	urlnya := "/proxy/seller/v1/orders/items/" + NoOrder

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()
	url += "&order-item-id=" + NoOrder

	url += "&storeId=10001"
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetOrderDetailBlibli")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "ERROR GET DETAIL ORDER BLIBLI. (CONNECTION)"
		tokenService.SaveErrorString(ChannelName, "GAGAL KONEK", nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersDetailBlibli(data)
		response.ResponseCode = http.StatusOK
		//response.ResponseDesc = enums.SUCCESS
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		response.Message = ""
		response.Result = string(data)
		//response.Result = datas
		if datas.ErrorMessage != "" {
			tokenService.SaveErrorString(ChannelName, datas.ErrorMessage, nmRoute)
			response.Message = datas.ErrorMessage
			response.Result = datas
			response.ResponseDesc = datas.ErrorCode
		} else {
			response.ResponseDesc = datas.Content.Status
		}

	}

	c.JSON(http.StatusOK, response)
	return

}

func GetOrderDetailAirwayBill(c *gin.Context) {
	var response response.ResponseCrud
	nmRoute := "GetOrderDetailAirwayBill"
	NoOrder := c.Param("noorder")

	urlnya := "/proxy/mta/api/businesspartner/v1/order/orderDetail"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	NoOrderObj := strings.Split(NoOrder, "-")

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()
	url += "&orderNo=" + NoOrderObj[0]
	url += "&orderItemNo=" + NoOrderObj[1]

	url += "&storeId=" + os.Getenv("STORE_ID_BLIBLI")
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal GetOrderDetailAirwayBill")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "ERROR GET DETAIL ORDER BLIBLI (CONNECTION)"
		tokenService.SaveErrorString(ChannelName, "GAGAL KONEK", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonDetailAirwaybillBlibli(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.ResponseDesc = ""
		fmt.Println(string(data))

		NoResinya := datas.Value.AwbNumber
		Kurirnya := datas.Value.LogisticsService

		if NoResinya != "" && Kurirnya != "" {
			response.ResponseDesc = NoResinya + "^" + Kurirnya
		}

		response.Message = ""
		response.Result = string(data)
		if datas.ErrorMessage != "" {
			response.Message = datas.ErrorMessage
			response.Result = datas
			if response.Message == "" {
				response.Message = "ERROR GET DETAIL ORDER BLIBLI"
				tokenService.SaveErrorString(ChannelName, "ERROR GET DETAIL ORDER BLIBLI", nmRoute)
			} else {
				response.Message = datas.ErrorMessage + " (BLIBLI)"
				tokenService.SaveErrorString(ChannelName, datas.ErrorMessage, nmRoute)
			}

		}

	}

	c.JSON(http.StatusOK, response)
	return

}

func OrderCombineBlibli(c *gin.Context) {
	nmRoute := "OrderCombineBlibli"
	var response response.ResponseCrud
	NoOrder := c.Param("noorder")

	NoOrderObj := strings.Split(NoOrder, "-")

	urlnya := "/proxy/mta/api/businesspartner/v1/order/getCombineShipping"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()
	url += "&storeId=" + os.Getenv("STORE_ID_BLIBLI")
	url += "&businessPartnerCode=" + os.Getenv("PARTNER_ID_BLIBLI")
	url += "&orderItemNo=" + NoOrderObj[0]

	//fmt.Println(url)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "Req Combine BLIBLI GAGAL (connection)"
		tokenService.SaveErrorString(ChannelName, "GAGAL KONEK", nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersBlibli(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Result = string(data)
		if datas.ErrorMessage != "" {
			tokenService.SaveErrorString(ChannelName, datas.ErrorMessage, nmRoute)
			response.Message = datas.ErrorMessage + " (BLIBLI)"
			fmt.Println(string(data))
		}

	}

	c.JSON(http.StatusOK, response)
	return

}

func CreatePackageBlibli(c *gin.Context) {
	nmRoute := "CreatePackageBlibli"
	var response response.ResponseCrud
	NoOrder := c.Param("noorder")
	IdItem := ""
	NoOrderObj := strings.Split(NoOrder, "-")

	if len(NoOrderObj) > 0 {
		IdItem = NoOrderObj[1]
	}

	//gabungkan iditem jika lebih dari satu

	objOrder, _ := tokenRepository.FindSalesOrderArray(NoOrder)

	if len(objOrder) > 1 {

		indexItem, _ := strconv.Atoi(IdItem)
		for indexLoop, _ := range objOrder {
			//fmt.Println(indexLoop)
			if indexLoop > 0 {
				IdItem += `","` + strconv.Itoa(indexItem)
			}
			indexItem++

		}

	}

	//cari

	//urlnya := "/proxy/mta/api/businesspartner/v1/order/getCombineShipping/" + NoOrder
	urlnya := "/proxy/mta/api/businesspartner/v1/order/createPackage"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()
	url += "&storeId=" + os.Getenv("STORE_ID_BLIBLI")
	url += "&businessPartnerCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	bodynya := `
	{
		"orderItemIds": [
		  "` + IdItem + `"
		]
	  }
	  `

	fmt.Println(nmRoute + " BLIBLI")
	fmt.Println(bodynya)

	var body_url = []byte(bodynya)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal OrderCombineBlibli")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "Req Package BLIBLI GAGAL (connection)"
		tokenService.SaveErrorString(ChannelName, "GAGAL KONEK", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonPackageBlibli(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		response.Message = ""
		response.Result = string(data)

		if datas.ErrorMessage != "" {
			response.Message = datas.ErrorMessage + " (BLIBLI)"
			tokenService.SaveErrorString(ChannelName, datas.ErrorMessage, nmRoute)
		} else {
			// request picking jika bukan dropship

			PackageNya := datas.Value.PackageId

			var objUpdateKurir models.TableSalesOrder
			objUpdateKurir.NoOrder = NoOrder
			objUpdateKurir.ExpeditionType = PackageNya

			result := tokenRepository.UpdateKurir(objUpdateKurir)
			if result != nil {
				fmt.Println("UpdateKurir Buat Package Blibli GAGAL")
			} else {
				fmt.Println("UpdateKurir Buat Package Blibli SUKSES")
			}

			objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
			fmt.Println(objCek)
			if objCek.DeliveryType != "DROPSHIP" {
				//lanjut request picking
				cekShip := ReqShippingBlibli(PackageNya)
				if cekShip != "" {
					response.Message = cekShip
				}
			}

		}

	}

	c.JSON(http.StatusOK, response)
	return

}

func ReqShippingBlibli(PackageId string) string {
	nmRoute := "ReqShippingBlibli"
	var response response.ResponseCrud
	pesan := ""
	urlnya := "/proxy/seller/v1/orders/regular/" + PackageId + "/fulfill"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()
	url += "&storeId=" + os.Getenv("STORE_ID_BLIBLI")
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")
	url += "&package-id=" + PackageId

	bodynya := `
	{}
	  `

	var body_url = []byte(bodynya)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		//response.Message = url
		pesan = "Request Pickup BLIBLI GAGAL (connection)"
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonPackageBlibli(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Result = string(data)
		fmt.Println("==========" + nmRoute + "==========")
		fmt.Println(string(data))
		if datas.ErrorMessage != "" {
			tokenService.SaveErrorString(ChannelName, datas.ErrorMessage, nmRoute)
			pesan = datas.ErrorMessage + " (BLIBLI)"
		}

	}

	return pesan

}

func ReqShippingBlibliTest(c *gin.Context) {
	nmRoute := "ReqShippingBlibliTest"
	PackageId := c.Param("package")
	var response response.ResponseCrud
	pesan := ""
	urlnya := "/proxy/seller/v1/orders/regular/" + PackageId + "/fulfill"
	//urlnya := "/proxy/seller/v1/orders/shipping-by-seller/" + PackageId + "/ready-to-ship"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()
	url += "&storeId=" + os.Getenv("STORE_ID_BLIBLI")
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")
	url += "&package-id=" + PackageId

	bodynya := `
	{}
	  `

	var body_url = []byte(bodynya)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		//response.Message = url
		pesan = "Request Pickup BLIBLI GAGAL (connection)"
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonPackageBlibli(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Result = string(data)
		fmt.Println("==========" + nmRoute + "==========")
		fmt.Println(string(data))
		if datas.ErrorMessage != "" {
			tokenService.SaveErrorString(ChannelName, datas.ErrorMessage, nmRoute)
			pesan = datas.ErrorMessage + " (BLIBLI)"
		}

	}

	fmt.Println(pesan)
	c.JSON(http.StatusOK, response)
	return

}

func UpdateStockBlibli(c *gin.Context) {
	nmRoute := "UpdateStockBlibli"
	var response response.ResponseCrud
	SkuBlibli := c.Param("skublibli")
	Stocknya := c.Param("stock")

	urlnya := "/proxy/seller/v1/products/" + SkuBlibli + "/stock"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()
	url += "&storeId=" + os.Getenv("STORE_ID_BLIBLI")
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&blibli-sku=" + SkuBlibli
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	bodynya := `
	{
		"availableStock": ` + Stocknya + `
	  }
	  `

	var body_url = []byte(bodynya)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "Update Stock BLIBLI GAGAL (connection)"
		tokenService.SaveErrorString(ChannelName, "GAGAL KONEK", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonPackageBlibli(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		if datas.ErrorMessage != "" {

			tokenService.SaveErrorString(ChannelName, datas.ErrorCode+" | "+datas.ErrorMessage, nmRoute)
			fmt.Println(string(data))
			response.Message = datas.ErrorMessage + " (BLBILI)"

			if datas.ErrorCode == "LIMIT_EXCEEDED" {
				//tunggu 1 menit hit lagi
				time.Sleep(60 * time.Second)
				UpdateStockBlibliLoop(SkuBlibli, Stocknya)

			}
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStockBlibliLoop(skublibli string, stock string) {
	nmRoute := "UpdateStockBlibliLoop"
	var response response.ResponseCrud
	SkuBlibli := skublibli
	Stocknya := stock

	urlnya := "/proxy/seller/v1/products/" + SkuBlibli + "/stock"

	url := os.Getenv("URL_API_BLIBLI") + urlnya

	token := authBlibli()

	url += "?channelId=" + os.Getenv("ID_TOKO_BLIBLI")
	url += "&requestId=" + os.Getenv("ID_TOKO_BLIBLI") + "-" + uuid.New().String()
	url += "&storeId=" + os.Getenv("STORE_ID_BLIBLI")
	url += "&username=" + os.Getenv("USER_SELLER_BLIBLI")
	url += "&blibli-sku=" + SkuBlibli
	url += "&storeCode=" + os.Getenv("PARTNER_ID_BLIBLI")

	bodynya := `
	{
		"availableStock": ` + Stocknya + `
	  }
	  `

	var body_url = []byte(bodynya)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Seller-Key", os.Getenv("SELLER_KEY_BLIBLI"))

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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "Update Stock BLIBLI GAGAL (connection)"
		tokenService.SaveErrorString(ChannelName, "GAGAL KONEK", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonPackageBlibli(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		if datas.ErrorMessage != "" {
			tokenService.SaveErrorString(ChannelName, datas.ErrorMessage, nmRoute)
			fmt.Println(string(data))
			response.Message = datas.ErrorMessage + " (BLBILI)"

			if datas.ErrorCode == "LIMIT_EXCEEDED" {
				//tunggu 1 menit hit lagi
				//time.Sleep(60 * time.Second)
				//UpdateStockBlibliLoop(SkuBlibli, Stocknya)

			}
		}

	}

}

type IP struct {
	Query string
}

//var test = make(chan string)

func GetIP(c *gin.Context) {
	var response response.ResponseCrud
	//test := make(chan string)
	// for i := 0; i < 10; i++ {
	// 	test <- "ping " + strconv.Itoa(i)
	// 	if i == 5 {
	// 		fmt.Println(test)
	// 	}
	// }

	// fmt.Println(test)
	//test := ""

	// go func() { test <- "" }()
	// msg := <-test

	// for i := 0; i < 20; i++ {
	// 	//test = "ping " + strconv.Itoa(i)
	// 	//test <- "ping "
	// 	a := "aaa"
	// 	b := "bbb"

	// 	if msg == "" {
	// 		fmt.Println("test kosong ")
	// 	}

	// 	if i == 0 {

	// 		go func(a, msg string) { test <- "ping " + a + " | " + msg }(a, msg)
	// 	} else {
	// 		go func(b, msg string) { test <- "ping " + b + " | " + msg }(b, msg)
	// 	}
	// 	msg = <-test
	// }
	// //msg := <-test
	// fmt.Println(msg)

	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "atas"
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "bawah"
	}

	var ip IP
	json.Unmarshal(body, &ip)
	fmt.Println(ip.Query)
	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "sukses"
	response.Result = ip

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}

	c.JSON(http.StatusOK, response)
	return
}
