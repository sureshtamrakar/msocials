package routes

import (
	"github.com/gin-gonic/gin"
	controller_login "github.com/sureshtamrakar/socials/controller/login"
	controller_register "github.com/sureshtamrakar/socials/controller/register"
	controller_user "github.com/sureshtamrakar/socials/controller/user"
	controller_user_friend "github.com/sureshtamrakar/socials/controller/user/friend"
	controller_user_post "github.com/sureshtamrakar/socials/controller/user/post"
	controller_user_postlike "github.com/sureshtamrakar/socials/controller/user/postlike"
	controller_user_postshared "github.com/sureshtamrakar/socials/controller/user/postshared"
	controller_user_request "github.com/sureshtamrakar/socials/controller/user/request"
)

func AddRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", controller_login.Login)                     // login user
	r.POST("/register", controller_register.CreateUser)          // register user
	r.GET("/user", controller_user.Get)                          // get user
	r.POST("/request", controller_user_request.Request)          // submit friend request
	r.GET("/list", controller_user_request.RequestList)          // list friend request
	r.POST("/accept", controller_user_friend.RequestAccept)      // accept friend request
	r.GET("/friend-list", controller_user_friend.List)           // generates friend lis
	r.POST("/post", controller_user_post.Create)                 // create post
	r.GET("/post", controller_user_post.LoadAll)                 //loads post from friends and self
	r.POST("/post/:id", controller_user_postlike.Create)         // like a post
	r.GET("/like", controller_user_postlike.List)                // list posts you liked
	r.POST("/post/share/:id", controller_user_postshared.Create) // share a post
	r.GET("/timeline", controller_user.Timeline)                 // timeline

	return r
}
