package lazadaController

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
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/helpers"
	"github.com/rals/dearme-channel/models"
	"github.com/rals/dearme-channel/models/response"
	"github.com/rals/dearme-channel/repositories/tokenRepository"
	"github.com/rals/dearme-channel/services/tokenService"
	"github.com/rals/dearme-channel/utils"
	// "encoding/json"
)

var ChannelName = "lazada"
var indexx = 1

func AuthLazada(path string, token string, option string) (string, string, string) {
	now := time.Now()
	timest := strconv.FormatInt(now.Unix(), 10)
	timest = timest + "000"
	//timest = "1657156881848"
	partner_key := os.Getenv("CLIENT_ID_LAZADA")
	partner_secret := os.Getenv("CLIENT_KEY_LAZADA")
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
	input := ""
	for _, key := range keys {
		input = input + key + parameter[key]
	}
	input = path + input

	h := hmac.New(sha256.New, []byte(partner_secret))

	balik := ""
	concatenated := input
	//fmt.Println(input)

	if _, err := h.Write([]byte(input)); err != nil {
		return balik, timest, concatenated
	} else {
		balik = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	}

	return balik, timest, concatenated

}
func parseJsonTokenLazada(jsonBuffer []byte) models.TokenLazada {

	TokenLazada := models.TokenLazada{}

	err := json.Unmarshal(jsonBuffer, &TokenLazada)
	if err != nil {
		return TokenLazada
	}

	// the array is now filled with users
	return TokenLazada

}
func parseJsonErrorStock(jsonBuffer []byte) models.ErrorUpdateStockLazada {

	ErrorUpdateStockLazada := models.ErrorUpdateStockLazada{}

	err := json.Unmarshal(jsonBuffer, &ErrorUpdateStockLazada)
	if err != nil {
		return ErrorUpdateStockLazada
	}

	return ErrorUpdateStockLazada

}

func parseJsonProductsLazada(jsonBuffer []byte) models.ProductsLazada {

	ProductsLazada := models.ProductsLazada{}

	err := json.Unmarshal(jsonBuffer, &ProductsLazada)
	if err != nil {
		return ProductsLazada
	}

	// the array is now filled with users
	return ProductsLazada

}

func parseJsonErrorLazada(jsonBuffer []byte) models.ErrorLazada {

	ErrorLazada := models.ErrorLazada{}

	err := json.Unmarshal(jsonBuffer, &ErrorLazada)
	if err != nil {
		return ErrorLazada
	}

	// the array is now filled with users
	return ErrorLazada

}

func parseJsonProductDetailLazada(jsonBuffer []byte) models.ProductDetailLazada {

	ProductDetailLazada := models.ProductDetailLazada{}

	err := json.Unmarshal(jsonBuffer, &ProductDetailLazada)
	if err != nil {
		return ProductDetailLazada
	}

	// the array is now filled with users
	return ProductDetailLazada

}

func parseJsonOrdersLazada(jsonBuffer []byte) models.OrdersLazada {

	OrdersLazada := models.OrdersLazada{}

	err := json.Unmarshal(jsonBuffer, &OrdersLazada)
	if err != nil {
		return OrdersLazada
	}

	// the array is now filled with users
	return OrdersLazada

}

func parseJsonOrdersDetailLazada(jsonBuffer []byte) models.OrdersDetailGenerated {

	OrdersDetailGenerated := models.OrdersDetailGenerated{}

	err := json.Unmarshal(jsonBuffer, &OrdersDetailGenerated)
	if err != nil {
		return OrdersDetailGenerated
	}

	// the array is now filled with users
	return OrdersDetailGenerated

}

func parseJsonOrdersDetailItemLazada(jsonBuffer []byte) models.OrderDetailItemLazada {

	OrderDetailItemLazada := models.OrderDetailItemLazada{}

	err := json.Unmarshal(jsonBuffer, &OrderDetailItemLazada)
	if err != nil {
		return OrderDetailItemLazada
	}

	// the array is now filled with users
	return OrderDetailItemLazada

}

func parseJsonTransactionmLazada(jsonBuffer []byte) models.DetalTransactionLazada {

	DetalTransactionLazada := models.DetalTransactionLazada{}

	err := json.Unmarshal(jsonBuffer, &DetalTransactionLazada)
	if err != nil {
		return DetalTransactionLazada
	}

	// the array is now filled with users
	return DetalTransactionLazada

}

func parseJsonSetStatusLazada(jsonBuffer []byte) models.SetPackedLazada {

	SetPackedLazada := models.SetPackedLazada{}

	err := json.Unmarshal(jsonBuffer, &SetPackedLazada)
	if err != nil {
		return SetPackedLazada
	}

	// the array is now filled with users
	return SetPackedLazada

}

