package testController

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/models/response"
	"github.com/rals/dearme-channel/utils"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	// "encoding/json"
)

type TestBody struct {
	Name    string      `json:"name"`
	Age     interface{} `json:"age"`
	Email   []string    `json:"email"`
	Address AddJmbn     `json:"address"`
}

// type Body struct {
// 	Code   string
// 	Shop   string
// 	Partner   string
// }

type AddJmbn struct {
	Alamat1       string      `json:"alamat1"`
	Alamat2       string      `json:"alamat2"`
	Alamat3       string      `json:"alamat3"`
	Addressnumber interface{} `json:"addressnumber"`
	Addresscode   interface{} `json:"addresscode"`
	ViewTest      []ViewTest  `json:"viewtest"`
}

type ViewTest struct {
	StoreCode     string `json:"store_code"`
	TransactionNo string `json:"transaction_no"`
	TransferOut   string `json:"transfer_out"`
	ReceivingNo   string `json:"receiving_no"`
}

func BodyJson(c *gin.Context) {
	var response response.ResponseCrud

	var json TestBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatus(400)
		return
	}

	obj := json.Address.ViewTest
	for _, objDetail := range obj {
		// spew.Dump(objDetail)
		fmt.Printf("%+v\n", objDetail)
	}

	response.ResponseCode = http.StatusInternalServerError
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	response.Message = enums.MSG_SUCCESS_INSERT
	response.Result = json

	c.JSON(http.StatusOK, response)
	return
}

