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
	r.POST("/login.check", auth.CheckLogin)

	r.GET("/auth/model", auth.ModelList)
	r.GET("/auth/model/:id", auth.ModelInfo)
	r.POST("/auth/model/add", auth.AddModel)
	r.POST("/auth/model/edit", auth.EditModel)
	r.POST("/auth/model/del", auth.DeleteModel)

	r.GET("/auth/action/info/:id", auth.ActionInfo)
	r.GET("/auth/action/model/:model_id", auth.ActionList)
	r.POST("/auth/action/add", auth.AddAction)
	r.POST("/auth/action/edit", auth.EditAction)
	r.POST("/auth/action/del", auth.DeleteAction)

	r.Static("/apidoc", "./apidoc")

	r.GET("/auth/route", auth.RouteList)
	r.GET("/auth/route/:id", auth.RouteInfo)
	r.POST("/auth/route", auth.AddRoute)
	r.PUT("/auth/route/:id", auth.EditAuth)
	r.DELETE("/auth/route/:id", auth.DeleteRoute)

	r.POST("/auth/assignment", auth.Assignment)

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
