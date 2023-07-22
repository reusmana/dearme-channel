package zaloraController

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/models"
	"github.com/rals/dearme-channel/models/response"
	"github.com/rals/dearme-channel/repositories/tokenRepository"

	"github.com/google/uuid"
	"github.com/rals/dearme-channel/helpers"
	"github.com/rals/dearme-channel/services/tokenService"
	"github.com/rals/dearme-channel/utils"
	// "encoding/json"
)

//var pathz = "/u01/rals-wms-api/packing/printlabel/"

var pathz = "c:/xampp/htdocs/wms-api.ramayana.co.id/api-wms/"

func parseJsonDocZaloraV2(jsonBuffer []byte) models.DocumentZaloraV2 {

	DocumentZaloraV2 := models.DocumentZaloraV2{}

	err := json.Unmarshal(jsonBuffer, &DocumentZaloraV2)
	if err != nil {
		return DocumentZaloraV2
	}

	// the array is now filled with users
	return DocumentZaloraV2

}

func parseJsonDetailStockV2(jsonBuffer []byte) models.DetailStockZaloraV2 {

	DetailStockZaloraV2 := models.DetailStockZaloraV2{}

	err := json.Unmarshal(jsonBuffer, &DetailStockZaloraV2)
	if err != nil {
		return DetailStockZaloraV2
	}

	// the array is now filled with users
	return DetailStockZaloraV2

}

func parseJsonStockZaloraV2(jsonBuffer []byte) models.UpdateStockZaloraV2 {

	UpdateStockZaloraV2 := models.UpdateStockZaloraV2{}

	err := json.Unmarshal(jsonBuffer, &UpdateStockZaloraV2)
	if err != nil {
		return UpdateStockZaloraV2
	}

	// the array is now filled with users
	return UpdateStockZaloraV2

}

func parseJsonPackedV2(jsonBuffer []byte) models.ZaloraPacked {

	ZaloraPacked := models.ZaloraPacked{}

	err := json.Unmarshal(jsonBuffer, &ZaloraPacked)
	if err != nil {
		return ZaloraPacked
	}

	// the array is now filled with users
	return ZaloraPacked

}

func parseJsonErrorV3(jsonBuffer []byte) models.ErrorZaloraV3 {

	ErrorZaloraV3 := models.ErrorZaloraV3{}

	err := json.Unmarshal(jsonBuffer, &ErrorZaloraV3)
	if err != nil {
		return ErrorZaloraV3
	}

	// the array is now filled with users
	return ErrorZaloraV3

}

func parseJsonErrorV2(jsonBuffer []byte) models.ErrorZaloraV2 {

	ErrorZaloraV2 := models.ErrorZaloraV2{}

	err := json.Unmarshal(jsonBuffer, &ErrorZaloraV2)
	if err != nil {
		return ErrorZaloraV2
	}

	// the array is now filled with users
	return ErrorZaloraV2

}

func parseJsonToken(jsonBuffer []byte) models.TokenZalora {

	TokenZalora := models.TokenZalora{}

	err := json.Unmarshal(jsonBuffer, &TokenZalora)
	if err != nil {
		return TokenZalora
	}

	// the array is now filled with users
	return TokenZalora

}

func parseJsonProductsV2Zalora(jsonBuffer []byte) models.ProductV2Zalora {

	ProductV2Zalora := models.ProductV2Zalora{}

	err := json.Unmarshal(jsonBuffer, &ProductV2Zalora)
	if err != nil {
		return ProductV2Zalora
	}

	// the array is now filled with users
	return ProductV2Zalora

}

func parseJsonOrdersZalora(jsonBuffer []byte) models.OrdersZalora {

	OrdersZalora := models.OrdersZalora{}

	err := json.Unmarshal(jsonBuffer, &OrdersZalora)
	if err != nil {
		return OrdersZalora
	}

	// the array is now filled with users
	return OrdersZalora

}

func parseJsonOrderDetailZaloraV2(jsonBuffer []byte) models.OrderDetailZaolraV2 {

	OrderDetailZaolraV2 := models.OrderDetailZaolraV2{}

	err := json.Unmarshal(jsonBuffer, &OrderDetailZaolraV2)
	if err != nil {
		return OrderDetailZaolraV2
	}

	// the array is now filled with users
	return OrderDetailZaolraV2

}

func parseJsonShipmentProvidesResponse(jsonBuffer []byte) models.SuccessResponseShipmentProvidesZalora {

	SuccessResponseShipmentProvidesZalora := models.SuccessResponseShipmentProvidesZalora{}

	err := json.Unmarshal(jsonBuffer, &SuccessResponseShipmentProvidesZalora)
	if err != nil {
		return SuccessResponseShipmentProvidesZalora
	}

	// the array is now filled with users
	return SuccessResponseShipmentProvidesZalora

}

func parseJsonSuccessResponse(jsonBuffer []byte) models.SuccessResponseHeader {

	SuccessResponseHeader := models.SuccessResponseHeader{}

	err := json.Unmarshal(jsonBuffer, &SuccessResponseHeader)
	if err != nil {
		return SuccessResponseHeader
	}

	// the array is now filled with users
	return SuccessResponseHeader

}

func parseJsonOrderSingle(jsonBuffer []byte) models.SuccessResponseHeaderOrderSingle {

	SuccessResponseHeaderOrderSingle := models.SuccessResponseHeaderOrderSingle{}

	err := json.Unmarshal(jsonBuffer, &SuccessResponseHeaderOrderSingle)
	if err != nil {
		return SuccessResponseHeaderOrderSingle
	}

	// the array is now filled with users
	return SuccessResponseHeaderOrderSingle

}

func parseJsonZaloraResponseError(jsonBuffer []byte) models.ErrorResponseHeader {

	ErrorResponseHeader := models.ErrorResponseHeader{}

	err := json.Unmarshal(jsonBuffer, &ErrorResponseHeader)
	if err != nil {
		return ErrorResponseHeader
	}

	// the array is now filled with users
	return ErrorResponseHeader

}
func parseJsonItemZaloraResponse(jsonBuffer []byte) models.SuccessResponseItemZalora {

	SuccessResponseItemZalora := models.SuccessResponseItemZalora{}

	err := json.Unmarshal(jsonBuffer, &SuccessResponseItemZalora)
	if err != nil {
		return SuccessResponseItemZalora
	}

	// the array is now filled with users
	return SuccessResponseItemZalora

}

func parseJsonItemSingleZaloraResponse(jsonBuffer []byte) models.SuccessResponseItemZaloras {

	SuccessResponseItemZaloras := models.SuccessResponseItemZaloras{}

	err := json.Unmarshal(jsonBuffer, &SuccessResponseItemZaloras)
	if err != nil {
		return SuccessResponseItemZaloras
	}

	// the array is now filled with users
	return SuccessResponseItemZaloras

}

func parseJsonDetailOrderZalora(jsonBuffer []byte) models.DetailOrderZalora {

	DetailOrderZalora := models.DetailOrderZalora{}

	err := json.Unmarshal(jsonBuffer, &DetailOrderZalora)
	if err != nil {
		return DetailOrderZalora
	}

	// the array is now filled with users
	return DetailOrderZalora

}

func parseJsonDocumentShipZalora(jsonBuffer []byte) models.DocumentShipZalora {

	DocumentShipZalora := models.DocumentShipZalora{}

	err := json.Unmarshal(jsonBuffer, &DocumentShipZalora)
	if err != nil {
		return DocumentShipZalora
	}

	// the array is now filled with users
	return DocumentShipZalora

}

func parseJsonProductZalora(jsonBuffer []byte) models.SuccessResponseProductZalora {

	SuccessResponseProductZalora := models.SuccessResponseProductZalora{}

	err := json.Unmarshal(jsonBuffer, &SuccessResponseProductZalora)
	if err != nil {
		return SuccessResponseProductZalora
	}

	// the array is now filled with users
	return SuccessResponseProductZalora

}

func parseJsonProductStockZalora(jsonBuffer []byte) models.SuccessResponseProductStockZalora {

	SuccessResponseProductStockZalora := models.SuccessResponseProductStockZalora{}

	err := json.Unmarshal(jsonBuffer, &SuccessResponseProductStockZalora)
	if err != nil {
		return SuccessResponseProductStockZalora
	}

	// the array is now filled with users
	return SuccessResponseProductStockZalora

}

func parseJsonPickupZalora(jsonBuffer []byte) models.SuccessResponsePickZalora {

	SuccessResponsePickZalora := models.SuccessResponsePickZalora{}

	err := json.Unmarshal(jsonBuffer, &SuccessResponsePickZalora)
	if err != nil {
		return SuccessResponsePickZalora
	}

	// the array is now filled with users
	return SuccessResponsePickZalora

}

func parseJsonBatchStockZalora(jsonBuffer []byte) models.UpdateStockBatchZalora {

	UpdateStockBatchZalora := models.UpdateStockBatchZalora{}

	err := json.Unmarshal(jsonBuffer, &UpdateStockBatchZalora)
	if err != nil {
		return UpdateStockBatchZalora
	}

	// the array is now filled with users
	return UpdateStockBatchZalora

}

func TestApi(c *gin.Context) {
	var response response.ResponseCrud

	var wg sync.WaitGroup // New wait group
	wg.Add(1)             // Using two goroutines

	// printCoba(20, "test go", &wg)
	// go printCoba(50, "test", &wg)

	// go func() {
	// 	printCoba(10, "test")
	// 	wg.Done()
	// }()
	go func() {
		printCoba(5, "test go")
		printCoba(7, "test goss")
		wg.Done()
		//wg.Done()
	}()

	printCoba(10, "test")

	wg.Wait()
	// var input string
	// fmt.Scanln(&input)

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = ""
	c.JSON(http.StatusOK, response)
	return
}

func printCoba(till int, message string) {
	//func printCoba(till int, message string) {
	for i := 0; i < till; i++ {
		if message == "test go" {
			//fmt.Println("sleep " + message)
			time.Sleep(7 * time.Second)
		}
		if message == "test" {
			//fmt.Println("sleep " + message)
			time.Sleep(5 * time.Second)
		}
		if message == "test goss" {
			//fmt.Println("sleep " + message)
			time.Sleep(9 * time.Second)
		}
		fmt.Println((i + 1), message)
	}

}