func GetTokenLazada(c *gin.Context) {

	var response response.ResponseCrud

	urllazada := os.Getenv("URL_AUTH_LAZADA") + "oauth/authorize"

	urllazada += "?response_type=code&force_auth=true&redirect_uri=" + os.Getenv("URI_CALLBACK_LAZADA") + "saveToken" + "&client_id=" + os.Getenv("CLIENT_ID_LAZADA")

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "Req Token LAZADA GAGAL (connection)"
		response.Result = err

	} else {
		fmt.Println("masuk siniii")
		defer resp.Body.Close()
		//data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJsonErrorTokped(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urllazada
		//response.Result = datas
		// fmt.Println(string(data))

	}

	c.JSON(http.StatusOK, response)
	return
}

func SaveTokenLazada(c *gin.Context) {

	var response response.ResponseCrud
	code := c.Query("code")
	urllazada := os.Getenv("URL_API_TOKEN_LAZADA")
	endpoint := "/auth/token/create"
	signMethode := "sha256"
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#code#"+code)
	//sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&code=" + code

	req, err := http.NewRequest("POST", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "Req Token LAZADA GAGAL (connection)"
		response.Result = err

	} else {
		fmt.Println("masuk siniii")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonTokenLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urllazada
		response.Result = string(data)

		var objToken models.TokenJdId

		if datas.AccessToken != "" {
			objToken.AccessToken = datas.AccessToken
			objToken.RefreshToken = datas.RefreshToken

			tokenService.SaveTokenJdId(objToken, ChannelName)
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

func RefreshToken(c *gin.Context) {

	var response response.ResponseCrud
	ObjToken := tokenService.FindToken("lazada")
	refreshtoken := fmt.Sprintf("%v", ObjToken.Value2)

	urllazada := os.Getenv("URL_API_TOKEN_LAZADA")
	endpoint := "/auth/token/refresh"
	signMethode := "sha256"
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#refresh_token#"+refreshtoken)
	//sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&refresh_token=" + refreshtoken

	req, err := http.NewRequest("POST", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "Req Token LAZADA GAGAL (connection)"
		response.Result = err

	} else {
		fmt.Println("masuk siniii")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonTokenLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urllazada
		response.Result = string(data)

		var objToken models.TokenJdId

		if datas.AccessToken != "" {
			objToken.AccessToken = datas.AccessToken
			objToken.RefreshToken = datas.RefreshToken

			tokenService.SaveTokenJdId(objToken, ChannelName)
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

func RefreshTokenAuto() {

	var response response.ResponseCrud
	ObjToken := tokenService.FindToken("lazada")
	refreshtoken := fmt.Sprintf("%v", ObjToken.Value2)

	urllazada := os.Getenv("URL_API_TOKEN_LAZADA")
	endpoint := "/auth/token/refresh"
	signMethode := "sha256"
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#refresh_token#"+refreshtoken)
	//sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&refresh_token=" + refreshtoken

	req, err := http.NewRequest("POST", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "Req Token LAZADA GAGAL (connection)"
		response.Result = err

	} else {
		fmt.Println("masuk siniii")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonTokenLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urllazada
		response.Result = string(data)

		var objToken models.TokenJdId

		if datas.AccessToken != "" {
			objToken.AccessToken = datas.AccessToken
			objToken.RefreshToken = datas.RefreshToken

			tokenService.SaveTokenJdId(objToken, ChannelName)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			response.Message = "Save Token Sukses"
			response.Result = datas
		}

		fmt.Println(string(data))

	}

}

func GetProducts(c *gin.Context) {
	nmRoute := "GetProducts"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	limit := os.Getenv("LIMIT_PRODUCT_LAZADA")
	offset := 0

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/products/get"
	signMethode := "sha256"

	setelah := time.Now().Add(time.Duration(-72) * time.Hour).Format(time.RFC3339) // 3 hari sebelumnya
	sebelum := time.Now().Format(time.RFC3339)
	//setelah := time.Now().Add(time.Duration(-72) * time.Hour).Unix() // 3 hari sebelumnya
	//sebelum := time.Now().Unix()

	tambahan := "#limit#" + limit
	tambahan += "#offset#" + strconv.Itoa(offset)
	//tambahan += "#update_before#" + sebelum
	// tambahan += "#create_before#" + sebelum

	//tambahan += "#create_after#" + setelah
	//tambahan += "#update_after#" + setelah

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&limit=" + limit
	urllazada += "&offset=" + strconv.Itoa(offset)
	//urllazada += "&update_before=" + sebelum
	// urllazada += "&create_before=" + sebelum
	//urllazada += "&create_after=" + setelah
	//urllazada += "&update_after=" + setelah
	fmt.Println(urllazada)
	fmt.Println("======= GET Products LAZADA =======")

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductsLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

			for _, val := range datas.Data.Products {
				//fmt.Println(" Item ID " + strconv.Itoa(int(val.ItemID)))
				GetProductDetailLoop(int(val.ItemID))
				//fmt.Println("[" + strconv.Itoa(indexx) + "] " + strconv.Itoa(int(val.ItemID)))
				//indexx++
				// if strconv.Itoa(int(val.ItemID)) == "6517166032" {
				// 	fmt.Println("ada " + strconv.Itoa(int(val.ItemID)))

				// }
			}

			if len(datas.Data.Products) > 0 {

				limits, _ := strconv.Atoi(limit)

				offset = limits
				GetProductsLoop(limits, offset, sebelum, setelah)

			}

		}

	}

	fmt.Println("======= SELESAI Products LAZADA =======")

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsAuto(wg *sync.WaitGroup) {
	nmRoute := "GetProductsAuto"
	var response response.ResponseCrud
	fmt.Println("mulai product lazada " + time.Now().String())
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	limit := os.Getenv("LIMIT_PRODUCT_LAZADA")
	offset := 0

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/products/get"
	signMethode := "sha256"

	setelah := time.Now().Add(time.Duration(-72) * time.Hour).Format(time.RFC3339) // 3 hari sebelumnya
	sebelum := time.Now().Format(time.RFC3339)

	tambahan := "#limit#" + limit
	tambahan += "#offset#" + strconv.Itoa(offset)
	// tambahan += "#update_before#" + sebelum
	// tambahan += "#create_before#" + sebelum

	// tambahan += "#create_after#" + setelah
	// tambahan += "#update_after#" + setelah

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&limit=" + limit
	urllazada += "&offset=" + strconv.Itoa(offset)
	// urllazada += "&update_before=" + sebelum
	// urllazada += "&create_before=" + sebelum
	// urllazada += "&create_after=" + setelah
	// urllazada += "&update_after=" + setelah

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductsLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

			for _, val := range datas.Data.Products {
				//fmt.Println(" Item ID " + strconv.Itoa(int(val.ItemID)))
				GetProductDetailLoop(int(val.ItemID))
			}

			if len(datas.Data.Products) > 0 {

				limits, _ := strconv.Atoi(limit)

				offset = limits
				GetProductsLoop(limits, offset, sebelum, setelah)

			}

		}

	}

	fmt.Println("selesai product lazada " + time.Now().String())
	wg.Done()
}

func GetProductsLoop(limit int, offset int, sebelum, setelah string) {
	nmRoute := "GetProducts"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/products/get"
	signMethode := "sha256"

	tambahan := "#limit#" + strconv.Itoa(limit)
	tambahan += "#offset#" + strconv.Itoa(offset)
	//tambahan += "#update_before#" + sebelum
	//tambahan += "#create_before#" + sebelum

	// tambahan += "#create_after#" + setelah
	// tambahan += "#update_after#" + setelah
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&limit=" + strconv.Itoa(limit)
	urllazada += "&offset=" + strconv.Itoa(offset)
	//urllazada += "&update_before=" + sebelum
	//urllazada += "&create_before=" + sebelum
	// urllazada += "&create_after=" + setelah
	// urllazada += "&update_after=" + setelah

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductsLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

			for _, val := range datas.Data.Products {
				//fmt.Println(" Item ID " + strconv.Itoa(int(val.ItemID)))
				GetProductDetailLoop(int(val.ItemID))
				// offset++
				//fmt.Println("[" + strconv.Itoa(indexx) + "] " + strconv.Itoa(int(val.ItemID)))
				//indexx++
				// 	if strconv.Itoa(int(val.ItemID)) == "6517166032" {
				// 		fmt.Println("ada " + strconv.Itoa(int(val.ItemID)))

				// 	}
			}
			if len(datas.Data.Products) > 0 {
				offset += limit
				//fmt.Println("offset " + strconv.Itoa(offset))
				GetProductsLoop(limit, offset, sebelum, setelah)
			}

		}

	}
}

func GetProductDetailLoop(itemid int) {
	nmRoute := "GetProductDetailLoop"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/product/item/get"
	signMethode := "sha256"

	tambahan := "#item_id#" + strconv.Itoa(itemid)
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&item_id=" + strconv.Itoa(itemid)

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

			tokenService.SaveSkuMappingLazada(datas, itemid)

		}

	}
}

