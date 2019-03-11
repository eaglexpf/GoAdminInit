package auth

import (
	"fmt"

	"bytes"

	"strconv"

	"github.com/eaglexpf/GoAdminInit/models"
	//	"github.com/jinzhu/gorm"
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
	var table = AuthAssignmentType{}
	err := models.DB.Delete(table, "user_id=?", user_id).Error
	if err != nil {
		return err
	}
	if len(routeData) == 0 {
		return nil
	}
	fmt.Println(user_id, routeData)
	var buffer bytes.Buffer
	buffer.WriteString("insert into ")
	buffer.WriteString(table.TableName())
	buffer.WriteString(" values ")
	for key, item := range routeData {
		buffer.WriteString("(")
		buffer.WriteString(strconv.Itoa(user_id))
		buffer.WriteString(",")
		buffer.WriteString(strconv.Itoa(item))
		buffer.WriteString(")")
		if key != len(routeData)-1 {
			buffer.WriteString(",")
		}
	}
	fmt.Println(buffer.String())
	err = models.DB.Exec(buffer.String()).Error
	return err
}

func GetUserRoute(user_id int) {
	var data []AuthAssignmentType
	models.DB.Where("user_id=?", user_id).Find(&data)
	fmt.Println(data)
}
