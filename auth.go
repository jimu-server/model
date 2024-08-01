package model

type AuthUserRole struct {
	Id          string `gorm:"column:id;primaryKey" column:"id" json:"id,omitempty"`
	UserId      string `gorm:"column:user_id" column:"user_id" json:"userId,omitempty"`
	RoleId      string `gorm:"column:role_id" column:"role_id" json:"roleId,omitempty"`
	OrgId       string `gorm:"column:org_id" column:"org_id" json:"orgId,omitempty"`
	FirstChoice bool   `gorm:"column:first_choice" column:"first_choice" json:"firstChoice,omitempty"`
	CreateTime  string `gorm:"column:create_time" column:"create_time" json:"createTime,omitempty"`
}

func (receiver *AuthUserRole) TableName() string {
	return "app_org_user_role"
}

// AuthTool 组织用户角色工具栏授权关系
type AuthTool struct {
	Id         string `gorm:"column:id;primaryKey" column:"id" json:"id,omitempty"`
	OrgId      string `gorm:"column:org_id" column:"org_id" json:"orgId,omitempty"`                // 组织
	UserId     string `gorm:"column:user_id" column:"user_id" json:"userId,omitempty"`             // 用户
	RoleId     string `gorm:"column:role_id" column:"role_id" json:"roleId,omitempty"`             // 角色
	ToolId     string `gorm:"column:tool_id" column:"tool_id" json:"toolId,omitempty"`             // 工具
	Status     bool   `gorm:"column:status" column:"status" json:"status" json:"status,omitempty"` // 是否启用
	CreateTime string `gorm:"column:create_time" column:"create_time" json:"createTime" json:"createTime,omitempty"`
}

func (receiver *AuthTool) TableName() string {
	return "app_org_user_role_tool_auth"
}

// AuthRouter 组织用户角色路由授权关系
type AuthRouter struct {
	Id         string `gorm:"column:id;primaryKey" column:"id" json:"id,omitempty"`
	OrgId      string `gorm:"column:org_id" column:"org_id" json:"orgId,omitempty"`                // 组织
	UserId     string `gorm:"column:user_id" column:"user_id" json:"userId,omitempty"`             // 用户
	RoleId     string `gorm:"column:role_id" column:"role_id" json:"roleId,omitempty"`             // 角色
	ToolId     string `gorm:"column:tool_id" column:"tool_id" json:"toolId,omitempty"`             // 工具
	RouterId   string `gorm:"column:router_id" column:"router_id" json:"routerId,omitempty"`       // 路由
	Status     bool   `gorm:"column:status" column:"status" json:"status" json:"status,omitempty"` // 是否启用
	CreateTime string `gorm:"column:create_time" column:"create_time" json:"createTime" json:"createTime,omitempty"`
}

func (receiver *AuthRouter) TableName() string {
	return "app_org_user_role_router_auth"
}

// AuthFun 组织用户角色授权关系
type AuthFun struct {
	Id         string `gorm:"column:id;primaryKey" column:"id" json:"id,omitempty"`
	OrgId      string `gorm:"column:org_id" column:"org_id" json:"orgId,omitempty"`                // 组织
	UserId     string `gorm:"column:user_id" column:"user_id" json:"userId,omitempty"`             // 用户
	RoleId     string `gorm:"column:role_id" column:"role_id" json:"roleId,omitempty"`             // 角色
	FunId      string `gorm:"column:fun_id" column:"fun_id" json:"funId,omitempty"`                // 功能
	Status     bool   `gorm:"column:status" column:"status" json:"status" json:"status,omitempty"` // 是否启用
	CreateTime string `gorm:"column:create_time" column:"create_time" json:"createTime" json:"createTime,omitempty"`
}

func (receiver *AuthFun) TableName() string {
	return "app_org_user_role_router_auth"
}

type AuthRoleTool struct {
	Id     string `gorm:"column:id;primaryKey" column:"id" json:"id,omitempty"`
	OrgId  string `gorm:"column:org_id" column:"org_id" json:"orgId,omitempty"`                // 组织
	RoleId string `gorm:"column:role_id" column:"role_id" json:"roleId,omitempty"`             // 角色
	ToolId string `gorm:"column:tool_id" column:"tool_id" json:"toolId,omitempty"`             // 工具
	Status bool   `gorm:"column:status" column:"status" json:"status" json:"status,omitempty"` // 是否启用
}

func (receiver *AuthFun) AuthRoleTool() string {
	return "app_org_role_tool_auth"
}

type AuthRoleRouter struct {
	Id       string `gorm:"column:id;primaryKey" column:"id" json:"id,omitempty"`
	OrgId    string `gorm:"column:org_id" column:"org_id" json:"orgId,omitempty"`                // 组织
	RoleId   string `gorm:"column:role_id" column:"role_id" json:"roleId,omitempty"`             // 角色
	ToolId   string `gorm:"column:tool_id" column:"tool_id" json:"toolId,omitempty"`             // 工具
	RouterId string `gorm:"column:router_id" column:"router_id" json:"routerId,omitempty"`       // 路由
	Status   bool   `gorm:"column:status" column:"status" json:"status" json:"status,omitempty"` // 是否启用
}

func (receiver *AuthRoleRouter) AuthRoleTool() string {
	return "app_org_role_router_auth"
}