func GetProductDetail(c *gin.Context) {
	nmRoute := "GetProductDetail"
	itemid := c.Param("itemid")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/product/item/get"
	signMethode := "sha256"

	tambahan := "#item_id#" + itemid
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&item_id=" + itemid

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetProductDetailSku(c *gin.Context) {
	nmRoute := "GetProductDetailSku"
	itemid := c.Param("sku")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/product/item/get"
	signMethode := "sha256"

	tambahan := "#seller_sku#" + itemid
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&seller_sku=" + itemid

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetOrder(c *gin.Context) {
	nmRoute := "GetOrder"
	statusOrder := c.Param("status")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/orders/get"
	signMethode := "sha256"

	offset := 0
	limit := 50

	updateAfter := time.Now().Add(time.Duration(-72) * time.Hour).Format(time.RFC3339) //3 hari sebelumnya

	tambahan := "#limit#" + strconv.Itoa(limit)
	tambahan += "#offset#" + strconv.Itoa(offset)
	tambahan += "#update_after#" + updateAfter
	tambahan += "#status#" + statusOrder

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&limit=" + strconv.Itoa(limit)
	urllazada += "&offset=" + strconv.Itoa(offset)
	urllazada += "&update_after=" + url.QueryEscape(updateAfter)
	urllazada += "&status=" + statusOrder

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {
			if len(datas.Data.Orders) > 0 {
				for _, val := range datas.Data.Orders {
					GetOrderDetailItemLoop(strconv.Itoa(int(val.OrderID)))
				}
				fmt.Println("masuk loop awal")
				offset += limit
				GetOrderLoop(limit, offset, statusOrder)
			}
		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetOrderAuto(statusOrder string) {
	nmRoute := "GetOrderAuto"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/orders/get"
	signMethode := "sha256"

	offset := 0
	limit := 50

	updateAfter := time.Now().Add(time.Duration(-72) * time.Hour).Format(time.RFC3339) //3 hari sebelumnya

	tambahan := "#limit#" + strconv.Itoa(limit)
	tambahan += "#offset#" + strconv.Itoa(offset)
	tambahan += "#update_after#" + updateAfter
	tambahan += "#status#" + statusOrder

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&limit=" + strconv.Itoa(limit)
	urllazada += "&offset=" + strconv.Itoa(offset)
	urllazada += "&update_after=" + url.QueryEscape(updateAfter)
	urllazada += "&status=" + statusOrder

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {
			if len(datas.Data.Orders) > 0 {
				for _, val := range datas.Data.Orders {
					//cek jika ada di salesorder skip
					objCek, _ := tokenRepository.FindSalesOrder(strconv.Itoa(int(val.OrderID)))
					if objCek.NoOrder == "" {
						GetOrderDetailItemLoop(strconv.Itoa(int(val.OrderID)))
					} else if statusOrder == "canceled" && objCek.StatusProcessOrder == "0" {
						GetOrderDetailItemLoop(strconv.Itoa(int(val.OrderID)))
					}

				}
				fmt.Println("masuk loop awal")
				offset += limit
				GetOrderLoop(limit, offset, statusOrder)
			}
		}

	}
}

func GetOrderLoop(limit int, offset int, statusOrder string) {

	nmRoute := "GetOrderLoop"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/orders/get"
	signMethode := "sha256"

	updateAfter := time.Now().Add(time.Duration(-72) * time.Hour).Format(time.RFC3339) //3 hari sebelumnya

	tambahan := "#limit#" + strconv.Itoa(limit)
	tambahan += "#offset#" + strconv.Itoa(offset)
	tambahan += "#update_after#" + updateAfter
	tambahan += "#status#" + statusOrder

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&limit=" + strconv.Itoa(limit)
	urllazada += "&offset=" + strconv.Itoa(offset)
	urllazada += "&update_after=" + url.QueryEscape(updateAfter)
	urllazada += "&status=" + statusOrder

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

			if len(datas.Data.Orders) > 0 {

				for _, val := range datas.Data.Orders {
					//GetOrderDetailItemLoop(strconv.Itoa(int(val.OrderID)))
					//cek jika ada di salesorder skip
					objCek, _ := tokenRepository.FindSalesOrder(strconv.Itoa(int(val.OrderID)))
					if objCek.NoOrder == "" {
						GetOrderDetailItemLoop(strconv.Itoa(int(val.OrderID)))
					} else if statusOrder == "canceled" && objCek.StatusProcessOrder == "0" {
						GetOrderDetailItemLoop(strconv.Itoa(int(val.OrderID)))
					}
				}

				fmt.Println("masuk loop")
				offset += limit
				GetOrderLoop(limit, offset, statusOrder)
			}
		}

	}

}

func GetOrderDetailItemLoop(orderId string) {
	nmRoute := "GetOrderDetailItem"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/order/items/get"
	signMethode := "sha256"

	tambahan := "#order_id#" + orderId

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&order_id=" + orderId

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersDetailItemLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {
			//fmt.Println("=======================")
			//fmt.Println("ORDERID : " + orderId)
			objCust := GetOrderDetailLoop(orderId)

			NoOrder := strconv.Itoa(int(objCust.Data.OrderNumber))
			if NoOrder != "" {
				//objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
				CariLogOrder := tokenRepository.CariOrderLogAPI(os.Getenv("KODE_LAZADA"), NoOrder)
				if CariLogOrder.UuidLogOrder == "" {
					objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
					if objCek.UuidSalesOrder == "" {
						var objlogOrder models.TableLogOrder
						objlogOrder.UuidLogOrder = uuid.New().String()
						objlogOrder.ChannelCode = os.Getenv("KODE_LAZADA")
						objlogOrder.Orderid = NoOrder
						objlogOrder.Response = string(data)
						objlogOrder.CreatedDate = time.Now()
						tokenRepository.SaveOrderAPI(objlogOrder)
					}

				}

			}

			tokenService.SaveSalesOrderLazada(datas, objCust)
			tokenService.UpdateAmountLazada(datas, objCust)
			// for _, val := range datas.Data {
			// 	fmt.Println(val.Sku)
			// }
			//fmt.Println("=======================")

		}

	}

}
func GetOrderDetailLoop(orderId string) models.OrdersDetailGenerated {
	nmRoute := "GetOrderDetailLoop"
	var response response.ResponseCrud

	var objDetail models.OrdersDetailGenerated
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/order/get"
	signMethode := "sha256"

	tambahan := "#order_id#" + orderId

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&order_id=" + orderId

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersDetailLazada(data)
		objDetail = datas
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

		}

	}

	return objDetail
}

