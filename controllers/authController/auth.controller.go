package authController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rals/dearme-channel/services/authService"
	"github.com/rals/dearme-channel/utils"
)

func Login(c *gin.Context) {

	login := utils.GetJSONRawBody(c)
	username := login["username"]
	password := login["password"]

	result, err := authService.CheckLogin(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
	return
}
