package tokpedController

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/helpers"
	"github.com/rals/dearme-channel/models"
	"github.com/rals/dearme-channel/models/response"
	"github.com/rals/dearme-channel/repositories/stockRepository"
	"github.com/rals/dearme-channel/repositories/tokenRepository"
	"github.com/rals/dearme-channel/services/tokenService"
	"github.com/rals/dearme-channel/utils"
	// "encoding/json"
)

var ChannelName = "tokopedia"

func AuthTokped() string {
	login := os.Getenv("CLIENT_ID_TOKPED") + ":" + os.Getenv("CLIENT_KEY_TOKPED")
	token := base64.StdEncoding.EncodeToString([]byte(login))
	return token
}

func parseJsonErrorTokped(jsonBuffer []byte) models.ErrorTokped {

	ErrorTokped := models.ErrorTokped{}

	err := json.Unmarshal(jsonBuffer, &ErrorTokped)
	if err != nil {
		return ErrorTokped
	}

	// the array is now filled with users
	return ErrorTokped

}

func parseJsonTokenTokped(jsonBuffer []byte) models.TokenTokped {

	TokenTokped := models.TokenTokped{}

	err := json.Unmarshal(jsonBuffer, &TokenTokped)
	if err != nil {
		return TokenTokped
	}

	// the array is now filled with users
	return TokenTokped

}

func parseJsonProductTokped(jsonBuffer []byte) models.ProductTokped {

	ProductTokped := models.ProductTokped{}

	err := json.Unmarshal(jsonBuffer, &ProductTokped)
	if err != nil {
		return ProductTokped
	}

	// the array is now filled with users
	return ProductTokped

}

func parseJsonProductDetailTokped(jsonBuffer []byte) models.ProductDetailTokped {

	ProductDetailTokped := models.ProductDetailTokped{}

	err := json.Unmarshal(jsonBuffer, &ProductDetailTokped)
	if err != nil {
		return ProductDetailTokped
	}

	// the array is now filled with users
	return ProductDetailTokped

}

func parseJsonProductTokped2(jsonBuffer []byte) models.ProductTokped2 {

	ProductTokped2 := models.ProductTokped2{}

	err := json.Unmarshal(jsonBuffer, &ProductTokped2)
	if err != nil {
		return ProductTokped2
	}

	// the array is now filled with users
	return ProductTokped2

}

func parseJsonProductTokped3(jsonBuffer []byte) models.ProductTokped3 {

	ProductTokped3 := models.ProductTokped3{}

	err := json.Unmarshal(jsonBuffer, &ProductTokped3)
	if err != nil {
		return ProductTokped3
	}

	// the array is now filled with users
	return ProductTokped3

}

func parseJsonStockTokped(jsonBuffer []byte) models.StockTokped {

	StockTokped := models.StockTokped{}

	err := json.Unmarshal(jsonBuffer, &StockTokped)
	if err != nil {
		return StockTokped
	}

	// the array is now filled with users
	return StockTokped

}

func parseJsonOrdersTokped(jsonBuffer []byte) models.OrdersTokped {

	OrdersTokped := models.OrdersTokped{}

	err := json.Unmarshal(jsonBuffer, &OrdersTokped)
	if err != nil {
		return OrdersTokped
	}

	// the array is now filled with users
	return OrdersTokped

}

// func parseJsonOrderDetailTokped(jsonBuffer []byte) models.OrderDetailTokped {

// 	OrderDetailTokped := models.OrderDetailTokped{}

// 	err := json.Unmarshal(jsonBuffer, &OrderDetailTokped)
// 	if err != nil {
// 		return OrderDetailTokped
// 	}

// 	// the array is now filled with users
// 	return OrderDetailTokped

// }

func parseJsonOrderDetailTokped(jsonBuffer []byte) models.OrderDetailTokpedSingle {

	OrderDetailTokpedSingle := models.OrderDetailTokpedSingle{}

	err := json.Unmarshal(jsonBuffer, &OrderDetailTokpedSingle)
	if err != nil {
		return OrderDetailTokpedSingle
	}

	// the array is now filled with users
	return OrderDetailTokpedSingle

}

func parseJsonDetailBookingTokped(jsonBuffer []byte) models.DetailBookingTokped {

	DetailBookingTokped := models.DetailBookingTokped{}

	err := json.Unmarshal(jsonBuffer, &DetailBookingTokped)
	if err != nil {
		return DetailBookingTokped
	}

	// the array is now filled with users
	return DetailBookingTokped

}

