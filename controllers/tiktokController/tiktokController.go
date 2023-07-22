package tiktokController

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	"github.com/rals/dearme-channel/services/tokenService"

	"github.com/rals/dearme-channel/helpers"
	"github.com/rals/dearme-channel/utils"
	// "encoding/json"
)

var ChannelName = "tiktok"

func parseAuthTiktok(jsonBuffer []byte) models.AuhtTiktok {

	AuhtTiktok := models.AuhtTiktok{}

	err := json.Unmarshal(jsonBuffer, &AuhtTiktok)
	if err != nil {
		return AuhtTiktok
	}

	// the array is now filled with users
	return AuhtTiktok

}

func Callback(c *gin.Context) {
	nmRoute := "Callback Tiktok"
	code := c.Query("code")

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	partner_secret := os.Getenv("PARTNER_SECRET_TIKTOK")

	var response response.ResponseCrud

	url := os.Getenv("URL_AUTH_TIKTOK") + "/api/token/getAccessToken"
	//fmt.Println(url)
	var jsonString = `{"app_key":"` + partner_key + `",
	"app_secret":"` + partner_secret + `",
	"auth_code":"` + code + `",
	"grant_type":"authorized_code"}`

	fmt.Println(jsonString)

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		fmt.Println(string(data))
		datas := parseAuthTiktok(data)
		response.Result = datas

		var objToken models.TokenBukalapak

		if datas.Data.AccessToken != "" {
			objToken.AccessToken = datas.Data.AccessToken
			objToken.RefreshToken = datas.Data.RefreshToken

			tokenService.SaveTokenBukalapak(objToken, ChannelName)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			response.Message = "Save Token Sukses"
			response.Result = datas
		}

		// fmt.Println("access_tokennya :" + objToken.AccessToken)
		// fmt.Println("refresh_token :" + objToken.RefreshToken)

		//insert token

	}

	c.JSON(http.StatusOK, response)
	return
}

func RefreshTokenTiktok(c *gin.Context) {
	nmRoute := "RefreshTokenTiktok Tiktok"

	ObjToken := tokenService.FindToken("tiktok")
	RefreshToken := fmt.Sprintf("%v", ObjToken.Value2)

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	partner_secret := os.Getenv("PARTNER_SECRET_TIKTOK")

	var response response.ResponseCrud

	url := os.Getenv("URL_AUTH_TIKTOK") + "/api/token/refreshToken"
	//fmt.Println(url)
	var jsonString = `{"app_key":"` + partner_key + `",
	"app_secret":"` + partner_secret + `",
	"refresh_token":"` + RefreshToken + `",
	"grant_type":"refresh_token"}`

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		fmt.Println(string(data))
		datas := parseAuthTiktok(data)
		response.Result = datas

		var objToken models.TokenBukalapak

		if datas.Data.AccessToken != "" {
			objToken.AccessToken = datas.Data.AccessToken
			objToken.RefreshToken = datas.Data.RefreshToken

			tokenService.SaveTokenBukalapak(objToken, ChannelName)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			response.Message = "Refresh Token Sukses"
			response.Result = datas
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetAuthorizeShopTiktok(c *gin.Context) {
	nmRoute := "GetAuthorizeShopTiktok Tiktok"
	//partner_secret := os.Getenv("PARTNER_SECRET_TIKTOK")
	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/shop/get_authorized_shop"

	tambahan += ""
	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseOrdersTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		// if datas.Message != "Success" { //save error
		// 	tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
		// } else {

		// }

		response.Result = string(data)

		response.Message = ""

	}
	//fmt.Println(concatenated)

	c.JSON(http.StatusOK, response)
	return
}

func UrlLoginTiktok(c *gin.Context) {
	nmRoute := "AuthTiktok Bukalapak"
	//partner_secret := os.Getenv("PARTNER_SECRET_TIKTOK")
	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	partner_state := os.Getenv("PARTNER_STATE_TIKTOK")

	var response response.ResponseCrud

	url := os.Getenv("URL_AUTH_TIKTOK") + "/oauth/authorize?app_key=" + partner_key + "&state=" + partner_state

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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		// response.Message = url
		response.Result = string(data)
		signnya, baseString, _ := AuthTiktok("/api/products/categories", "tokennya", "category_id#602362#shop_id#7493990568670889543")
		//fmt.Println(signnya)
		response.Result = signnya
		response.Message = baseString
		response.Message = url
		//insert token

	}

	c.JSON(http.StatusOK, response)
	return
}

