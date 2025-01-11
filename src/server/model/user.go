package model

import (
	"fmt"
	"server/global"
	"time"
)

// CREATE TABLE `plg_users` (
// 	`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
// 	`created_at` datetime NOT NULL,
// 	`updated_at` datetime NOT NULL,
// 	`deleted_at` datetime NOT NULL,
// 	`account` varchar(64) NOT NULL,
// 	`phone` varchar(64) NOT NULL,
// 	`password` varchar(64) NOT NULL,
// 	`nickname` varchar(256) NOT NULL,
// 	`login_ip` varchar(64) NOT NULL,
// 	`register_ip` varchar(64) NOT NULL,
// 	`status` int NOT NULL DEFAULT ' 1' COMMENT '状态'
//   );

type User struct {
	ID         uint64     `gorm:"column:id"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
	Account    string     `gorm:"column:account"`
	Phone      string     `gorm:"column:phone"`
	Password   string     `gorm:"column:password"`
	Nickname   string     `gorm:"column:nickname"`
	LoginIP    string     `gorm:"column:login_ip"`
	RegisterIP string     `gorm:"column:register_ip"`
	Status     int        `gorm:"column:status"`
}

func (User) TableName() string {
	return "plg_users"
}

// 判断是否存在
func UserIsExist(account string) bool {
	var count int64
	global.G_DB.Where(User{Account: account}).Count(&count)
	return count > 0
}

// 判断参数是否合法
func (u *User) Check() bool {
	return true
}

// 添加
func AddUser(u *User) error {
	if UserIsExist(u.Account) {
		return fmt.Errorf("%s", u.Account+"已存在")
	}

	if !u.Check() {
		return fmt.Errorf("%s", u.Account+"参数不合法")
	}
	return global.G_DB.Create(u).Error
}

// 登陆
func UserLogin(account, password string) (*User, error) {
	var user User
	err := global.G_DB.Where(User{Account: account, Password: password}).First(&user).Error
	return &user, err
}
