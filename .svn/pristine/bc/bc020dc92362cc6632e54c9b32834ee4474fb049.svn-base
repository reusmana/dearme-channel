package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetJSONRawBody(c *gin.Context) map[string]string {

	jsonBody := make(map[string]string)
	err := json.NewDecoder(c.Request.Body).Decode(&jsonBody)
	if err != nil {

		// log.Error("empty json body")
		return nil
	}

	return jsonBody
}

func GetJSONRawBodyobject(c *gin.Context) map[string][]string {

	jsonBody := make(map[string][]string)
	err := json.NewDecoder(c.Request.Body).Decode(&jsonBody)
	if err != nil {

		// log.Error("json kosong")
		return nil
	}

	return jsonBody
}

func GetJSONToInterface(params interface{}) string {

	jsonBody, _ := json.MarshalIndent(&params, "", "  ")
	return string(jsonBody)
}

func GetByteToInterface(data []byte) map[string]interface{} {

	var convertData map[string]interface{}
	err := json.Unmarshal([]byte(data), &convertData)
	if err != nil {
		panic(err)
	}

	return convertData
}