// func AuthZalora(action string, key1 string, value1 string,key2 string, value2 string,key3 string, value3 string) (string, string) {
func AuthZalora(action string, key1 string) (string, string) {
	urlzalora := os.Getenv("URL_API_ZALORA")

	format := "JSON"
	//timestamp := time.Now().Format(time.RFC3339)
	timestamp := time.Now().Format("2006-01-02T15:04:05-0700")

	userid := os.Getenv("PARTNER_MAIL_ZALORA")

	version := "1.0"

	parameter := map[string]string{
		"Action":    action,
		"Format":    format,
		"Timestamp": timestamp,
		"UserID":    userid,
		"Version":   version,
	}

	paramTambah := strings.Split(key1, "#")

	index := 2
	keyisi := "kosong"
	for _, value := range paramTambah {

		if (index % 2) == 0 {
			keyisi = value
		} else {

			//value = strings.ReplaceAll(value, " ", "%20")
			parameter[keyisi] = value
		}
		index++

	}

	// if key1 != "" {
	// 	parameter[key1] = value1
	// }

	parameter_sort := make([]string, 0, len(parameter))
	for k := range parameter {
		parameter_sort = append(parameter_sort, k)
	}
	sort.Strings(parameter_sort)
	concatenated := ""
	i := 0

	for _, k := range parameter_sort {
		urls := url.QueryEscape(parameter[k])
		if i == 0 {
			concatenated += k + "=" + urls
		} else {
			//concatenated += "&" + k + "=" + urls
			urls = strings.ReplaceAll(urls, "+", "%20")
			concatenated += "&" + k + "=" + urls
		}
		i++
	}

	api_key := os.Getenv("PARTNER_KEY_ZALORA")

	h := hmac.New(sha256.New, []byte(api_key))
	h.Write([]byte(concatenated))
	sign := hex.EncodeToString(h.Sum(nil))
	concatenated += "&Signature=" + sign

	return urlzalora, concatenated

}

func GetOrders(c *gin.Context) {
	var response response.ResponseCrud
	statusZalora := c.Param("status")

	Tarik, _ := strconv.Atoi(os.Getenv("GET_ORDERS_TIME_ZALORA"))
	createdAfter := time.Now().Add(time.Duration(-Tarik) * time.Hour).Format("2006-01-02T15:04:05-0700")

	// urlzalora, concatenated := AuthZalora("GetOrders", "CreatedAfter#"+createdAfter+"#Status#pending")
	// urlzalora, concatenated := AuthZalora("GetOrders", "CreatedAfter#"+createdAfter+"")
	parameter := ""
	parameter += "CreatedAfter#" + createdAfter
	parameter += "#"
	parameter += "Limit#100"
	parameter += "#"
	parameter += "Status#" + statusZalora
	fmt.Println("======= Zalora " + statusZalora + " ======= ")
	//pending
	//ready_to_ship
	//canceled

	urlzalora, concatenated := AuthZalora("GetOrders", parameter)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetOrders Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonSuccessResponse(data)
		//disini
		// response.Message = ""
		response.Message = url

		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "GetOrders")
			fmt.Println("Error GetOrders Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		}

		//fmt.Println("===========")
		//fmt.Println(datas)
		//fmt.Println(len(datas.SuccessResponse.Body.Orders))
		//fmt.Println("===========")
		response.Result = datas
		//response.Result = string(data)
		index := 0
		if len(datas.SuccessResponse.Body.Orders) > 0 {

			for _, val := range datas.SuccessResponse.Body.Orders {

				//cek jika sudah ada lewatin kalo bukan canceled
				if statusZalora != "canceled" {
					noordernya := val.Order.OrderNumber + "-" + val.Order.OrderId
					objCek, _ := tokenRepository.FindSalesOrder(noordernya)
					if objCek.NoOrder == "" {
						GetDetailItem(c, val.Order.OrderId, val)
					} else if statusZalora == "delivered" {
						GetDetailItem(c, val.Order.OrderId, val)
					}
				} else {
					GetDetailItem(c, val.Order.OrderId, val)
				}

				index++
			}
			Countnya, _ := strconv.Atoi(datas.SuccessResponse.Head.TotalCount)
			if Countnya > index {
				//masuk order loop
				fmt.Println("masuk loop " + strconv.Itoa(index))
				GetOrdersLoop(c, index, parameter, statusZalora)
			}

		} else {

			//jika order bukan object
			var dataCustObj models.OrdersListZalora
			//var dataCustObj models.BodyList

			datasx := parseJsonOrderSingle(data)
			if datasx.SuccessResponse.Body.Orders.Order.OrderId != "" {
				OrderIdnya := datasx.SuccessResponse.Body.Orders.Order.OrderId
				dataCustObj.Order = datasx.SuccessResponse.Body.Orders.Order
				// dataCustObj.Orders[0].Order = datasx.SuccessResponse.Body.Orders.Order
				//cek jika sudah ada lewatin kalo bukan canceled
				if statusZalora != "canceled" {
					noordernya := datasx.SuccessResponse.Body.Orders.Order.OrderNumber + "-" + OrderIdnya
					objCek, _ := tokenRepository.FindSalesOrder(noordernya)
					if objCek.NoOrder == "" {
						GetDetailItem(c, OrderIdnya, dataCustObj)
					} else if statusZalora == "delivered" {
						GetDetailItem(c, OrderIdnya, dataCustObj)
					}
				} else {
					GetDetailItem(c, OrderIdnya, dataCustObj)
				}
			}

		}

		if statusZalora == "pending" {
			GetProcessedOrder("canceled")
			//GetProcessedOrder("ready_to_ship") //gak perlu cek ke sini
		}

		// if statusZalora == "ready_to_ship" {
		// 	GetProcessedOrder("canceled")

		// }

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrdersAuto(c *gin.Context) {
	var response response.ResponseCrud
	statusZalora := c.Param("status")

	Tarik, _ := strconv.Atoi(os.Getenv("GET_ORDERS_TIME_ZALORA"))
	createdAfter := time.Now().Add(time.Duration(-Tarik) * time.Hour).Format("2006-01-02T15:04:05-0700")

	// urlzalora, concatenated := AuthZalora("GetOrders", "CreatedAfter#"+createdAfter+"#Status#pending")
	// urlzalora, concatenated := AuthZalora("GetOrders", "CreatedAfter#"+createdAfter+"")
	parameter := ""
	parameter += "CreatedAfter#" + createdAfter
	parameter += "#"
	parameter += "Limit#100"
	parameter += "#"
	parameter += "Status#" + statusZalora
	fmt.Println("======= Zalora " + statusZalora + " ======= ")
	//pending
	//ready_to_ship
	//canceled

	urlzalora, concatenated := AuthZalora("GetOrders", parameter)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetOrders Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonSuccessResponse(data)
		//disini
		// response.Message = ""
		response.Message = url

		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "GetOrders")
			fmt.Println("Error GetOrders Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		}

		//fmt.Println("===========")
		//fmt.Println(datas)
		//fmt.Println(len(datas.SuccessResponse.Body.Orders))
		//fmt.Println("===========")
		response.Result = datas
		//response.Result = string(data)
		index := 0
		if len(datas.SuccessResponse.Body.Orders) > 0 {

			for _, val := range datas.SuccessResponse.Body.Orders {

				//cek jika sudah ada lewatin kalo bukan canceled
				if statusZalora != "canceled" {
					noordernya := val.Order.OrderNumber + "-" + val.Order.OrderId
					objCek, _ := tokenRepository.FindSalesOrder(noordernya)
					if objCek.NoOrder == "" {
						GetDetailItem(c, val.Order.OrderId, val)
					} else if statusZalora == "delivered" {
						GetDetailItem(c, val.Order.OrderId, val)
					}
				} else {
					GetDetailItem(c, val.Order.OrderId, val)
				}

				index++
			}
			Countnya, _ := strconv.Atoi(datas.SuccessResponse.Head.TotalCount)
			if Countnya > index {
				//masuk order loop
				fmt.Println("masuk loop " + strconv.Itoa(index))
				GetOrdersLoop(c, index, parameter, statusZalora)
			}

		} else {

			//jika order bukan object
			var dataCustObj models.OrdersListZalora
			//var dataCustObj models.BodyList

			datasx := parseJsonOrderSingle(data)
			if datasx.SuccessResponse.Body.Orders.Order.OrderId != "" {
				OrderIdnya := datasx.SuccessResponse.Body.Orders.Order.OrderId
				dataCustObj.Order = datasx.SuccessResponse.Body.Orders.Order
				// dataCustObj.Orders[0].Order = datasx.SuccessResponse.Body.Orders.Order
				//cek jika sudah ada lewatin kalo bukan canceled
				if statusZalora != "canceled" {
					noordernya := datasx.SuccessResponse.Body.Orders.Order.OrderNumber + "-" + OrderIdnya
					objCek, _ := tokenRepository.FindSalesOrder(noordernya)
					if objCek.NoOrder == "" {
						GetDetailItem(c, OrderIdnya, dataCustObj)
					} else if statusZalora == "delivered" {
						GetDetailItem(c, OrderIdnya, dataCustObj)
					}
				} else {
					GetDetailItem(c, OrderIdnya, dataCustObj)
				}
			}

		}

		if statusZalora == "pending" {
			//GetProcessedOrderAuto("ready_to_ship")
			GetProcessedOrderAuto("canceled")
		}

		// if statusZalora == "ready_to_ship" {
		// 	GetProcessedOrderAuto("canceled")

		// }

	}

	//c.JSON(http.StatusOK, response)
	//return
}

func GetProcessedOrderAuto(status string) {
	urlApi := os.Getenv("URL_WMS_ZALORA") + "GetOrdersAuto/" + status

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
		fmt.Println("gagal GetProcessedOrder Zalora")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

	} else {
		fmt.Println("sukses GetProcessedOrder Zalora")
		defer resp.Body.Close()
	}

}

func GetProcessedOrder(status string) {
	urlApi := os.Getenv("URL_ORDER_ZALORA") + status

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
		fmt.Println("gagal GetProcessedOrder Zalora")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

	} else {
		fmt.Println("sukses GetProcessedOrder Zalora")
		defer resp.Body.Close()
	}

}

