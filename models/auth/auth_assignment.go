package auth

import (
	"fmt"

	"github.com/eaglexpf/GoAdminInit/models"
	"github.com/jinzhu/gorm"
)

type AuthAssignmentType struct {
	UserID  int `json:"user_id" gorm:"Column:user_id;Type:int(11);NOT NULL;"`
	RouteID int `json:"route_id" gorm:"Column:route_id;Type:int(11);NOT NULL;"`
}

func (AuthAssignmentType) TableName() string {
	return models.TablePrefix + "auth_assignment"
}

func init() {
	var table = &AuthAssignmentType{}
	if !models.DB.HasTable(table.TableName()) {
		models.DB.CreateTable(table)
	}
	var user_id int
	var routeData = []int{1, 2, 3, 4, 5, 6, 7}
}

func ExistAssignment(user_id, route_id int) bool {
	var data = &AuthAssignmentType{}
	models.DB.Where("user_id=? and route_id=?", user_id, route_id).First(data)
	if data.UserID > 0 && data.RouteID > 0 {
		return true
	}
	return false
}

func UserRouteAssignment(user_id int, routeData []int) error {
	err := models.DB.Delete("user_id=?", user_id).Error
	if err != nil {
		return err
	}
	sql := ""
	for key, item := range routeData {
		sql += "(" + user_id + "," + item + ")"
		if key != len(routeData) {
			sql += ","
		}
	}
	fmt.Println(sql)
	return nil
}

func GetUserRoute(user_id int) {
	var data []AuthAssignmentType
	models.DB.Where("user_id=?", user_id).Find(&data)
	fmt.Println(data)
}
