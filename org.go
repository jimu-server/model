package model

type Org struct {
	Id         string `column:"id" json:"id"`
	Pid        string `column:"pid" json:"pid"`                // 上级
	Name       string `column:"name" json:"name"`              // 名称
	CreateTime string `column:"create_time" json:"createTime"` // 成立时间
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

type OrgUserRole struct {
	ID string ` column:"id"json:"id"`
	// 角色所属组织id
	OrgID string ` column:"org_id"json:"orgId"`
	// 角色id
	RoleID string `column:"role_id"json:"roleId"`
	// 用户id
	UserID string `column:"user_id"json:"userId"`
	// 创建时间
	CreateTime string `column:"create_time"json:"createTime"`
}
