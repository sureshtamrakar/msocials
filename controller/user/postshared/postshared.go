package controller_user_postshared

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models_user_friend "github.com/sureshtamrakar/socials/models/user_friend"
	models_user_post "github.com/sureshtamrakar/socials/models/user_post"
	models_user_share "github.com/sureshtamrakar/socials/models/user_share"
	util_auth "github.com/sureshtamrakar/socials/util/auth"
)

func Create(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)

	}

	postId := c.Param("id")
	i1, _ := strconv.Atoi(postId)
	post, err := models_user_post.Load(i1)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	val, err := models_user_friend.LoadAll(c.GetString("uuid"))

	if err != nil || len(val) == 0 {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	for _, v := range val {
		if post.UserUuid != v.UserUuid {
			c.JSON(http.StatusInternalServerError, "user is not your friend")
			return
		}
	}
	err = models_user_share.Create(i1, c.GetString("uuid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "could not share post")
		return
	}
}
