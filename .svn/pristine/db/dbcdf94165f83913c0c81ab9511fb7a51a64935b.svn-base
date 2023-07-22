package jdidController

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	urls "net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/helpers"
	"github.com/rals/dearme-channel/models"
	"github.com/rals/dearme-channel/models/response"
	"github.com/rals/dearme-channel/repositories/tokenRepository"
	"github.com/rals/dearme-channel/services/tokenService"
	"github.com/rals/dearme-channel/utils"
	// "encoding/json"
)

var ChannelName = "jdid"

func ReqToken(c *gin.Context) {

	url_client := os.Getenv("URI_CALLBACK_JDID") + "getCode"
	client_id := os.Getenv("CLIENT_ID_JDID")
	scope := "snsapi_base"
	response_type := "code"

	var response response.ResponseCrud

	url := os.Getenv("URL_AUTH_JDID") + "oauth2/to_login"

	// oauth.jd.id/oauth2/to_login?app_key=xxxx&response_type=code&redirect_uri=xxxx&state=20200428&scope=snsapi_base

	tambah := "?app_key=" + client_id
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
		//data, _ := ioutil.ReadAll(resp.Body)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		//response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func parseAuthJdid(jsonBuffer []byte) models.TokenJdId {

	TokenJdId := models.TokenJdId{}

	err := json.Unmarshal(jsonBuffer, &TokenJdId)
	if err != nil {
		return TokenJdId
	}

	// the array is now filled with users
	return TokenJdId

}

func GetToken(c *gin.Context) {
	nmRoute := "GetToken"
	var response response.ResponseCrud

	code := c.Query("code")

	//https://oauth.jd.id/oauth2/access_token?app_key=xxxx&app_secret=xxxxxx&grant_type=authorization_code&code=xxxx

	url_client := os.Getenv("URI_CALLBACK_JDID") + "saveToken"
	client_secret := os.Getenv("CLIENT_SECRET_JDID")
	client_id := os.Getenv("CLIENT_ID_JDID")

	url := os.Getenv("URL_AUTH_JDID") + "oauth2/access_token"

	tambah := "?app_key=" + client_id
	tambah += "&redirect_uri=" + url_client
	tambah += "&app_secret=" + client_secret
	tambah += "&code=" + code
	tambah += "&grant_type=authorization_code"

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
		log.Println("Error on response.\n[ERRO] -", err)
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
		fmt.Println(string(data))
		//fmt.Println("sukses")
		datass := parseAuthJdid(data)
		var objToken models.TokenJdId

		if datass.AccessToken != "" {
			objToken.AccessToken = datass.AccessToken
			objToken.TokenType = datass.TokenType
			objToken.CreatedAt = datass.CreatedAt
			objToken.ExpiresIn = datass.ExpiresIn
			objToken.RefreshToken = datass.RefreshToken
			objToken.Scope = datass.Scope

			tokenService.SaveTokenJdId(objToken, ChannelName)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			response.Message = "Save Token Sukses"
			response.Result = datass
		}

		// //insert token

	}

	c.JSON(http.StatusOK, response)
	return
}

func RefreshToken(c *gin.Context) {
	nmRoute := "RefreshToken"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")
	refreshtoken := fmt.Sprintf("%v", ObjToken.Value2)

	//https://oauth.jd.id/oauth2/access_token?app_key=xxxx&app_secret=xxxxxx&grant_type=authorization_code&code=xxxx

	client_secret := os.Getenv("CLIENT_SECRET_JDID")
	client_id := os.Getenv("CLIENT_ID_JDID")

	url := os.Getenv("URL_AUTH_JDID") + "oauth2/refresh_token"

	tambah := "?app_key=" + client_id
	tambah += "&app_secret=" + client_secret
	tambah += "&refresh_token=" + refreshtoken
	tambah += "&grant_type=refresh_token"

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
		log.Println("Error on response.\n[ERRO] -", err)
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
		fmt.Println(string(data))
		//fmt.Println("sukses")
		datass := parseAuthJdid(data)
		var objToken models.TokenJdId

		if datass.AccessToken != "" {
			objToken.AccessToken = datass.AccessToken
			objToken.TokenType = datass.TokenType
			objToken.CreatedAt = datass.CreatedAt
			objToken.ExpiresIn = datass.ExpiresIn
			objToken.RefreshToken = datass.RefreshToken
			objToken.Scope = datass.Scope

			tokenService.SaveTokenJdId(objToken, ChannelName)
			response.ResponseCode = http.StatusOK
			response.ResponseDesc = enums.SUCCESS
			response.ResponseTime = utils.DateToStdNow()
			response.Message = "Save Token Sukses"
			response.Result = datass
		}

		// //insert token

	}

	c.JSON(http.StatusOK, response)
	return
}

