package auth

import (
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/eaglexpf/GoAdminInit/models/auth"
	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /auth/action/model/:model_id 模块方法-1.列表
 * @apiDescription 根据模块id获取该模块下的所有方法列表
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {int} model_id 模块id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 */
func ActionList(c *gin.Context) {
	model_id := com.StrTo(c.Param("model_id")).MustInt()

	valid := validation.Validation{}
	valid.Min(model_id, 1, "model_id").Message("model_id不能为空")

	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "model_id不能为空",
		})
		return
	}
	data := auth.ActionListByModelID(model_id)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

/**
 * @api {get} /auth/action/info/:id 模块方法-2.详情
 * @apiDescription 根据方法id获取方法详情
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {int} id 方法id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 */
func ActionInfo(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id不能为空")

	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "id不能为空",
		})
		return
	}
	data, err := auth.ActionInfoByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

/**
 * @api {post} /auth/action/add 模块方法-3.添加
 * @apiDescription 添加权限方法
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {int} model_id 模块id
 * @apiParam {string} name 方法名称
 * @apiParam {string} [description] 方法描述
 * @apiParam {string} [status] 方法状态
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 **/
func AddAction(c *gin.Context) {
	model_id := com.StrTo(c.PostForm("model_id")).MustInt()
	name := c.PostForm("name")
	description := c.PostForm("description")
	status := com.StrTo(c.PostForm("status")).MustInt()

	valid := validation.Validation{}
	valid.Min(model_id, 1, "model_id").Message("model_id不能小于1")
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
	if !auth.ExistModelByID(model_id) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "无效的model_id",
		})
		return
	}
	err := auth.CreateAction(model_id, name, description, status)
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
 * @api {post} /auth/action/edit 模块方法-4.修改
 * @apiDescription 修改权限方法
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {int} id 方法id
 * @apiParam {int} model_id 模块id
 * @apiParam {string} name 方法名称
 * @apiParam {string} [description] 方法描述
 * @apiParam {string} [status] 方法状态
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 **/
func EditAction(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	model_id := com.StrTo(c.PostForm("model_id")).MustInt()
	name := c.PostForm("name")
	description := c.PostForm("description")
	status := com.StrTo(c.PostForm("status")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id不能小于1")
	valid.Min(model_id, 1, "model_id").Message("model_id不能小于1")
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

	if !auth.ExistActionByID(id) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "无效的id",
		})
		return
	}

	if !auth.ExistModelByID(model_id) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "无效的model_id",
		})
		return
	}

	var data = make(map[string]interface{})
	data["model_id"] = model_id
	data["name"] = name
	data["description"] = description
	data["status"] = status
	err := auth.EditAction(id, data)
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
 * @api {post} /auth/action/del 模块方法-5.删除
 * @apiDescription 删除权限方法
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {int} id 方法id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 *
 **/
func DeleteAction(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id不能小于1")

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

	err := auth.DeleteAction(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "succcess",
	})
}