func GetOrderDetailItemsLoop(orderId string) models.OrderDetailItemLazada {
	nmRoute := "GetOrderDetailItemsLoop"
	var response response.ResponseCrud

	var objDetail models.OrderDetailItemLazada
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/order/items/get"
	signMethode := "sha256"

	tambahan := "#order_id#" + orderId

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&order_id=" + orderId

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersDetailItemLazada(data)
		objDetail = datas
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

		}

	}

	return objDetail
}

func GetOrderDetail(c *gin.Context) {
	nmRoute := "GetOrderDetail"
	var response response.ResponseCrud

	orderId := c.Param("orderid")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/order/get"
	signMethode := "sha256"

	tambahan := "#order_id#" + orderId

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&order_id=" + orderId

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersDetailLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

			statusesx := "canceled"
			for index, _ := range datas.Data.Statuses {
				//if statusesx == "" {
				statusesx = datas.Data.Statuses[index]
				//}
				if datas.Data.Statuses[index] == "pending" {
					statusesx = datas.Data.Statuses[index]
					//break
				}

				if datas.Data.Statuses[index] == "canceled" {
					GetOrderDetailItemLoop(orderId)
				} else {
					GetOrderDetailItemLoop(orderId)
				}

			}
			response.ResponseDesc = statusesx
			//response.ResponseDesc = datas.Data.Statuses[0]
		}

	}
	c.JSON(http.StatusOK, response)
	return
}
func GetOrderDetailItem(c *gin.Context) {
	nmRoute := "GetOrderDetailItem"
	var response response.ResponseCrud

	orderId := c.Param("orderid")
	status := c.Param("status")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/order/items/get"
	signMethode := "sha256"

	tambahan := "#order_id#" + orderId

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&order_id=" + orderId

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersDetailItemLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

			if status == "PACK" {

				if len(datas.Data) > 0 {

					orderitemids := ""
					// for indexs, val := range datas.Data {
					for _, val := range datas.Data {
						if val.Status == "pending" {

							if orderitemids == "" {
								orderitemids += strconv.Itoa(int(val.OrderItemID))
							} else {
								orderitemids += "," + strconv.Itoa(int(val.OrderItemID))
							}
							// if indexs > 0  {
							// 	orderitemids += "," + strconv.Itoa(int(val.OrderItemID))
							// } else {
							// 	orderitemids += strconv.Itoa(int(val.OrderItemID))
							// }

						}

					}
					orderitemids = "[" + orderitemids + "]"

					//fmt.Println("order_item_ids=" + orderitemids)
					pesanPack, objShip := GetPackedStatus(orderId, orderitemids, "dropship", "pack")
					if pesanPack == "" {
						fmt.Println("set packing SUKSES")
						//update kurir
						fmt.Println(objShip.Data.OrderItems[0].ShipmentProvider)

						var objUpdateKurir models.TableSalesOrder
						objUpdateKurir.NoOrder = orderId
						objUpdateKurir.ExpeditionType = objShip.Data.OrderItems[0].ShipmentProvider
						objUpdateKurir.CreatedBy = "LAZADA"
						result := tokenRepository.UpdateKurirNew(objUpdateKurir)
						if result != nil {
							fmt.Println("UpdateKurir GetPackedStatus LAZADA GAGAL")
						} else {
							fmt.Println("UpdateKurir GetPackedStatus LAZADA SUKSES")
						}

						//set readytoship
						pesanRts, _ := GetPackedStatus(orderId, orderitemids, "dropship", "rts")
						if pesanRts == "" {
							fmt.Println("set rts SUKSES")
						} else {
							fmt.Println("set rts GAGAL")
							response.Message = pesanPack
						}

					} else {
						fmt.Println("set packing GAGAL")
						response.Message = pesanPack
					}

				}

			}

		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func GetPackedStatus(orderId string, orderIds string, statuship string, param1 string) (string, models.SetPackedLazada) {
	nmRoute := "GetPackedStatus"
	var response response.ResponseCrud

	var obj models.SetPackedLazada
	pesan := ""
	//orderIds := "[901656018169007,901656018269007]"

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/order/pack"

	if param1 == "rts" {
		endpoint = "/order/rts"
	}

	signMethode := "sha256"

	tambahan := "#order_item_ids#" + orderIds
	tambahan += "#delivery_type#" + statuship

	fmt.Println(orderIds)
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&order_item_ids=" + url.QueryEscape(orderIds)
	urllazada += "&delivery_type=" + statuship

	req, err := http.NewRequest("POST", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		pesan = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonSetStatusLazada(data)
		obj = datas
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			pesan = pesanErr
			//response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {
			fmt.Println(string(data))

		}

	}

	return pesan, obj

}

