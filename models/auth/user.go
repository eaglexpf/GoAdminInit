package auth

import (
	"fmt"

	"github.com/eaglexpf/GoAdminInit/models"
)

type UserTable struct {
	ID             int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	UserName       string `gorm:"Column:username;Type:varchar(32);NOT NULL;UNIQUE_INDEX;"`
	AuthKey        string `gorm:"Column:auth_key;Type:varchar(32);NOT NULL;"`
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

func init() {
	user := &UserTable{}
	var hasUserTable = models.DB.HasTable(user.TableName())
	fmt.Println("user表测试：", hasUserTable)
	if !hasUserTable {
		models.DB.CreateTable(user)
	}
}

func ExitUserByUserName(username string) bool {
	var user UserTable
	models.DB.Where("username=?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func ExitUserByEmail(email string) bool {
	var user UserTable
	models.DB.Where("email=?", email).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}
