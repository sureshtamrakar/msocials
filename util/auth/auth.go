package util_auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	controller_login "github.com/sureshtamrakar/socials/controller/login"
	models_user "github.com/sureshtamrakar/socials/models/user"
)

func ExtractToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func Authenticate(c *gin.Context) bool {
	tokenString := ExtractToken(c)
	pubKey, err := ioutil.ReadFile("keys/public.pem")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "unable to read key")
		return false
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "unable to parse key")
		return false
	}
	value := &controller_login.SignedDetails{}
	token, err := jwt.ParseWithClaims(tokenString, value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, "Signature invalid")
			return false
		}
		c.JSON(http.StatusBadRequest, "unable to authenticate")
		return false
	}
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, "unable to authenticate")
		return false
	}
	val, err := models_user.GetUuid(value.Email)
	if err != nil {

		c.JSON(http.StatusUnauthorized, "Email not found")
		return false
	}
	c.Set("uuid", val.Uuid)
	return true
}
