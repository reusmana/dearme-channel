package bukalapakController

import (
	"bytes"
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
	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/models"
	"github.com/rals/dearme-channel/models/response"
	"github.com/rals/dearme-channel/repositories/tokenRepository"
	"github.com/rals/dearme-channel/services/tokenService"

	"github.com/rals/dearme-channel/helpers"
	"github.com/rals/dearme-channel/utils"
	// "encoding/json"
)

var ChannelName = "bukalapak"

func AuthBukalapak() (string, string) {
	urlbukalapak := os.Getenv("URL_API_BUKALAPAK")
	token := "asasasdadadadad"

	return urlbukalapak, token

}

func ReqToken(c *gin.Context) {

	url_client := os.Getenv("URI_CALLBACK_BUKALAPAK") + "getCode"
	client_id := os.Getenv("CLIENT_ID_BUKALAPAK")
	scope := "public user store"
	response_type := "code"

	var response response.ResponseCrud

	url := os.Getenv("URL_AUTH_BUKALAPAK") + "/oauth/authorize"

	tambah := "?client_id=" + client_id
	tambah += "&redirect_uri=" + url_client
	tambah += "&scope=" + scope
	tambah += "&response_type=" + response_type

	url = url + tambah

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
		fmt.Println("Gagal !")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func parseAuthBukalapak(jsonBuffer []byte) models.TokenBukalapak {

	TokenBukalapak := models.TokenBukalapak{}

	err := json.Unmarshal(jsonBuffer, &TokenBukalapak)
	if err != nil {
		return TokenBukalapak
	}

	// the array is now filled with users
	return TokenBukalapak

}

func parseErrorBukalapak(jsonBuffer []byte) models.ErrorBukalapak {

	ErrorBukalapak := models.ErrorBukalapak{}

	err := json.Unmarshal(jsonBuffer, &ErrorBukalapak)
	if err != nil {
		return ErrorBukalapak
	}

	// the array is now filled with users
	return ErrorBukalapak

}

func SaveToken(c *gin.Context) {
	nmRoute := "SaveToken Bukalapak"
	code := c.Query("code")
	url_client := os.Getenv("URI_CALLBACK_BUKALAPAK") + "saveToken"
	client_secret := os.Getenv("CLIENT_SECRET_BUKALAPAK")
	client_id := os.Getenv("CLIENT_ID_BUKALAPAK")

	var response response.ResponseCrud

	url := os.Getenv("URL_AUTH_BUKALAPAK") + "/oauth/token"

	var jsonString = `{"grant_type":"authorization_code"
	,"client_id":"` + client_id + `",
	"client_secret":"` + client_secret + `",
	"code":"` + code + `",
	"redirect_uri":"` + url_client + `"}`

	// fmt.Println(jsonString)

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
		response.Result = string(data)

		datass := parseAuthBukalapak(data)
		var objToken models.TokenBukalapak

		if datass.AccessToken != "" {
			objToken.AccessToken = datass.AccessToken
			objToken.TokenType = datass.TokenType
			objToken.CreatedAt = datass.CreatedAt
			objToken.ExpiresIn = datass.ExpiresIn
			objToken.RefreshToken = datass.RefreshToken
			objToken.Scope = datass.Scope

			tokenService.SaveTokenBukalapak(objToken, ChannelName)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			response.Message = "Save Token Sukses"
			response.Result = datass
		}

		//insert token

	}

	c.JSON(http.StatusOK, response)
	return
}

