package auth

import (
	"fmt"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/eaglexpf/GoAdminInit/models/auth"
	"github.com/gin-gonic/gin"
)

/**
 * @api {post} /auth/model/add 添加模块
 * @apiDescription 新加权限模块
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {string} name 模块名称
 * @apiParam {string} [description] 描述
 * @apiParam {int} [parent_id] 父节点id
 * @apiParam {int} [status] 模块状态
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 **/
func AddModel(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	parent_id := com.StrTo(c.PostForm("parent_id")).MustInt()
	status := com.StrTo(c.PostForm("status")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")

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
	fmt.Println(parent_id, name, description, status)
	err := auth.CreateModel(parent_id, name, description, status)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
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
 * @api {post} /auth/model/edit 修改模块
 * @apiDescription 修改权限模块
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {int} id 模块id
 * @apiParam {string} name 模块名称
 * @apiParam {string} [description] 描述
 * @apiParam {int} [parent_id] 父节点id
 * @apiParam {int} [status] 模块状态
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 **/
func EditModel(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	name := c.PostForm("name")
	description := c.PostForm("description")
	parent_id := com.StrTo(c.PostForm("parent_id")).MustInt()
	status := com.StrTo(c.PostForm("status")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id不能小于1")
	valid.Required(name, "name").Message("name不能为空")

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
	if !auth.ExistModelByID(id) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "无效的模块id",
		})
		return
	}
	var data = make(map[string]interface{})
	data["parent_id"] = parent_id
	data["name"] = name
	data["description"] = description
	data["status"] = status
	err := auth.EditModel(id, data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
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
 * @api {post} /auth/model/del 删除模块
 * @apiDescription 删除权限模块
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {int} id 模块id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 **/
func DeleteModel(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于1")

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
	if !auth.ExistModelByID(id) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "无效的模块id",
		})
		return
	}

	err := auth.DeleteModel(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