func AuthTiktok(path string, token string, option string) (string, string, string) {
	now := time.Now()
	timest := strconv.FormatInt(now.Unix(), 10)

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	partner_secret := os.Getenv("PARTNER_SECRET_TIKTOK")

	parameter := map[string]string{
		"app_key":   partner_key,
		"timestamp": timest,
	}

	paramTambah := strings.Split(option, "#")
	index := 2
	keyisi := "kosong"
	for _, value := range paramTambah {

		if (index % 2) == 0 {
			keyisi = value
		} else {

			parameter[keyisi] = value
		}
		index++

	}

	keys := make([]string, len(parameter))
	idx := 0
	for k, _ := range parameter {
		keys[idx] = k
		idx++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	input := path
	for _, key := range keys {
		input = input + key + parameter[key]
	}

	input = partner_secret + input + partner_secret

	h := hmac.New(sha256.New, []byte(partner_secret))

	balik := ""
	concatenated := ""

	if _, err := h.Write([]byte(input)); err != nil {
		return balik, timest, concatenated
	} else {
		balik = hex.EncodeToString(h.Sum(nil))
	}

	return balik, timest, concatenated

}

func parseOrdersTiktok(jsonBuffer []byte) models.RespOrdersTiktok {

	RespOrdersTiktok := models.RespOrdersTiktok{}

	err := json.Unmarshal(jsonBuffer, &RespOrdersTiktok)
	if err != nil {
		return RespOrdersTiktok
	}

	// the array is now filled with users
	return RespOrdersTiktok

}

func GetProcessedOrderAutoTiktok(status string) {
	urlApi := os.Getenv("URL_WMS_TIKTOK") + "getOrders/" + status

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
		fmt.Println("gagal GetProcessedOrderAutoTiktok")
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

	} else {
		fmt.Println("sukses GetProcessedOrderAutoTiktok")
		defer resp.Body.Close()
	}

}

func GetOrdersTiktok(c *gin.Context) {
	nmRoute := "GetOrdersTiktok Tiktok"
	//partner_secret := os.Getenv("PARTNER_SECRET_TIKTOK")
	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")
	pagesize := os.Getenv("PAGE_SIZE_ORDER_TIKTOK")
	Status := c.Param("status")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/orders/search"

	tambahan += "shop_id#" + shop_id
	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari yg lalu
	var jsonString = `{
		"create_time_from":` + kmrn + `,
		"create_time_to":` + skrg + `,
		"update_time_from":` + kmrn + `,
		"update_time_to":` + skrg + `,
		"sort_type":1,
		"sort_by":"CREATE_TIME",
		"page_size":` + pagesize + `,
		"order_status":` + Status + `}`

	//fmt.Println(jsonString)

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseOrdersTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
		} else {

			index := 0
			orderidnya := ""
			if len(datas.Data.OrderList) > 0 {

				for _, val := range datas.Data.OrderList {

					if Status == "140" { //CANCEL
						objCekCancel, _ := tokenRepository.FindSalesOrder(val.OrderId)
						if objCekCancel.StatusProcessOrder != "3" && objCekCancel.StatusProcessOrder != "" {
							GetOrderDetailTiktokAuto(val.OrderId)
						}
					} else if Status == "111" { //PENDING
						//fmt.Println("GetOrdersTiktok " + strconv.Itoa(index) + " | " + val.OrderId)

						objCekData, _ := tokenRepository.FindSalesOrder(val.OrderId)
						if objCekData.StatusProcessOrder == "" {
							fmt.Println("GetOrdersTiktok " + strconv.Itoa(index) + " | " + val.OrderId)
							GetOrderDetailTiktokAuto(val.OrderId)
						}

						//GetOrderDetailTiktokAuto(val.OrderId)
					}
					//GetOrderDetailTiktokAuto(val.OrderId)

					if index != 0 {
						orderidnya += ","
					}
					orderidnya += `"` + val.OrderId + `"`

					index++
				}

				// GetOrderDetailTiktokAuto(orderidnya)

				if index < datas.Data.Total {
					fmt.Println("masuk looping")
					GetOrdersLoopTiktok(datas.Data.NextCursor, index, Status)
				}

			}

		}

		response.Result = datas

		response.Message = ""

	}
	//fmt.Println(concatenated)

	c.JSON(http.StatusOK, response)
	return
}