func RefreshToken(c *gin.Context) {
	nmRoute := "RefreshToken Bukalapak"

	client_secret := os.Getenv("CLIENT_SECRET_BUKALAPAK")
	client_id := os.Getenv("CLIENT_ID_BUKALAPAK")

	ObjToken := tokenService.FindToken("bukalapak")
	refreshtoken := fmt.Sprintf("%v", ObjToken.Value2)

	var response response.ResponseCrud

	url := os.Getenv("URL_AUTH_BUKALAPAK") + "/oauth/token"

	var jsonString = `{"grant_type":"refresh_token"
	,"client_id":"` + client_id + `",
	"client_secret":"` + client_secret + `",
	"refresh_token":"` + refreshtoken + `"}`

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
		response.Result = string(data)

		datass := parseAuthBukalapak(data)
		var objToken models.TokenBukalapak

		if datass.AccessToken != "" {
			objToken.AccessToken = datass.AccessToken
			objToken.TokenType = datass.TokenType
			objToken.CreatedAt = datass.CreatedAt
			objToken.ExpiresIn = datass.ExpiresIn
			objToken.RefreshToken = datass.RefreshToken
			objToken.Scope = datass.Scope

			tokenService.SaveTokenBukalapak(objToken, ChannelName)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			response.Message = "Save Token Sukses"
			response.Result = datass
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func RefreshTokenAuto() {
	nmRoute := "RefreshTokenAuto Bukalapak"

	client_secret := os.Getenv("CLIENT_SECRET_BUKALAPAK")
	client_id := os.Getenv("CLIENT_ID_BUKALAPAK")

	ObjToken := tokenService.FindToken("bukalapak")
	refreshtoken := fmt.Sprintf("%v", ObjToken.Value2)

	var response response.ResponseCrud

	url := os.Getenv("URL_AUTH_BUKALAPAK") + "/oauth/token"

	var jsonString = `{"grant_type":"refresh_token"
	,"client_id":"` + client_id + `",
	"client_secret":"` + client_secret + `",
	"refresh_token":"` + refreshtoken + `"}`

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
		response.Result = string(data)

		datass := parseAuthBukalapak(data)
		var objToken models.TokenBukalapak

		if datass.AccessToken != "" {
			objToken.AccessToken = datass.AccessToken
			objToken.TokenType = datass.TokenType
			objToken.CreatedAt = datass.CreatedAt
			objToken.ExpiresIn = datass.ExpiresIn
			objToken.RefreshToken = datass.RefreshToken
			objToken.Scope = datass.Scope

			tokenService.SaveTokenBukalapak(objToken, ChannelName)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			response.Message = "Save Token Sukses"
			response.Result = datass
			fmt.Println(datass)
		}

	}

}

func parseProductsBukalapak(jsonBuffer []byte) models.ProductsBukalapak {

	ProductsBukalapak := models.ProductsBukalapak{}

	err := json.Unmarshal(jsonBuffer, &ProductsBukalapak)
	if err != nil {
		return ProductsBukalapak
	}

	// the array is now filled with users
	return ProductsBukalapak

}

func GetProductsAuto(wg *sync.WaitGroup) {
	nmRoute := "GetProducts Bukalapak"
	fmt.Println("mulai product bukalapak " + time.Now().String())
	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/stores/me/products"

	urlBukalapak += "?limit=" + os.Getenv("PRODUCT_LIMIT_BUKALAPAK")

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseProductsBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		if len(datas.Data) > 0 {

			for _, val := range datas.Data {
				//fmt.Println("=================================")
				//fmt.Println("SkuName " + val.SkuName + " | Name " + val.Name + " | SkuID " + strconv.Itoa(int(val.SkuID)))
				GetProductDetailLoop(val.ID, val.SkuName)
				if len(val.Variants) > 0 {
					//fmt.Println("Ada variant")
					for _, valVar := range val.Variants {
						GetProductDetailBySkuLoop(val.ID, strconv.Itoa(int(valVar.ID)), val.Name)
					}
				}
				//fmt.Println("=================================")
			}

			fmt.Println("masuk loop")
			offset := os.Getenv("PRODUCT_LIMIT_BUKALAPAK")
			intoffset, _ := strconv.Atoi(offset)
			GetProductsLoop(intoffset)
		} else {

			if len(datasErr.Errors) > 0 {
				response.Result = datasErr
				response.Message = PesanError(datasErr)
				tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
			}

		}

	}

	fmt.Println("selesai product bukalapak " + time.Now().String())
	wg.Done()

}

