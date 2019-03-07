package auth

import (
	"fmt"
	"time"

	"github.com/eaglexpf/GoAdminInit/models"
	"github.com/jinzhu/gorm"
)

type AuthRouteTable struct {
	ID          int    `json:"id" gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	Name        string `json:"name" gorm:"Column:name;Type:varchar(100);NOT NULL;"`
	Description string `json:"description" gorm:"Column:description;Type:varchar(255);"`
	Route       string `json:"route" gorm:"Column:route;Type:varchar(50);"`
	ParentID    int    `json:"parent_id" gorm:"Column:parent_id;Type:int(11);NOT NULL;"`
	Status      int    `json:"status" gorm:"Column:status;Type:int(4);NOT NULL;"`
	CreateAt    int    `json:"create_at" gorm:"Column:create_at;Type:int(11);NOT NULL;"`
	UpdateAt    int    `json:"update_at" gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

type AuthRouteList struct {
	AuthRouteTable
	Children []AuthRouteTable
}

func (AuthRouteTable) TableName() string {
	return models.TablePrefix + "auth_route"
}

func (AuthRouteTable) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateAt", time.Now().Unix())
	return nil
}

func (AuthRouteTable) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateAt", time.Now().Unix())
	return nil
}

func init() {
	var table = &AuthRouteTable{}
	if !models.DB.HasTable(table.TableName()) {
		models.DB.CreateTable(table)
	}
	GetRouteAll()
}

func CreateRoute(name, description, route string, parent_id, status int) error {
	return models.DB.Create(&AuthRouteTable{
		Name:        name,
		Description: description,
		Route:       route,
		ParentID:    parent_id,
		Status:      status,
	}).Error
}

func EditRoute(id int, data interface{}) error {
	return models.DB.Model(&AuthRouteTable{}).Where("id=?", id).Updates(data).Error
}

func DeleteRoute(id int) error {
	route, err := GetRouteInfoByID(id)
	if err != nil {
		return err
	}
	tx := models.DB.Begin()
	err = tx.Where("id=?", id).Delete(&AuthRouteTable{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&AuthRouteTable{}).Where("parent_id=?", id).Update("ParentID", route.ParentID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func ExistRouteByID(id int) bool {
	var table AuthRouteTable
	models.DB.Where("id=?", id).First(&table)
	if table.ID > 0 {
		return true
	}
	return false
}

func GetRouteInfoByID(id int) (AuthRouteTable, error) {
	var table AuthRouteTable
	err := models.DB.Where("id=?", id).First(&table).Error
	return table, err
}

func GetRouteInfoListByID(id int) []AuthRouteTable {
	var table []AuthRouteTable
	models.DB.Where("parent_id=?", id).Find(&table)
	return table
}

func GetRouteAll() []AuthRouteTable {
	var table []AuthRouteTable
	var data []AuthRouteList
	models.DB.Order("parent_id asc").Find(&table)
	//	fmt.Println(table)
	for _, route := range table {
		//		fmt.Println(route)
		if route.ParentID == 0 {
			var item = AuthRouteList{route}
			append(data, route)
		}
	}
	fmt.Println(data)
	return table
}
