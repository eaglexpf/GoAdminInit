package auth

import (
	//	"fmt"

	"time"

	"github.com/eaglexpf/GoAdminInit/models"
	"github.com/jinzhu/gorm"
)

type AuthItemTable struct {
	Name     string `gorm:"Column:name;Type:varchar(255);NOT NULL;UNIQUE_INDEX;"`
	Type     int    `gorm:"Column:type;Type:int(11);NOT NULL;"`
	Status   int    `gorm:"Column:status;Type:int(11);NOT NULL;"`
	CreateAt int    `gorm:"Column:create_at;Type:int(11);NOT NULL;"`
	UpdateAt int    `gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

func (AuthItemTable) TableName() string {
	return models.TablePrefix + "auth_item"
}

func (AuthItemTable) BeforeCreate(scope *gorm.Scope) error {
	var t = time.Now()
	scope.SetColumn("create_at", t.Unix())
	return nil
}

func (AuthItemTable) BeforeUpdate(scope *gorm.Scope) error {
	var t = time.Now()
	scope.SetColumn("update_at", t.Unix())
	return nil
}

//初始化，没有表则创建
func init() {
	auth_item := &AuthItemTable{}
	var hasTable = models.DB.HasTable(auth_item.TableName())
	if !hasTable {
		models.DB.CreateTable(auth_item)
	}
}

func ExistAuthItemByName(name string) bool {
	var auth_item AuthItemTable
	err := models.DB.Where("name=?", name).First(&auth_item).Error
	if err != nil {
		return true
	}
	return false
}

func AddAuthItem(name string) error {
	err := models.DB.Create(&AuthItemTable{
		Name: name,
	}).Error
	return err
}

func GetAuthItemList() (data []AuthItemTable) {
	auth_item := AuthItemTable{}
	models.DB.Table(auth_item.TableName()).Find(data)
	return
}

func GetAuthItemByID(id int) AuthItemTable {
	var data AuthItemTable
	models.DB.First(&data, id)
	return data
}

func GetAuthItemByName(name string) AuthItemTable {
	var data AuthItemTable
	models.DB.Where("name=?", name).First(&data)
	return data
}

func DeleteAuthItemmByID(id int) error {
	err := models.DB.Where("id=?", id).Delete(AuthItemTable{}).Error
	return err
}

func DeleteAuthItemByName(name string) error {
	err := models.DB.Where("name=?", name).Delete(AuthItemTable{}).Error
	return err
}