func GetOrdersLoop(c *gin.Context, index int, parameter string, statusZalora string) {
	nmRoute := "GetOrdersLoop"
	var response response.ResponseCrud
	parameter = parameter + "#Offset#" + strconv.Itoa(index)
	urlzalora, concatenated := AuthZalora("GetOrders", parameter)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal " + nmRoute + " Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonSuccessResponse(data)
		response.Message = ""

		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, nmRoute)
			fmt.Println("Error " + nmRoute + " Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		}

		//fmt.Println("===========")
		//fmt.Println(datas)
		//fmt.Println(len(datas.SuccessResponse.Body.Orders))
		//fmt.Println("===========")
		response.Result = datas
		//response.Result = string(data)

		if len(datas.SuccessResponse.Body.Orders) > 0 {

			for _, val := range datas.SuccessResponse.Body.Orders {

				//cek jika sudah ada lewatin kalo bukan canceled
				if statusZalora != "canceled" {
					noordernya := val.Order.OrderNumber + "-" + val.Order.OrderId
					objCek, _ := tokenRepository.FindSalesOrder(noordernya)
					if objCek.NoOrder == "" {
						GetDetailItem(c, val.Order.OrderId, val)
					} else if statusZalora == "delivered" {
						GetDetailItem(c, val.Order.OrderId, val)
					}
				} else {
					GetDetailItem(c, val.Order.OrderId, val)
				}

				//GetDetailItem(c, val.Order.OrderId, val)
				index++
			}
			Countnya, _ := strconv.Atoi(datas.SuccessResponse.Head.TotalCount)
			if Countnya > index {
				//masuk order loop
				fmt.Println("lanjut loop " + strconv.Itoa(index))
				GetOrdersLoop(c, index, parameter, statusZalora)
			}

		} else {

			//jika order bukan object

			var dataCustObj models.OrdersListZalora
			datasx := parseJsonOrderSingle(data)
			if datasx.SuccessResponse.Body.Orders.Order.OrderId != "" {
				OrderIdnya := datasx.SuccessResponse.Body.Orders.Order.OrderId
				dataCustObj.Order = datasx.SuccessResponse.Body.Orders.Order
				//cek jika sudah ada lewatin kalo bukan canceled
				if statusZalora != "canceled" {
					noordernya := datasx.SuccessResponse.Body.Orders.Order.OrderNumber + "-" + OrderIdnya
					objCek, _ := tokenRepository.FindSalesOrder(noordernya)
					if objCek.NoOrder == "" {
						GetDetailItem(c, OrderIdnya, dataCustObj)
					} else if statusZalora == "delivered" {
						GetDetailItem(c, OrderIdnya, dataCustObj)
					}
				} else {

					GetDetailItem(c, OrderIdnya, dataCustObj)
				}
			}

		}

	}

	//c.JSON(http.StatusOK, response)
	//return
}

func GetOrderDetail(c *gin.Context) {
	var response response.ResponseCrud

	NoOrder := c.Param("noorder")
	NoOrderObj := strings.Split(NoOrder, "-")
	if len(NoOrderObj) > 0 {
		NoOrder = NoOrderObj[1]
	}

	urlzalora, concatenated := AuthZalora("GetOrder", "OrderId#"+NoOrder)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetOrderDetail Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonDetailOrderZalora(data)
		response.Message = ""
		response.Result = datas
		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "GetOrderDetail")
			fmt.Println("Error GetOrderDetail Zaloras")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
			response.ResponseDesc = datasx.ErrorResponse.Head.ErrorMessage
		} else {
			response.ResponseDesc = datas.SuccessResponse.Body.Orders.Order.Statuses[0].Status

		}

		// fmt.Println("===========")
		// fmt.Println(string(data))
		// // fmt.Println(len(datas.SuccessResponse.Body.Orders))
		// fmt.Println("===========")

		//response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetDetailItem(c *gin.Context, noorder string, dataCust models.OrdersListZalora) {
	var response response.ResponseCrud

	urlzalora, concatenated := AuthZalora("GetOrderItems", "OrderId#"+noorder)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated
	fmt.Println(noorder)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetDetailItem Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonItemZaloraResponse(data)
		// fmt.Println(string(data))
		response.Message = url

		if len(datas.SuccessResponse.Body.OrderItems.OrderItem) > 0 {
			tokenService.SaveSalesOrderZalora("zalora", datas, dataCust)
			fmt.Println("datas object")
		} else {
			datasx := parseJsonItemSingleZaloraResponse(data)
			fmt.Println("datasx single")
			tokenService.SaveSalesOrderZaloraSingle("zalora", datasx, dataCust)
		}

		// tokenService.SaveSalesOrderZalora("zalora", datas, dataCust)

		// fmt.Println("******************")
		// // fmt.Println(string(data))
		// fmt.Println(datas)
		// fmt.Println("******************")
		// response.Result = string(data)
		response.Result = datas

	}
	// c.JSON(http.StatusOK, response)
	// return
}

func GetOrderItemsDetail(c *gin.Context) {
	var response response.ResponseCrud

	noorder := c.Param("noorder")
	noorderx := noorder
	NoOrderObj := strings.Split(noorderx, "-")
	if len(NoOrderObj) > 0 {
		noorderx = NoOrderObj[1]
	}

	urlzalora, concatenated := AuthZalora("GetOrderItems", "OrderId#"+noorderx)

	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetOrderItemsDetail Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonItemZaloraResponse(data)
		// fmt.Println(string(data))
		response.Message = url

		if len(datas.SuccessResponse.Body.OrderItems.OrderItem) > 0 {
			//tokenService.SaveSalesOrderZalora("zalora", datas, dataCust)
			response.Result = datas
		} else {
			datasx := parseJsonItemSingleZaloraResponse(data)
			response.Result = datasx
			//tokenService.SaveSalesOrderZaloraSingle("zalora", datasx, dataCust)
		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetDetailItemView(noorder string) models.OrdersItemsZalora {
	var response response.ResponseCrud
	var obj models.OrdersItemsZalora
	nmRoute := "GetDetailItemView"
	urlzalora, concatenated := AuthZalora("GetOrderItems", "OrderId#"+noorder)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetDetailItemView Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonItemZaloraResponse(data)
		response.Message = ""
		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, nmRoute)
			fmt.Println("Error " + nmRoute + " Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		} else {
			obj = datas.SuccessResponse.Body.OrderItems

		}

		// fmt.Println("******************")
		// // fmt.Println(string(data))
		// fmt.Println(datas)
		// fmt.Println("******************")
		// response.Result = string(data)

	}
	// c.JSON(http.StatusOK, response)
	return obj
}

func GetDetailItemViewSingle(noorder string) models.OrdersItemsZaloras {
	var response response.ResponseCrud
	var obj models.OrdersItemsZaloras
	nmRoute := "GetDetailItemViewSingle"
	urlzalora, concatenated := AuthZalora("GetOrderItems", "OrderId#"+noorder)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetDetailItemViewSingle Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonItemSingleZaloraResponse(data)
		response.Message = ""
		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, nmRoute)
			fmt.Println("Error " + nmRoute + " Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		} else {
			obj = datas.SuccessResponse.Body.OrderItems

		}

		// fmt.Println("******************")
		// // fmt.Println(string(data))
		// fmt.Println(datas)
		// fmt.Println("******************")
		// response.Result = string(data)

	}
	// c.JSON(http.StatusOK, response)
	return obj
}

func GetProducts(c *gin.Context) {
	var response response.ResponseCrud

	Tarik, _ := strconv.Atoi(os.Getenv("GET_ORDERS_TIME_ZALORA"))
	createdAfter := time.Now().Add(time.Duration(-Tarik) * time.Hour).Format("2006-01-02T15:04:05-0700")

	parameter := ""
	parameter += "UpdatedAfter#" + createdAfter
	parameter += "#"
	parameter += "CreatedAfter#" + createdAfter
	parameter += "#"
	parameter += "Limit#100"

	// urlzalora, concatenated := AuthZalora("GetProducts", "Limit#100")
	urlzalora, concatenated := AuthZalora("GetProducts", parameter)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetProducts Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonProductZalora(data)
		response.Message = ""
		response.Result = datas
		// if datas.SuccessResponse.Head.RequestAction == "" {
		// 	datasx := parseJsonZaloraResponseError(data)
		// 	response.Result = datasx
		// 	tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "GetProducts")
		// 	fmt.Println("Error GetProducts Zalora")
		// 	response.Message = datasx.ErrorResponse.Head.ErrorMessage
		// } else {
		// 	if len(datas.SuccessResponse.Body.Products.Product) > 0 {
		// 		tokenService.SaveSkuMappingZalora(datas.SuccessResponse.Body.Products)
		// 	}
		// 	GetProductsLoop(parameter, 100)
		// }
		fmt.Println("finish product zalora")
		// fmt.Println("===========")
		// fmt.Println(string(data))
		// // fmt.Println(len(datas.SuccessResponse.Body.Orders))
		// fmt.Println("===========")

		//response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}
func GetProductsAuto(wg *sync.WaitGroup) {
	var response response.ResponseCrud
	fmt.Println("mulai product zalora " + time.Now().String())
	Tarik, _ := strconv.Atoi(os.Getenv("GET_ORDERS_TIME_ZALORA"))
	createdAfter := time.Now().Add(time.Duration(-Tarik) * time.Hour).Format("2006-01-02T15:04:05-0700")

	parameter := ""
	parameter += "UpdatedAfter#" + createdAfter
	parameter += "#"
	parameter += "CreatedAfter#" + createdAfter
	parameter += "#"
	parameter += "Limit#100"

	// urlzalora, concatenated := AuthZalora("GetProducts", "Limit#100")
	urlzalora, concatenated := AuthZalora("GetProducts", parameter)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetProducts Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonProductZalora(data)
		response.Message = ""
		response.Result = datas
		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "GetProducts")
			//fmt.Println("Error GetProducts Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		} else {
			if len(datas.SuccessResponse.Body.Products.Product) > 0 {
				tokenService.SaveSkuMappingZalora(datas.SuccessResponse.Body.Products)
			}
			GetProductsLoop(parameter, 100)
		}
		fmt.Println("finish product zalora " + time.Now().String())

	}

	wg.Done()
}

