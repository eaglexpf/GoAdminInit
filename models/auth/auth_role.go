package auth

import (
	"time"

	"github.com/eaglexpf/GoAdminInit/models"
	"github.com/jinzhu/gorm"
)

type AuthRoleTable struct {
	ID          int    `json:"id" gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	Name        string `json:"name" gorm:"Column:name;Type:varchar(100);NOT NULL;"`
	Description string `json:"description" gorm:"Column:description;Type:varchar(255);"`
	Status      int    `json:"status" gorm:"Column:status;Type:int(4);NOT NULL;"`
	CreateAt    int    `json:"create_at" gorm:"Column:create_at;Type:int(11);NOT NULL;"`
	UpdateAt    int    `json:"update_at" gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

func (AuthRoleTable) TableName() string {
	return models.TablePrefix + "auth_role"
}

func (AuthRoleTable) BeforCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateAt", time.Now().Unix())
	return nil
}

func (AuthRoleTable) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateAt", time.Now().Unix())
	return nil
}

func init() {
	var table = &AuthRoleTable{}
	if !models.DB.HasTable(table.TableName()) {
		models.DB.CreateTable(table)
	}
}

func CreateRole(name, description string, status int) error {
	return models.DB.Create(&AuthRoleTable{
		Name:        name,
		Description: description,
		Status:      status,
	}).Error
}

func UpdateRole(id int, data interface{}) error {
	return models.DB.Model(&AuthRoleTable{}).Where("id=?", id).Update(data).Error
}

func DeleteRole(id int) error {
	return models.DB.Where("id=?", id).Delete(&AuthRoleTable{}).Error
}