func GetOrdersLoopTiktok(cursor string, index int, status string) {
	nmRoute := "GetOrdersLoopTiktok Tiktok"
	//partner_secret := os.Getenv("PARTNER_SECRET_TIKTOK")
	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")
	pagesize := os.Getenv("PAGE_SIZE_ORDER_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/orders/search"

	tambahan += "shop_id#" + shop_id
	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari yg lalu
	var jsonString = `{
		"create_time_from":` + kmrn + `,
		"create_time_to":` + skrg + `,
		"update_time_from":` + kmrn + `,
		"update_time_to":` + skrg + `,
		"sort_type":1,
		"sort_by":"CREATE_TIME",
		"page_size":` + pagesize + `,
		"cursor":"` + cursor + `",
		"order_status":` + status + `}`

	//fmt.Println(jsonString)

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseOrdersTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
		} else {

			orderidnya := ""
			indextemp := index
			if len(datas.Data.OrderList) > 0 {

				for _, val := range datas.Data.OrderList {

					if status == "140" { //CANCEL
						objCekCancel, _ := tokenRepository.FindSalesOrder(val.OrderId)
						if objCekCancel.StatusProcessOrder != "3" && objCekCancel.StatusProcessOrder != "" {
							GetOrderDetailTiktokAuto(val.OrderId)
						}
					} else if status == "111" { //PENDING
						//fmt.Println("GetOrdersLoopTiktok " + strconv.Itoa(index) + " | " + val.OrderId)
						objCekData, _ := tokenRepository.FindSalesOrder(val.OrderId)
						if objCekData.StatusProcessOrder == "" {
							fmt.Println("GetOrdersLoopTiktok " + strconv.Itoa(index) + " | " + val.OrderId)
							GetOrderDetailTiktokAuto(val.OrderId)
						}
						//GetOrderDetailTiktokAuto(val.OrderId)
					}
					//GetOrderDetailTiktokAuto(val.OrderId)

					//fmt.Println("GetOrdersLoopTiktok " + strconv.Itoa(index) + " | " + val.OrderId)
					//GetOrderDetailTiktokAuto(val.OrderId)
					if index != indextemp {
						orderidnya += ","
					}
					orderidnya += `"` + val.OrderId + `"`

					index++
				}

				//GetOrderDetailTiktokAuto(orderidnya)
				if index < datas.Data.Total {
					fmt.Println("masuk looping")
					GetOrdersLoopTiktok(datas.Data.NextCursor, index, status)
				}

			}

		}

		// fmt.Println("======== orderss looping ========")
		// fmt.Println(string(data))
		// fmt.Println("======== orderss looping ========")

	}
	//fmt.Println(concatenated)
}

func parseOrderDetailTiktok(jsonBuffer []byte) models.RespOrderDetailTiktok {

	RespOrderDetailTiktok := models.RespOrderDetailTiktok{}

	err := json.Unmarshal(jsonBuffer, &RespOrderDetailTiktok)
	if err != nil {
		return RespOrderDetailTiktok
	}

	// the array is now filled with users
	return RespOrderDetailTiktok

}

func parseConfigPickupTiktok(jsonBuffer []byte) models.ConfigPickupTiktok {

	ConfigPickupTiktok := models.ConfigPickupTiktok{}

	err := json.Unmarshal(jsonBuffer, &ConfigPickupTiktok)
	if err != nil {
		return ConfigPickupTiktok
	}

	// the array is now filled with users
	return ConfigPickupTiktok

}

func parseDocumentShipTiktok(jsonBuffer []byte) models.DocumentShipTiktok {

	DocumentShipTiktok := models.DocumentShipTiktok{}

	err := json.Unmarshal(jsonBuffer, &DocumentShipTiktok)
	if err != nil {
		return DocumentShipTiktok
	}

	// the array is now filled with users
	return DocumentShipTiktok

}

func parseRespOrderDetailTiktok2(jsonBuffer []byte) models.RespOrderDetailTiktok2 {

	RespOrderDetailTiktok2 := models.RespOrderDetailTiktok2{}

	err := json.Unmarshal(jsonBuffer, &RespOrderDetailTiktok2)
	if err != nil {
		return RespOrderDetailTiktok2
	}

	// the array is now filled with users
	return RespOrderDetailTiktok2

}

func parsePackageTiktok(jsonBuffer []byte) models.PackageTiktok {

	PackageTiktok := models.PackageTiktok{}

	err := json.Unmarshal(jsonBuffer, &PackageTiktok)
	if err != nil {
		return PackageTiktok
	}

	// the array is now filled with users
	return PackageTiktok

}