func GetProductsLoop(parameter string, offset int) bool {
	var response response.ResponseCrud

	parameters := parameter + "#Offset#" + strconv.Itoa(offset)
	// urlzalora, concatenated := AuthZalora("GetProducts", "Limit#100#Offset#"+strconv.Itoa(offset))
	urlzalora, concatenated := AuthZalora("GetProducts", parameters)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetProductsLoop Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonProductZalora(data)
		response.Message = ""
		response.Result = datas
		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "GetProductsLoop")
			fmt.Println("Error GetProductsLoop Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		}

		if len(datas.SuccessResponse.Body.Products.Product) > 0 && offset < 3000 {
			tokenService.SaveSkuMappingZalora(datas.SuccessResponse.Body.Products)
			offset += 100
			//fmt.Println("offset: " + strconv.Itoa(offset))
			GetProductsLoop(parameter, offset)

		}

	}

	return true

}

func GetProductStock(c *gin.Context) {
	var response response.ResponseCrud

	sku := c.Param("sku")
	urlzalora, concatenated := AuthZalora("GetProductStocks", "Search#"+sku)

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetProductStock Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonProductStockZalora(data)
		response.Message = url
		response.Result = datas
		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "GetProductStock")
			fmt.Println("Error GetProductStock Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		}
		// fmt.Println("===========")
		// fmt.Println(string(data))
		// // fmt.Println(len(datas.SuccessResponse.Body.Orders))
		// fmt.Println("===========")

		//response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStockProduct(c *gin.Context) {
	var response response.ResponseCrud

	sku := c.Param("sku")
	stock := c.Param("stock")

	//cari sku_mapping

	//objSku, _ := tokenRepository.CariSkuMapping(sku, "R991")
	objSku, _ := tokenRepository.CariSkuMappingActive(sku, "R991")

	var objRest models.TableLogStock

	if objSku.SkuNo != "" {

		//fmt.Println("lewat sini  ")
		urlzalora, concatenated := AuthZalora("ProductStockUpdate", "")
		//fmt.Println("masuk siniss")
		//concatenated += "&CreatedAfter=" + CreatedAfter
		url := urlzalora + "?" + concatenated

		xmlString := ""

		skustock := objSku.IdSkuParent
		skustock = sku

		xmlString = `<?xml version="1.0" encoding="UTF-8" ?>
		<Request>
		<Product>
		<SellerSku>` + skustock + `</SellerSku>
		<Quantity>` + stock + `</Quantity>
		</Product>
		</Request>`

		objRest.UuidLog = uuid.New().String()
		objRest.ChannelCode = os.Getenv("KODE_ZALORA")
		objRest.Sku = skustock
		objRest.Body = xmlString

		if s, err := strconv.ParseFloat(stock, 64); err == nil {
			objRest.Stock = s
		}

		//fmt.Println(xmlString)
		var body_url = []byte(xmlString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))

		//req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
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
			fmt.Println("gagal GetProductStock Zalora")
			response.ResponseCode = http.StatusInternalServerError
			response.ResponseDesc = enums.ERROR
			response.ResponseTime = utils.DateToStdNow()
			response.Message = "gagal GetProductStock Zalora"
			response.Result = concatenated
			objRest.Response = "Koneksi UpdateStockProduct Gagal"

		} else {
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			datas := parseJsonProductStockZalora(data)
			response.Message = ""
			response.Result = string(data)
			objRest.Response = string(data)
			if datas.SuccessResponse.Head.RequestAction == "" {
				datasx := parseJsonZaloraResponseError(data)
				response.Result = datasx
				tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage+"|"+sku+"|"+stock, "ProductStockUpdate")
				fmt.Println("Error ProductStockUpdate Zalora")
				response.Message = datasx.ErrorResponse.Head.ErrorMessage
			}
			// fmt.Println("===========")
			// fmt.Println(string(data))
			// // fmt.Println(len(datas.SuccessResponse.Body.Orders))
			// fmt.Println("===========")

			//response.Result = string(data)

		}

		objRest.CreatedBy = "API"
		objRest.CreatedDate = time.Now()
		tokenRepository.SaveStockAPI(objRest)

	} else { //objSku.SkuNo != ""
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "SKU TIDAK DITEMUKAN"
	}

	c.JSON(http.StatusOK, response)
	return
}
func UpdateStockZaloraBatch(param string) {
	var response response.ResponseCrud

	//fmt.Println("lewat sini  ")
	urlzalora, concatenated := AuthZalora("ProductStockUpdate", "")
	//fmt.Println("masuk siniss")
	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	xmlString := ""

	xmlString = `<?xml version="1.0" encoding="UTF-8" ?>
		<Request>
		` + param + `
		</Request>`

	//fmt.Println(xmlString)
	var body_url = []byte(xmlString)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))

	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal UpdateStockZaloraBatch Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "gagal UpdateStockZaloraBatch Zalora"
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonProductStockZalora(data)
		response.Message = ""
		response.Result = datas
		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "UpdateStockZaloraBatch")
			fmt.Println("Error UpdateStockZaloraBatch Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		}

	}

}
func BatchUpdateStockZalora(c *gin.Context) {
	var response response.ResponseCrud

	//parseJsonBatchStockZalora
	type SkuList struct {
		Sku   string `json:"sku"`
		Stock int16  `json:"stock"`
	}

	type FormValue struct {
		SkuList []SkuList `json:"sku_list"`
	}

	var request FormValue = FormValue{}
	if err := c.Bind(&request); err != nil {
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "Format JSON Tidak Valid"
		c.JSON(http.StatusOK, response)
		return
	}

	if len(request.SkuList) < 1 {
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "Format JSON Tidak Valid"
		c.JSON(http.StatusOK, response)
		return
	}

	BodyProduct := ""
	index := 1
	for _, val := range request.SkuList {
		objSku, _ := tokenRepository.CariSkuMapping(val.Sku, "R991")
		if objSku.SkuNo != "" {
			BodyProduct = BodyProduct + `
			<Product>
			<SellerSku>` + val.Sku + `</SellerSku>
			<Quantity>` + strconv.Itoa(int(val.Stock)) + `</Quantity>
			</Product>
			`
			index++
		}

		if index == 50 {
			UpdateStockZaloraBatch(BodyProduct)
			BodyProduct = ""
		}

	}

	if BodyProduct != "" {
		UpdateStockZaloraBatch(BodyProduct)
		BodyProduct = ""
	}

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "CEK CEK"
	response.Result = request

	c.JSON(http.StatusOK, response)
	return
}

func UpdateOtomatisStockZalora(c *gin.Context) {

	ObjMapping := tokenService.CariSkuMappingObjGroup(os.Getenv("KODE_ZALORA")) //param by channel name

	if len(ObjMapping) > 0 {
		for _, value := range ObjMapping {

			helpers.UpdateStock(value.SkuNo, "API_CHANNEL", os.Getenv("KODE_ZALORA"))

			fmt.Println(len(ObjMapping))
			fmt.Println("SKU " + value.SkuNo)
		}
	}

}

func UpdateOtomatisStockZaloraOLD(c *gin.Context) {
	var response response.ResponseCrud
	fmt.Println(os.Getenv("KODE_ZALORA"))
	ObjMapping := tokenService.CariSkuMappingObj(os.Getenv("KODE_ZALORA")) //param by channel name
	//spew.Dump(ObjMapping)
	if len(ObjMapping) > 0 {
		for _, value := range ObjMapping {
			//cari status uid AVLB
			objAlb := tokenService.CariUidAvlb(value.SkuNo)
			objBuffer := tokenRepository.CariBufferStock(value.SkuNo)

			IdParent := value.IdSkuParent

			StockBuffers := 0
			if objBuffer.SkuNo != "" {
				StockBuffers = int(objBuffer.BufferStock)
			}

			StockBuffer := (len(objAlb)) - StockBuffers
			if StockBuffer < 1 {
				StockBuffer = 0
			}

			UpdateStockProductLoop(IdParent, strconv.Itoa(int(StockBuffer)))

		}
	}
	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "FINISH"

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStockProductLoop(sku string, stock string) {
	var response response.ResponseCrud

	//cari sku_mapping

	objSku, _ := tokenRepository.CariSkuMappingActive(sku, "R991")
	if objSku.SkuNo != "" {

		urlzalora, concatenated := AuthZalora("ProductStockUpdate", "")

		url := urlzalora + "?" + concatenated

		xmlString := ""

		xmlString = `<?xml version="1.0" encoding="UTF-8" ?>
		<Request>
		<Product>
		<SellerSku>` + sku + `</SellerSku>
		<Quantity>` + stock + `</Quantity>
		</Product>
		</Request>`

		//fmt.Println(xmlString)
		var body_url = []byte(xmlString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))

		//req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
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
			fmt.Println("gagal UpdateStockProductLoop Zalora")
			response.ResponseCode = http.StatusInternalServerError
			response.ResponseDesc = enums.ERROR
			response.ResponseTime = utils.DateToStdNow()
			response.Message = url
			response.Result = concatenated

		} else {
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			datas := parseJsonProductStockZalora(data)
			response.Message = ""
			response.Result = datas
			if datas.SuccessResponse.Head.RequestAction == "" {
				datasx := parseJsonZaloraResponseError(data)
				response.Result = datasx
				tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage+"|"+sku+"|"+stock, "UpdateStockProductLoop")
				fmt.Println("Error UpdateStockProductLoop Zalora")
				response.Message = datasx.ErrorResponse.Head.ErrorMessage
			}

		}
	} else { //objSku.SkuNo != ""
		fmt.Println("SKU TIDAK DITEMUKAN")
	}
}