func AuthJdid(path string, token string, option string) (string, string, string) {

	now := time.Now().Format("2006-01-02 15:04:05.000+0700")

	//now := "2022-06-23 11:46:12.201+0700"
	partner_key := os.Getenv("CLIENT_ID_JDID")
	partner_secret := os.Getenv("CLIENT_SECRET_JDID")

	parameter := map[string]string{
		"app_key":   partner_key,
		"method":    path,
		"timestamp": now,
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

	input = partner_secret + input + partner_secret

	balik := input
	sign := ""
	hash := md5.Sum([]byte(input))
	sign = hex.EncodeToString(hash[:])

	return strings.ToUpper(sign), now, balik

}

func parseJsonOrdersJdId(jsonBuffer []byte) models.OrdersDataJdId {

	OrdersDataJdId := models.OrdersDataJdId{}

	err := json.Unmarshal(jsonBuffer, &OrdersDataJdId)
	if err != nil {
		return OrdersDataJdId
	}

	// the array is now filled with users
	return OrdersDataJdId

}

func parseJsonStockJdId(jsonBuffer []byte) models.UpdateStockJdID {

	UpdateStockJdID := models.UpdateStockJdID{}

	err := json.Unmarshal(jsonBuffer, &UpdateStockJdID)
	if err != nil {
		return UpdateStockJdID
	}

	// the array is now filled with users
	return UpdateStockJdID

}

func parseJsonErrorStockJdId(jsonBuffer []byte) models.ErrorUpdateStockJdID {

	ErrorUpdateStockJdID := models.ErrorUpdateStockJdID{}

	err := json.Unmarshal(jsonBuffer, &ErrorUpdateStockJdID)
	if err != nil {
		return ErrorUpdateStockJdID
	}

	// the array is now filled with users
	return ErrorUpdateStockJdID

}

func parseJsonProductJdId(jsonBuffer []byte) models.HeaderProductJdId {

	HeaderProductJdId := models.HeaderProductJdId{}

	err := json.Unmarshal(jsonBuffer, &HeaderProductJdId)
	if err != nil {
		return HeaderProductJdId
	}

	// the array is now filled with users
	return HeaderProductJdId

}

func parseJsonDetailProductJdId(jsonBuffer []byte) models.DetailProductJdId {

	DetailProductJdId := models.DetailProductJdId{}

	err := json.Unmarshal(jsonBuffer, &DetailProductJdId)
	if err != nil {
		return DetailProductJdId
	}

	// the array is now filled with users
	return DetailProductJdId

}

func parseJsonDetailStockJdId(jsonBuffer []byte) models.DetailStockJdId {

	DetailStockJdId := models.DetailStockJdId{}

	err := json.Unmarshal(jsonBuffer, &DetailStockJdId)
	if err != nil {
		return DetailStockJdId
	}

	// the array is now filled with users
	return DetailStockJdId

}

func parseJsonDetailOrderJdID(jsonBuffer []byte) models.OrderDetailJdId {

	OrderDetailJdId := models.OrderDetailJdId{}

	err := json.Unmarshal(jsonBuffer, &OrderDetailJdId)
	if err != nil {
		return OrderDetailJdId
	}

	// the array is now filled with users
	return OrderDetailJdId

}

func parseJsonPickingJdID(jsonBuffer []byte) models.ReqPickJdId {

	ReqPickJdId := models.ReqPickJdId{}

	err := json.Unmarshal(jsonBuffer, &ReqPickJdId)
	if err != nil {
		return ReqPickJdId
	}

	// the array is now filled with users
	return ReqPickJdId

}

func parseJsonErrorJdID(jsonBuffer []byte) models.ErrorJdId {

	ErrorJdId := models.ErrorJdId{}

	err := json.Unmarshal(jsonBuffer, &ErrorJdId)
	if err != nil {
		return ErrorJdId
	}

	// the array is now filled with users
	return ErrorJdId

}

func parseJsonPrintDocJdID(jsonBuffer []byte) models.PrintDocShipJdId {

	PrintDocShipJdId := models.PrintDocShipJdId{}

	err := json.Unmarshal(jsonBuffer, &PrintDocShipJdId)
	if err != nil {
		return PrintDocShipJdId
	}

	// the array is now filled with users
	return PrintDocShipJdId

}

func GetProducts(c *gin.Context) {
	nmRoute := "GetProducts"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	fmt.Println("======= GET Products JDID =======")

	Size := os.Getenv("LIMIT_PRODUCT_JDID")
	bodynya := `{"page":"1","size":"` + Size + `"}`
	fmt.Println(bodynya)
	tambahan := ""
	versi := "2.0"
	methode := "jingdong.seller.product.getWareInfoListByVendorId"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		datas := parseJsonProductJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		// response.Result = string(data)
		response.Result = datas

		index := 0
		fmt.Println(string(data))
		CekErr := datas.JingdongSellerProductGetWareInfoListByVendorIDResponse.ReturnType.IsSuccess
		if CekErr == false {
			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)
		} else {
			CountobjProduct := datas.JingdongSellerProductGetWareInfoListByVendorIDResponse.ReturnType.Model.SpuInfoVoList
			//TotalData := datas.JingdongSellerProductGetWareInfoListByVendorIDResponse.ReturnType.Model.TotalNum
			//cek detail spuID
			for _, val := range CountobjProduct {

				// fmt.Println("SPUID " + strconv.Itoa(val.SpuID))
				// fmt.Println("STATUS " + strconv.Itoa(val.WareStatus))
				GetProductByIDLoop(val.SpuID, val.WareStatus)
				// index++
			}

			//count CountobjProductll
			if len(CountobjProduct) > 0 {
				index++
				fmt.Println("======= CARI LAGI =======")
				fmt.Println("index : " + strconv.Itoa(index) + " Size " + Size)
				GetProductsLoop(index, Size)
			}
		}

	}

	fmt.Println("======= SELESAI Products JDID =======")

	c.JSON(http.StatusOK, response)
	return
}

