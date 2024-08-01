package model

type Role struct {
	Id         string `gorm:"column:id;primaryKey" column:"id" json:"id"`
	Name       string `gorm:"column:org_id" column:"name" json:"name"`                   // 角色名称
	RoleKey    string `gorm:"column:role_key" column:"role_key" json:"roleKey"`          // 角色标识
	Status     bool   `gorm:"column:status" column:"status" json:"status"`               // 是否启用
	CreateTime string `gorm:"column:create_time" column:"create_time" json:"createTime"` // 成立时间
}

func (r Role) TableName() string {
	return "app_role"
}
