package controller_user_friend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models_user_friend "github.com/sureshtamrakar/socials/models/user_friend"
	models_user_request "github.com/sureshtamrakar/socials/models/user_request"
	util_auth "github.com/sureshtamrakar/socials/util/auth"
)

type AcceptRequest struct {
	Uuid string `json:"uuid"`
}

func RequestAccept(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)

	}
	var req AcceptRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	count, err := models_user_request.Load(c.GetString("uuid"), req.Uuid)
	if err != nil || count != 1 {
		c.JSON(http.StatusInternalServerError, "no request from user or already in friend list")
		return
	}
	err = models_user_friend.Create(c.GetString("uuid"), req.Uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "could not accept friend request")
		return
	}
	models_user_request.Update(c.GetString("uuid"), req.Uuid)
}

func List(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)

	}
	val := models_user_friend.List(c.GetString("uuid"))

	c.JSON(http.StatusOK, val)
	return
}