func GetProductsAuto(wg *sync.WaitGroup) {
	nmRoute := "GetProductsAuto"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	fmt.Println("mulai product jdid " + time.Now().String())

	Size := os.Getenv("LIMIT_PRODUCT_JDID")
	bodynya := `{"page":"1","size":"` + Size + `"}`

	tambahan := ""
	versi := "2.0"
	methode := "jingdong.seller.product.getWareInfoListByVendorId"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		datas := parseJsonProductJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		// response.Result = string(data)
		response.Result = datas

		index := 0

		CekErr := datas.JingdongSellerProductGetWareInfoListByVendorIDResponse.ReturnType.IsSuccess
		if CekErr == false {
			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)
		} else {
			CountobjProduct := datas.JingdongSellerProductGetWareInfoListByVendorIDResponse.ReturnType.Model.SpuInfoVoList
			//TotalData := datas.JingdongSellerProductGetWareInfoListByVendorIDResponse.ReturnType.Model.TotalNum
			//cek detail spuID
			for _, val := range CountobjProduct {

				// fmt.Println("SPUID " + strconv.Itoa(val.SpuID))
				// fmt.Println("STATUS " + strconv.Itoa(val.WareStatus))
				GetProductByIDLoop(val.SpuID, val.WareStatus)
				//index++
			}

			//count CountobjProductll
			if len(CountobjProduct) > 0 {
				index++
				//fmt.Println("======= CARI LAGI =======")
				GetProductsLoop(index, Size)
			}
		}

	}

	fmt.Println("selesai product jdid " + time.Now().String())
	wg.Done()
}

func GetProductsLoop(index int, size string) {
	nmRoute := "GetProductsLoop"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)

	bodynya := `{"page":"` + strconv.Itoa(index) + `","size":"` + size + `"}`

	tambahan := ""
	versi := "2.0"
	methode := "jingdong.seller.product.getWareInfoListByVendorId"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		datas := parseJsonProductJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""

		CekErr := datas.JingdongSellerProductGetWareInfoListByVendorIDResponse.ReturnType.IsSuccess
		if CekErr == false {

			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)

		} else {

			CountobjProduct := datas.JingdongSellerProductGetWareInfoListByVendorIDResponse.ReturnType.Model.SpuInfoVoList
			//TotalData := datas.JingdongSellerProductGetWareInfoListByVendorIDResponse.ReturnType.Model.TotalNum
			//cek detail spuID
			for _, val := range CountobjProduct {
				// fmt.Println("SPUID LOOP" + strconv.Itoa(val.SpuID))
				// fmt.Println("STATUS LOOP" + strconv.Itoa(val.WareStatus))
				GetProductByIDLoop(val.SpuID, val.WareStatus)
				// index++
			}

			// if index < TotalData {
			if len(CountobjProduct) > 0 {
				index++
				//fmt.Println("======= CARI LAGI LOOP=======")
				//fmt.Println("index : " + strconv.Itoa(index) + " Size " + size)
				GetProductsLoop(index, size)
			}

		}

	}

}

