package model

type Org struct {
	Id         string `gorm:"column:id;primaryKey" column:"id" json:"id"`
	Pid        string `gorm:"column:pid" column:"pid" json:"pid"`                        // 上级
	Name       string `gorm:"column:name" column:"name" json:"name"`                     // 名称
	CreateTime string `gorm:"column:create_time" column:"create_time" json:"createTime"` // 成立时间
}

func (o *Org) GetId() string {
	return o.Id
}
func (o *Org) GetPid() string {
	return o.Pid
}
func (o *Org) GetName() string {
	return o.Name
}
func (o *Org) TableName() string {
	return "app_org"
}

type OrgUserRole struct {
	ID string `gorm:"column:id;primaryKey" column:"id"json:"id"`
	// 角色所属组织id
	OrgID string `gorm:"column:org_id" column:"org_id"json:"orgId"`
	// 角色id
	RoleID string `gorm:"column:role_id" column:"role_id"json:"roleId"`
	// 用户id
	UserID string `gorm:"column:user_id" column:"user_id"json:"userId"`
	// 创建时间
	CreateTime string `gorm:"column:create_time" column:"create_time"json:"createTime"`
}

func (o *OrgUserRole) TableName() string {
	return "app_org_user_role"
}
