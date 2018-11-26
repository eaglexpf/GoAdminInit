package install

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/eaglexpf/GoAdminInit/models"
	"github.com/eaglexpf/GoAdminInit/pkg/e"
	"github.com/eaglexpf/GoAdminInit/pkg/setting"
	"github.com/gin-gonic/gin"
)

var tablePrefix string = ""

func init() {
	sec, err := setting.Cfg.GetSection("database")
	if err == nil {
		tablePrefix = sec.Key("TABLE_PREFIX").MustString("")
	}
}

type UserTable struct {
	ID             int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	UserName       string `gorm:"Column:user_name;Type:varchar(32);NOT NULL;UNIQUE_INDEX;"`
	AuthKey        string `gorm:"Column:auth_key;Type:varchar(32);NOT NULL;"`
	PassHash       string `gorm:"Column:pass_hash;Type:varchar(256);NOT NULL;"`
	ResetPassToken string `gorm:"Column:reset_pass_token;Type:varchar(256);"`
	Email          string `gorm:"Column:email;Type:varchar(256);NOT NULL;UNIQUE_INDEX;"`
	Status         int    `gorm:"Column:status;Type:int(11);NOT NULL;"`
	CreateAt       int    `gorm:"Column:crate_at;Type:int(11);NOT NULL;"`
	UpdateAt       int    `gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

func (UserTable) TableName() string {
	return tablePrefix + "users"
}

type AuthItemTable struct {
	ID       int    `gorm:"Column:id;PRIMARY_KEY;AUTO_INCREMENT;"`
	Name     string `gorm:"Column:name;NOT NULL;"`
	Type     int    `gorm:"Column:type;NOT NULL;"`
	Status   int    `gorm:"Column:status;Type:int(11);NOT NULL;"`
	CreateAt int    `gorm:"Column:crate_at;Type:int(11);NOT NULL;"`
	UpdateAt int    `gorm:"Column:update_at;Type:int(11);NOT NULL;"`
}

func (AuthItemTable) TableName() string {
	return tablePrefix + "auth_item"
}

type AuthItemChildTable struct {
	Parent int `gorm:"Column:parent_id;Type:int(11);NOT NULL;"`
	Child  int `gorm:"Column:child_id;Type:int(11);NOT NULL;"`
}

func (AuthItemChildTable) TableName() string {
	return tablePrefix + "auth_item_child"
}

type AuthAssignmentTable struct {
	UserID int `gorm:"Column:user_id;Type:int(11);NOT NULL;"`
	ItemID int `gorm:"Column:item_id;Type:int(11);NOT NULL;"`
}

func (AuthAssignmentTable) TableName() string {
	return tablePrefix + "auth_assignment"
}

func AddUser(c *gin.Context) {
	username := c.Query("username")
	email := c.Query("email")
	valid := validation.Validation{}
	valid.Required(username, "username").Message("username不能为空")
	valid.Required(email, "emial").Message("email不能为空")
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.ERROR),
			"data": valid.Errors,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
	})
}

func Install(c *gin.Context) {
	user := &UserTable{}
	var hasUserTable = models.DB.HasTable(user.TableName())
	if !hasUserTable {
		models.DB.CreateTable(user)
	}
	item := &AuthItemTable{}
	var hasAuthItemTable = models.DB.HasTable(item.TableName())
	if !hasAuthItemTable {
		models.DB.CreateTable(item)
	}
	itemChild := &AuthItemChildTable{}
	var hasAuthItemChildTable = models.DB.HasTable(itemChild.TableName())
	if !hasAuthItemChildTable {
		models.DB.CreateTable(itemChild)
	}
	assignment := &AuthAssignmentTable{}
	var hasAuthAssignmentTable = models.DB.HasTable(assignment.TableName())
	if !hasAuthAssignmentTable {
		models.DB.CreateTable(assignment)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
	})
}
