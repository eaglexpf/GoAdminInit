package auth

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/eaglexpf/GoAdminInit/models/auth"
	"github.com/gin-gonic/gin"
)

func AddAuthItemChild(c *gin.Context) {
	parent := c.PostForm("parent")
	child := c.PostForm("child")
	valid := validation.Validation{}
	valid.Required(parent, "parent").Message("parent不能为空")
	valid.Required(child, "child").Message("child不能为空")
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": valid.Errors,
		})
		return
	}
	hasParent := auth.ExistAuthItemByName(parent)
	if !hasParent {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": "无效的parent",
		})
		return
	}
	hasChild := auth.ExistAuthItemByName(child)
	if !hasChild {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "无效的child",
		})
		return
	}
	err := auth.AddAuthItemChild(parent, child)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "添加错误",
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func GetAuthItemChildList(c *gin.Context) {
	parent := c.PostForm("parent")
	data := auth.GetAuthItemChildListByParent(parent)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
