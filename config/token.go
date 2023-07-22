package config

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/rals/dearme-channel/models"
)

func CreateToken(userid string) (models.TokenDetails, error) {
	var objToken models.TokenDetails

	objToken.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	objToken.TokenUuid = uuid.New().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = objToken.TokenUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = objToken.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	objToken.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return objToken, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = objToken.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = objToken.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	objToken.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return objToken, err
	}

	return objToken, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// get the token from the request body
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenMetadata(r *http.Request) (models.AccessDetails, error) {
	fmt.Println("WE ENTERED METADATA")
	var objAccessDetails models.AccessDetails

	token, err := VerifyToken(r)
	if err != nil {
		return objAccessDetails, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return objAccessDetails, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return objAccessDetails, err
		}
		// return &objAccessDetails{
		// 	TokenUuid: accessUuid,
		// 	UserId:    userId,
		// }, nil

		objAccessDetails.TokenUuid = accessUuid
		objAccessDetails.UserId = userId
	}

	return objAccessDetails, err
}
