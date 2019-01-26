package auth

import (
	"time"

	"github.com/eaglexpf/GoAdminInit/models"
	"github.com/jinzhu/gorm"
)

type AuthModelTable struct {
	ID          int    `json:"id" gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	ParentID    int    `json:"parent_id" gorm:"Column:parent_id;Type:int(11);NOT NULL;"`
	Name        string `json:"name" gorm:"Column:name;Type:varchar(100);NOT NULL;"`
	Description string `json:"description" gorm:"Column:description;Type:varchar(255);"`
	Status      int    `json:"status" gorm:"Column:status;Type:int(4);NOT NULL;"`
	CreateAt    int    `json:"create_at" gorm:"Column:create_at;Type:int(11);NOT NULL;"`
	UpdateAt    int    `json:"update_at" gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

func (AuthModelTable) TableName() string {
	return models.TablePrefix + "auth_model"
}
func (AuthModelTable) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateAt", time.Now().Unix())
	return nil
}
func (AuthModelTable) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateAt", time.Now().Unix())
	return nil
}

func init() {
	var table = &AuthModelTable{}
	if !models.DB.HasTable(table.TableName()) {
		models.DB.CreateTable(table)
	}
}

func CreateModel(parent_id int, name string, description string, status int) error {
	return models.DB.Create(&AuthModelTable{
		ParentID:    parent_id,
		Name:        name,
		Description: description,
		Status:      status,
	}).Error
}

func EditModel(id int, data interface{}) error {
	return models.DB.Model(&AuthModelTable{}).Where("id=?", id).Updates(data).Error
}

func DeleteModel(id int) error {
	return models.DB.Where("id=?", id).Delete(&AuthModelTable{}).Error
}

func ExistModelByID(id int) bool {
	var table AuthModelTable
	models.DB.Where("id=?", id).First(&table)
	if table.ID > 0 {
		return true
	}
	return false
}

func GetModelInfoByID(id int) (AuthModelTable, error) {
	var table AuthModelTable
	err := models.DB.Where("id=?", id).First(&table).Error
	return table, err
}

func GetModelAll() []AuthModelTable {
	var table []AuthModelTable
	models.DB.Find(&table)
	return table
}
