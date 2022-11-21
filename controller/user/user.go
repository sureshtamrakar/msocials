package controller_user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models_user "github.com/sureshtamrakar/socials/models/user"
	models_user_post "github.com/sureshtamrakar/socials/models/user_post"
	models_user_share "github.com/sureshtamrakar/socials/models/user_share"
	util_auth "github.com/sureshtamrakar/socials/util/auth"
)

func Get(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		return
	}
	id := c.GetString("uuid")
	val, err := models_user.Load(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "id not found")
		return
	}
	c.JSON(http.StatusOK, val)
	return

}
func Timeline(c *gin.Context) {
	if !util_auth.Authenticate(c) {
		return
	}
	sharedPost, err := models_user_share.LoadAll(c.GetString("uuid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	post, err := models_user_post.LoadPost(c.GetString("uuid"))
	if err != nil {

		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	sharedPost = append(sharedPost, post...)
	c.JSON(http.StatusOK, sharedPost)
	return
}