func GetOrderDetailTiktok(c *gin.Context) {
	nmRoute := "GetOrderDetailTiktok Tiktok"
	orderid := c.Param("orderid")

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/orders/detail/query"

	tambahan += "shop_id#" + shop_id
	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	var jsonString = `{
		"order_id_list":["` + orderid + `"]}`

	//fmt.Println(jsonString)

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseOrderDetailTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)
		// fmt.Println("====== " + nmRoute + " ======")
		// fmt.Println(string(data))

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
			response.Message = datas.Message + "(TIKTOK)"
		} else {
			if len(datas.Data.OrderList) > 0 {
				response.ResponseDesc = strconv.Itoa(int(datas.Data.OrderList[0].OrderStatus))
			}
		}

		objDataDetail := GetDataDetailTiktok(orderid)
		fmt.Println(objDataDetail)

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetOrderDetailTiktokAuto(Orderid string) {
	nmRoute := "GetOrderDetailTiktokAuto Tiktok"
	orderid := Orderid
	fmt.Println(orderid)

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/orders/detail/query"

	tambahan += "shop_id#" + shop_id
	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	var jsonString = `{
		"order_id_list":["` + orderid + `"]}`

	//fmt.Println(jsonString)

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseOrderDetailTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		fmt.Println(datas)
		response.Message = ""

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)

		} else {
			fmt.Println(datas.Data.OrderList)
			fmt.Println("masuk SaveSalesOrderTiktok " + nmRoute)
			tokenService.SaveSalesOrderTiktok("tiktok", datas.Data.OrderList)
		}

		// indexdetail := 0
		// for _, val := range datas.Data.OrderList {
		// 	fmt.Println(" Order ID: " + val.OrderId)

		// 	for _, valSku := range datas.Data.OrderList[indexdetail].ItemList {
		// 		fmt.Println(" SKU ID : " + valSku.SkuId)
		// 		fmt.Println(" QTY : " + strconv.Itoa(valSku.Quantity))
		// 	}

		// 	indexdetail++
		// }

	}
}

func parseProductsTiktok(jsonBuffer []byte) models.RespProdutcsTiktok {

	RespProdutcsTiktok := models.RespProdutcsTiktok{}

	err := json.Unmarshal(jsonBuffer, &RespProdutcsTiktok)
	if err != nil {
		return RespProdutcsTiktok
	}

	// the array is now filled with users
	return RespProdutcsTiktok

}

func GetProductsTiktok(c *gin.Context) {
	nmRoute := "GetProductsTiktok Tiktok"

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")
	pagesize := os.Getenv("PAGE_SIZE_ORDER_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/products/search"

	tambahan += "shop_id#" + shop_id
	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	var jsonString = `{
		"page_size":` + pagesize + `,
		"page_number":1,
		"search_status":0}`

	//fmt.Println(jsonString)

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseProductsTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//fmt.Println(string(data))
		response.Result = datas
		index := 0

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)

		} else {
			fmt.Println(nmRoute + " Total:" + strconv.Itoa(int(datas.Data.Total)))
			if len(datas.Data.Products) > 0 {

				tokenService.SaveSkuMappingTiktok(datas.Data.Products)
				// for _, Isiproducts := range datas.Data.Products {
				// 	GetProductDetailTiktokloop(Isiproducts.Id)

				// for _, Isiproducts := range datas.Data.Products {
				// 	fmt.Println("===================================")
				// 	fmt.Println(index)
				// 	fmt.Println("Product ID :" + Isiproducts.Id)
				// 	//cari detail sku
				// 	for _, detailSKU := range Isiproducts.Skus {
				// 		fmt.Println("ID SKUNYA :" + detailSKU.Id)
				// 		fmt.Println("SKUNYA :" + detailSKU.SellerSku)
				// 	}
				// 	fmt.Println("===================================")
				// 	index++
				// }

				// if int64(index) < datas.Data.Total {
				// 	fmt.Println("masuk looping")
				// 	fmt.Println("index :" + strconv.Itoa(index))
				// 	fmt.Println("total :" + strconv.Itoa(int(datas.Data.Total)))
				page_number := 1
				fmt.Println(nmRoute + " " + strconv.Itoa(page_number))
				page_number++
				GetProductsTiktokloop(strconv.Itoa(page_number), index)

			}

		}

		response.Message = ""

		fmt.Println(nmRoute + " finish")

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetProductsTiktokAuto(wg *sync.WaitGroup) {
	nmRoute := "GetProductsTiktokAuto Tiktok"
	fmt.Println("mulai product tiktok " + time.Now().String())
	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")
	pagesize := os.Getenv("PAGE_SIZE_ORDER_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/products/search"

	tambahan += "shop_id#" + shop_id
	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari yg lalu

	var jsonString = `{
		"page_size":` + pagesize + `,
		"page_number":1,
		"search_status":0,
		"update_time_from":` + kmrn + `,
		"update_time_to":` + skrg + `
		}`

	//fmt.Println(jsonString)

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseProductsTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//fmt.Println(string(data))
		response.Result = datas
		index := 0

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)

		} else {
			//fmt.Println(nmRoute + " Total:" + strconv.Itoa(int(datas.Data.Total)))
			if len(datas.Data.Products) > 0 {

				tokenService.SaveSkuMappingTiktok(datas.Data.Products)

				page_number := 1
				//fmt.Println(nmRoute + " " + strconv.Itoa(page_number))
				page_number++
				GetProductsTiktokloop(strconv.Itoa(page_number), index)

			}

		}

		response.Message = ""

		//fmt.Println(nmRoute + " finish")

	}
	fmt.Println("selesai product tiktok " + time.Now().String())
	wg.Done()

}

