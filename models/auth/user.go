package auth

import (
	"fmt"

	"crypto/md5"
	"time"

	"github.com/google/uuid"

	"github.com/eaglexpf/GoAdminInit/models"
	"github.com/jinzhu/gorm"
)

type UserTable struct {
	ID             int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	UserName       string `gorm:"Column:username;Type:varchar(36);NOT NULL;UNIQUE_INDEX;"`
	AuthKey        string `gorm:"Column:auth_key;Type:varchar(36);NOT NULL;"`
	PassHash       string `gorm:"Column:pass_hash;Type:varchar(256);NOT NULL;"`
	ResetPassToken string `gorm:"Column:reset_pass_token;Type:varchar(256);"`
	Email          string `gorm:"Column:email;Type:varchar(256);NOT NULL;UNIQUE_INDEX;"`
	Status         int    `gorm:"Column:status;Type:int(11);NOT NULL;"`
	CreateAt       int    `gorm:"Column:crate_at;Type:int(11);NOT NULL;"`
	UpdateAt       int    `gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

func (UserTable) TableName() string {
	return models.TablePrefix + "user"
}

func (UserTable) BeforeCreate(scope *gorm.Scope) error {
	var t = time.Now()
	scope.SetColumn("CreateAt", t.Unix())
	return nil
}

func (UserTable) BeforeUpdate(scope *gorm.Scope) error {
	var t = time.Now()
	scope.SetColumn("UpdateAt", t.Unix())
	return nil
}

//初始化
func init() {
	user := &UserTable{}
	var hasUserTable = models.DB.HasTable(user.TableName())
	if !hasUserTable {
		models.DB.CreateTable(user)
	}
}

//判断用户是否存在——根据用户名
func ExistUserByUserName(username string) bool {
	var user UserTable
	models.DB.Where("username=?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

//判断用户是否存在——根据邮箱
func ExistUserByEmail(email string) bool {
	var user UserTable
	models.DB.Where("email=?", email).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func AddUser(username, password, email string) error {
	uid := uuid.New().String()
	data := append([]byte(uid), []byte(password)...)
	hash := fmt.Sprintf("%x", md5.Sum(data))
	err := models.DB.Create(&UserTable{
		UserName: username,
		AuthKey:  uid,
		PassHash: hash,
		Email:    email,
	}).Error
	return err
}

func GetUserByID(id int) (UserTable, error) {
	var user UserTable
	err := models.DB.Where("id=?", id).First(&user).Error
	return user, err
}

func GetUserByUserName(username string) (UserTable, error) {
	var user UserTable
	err := models.DB.Where("username=?", username).First(&user).Error
	return user, err
}

func GetUserByEmail(email string) (UserTable, error) {
	var user UserTable
	err := models.DB.Where("email=?", email).First(&user).Error
	return user, err
}

//检查用户密码是否正确
func CheckUserPassword(user UserTable, password string) bool {
	data := append([]byte(user.AuthKey), []byte(password)...)
	hash := fmt.Sprintf("%x", md5.Sum(data))
	if hash == user.PassHash {
		return true
	}
	return false
}