func GetResi(c *gin.Context) {
	nmRoute := "GetResi"
	var response response.ResponseCrud

	orderId := c.Param("orderid")
	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/order/items/get"
	signMethode := "sha256"

	tambahan := "#order_id#" + orderId

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&order_id=" + orderId

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonOrdersDetailItemLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

			statuses := ""
			for _, val := range datas.Data {
				if val.TrackingCode != "" {

					shipmentnya := val.ShipmentProvider
					CekKurirJdId := strings.Split(val.ShipmentProvider, ",")

					if len(CekKurirJdId) > 1 {
						shipmentnya = strings.Replace(CekKurirJdId[1], " Delivery: ", "", -1)
					}

					statuses = val.TrackingCode + "^" + shipmentnya
					break
				}

			}
			response.ResponseDesc = statuses
			//response.ResponseDesc = datas.Data[0].TrackingCode

		}

	}
	c.JSON(http.StatusOK, response)
	return
}

func UpdateStock(c *gin.Context) {
	nmRoute := "UpdateStock"
	ItemId := c.Param("itemid")
	SkuId := c.Param("skuid")
	Qty := c.Param("qty")
	Sku := c.Param("sku")
	var response response.ResponseCrud

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/product/price_quantity/update"
	signMethode := "sha256"

	bodyXml := `
			<Request>
				<Product>
					<Skus>
					<Sku>
						<ItemId>` + ItemId + `</ItemId>
						<SkuId>` + SkuId + `</SkuId>
						<Quantity>` + Qty + `</Quantity>
					</Sku>
					</Skus>
				</Product>
		</Request>`

	if ItemId == "6517166032" && SkuId == "12395682763" {
		fmt.Println("masuk sisni")
		// endpoint = "/product/stock/sellable/adjust"
		// bodyXml = `
		// 	<Request>
		// 		<Product>
		// 			<Skus>
		// 			<Sku>
		// 				<ItemId>` + ItemId + `</ItemId>
		// 				<SkuId>` + SkuId + `</SkuId>
		// 				<SellableQuantity>1</SellableQuantity>
		// 			</Sku>
		// 			</Skus>
		// 		</Product>
		// </Request>`

	}

	var objRest models.TableLogStock
	objRest.UuidLog = uuid.New().String()
	objRest.ChannelCode = os.Getenv("KODE_LAZADA")
	objRest.Sku = Sku
	objRest.Body = bodyXml
	if s, err := strconv.ParseFloat(Qty, 64); err == nil {
		objRest.Stock = s
	}

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	// urllazada := os.Getenv("URL_API_LAZADA")
	// endpoint := "/product/price_quantity/update"
	// signMethode := "sha256"

	// bodyXml := `
	// 		<Request>
	// 			<Product>
	// 				<Skus>
	// 				<Sku>
	// 					<ItemId>` + ItemId + `</ItemId>
	// 					<SkuId>` + SkuId + `</SkuId>
	// 					<Quantity>` + Qty + `</Quantity>
	// 				</Sku>
	// 				</Skus>
	// 			</Product>
	// 	</Request>`

	tambahan := "#payload#" + bodyXml
	//tambahan = ""
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&payload=" + url.QueryEscape(bodyXml)

	var body_url = []byte(bodyXml)

	req, err := http.NewRequest("POST", urllazada, bytes.NewBuffer(body_url))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)
		objRest.Response = nmRoute + " Gagal Koneksi"

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJsonOrdersDetailItemLazada(data)
		//datasErr := parseJsonErrorLazada(data)
		datasErr := parseJsonErrorStock(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)
		objRest.Response = string(data)
		if datasErr.Message != "" {
			respUpd := "test"
			if len(datasErr.Detail) > 0 {
				if datasErr.Detail[0].SellerSku == "INV_NEGATIVE_SELLABLE" {
					//fmt.Println("update stock")
					qtyold := GetProductDetailStock(ItemId, SkuId)
					//fmt.Println(" qty " + qtyold)
					//respUpd = UpdateStockSalesAble(ItemId, SkuId, Qty, qtyold)
					responnya := ""
					respUpd, responnya = UpdateStockSalesAbleNew(ItemId, SkuId, Qty, qtyold, Sku)

					response.Result = responnya
					//fmt.Println(respUpd)
				}
				//cek stock sellable

			}
			if respUpd != "" {
				pesanErr := datasErr.Message + " | " + ItemId + " | " + SkuId + " | " + Qty
				response.Message = pesanErr
				tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)
			}
			// pesanErr := datasErr.Message + " | " + ItemId + " | " + SkuId + " | " + Qty
			// response.Message = pesanErr
			// //response.Result = datasErr
			// tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {

		}

	}

	objRest.CreatedBy = "API"
	objRest.CreatedDate = time.Now()
	tokenRepository.SaveStockAPI(objRest)
	c.JSON(http.StatusOK, response)
	return
}