func GetDataRestFullORI(c *gin.Context) {
	var response response.ResponseCrud

	//test auth shopee
	timest := strconv.FormatInt(time.Now().Unix(), 10)
	host := "https://partner.test-stable.shopeemobile.com"
	paths := "/api/v2/shop/auth_partner"
	redirectUrl := "https://wms.ramayana.co.id/"
	partnerId := strconv.Itoa(1003569)
	partnerKey := "8f0b566912e0b7262489d2b7240949fe50226bc5821b3c208a2e8046620fd05b"
	baseString := fmt.Sprintf("%s%s%s", partnerId, paths, timest)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&sign=%s&redirect=%s", partnerId, timest, sign, redirectUrl)
	// urlshopee := host + paths + "?partner_id=" + partnerId + "&timestamp=" + timest + "&sign=" + sign + "&redirect=" + redirectUrl

	fmt.Println(urlshopee)
	//panic("")
	response.ResponseCode = http.StatusOK
	response.ResponseDesc = enums.SUCCESS
	response.ResponseTime = utils.DateToStdNow()
	// response.Message = enums.MSG_SUCCESS_INSERT
	response.Message = "timestamp: " + timest + " | timenow: " + time.Now().String()
	response.Result = time.Now().Unix()
	c.JSON(http.StatusOK, response)
	return

	uriApi := os.Getenv("URI_API_WMS") + ":" + os.Getenv("PORT_API_WMS")
	endPoint := "/api/v1/channel?"
	token := os.Getenv("TOKEN_API_WMS")
	bearer := "Bearer " + token

	limit := "10"
	skip := "0"

	params := "limit=" + url.QueryEscape(limit) + "&" +
		"skip=" + url.QueryEscape(skip)
	path := fmt.Sprintf(uriApi+endPoint+"%s", params)

	req, err := http.NewRequest("GET", path, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
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
		response.Message = enums.UNAUTHORIZED_OAUTH2_TOKEN
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		convertData := utils.GetByteToInterface(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = enums.MSG_SUCCESS_INSERT
		response.Result = convertData
	}

	c.JSON(http.StatusOK, response)
	return
}

func GetDataRestFull(c *gin.Context) {
	var response response.ResponseCrud

	//test auth shopee
	//timest := strconv.FormatInt(time.Unix(), 10)
	now := time.Now()

	timest := strconv.FormatInt(now.Unix(), 10)

	fmt.Println("Test Jam : ")
	fmt.Println(now)
	// fmt.Println(now.Add(time.Duration(-18000)*time.Second))

	// timest := 1634010098
	host := "https://partner.test-stable.shopeemobile.com"
	paths := "/api/v2/shop/auth_partner"
	redirectUrl := "https://supplier.ramayana.co.id/getauth.shopee/"
	partnerId := strconv.Itoa(1003569)
	partnerKey := "8f0b566912e0b7262489d2b7240949fe50226bc5821b3c208a2e8046620fd05b"
	baseString := fmt.Sprintf("%s%s%s", partnerId, paths, timest)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))
	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&sign=%s&redirect=%s", partnerId, timest, sign, redirectUrl)

	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	//req.Header.Set("Authorization", bearer)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("X-Client-Cert", "MIIHRTCCBi2gAwIBAgISBCpvHxBRDrBPkVTwsqBJIwtxMA0GCSqGSIb3DQEBCwUAMDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQDEwJSMzAeFw0yMTEwMDYwMTA4MzFaFw0yMjAxMDQwMTA4MzBaMCYxJDAiBgNVBAMMGyouYWRzLnRlc3Qtc3RhYmxlLnNob3BlZS5pbzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAOe8yacEdrkeQPFXUc0TjMc1NB9c0g2rQQuTLgL74mLpjFWVzWBlYqOSdQhqinJ9xZMKPTAMV+xdv9E8DzOSPTWRw9aOWDga+S1copnB712hNYrmrywgoYS6Id0BXj6KsBCcRQWnVtCEuUdM7y8wad4a3fW6RdHurjPJRpOlHLmw2dL+x7wPbvfCikXt+uGjdWRuWAnsP4RdKpQLHKUyq0O6pTXhOzRWL+7xXW+sS07PtGQVSQ9tCaW2DvIK4K/5sFI3KkEOTbOlkZqiNZljyrZNTTAVeP6EC+h8oaF6GiJJnbXTij5phQixNpoPkBwA6wCVDV3OokS/w3+3N7C5KRUCAwEAAaOCBF8wggRbMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUoYe+nFwMqfUJ8EWtBAtSIdu63fowHwYDVR0jBBgwFoAUFC6zF7dYVsuuUAlA5h+vnYsUwsYwVQYIKwYBBQUHAQEESTBHMCEGCCsGAQUFBzABhhVodHRwOi8vcjMuby5sZW5jci5vcmcwIgYIKwYBBQUHMAKGFmh0dHA6Ly9yMy5pLmxlbmNyLm9yZy8wggIuBgNVHREEggIlMIICIYIbKi5hZHMudGVzdC1zdGFibGUuc2hvcGVlLmlvgisqLmFwaS5mZHMuZGVlcC50ZXN0LXN0YWJsZS5zaG9wZWVtb2JpbGUuY29tgh0qLmNsb3VkLnRlc3Qtc3RhYmxlLnNob3BlZS5pb4ImKi5kYi5jbWRiLnRlc3Qtc3RhYmxlLnNob3BlZW1vYmlsZS5jb22CIyouZGVlcC50ZXN0LXN0YWJsZS5zaG9wZWVtb2JpbGUuY29tghoqLmRzLnRlc3Qtc3RhYmxlLnNob3BlZS5pb4IcKi5yY21kLnRlc3Qtc3RhYmxlLnNob3BlZS5pb4IYKi5zYW9wcy5zaG9wZWVtb2JpbGUuY29tgiIqLnNzYy50ZXN0LXN0YWJsZS5zaG9wZWVtb2JpbGUuY29tghoqLnN6LnRlc3Qtc3RhYmxlLnNob3BlZS5pb4IXKi50ZXN0LXN0YWJsZS5zaG9wZWUuaW+CHioudGVzdC1zdGFibGUuc2hvcGVlbW9iaWxlLmNvbYIpYXBpLmZkcy5kZWVwLnRlc3Qtc3RhYmxlLnNob3BlZW1vYmlsZS5jb22CJGRiLmNtZGIudGVzdC1zdGFibGUuc2hvcGVlbW9iaWxlLmNvbYIWc2FvcHMuc2hvcGVlbW9iaWxlLmNvbYIVdGVzdC1zdGFibGUuc2hvcGVlLmlvghx0ZXN0LXN0YWJsZS5zaG9wZWVtb2JpbGUuY29tMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYIKwYBBQUHAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3JnMIIBAwYKKwYBBAHWeQIEAgSB9ASB8QDvAHUA36Veq2iCTx9sre64X04+WurNohKkal6OOxLAIERcKnMAAAF8U1wCFwAABAMARjBEAiAdfPPf2BSpRd0bZx4GElToZQW/K38m5nlHyMnUXEudoQIgSKhmv7BJylZzztEO1+ggdUA6/j35xaMaYOnUJQVtpIoAdgBGpVXrdfqRIDC1oolp9PN9ESxBdL79SbiFq/L8cP5tRwAAAXxTXAJHAAAEAwBHMEUCICietJFGcTnaIXt4fCsUvjFbwW5i1MdfbA0cU7TUarh2AiEA4ozpcCEeZk9K0ymj1JTMij2sBoclJnL0gsTE8QRXXn8wDQYJKoZIhvcNAQELBQADggEBAC3Dw2os2WKsXfoK3v9gcLi2lhm6N6igGfI8e/hETVp+assBzp9SIzWeO8lc098JQDtF05YYRMoefbR+jq7mfgjshV31YmZMg0ybAAwqNv0S0DB7e8CxMYF3fYPVd0FxtuFE6HGr5CrN7MOMixqlqi6hxgsFFmnb5WezVEFgwwSdKsbc4KRAz7i8RtvBbVP1kY7tXkBLqm5OSG+gBnbE2CJ6uU5Mhyh6g5lhME0bSk6UBxweEd5yzfiabAzO10582nPNZHP7i/SA2ZS8iNwhBGtdNAKw4ddk/bMI5lMEGLGMkYeIpiPUJpOmTS90P2aw1zPXrgcE4wPdMwJaIXpXw54=")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)
	// response.ResponseCode = http.StatusOK
	// 	response.ResponseDesc = enums.SUCCESS
	// 	response.ResponseTime = utils.DateToStdNow()
	// 	response.Message = enums.MSG_SUCCESS_INSERT
	// 	response.Result = urlshopee
	// // path := urlshopee
	// c.JSON(http.StatusOK, response)
	// return
	if err != nil {
		fmt.Println("Masuk")
		log.Println("Error on response.\n[ERRO] -", err)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.ERROR
		response.ResponseTime = utils.DateToStdNow()
		response.Message = timest
		response.Result = err

	} else {
		defer resp.Body.Close()
		fmt.Println("Masuk ghfdgsjdghjsghd")
		data, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("Datanya : ")
		fmt.Println(resp.Body)
		fmt.Println("Datanya : ")

		response.Result = data

		// data, _ := ioutil.ReadAll(resp.Body)

		// convertData := utils.GetByteToInterface(data)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		// response.Result = resp
	}

	c.JSON(http.StatusOK, response)
	return
}