func CekDetailItem(c *gin.Context) {
	var response response.ResponseCrud

	urlzalora, concatenated := AuthZalora("SetStatusToReadyToShip", "Limit#100")

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal CekDetailItem Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonProductZalora(data)
		response.Message = ""
		response.Result = datas
		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "CekDetailItem")
			fmt.Println("Error CekDetailItem Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		} else {
		}
	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStatusZalora(c *gin.Context) {
	var response response.ResponseCrud
	nmRoute := "UpdateStatusZalora"

	statusZalora := c.Param("status")
	fmt.Println(statusZalora)
	noorder := c.Param("noorder")
	noorderx := noorder
	NoOrderObj := strings.Split(noorderx, "-")
	if len(NoOrderObj) > 0 {
		noorderx = NoOrderObj[1]
	}

	if noorder != "242967365-8313441" {

		// response.ResponseCode = http.StatusInternalServerError
		// response.ResponseDesc = enums.ERROR
		// response.ResponseTime = utils.DateToStdNow()
		// response.Message = "Update Status Zalora Gagal (SYSTEM UPGRADE)"
		// response.Result = noorderx

		// c.JSON(http.StatusOK, response)
		// return

		DetailItem := GetDetailItemView(noorderx)

		fmt.Println(DetailItem)
		paramTambah := ""
		OrderItemIds := ""
		index := 1
		ShipmentProvider := ""
		fmt.Println(len(DetailItem.OrderItem))
		if len(DetailItem.OrderItem) > 0 {
			fmt.Println("atas")
			for _, valuenya := range DetailItem.OrderItem {
				OrderItemIds += valuenya.OrderItemId
				if index < len(DetailItem.OrderItem) {
					OrderItemIds += ","
				}
				index++
			}

			if statusZalora == "Shipped" {
				OrderItemIds = DetailItem.OrderItem[0].OrderItemId
			}

			ShipmentProvider = DetailItem.OrderItem[0].ShipmentProvider

		} else {
			//single
			fmt.Println("bawah")
			DetailItemSingle := GetDetailItemViewSingle(noorderx)
			fmt.Println(DetailItemSingle.OrderItem.ShipmentProvider)
			OrderItemIds = DetailItemSingle.OrderItem.OrderItemId
			ShipmentProvider = DetailItemSingle.OrderItem.ShipmentProvider
		}

		if statusZalora == "Shipped" {

		} else {
			OrderItemIds = "[" + OrderItemIds + "]"
		}

		response.Result = DetailItem

		if statusZalora == "Shipped" {
			paramTambah += "OrderItemId#" + OrderItemIds
		} else {
			paramTambah += "OrderItemIds#" + OrderItemIds
		}
		deliveryType := "pickup"

		paramTambah += "#DeliveryType#" + deliveryType
		paramTambah += "#ShippingProvider#" + ShipmentProvider

		fmt.Println(len(DetailItem.OrderItem))
		fmt.Println(OrderItemIds)
		urlzalora, concatenated := AuthZalora("SetStatusTo"+statusZalora, paramTambah)
		url := urlzalora + "?" + concatenated
		fmt.Println(url)
		response.Message = url
		response.Result = DetailItem
		fmt.Println(nmRoute)
		req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
			fmt.Println("gagal " + nmRoute + " Zalora")
			response.ResponseCode = http.StatusInternalServerError
			response.ResponseDesc = enums.ERROR
			response.ResponseTime = utils.DateToStdNow()
			// response.Message = url
			response.Message = "Update Status Zalora Gagal"
			response.Result = concatenated

		} else {
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			datas := parseJsonPickupZalora(data)
			response.Message = ""
			response.Result = datas
			fmt.Println(string(data))
			if datas.SuccessResponse.Head.RequestAction == "" {
				datasx := parseJsonZaloraResponseError(data)
				if datasx.ErrorResponse.Head.ErrorCode == "24" {
					objDrop, pesanZal := UpdateStatusZaloraParam(statusZalora, noorder, "dropship")

					response.Result = objDrop
					// if len(objDrop.OrderItem)>1 {
					// 	response.Message = pesanZal
					// }
					if pesanZal != "" {
						response.Message = pesanZal + " (ZALORA)"
					}

				} else {
					response.Message = datasx.ErrorResponse.Head.ErrorMessage + " (ZALORA)"
					response.Result = datasx
					tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, nmRoute)
					fmt.Println("Error " + nmRoute + " Zalora")
					if response.Message == "" {
						response.Message = "Update Status Zalora Gagal"
					}

				}
				//response.Message = datasx.ErrorResponse.Head.ErrorMessage
			} else {

			}

		}

	} else {
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "Update Status Zalora Gagal (TESTING)"
	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStatusZaloraParam(statusZalora string, noorder string, deliveryType string) (models.OrderItemsPickZalora, string) {
	var response response.ResponseCrud
	nmRoute := "UpdateStatusZaloraParam"
	var obj models.OrderItemsPickZalora
	pesan := ""

	noorderx := noorder
	NoOrderObj := strings.Split(noorderx, "-")
	if len(NoOrderObj) > 0 {
		noorderx = NoOrderObj[1]
	}

	DetailItem := GetDetailItemView(noorderx)
	paramTambah := ""
	OrderItemIds := ""
	index := 1
	ShipmentProvider := ""
	fmt.Println(len(DetailItem.OrderItem))

	if len(DetailItem.OrderItem) > 0 {
		fmt.Println("atas")
		for _, valuenya := range DetailItem.OrderItem {
			OrderItemIds += valuenya.OrderItemId
			if index < len(DetailItem.OrderItem) {
				OrderItemIds += ","
			}
			index++
		}

		if statusZalora == "Shipped" {
			OrderItemIds = DetailItem.OrderItem[0].OrderItemId
		}

		ShipmentProvider = DetailItem.OrderItem[0].ShipmentProvider

	} else {
		//single
		fmt.Println("bawah")
		DetailItemSingle := GetDetailItemViewSingle(noorderx)
		fmt.Println(DetailItemSingle.OrderItem.ShipmentProvider)
		OrderItemIds = DetailItemSingle.OrderItem.OrderItemId
		ShipmentProvider = DetailItemSingle.OrderItem.ShipmentProvider
	}

	if statusZalora == "Shipped" {

	} else {
		OrderItemIds = "[" + OrderItemIds + "]"
	}

	response.Result = DetailItem
	if statusZalora == "Shipped" {
		paramTambah += "OrderItemId#" + OrderItemIds
	} else {
		paramTambah += "OrderItemIds#" + OrderItemIds
	}

	paramTambah += "#DeliveryType#" + deliveryType
	paramTambah += "#ShippingProvider#" + ShipmentProvider

	fmt.Println(len(DetailItem.OrderItem))
	fmt.Println(OrderItemIds)
	urlzalora, concatenated := AuthZalora("SetStatusTo"+statusZalora, paramTambah)
	url := urlzalora + "?" + concatenated
	fmt.Println(url)
	response.Message = url
	response.Result = DetailItem
	fmt.Println(nmRoute)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal " + nmRoute + " Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonPickupZalora(data)
		response.Message = ""
		response.Result = datas
		fmt.Println(string(data))
		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, nmRoute)
			fmt.Println("Error " + nmRoute + " Zalora")

			response.Message = datasx.ErrorResponse.Head.ErrorMessage + " (ZALORA)"

			if response.Message == "" {
				response.Message = "Update Status Zalora Gagal"
			}

			pesan = datasx.ErrorResponse.Head.ErrorMessage
		} else {
			obj = datas.SuccessResponse.Body.OrderItems
		}
	}

	return obj, pesan
}

func GetShipmentProviders(c *gin.Context) {
	var response response.ResponseCrud

	urlzalora, concatenated := AuthZalora("GetShipmentProviders", "")

	//concatenated += "&CreatedAfter=" + CreatedAfter
	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetShipmentProviders Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonShipmentProvidesResponse(data)
		response.Message = ""
		response.Result = datas
		// if datas.SuccessResponse.Head.RequestAction == "" {
		// 	datasx := parseJsonZaloraResponseError(data)
		// 	response.Result = datasx
		// 	tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "GetProducts")
		// 	fmt.Println("Error GetProducts Zalora")
		// 	response.Message = datasx.ErrorResponse.Head.ErrorMessage
		// } else {
		// 	if len(datas.SuccessResponse.Body.Products.Product) > 0 {
		// 		tokenService.SaveSkuMappingZalora(datas.SuccessResponse.Body.Products)
		// 	}
		// 	GetProductsLoop(100)
		// }

		//fmt.Println("finish product zalora")
		// fmt.Println("===========")
		// fmt.Println(string(data))
		// // fmt.Println(len(datas.SuccessResponse.Body.Orders))
		// fmt.Println("===========")

		//response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetResiZalora(c *gin.Context) {
	var response response.ResponseCrud

	noorder := c.Param("noorder")
	noorderx := noorder
	NoOrderObj := strings.Split(noorderx, "-")
	if len(NoOrderObj) > 0 {
		noorderx = NoOrderObj[1]
	}

	urlzalora, concatenated := AuthZalora("GetOrderItems", "OrderId#"+noorderx)

	url := urlzalora + "?" + concatenated

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal GetResiZalora Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = enums.ERROR
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonItemZaloraResponse(data)
		// fmt.Println(string(data))
		response.Message = ""
		resinya := ""

		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, "GetResiZalora")
			fmt.Println("Error GetResiZalora Zalora")
			response.Message = datasx.ErrorResponse.Head.ErrorMessage
		} else {

			if len(datas.SuccessResponse.Body.OrderItems.OrderItem) > 0 {
				//tokenService.SaveSalesOrderZalora("zalora", datas, dataCust)
				response.Result = datas
				resinya = datas.SuccessResponse.Body.OrderItems.OrderItem[0].TrackingCode
			} else {
				datasx := parseJsonItemSingleZaloraResponse(data)
				response.Result = datasx
				resinya = datasx.SuccessResponse.Body.OrderItems.OrderItem.TrackingCode
				//tokenService.SaveSalesOrderZaloraSingle("zalora", datasx, dataCust)
			}

		}

		//balikan resi taro di ResponseDesc
		response.ResponseDesc = resinya

	}
	c.JSON(http.StatusOK, response)
	return
}

