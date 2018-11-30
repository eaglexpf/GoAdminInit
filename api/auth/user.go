package auth

import (
	//	"fmt"
	"net/http"

	"github.com/eaglexpf/GoAdminInit/pkg/e"

	"github.com/astaxie/beego/validation"
	"github.com/eaglexpf/GoAdminInit/models/auth"
	"github.com/eaglexpf/GoAdminInit/pkg/util"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	valid := validation.Validation{}
	valid.Required(username, "username").Message("username不能为空")
	valid.Required(password, "password").Message("password不能为空")
	valid.Required(email, "email").Message("email不能为空")
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.ERROR),
			"data": valid.Errors,
		})
		return
	}
	existUserByUserName := auth.ExistUserByUserName(username)
	if existUserByUserName {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "该账号已存在",
		})
		return
	}
	existUserByEmail := auth.ExistUserByEmail(email)
	if existUserByEmail {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "该邮箱已存在",
		})
		return
	}
	err := auth.AddUser(username, password, email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.ERROR),
			"data": err,
		})
		return
	}
	user, err := auth.GetUserByUserName(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.ERROR),
			"data": err,
		})
		return
	}
	//生成token
	token, err := util.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "token生成失败",
		})
		return
	}
	//发送邮件
	title := "注册认证"
	body := "恭喜你注册成功,点击下列链接认证：" + token
	to := []string{email}
	util.SendMail(to, title, body)

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	valid := validation.Validation{}
	valid.Required(username, "username").Message("username不能为空")
	valid.Required(password, "password").Message("password不能为空")
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.ERROR),
			"data": valid.Errors,
		})
		return
	}
	user, err := auth.GetUserByUserName(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "用户名不存在",
		})
		return
	}
	if !auth.CheckUserPassword(user, password) {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "密码错误",
		})
		return
	}

	token, err := util.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "token生成失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  e.SUCCESS,
		"msg":   e.GetMsg(e.SUCCESS),
		"token": token,
	})
}
