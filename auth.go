package model

type AuthUserRole struct {
	Id          string `column:"id"`
	UserId      string `column:"user_id"`
	RoleId      string `column:"role_id"`
	OrgId       string `column:"org_id"`
	FirstChoice bool   `column:"first_choice"`
	CreateTime  string `column:"create_time"`
}

// AuthTool 组织用户角色工具栏授权关系
type AuthTool struct {
	Id         string `column:"id"`
	OrgId      string `column:"org_id"`  // 组织
	UserId     string `column:"user_id"` // 用户
	RoleId     string `column:"role_id"` // 角色
	ToolId     string `column:"tool_id"` // 工具
	Status     int    `column:"status"`  // 状态
	CreateTime string `column:"create_time" json:"createTime"`
}

// AuthRouter 组织用户角色路由授权关系
type AuthRouter struct {
	Id         string `column:"id"`
	OrgId      string `column:"org_id"`    // 组织
	UserId     string `column:"user_id"`   // 用户
	RoleId     string `column:"role_id"`   // 角色
	ToolId     string `column:"tool_id"`   // 工具
	RouterId   string `column:"router_id"` // 路由
	Status     int    `column:"status"`    // 状态
	CreateTime string `column:"create_time" json:"createTime"`
}

// AuthFun 组织用户角色授权关系
type AuthFun struct {
	Id         string `column:"id"`
	OrgId      string `column:"org_id"`  // 组织
	UserId     string `column:"user_id"` // 用户
	RoleId     string `column:"role_id"` // 角色
	FunId      string `column:"fun_id"`  // 功能
	Status     int    `column:"status"`  // 状态
	CreateTime string `column:"create_time" json:"createTime"`
}

type AuthRoleTool struct {
	Id     string `column:"id"`
	OrgId  string `column:"org_id"`  // 组织
	RoleId string `column:"role_id"` // 角色
	ToolId string `column:"tool_id"` // 工具
}

type AuthRoleRouter struct {
	Id       string `column:"id"`
	OrgId    string `column:"org_id"`    // 组织
	RoleId   string `column:"role_id"`   // 角色
	ToolId   string `column:"tool_id"`   // 工具
	RouterId string `column:"router_id"` // 路由
}