func ReqPickupZalora(c *gin.Context) {
	var response response.ResponseCrud
	nmRoute := "ReqPickupZalora"

	noorder := c.Param("noorder")
	noorderx := noorder
	NoOrderObj := strings.Split(noorderx, "-")
	if len(NoOrderObj) > 0 {
		noorderx = NoOrderObj[1]
	}

	DetailItem := GetDetailItemView(noorderx)

	fmt.Println(DetailItem)
	paramTambah := ""
	OrderItemIds := ""
	index := 1

	fmt.Println(len(DetailItem.OrderItem))
	if len(DetailItem.OrderItem) > 0 {
		fmt.Println("atas")
		for _, valuenya := range DetailItem.OrderItem {
			OrderItemIds += valuenya.OrderItemId
			if index < len(DetailItem.OrderItem) {
				OrderItemIds += ","
			}
			index++
		}

	} else {
		//single
		fmt.Println("bawah")
		DetailItemSingle := GetDetailItemViewSingle(noorderx)
		fmt.Println(DetailItemSingle.OrderItem.ShipmentProvider)
		OrderItemIds = DetailItemSingle.OrderItem.OrderItemId
	}

	OrderItemIds = "[" + OrderItemIds + "]"
	response.Result = DetailItem

	paramTambah += "OrderItemIds#" + OrderItemIds

	fmt.Println(len(DetailItem.OrderItem))
	fmt.Println(OrderItemIds)
	urlzalora, concatenated := AuthZalora("CreatePickupRequest", paramTambah)
	url := urlzalora + "?" + concatenated
	fmt.Println(url)
	response.Message = url
	response.Result = DetailItem
	fmt.Println(nmRoute)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal " + nmRoute + " Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		// response.Message = url
		response.Message = nmRoute + " Zalora Gagal"
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonPickupZalora(data)
		response.Message = ""
		response.Result = datas
		fmt.Println(string(data))
		if datas.SuccessResponse.Head.RequestAction == "" {
			datasx := parseJsonZaloraResponseError(data)

			response.Message = datasx.ErrorResponse.Head.ErrorMessage + " (ZALORA)"
			response.Result = datasx
			tokenService.SaveErrorString("zalora", datasx.ErrorResponse.Head.ErrorMessage, nmRoute)
			fmt.Println("Error " + nmRoute + " Zalora")
			if response.Message == "" {
				response.Message = nmRoute + " Zalora Gagal"
			}
			//response.Message = datasx.ErrorResponse.Head.ErrorMessage
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetDocumentShip(c *gin.Context) {
	var response response.ResponseCrud
	nmRoute := "GetDocumentShip"

	noorder := c.Param("noorder")
	noorderx := noorder
	NoOrderObj := strings.Split(noorderx, "-")
	if len(NoOrderObj) > 0 {
		noorderx = NoOrderObj[1]
	}

	DetailItem := GetDetailItemView(noorderx)

	fmt.Println(DetailItem)
	paramTambah := ""
	OrderItemIds := ""
	index := 1

	fmt.Println(len(DetailItem.OrderItem))
	if len(DetailItem.OrderItem) > 0 {

		for _, valuenya := range DetailItem.OrderItem {
			OrderItemIds += valuenya.OrderItemId
			if index < len(DetailItem.OrderItem) {
				OrderItemIds += ","
			}
			index++
		}

	} else {
		//single

		DetailItemSingle := GetDetailItemViewSingle(noorderx)
		fmt.Println(DetailItemSingle.OrderItem.ShipmentProvider)
		OrderItemIds = DetailItemSingle.OrderItem.OrderItemId
	}

	OrderItemIds = "[" + OrderItemIds + "]"
	response.Result = DetailItem

	paramTambah += "OrderItemIds#" + OrderItemIds
	paramTambah += "#DocumentType#shippingLabel"

	urlzalora, concatenated := AuthZalora("GetDocument", paramTambah)
	url := urlzalora + "?" + concatenated

	response.Message = url
	response.Result = DetailItem

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
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
		fmt.Println("gagal " + nmRoute + " Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		// response.Message = url
		response.Message = nmRoute + " Zalora Gagal"
		response.Result = concatenated

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonDocumentShipZalora(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		response.Message = ""

		var decodedByte, _ = base64.StdEncoding.DecodeString(datas.SuccessResponse.Body.Documents.Document.File)
		var decodedString = string(decodedByte)
		fmt.Println("decoded:", decodedString)

		response.Result = datas

	}

	c.JSON(http.StatusOK, response)
	return
}

///////////// V2 ////////////

func ReqToken(c *gin.Context) {

	urlzalora := os.Getenv("URL_API_ZALORA")

	var response response.ResponseCrud

	paths := urlzalora + "/oauth/client-credentials"

	var params = url.Values{}
	params.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", paths, bytes.NewBufferString(params.Encode()))

	token := base64.StdEncoding.EncodeToString([]byte(os.Getenv("CLIENT_ID_ZALORA") + ":" + os.Getenv("CLIENT_SECRET_ZALORA")))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+token)
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
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonToken(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Result = datas
		response.Message = paths

		tokenService.SaveTokenZalora(datas)

	}

	c.JSON(http.StatusOK, response)
	return
}

var ChannelName = "zalora"

func GetProcessedOrderAutoV2(statusZaloras string) {
	var response response.ResponseCrud
	statusZalora := "status_" + statusZaloras
	nmRoute := "GetProcessedOrderAutoV2"
	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders"

	limit := 100

	Tarik, _ := strconv.Atoi(os.Getenv("GET_ORDERS_TIME_ZALORA"))
	createdAfter := time.Now().Add(time.Duration(-Tarik) * time.Hour).Format("2006-01-02 15:04:05")

	parameter := ""
	parameter += "dateStart=" + url.QueryEscape(createdAfter)
	parameter += "&limit=" + strconv.Itoa(limit)
	parameter += "&section=" + statusZalora
	fmt.Println("======= Zalora " + statusZalora + " ======= ")

	paths = paths + "?" + parameter
	fmt.Println(paths)

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetOrders Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//datas := parseJsonSuccessResponse(data)
		datas := parseJsonOrdersZalora(data)

		response.Message = paths
		//response.Result = string(data)
		response.Result = datas

		datasErr := parseJsonErrorV2(data)

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)
			response.Result = string(data)
		}

		fmt.Println(len(datas.Items))

		for _, val := range datas.Items {
			GetOrderDetailV2Loop(strconv.Itoa(val.ID))
		}

		if len(datas.Items) > 0 {
			GetOrdersV2Loop(limit, limit, createdAfter, statusZalora)
		}

		if statusZalora == "pending" {
			GetProcessedOrderAutoV2("canceled")
		}

	}

}
func GetOrdersV2(c *gin.Context) {
	var response response.ResponseCrud
	statusZalora := "status_" + c.Param("status")
	nmRoute := "GetOrdersV2"
	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders"

	limit := 100

	Tarik, _ := strconv.Atoi(os.Getenv("GET_ORDERS_TIME_ZALORA"))
	createdAfter := time.Now().Add(time.Duration(-Tarik) * time.Hour).Format("2006-01-02 15:04:05")

	parameter := ""
	parameter += "dateStart=" + url.QueryEscape(createdAfter)
	parameter += "&limit=" + strconv.Itoa(limit)
	parameter += "&section=" + statusZalora
	fmt.Println("======= Zalora " + statusZalora + " ======= ")

	paths = paths + "?" + parameter
	fmt.Println(paths)

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetOrders Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//datas := parseJsonSuccessResponse(data)
		datas := parseJsonOrdersZalora(data)

		response.Message = paths
		//response.Result = string(data)
		response.Result = datas

		datasErr := parseJsonErrorV2(data)

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)
			response.Result = string(data)
		}

		fmt.Println(len(datas.Items))

		for _, val := range datas.Items {
			GetOrderDetailV2Loop(strconv.Itoa(val.ID))
		}

		if len(datas.Items) > 0 {
			GetOrdersV2Loop(limit, limit, createdAfter, statusZalora)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrdersV2Loop(limit, offset int, createdAfter, statusZalora string) {
	var response response.ResponseCrud
	nmRoute := "GetOrdersV2Loop"
	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders"

	parameter := ""
	parameter += "dateStart=" + url.QueryEscape(createdAfter)
	parameter += "&limit=" + strconv.Itoa(limit)
	parameter += "&offset=" + strconv.Itoa(offset)
	parameter += "&section=" + statusZalora
	fmt.Println("======= Zalora " + statusZalora + " ======= ")

	paths = paths + "?" + parameter
	fmt.Println(paths)

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetOrdersV2Loop Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonOrdersZalora(data)

		response.Message = paths
		response.Result = datas

		datasErr := parseJsonErrorV2(data)

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)
			response.Result = string(data)
		}

		fmt.Println(len(datas.Items))

		for _, val := range datas.Items {
			GetOrderDetailV2Loop(strconv.Itoa(val.ID))
		}

		offset += limit
		if len(datas.Items) > 0 {
			GetOrdersV2Loop(limit, offset, createdAfter, statusZalora)
		}

	}

}

func GetOrderDetailV2(c *gin.Context) {
	nmRoute := "GetOrderDetailV2"
	var response response.ResponseCrud

	orderid := c.Param("orderid")
	NoOrderObj := strings.Split(orderid, "-")
	if len(NoOrderObj) > 0 {
		orderid = NoOrderObj[1]
	}

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders/" + orderid

	parameter := ""
	parameter += "orderId=" + url.QueryEscape(orderid)

	paths = paths + "?" + parameter

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetOrders Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonOrderDetailZaloraV2(data)
		datasErr := parseJsonErrorV2(data)
		response.Message = ""
		response.Result = datas
		response.Result = string(data)

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)
			response.Result = string(data)
			response.Message = datasErr.Message + " (ZALORA)"
			response.ResponseDesc = datasErr.Message + " (ZALORA)"
		} else {

			pending := "pending"
			if datas.StatusList.Pending == 0 {
				pending = ""
			}
			if c.Param("orderid") == "243497665-8189098" {
				pending = "pending"
			}
			response.ResponseDesc = pending
			response.Result = datas
			//update  amount
			tokenService.UpdateAmountZalora(c.Param("orderid"), datas)
			//response.Result = cekAmount
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrderDetailV2Loop(orderid string) {
	var response response.ResponseCrud

	nmRoute := "GetOrderDetailV2Loop"

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders/" + orderid

	parameter := ""
	parameter += "orderId=" + url.QueryEscape(orderid)

	paths = paths + "?" + parameter

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetOrders Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonOrderDetailZaloraV2(data)

		response.Message = paths
		response.Result = datas

		fmt.Println("================")
		tokenService.SaveSalesOrderZaloraV2(strconv.Itoa(datas.ID), datas)
		fmt.Println("================")
		datasErr := parseJsonErrorV2(data)

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)
			response.Result = string(data)
		}

	}

}

func GetOrderDetailItemV2(c *gin.Context) {
	var response response.ResponseCrud

	nmRoute := "GetOrderDetailItemV2"

	orderid := c.Param("orderitem")

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders/" + orderid

	parameter := ""
	parameter += "orderId=" + url.QueryEscape(orderid)

	paths = paths + "?" + parameter

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetOrders Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonOrderDetailZaloraV2(data)
		datasErr := parseJsonErrorV2(data)

		response.Message = paths
		response.Result = datas

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)
			response.Result = string(data)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}
