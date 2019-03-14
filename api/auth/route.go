package auth

import (
	"net/http"

	"fmt"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/eaglexpf/GoAdminInit/models/auth"
	"github.com/gin-gonic/gin"
)

/**
 * @api {get} /auth/route 路由-1、列表
 * @apiDescription 获取路由列表
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {array} data 返回数据
 *
 **/
func RouteList(c *gin.Context) {
	data := auth.GetRouteAll()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

/**
 * @api {get} /auth/route/:id 路由-2、详情
 * @apiDescription 查询路由详情
 * @apiGroup AUTH
 * @apiVersion 0.1.0
 *
 * @apiParam {int} id 路由id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {array} data 返回数据
 *
 **/
func RouteInfo(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id不能为空")

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

	if !auth.ExistRouteByID(id) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "无效的id",
			"data": make(map[string]string),
		})
		return
	}

	data, err := auth.GetRouteInfoByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": make(map[string]string),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

/**
 * @api {post} /auth/route 路由-3、新建
 * @apiDescription 新建路由
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {string} name 路由名称
 * @apiParam {string} [desc] 路由描述
 * @apiParam {string} [route] 路由地址
 * @apiParam {int} [parent_id] 父节点id
 * @apiParam {int} [status] 路由状态
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {array} data 返回数据
 *
 **/
func AddRoute(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	route := c.PostForm("route")
	parent_id := com.StrTo(c.PostForm("parent_id")).MustInt()
	status := com.StrTo(c.PostForm("status")).MustInt()
	fmt.Println(name)
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name不能为空")

	if valid.HasErrors() {
		var errData = make(map[string]string)
		for _, err := range valid.Errors {
			errData[err.Key] = err.Message
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "参数错误",
			"data": errData,
		})
		return
	}
	if parent_id > 0 {
		if !auth.ExistRouteByID(parent_id) {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "无效的parent_id",
			})
			return
		}
	}

	err := auth.CreateRoute(name, desc, route, parent_id, status)
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
 * @api {put} /auth/route/:id 路由-4、修改
 * @apiDescription 修改路由信息
 * @apiVersion 0.1.0
 * @apiGroup AUTH
 *
 * @apiParam {int} id 路由id
 * @apiParam {string} name 路由名称
 * @apiParam {string} [desc] 路由描述
 * @apiParam {string} [route] 路由地址
 * @apiParam {int} [parent_id] 父节点id
 * @apiParam {int} [status] 路由状态
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {array} data 返回数据
 *
 **/
func EditAuth(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	route := c.PostForm("route")
	parent_id := com.StrTo(c.PostForm("parent_id")).MustInt()
	status := com.StrTo(c.PostForm("status")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id不能为空")

	if valid.HasErrors() {
		errData := make(map[string]string)
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
	if parent_id > 0 {
		if !auth.ExistRouteByID(parent_id) {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "无效的parent_id",
				"data": make(map[string]string),
			})
			return
		}
	}

	oldData, err := auth.GetRouteInfoByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": make(map[string]string),
		})
		return
	}

	data := make(map[string]interface{})
	if name != "" && name != oldData.Name {
		data["name"] = name
	}
	if desc != oldData.Description {
		data["description"] = desc
	}
	if route != oldData.Route {
		data["route"] = route
	}
	if parent_id != oldData.ParentID {
		data["parent_id"] = parent_id
	}
	if status != oldData.Status {
		data["status"] = status
	}

	err = auth.EditRoute(id, data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": make(map[string]string),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": make(map[string]string),
	})

}

/**
 * @api {delete} /auth/route/:id 路由-5、删除
 * @apiDescription 删除路由
 * @apiGroup AUTH
 * @apiVersion 0.1.0
 *
 * @apiParam {int} id 路由id
 *
 * @apiSuccess {int} code 状态值
 * @apiSuccess {string} msg 状态描述
 * @apiSuccess {array} data 返回数据
 *
 **/
func DeleteRoute(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id不能为空")

	if valid.HasErrors() {
		errData := make(map[string]string)
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
	err := auth.DeleteRoute(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": make(map[string]string),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": make(map[string]string),
	})
}
