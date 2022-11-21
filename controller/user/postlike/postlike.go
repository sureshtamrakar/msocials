package controller_user_postlike

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models_user_post "github.com/sureshtamrakar/socials/models/user_post"
	util_auth "github.com/sureshtamrakar/socials/util/auth"
)

func Create(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)

	}
	postId := c.Param("id")
	uuid := c.GetString("uuid")

	i1, err := strconv.Atoi(postId)
	if err == nil {
		fmt.Println(i1)
	}
	_, err = models_user_post.Load(i1)
	if err != nil {

		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	models_user_post.Like(i1, uuid)
}

func List(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)

	}
	uuid := c.GetString("uuid")

	val, err := models_user_post.LikeList(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, val)
}