func GetProductsAutoV2(statusproduct string, wg *sync.WaitGroup) {
	var response response.ResponseCrud

	nmRoute := "GetProductsAutoV2"

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	//statusproduct := c.Param("status")

	paths := urlzalora + "/v2/product-sets"

	limit := 100
	offset := 0
	// inactive-all
	// deleted-all
	// active
	// pending
	// rejected

	Tarik, _ := strconv.Atoi(os.Getenv("GET_ORDERS_TIME_ZALORA"))
	createdAfter := time.Now().Add(time.Duration(-Tarik) * time.Hour).Format("2006-01-02")

	parameter := ""
	parameter += "createDateStart=" + url.QueryEscape(createdAfter)
	parameter += "&updateDateStart=" + url.QueryEscape(createdAfter)
	parameter += "&limit=" + strconv.Itoa(limit)
	parameter += "&offset=" + strconv.Itoa(offset)
	parameter += "&status=" + statusproduct

	paths = paths + "?" + parameter
	fmt.Println(paths)

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetProductsV2 Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonProductsV2Zalora(data)

		response.Message = paths
		//response.Result = string(data)
		response.Result = datas

		datasErr := parseJsonErrorV2(data)

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)
			response.Result = string(data)
		}

		if len(datas.Items) > 0 {
			//tokenService.SaveSkuMappingZaloraV2(datas, statusproduct)
			for _, valItems := range datas.Items {
				GetProductStockV2Loop(strconv.Itoa(valItems.ID), statusproduct)
			}

			offset = limit
			GetProductsV2Loop(limit, offset, createdAfter, statusproduct)
		}

	}

	if statusproduct == "inactive-all" {
		GetProductsV2Loop(limit, 0, createdAfter, "deleted-all")
		GetProductsV2Loop(limit, 0, createdAfter, "active")
		GetProductsV2Loop(limit, 0, createdAfter, "pending")
		GetProductsV2Loop(limit, 0, createdAfter, "rejected")
		GetProductsV2Loop(limit, 0, createdAfter, "sold-out")

	}

	// inactive-all
	// deleted-all
	// active
	// pending
	// rejected

	wg.Done()
}

func GetProductsV2(c *gin.Context) {
	var response response.ResponseCrud

	nmRoute := "GetProductsV2"

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	statusproduct := c.Param("status")

	paths := urlzalora + "/v2/product-sets"

	limit := 100
	offset := 0
	// inactive-all
	// deleted-all
	// active
	// pending
	// rejected

	Tarik, _ := strconv.Atoi(os.Getenv("GET_ORDERS_TIME_ZALORA"))
	createdAfter := time.Now().Add(time.Duration(-Tarik) * time.Hour).Format("2006-01-02")
	//createdBefore := time.Now().Format("2006-01-02")

	parameter := ""
	parameter += "createDateStart=" + url.QueryEscape(createdAfter)
	//parameter += "&createDateEnd=" + url.QueryEscape(createdBefore)

	parameter += "&updateDateStart=" + url.QueryEscape(createdAfter)
	//parameter += "&updateDateEnd=" + url.QueryEscape(createdBefore)

	//parameter += "&parentSku=08099534"

	parameter += "&limit=" + strconv.Itoa(limit)
	parameter += "&offset=" + strconv.Itoa(offset)
	parameter += "&status=" + statusproduct

	paths = paths + "?" + parameter
	fmt.Println(paths)

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetProductsV2 Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonProductsV2Zalora(data)

		response.Message = paths
		//response.Result = string(data)
		response.Result = datas
		//response.Result = string(data)

		datasErr := parseJsonErrorV2(data)

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)
			response.Result = string(data)
		}

		// //coba
		// if len(datas.Items) > 0 {
		// 	for _, valItems := range datas.Items {
		// 		GetProductStockV2Loop(strconv.Itoa(valItems.ID), statusproduct)
		// 	}
		// }

		if len(datas.Items) > 0 {
			for _, valItems := range datas.Items {
				GetProductStockV2Loop(strconv.Itoa(valItems.ID), statusproduct)
			}

			//tokenService.SaveSkuMappingZaloraV2(datas, statusproduct)
			offset = limit
			GetProductsV2Loop(limit, offset, createdAfter, statusproduct)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsV2Loop(limit, offset int, createdAfter, statusproduct string) {
	var response response.ResponseCrud

	nmRoute := "GetProductsV2Loop"

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/product-sets"

	parameter := ""
	parameter += "createDateStart=" + url.QueryEscape(createdAfter)
	parameter += "&updateDateStart=" + url.QueryEscape(createdAfter)
	parameter += "&limit=" + strconv.Itoa(limit)
	parameter += "&offset=" + strconv.Itoa(offset)
	parameter += "&status=" + statusproduct

	paths = paths + "?" + parameter
	fmt.Println(paths)

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetProductsV2 Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonProductsV2Zalora(data)

		response.Message = paths
		//response.Result = string(data)
		response.Result = datas

		datasErr := parseJsonErrorV2(data)

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)
			response.Result = string(data)
		}

		if len(datas.Items) > 0 {
			//tokenService.SaveSkuMappingZaloraV2(datas, statusproduct)
			for _, valItems := range datas.Items {
				GetProductStockV2Loop(strconv.Itoa(valItems.ID), statusproduct)
			}

			offset += limit
			GetProductsV2Loop(limit, offset, createdAfter, statusproduct)
		}

	}

}

