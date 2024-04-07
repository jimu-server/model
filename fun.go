package model

// FuncRouter 功能路由
type FunRouter struct {
	Id         string `column:"id" json:"id"`
	Method     string `column:"method" json:"method"`          // 接口类型
	Name       string `column:"name" json:"name"`              // 功能名称
	Path       string `column:"path" json:"path"`              // 功能路径
	Status     bool   `column:"status" json:"status"`          // 状态
	CreateTime string `column:"create_time" json:"createTime"` // 成立时间
}
