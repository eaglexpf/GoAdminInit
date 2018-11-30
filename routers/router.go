package routers

import (
	"github.com/eaglexpf/GoAdminInit/api/auth"
	"github.com/eaglexpf/GoAdminInit/pkg/install"
	"github.com/eaglexpf/GoAdminInit/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.PUT("/register", auth.AddUser)
	r.POST("/login", auth.Login)
	r.GET("/getAuthItemList", auth.GetAuthItemList)
	r.PUT("/addAuthItem", auth.AddAuthItem)

	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	})
	r.GET("/install", install.Install)
	r.GET("/add", install.AddUser)
	r.GET("/addUser", auth.AddUser)

	return r
}
