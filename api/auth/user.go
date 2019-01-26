package auth

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/eaglexpf/GoAdminInit/models/auth"
	"github.com/eaglexpf/GoAdminInit/pkg/util"
	"github.com/gin-gonic/gin"
)

/**
 * @api {post} /register 注册
 * @apiDescription 注册接口
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {string} username 账号
 * @apiParam {string} password 密码
 * @apiParam {string} email 邮箱
 * @apiParam {string} uuid 客户端唯一标识
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 */
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	uuid := c.PostForm("uuid")

	valid := validation.Validation{}
	valid.Required(username, "username").Message("账号不能为空")
	valid.Required(password, "password").Message("密码不能为空")
	valid.Required(email, "email").Message("email不能为空")
	valid.Email(email, "emial").Message("请输入正确的邮箱")
	valid.Required(uuid, "uuid").Message("uuid不能为空")

	if valid.HasErrors() {
		var errData = make(map[string]interface{})
		for _, err := range valid.Errors {
			errData[err.Key] = err.Message
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
			"data": errData,
		})
		return
	}

	if auth.ExistUserByUserName(username) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "该账号已存在，请换一个名字",
		})
		return
	}

	if auth.ExistUserByEmail(email) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "该邮箱已存在，请换一个邮箱",
		})
		return
	}
	err := auth.AddUser(username, password, email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

/**
 * @api {post} /login 登录
 * @apiDescription 登录接口
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {string} username 账号
 * @apiParam {string} password 密码
 * @apiParam {string} uuid 客户端唯一标识
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 */
func LoginPwd(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	uuid := c.PostForm("uuid")

	valid := validation.Validation{}
	valid.Required(username, "username").Message("账号不能为空")
	valid.Required(password, "password").Message("密码不能为空")
	valid.Required(uuid, "uuid").Message("uuid不能为空")

	if valid.HasErrors() {
		var errData = make(map[string]string)
		for _, err := range valid.Errors {
			errData[err.Key] = err.Message
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
			"data": errData,
		})
		return
	}
	user, err := auth.GetUserByUserName(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "账号不存在",
		})
		return
	}
	if !auth.CheckUserPassword(user, password) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}

	err = auth.UUIDLogin(user.ID, uuid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	var data = struct {
		UserName string
		UUID     string
	}{
		user.UserName,
		uuid,
	}
	token, err := util.GenerateToken(user.ID, uuid, data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "token生成失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "success",
		"token": token,
	})
}

/**
 * @api {post} /login.uuid uuid登录
 * @apiDescription uuid登录接口
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {string} uuid 客户端唯一标识
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 */
func LoginUUID(c *gin.Context) {
	uuid := c.PostForm("uuid")

	valid := validation.Validation{}
	valid.Required(uuid, "uuid").Message("uuid不能为空")

	if valid.HasErrors() {
		var errData = make(map[string]string)
		for _, err := range valid.Errors {
			errData[err.Key] = err.Message
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
			"data": errData,
		})
		return
	}
	user, err := auth.GetUserByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	fmt.Println(user)
	var data = struct {
		UserName string
		UUID     string
	}{
		user.User.UserName,
		user.UUID,
	}
	token, err := util.GenerateToken(user.UserId, user.UUID, data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "token生成失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "success",
		"token": token,
	})
}