func GetProductByIDLoop(idnya int, statusid int) {
	nmRoute := "GetProductByID"

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)

	bodynya := `{"spuId":"` + strconv.Itoa(idnya) + `"}`

	tambahan := ""
	versi := "2.0"
	methode := "jingdong.seller.product.getSkuInfoBySpuIdAndVenderId"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		datas := parseJsonDetailProductJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas

		//fmt.Println(statusid)
		//datasnew := datas.JingdongSellerProductGetSkuInfoBySpuIDAndVenderIDResponse.ReturnType.Model
		CekErr := datas.JingdongSellerProductGetSkuInfoBySpuIDAndVenderIDResponse.ReturnType.Success
		if CekErr == false {

			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)

		} else {

			//if idnya == 678426106 || idnya == 678423640 {
			tokenService.SaveSkuMappingJdId(datas, idnya, statusid)
			//}

		}

	}

}

func GetProductByID(c *gin.Context) {
	nmRoute := "GetProductByID"
	idnya := c.Param("code")

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)

	bodynya := `{"spuId":"` + idnya + `"}`

	tambahan := ""
	versi := "2.0"
	methode := "jingdong.seller.product.getSkuInfoBySpuIdAndVenderId"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	//tambahan += "#method#" + methode
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		datas := parseJsonDetailProductJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		CekErr := datas.JingdongSellerProductGetSkuInfoBySpuIDAndVenderIDResponse.ReturnType.Success
		if CekErr == false {

			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)

		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateStockJdID(c *gin.Context) {
	nmRoute := "UpdateStockJdID"
	idnya := c.Param("skuid")
	stocknya := c.Param("stock")
	//sini
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)

	bodynya := `{"wareStockUpdateListStr":"[{'skuId':` + idnya + `,'realNum':` + stocknya + `}]"}`
	//{"wareStockUpdateListStr":"[{'skuId':678337580,'realNum':1}]"}

	tambahan := ""
	versi := "2.0"
	methode := "jingdong.epistock.updateEpiMerchantWareStock"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	//tambahan += "#method#" + methode
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		response.Message = "ERROR KONEKSI"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonStockJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		CekErr := datas.JingdongEpistockUpdateEpiMerchantWareStockResponse.EptRemoteResult.Success

		if CekErr == false {
			datasErr := parseJsonErrorStockJdId(data)
			pesanErr := datasErr.JingdongEpistockUpdateEpiMerchantWareStockResponse.EptRemoteResult.Message
			if pesanErr == "Failure operation!!" {
				//inaktif sku

			}
			if pesanErr == "" {
				pesanErr = "ERROR UPDATE STOCK"
			} else {
				pesanErr += " (JDID)"
			}

			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, idnya+"|"+pesanErr, nmRoute)
		}

		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func CekStockJdID(c *gin.Context) {
	nmRoute := "CekStockJdID"
	idnya := c.Param("skuid")

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println(idnya)
	bodynya := `{"wareStockQueryListStr":"[{'skuId':` + idnya + `}]"}`
	// {"wareStockQueryListStr":"[{'skuId':678337580}]"}

	//{"wareStockUpdateListStr":"[{'skuId':678337580,'realNum':1}]"}

	tambahan := ""
	methode := "jingdong.epistock.queryEpiMerchantWareStock"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	//tambahan += "#method#" + methode
	tambahan += "#v#2.0"

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	var param = urls.Values{}
	param.Set("app_key", os.Getenv("CLIENT_ID_JDID"))
	param.Set("360buy_param_json", bodynya)
	param.Set("access_token", token)
	param.Set("method", methode)
	param.Set("timestamp", timest)
	param.Set("v", "2.0")
	param.Set("sign", sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		datas := parseJsonDetailStockJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		//datas map[string]interface{}

		//fmt.Println(datas.JingdongEpistockQueryEpiMerchantWareStockResponse.EptRemoteResult.Model)
		response.Message = ""

		//fmt.Println(string(data))
		response.Result = datas

	}

	c.JSON(http.StatusOK, response)
	return
}

