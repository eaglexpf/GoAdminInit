package auth

import (
	"net/http"

	"github.com/eaglexpf/GoAdminInit/pkg/e"

	"github.com/astaxie/beego/validation"
	"github.com/eaglexpf/GoAdminInit/models/auth"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	username := c.Query("username")
	valid := validation.Validation{}
	valid.Required(username, "username").Message("username不能为空")
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.ERROR),
			"data": valid.Errors,
		})
		return
	}
	isUserByUserName := auth.ExitUserByUserName(username)
	if !isUserByUserName {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "没有该账号",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
	})
}
