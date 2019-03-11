package auth

import (
	"net/http"

	//	"fmt"

	//	"github.com/Unknwon/com"
	//	"github.com/astaxie/beego/validation"
	"github.com/eaglexpf/GoAdminInit/models/auth"
	"github.com/gin-gonic/gin"
)

type AssignmentType struct {
	UserID   int   `form:"user_id" binding:"required"`
	RouteIDS []int `form:"route_ids[]"`
}

/**
 * @api {post} /auth/assignment 路由分配
 * @apiDescription 分配权限
 * @apiVersion 0.0.1
 * @apiGroup AUTH
 *
 * @apiParam {int} user_id 用户id
 * @apiParam {Array} route_ids 路由id集合
 *
 **/
func Assignment(c *gin.Context) {
	var data AssignmentType
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	err = auth.UserRouteAssignment(data.UserID, data.RouteIDS)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	auth.GetRouteList()

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