func GetProductsTiktokloop(page_number string, index int) {
	nmRoute := "GetProductsTiktokloop Tiktok"

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")
	pagesize := os.Getenv("PAGE_SIZE_ORDER_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/products/search"

	tambahan += "shop_id#" + shop_id
	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	skrg := strconv.FormatInt(time.Now().Unix(), 10)
	kmrn := strconv.FormatInt(time.Now().Add(time.Duration(-72)*time.Hour).Unix(), 10) //3 hari yg lalu

	var jsonString = `{
		"page_size":` + pagesize + `,
		"page_number":` + page_number + `,
		"search_status":0,
		"update_time_from":` + kmrn + `,
		"update_time_to":` + skrg + `}`

	//fmt.Println(jsonString)

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseProductsTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//fmt.Println(string(data))
		response.Result = datas

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)

		} else {
			//fmt.Println(len(datas.Data.Products))
			if len(datas.Data.Products) > 0 {

				tokenService.SaveSkuMappingTiktok(datas.Data.Products)

				// for _, Isiproducts := range datas.Data.Products {
				// 	GetProductDetailTiktokloop(Isiproducts.Id)

				// 	fmt.Println("===================================")
				// 	fmt.Println(index)
				// 	fmt.Println("Product ID :" + Isiproducts.Id)
				// 	//cari detail sku
				// 	for _, detailSKU := range Isiproducts.Skus {
				// 		fmt.Println("ID SKUNYA :" + detailSKU.Id)
				// 		fmt.Println("SKUNYA :" + detailSKU.SellerSku)
				// 	}
				// 	fmt.Println("===================================")

				// 	index++
				// }

				page_number_next, _ := strconv.Atoi(page_number)
				//fmt.Println(nmRoute + " " + strconv.Itoa(page_number_next))
				page_number_next++
				GetProductsTiktokloop(strconv.Itoa(page_number_next), index)
			}

			// if int64(index) < datas.Data.Total {
			// 	fmt.Println("masuk looping lagi")
			// 	fmt.Println("index :" + strconv.Itoa(index))
			// 	fmt.Println("total :" + strconv.Itoa(int(datas.Data.Total)))
			// GetProductsTiktokloop(strconv.Itoa(index), index)

			//}

		}

		response.Message = ""

	}
}

func parseProductDetailTiktok(jsonBuffer []byte) models.RespProdutcDetailTiktok {

	RespProdutcDetailTiktok := models.RespProdutcDetailTiktok{}

	err := json.Unmarshal(jsonBuffer, &RespProdutcDetailTiktok)
	if err != nil {
		return RespProdutcDetailTiktok
	}

	// the array is now filled with users
	return RespProdutcDetailTiktok

}

func GetProductDetailTiktok(c *gin.Context) {
	nmRoute := "GetProductDetailTiktok Tiktok"
	product_id := c.Param("product_id")

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/products/details"

	tambahan += "shop_id#" + shop_id
	tambahan += "#product_id#" + product_id

	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id
	url_last += "&product_id=" + product_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseProductDetailTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//fmt.Println(string(data))
		response.Result = datas

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message+" | "+product_id, nmRoute)

		}

		response.Message = ""

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetProductDetailTiktokloop(product_id string) {
	nmRoute := "GetProductDetailTiktokloop Tiktok"

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/products/details"

	tambahan += "shop_id#" + shop_id
	tambahan += "#product_id#" + product_id

	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id
	url_last += "&product_id=" + product_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

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
		fmt.Println("gagal " + nmRoute)
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseProductDetailTiktok(data)
		response.Result = datas

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message+" | "+product_id, nmRoute)
		} else {
			//tokenService.SaveSkuMappingTiktok(datas.Data)
		}

	}

}

