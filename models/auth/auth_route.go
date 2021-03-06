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
	Children []AuthRouteList `json:"children"`
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
}

//添加路由
func CreateRoute(name, description, route string, parent_id, status int) error {
	return models.DB.Create(&AuthRouteTable{
		Name:        name,
		Description: description,
		Route:       route,
		ParentID:    parent_id,
		Status:      status,
	}).Error
}

//修改路由
func EditRoute(id int, data interface{}) error {
	return models.DB.Model(&AuthRouteTable{}).Where("id=?", id).Updates(data).Error
}

//删除路由
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

//根据路由id判断路由是否存在
func ExistRouteByID(id int) bool {
	var table AuthRouteTable
	models.DB.Where("id=?", id).First(&table)
	if table.ID > 0 {
		return true
	}
	return false
}

//根据路由id获取路由详情
func GetRouteInfoByID(id int) (AuthRouteTable, error) {
	var table AuthRouteTable
	err := models.DB.Where("id=?", id).First(&table).Error
	return table, err
}

//根据路由id获取子节点列表
func GetRouteInfoListByID(id int) []AuthRouteTable {
	var table []AuthRouteTable
	models.DB.Where("parent_id=?", id).Find(&table)
	return table
}

func tree(data []AuthRouteList, list *map[int][]AuthRouteList) []AuthRouteList {
	for key, item := range data {
		data[key].Children = tree((*list)[item.ID], list)
	}
	return data
}

//获取路由树节点列表
func GetRouteAll() []AuthRouteList {
	var table []AuthRouteList
	var data []AuthRouteList
	parentList := make(map[int][]AuthRouteList)
	models.DB.Order("parent_id asc").Find(&table)
	for _, route := range table {
		parentList[route.ParentID] = append(parentList[route.ParentID], route)
		if route.ParentID == 0 {
			data = append(data, route)
		}
	}
	data = tree(data, &parentList)
	return data
}

//获取路由有效节点列表
func GetRouteList() {
	var data []AuthRouteTable
	models.DB.Where("route!=?", "").Find(&data)
	fmt.Println(data)
}
