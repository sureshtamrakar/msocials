package controller_login

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	models_login "github.com/sureshtamrakar/socials/models/login"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type SignedDetails struct {
	Email string
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var req models_login.Entity

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	val, err := models_login.Login(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Email Not Found")
		return
	}
	errf := bcrypt.CompareHashAndPassword([]byte(val.Password), []byte(req.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusUnauthorized, "Password does not match!")
		return

	}

	prvKey, err := ioutil.ReadFile("keys/private.pem")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "unable to read key")
	}
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(prvKey)
	value := &SignedDetails{
		Email: val.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodRS256, value)
	token, err := claims.SignedString(key)
	c.Header("Access-Token", token)
	c.JSON(http.StatusOK, "Login Validated")
	return
}
