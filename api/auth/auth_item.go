package auth

import (
	"fmt"

	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/eaglexpf/GoAdminInit/models/auth"
	"github.com/gin-gonic/gin"
)

func GetAuthItemList(c *gin.Context) {
	data := auth.GetAuthItemList()
	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

func AddAuthItem(c *gin.Context) {
	name := c.PostForm("name")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": valid.Errors,
		})
		return
	}
	hasName := auth.ExistAuthItemByName(name)
	if hasName {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": "该路径已存在",
		})
		return
	}
	err := auth.AddAuthItem(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
