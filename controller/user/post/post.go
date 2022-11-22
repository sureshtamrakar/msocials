package controller_user_post

import (
	"net/http"

	models_user "github.com/sureshtamrakar/socials/models/user"
	models_user_post "github.com/sureshtamrakar/socials/models/user_post"
	util_auth "github.com/sureshtamrakar/socials/util/auth"

	"github.com/gin-gonic/gin"
)

type postRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Create(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)

	}
	var req postRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	uuid := c.GetString("uuid")
	val, err := models_user.Load(uuid)
	if err != nil {

		c.JSON(http.StatusInternalServerError, "id not found")
		return
	}
	err = models_user_post.Create(val.Name, uuid, req.Title, req.Description)
	if err != nil {

		c.JSON(http.StatusInternalServerError, "could not create post")
		return
	}

}

func LoadAll(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		c.JSON(http.StatusUnauthorized, nil)

	}
	uuid := c.GetString("uuid")

	val, err := models_user_post.LoadAll(uuid)
	if err != nil {

		c.JSON(http.StatusInternalServerError, "no post to load")
		return
	}
	c.JSON(http.StatusOK, val)
	return
}
