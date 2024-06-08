package model

type Router struct {
	Id         string `column:"id" json:"id"`                  // id
	Pid        string `column:"pid" json:"pid"`                // 父节点
	Title      string `column:"title" json:"title"`            // 标题
	Icon       string `column:"icon" json:"icon"`              // 图标组件
	Name       string `column:"name" json:"name"`              // 路由名称
	Component  string `column:"component" json:"component"`    // 路由组件
	Path       string `column:"path" json:"path"`              // 路由路径
	Remark     string `column:"remark" json:"remark"`          // 备注
	Status     bool   `column:"status" json:"status"`          //状态
	RouterType int    `column:"router_type" json:"routerType"` // 路由类型
	Sort       int    `column:"sort" json:"sort"`              // 排序
	ToolId     string `column:"tool_id" json:"toolId"`         // 工具id
	CreateTime string `column:"create_time" json:"createTime"` // 成立时间
}

func (receiver *Router) GetId() string {
	return receiver.Id
}

func (receiver *Router) GetPid() string {
	return receiver.Pid
}

func (receiver *Router) GetName() string {
	return receiver.Name
}