func UpdateStockV2(c *gin.Context) {
	nmRoute := "UpdateStockV2"
	ItemId := c.Param("itemid")
	SkuId := c.Param("skuid")
	Qty := c.Param("qty")
	var response response.ResponseCrud

	response.Message = ""
	//fmt.Println("update stock")
	qtyold := GetProductDetailStock(ItemId, SkuId)
	//fmt.Println(" qty " + qtyold)
	respUpd := UpdateStockSalesAble(ItemId, SkuId, "0", qtyold)
	response.Message = respUpd
	respUpdnew := ""
	if respUpd == "" {
		time.Sleep(1 * time.Second)
		respUpdnew = UpdateStockSalesAble(ItemId, SkuId, Qty, qtyold)
		if respUpdnew != "" {
			response.Message = respUpdnew
		}
	}

	fmt.Println(respUpd)
	fmt.Println(respUpdnew)

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()

	response.Result = nmRoute

	c.JSON(http.StatusOK, response)
	return
}

func GetProductDetailStock(itemid, skuid string) string {
	qtyold := ""
	nmRoute := "GetProductDetailStock"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/product/item/get"
	signMethode := "sha256"

	tambahan := "#item_id#" + itemid
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&item_id=" + itemid

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonProductDetailLazada(data)
		datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datasErr.Message != "" {
			pesanErr := datasErr.Message
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		}
		for _, val := range datas.Data.Skus {
			if strconv.Itoa(int(val.SkuID)) == skuid {
				qtyold = strconv.Itoa(val.Quantity)

			}
		}

	}

	return qtyold
}