func UpdateStatusZaloraV2(c *gin.Context) {
	var response response.ResponseCrud
	nmRoute := "UpdateStatusZaloraV2"

	statusZalora := c.Param("status")
	if statusZalora == "ReadyToShip" {
		statusZalora = "set-to-packed-by-marketplace"
	}

	fmt.Println(statusZalora)
	noorder := c.Param("noorder")
	noorderx := noorder
	NoOrderObj := strings.Split(noorderx, "-")
	if len(NoOrderObj) > 0 {
		noorderx = NoOrderObj[1]
	}

	DetailItem := GetOrderDetailV2View(noorderx)

	fmt.Println(DetailItem)
	//paramTambah := ""
	OrderItemIds := ""
	index := 1
	ShipmentProvider := ""
	deliveryType := "pickup"

	fmt.Println(len(DetailItem.Items))
	if len(DetailItem.Items) > 0 {

		for _, valuenya := range DetailItem.Items {
			//OrderItemIds += "{" + `"id":` + strconv.Itoa(valuenya.ID) + "}"
			if statusZalora == "set-to-packed-by-marketplace" {
				OrderItemIds += "{" + `"orderItemId":` + strconv.Itoa(valuenya.ID) + "}"
			} else {
				OrderItemIds += "{" + `"id":` + strconv.Itoa(valuenya.ID) + "}"
			}
			//OrderItemIds += strconv.Itoa(valuenya.ID)
			if index < len(DetailItem.Items) {
				OrderItemIds += ","
			}
			index++
		}

		ShipmentProvider = DetailItem.Items[0].Shipment.Provider.Name
		// deliveryType = DetailItem.Items[0].Shipment.Type
		// if deliveryType == "dropshipping" {
		// 	deliveryType = "dropship"
		// }

	}

	deliveryType = "dropship"

	OrderItemIds = `"orderItems": [` + OrderItemIds + "]"
	//OrderItemIds = `"orderItemIds": [` + OrderItemIds + "]"

	deliveryType = ` ,"deliveryType": "` + deliveryType + `"`
	//deliveryType = ""

	if statusZalora == "set-to-packed-by-marketplace" {
		ShipmentProvider = `,"shippingProvider": "` + ShipmentProvider + `"`
	} else {
		ShipmentProvider = "" //RTS
	}
	//ShipmentProvider = `,"shippingProvider": "` + ShipmentProvider + `"`
	//ShipmentProvider = ""

	jsonString := `{
			` + OrderItemIds + `
			` + deliveryType + `
			` + ShipmentProvider + `
			
	}`

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders/statuses/" + statusZalora
	//paths := urlzalora + "/v2/order-pickup-requests"

	fmt.Println(jsonString)
	fmt.Println(paths)

	var body_url = []byte(jsonString)

	// response.ResponseCode = http.StatusOK
	// response.ResponseDesc = enums.SUCCESS
	// response.ResponseTime = utils.DateToStdNow()
	// response.Message = paths
	// response.Result = jsonString

	req, err := http.NewRequest("POST", paths, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal UpdateStatusZaloraV2 Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "Koneksi Zalora Gagal"
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//datas := parseJsonProductsV2Zalora(data)
		datasErr := parseJsonErrorV3(data)

		psnerror := "picking gagal. (ZALORA)"

		fmt.Println("===== " + nmRoute + " =====")
		fmt.Println(string(data))
		fmt.Println("+++++ " + nmRoute + " +++++")

		if len(datasErr.Errors) > 0 {
			tokenService.SaveErrorString("zalora", datasErr.Errors[0].Detail, nmRoute)
			psnerror = datasErr.Errors[0].Detail + " (ZALORA)"
		}

		if statusZalora == "set-to-packed-by-marketplace" {

			dataspackded := parseJsonPackedV2(data)
			if len(dataspackded.OrderItemIds) > 0 {

				cekdoc := GetDocumentV2ZaloraLoop(noorder)
				if cekdoc == "" {

					time.Sleep(2 * time.Second)

					statusZalora = "set-to-ready-to-ship"
					hasil := UpdateStatusZaloraLoopV2(statusZalora, noorder)
					if hasil != "" {
						psnerror = hasil
					} else {
						psnerror = ""
					}

				} else {
					psnerror = "Create Document Shipping Gagal"
				}

			}

		}

		response.Message = psnerror
		response.Result = string(data)
		fmt.Println(datasErr)
		//response.Result = datas

	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStatusZaloraLoopV2(statusZalora, noorder string) string {
	var response response.ResponseCrud
	nmRoute := "UpdateStatusZaloraLoopV2"
	balik := "picking gagal"

	noorderx := noorder
	NoOrderObj := strings.Split(noorderx, "-")
	if len(NoOrderObj) > 0 {
		noorderx = NoOrderObj[1]
	}

	DetailItem := GetOrderDetailV2View(noorderx)

	OrderItemIds := ""
	index := 1
	ShipmentProvider := ""
	deliveryType := "pickup"

	if len(DetailItem.Items) > 0 {

		for _, valuenya := range DetailItem.Items {
			//OrderItemIds += "{" + `"id":` + strconv.Itoa(valuenya.ID) + "}"
			if statusZalora == "set-to-packed-by-marketplace" {
				OrderItemIds += "{" + `"orderItemId":` + strconv.Itoa(valuenya.ID) + "}"
			} else {
				OrderItemIds += "{" + `"id":` + strconv.Itoa(valuenya.ID) + "}"
			}
			//OrderItemIds += strconv.Itoa(valuenya.ID)
			if index < len(DetailItem.Items) {
				OrderItemIds += ","
			}
			index++
		}

		ShipmentProvider = DetailItem.Items[0].Shipment.Provider.Name
		// deliveryType = DetailItem.Items[0].Shipment.Type
		// if deliveryType == "dropshipping" {
		// 	deliveryType = "dropship"
		// }

	}

	deliveryType = "dropship"

	OrderItemIds = `"orderItems": [` + OrderItemIds + "]"
	//OrderItemIds = `"orderItemIds": [` + OrderItemIds + "]"

	deliveryType = ` ,"deliveryType": "` + deliveryType + `"`
	//deliveryType = ""

	if statusZalora == "set-to-packed-by-marketplace" {
		ShipmentProvider = `,"shippingProvider": "` + ShipmentProvider + `"`
	} else {
		ShipmentProvider = "" //RTS
	}
	//ShipmentProvider = `,"shippingProvider": "` + ShipmentProvider + `"`
	//ShipmentProvider = ""

	jsonString := `{
			` + OrderItemIds + `
			` + deliveryType + `
			` + ShipmentProvider + `
			
	}`

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders/statuses/" + statusZalora

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", paths, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal " + nmRoute + " Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//datas := parseJsonProductsV2Zalora(data)
		datasErr := parseJsonErrorV3(data)

		fmt.Println("===== " + nmRoute + " =====")
		fmt.Println(string(data))
		fmt.Println("+++++ " + nmRoute + " +++++")

		psnerror := ""
		if len(datasErr.Errors) > 0 { //
			tokenService.SaveErrorString("zalora", datasErr.Errors[0].Detail, nmRoute)
			psnerror = datasErr.Errors[0].Detail
		}

		if statusZalora == "set-to-ready-to-ship" {
			dataspackded := parseJsonPackedV2(data)
			fmt.Println("===== " + nmRoute + " LAGI =====")
			fmt.Println(dataspackded)
			fmt.Println("+++++ " + nmRoute + " LAGI +++++")

			if len(dataspackded.OrderItemIds) > 0 {
				// psnerror = "picking gagal"
			} else {
				psnerror = "picking gagal."
			}
		}

		if psnerror != "" {
			response.Message = psnerror
		} else {
			psnerror = ""
		}

		balik = psnerror

		fmt.Println("===== " + nmRoute + " LAGI LAGI =====")
		fmt.Println(balik)
		fmt.Println("+++++ " + nmRoute + " LAGI LAGI +++++")

		// response.Message = paths
		response.Result = string(data)
		//fmt.Println(datasErr)
		//response.Result = datas

	}

	return balik
}

func GetDocumentV2Zalora(c *gin.Context) {
	var response response.ResponseCrud
	nmRoute := "GetDocumentV2Zalora"

	statusZalora := c.Param("status")
	fmt.Println(statusZalora)
	noorder := c.Param("noorder")
	noorderx := noorder
	NoOrderObj := strings.Split(noorderx, "-")
	if len(NoOrderObj) > 0 {
		noorderx = NoOrderObj[1]
	}

	OrderItemIds := ` "orderIds": [` + noorderx
	//index := 1
	docnya := ` ,"documentType": "shippingLabel" `

	OrderItemIds += "]"

	jsonString := `{
		` + OrderItemIds + `
		` + docnya + `
		
		}`

	//fmt.Println(jsonString)
	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders/export-document"

	var body_url = []byte(jsonString)
	req, err := http.NewRequest("POST", paths, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetDocumentV2Zalora Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "Koneksi Zalora Gagal"
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonDocZaloraV2(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Result = string(data)
		response.Result = datas
	}

	c.JSON(http.StatusOK, response)
	return

}

func GetDocumentV2ZaloraLoop(noorder string) string {
	var response response.ResponseCrud
	nmRoute := "GetDocumentV2ZaloraLoop"

	status := ""

	noorderx := noorder
	NoOrderObj := strings.Split(noorderx, "-")
	if len(NoOrderObj) > 0 {
		noorderx = NoOrderObj[1]
	}

	OrderItemIds := ` "orderIds": [` + noorderx
	//index := 1
	docnya := ` ,"documentType": "shippingLabel" `

	OrderItemIds += "]"

	jsonString := `{
		` + OrderItemIds + `
		` + docnya + `
		
		}`

	//fmt.Println(jsonString)
	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders/export-document"

	var body_url = []byte(jsonString)
	req, err := http.NewRequest("POST", paths, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetDocumentV2Zalora Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "Koneksi Zalora Gagal"
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
		status = "koneksi"
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonDocZaloraV2(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Result = string(data)
		response.Result = datas

		if datas.Status == "" {
			status = "gagal"
		}

		fmt.Println("===== " + nmRoute + " =====")
		fmt.Println(string(data))
		fmt.Println(status)
		fmt.Println("+++++ " + nmRoute + " +++++")

	}

	return status

}

func GetOrderDetailV2View(orderid string) models.OrderDetailZaolraV2 {
	var obj models.OrderDetailZaolraV2
	nmRoute := "GetOrderDetailV2View"
	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders/" + orderid

	parameter := ""
	parameter += "orderId=" + url.QueryEscape(orderid)

	paths = paths + "?" + parameter

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		obj = parseJsonOrderDetailZaloraV2(data)

		datasErr := parseJsonErrorV2(data)

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)

		}
	}

	return obj

}

func GetResiZaloraV2(c *gin.Context) {

	nmRoute := "GetResiZaloraV2"
	var response response.ResponseCrud

	orderid := c.Param("noorder")
	NoOrderObj := strings.Split(orderid, "-")
	if len(NoOrderObj) > 0 {
		orderid = NoOrderObj[1]
	}
	fmt.Println(orderid)

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/orders/" + orderid

	parameter := ""
	parameter += "orderId=" + url.QueryEscape(orderid)

	paths = paths + "?" + parameter

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal GetOrders Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = paths
		tokenService.SaveErrorString("zalora", "koneksi", nmRoute)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonOrderDetailZaloraV2(data)
		datasErr := parseJsonErrorV2(data)
		response.Message = ""
		response.Result = datas

		if datasErr.Message != "" {
			tokenService.SaveErrorString("zalora", datasErr.Message, nmRoute)
			response.Result = string(data)
			response.Message = datasErr.Message + " (ZALORA)"
		} else {

			response.Message = "Resi Belum Tersedia (ZALORA)"

			if len(datas.Items) > 0 {
				resinya := ""

				if datas.Items[0].Shipment.TrackingCode != "" {
					resinya = datas.Items[0].Shipment.TrackingCode
					response.Message = ""
				}

				response.ResponseDesc = resinya

			}

		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStockProductV2(c *gin.Context) {
	var response response.ResponseCrud

	nmRoute := "UpdateStockProductV2"

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/stock/product"

	productid := c.Param("productid")
	stock := c.Param("stock")
	sku := c.Param("sku")

	xmlString := ""

	xmlString = `[
			{
			  "productId": ` + productid + `,
			  "quantity": ` + stock + `
			}
		  ]`

	//fmt.Println(xmlString)
	var objRest models.TableLogStock
	objRest.UuidLog = uuid.New().String()
	objRest.ChannelCode = os.Getenv("KODE_ZALORA")

	stocks, _ := strconv.ParseFloat(stock, 64)

	objRest.Stock = stocks
	objRest.Body = xmlString
	objRest.Sku = sku

	var body_url = []byte(xmlString)

	req, err := http.NewRequest("PUT", paths, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal " + nmRoute + " Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "gagal " + nmRoute + " Zalora"
		response.Result = "ERROR"
		objRest.Response = nmRoute + " Gagal Koneksi"
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonStockZaloraV2(data)
		response.Message = paths
		response.Result = string(data)
		response.Result = datas
		objRest.Response = string(data)
		if len(datas) < 1 {
			fmt.Println("gagal update stock")
			response.Result = string(data)
		} else {
			fmt.Println("berhasil update stock")
			response.Message = ""
		}

	}

	objRest.CreatedBy = "API"
	objRest.CreatedDate = time.Now()
	tokenRepository.SaveStockAPI(objRest)

	c.JSON(http.StatusOK, response)
	return
}

func GetProductStockV2(c *gin.Context) {
	var response response.ResponseCrud

	nmRoute := "GetProductStockV2"

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/stock/product-set/"

	productid := c.Param("productid")

	paths += productid

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal " + nmRoute + " Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "gagal " + nmRoute + " Zalora"
		response.Result = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonDetailStockV2(data)
		response.Message = paths
		response.Result = datas
		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductBySKUV2(c *gin.Context) {
	var response response.ResponseCrud

	nmRoute := "GetProductBySKUV2"

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	Skunya := c.Param("sku")
	statusnya := c.Param("status")

	paths := urlzalora + "/v2/product-sets"
	parameter := ""
	parameter += "&parentSku=" + Skunya
	parameter += "&status=" + statusnya

	paths += "?" + parameter

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal " + nmRoute + " Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "gagal " + nmRoute + " Zalora"
		response.Result = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonProductsV2Zalora(data)
		response.Message = paths
		response.Result = datas
		response.Result = string(data)

		// if len(datas.Items) > 0 {
		// 	for _, valItems := range datas.Items {
		// 		GetProductStockV2Loop(strconv.Itoa(valItems.ID), statusproduct)
		// 	}

		// }

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductStockV2Loop(productid, statusproduct string) {
	var response response.ResponseCrud

	nmRoute := "GetProductStockV2Loop"

	urlzalora := os.Getenv("URL_API_ZALORA")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	paths := urlzalora + "/v2/stock/product-set/"

	paths += productid

	req, err := http.NewRequest("GET", paths, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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
		fmt.Println("gagal " + nmRoute + " Zalora")
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = "gagal " + nmRoute + " Zalora"
		response.Result = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		datas := parseJsonDetailStockV2(data)
		response.Message = paths
		response.Result = string(data)
		response.Result = datas

		if len(datas) > 0 {
			//fmt.Println(productid + " | ada detail")
			tokenService.SaveSkuMappingZaloraV2(datas, statusproduct)
		}

	}

}