func GetToken(c *gin.Context) {

	code := c.Query("code")
	shop_id := c.Query("shop_id")

	var response response.ResponseCrud

	now := time.Now()

	timest := strconv.FormatInt(now.Unix(), 10)

	fmt.Println("Jam tembak : ")
	fmt.Println(now)
	fmt.Println("Jam TimeSt : ")
	fmt.Println(timest)

	host := "https://partner.test-stable.shopeemobile.com"
	paths := "/api/v2/auth/token/get"
	partnerId := strconv.Itoa(1003569)
	partnerKey := "8f0b566912e0b7262489d2b7240949fe50226bc5821b3c208a2e8046620fd05b"
	baseString := fmt.Sprintf("%s%s%s", partnerId, paths, timest)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))

	var jsonString = `{"code":"` + code + `","shop_id":` + shop_id + `,"partner_id":` + partnerId + `}`
	var body_url = []byte(jsonString)

	fmt.Println("check")
	fmt.Println(string(body_url))

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&sign=%s", partnerId, timest, sign)

	req, err := http.NewRequest("POST", urlshopee, bytes.NewBuffer(body_url))
	//req.Header.Set("Authorization", bearer)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("X-Client-Cert", "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")
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
		fmt.Println("Berhasil ! ")
		data, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("Data Body : ")
		fmt.Println(resp.Body)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		// response.Result = resp
		response.Result = string(data)
	}

	c.JSON(http.StatusOK, response)
	return
}