func UpdateStockTiktok(c *gin.Context) {
	nmRoute := "UpdateStockTiktok Tiktok"
	product_id := c.Param("product_id")
	sku_id := c.Param("sku_id")
	stocks := c.Param("stock")

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/products/stocks"

	tambahan += "shop_id#" + shop_id

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	var jsonString = `{
		"product_id":"` + product_id + `",
		"skus":[{
		 "id":"` + sku_id + `",
		 "stock_infos":[{
			"available_stock":` + stocks + `
		   }]
		 }]
		}`

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body_url))
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
		helpers.KirimEmail("UpdateStockTiktok", nmRoute+" Gagal Koneksi", "")
		log.Println("Error on response.\n[ERRO] -", err)
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseProductDetailTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//fmt.Println(string(data))
		response.Message = ""
		response.Result = datas

		if datas.Message != "Success" { //save error
			response.Message = "ERROR " + datas.Message
			tokenService.SaveErrorString(ChannelName, datas.Message+"|"+product_id+"|"+sku_id+"|"+stocks, nmRoute)

			//update status jadi NACT jika product status invalid
			if datas.Message == "product status invalid" {
				tokenService.UpdateProductNuse(product_id, sku_id, os.Getenv("KODE_TIKTOK"))
			}
		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func UpdateStockOtomatisTiktok(c *gin.Context) {

	ObjMapping := tokenService.CariSkuMappingObjGroup(os.Getenv("KODE_TIKTOK"))
	//spew.Dump(ObjMapping)
	if len(ObjMapping) > 0 {

		for _, value := range ObjMapping {

			helpers.UpdateStock(value.SkuNo, "API_CHANNEL", os.Getenv("KODE_TIKTOK"))

		}
	}

}

func UpdateStockTiktokLoop(product_id string, sku_id string, stocks string) {
	nmRoute := "UpdateStockTiktokLoop Tiktok"

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/products/stocks"

	tambahan += "shop_id#" + shop_id

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	var jsonString = `{
		"product_id":"` + product_id + `",
		"skus":[{
		 "id":"` + sku_id + `",
		 "stock_infos":[{
			"available_stock":` + stocks + `
		   }]
		 }]
		}`

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body_url))
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
		helpers.KirimEmail("UpdateStockTiktokLoop", nmRoute+" Gagal", "")
		log.Println("Error on response.\n[ERRO] -", err)
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseProductDetailTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//fmt.Println(string(data))

		response.Result = datas

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message+"|"+product_id+"|"+sku_id+"|"+stocks, nmRoute)
			helpers.KirimEmail("UpdateStockTiktokLoop", nmRoute+" | "+datas.Message+"|"+product_id+"|"+sku_id+"|"+stocks, "")
		}

	}
}

func parseHeaderiktok(jsonBuffer []byte) models.HeadRespTiktok {

	HeadRespTiktok := models.HeadRespTiktok{}

	err := json.Unmarshal(jsonBuffer, &HeadRespTiktok)
	if err != nil {
		return HeadRespTiktok
	}

	// the array is now filled with users
	return HeadRespTiktok

}

func ReqShippingTiktokOLD(c *gin.Context) { //31 JULI 2022 #deprecated
	nmRoute := "ReqShippingTiktok Tiktok"
	order_id := c.Param("order_id")

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/order/rts"

	tambahan += "shop_id#" + shop_id

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	var jsonString = `{
		"order_id":"` + order_id + `"
		}`

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		helpers.KirimEmail("ReqShippingTiktok", nmRoute+" "+order_id+" Koneksi Gagal", "")
		log.Println("Error on response.\n[ERRO] -", err)
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseHeaderiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		//fmt.Println(string(data))

		response.Result = datas

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
			response.Message = datas.Message + "(TIKTOK)"
			helpers.KirimEmail("ReqShippingTiktok", nmRoute+" "+order_id+" "+datas.Message, "")
		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func ReqShippingTiktok(c *gin.Context) {
	nmRoute := "ReqShippingTiktok Tiktok"
	order_id := c.Param("order_id")

	//ambil package
	packageId := ""
	TypeShip := ""
	PickStartTime := ""
	PickEndTime := ""
	objDataDetail := GetDataDetailTiktok(order_id)

	if len(objDataDetail.Data.OrderList) > 0 {
		if len(objDataDetail.Data.OrderList[0].PackageList) > 0 {
			packageId = objDataDetail.Data.OrderList[0].PackageList[0].PackageID
			TypeShip, PickStartTime, PickEndTime = GetPickupConfigTiktok(objDataDetail.Data.OrderList[0].PackageList[0].PackageID)
		}

	}
	fmt.Println("PackageId " + packageId)
	fmt.Println("TypeShip " + TypeShip)
	fmt.Println("PickStartTime " + PickStartTime)
	fmt.Println("PickEndTime " + PickEndTime)

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	//path := "/api/order/rts"
	path := "/api/fulfillment/rts" // ganti

	tambahan += "shop_id#" + shop_id

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	var jsonString = `{
		"order_id":"` + order_id + `"
		}`

	StringBody := ""

	if TypeShip == "1" {
		StringBody = `{
			"package_id": "` + packageId + `",
			"pick_up": {
			  "pick_up_end_time": ` + PickEndTime + `,
			  "pick_up_start_time": ` + PickStartTime + `
			},
			"pick_up_type": ` + TypeShip + `
		  }`
	} else {
		StringBody = `{
			"package_id": ` + packageId + `,
			"pick_up_type": ` + TypeShip + `
		  }`
	}

	fmt.Println(StringBody)
	jsonString = StringBody

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		helpers.KirimEmail("ReqShippingTiktok", nmRoute+" "+order_id+" Gagal", "")
		log.Println("Error on response.\n[ERRO] -", err)
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseHeaderiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		//fmt.Println(string(data))

		response.Result = datas

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
			response.Message = datas.Message + "(TIKTOK)"
		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func PackageShipTiktok(order_id string) {
	nmRoute := "PackageShipTiktok Tiktok"

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/fulfillment/pre_combine_pkg/confirm"

	tambahan += "shop_id#" + shop_id

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	var jsonString = `
		{
			"pre_combine_pkg_list": [
					{
						"pre_combine_pkg_id": "` + order_id + `",
						"order_id_list": [
							"` + order_id + `"
						]
					}
			]
		}
		`

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		helpers.KirimEmail("PackageShipTiktok", nmRoute+" "+order_id+" Gagal", "")
		log.Println("Error on response.\n[ERRO] -", err)
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parsePackageTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		fmt.Println("======= " + nmRoute + " =======")
		fmt.Println(string(data))
		fmt.Println(datas)
		response.Result = datas

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
		}

	}

}

