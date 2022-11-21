package controller_user_request

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models_user "github.com/sureshtamrakar/socials/models/user"
	models_user_request "github.com/sureshtamrakar/socials/models/user_request"
	util_auth "github.com/sureshtamrakar/socials/util/auth"
)

//User can request by email address
type UserRequest struct {
	Email string `json:"email"`
}

type AcceptRequest struct {
	Uuid string `json:"uuid"`
}

func Request(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		return
	}
	var req UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uuid, err := models_user.GetUuid(req.Email)
	if c.GetString("uuid") == uuid.Uuid {
		c.JSON(http.StatusInternalServerError, "Could not send friend request to yourself")
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not find email")
		return
	}

	count, err := models_user_request.Load(c.GetString("uuid"), uuid.Uuid)
	if err != nil || count >= 1 {
		c.JSON(http.StatusInternalServerError, "user has sent friend request to you already")
		return
	}
	models_user_request.Create(c.GetString("uuid"), uuid.Uuid)

}

func RequestList(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)

	}
	requestList, err := models_user_request.List(c.GetString("uuid"))
	if err != nil {

		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	var value []models_user.Entity

	for _, users := range requestList {
		user, err := models_user.Load(users.User_Uuid)
		if err == nil {
			value = append(value, user)
		}
	}
	c.JSON(http.StatusOK, value)
	return
}

// func RequestAccept(c *gin.Context) {
// 	if !util_auth.Authenticate(c) {
// 		c.JSON(http.StatusUnauthorized, nil)

// 	}
// 	var req AcceptRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	models_user_friend.Create(c.GetString("uuid"), req.Uuid)
// }