func GetTokenTokped(c *gin.Context) {

	var response response.ResponseCrud

	urltokped := os.Getenv("URL_AUTH_TOKPED")

	token := AuthTokped()
	// var jsonString = `{"refresh_token":"TEST"}`

	// var body_url = []byte(jsonString)
	urltokped += "?grant_type=client_credentials"
	req, err := http.NewRequest("POST", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Basic "+token)
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
		response.Message = "Req Token TOKOPEDIA GAGAL (connection)"
		response.Result = err

	} else {

		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonTokenTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)
		var objToken models.TokenTokped

		if datas.AccessToken != "" {
			objToken.AccessToken = datas.AccessToken
			objToken.EventCode = datas.EventCode
			objToken.SqCheck = datas.SqCheck
			objToken.ExpiresIn = datas.ExpiresIn
			objToken.LastLoginType = datas.LastLoginType
			objToken.TokenType = datas.TokenType

			tokenService.SaveTokenTokped(objToken, ChannelName)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			response.Message = "Save Token Sukses"
			response.Result = datas
		}

		fmt.Println(string(data))

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsAuto(wg *sync.WaitGroup) {
	nmRoute := "GetProductsAuto"
	var response response.ResponseCrud
	fmt.Println("mulai product tokped " + time.Now().String())
	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	shopid := os.Getenv("SHOP_ID_TOKPED")

	rows := os.Getenv("LIMIT_PRODUCT_TOKPED")
	starts := 1

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/list?shop_id=" + shopid + "&rows=" + rows + "&start=" + strconv.Itoa(starts)
	urltokped += "&order_by=10"

	//fmt.Println(urltokped)
	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println("======= GET Products TOKPED =======")

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductTokped(data)
		//datas := parseJsonProductTokped2(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		// fmt.Println(datas.Data.TotalData)
		// fmt.Println(len(datas.Data.Products))
		index := 0
		for _, val := range datas.Data.Products {

			for _, valChilds := range val.Childs {
				index++
				GetProductDetailLoop(strconv.Itoa(int(valChilds)))
			}
			if len(val.Childs) < 1 {
				tokenService.SaveSkuMappingSingleTokped(datas)
			}
			// index++
			// GetProductDetailLoop(strconv.Itoa(int(val.ID)))
		}
		//fmt.Println(len(datas.Data.Products))
		if len(datas.Data.Products) > 0 {
			//fmt.Println("rows " + rows)
			rowInt, _ := strconv.Atoi(rows)
			starts++

			GetProductsLoop(rowInt, starts, index)
		}

		//fmt.Println("total " + strconv.Itoa(index))
	}

	//fmt.Println("======= SELESAI Products TOKPED =======")
	fmt.Println("selesai product tokped " + time.Now().String())
	wg.Done()
}

func GetProducts(c *gin.Context) {
	nmRoute := "GetProducts"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	shopid := os.Getenv("SHOP_ID_TOKPED")

	rows := os.Getenv("LIMIT_PRODUCT_TOKPED")
	starts := 1

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/list?shop_id=" + shopid + "&rows=" + rows + "&start=" + strconv.Itoa(starts)
	urltokped += "&order_by=10"

	fmt.Println(urltokped)
	token := fmt.Sprintf("%v", ObjToken.Value1)
	fmt.Println("======= GET Products TOKPED =======")
	fmt.Println(ObjToken)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductTokped(data)
		//datas := parseJsonProductTokped2(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		// fmt.Println(datas.Data.TotalData)
		// fmt.Println(len(datas.Data.Products))
		index := 0
		for _, val := range datas.Data.Products {

			for _, valChilds := range val.Childs {
				index++
				GetProductDetailLoop(strconv.Itoa(int(valChilds)))
			}
			if len(val.Childs) < 1 {
				tokenService.SaveSkuMappingSingleTokped(datas)
			}
			// index++
			// GetProductDetailLoop(strconv.Itoa(int(val.ID)))
		}
		fmt.Println(len(datas.Data.Products))
		if len(datas.Data.Products) > 0 {
			fmt.Println("rows " + rows)
			rowInt, _ := strconv.Atoi(rows)
			starts++

			GetProductsLoop(rowInt, starts, index)
		}

		fmt.Println("total " + strconv.Itoa(index))
	}

	fmt.Println("======= SELESAI Products TOKPED =======")

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsLoop(rows int, starts int, index int) {
	nmRoute := "GetProductsLoop"
	var response response.ResponseCrud
	//fmt.Println("masuk loop")
	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	shopid := os.Getenv("SHOP_ID_TOKPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/list?shop_id=" + shopid + "&rows=" + strconv.Itoa(rows) + "&start=" + strconv.Itoa(starts)
	urltokped += "&order_by=10"
	//fmt.Println(urltokped)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		RespHead := resp.Header
		datas := parseJsonProductTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		// fmt.Println(datas.Data.TotalData)
		// fmt.Println(len(datas.Data.Products))

		if len(datas.Data.Products) < 1 {
			//fmt.Println(string(data))

		}
		for _, val := range datas.Data.Products {
			for _, valChilds := range val.Childs {
				index++
				GetProductDetailLoop(strconv.Itoa(int(valChilds)))
			}
			if len(val.Childs) < 1 {
				tokenService.SaveSkuMappingSingleTokped(datas)
			}
			// index++
			// GetProductDetailLoop(strconv.Itoa(int(val.ID)))
		}

		if len(datas.Data.Products) > 0 {
			starts++
			if RespHead["X-Ratelimit-Remaining"][0] == "0" {
				//fmt.Println("masuk sleep")
				time.Sleep(30 * time.Second)
			}
			//fmt.Println("total index " + strconv.Itoa(index))
			GetProductsLoop(rows, starts, index)
		}

	}

}

func GetProductDetailLoop(ProductId string) {
	nmRoute := "GetProductDetailLoop"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/info"
	urltokped += "?product_id=" + ProductId
	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		RespHead := resp.Header
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		if RespHead["X-Ratelimit-Remaining"][0] == "0" {
			//fmt.Println("masuk sleep detail loop")
			time.Sleep(10 * time.Second)
		}

		psnError := ""
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		if len(datas.Data) > 0 {
			tokenService.SaveSkuMappingTokped(datas, ProductId)
		}

	}

}