func UpdateStockSalesAble(ItemId, SkuId, QtyNew, QtyOld string) string {
	nmRoute := "UpdateStockSalesAble"
	var response response.ResponseCrud
	res := "test"
	signMethode := "sha256"
	if QtyNew == "0" {
		QtyNew = "-" + QtyOld
	}
	urllazada := os.Getenv("URL_API_LAZADA")

	endpoint := "/product/stock/sellable/adjust"
	bodyXml := `
			<Request>
				<Product>
					<Skus>
					<Sku>
						<ItemId>` + ItemId + `</ItemId>
						<SkuId>` + SkuId + `</SkuId>
						<SellableQuantity>` + QtyNew + `</SellableQuantity>
					</Sku>
					</Skus>
				</Product>
		</Request>`

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	tambahan := "#payload#" + bodyXml
	//tambahan = ""
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&payload=" + url.QueryEscape(bodyXml)

	var body_url = []byte(bodyXml)

	req, err := http.NewRequest("POST", urllazada, bytes.NewBuffer(body_url))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJsonOrdersDetailItemLazada(data)
		//datasErr := parseJsonErrorLazada(data)
		datasErr := parseJsonErrorStock(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)
		if datasErr.Message != "" {

			pesanErr := datasErr.Message + " | " + ItemId + " | " + SkuId + " | " + QtyNew
			response.Message = pesanErr
			//response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {
			res = ""
		}

	}

	return res
}

