package model

import "time"

type User struct {
	Id         string    `gorm:"column:id;primaryKey" column:"id" json:"id"`     // id
	Name       string    `gorm:"column:name" column:"name" json:"name"`          // 昵称
	Account    string    `gorm:"column:account" column:"account" json:"account"` // 账号
	Password   string    `gorm:"column:password" column:"password" json:"-"`     // 密码
	Phone      string    `gorm:"column:phone" column:"phone" json:"phone"`       // 电话
	Email      string    `gorm:"column:email" column:"email" json:"email"`       // 邮箱
	Picture    string    `gorm:"column:picture" column:"picture" json:"picture"` // 头像
	Gender     int       `gorm:"column:gender" column:"gender" json:"gender"`    // 性别
	Age        int       `gorm:"column:age" column:"age" json:"age"`             // 年龄
	Status     bool      `gorm:"column:status" column:"status" json:"status"`    // 是否启用
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`           // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`           // 更新时间
}

func (receiver *User) TableName() string {
	return "app_user"
}