func GetDataDetailTiktok(orderid string) models.RespOrderDetailTiktok2 {
	nmRoute := "GetDataDetailTiktok Tiktok"
	var objData models.RespOrderDetailTiktok2
	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/orders/detail/query"

	tambahan += "shop_id#" + shop_id
	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	var jsonString = `{
		"order_id_list":["` + orderid + `"]}`

	//fmt.Println(jsonString)

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseRespOrderDetailTiktok2(data)
		objData = datas
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)
		// fmt.Println("====== " + nmRoute + " ======")
		// fmt.Println(string(data))

		// fmt.Println("====== " + nmRoute + " ======")
		// if len(objData.Data.OrderList) > 0 {
		// 	if len(objData.Data.OrderList[0].PackageList) > 0 {
		// 		fmt.Println("PackageId " + objData.Data.OrderList[0].PackageList[0].PackageID)
		// 		GetPickupConfigTiktok(objData.Data.OrderList[0].PackageList[0].PackageID)
		// 	}

		// }
		// fmt.Println("++++++ " + nmRoute + " ++++++")
		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
			response.Message = datas.Message + "(TIKTOK)"
		} else {
			if len(datas.Data.OrderList) > 0 {
				response.ResponseDesc = strconv.Itoa(int(datas.Data.OrderList[0].OrderStatus))
			}
		}

	}

	return objData
}

func GetPickupConfigTiktok(packageid string) (string, string, string) {
	nmRoute := "GetPickupConfigTiktok Tiktok"
	TypeShip := ""
	PickStartTime := ""
	PickEndTime := ""
	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/fulfillment/package_pickup_config/list"

	tambahan += "shop_id#" + shop_id
	tambahan += "#package_id#" + packageid

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id
	url_last += "&package_id=" + packageid

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

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
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseConfigPickupTiktok(data)
		//objData = datas
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)
		// fmt.Println("====== " + nmRoute + " ======")
		// fmt.Println(string(data))

		fmt.Println("====== " + nmRoute + " ======")
		//fmt.Println(string(data))

		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
			response.Message = datas.Message + "(TIKTOK)"
		} else {

			if len(datas.Data.PickUpTimeList) > 0 {
				TypeShip = "1"
				PickStartTime = datas.Data.PickUpTimeList[0].StartTime
				PickEndTime = datas.Data.PickUpTimeList[0].EndTime
			} else {
				TypeShip = "2" //dropoff
			}
		}
		fmt.Println("++++++ " + nmRoute + " ++++++")

	}

	fmt.Println("TypeShip " + TypeShip)
	fmt.Println("PickStartTime " + PickStartTime)
	fmt.Println("PickEndTime " + PickEndTime)
	return TypeShip, PickStartTime, PickEndTime
}

