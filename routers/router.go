package routers

import (
	"github.com/eaglexpf/GoAdminInit/api/auth"
	//	"github.com/eaglexpf/GoAdminInit/pkg/install"
	"github.com/eaglexpf/GoAdminInit/middleware/cors"
	"github.com/eaglexpf/GoAdminInit/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Cors())

	gin.SetMode(setting.RunMode)

	r.POST("/register", auth.Register)
	r.POST("/login", auth.LoginPwd)
	r.POST("/login.uuid", auth.LoginUUID)

	r.POST("/auth/model/add", auth.AddModel)
	r.POST("/auth/model/edit", auth.EditModel)
	r.POST("/auth/model/del", auth.DeleteModel)

	r.POST("/auth/action/add", auth.AddAction)
	r.POST("/auth/action/edit", auth.EditAction)
	r.POST("/auth/action/del", auth.DeleteAction)

	r.Static("/apidoc", "./apidoc")

	//	r.PUT("/register", auth.AddUser)
	//	r.POST("/login", auth.Login)
	//	r.GET("/getAuthItemList", auth.GetAuthItemList)
	//	r.PUT("/addAuthItem", auth.AddAuthItem)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	})
	//	r.GET("/install", install.Install)
	//	r.GET("/add", install.AddUser)
	//	r.GET("/addUser", auth.AddUser)

	return r
}