func GetProducts(c *gin.Context) {
	nmRoute := "GetProducts Bukalapak"

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/stores/me/products"

	urlBukalapak += "?limit=" + os.Getenv("PRODUCT_LIMIT_BUKALAPAK")

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseProductsBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		if len(datas.Data) > 0 {

			for _, val := range datas.Data {
				//fmt.Println("=================================")
				//fmt.Println("SkuName " + val.SkuName + " | Name " + val.Name + " | SkuID " + strconv.Itoa(int(val.SkuID)))
				GetProductDetailLoop(val.ID, val.SkuName)
				if len(val.Variants) > 0 {
					//fmt.Println("Ada variant")
					for _, valVar := range val.Variants {
						GetProductDetailBySkuLoop(val.ID, strconv.Itoa(int(valVar.ID)), val.Name)
					}
				}
				//fmt.Println("=================================")
			}

			fmt.Println("masuk loop")
			offset := os.Getenv("PRODUCT_LIMIT_BUKALAPAK")
			intoffset, _ := strconv.Atoi(offset)
			GetProductsLoop(intoffset)
		} else {

			if len(datasErr.Errors) > 0 {
				response.Result = string(data)
				response.Message = PesanError(datasErr)
				tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
			}

		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsV2(c *gin.Context) {
	nmRoute := "GetProductsV2 Bukalapak"
	storeid := c.Param("storeid")
	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)
	fmt.Println("token bukalapak " + token)
	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/products"

	urlBukalapak += "?limit=" + os.Getenv("PRODUCT_LIMIT_BUKALAPAK")

	if storeid != "a" {
		urlBukalapak += "&store_id=" + storeid
	}

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlBukalapak
		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}
func GetProductsLoop(offset int) {
	nmRoute := "GetProductsLoop Bukalapak"

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/stores/me/products"

	urlBukalapak += "?limit=" + os.Getenv("PRODUCT_LIMIT_BUKALAPAK")
	urlBukalapak += "&offset=" + strconv.Itoa(offset)

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseProductsBukalapak(data)
		datasErr := parseErrorBukalapak(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlBukalapak
		response.Result = datas

		if len(datas.Data) > 0 {
			for _, val := range datas.Data {
				//fmt.Println("=================================")
				//fmt.Println("SkuName " + val.SkuName + " | Name " + val.Name + " | SkuID " + strconv.Itoa(int(val.SkuID)))
				GetProductDetailLoop(val.ID, val.SkuName)
				if len(val.Variants) > 0 {
					//fmt.Println("Ada variant")
					for _, valVar := range val.Variants {
						GetProductDetailBySkuLoop(val.ID, strconv.Itoa(int(valVar.ID)), val.Name)
					}
				}
				//fmt.Println("=================================")

			}

			offset += offset
			fmt.Println("masuk loop GetProductsLoop Bukalapak offset:" + strconv.Itoa(offset))
			GetProductsLoop(offset)

		} else {
			if len(datasErr.Errors) > 0 {
				response.Result = datasErr
				response.Message = PesanError(datasErr)
				tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
			}

		}

	}
}

func parseProductDetailBukalapak(jsonBuffer []byte) models.ProductDetailBukalapak {

	ProductDetailBukalapak := models.ProductDetailBukalapak{}

	err := json.Unmarshal(jsonBuffer, &ProductDetailBukalapak)
	if err != nil {
		return ProductDetailBukalapak
	}

	// the array is now filled with users
	return ProductDetailBukalapak

}

func GetProductDetail(c *gin.Context) {
	nmRoute := "GetProductDetail Bukalapak"
	ProductId := c.Param("id")
	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/products/" + ProductId

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseProductDetailBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		if datas.Data.Name == "" {
			if len(datasErr.Errors) > 0 {
				response.Message = PesanError(datasErr)
				response.Result = datasErr
				tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
			}

		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductDetailLoop(ProductId string, SkuBl string) {
	fmt.Println("ProductId " + ProductId)
	nmRoute := "GetProductDetailLoop Bukalapak"
	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/products/" + ProductId

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseProductDetailBukalapak(data)
		datasErr := parseErrorBukalapak(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		if datas.Data.Name == "" {
			if len(datasErr.Errors) > 0 {
				response.Message = PesanError(datasErr)
				response.Result = datasErr
				tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
			}

		} else {
			tokenService.SaveSkuMappingBukalapak(datas, SkuBl)
			//models.ProductDetailBukalapak
		}

	}

}

func parseProductDetailSkuBukalapak(jsonBuffer []byte) models.ProductDetailSkuBukalapak {

	ProductDetailSkuBukalapak := models.ProductDetailSkuBukalapak{}

	err := json.Unmarshal(jsonBuffer, &ProductDetailSkuBukalapak)
	if err != nil {
		return ProductDetailSkuBukalapak
	}

	// the array is now filled with users
	return ProductDetailSkuBukalapak

}

func GetProductDetailBySku(c *gin.Context) {
	nmRoute := "GetProductDetailBySku Bukalapak"
	ProductId := c.Param("id")
	SkuId := c.Param("skuid")
	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/products/" + ProductId + "/skus/" + SkuId

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseProductDetailSkuBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		if datas.Data.ProductID == "" {

			if len(datasErr.Errors) > 0 {
				response.Message = PesanError(datasErr)
				response.Result = datasErr
				tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
			}

		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProductDetailBySkuLoop(ProductId string, SkuId string, ProductName string) {
	nmRoute := "GetProductDetailBySkuLoop Bukalapak"

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/products/" + ProductId + "/skus/" + SkuId

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseProductDetailSkuBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		if datas.Data.ProductID == "" {
			if len(datasErr.Errors) > 0 {
				response.Message = PesanError(datasErr)
				response.Result = datasErr
				tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
			}

		} else {
			//Skunya := datas.Data.SkuName
			//Variant := datas.Data.VariantName
			// fmt.Println("ada detail variant")
			// fmt.Println("variant SKU " + Skunya + " variant " + Variant)

			//fmt.Println("variant SKU " + Skunya + " | varian " + Variant + " | SkuID " + strconv.Itoa(int(datas.Data.ID)))
			tokenService.SaveSkuMappingVarianBukalapak(datas, ProductName)

		}

	}

}

func parseUpdateStockBukalapak(jsonBuffer []byte) models.UpdateStockBukalapak {

	UpdateStockBukalapak := models.UpdateStockBukalapak{}

	err := json.Unmarshal(jsonBuffer, &UpdateStockBukalapak)
	if err != nil {
		return UpdateStockBukalapak
	}

	// the array is now filled with users
	return UpdateStockBukalapak

}

func UpdateStock(c *gin.Context) {
	nmRoute := "UpdateStock Bukalapak"
	ProductId := c.Param("id")
	SkuId := c.Param("skuid")
	Stock := c.Param("stock")
	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/products/" + ProductId + "/skus/" + SkuId

	var jsonString = `{
		"stock": ` + Stock + `
	  }`

	var body_url = []byte(jsonString)

	req, err := http.NewRequest("PATCH", urlBukalapak, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseUpdateStockBukalapak(data)
		datasErr := parseErrorBukalapak(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		if len(datasErr.Errors) > 0 {
			response.Result = datasErr
			response.Message = PesanError(datasErr)
			tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func parseOrdersBukalapak(jsonBuffer []byte) models.OrdersBukalapak {

	OrdersBukalapak := models.OrdersBukalapak{}

	err := json.Unmarshal(jsonBuffer, &OrdersBukalapak)
	if err != nil {
		return OrdersBukalapak
	}

	// the array is now filled with users
	return OrdersBukalapak

}

func GetOrders(c *gin.Context) {
	nmRoute := "GetOrders Bukalapak"
	Status := c.Param("status")

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/transactions"

	skrg := time.Now().Format("2006-01-02T15:04:05-0700")
	kmrn := time.Now().Add(time.Duration(-72) * time.Hour).Format("2006-01-02T15:04:05-0700")

	// urlBukalapak += "?states=" + Status
	// urlBukalapak += "&limit=" + os.Getenv("ORDER_LIMIT_BUKALAPAK")

	urlBukalapak += "?limit=" + os.Getenv("ORDER_LIMIT_BUKALAPAK")
	urlBukalapak += "&offset=0"
	urlBukalapak += "&states=" + Status
	urlBukalapak += "&start_time=" + kmrn
	urlBukalapak += "&end_time=" + skrg

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseOrdersBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlBukalapak
		response.Result = datas

		if len(datas.Data) > 0 {

			for _, val := range datas.Data {
				GetOrderDetailLoop(strconv.Itoa(int(val.ID)))
			}

			limit, _ := strconv.Atoi(os.Getenv("ORDER_LIMIT_BUKALAPAK"))
			offset := limit
			GetOrderLoop(Status, skrg, kmrn, limit, offset)
		}

		if len(datasErr.Errors) > 0 {
			response.Message = PesanError(datasErr)
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrdersAuto(Status string) {
	nmRoute := "GetOrders Bukalapak"

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/transactions"

	skrg := time.Now().Format("2006-01-02T15:04:05-0700")
	kmrn := time.Now().Add(time.Duration(-72) * time.Hour).Format("2006-01-02T15:04:05-0700")

	// urlBukalapak += "?states=" + Status
	// urlBukalapak += "&limit=" + os.Getenv("ORDER_LIMIT_BUKALAPAK")

	urlBukalapak += "?limit=" + os.Getenv("ORDER_LIMIT_BUKALAPAK")
	urlBukalapak += "&offset=0"
	urlBukalapak += "&states=" + Status
	urlBukalapak += "&start_time=" + kmrn
	urlBukalapak += "&end_time=" + skrg

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseOrdersBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlBukalapak
		response.Result = datas

		if len(datas.Data) > 0 {

			for _, val := range datas.Data {
				GetOrderDetailLoop(strconv.Itoa(int(val.ID)))
			}

			limit, _ := strconv.Atoi(os.Getenv("ORDER_LIMIT_BUKALAPAK"))
			offset := limit
			GetOrderLoop(Status, skrg, kmrn, limit, offset)
		}

		if len(datasErr.Errors) > 0 {
			response.Message = PesanError(datasErr)
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
		}

	}

}

func GetOrderLoop(Status string, skrg string, kmrn string, limit int, offset int) {
	nmRoute := "GetOrderLoop Bukalapak"

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/transactions"

	urlBukalapak += "?states=" + Status
	urlBukalapak += "&limit=" + strconv.Itoa(limit)
	urlBukalapak += "&offset=" + strconv.Itoa(offset)
	urlBukalapak += "&start_time=" + kmrn
	urlBukalapak += "&end_time=" + skrg

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseOrdersBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)

		if len(datas.Data) > 0 {
			fmt.Println("masuk " + nmRoute)
			for _, val := range datas.Data {

				GetOrderDetailLoop(strconv.Itoa(int(val.ID)))
			}

			offset += limit
			GetOrderLoop(Status, skrg, kmrn, limit, offset)
		}

		if len(datasErr.Errors) > 0 {
			response.Message = PesanError(datasErr)
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
		}

	}

}

func parseOrderDetailBukalapak(jsonBuffer []byte) models.OrderDetailBukalapak {

	OrderDetailBukalapak := models.OrderDetailBukalapak{}

	err := json.Unmarshal(jsonBuffer, &OrderDetailBukalapak)
	if err != nil {
		return OrderDetailBukalapak
	}

	// the array is now filled with users
	return OrderDetailBukalapak

}

func GetOrderDetail(c *gin.Context) {
	nmRoute := "GetOrderDetail Bukalapak"
	OrderIds := c.Param("orderid")

	OrderId := OrderIds[2:len(OrderIds)] //potong 2 depan

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/transactions/" + OrderId

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseOrderDetailBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		if datas.Data.State != "" {
			response.ResponseDesc = datas.Data.State
		}

		if len(datasErr.Errors) > 0 {
			response.Message = PesanError(datasErr)
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrderDetailLoop(orderId string) {
	nmRoute := "GetOrderDetailLoop Bukalapak"
	fmt.Println("GetOrderDetailLoop " + orderId)
	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/transactions/" + orderId

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseOrderDetailBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		if len(datas.Data.Items) > 0 {
			tokenService.SaveSalesOrderBukalapak(datas)
		}

		if len(datasErr.Errors) > 0 {
			response.Message = PesanError(datasErr)
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
		}

	}

}

func UpdateStatusOrder(c *gin.Context) {
	nmRoute := "UpdateStatusOrder Bukalapak"
	OrderIds := c.Param("orderid")
	Status := c.Param("status")

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	OrderId := OrderIds[2:len(OrderIds)] //potong 2 depan

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/transactions/" + OrderId + "/status"
	fmt.Println(urlBukalapak)
	jsonString := ""
	jsonString = `{
		"state": "` + Status + `"
	  }`

	if Status == "delivered" {

		param_kurir := ""
		objCek, _ := tokenRepository.FindSalesOrder(OrderIds)
		if objCek.ExpeditionType != "" {
			param_kurir = objCek.ExpeditionType
		}

		jsonString = `{
					"state": "` + Status + `",
					"state_options": {
					"carrier": "` + param_kurir + `"
					}
					}`
	}

	//fmt.Println(jsonString)
	var body_url = []byte(jsonString)

	req, err := http.NewRequest("PUT", urlBukalapak, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		response.Message = "GAGAL"
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseOrderDetailBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = string(data)

		if datas.Data.TransactionID == "" {
			response.Message = "ERROR "
		}
		if len(datasErr.Errors) > 0 {
			response.Message = PesanError(datasErr)
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func parseBookingBukalapak(jsonBuffer []byte) models.BookingBukalapak {

	BookingBukalapak := models.BookingBukalapak{}

	err := json.Unmarshal(jsonBuffer, &BookingBukalapak)
	if err != nil {
		return BookingBukalapak
	}

	// the array is now filled with users
	return BookingBukalapak

}

func ReqPicking(c *gin.Context) {
	nmRoute := "ReqPicking Bukalapak"
	OrderIds := c.Param("orderid")
	TypeShip := c.Param("type")

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	OrderId := OrderIds[2:len(OrderIds)] //potong 2 depan

	param_kurir := ""
	objCek, _ := tokenRepository.FindSalesOrder(OrderIds)
	if objCek.ExpeditionType != "" {
		param_kurir = objCek.ExpeditionType
		//param_kurir = "sicepat"
	}

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/_partners/logistic-bookings"

	if TypeShip != "pickup" {
		TypeShip = "drop off"
	}

	//cari kurir group
	objKurir := GetDetailKurirLoop()

	if len(objKurir.Data) > 0 {

		for _, val := range objKurir.Data {
			if param_kurir == val.Carrier {
				param_kurir = val.CourierGroup
				break
			}

		}

	}

	jsonString := `{
		"transaction_id": ` + OrderId + `,
		"courier_selection": "` + param_kurir + `"
	  }`
	// fmt.Println(TypeShip)
	// fmt.Println(urlBukalapak)
	//fmt.Println(jsonString)
	var body_url = []byte(jsonString)

	req, err := http.NewRequest("POST", urlBukalapak, bytes.NewBuffer(body_url))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		response.Message = "GAGAL"
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseBookingBukalapak(data)
		datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		if datas.Data.BookingCode != "" {
			//update jadi delivered

		}

		if datas.Data.Courier == "" {
			response.Message = "ERROR "
			response.Result = string(data)
		}
		if len(datasErr.Errors) > 0 {
			response.Message = PesanError(datasErr)
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, PesanError(datasErr), nmRoute)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrderResi(c *gin.Context) {
	nmRoute := "GetOrderResi Bukalapak"
	OrderIds := c.Param("orderid")

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	OrderId := OrderIds[2:len(OrderIds)] //potong 2 depan

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/transactions/" + OrderId

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		response.Message = "GAGAL"
		tokenService.SaveErrorString(ChannelName, "Gagal Konek", nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseOrderDetailBukalapak(data)
		//datasErr := parseErrorBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = ""
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		response.ResponseDesc = datas.Data.Delivery.TrackingNumber

	}

	c.JSON(http.StatusOK, response)
	return
}

func parseDetailStoreBukalapak(jsonBuffer []byte) models.DetailStoreBukalapak {

	DetailStoreBukalapak := models.DetailStoreBukalapak{}

	err := json.Unmarshal(jsonBuffer, &DetailStoreBukalapak)
	if err != nil {
		return DetailStoreBukalapak
	}

	// the array is now filled with users
	return DetailStoreBukalapak

}
func DetailStore(c *gin.Context) {
	nmRoute := "DetailStore Bukalapak"

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/me"

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseDetailStoreBukalapak(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlBukalapak
		response.Result = datas

	}

	c.JSON(http.StatusOK, response)
	return
}

func PesanError(obj models.ErrorBukalapak) string {
	error := "Error Bukalapak"
	if len(obj.Errors) > 0 {
		error = obj.Errors[0].Message
	}
	return error
}

func parseDetailKurirBukalapak(jsonBuffer []byte) models.DetailKurirBukalapak {

	DetailKurirBukalapak := models.DetailKurirBukalapak{}

	err := json.Unmarshal(jsonBuffer, &DetailKurirBukalapak)
	if err != nil {
		return DetailKurirBukalapak
	}

	// the array is now filled with users
	return DetailKurirBukalapak

}

func GetDetailKurirLoop() models.DetailKurirBukalapak {
	nmRoute := "GetDetailKurirLoop Bukalapak"
	var obj models.DetailKurirBukalapak
	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/info/carriers"

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseDetailKurirBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Result = datas
		obj = datas

	}

	return obj
}

func GetDetailKurir(c *gin.Context) {
	nmRoute := "GetDetailKurir Bukalapak"

	ObjToken := tokenService.FindToken("bukalapak")
	token := fmt.Sprintf("%v", ObjToken.Value1)

	var response response.ResponseCrud

	urlBukalapak := os.Getenv("URL_API_BUKALAPAK") + "/info/carriers"

	req, err := http.NewRequest("GET", urlBukalapak, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
		datas := parseDetailKurirBukalapak(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlBukalapak
		response.Result = datas

	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateAllStockBukalapak(c *gin.Context) {
	var response response.ResponseCrud

	ObjMapping := tokenService.CariSkuMappingObjGroup(os.Getenv("KODE_BUKALAPK")) //param by channel name

	if len(ObjMapping) > 0 {
		for _, value := range ObjMapping {
			helpers.UpdateStock(value.SkuNo, "API_CHANNEL", os.Getenv("KODE_BUKALAPK"))

		}
	}

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "DISINISS"

	c.JSON(http.StatusOK, response)
	return
}