func parseTrackinfoTiktok(jsonBuffer []byte) models.RespTrackingTiktok {

	RespTrackingTiktok := models.RespTrackingTiktok{}

	err := json.Unmarshal(jsonBuffer, &RespTrackingTiktok)
	if err != nil {
		return RespTrackingTiktok
	}

	// the array is now filled with users
	return RespTrackingTiktok

}

func GetTrackingNoTiktok(c *gin.Context) {
	nmRoute := "GetTrackingNoTiktok Tiktok"
	order_id := c.Param("order_id")

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/logistics/ship/get"

	tambahan += "shop_id#" + shop_id
	tambahan += "#order_id#" + order_id

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id
	url_last += "&order_id=" + order_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

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
		helpers.KirimEmail("GetTrackingNoTiktok", nmRoute+" "+order_id+" Gagal", "")
		log.Println("Error on response.\n[ERRO] -", err)
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseTrackinfoTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//fmt.Println(string(data))
		//response.Result = string(data)
		response.Result = datas

		// if datas.Message != "Success" { //save error
		// 	tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
		// }

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetResiTiktok(c *gin.Context) {
	nmRoute := "GetResiTiktok Tiktok"
	orderid := c.Param("order_id")

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")

	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/orders/detail/query"

	tambahan += "shop_id#" + shop_id
	//tambahan += "#access_token#" + token

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

	var jsonString = `{
		"order_id_list":["` + orderid + `"]}`

	//fmt.Println(jsonString)

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body_url))
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
		helpers.KirimEmail("GetResiTiktok", nmRoute+" "+orderid+" Gagal", "")
		log.Println("Error on response.\n[ERRO] -", err)
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseOrderDetailTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		//response.Result = datas

		response.Message = "Resi Masih Kosong (TIKTOK)"
		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
			response.Message = "Resi Masih Kosong (TIKTOK)"
		} else {
			response.Message = ""
			response.ResponseDesc = datas.Data.OrderList[0].TrackingNumber
		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetDocumnetShipTiktok(c *gin.Context) {
	nmRoute := "GetDocumnetShipTiktok Tiktok"
	orderid := c.Param("order_id")

	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")
	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/logistics/shipping_document"

	tambahan += "shop_id#" + shop_id
	tambahan += "#order_id#" + orderid
	tambahan += "#document_type#SHIPPING_LABEL"
	tambahan += "#document_size#A6"

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id
	url_last += "&order_id=" + orderid
	url_last += "&document_type=SHIPPING_LABEL"
	url_last += "&document_size=A6"

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

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
		helpers.KirimEmail("GetDocumnetShipTiktok", nmRoute+" "+orderid+" Gagal", "")
		log.Println("Error on response.\n[ERRO] -", err)
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseDocumentShipTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Result = datas

		// response.Message = "Resi Masih Kosong (TIKTOK)"
		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
			response.Message = "Resi Masih Kosong (TIKTOK)"
		} else {
			response.Message = ""
		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetDocumnetShipTiktokAuto(orderid string) string {
	nmRoute := "GetDocumnetShipTiktokAuto Tiktok"
	document := ""
	partner_key := os.Getenv("PARTNER_KEY_TIKTOK")
	shop_id := os.Getenv("SHOP_ID_TIKTOK")
	tambahan := ""
	ObjToken := tokenService.FindToken("tiktok")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud
	path := "/api/logistics/shipping_document"

	tambahan += "shop_id#" + shop_id
	tambahan += "#order_id#" + orderid
	tambahan += "#document_type#SHIPPING_LABEL"
	tambahan += "#document_size#A6"

	sign, timestamp, _ := AuthTiktok(path, token, tambahan)

	url_last := ""
	url_last += "?app_key=" + partner_key
	url_last += "&access_token=" + token
	url_last += "&sign=" + sign
	url_last += "&timestamp=" + timestamp
	url_last += "&shop_id=" + shop_id
	url_last += "&order_id=" + orderid
	url_last += "&document_type=SHIPPING_LABEL"
	url_last += "&document_size=A6"

	url := os.Getenv("URL_API_TIKTOK") + path + url_last

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
		helpers.KirimEmail("GetDocumnetShipTiktokAuto", nmRoute+" "+orderid+" Gagal", "")
		log.Println("Error on response.\n[ERRO] -", err)
		fmt.Println("gagal " + nmRoute)
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)
		response.Message = "ERROR"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseDocumentShipTiktok(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Result = datas

		// response.Message = "Resi Masih Kosong (TIKTOK)"
		if datas.Message != "Success" { //save error
			tokenService.SaveErrorString(ChannelName, datas.Message, nmRoute)
			response.Message = datas.Message + " (TIKTOK)"
		} else {
			response.Message = ""

			document = datas.Data.DocURL
		}

	}

	return document

}