func GetProductDetail(c *gin.Context) {
	nmRoute := "GetProductDetail"
	ProductId := c.Param("id")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/info"
	urltokped += "?product_id=" + ProductId
	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		response.Message = psnError

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetShopInfo(c *gin.Context) {
	nmRoute := "GetShopInfo"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/v1/shop/fs/" + fsid + "/shop-info"

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urltokped
		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsV2(c *gin.Context) {
	nmRoute := "GetProductsV2"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	rows := os.Getenv("LIMIT_PRODUCT_TOKPED")
	starts := 1

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/v2/products/fs/" + fsid + "/" + strconv.Itoa(starts) + "/" + rows + ""

	fmt.Println(urltokped)
	token := fmt.Sprintf("%v", ObjToken.Value1)
	fmt.Println("======= GET Products V2 TOKPED =======")

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductTokped2(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		index := 0
		cek := 0
		for _, val := range datas.Data {
			index++
			GetProductDetailLoop(strconv.Itoa(int(val.ProductID)))
		}
		fmt.Println(len(datas.Data))
		if len(datas.Data) > 0 {
			fmt.Println("rows " + rows)
			rowInt, _ := strconv.Atoi(rows)
			starts++

			GetProductsLoopV2(rowInt, starts, index, cek)
		}

		fmt.Println("total " + strconv.Itoa(index))
	}

	fmt.Println("======= SELESAI Products V2 TOKPED =======")

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsLoopV2(rows int, starts int, index int, cek int) {
	nmRoute := "GetProductsLoopV2"
	var response response.ResponseCrud
	fmt.Println("masuk loop")
	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/v2/products/fs/" + fsid + "/" + strconv.Itoa(starts) + "/" + strconv.Itoa(rows) + ""
	fmt.Println(urltokped)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		RespHead := resp.Header
		datas := parseJsonProductTokped2(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		// fmt.Println(datas.Data.TotalData)
		// fmt.Println(len(datas.Data.Products))

		if len(datas.Data) < 1 {
			fmt.Println(string(data))
			if cek < 4 {
				cek++
				fmt.Println("masuk sleep cek")
				time.Sleep(10 * time.Second)
				GetProductsLoopV2(rows, starts, index, cek)
			}
		}
		for _, val := range datas.Data {
			index++
			GetProductDetailLoop(strconv.Itoa(int(val.ProductID)))
		}

		if len(datas.Data) > 0 {
			starts++

			if RespHead["X-Ratelimit-Remaining"][0] == "0" {
				fmt.Println("masuk sleep")
				time.Sleep(30 * time.Second)
			}
			fmt.Println("total index " + strconv.Itoa(index))
			cek = 0
			GetProductsLoopV2(rows, starts, index, cek)
		}

	}

}

func GetProductsV3(c *gin.Context) {
	nmRoute := "GetProductsV3"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	shopid := os.Getenv("SHOP_ID_TOKPED")
	rows := os.Getenv("LIMIT_PRODUCT_TOKPED")
	starts := 1
	//rows = "50"

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/info?shop_id=" + shopid

	urltokped += "&page=" + strconv.Itoa(starts)
	urltokped += "&per_page=" + rows

	fmt.Println(urltokped)
	token := fmt.Sprintf("%v", ObjToken.Value1)
	fmt.Println("======= GET Products V3 TOKPED =======")

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductTokped3(data)
		//datas := parseJsonProductDetailTokped(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		index := 0

		psnError := ""
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		for _, val := range datas.Data {
			index++
			fmt.Println("ID Parent" + strconv.Itoa(int(val.Basic.ProductID)))
			fmt.Println("sku parent " + val.Other.Sku)
			if len(val.Variant.ChildrenID) < 1 {
				fmt.Println("save single mapping sku")
				//save single mapping sku

				skuparam := val.Other.Sku
				if skuparam != "" {
					//tokenService.SaveSkuMappingSingleTokpedNew(datas, skuparam)
				}

			} else {
				fmt.Println("ada children id")
			}
			// for _, valVariant := range val.Variant.ChildrenID {
			// 	GetProductDetailLoop(strconv.Itoa(int(valVariant)))
			// }
			//GetProductDetailLoop(strconv.Itoa(int(val.Basic.ProductID)))
		}
		// fmt.Println(len(datas.Data))
		// if len(datas.Data) > 0 {
		// 	fmt.Println("rows " + rows)
		// 	rowInt, _ := strconv.Atoi(rows)
		// 	starts++

		// 	GetProductsLoopV3(rowInt, starts, index)
		// }

		fmt.Println("total " + strconv.Itoa(index))
	}

	fmt.Println("======= SELESAI Products V3 TOKPED =======")

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsLoopV3(rows int, starts int, index int) {
	nmRoute := "GetProductsLoopV3"
	var response response.ResponseCrud
	fmt.Println("masuk loop")
	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	shopid := os.Getenv("SHOP_ID_TOKPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/info?shop_id=" + shopid
	urltokped += "&page=" + strconv.Itoa(starts)
	urltokped += "&per_page=" + strconv.Itoa(rows)

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		RespHead := resp.Header
		datas := parseJsonProductTokped3(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		// fmt.Println(datas.Data.TotalData)
		// fmt.Println(len(datas.Data.Products))
		psnError := ""
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		if len(datas.Data) < 1 {
			fmt.Println(string(data))

		}
		for _, val := range datas.Data {
			index++
			fmt.Println("ID Parent loop" + strconv.Itoa(int(val.Basic.ProductID)))
			fmt.Println("sku parent loop" + val.Other.Sku)
			if len(val.Variant.ChildrenID) < 1 {
				fmt.Println("save single mapping sku loop")
				//save single mapping sku
			}
			for _, valVariant := range val.Variant.ChildrenID {
				GetProductDetailLoop(strconv.Itoa(int(valVariant)))
			}
			//GetProductDetailLoop(strconv.Itoa(int(val.Basic.ProductID)))
		}

		if len(datas.Data) > 0 {
			starts++

			if RespHead["X-Ratelimit-Remaining"][0] == "0" {
				fmt.Println("masuk sleep")
				time.Sleep(30 * time.Second)
			}
			fmt.Println("total index " + strconv.Itoa(index))
			//GetProductsLoopV3(rows, starts, index)
		}

	}

}

func UpdateStock(c *gin.Context) {
	nmRoute := "UpdateStock"
	ProductId := c.Param("id")
	Stock := c.Param("stock")
	Sku := c.Param("sku")
	var response response.ResponseCrud
	var objRest models.TableLogStock

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	shopid := os.Getenv("SHOP_ID_TOKPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/stock/update?shop_id=" + shopid
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var jsonString = `[{"product_id":` + ProductId + `,"new_stock":` + Stock + `}]`
	var body_url = []byte(jsonString)

	objRest.UuidLog = uuid.New().String()
	objRest.ChannelCode = os.Getenv("KODE_TOKPED")
	objRest.Sku = Sku
	objRest.Body = jsonString
	if s, err := strconv.ParseFloat(Stock, 64); err == nil {
		objRest.Stock = s
	}

	req, err := http.NewRequest("POST", urltokped, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)
		objRest.Response = nmRoute + " Gagal Koneksi"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonStockTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas
		objRest.Response = string(data)
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		response.Message = psnError

	}

	objRest.CreatedBy = "API"
	objRest.CreatedDate = time.Now()
	tokenRepository.SaveStockAPI(objRest)

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStockDecre(c *gin.Context) {
	nmRoute := "UpdateStockDecre"
	ProductId := c.Param("id")
	Stock := c.Param("stock")
	Sku := c.Param("sku")
	StockOld := c.Param("stockold")
	var response response.ResponseCrud
	var objRest models.TableLogStock

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	shopid := os.Getenv("SHOP_ID_TOKPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/stock/update?shop_id=" + shopid
	token := fmt.Sprintf("%v", ObjToken.Value1)

	//var jsonString = `[{"product_id":` + ProductId + `,"new_stock":` + Stock + `}]`
	var jsonString = `
	{
		"fs_id": ` + fsid + `,
		"product_changes_data": [
		  {
			"action": "stock_decrement",
			"product_id": ` + ProductId + `,
			"shop_id": ` + shopid + `,
			"warehouse_id": 683587,
			"value": "` + Stock + `",
			"previous_value": "` + StockOld + `",
			"is_default_warehouse": true
		  }
		]
	  }
	`
	var body_url = []byte(jsonString)

	objRest.UuidLog = uuid.New().String()
	objRest.ChannelCode = os.Getenv("KODE_TOKPED")
	objRest.Sku = Sku
	objRest.Body = jsonString
	if s, err := strconv.ParseFloat(Stock, 64); err == nil {
		objRest.Stock = s
	}

	req, err := http.NewRequest("POST", urltokped, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)
		objRest.Response = nmRoute + " Gagal Koneksi"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonStockTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas
		objRest.Response = string(data)
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		response.Message = psnError

	}

	objRest.CreatedBy = "API"
	objRest.CreatedDate = time.Now()
	tokenRepository.SaveStockAPI(objRest)

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStockAuto0(ProductId, Stock string) {
	nmRoute := "UpdateStockAuto0"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	shopid := os.Getenv("SHOP_ID_TOKPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/stock/update?shop_id=" + shopid
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var jsonString = `[{"product_id":` + ProductId + `,"new_stock":` + Stock + `}]`
	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", urltokped, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		RespHead := resp.Header
		datas := parseJsonStockTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas
		if RespHead["X-Ratelimit-Remaining"][0] == "0" {
			fmt.Println("masuk sleep")
			time.Sleep(10 * time.Second)
		}
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		response.Message = psnError

	}

}

func GetProductsBySku(c *gin.Context) {
	nmRoute := "GetProductsBySku"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	//sku := "a08173598"
	sku := c.Query("sku")

	//ambil buat jadi 0 dulu
	var ObjSkuMapping []models.TableSkuMapping
	kode_tokped := os.Getenv("KODE_TOKPED")

	ObjSkuMapping, _ = stockRepository.CariSkuParentbychannel(sku, kode_tokped)
	if len(ObjSkuMapping) > 0 {
		for _, ObjSkuMapping := range ObjSkuMapping {
			IdParent := ObjSkuMapping.IdSkuParent
			UpdateStockAuto0(IdParent, "0")
		}
	}

	sku = "a" + sku
	token := fmt.Sprintf("%v", ObjToken.Value1)
	fmt.Println("======= GET Products By SKU TOKPED =======")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/info?sku=" + sku

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		RespHead := resp.Header
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailTokped(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		//index := 0

		psnError := ""
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			if psnError == "Data Not Found" {
				GetProductsBySku2(c.Query("sku"))
			} else {
				tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
			}

		}

		if len(RespHead["X-Ratelimit-Remaining"]) > 0 {
			if RespHead["X-Ratelimit-Remaining"][0] == "0" {
				fmt.Println("masuk sleep " + nmRoute)
				time.Sleep(3 * time.Second)
			}
		}

		if len(datas.Data) > 0 {

			//save single
			tokenService.SaveSkuMappingSingleTokpedNew(datas, c.Query("sku"))

		}

	}

	fmt.Println("======= SELESAI Products By SKU TOKPED =======")

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsBySkuALL(c *gin.Context) {
	nmRoute := "GetProductsBySkuALL"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	//sku := "a08173598"
	//sku := c.Query("sku")

	//ambil buat jadi 0 dulu
	var ObjSkuMapping []models.TableSkuMapping
	kode_tokped := os.Getenv("KODE_TOKPED")

	datas, _ := tokenRepository.CariSkuWMSStock()

	fmt.Println("======= " + utils.DateToStdNow() + " MULAI " + nmRoute + " By SKU TOKPED =======")

	for _, val := range datas {

		sku := val.SkuNo

		ObjSkuMapping, _ = stockRepository.CariSkuParentbychannel(sku, kode_tokped)
		if len(ObjSkuMapping) > 0 {
			for _, ObjSkuMapping := range ObjSkuMapping {
				IdParent := ObjSkuMapping.IdSkuParent
				UpdateStockAuto0(IdParent, "0")
			}
		}

		sku = "a" + sku
		token := fmt.Sprintf("%v", ObjToken.Value1)
		fmt.Println("======= " + nmRoute + "By SKU " + sku + " TOKPED =======")

		urltokped := os.Getenv("URL_API_TOKPED")
		urltokped += "/inventory/v1/fs/" + fsid + "/product/info?sku=" + sku

		req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
		req.Header.Set("Authorization", "Bearer "+token)
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
			response.Result = err
			tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

		} else {
			defer resp.Body.Close()
			RespHead := resp.Header
			data, _ := ioutil.ReadAll(resp.Body)
			datas := parseJsonProductDetailTokped(data)

			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			response.Message = ""
			response.Result = datas

			//index := 0

			psnError := ""
			if datas.Header.Reason != "" {
				psnError = datas.Header.Reason
				if psnError == "Data Not Found" {
					GetProductsBySku2(c.Query("sku"))
				} else {
					tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
				}

			}

			if len(RespHead["X-Ratelimit-Remaining"]) > 0 {
				if RespHead["X-Ratelimit-Remaining"][0] == "0" {
					fmt.Println("masuk sleep " + nmRoute)
					time.Sleep(3 * time.Second)
				}
			}

			if len(datas.Data) > 0 {

				//save single
				tokenService.SaveSkuMappingSingleTokpedNew(datas, c.Query("sku"))

			}

			time.Sleep(1 * time.Second)

		}

		fmt.Println("======= SELESAI " + nmRoute + " By SKU " + sku + " TOKPED =======")

	}

	fmt.Println("======= " + utils.DateToStdNow() + " SELESAI " + nmRoute + " By SKU TOKPED =======")

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsByUrl(c *gin.Context) {
	nmRoute := "GetProductsByUrl"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	//sku := "a08173598"
	sku := c.Query("url")

	//ambil buat jadi 0 dulu
	// var ObjSkuMapping []models.TableSkuMapping
	// kode_tokped := os.Getenv("KODE_TOKPED")

	// ObjSkuMapping, _ = stockRepository.CariSkuParentbychannel(sku, kode_tokped)
	// if len(ObjSkuMapping) > 0 {
	// 	for _, ObjSkuMapping := range ObjSkuMapping {
	// 		IdParent := ObjSkuMapping.IdSkuParent
	// 		UpdateStockAuto0(IdParent, "0")
	// 	}
	// }

	token := fmt.Sprintf("%v", ObjToken.Value1)
	fmt.Println("======= GET Products By URL TOKPED =======")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/info?product_url=" + sku

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		RespHead := resp.Header
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailTokped(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		//index := 0

		psnError := ""
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)

		}

		if len(RespHead["X-Ratelimit-Remaining"]) > 0 {
			if RespHead["X-Ratelimit-Remaining"][0] == "0" {
				fmt.Println("masuk sleep " + nmRoute)
				time.Sleep(3 * time.Second)
			}
		}

		if len(datas.Data) > 0 {

			//save single
			skunya := datas.Data[0].Other.Sku
			if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else {

				if len([]rune(skunya)) < 8 {
					skunya = "0" + skunya
				}
			}

			tokenService.SaveSkuMappingSingleTokpedNew(datas, skunya)

		}

	}

	fmt.Println("======= SELESAI Products By URL TOKPED =======")

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsBySku2(sku string) {
	nmRoute := "GetProductsBySku2"

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	token := fmt.Sprintf("%v", ObjToken.Value1)

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/info?sku=" + sku

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		RespHead := resp.Header
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailTokped(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		//index := 0

		psnError := ""
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			if psnError == "Data Not Found" {

			} else {
				tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
			}

		}

		if len(RespHead["X-Ratelimit-Remaining"]) > 0 {
			if RespHead["X-Ratelimit-Remaining"][0] == "0" {
				fmt.Println("masuk sleep " + nmRoute)
				time.Sleep(3 * time.Second)
			}
		}

		if len(datas.Data) > 0 {

			//save single
			tokenService.SaveSkuMappingSingleTokpedNew(datas, sku)

		}

	}

}

func GetProductsBySkuAuto(skunya string) {
	nmRoute := "GetProductsBySkuAuto"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	sku := skunya

	//ambil buat jadi 0 dulu
	// var ObjSkuMapping []models.TableSkuMapping
	// kode_tokped := os.Getenv("KODE_TOKPED")

	// ObjSkuMapping, _ = stockRepository.CariSkuParentbychannel(sku, kode_tokped)
	// if len(ObjSkuMapping) > 0 {
	// 	for _, ObjSkuMapping := range ObjSkuMapping {
	// 		IdParent := ObjSkuMapping.IdSkuParent
	// 		UpdateStockAuto0(IdParent, "0")
	// 	}
	// }

	sku = "a" + sku

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println("======= GET Products By SKU TOKPED =======")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/info?sku=" + sku

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		RespHead := resp.Header
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailTokped(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		//index := 0

		psnError := ""
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			if psnError == "Data Not Found" {
				GetProductsBySku2(skunya)
			} else {
				tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
			}

		}

		if len(RespHead["X-Ratelimit-Remaining"]) > 0 {
			if RespHead["X-Ratelimit-Remaining"][0] == "0" {
				fmt.Println("masuk sleep " + nmRoute)
				time.Sleep(3 * time.Second)
			}
		}

		if len(datas.Data) > 0 {

			//save single
			tokenService.SaveSkuMappingSingleTokpedNew(datas, skunya)
			//helpers.UpdateStock(skunya, "API_CHANNEL", kode_tokped)
		}

	}

	//fmt.Println("======= SELESAI Products By SKU TOKPED =======")

}
func GetAllSkuWms(c *gin.Context) {
	var response response.ResponseCrud
	fmt.Println("======= Mulai Products By SKU TOKPED =======")
	datas, _ := tokenRepository.CariSkuWMSAll()
	tulis := ""
	for _, val := range datas {
		tulis = val.SkuNo
		GetProductsBySkuAuto(tulis)
	}
	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = ""
	//response.Result = tulis
	fmt.Println("======= SELESAI Products By SKU TOKPED =======")
	c.JSON(http.StatusOK, response)
	return
}

func GetProductsAutoV2(wg *sync.WaitGroup) {
	fmt.Println("mulai product tokped " + time.Now().String())
	tokenService.SaveErrorString(ChannelName, "mulai product tokped", "GetProductsAutoV2")

	// datas, _ := tokenRepository.CariSkuWMSAll()
	datas, _ := tokenRepository.CariSkuWMSStock()

	tulis := ""
	for _, val := range datas {
		tulis = val.SkuNo
		var wgx sync.WaitGroup
		wgx.Add(1)
		// go func() {
		// 	GetProductsBySkuAutoV2(tulis, &wgx)
		// }()
		go func() {
			GetProductsBySku2V2(tulis, &wgx)
		}()
		wgx.Wait()
	}
	fmt.Println("selesai product tokped " + time.Now().String())
	tokenService.SaveErrorString(ChannelName, "selesai product tokped", "GetProductsAutoV2")
	wg.Done()
}

func GetProductsBySkuAutoV2(skunya string, wgx *sync.WaitGroup) {
	nmRoute := "GetProductsBySkuAutoV2"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")
	sku := skunya

	sku = "a" + sku

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println("======= GET Products By SKU TOKPED =======")

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/info?sku=" + sku

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		RespHead := resp.Header
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailTokped(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		//index := 0

		psnError := ""
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			if psnError == "Data Not Found" {
				//GetProductsBySku2(skunya)
			} else {
				tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
			}

		}

		if len(RespHead["X-Ratelimit-Remaining"]) > 0 {
			if RespHead["X-Ratelimit-Remaining"][0] == "0" {
				fmt.Println("masuk sleep " + nmRoute)
				time.Sleep(3 * time.Second)
			}
		}

		if len(datas.Data) > 0 {

			//save single
			tokenService.SaveSkuMappingSingleTokpedNew(datas, skunya)
		}

	}

	wgx.Done()
}
func GetProductsBySku2V2(sku string, wgx *sync.WaitGroup) {
	nmRoute := "GetProductsBySku2V2"

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	token := fmt.Sprintf("%v", ObjToken.Value1)

	urltokped := os.Getenv("URL_API_TOKPED")
	urltokped += "/inventory/v1/fs/" + fsid + "/product/info?sku=" + sku

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		RespHead := resp.Header
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailTokped(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		//index := 0

		psnError := ""
		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			if psnError == "Data Not Found" {

			} else {
				tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
			}

		}

		if len(RespHead["X-Ratelimit-Remaining"]) > 0 {
			if RespHead["X-Ratelimit-Remaining"][0] == "0" {
				//fmt.Println("masuk sleep " + nmRoute)
				time.Sleep(3 * time.Second)
			}
		}

		if len(datas.Data) > 0 {

			//save single
			tokenService.SaveSkuMappingSingleTokpedNew(datas, sku)

		}

	}
	wgx.Done()
}

func GetOrders(c *gin.Context) {
	nmRoute := "GetOrders"
	Status := c.Param("status")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fmt.Println(ObjToken)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	page := 1
	per_page := 50

	//skrg := strconv.FormatInt(time.Now().Unix(), 10)
	skrg := strconv.FormatInt(time.Now().Add(time.Duration(7)*time.Hour).Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-65)*time.Hour).Unix(), 10) //3 hari lalu

	//kmrn = "1616904016"
	//skrg = "1617163216"

	urltokped += "/v2/order/list?fs_id=" + fsid
	urltokped += "&page=" + strconv.Itoa(page)
	urltokped += "&per_page=" + strconv.Itoa(per_page)
	urltokped += "&from_date=" + kmrn
	urltokped += "&to_date=" + skrg
	urltokped += "&status=" + Status
	fmt.Println(Status)

	token := fmt.Sprintf("%v", ObjToken.Value1)

	fmt.Println(urltokped)
	fmt.Println(token)
	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas
		// response.Result = string(data)

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		// if len(datas.Data) > 0 {
		// 	for _, val := range datas.Data {
		// 		// 	GetOrderDetailLoop(strconv.Itoa(val.OrderID), val.Logistics.ShippingAgency)
		// 		fmt.Println(val)
		// 	}

		// }

		response.Message = psnError
		response.Message = urltokped
	}

	c.JSON(http.StatusOK, response)
	return
}
func GetOrdersAuto(Status string) {
	nmRoute := "GetOrdersAuto"

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")
	page := 1
	per_page := 50

	// skrg := strconv.FormatInt(time.Now().Unix(), 10)
	// kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari lalu

	skrg := strconv.FormatInt(time.Now().Add(time.Duration(7)*time.Hour).Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-65)*time.Hour).Unix(), 10) //3 hari lalu

	//kmrn = "1616904016"
	//skrg = "1617163216"

	urltokped += "/v2/order/list?fs_id=" + fsid
	urltokped += "&page=" + strconv.Itoa(page)
	urltokped += "&per_page=" + strconv.Itoa(per_page)
	urltokped += "&from_date=" + kmrn
	urltokped += "&to_date=" + skrg
	urltokped += "&status=" + Status
	//fmt.Println(Status)

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		RespHead := resp.Header
		datas := parseJsonOrdersTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		if len(RespHead["X-Ratelimit-Remaining"]) > 0 {
			if RespHead["X-Ratelimit-Remaining"][0] == "0" {
				fmt.Println("masuk sleep " + nmRoute)
				time.Sleep(10 * time.Second)
			}
		}

		if len(datas.Data) > 0 {

			for _, val := range datas.Data {
				GetOrderDetailLoop(strconv.Itoa(val.OrderID), val.Logistics.ShippingAgency)
			}

			page++
			GetOrdersAutoLoop(page, per_page, Status)
		}

		response.Message = psnError
		response.Message = urltokped
	}

	fmt.Println("GetOrdersAuto TOKOPEDIA " + Status + " SUKSES")

}