func RefreshToken(c *gin.Context) {

	refresh_token := c.Query("refresh_token")
	shop_id := c.Query("shop_id")

	var response response.ResponseCrud

	now := time.Now()

	timest := strconv.FormatInt(now.Unix(), 10)

	fmt.Println("Jam tembak : ")
	fmt.Println(now)
	fmt.Println("Jam TimeSt : ")
	fmt.Println(timest)

	host := "https://partner.test-stable.shopeemobile.com"
	paths := "/api/v2/auth/access_token/get"
	partnerId := strconv.Itoa(1003569)
	partnerKey := "8f0b566912e0b7262489d2b7240949fe50226bc5821b3c208a2e8046620fd05b"
	baseString := fmt.Sprintf("%s%s%s", partnerId, paths, timest)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))

	var jsonString = `{"refresh_token":"` + refresh_token + `","shop_id":` + shop_id + `,"partner_id":` + partnerId + `}`
	var body_url = []byte(jsonString)

	fmt.Println("check")
	fmt.Println(string(body_url))

	urlshopee := fmt.Sprintf(host+paths+"?partner_id=%s&timestamp=%s&sign=%s", partnerId, timest, sign)

	req, err := http.NewRequest("POST", urlshopee, bytes.NewBuffer(body_url))
	//req.Header.Set("Authorization", bearer)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("X-Client-Cert", "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")
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
		fmt.Println("Berhasil ! ")
		data, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("Data Body : ")
		fmt.Println(resp.Body)

		response.ResponseCode = http.StatusOK
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = urlshopee
		// response.Result = resp
		response.Result = string(data)
	}

	c.JSON(http.StatusOK, response)
	return
}

func GetDataRestFullById(c *gin.Context) {
	var response response.ResponseCrud

	byId := c.Param("byId")

	// limit := c.Query("limit")
	// skip := c.Query("skip")

	uriApi := os.Getenv("URI_API_WMS") + ":" + os.Getenv("PORT_API_WMS")
	endPoint := "/api/v1/channel/"
	token := os.Getenv("TOKEN_API_WMS")
	bearer := "Bearer " + token

	path := uriApi + endPoint + byId

	req, err := http.NewRequest("GET", path, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
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
		response.Message = enums.UNAUTHORIZED_OAUTH2_TOKEN
		response.Result = err

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		convertData := utils.GetByteToInterface(data)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = enums.MSG_SUCCESS_INSERT
		response.Result = convertData
	}

	c.JSON(http.StatusOK, response)
	return
}

func PostDataRestFull(c *gin.Context) {
	var response response.ResponseCrud

	value := url.Values{}
	value.Add("channel_code", c.Request.FormValue("channel_code"))
	value.Add("channel_name", c.Request.FormValue("channel_name"))
	value.Add("channel_addr", c.Request.FormValue("channel_addr"))
	value.Add("uuid_status", c.Request.FormValue("uuid_status"))
	value.Add("created_by", c.Request.FormValue("created_by"))

	uriApi := os.Getenv("URI_API_WMS") + ":" + os.Getenv("PORT_API_WMS")
	endPoint := "/api/v1/channel"
	token := os.Getenv("TOKEN_API_WMS")
	bearer := "Bearer " + token

	req, err := http.NewRequest("POST", uriApi+endPoint, strings.NewReader(value.Encode()))
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
		response.Message = enums.UNAUTHORIZED_OAUTH2_TOKEN
		response.Result = err

	} else {
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		convertData := utils.GetByteToInterface(body)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = enums.MSG_SUCCESS_INSERT
		response.Result = convertData
	}

	c.JSON(http.StatusOK, response)
	return

}

func PutDataRestFull(c *gin.Context) {
	var response response.ResponseCrud

	value := url.Values{}
	value.Add("uuid_channel", c.Request.FormValue("uuid_channel"))
	value.Add("channel_code", c.Request.FormValue("channel_code"))
	value.Add("channel_name", c.Request.FormValue("channel_name"))
	value.Add("channel_addr", c.Request.FormValue("channel_addr"))
	value.Add("uuid_status", c.Request.FormValue("uuid_status"))
	value.Add("updated_by", c.Request.FormValue("created_by"))

	uriApi := os.Getenv("URI_API_WMS") + ":" + os.Getenv("PORT_API_WMS")
	endPoint := "/api/v1/channel"
	token := os.Getenv("TOKEN_API_WMS")
	bearer := "Bearer " + token

	req, err := http.NewRequest("PUT", uriApi+endPoint, strings.NewReader(value.Encode()))
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
		response.Message = enums.UNAUTHORIZED_OAUTH2_TOKEN
		response.Result = err

	} else {
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		convertData := utils.GetByteToInterface(body)

		response.ResponseCode = http.StatusInternalServerError
		response.ResponseDesc = enums.SUCCESS
		response.ResponseTime = utils.DateToStdNow()
		response.Message = enums.MSG_SUCCESS_INSERT
		response.Result = convertData
	}

	c.JSON(http.StatusOK, response)
	return

}
