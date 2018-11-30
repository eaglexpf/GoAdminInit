package auth

import (
	//	"fmt"

	"github.com/eaglexpf/GoAdminInit/models"
)

type AuthItemChildTable struct {
	Parent string `gorm:"Column:parent_id;Type:varchar(255);NOT NULL;"`
	Child  string `gorm:"Column:child_id;Type:varchar(255);NOT NULL;"`
}

func (AuthItemChildTable) TableName() string {
	return models.TablePrefix + "auth_item_child"
}

func init() {
	auth_item_child := &AuthItemChildTable{}
	hasTable := models.DB.HasTable(auth_item_child)
	if !hasTable {
		models.DB.CreateTable(auth_item_child)
	}
}

func AddAuthItemChild(parent, child string) error {
	err := models.DB.Create(&AuthItemChildTable{
		Parent: parent,
		Child:  child,
	}).Error
	return err
}

func GetAuthItemChildList() (data []AuthItemChildTable) {
	auth_item_child := AuthItemChildTable{}
	models.DB.Table(auth_item_child.TableName()).Find(&data)
	return
}

func GetAuthItemChildListByParent(parent string) (data []AuthItemChildTable) {
	auth_item_child := AuthItemChildTable{}
	models.DB.Table(auth_item_child.TableName()).Where("parent=?", parent).Find(&data)
	return
}