func UpdateStockSalesAbleNew(ItemId, SkuId, QtyNew, QtyOld, Sku string) (string, string) {
	nmRoute := "UpdateStockSalesAbleNew"
	var response response.ResponseCrud
	res := "test"
	respon := ""
	signMethode := "sha256"
	if QtyNew == "0" {
		QtyNew = "-" + QtyOld
	}
	urllazada := os.Getenv("URL_API_LAZADA")

	endpoint := "/product/stock/sellable/adjust"
	bodyXml := `
			<Request>
				<Product>
					<Skus>
					<Sku>
						<ItemId>` + ItemId + `</ItemId>
						<SkuId>` + SkuId + `</SkuId>
						<SellableQuantity>` + QtyNew + `</SellableQuantity>
					</Sku>
					</Skus>
				</Product>
		</Request>`

	var objRest models.TableLogStock
	objRest.UuidLog = uuid.New().String()
	objRest.ChannelCode = os.Getenv("KODE_LAZADA")
	objRest.Sku = Sku
	objRest.Body = bodyXml
	if s, err := strconv.ParseFloat(QtyNew, 64); err == nil {
		objRest.Stock = s
	}

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	tambahan := "#payload#" + bodyXml
	//tambahan = ""
	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&payload=" + url.QueryEscape(bodyXml)

	var body_url = []byte(bodyXml)

	req, err := http.NewRequest("POST", urllazada, bytes.NewBuffer(body_url))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)
		objRest.Response = nmRoute + " Gagal Koneksi"
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJsonOrdersDetailItemLazada(data)
		//datasErr := parseJsonErrorLazada(data)
		datasErr := parseJsonErrorStock(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)
		respon = string(data)
		objRest.Response = string(data)

		if datasErr.Message != "" {

			pesanErr := datasErr.Message + " | " + ItemId + " | " + SkuId + " | " + QtyNew
			response.Message = pesanErr
			//response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		} else {
			res = ""
		}

	}

	objRest.CreatedBy = "API"
	objRest.CreatedDate = time.Now()
	tokenRepository.SaveStockAPI(objRest)

	return res, respon
}

func GetDocumentShip(c *gin.Context) {
	nmRoute := "GetDocumentShip"
	var response response.ResponseCrud

	orderId := c.Param("orderid")

	objDetailItems := GetOrderDetailItemsLoop(orderId)

	orderitemids := ""
	if len(objDetailItems.Data) > 0 {

		for indexs, val := range objDetailItems.Data {
			if indexs > 0 {
				orderitemids += "," + strconv.Itoa(int(val.OrderItemID))
			} else {
				orderitemids += strconv.Itoa(int(val.OrderItemID))
			}

		}

	}

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/order/document/get"
	signMethode := "sha256"
	orderIds := "[" + orderitemids + "]"
	tambahan := "#order_item_ids#" + orderIds
	tambahan += "#doc_type#shippingLabel"
	fmt.Println(orderIds)

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&order_item_ids=" + url.QueryEscape(orderIds)
	urllazada += "&doc_type=shippingLabel"

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJsonOrdersDetailItemLazada(data)
		//datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)
		// if datasErr.Message != "" {
		// 	pesanErr := datasErr.Message
		// 	response.Message = pesanErr
		// 	response.Result = datasErr
		// 	tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)

		// } else {

		// 	response.ResponseDesc = datas.Data[0].TrackingCode

		// }

	}
	c.JSON(http.StatusOK, response)
	return
}

func UpdateAllStockLazada(c *gin.Context) {
	var response response.ResponseCrud

	ObjMapping := tokenService.CariSkuMappingObjGroup(os.Getenv("KODE_LAZADA")) //param by channel name

	if len(ObjMapping) > 0 {
		for _, value := range ObjMapping {
			helpers.UpdateStock(value.SkuNo, "API_CHANNEL", os.Getenv("KODE_LAZADA"))

		}
	}

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "DISINISS"

	c.JSON(http.StatusOK, response)
	return
}

func GetTransactionDetails(c *gin.Context) {
	nmRoute := "GetTransactionDetails"
	var response response.ResponseCrud

	OrderId := c.Query("OrderId")

	ObjToken := tokenService.FindToken(ChannelName)
	token := fmt.Sprintf("%v", ObjToken.Value1)

	urllazada := os.Getenv("URL_API_LAZADA")
	endpoint := "/finance/transaction/details/get"
	signMethode := "sha256"

	starttime := time.Now().Add(time.Duration(-720) * time.Hour).Format("2006-01-02") //30 hari sebelumnya
	endtime := time.Now().Format("2006-01-02")
	limit := 10
	tambahan := "#start_time#" + starttime
	tambahan += "#end_time#" + endtime
	tambahan += "#limit#" + strconv.Itoa(limit)

	if OrderId != "" {
		tambahan += "#trade_order_id#" + OrderId
	}

	sign, timest, _ := AuthLazada(endpoint, "", "sign_method#"+signMethode+"#access_token#"+token+tambahan)

	urllazada += endpoint
	urllazada += "?app_key=" + os.Getenv("CLIENT_ID_LAZADA") + "&timestamp=" + timest + "&sign_method=" + signMethode + "&sign=" + sign
	urllazada += "&access_token=" + token
	urllazada += "&start_time=" + url.QueryEscape(starttime)
	urllazada += "&end_time=" + url.QueryEscape(endtime)
	urllazada += "&limit=" + strconv.Itoa(limit)
	if OrderId != "" {
		urllazada += "&trade_order_id=" + url.QueryEscape(OrderId)
	}

	req, err := http.NewRequest("GET", urllazada, bytes.NewBuffer(nil))
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
		response.Message = "ERROR"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonTransactionmLazada(data)
		//datasErr := parseJsonErrorLazada(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

	}
	c.JSON(http.StatusOK, response)
	return
}
