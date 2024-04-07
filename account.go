package model

type ImAccount struct {
	Id       string `column:"id" json:"id"`           // id
	Name     string `column:"name" json:"name"`       // 昵称
	Account  string `column:"account" json:"account"` // 账号
	Password string `column:"password" json:"-"`      // 密码
	Picture  string `column:"picture" json:"picture"` // 头像
	Gender   string `column:"gender" json:"gender"`   // 性别
	Age      int    `column:"age" json:"age"`         // 年龄
	Created  string `column:"created" json:"created"` // 创建时间
}

type User struct {
	Id         string `column:"id" json:"id"`                   // id
	Name       string `column:"name" json:"name"`               // 昵称
	Account    string `column:"account" json:"account"`         // 账号
	Password   string `column:"password" json:"-"`              // 密码
	Picture    string `column:"picture" json:"picture"`         // 头像
	Gender     int    `column:"gender" json:"gender"`           // 性别
	Age        int    `column:"age" json:"age"`                 // 年龄
	CreateTime string `column:"create_time" json:"create_time"` // 创建时间
}