func UpdateSkuInfo(c *gin.Context) {
	nmRoute := "CekStockJdID"
	idnya := c.Param("skuid")
	codenya := c.Param("code")
	skuwms := c.Param("skuwms")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println(idnya)
	bodynya := `{"spuId":` + codenya + `,"skuId":` + idnya + `,"sellerSkuId":"` + skuwms + `"}`
	// {"wareStockQueryListStr":"[{'skuId':678337580}]"}

	//{"wareStockUpdateListStr":"[{'skuId':678337580,'realNum':1}]"}

	tambahan := ""
	methode := "jingdong.seller.product.sku.write.updateSkuInfo"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	//tambahan += "#method#" + methode
	tambahan += "#v#2.0"

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	var param = urls.Values{}
	param.Set("app_key", os.Getenv("CLIENT_ID_JDID"))
	param.Set("360buy_param_json", bodynya)
	param.Set("access_token", token)
	param.Set("method", methode)
	param.Set("timestamp", timest)
	param.Set("v", "2.0")
	param.Set("sign", sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		//datas := parseJsonDetailStockJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()

		//datas map[string]interface{}

		//fmt.Println(datas.JingdongEpistockQueryEpiMerchantWareStockResponse.EptRemoteResult.Model)
		response.Message = ""

		//fmt.Println(string(data))
		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetOrders(c *gin.Context) {
	nmRoute := "GetOrders"
	status := c.Param("status")
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println(idnya)

	skrg := time.Now().Format("2006-01-02 15:04:05")
	kmrn := time.Now().Add(time.Duration(-72) * time.Hour).Format("2006-01-02 15:04:05") //3 hari lalu
	pageSize := 50
	pageNo := 1
	// bodynya := `{"orderStatus":` + status + `,"createdTimeBegin":"` + kmrn + `", "createdTimeEnd":"` + skrg + `","pageSize": ` + strconv.Itoa(pageSize) + `,"pageNo":` + strconv.Itoa(pageNo) + ` }`
	// fmt.Println(bodynya)
	bodynya := `{"orderStatus":` + status + `,"updateTimeBegin":"` + kmrn + `", "updateTimeEnd":"` + skrg + `","pageSize": ` + strconv.Itoa(pageSize) + `,"pageNo":` + strconv.Itoa(pageNo) + ` }`

	tambahan := ""
	versi := "2.0"
	methode := "jingdong.seller.order.getOrderIdListByCondition"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	//tambahan += "#method#" + methode
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		datas := parseJsonOrdersJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		CekErr := datas.JingdongSellerOrderGetOrderIDListByConditionResponse.Result.Success
		if CekErr == false {
			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			if pesanErr == "" {
				pesanErr = "Error GetOrders"
			} else {
				pesanErr += " (JDID)"
			}

			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)
		} else {
			dataIns := datas.JingdongSellerOrderGetOrderIDListByConditionResponse.Result.Model
			statusDetail := ""
			for key, _ := range dataIns {
				statusDetail = "ada"
				objCek, _ := tokenRepository.FindSalesOrder(strconv.Itoa(dataIns[key]))
				if objCek.NoOrder == "" {
					GetOrderDetailLoop(dataIns[key])
				} else if status == "5" && objCek.StatusProcessOrder == "0" {
					GetOrderDetailLoop(dataIns[key])
				}
				//GetOrderDetailLoop(dataIns[key])
			}

			if statusDetail != "" {
				GetOrdersLoop(status, skrg, kmrn, pageSize, pageNo)
			}

		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetProcessedOrderAutoJdId(status string) {
	nmRoute := "GetProcessedOrderAutoJdId"

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println(idnya)

	skrg := time.Now().Format("2006-01-02 15:04:05")
	// kmrn := time.Now().Add(time.Duration(-72) * time.Hour).Format("2006-01-02 15:04:05") //3 hari lalu
	kmrn := time.Now().Add(time.Duration(-192) * time.Hour).Format("2006-01-02 15:04:05") //8 hari lalu
	pageSize := 50
	pageNo := 1
	//bodynya := `{"orderStatus":` + status + `,"createdTimeBegin":"` + kmrn + `", "createdTimeEnd":"` + skrg + `","pageSize": ` + strconv.Itoa(pageSize) + `,"pageNo":` + strconv.Itoa(pageNo) + ` }`
	bodynya := `{"orderStatus":` + status + `,"updateTimeBegin":"` + kmrn + `", "updateTimeEnd":"` + skrg + `","pageSize": ` + strconv.Itoa(pageSize) + `,"pageNo":` + strconv.Itoa(pageNo) + ` }`

	// fmt.Println(bodynya)
	tambahan := ""
	versi := "2.0"
	methode := "jingdong.seller.order.getOrderIdListByCondition"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	//tambahan += "#method#" + methode
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		datas := parseJsonOrdersJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		CekErr := datas.JingdongSellerOrderGetOrderIDListByConditionResponse.Result.Success
		if CekErr == false {
			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			if pesanErr == "" {
				pesanErr = "Error GetOrders"
			} else {
				pesanErr += " (JDID)"
			}

			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)
		} else {
			dataIns := datas.JingdongSellerOrderGetOrderIDListByConditionResponse.Result.Model
			statusDetail := ""
			for key, _ := range dataIns {
				statusDetail = "ada"
				objCek, _ := tokenRepository.FindSalesOrder(strconv.Itoa(dataIns[key]))
				if objCek.NoOrder == "" {
					GetOrderDetailLoop(dataIns[key])
				} else if status == "5" && objCek.StatusProcessOrder == "0" {
					GetOrderDetailLoop(dataIns[key])
				}
				//GetOrderDetailLoop(dataIns[key])
			}

			if statusDetail != "" {
				GetOrdersLoop(status, skrg, kmrn, pageSize, pageNo)
			}

		}

	}

}

func GetOrdersLoop(status string, skrg string, kmrn string, pageSize int, pageNo int) {
	nmRoute := "GetOrdersLoop"

	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	pageNo++

	//bodynya := `{"orderStatus":` + status + `,"createdTimeBegin":"` + kmrn + `", "createdTimeEnd":"` + skrg + `","pageSize": ` + strconv.Itoa(pageSize) + `,"pageNo":` + strconv.Itoa(pageNo) + ` }`
	//fmt.Println(bodynya)
	bodynya := `{"orderStatus":` + status + `,"updateTimeBegin":"` + kmrn + `", "updateTimeEnd":"` + skrg + `","pageSize": ` + strconv.Itoa(pageSize) + `,"pageNo":` + strconv.Itoa(pageNo) + ` }`

	tambahan := ""
	versi := "2.0"
	methode := "jingdong.seller.order.getOrderIdListByCondition"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	//tambahan += "#method#" + methode
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		datas := parseJsonOrdersJdId(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		CekErr := datas.JingdongSellerOrderGetOrderIDListByConditionResponse.Result.Success
		if CekErr == false {
			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			if pesanErr == "" {
				pesanErr = "Error GetOrders"
			} else {
				pesanErr += " (JDID)"
			}

			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)
		} else {
			dataIns := datas.JingdongSellerOrderGetOrderIDListByConditionResponse.Result.Model

			statusDetail := ""
			for key, _ := range dataIns {

				objCek, _ := tokenRepository.FindSalesOrder(strconv.Itoa(dataIns[key]))
				if objCek.NoOrder == "" {
					GetOrderDetailLoop(dataIns[key])
				} else if status == "5" && objCek.StatusProcessOrder == "0" {
					GetOrderDetailLoop(dataIns[key])
				}

				//GetOrderDetailLoop(dataIns[key])
				statusDetail = "ada"
			}
			if statusDetail != "" {

				GetOrdersLoop(status, skrg, kmrn, pageSize, pageNo)
			}

		}

	}

}