func GetOrdersAutoLoop(page int, per_page int, Status string) {
	nmRoute := "GetOrdersAutoLoop"

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	// skrg := strconv.FormatInt(time.Now().Unix(), 10)
	// kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari lalu

	skrg := strconv.FormatInt(time.Now().Add(time.Duration(7)*time.Hour).Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-65)*time.Hour).Unix(), 10) //3 hari lalu

	urltokped += "/v2/order/list?fs_id=" + fsid
	urltokped += "&page=" + strconv.Itoa(page)
	urltokped += "&per_page=" + strconv.Itoa(per_page)
	urltokped += "&from_date=" + kmrn
	urltokped += "&to_date=" + skrg
	urltokped += "&status=" + Status

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		RespHead := resp.Header
		datas := parseJsonOrdersTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		if RespHead["X-Ratelimit-Remaining"][0] == "0" {
			fmt.Println("masuk sleep " + nmRoute)
			time.Sleep(10 * time.Second)
		}

		if len(datas.Data) > 0 {

			for _, val := range datas.Data {
				GetOrderDetailLoop(strconv.Itoa(val.OrderID), val.Logistics.ShippingAgency)
			}

			page++
			GetOrdersAutoLoop(page, per_page, Status)

		}

		response.Message = psnError
		response.Message = urltokped
	}

}

