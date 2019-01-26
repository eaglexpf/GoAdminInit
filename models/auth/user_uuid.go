package auth

import (
	"fmt"
	"time"

	"github.com/eaglexpf/GoAdminInit/models"
	"github.com/jinzhu/gorm"
)

type UserUuidTable struct {
	User     UserTable `json:"user" gorm:"ForeignKey:UserId;AssociationForeignKey:ID"`
	UserId   int       `json:"user_id" gorm:"index;Column:user_id;Type:int(11);NOT NULL;user_id_idx;"`
	UUID     string    `json:"uuid" gorm:"Column:uuid;Type:varchar(50);NOT NULL;UNIQUE_INDEX;"`
	CreateAt int       `json:"create_at" gorm:"Column:create_at;Type:int(11);NOT NULL;"`
	ValidAt  int       `json:"valid_at" gorm:"Column:valid_at;Type:int(11);Not NULL;"`
}

func (UserUuidTable) TableName() string {
	return models.TablePrefix + "user_uuid"
}

func (UserUuidTable) BeforeCreate(scope *gorm.Scope) error {
	var t = time.Now()
	scope.SetColumn("CreateAt", t.Unix())
	valid_t := t.AddDate(0, 0, 30)
	scope.SetColumn("ValidAt", valid_t.Unix())
	return nil
}

func init() {
	table := &UserUuidTable{}
	if !models.DB.HasTable(table.TableName()) {
		models.DB.CreateTable(table)
	}
}

func GetUserByUUID(uuid string) (UserUuidTable, error) {
	var table UserUuidTable
	err := models.DB.Where("uuid=?", uuid).First(&table).Error
	if err == nil {
		err = models.DB.Model(&table).Related(&table.User, "User").Error
	}
	return table, err
}

func ExistUUIDByUserID(user_id int) bool {
	var table UserUuidTable
	models.DB.Where("user_id=?", user_id).First(&table)
	if table.UserId > 0 {
		return true
	}
	return false
}

func ExistUUIDByUUID(uuid string) bool {
	var table UserUuidTable
	models.DB.Where("uuid=?", uuid).First(&table)
	if table.UserId > 0 {
		return true
	}
	return false
}

func AddUserUUID(user_id int, uuid string) error {
	return models.DB.Create(&UserUuidTable{
		UserId: user_id,
		UUID:   uuid,
	}).Error
}

func DeleteUUIDByUserID(user_id int) error {
	return models.DB.Where("user_id=?", user_id).Delete(&UserUuidTable{}).Error
}

func DeleteUUIDByUserIDForOver(user_id int) error {
	return models.DB.Where("user_id=? and valid_at<", user_id, time.Now().Unix()).Delete(&UserUuidTable{}).Error
}

func DeleteUUIDByUUID(uuid string) error {
	return models.DB.Where("uuid=?", uuid).Delete(&UserUuidTable{}).Error
}

func UUIDLogin(user_id int, uuid string) error {
	tx := models.DB.Begin()
	err := tx.Where("uuid=?", uuid).Delete(&UserUuidTable{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Where("user_id=? and valid_at<?", user_id, time.Now().Unix()).Delete(&UserUuidTable{}).Error
	fmt.Println()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Create(&UserUuidTable{
		UserId: user_id,
		UUID:   uuid,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
