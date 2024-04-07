package model

type Role struct {
	Id         string `column:"id" json:"id"`
	Name       string `column:"name" json:"name"`              // 角色名称
	RoleKey    string `column:"role_key" json:"roleKey"`       // 角色标识
	CreateTime string `column:"create_time" json:"createTime"` // 成立时间
}