func GetOrderDetail(c *gin.Context) {
	nmRoute := "GetOrderDetail"
	OrderId := c.Param("orderid")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	urltokped += "/v2/fs/" + fsid + "/order"
	urltokped += "?order_id=" + OrderId

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrderDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas
		//response.Result = string(data)

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		} else {
			response.ResponseDesc = strconv.Itoa(datas.Data.OrderStatus)
			// response.ResponseDesc = datas
		}

		response.Message = psnError

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrderDetailByInvoice(c *gin.Context) {
	nmRoute := "GetOrderDetail"
	InvoiceNum := c.Query("invoice_num")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	urltokped += "/v2/fs/" + fsid + "/order"
	urltokped += "?invoice_num=" + InvoiceNum

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrderDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas
		//response.Result = string(data)

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		} else {
			response.ResponseDesc = strconv.Itoa(datas.Data.OrderStatus)
			// response.ResponseDesc = datas
		}

		response.Message = psnError

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrderDetailLoopNew(OrderId string, Kurir string) {
	nmRoute := "GetOrderDetailLoop"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	urltokped += "/v2/fs/" + fsid + "/order"
	urltokped += "?order_id=" + OrderId

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		RespHead := resp.Header
		datas := parseJsonOrderDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		if datas.Data.OrderID != 0 {
			tokenService.SaveSalesOrderTokped(datas, Kurir)
		}
		if RespHead["X-Ratelimit-Remaining"][0] == "0" {
			fmt.Println("masuk sleep " + nmRoute)
			time.Sleep(10 * time.Second)
		}

		response.Message = psnError

	}

}