func GetOrderDetailLoop(idnya int) {
	nmRoute := "GetOrderDetailLoop"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println(idnya)
	bodynya := `{"orderId":` + strconv.Itoa(idnya) + `}`

	tambahan := ""
	versi := "2.0"
	methode := "jingdong.seller.order.getOrderInfoByOrderId"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	//tambahan += "#method#" + methode
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		datas := parseJsonDetailOrderJdID(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		CekErr := datas.JingdongSellerOrderGetOrderInfoByOrderIDResponse.Result.Success
		if CekErr == false {
			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			if pesanErr == "" {
				pesanErr = "Error Ambil Detail Orderan"
			} else {
				pesanErr += " (JDID)"
			}

			response.Message = pesanErr
			response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)
		} else {
			//save disini
			tokenService.SaveSalesOrderJdId(datas)

		}

	}

}

func GetOrderDetail(c *gin.Context) {
	nmRoute := "GetOrderDetail"
	var response response.ResponseCrud
	idnya := c.Param("orderid")
	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println(idnya)
	bodynya := `{"orderId":` + idnya + `}`

	versi := "2.0"
	tambahan := ""
	methode := "jingdong.seller.order.getOrderInfoByOrderId"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		response.Message = "Gagal Cek Order Detail JDID (connection)"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonDetailOrderJdID(data)
		response.ResponseCode = http.StatusOK
		// response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		// response.Result = string(data)
		response.Result = datas

		response.ResponseDesc = strconv.Itoa(datas.JingdongSellerOrderGetOrderInfoByOrderIDResponse.Result.Model.OrderState)
		//response.Result = datas
		CekErr := datas.JingdongSellerOrderGetOrderInfoByOrderIDResponse.Result.Success
		if CekErr == false {
			fmt.Println(string(data))
			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc

			if pesanErr == "" {
				pesanErr = datas.JingdongSellerOrderGetOrderInfoByOrderIDResponse.Result.Message
			}

			if pesanErr == "" {
				pesanErr = "Error Ambil Detail Orderan"
			} else {
				pesanErr += " (JDID)"
			}

			response.Message = pesanErr
			//response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)
		} else {
			tokenService.SaveSalesOrderJdId(datas)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func PrintShipping(c *gin.Context) {
	nmRoute := "PrintShipping"
	var response response.ResponseCrud
	idnya := c.Param("orderid")
	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println(idnya)
	bodynya := `{"orderId":` + idnya + `,"printType":1,"printNum":1}`

	versi := "2.0"
	tambahan := ""
	methode := "jingdong.seller.order.printOrder"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		response.Message = "Gagal Print Shipping JDID (connection)"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonPrintDocJdID(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		//response.Result = string(data)

		CekErr := datas.JingdongSellerOrderPrintOrderResponse.Result.Success
		if CekErr == false {
			fmt.Println(string(data))
			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			if pesanErr == "" {
				pesanErr = datas.JingdongSellerOrderPrintOrderResponse.Result.Message
			}

			if pesanErr == "" {
				pesanErr = "Error Print Shipping"
			} else {
				pesanErr += " (JDID)"
			}

			response.Message = pesanErr
			response.Result = string(data)
			fmt.Println(string(data))
			//response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func ReqShippingJdId(c *gin.Context) {
	nmRoute := "ReqShippingJdId"
	var response response.ResponseCrud
	idnya := c.Param("orderid")
	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println(idnya)
	bodynya := `{"orderId":` + idnya + `}`

	versi := "2.0"
	tambahan := ""
	methode := "jingdong.seller.order.sendGoodsOpenApi"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		response.Message = "Gagal Request Shipping JDID (connection)"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonPickingJdID(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		//response.Result = string(data)
		CekErr := datas.JingdongSellerOrderSendGoodsOpenAPIResponse.Result.Success
		if CekErr == false {
			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc
			if pesanErr == "" {
				pesanErr = datas.JingdongSellerOrderSendGoodsOpenAPIResponse.Result.Message
			}

			if pesanErr == "" {
				pesanErr = "Request Picking JDID GAGAL"
			} else {
				pesanErr += " (JDID)"
			}

			response.Message = pesanErr
			response.Result = string(data)
			fmt.Println(string(data))
			tokenService.SaveErrorString(ChannelName, pesanErr, nmRoute)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetResi(c *gin.Context) {
	nmRoute := "GetResi"
	var response response.ResponseCrud
	idnya := c.Param("orderid")
	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)
	//fmt.Println(idnya)
	bodynya := `{"orderId":` + idnya + `}`

	versi := "2.0"
	tambahan := ""
	methode := "jingdong.seller.order.getOrderInfoByOrderId"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	tambahan += "#v#" + versi

	sign, timest, _ := AuthJdid(methode, token, tambahan)

	param := ReqCommmon(token, methode, bodynya, timest, versi, sign)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
		response.Message = "Gagal Cek Resi JDID (connection)"
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		datas := parseJsonDetailOrderJdID(data)
		response.ResponseCode = http.StatusOK
		// response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = ""
		response.Result = datas
		//response.Result = datas
		CekErr := datas.JingdongSellerOrderGetOrderInfoByOrderIDResponse.Result.Success
		response.ResponseDesc = ""

		resinya := datas.JingdongSellerOrderGetOrderInfoByOrderIDResponse.Result.Model.ExpressNo
		kurirnya := datas.JingdongSellerOrderGetOrderInfoByOrderIDResponse.Result.Model.CarrierCompany
		if resinya != "" && kurirnya != "" {
			response.ResponseDesc = resinya + "^" + kurirnya
		}

		if CekErr == false {
			fmt.Println(string(data))
			datasErr := parseJsonErrorJdID(data)
			pesanErr := datasErr.ErrorResponse.EnDesc

			if pesanErr == "" {
				pesanErr = datas.JingdongSellerOrderGetOrderInfoByOrderIDResponse.Result.Message
			}

			if pesanErr == "" {
				pesanErr = "Error Ambil Resi"
			} else {
				pesanErr += " (JDID)"
			}

			response.Message = pesanErr
			//response.Result = datasErr
			tokenService.SaveErrorString(ChannelName, datasErr.ErrorResponse.EnDesc, nmRoute)
		}

	}

	c.JSON(http.StatusOK, response)
	return
}

func GetSellerInfo(c *gin.Context) {
	nmRoute := "GetSellerInfo"
	var response response.ResponseCrud

	ObjToken := tokenService.FindToken("jdid")

	url := os.Getenv("URL_API_JDID")

	token := fmt.Sprintf("%v", ObjToken.Value1)

	bodynya := `{}`

	//var body_url = []byte(bodynya)

	tambahan := ""
	methode := "jingdong.seller.getAccountByPin"
	tambahan += `360buy_param_json#` + bodynya
	tambahan += "#access_token#" + token
	tambahan += "#method#" + methode
	//tambahan += "#v#2.0"

	sign, timest, raw := AuthJdid("", token, tambahan)

	//url += "?v=2.0"
	url += "?method=" + methode
	url += "&app_key=" + os.Getenv("CLIENT_ID_JDID")
	url += "&access_token=" + token
	url += "&360buy_param_json=" + bodynya
	url += "&timestamp=" + timest
	url += "&sign=" + sign

	//url = `https://open-api.jd.id/routerjson?v=2.0&method=jingdong.seller.product.getWareInfoListByVendorId&app_key=DF4C6F9506BEA85D4E5318E03F772850&access_token=b3e0f220fdeb490aaafcd6c5281dbe75nlmg&360buy_param_json={"page":"1","size":"100"}&timestamp=2022-06-22 15:12:56.074+0700&sign=859335C6545C28321BB510A05AEE489D`
	fmt.Println(url)
	// response.ResponseCode = http.StatusOK
	// response.ResponseDesc = enums.SUCCESS
	// response.ResponseTime = utils.DateToStdNow()
	response.Result = raw
	// response.Message = url

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
		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Result = err
		tokenService.SaveErrorString(ChannelName, "GAGAL GET "+nmRoute, nmRoute)

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := utils.GetByteToInterface(data)
		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = url
		fmt.Println("=================")
		fmt.Println(data)
		if string(data) != "" {
			response.Result = string(data)
		}
		response.Result = string(data)

	}

	c.JSON(http.StatusOK, response)
	return
}

func ReqCommmon(token string, methode string, bodynya string, timest string, versi string, sign string) urls.Values {
	var params = urls.Values{}
	params.Set("app_key", os.Getenv("CLIENT_ID_JDID"))
	params.Set("360buy_param_json", bodynya)
	params.Set("access_token", token)
	params.Set("method", methode)
	params.Set("timestamp", timest)
	params.Set("v", versi)
	params.Set("sign", sign)

	return params
}

func UpdateAllStockJDID(c *gin.Context) {
	var response response.ResponseCrud

	ObjMapping := tokenService.CariSkuMappingObjGroup(os.Getenv("KODE_JDID")) //param by channel name

	if len(ObjMapping) > 0 {
		for _, value := range ObjMapping {
			helpers.UpdateStock(value.SkuNo, "API_CHANNEL", os.Getenv("KODE_JDID"))

		}
	}

	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = "DISINISS"

	c.JSON(http.StatusOK, response)
	return
}
