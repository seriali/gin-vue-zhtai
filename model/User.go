package model

import (
	"gin-vue-zhtai-server/utils/message"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255); not null;comment:用户名"`
	Password string `json:"-" gorm:"type:varchar(255); not null; comment:密码"`
	Nickname string `json:"nickname" gorm:"type:varchar(255); not null; comment:用户昵称"`
	Role     int    `json:"role" gorm:"type:int"`
}

// 登录验证用户名和密码
func CheckLogin(username, password string) int {
	var u User
	err := db.Where("username = ?", username).First(&u).Error
	if err != nil {
		return message.UsernameNotExit
	}
	if u.Password != password {
		return message.UserPasswordWrong
	}
	return message.SUCCESS
}
func Login(u *User) (error, *User) {
	var user User
	err := db.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}