func GetOrderDetailLoop(OrderId string, Kurir string) {
	nmRoute := "GetOrderDetailLoop"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	urltokped += "/v2/fs/" + fsid + "/order"
	urltokped += "?order_id=" + OrderId

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		RespHead := resp.Header
		datas := parseJsonOrderDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		if datas.Data.OrderID != 0 {
			tokenService.SaveSalesOrderTokped(datas, Kurir)
		}
		if RespHead["X-Ratelimit-Remaining"][0] == "0" {
			fmt.Println("masuk sleep " + nmRoute)
			time.Sleep(10 * time.Second)
		}

		response.Message = psnError

	}

}

func GetOrderidByInvoice(c *gin.Context) {
	nmRoute := "GetOrderidByInvoice"
	OrderId := c.Query("orderid")

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	urltokped += "/v2/fs/" + fsid + "/order"
	urltokped += "?invoice_num=" + OrderId

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrderDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		} else {
			response.ResponseDesc = strconv.Itoa(datas.Data.OrderStatus)
			if datas.Data.OrderID != 0 {
				Kurir := datas.Data.OrderInfo.ShippingInfo.LogisticName
				//fmt.Println("kurir " + Kurir)
				tokenService.SaveSalesOrderTokped(datas, Kurir)
			}

		}

		response.Message = psnError

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrderidByInvoiceAuto(OrderId string) string {
	nmRoute := "GetOrderidByInvoiceAuto"

	OrderIdnya := ""
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	urltokped += "/v2/fs/" + fsid + "/order"
	urltokped += "?invoice_num=" + OrderId

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrderDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		} else {
			response.ResponseDesc = strconv.Itoa(datas.Data.OrderStatus)
			OrderIdnya = strconv.Itoa(datas.Data.OrderID)
		}

		response.Message = psnError

	}

	return OrderIdnya
}

