package auth

import (
	"time"

	"github.com/eaglexpf/GoAdminInit/models"
	"github.com/jinzhu/gorm"
)

type AuthActionTable struct {
	ID          int    `json:"id" gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	ModelID     int    `json:"model_id" gorm:"Column:model_id;Type:int(11);NOT NULL;"`
	Name        string `json:"name" gorm:"Column:name;Type:varchar(100);NOT NULL;"`
	Description string `json:"description" gorm:"Column:description;Type:varchar(255);"`
	Status      int    `json:"status" gorm:"Column:status;Type:int(4);NOT NULL;"`
	CreateAt    int    `json:"create_at" gorm:"Column:create_at;Type:int(11);NOT NULL;"`
	UpdateAt    int    `json:"update_at" gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

func (AuthActionTable) TableName() string {
	return models.TablePrefix + "auth_action"
}

func (AuthActionTable) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateAt", time.Now().Unix())
	return nil
}

func (AuthActionTable) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateAt", time.Now().Unix())
	return nil
}

func init() {
	var table = &AuthActionTable{}
	if !models.DB.HasTable(table.TableName()) {
		models.DB.CreateTable(table)
	}
}

func CreateAction(model_id int, name string, description string, status int) error {
	return models.DB.Create(&AuthActionTable{
		ModelID:     model_id,
		Name:        name,
		Description: description,
		Status:      status,
	}).Error
}

func EditAction(id int, data interface{}) error {
	return models.DB.Model(&AuthActionTable{}).Where("id=?", id).Updates(data).Error
}

func DeleteAction(id int) error {
	return models.DB.Where("id=?", id).Delete(&AuthActionTable{}).Error
}

func ExistActionByID(id int) bool {
	var table AuthActionTable
	models.DB.Where("id=?", id).First(&table)
	if table.ID > 0 {
		return true
	}
	return false
}
