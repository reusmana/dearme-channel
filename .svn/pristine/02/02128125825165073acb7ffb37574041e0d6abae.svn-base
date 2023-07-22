package authService

import (
	"net/http"

	"github.com/rals/dearme-channel/config"
	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/models"
	"github.com/rals/dearme-channel/models/response"
	"github.com/rals/dearme-channel/repositories/authRepository"
	"github.com/rals/dearme-channel/utils"
)

func CheckLogin(username string, password string) (response.ResponseView, error) {
	var res response.ResponseView
	var obj models.Auth

	// fmt.Println("username: ", username)
	objUser, objModulars, errIntf := authRepository.CheckLogin(username, password)

	// spew.Dump(objUser)

	if errIntf != nil {
		res.ResponseCode = http.StatusFound
		res.ResponseDesc = enums.FAILED
		res.ResponseTime = utils.DateToStdNow()
		res.Result = enums.DATA_NOT_FOUND
		return res, nil
	}

	t, errToken := config.CreateToken(username)
	if errToken != nil {
		res.ResponseCode = http.StatusBadRequest
		res.ResponseDesc = enums.FAILED
		res.ResponseTime = utils.DateToStdNow()
		res.Result = enums.UNAUTHORIZED_OAUTH2_TOKEN
		return res, errToken
	}
	//set token to model
	obj.Token = t.AccessToken
	obj.RefreshToken = t.RefreshToken
	obj.IdMasterUsers = objUser.IdMasterUsers
	obj.IdMasterRoles = objUser.IdMasterRoles
	obj.Username = objUser.Username
	obj.Fullname = objUser.Fullname
	obj.Email = objUser.Email
	obj.EmailVerifiedAt = objUser.EmailVerifiedAt
	obj.Password = objUser.Password
	obj.UrlPhoto = objUser.UrlPhoto
	obj.LoginCount = objUser.LoginCount
	obj.RememberToken = objUser.RememberToken
	obj.Sequence = objUser.Sequence
	obj.IsActive = objUser.IsActive
	obj.CreatedBy = objUser.CreatedBy
	obj.CreatedDate = objUser.CreatedDate
	obj.MasterModulars = objModulars

	res.ResponseCode = http.StatusOK
	res.ResponseDesc = enums.SUCCESS
	res.ResponseTime = utils.DateToStdNow()
	res.Result = obj

	return res, nil
}