func GetLabelShipping(c *gin.Context) {
	nmRoute := "GetLabelShipping"
	OrderId := c.Param("orderid")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	urltokped += "/v1/order/" + OrderId + "/fs/" + fsid + "/shipping-label"

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrderDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = string(data)

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		response.Message = psnError

	}

	c.JSON(http.StatusOK, response)
	return
}

func AcceptOrder(c *gin.Context) {
	nmRoute := "AcceptOrder"
	OrderId := c.Param("orderid")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	urltokped += "/v1/order/" + OrderId + "/fs/" + fsid + "/ack"

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("POST", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrderDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = string(data)

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		response.Message = psnError

	}

	c.JSON(http.StatusOK, response)
	return
}

func ReqShip(c *gin.Context) {
	nmRoute := "ReqShip"
	OrderId := c.Param("orderid")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	urltokped += "/inventory/v1/fs/" + fsid + "/pick-up"

	token := fmt.Sprintf("%v", ObjToken.Value1)
	shopid := os.Getenv("SHOP_ID_TOKPED")

	var bodyString = `{
		"order_id": ` + OrderId + `,
		"shop_id": ` + shopid + `
		}'`
	var body_url = []byte(bodyString)

	req, err := http.NewRequest("POST", urltokped, bytes.NewBuffer(body_url))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrderDetailTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = string(data)

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		response.Message = psnError

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOnlineBooking(c *gin.Context) {
	nmRoute := "GetOnlineBooking"
	OrderId := c.Param("orderid")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	urltokped += "/v1/fs/" + fsid + "/fulfillment_order?order_id=" + OrderId

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.Result = err
		response.Message = "ERROR CONNECTION "
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonDetailBookingTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		response.Result = datas

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		}

		response.Message = psnError

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetResi(c *gin.Context) {
	nmRoute := "GetResi"
	OrderId := c.Param("orderid")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)

	fsid := os.Getenv("ID_TOKOPED")

	urltokped := os.Getenv("URL_API_TOKPED")

	// urltokped += "/v2/fs/" + fsid + "/order"
	// urltokped += "?order_id=" + OrderId
	urltokped += "/v1/fs/" + fsid + "/fulfillment_order?order_id=" + OrderId

	token := fmt.Sprintf("%v", ObjToken.Value1)

	req, err := http.NewRequest("GET", urltokped, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+token)
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
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		response.Message = "GAGAL GET RESI TOKOPEDIA"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJsonOrderDetailTokped(data)
		datas := parseJsonDetailBookingTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		psnError := ""

		// response.Result = string(datas)
		response.Result = datas

		if datas.Header.Reason != "" {
			psnError = datas.Header.Reason
			tokenService.SaveErrorString(ChannelName, psnError, nmRoute)
		} else {
			if len(datas.Data.OrderData) > 0 {
				response.ResponseDesc = datas.Data.OrderData[0].BookingData.BookingCode
			}

			//response.ResponseDesc = datas.Data.OrderInfo.ShippingInfo.Awb

		}

		response.Message = psnError

	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateAllStockTokped(c *gin.Context) {
	var response response.ResponseCrud

	ObjMapping := tokenService.CariSkuMappingObjGroup(os.Getenv("KODE_TOKPED")) //param by channel name

	if len(ObjMapping) > 0 {
		for _, value := range ObjMapping {
			helpers.UpdateStock(value.SkuNo, "API_CHANNEL", os.Getenv("KODE_TOKPED"))

		}
	}

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "DISINISS"

	c.JSON(http.StatusOK, response)
	return
}
