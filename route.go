package model

type Router struct {
	Id         string `gorm:"column:id;primaryKey" column:"id" json:"id"`                // id
	Pid        string `gorm:"column:org_id" column:"pid" json:"pid"`                     // 父节点
	Title      string `gorm:"column:title" column:"title" json:"title"`                  // 标题
	Icon       string `gorm:"column:icon" column:"icon" json:"icon"`                     // 图标组件
	Name       string `gorm:"column:name" column:"name" json:"name"`                     // 路由名称
	Component  string `gorm:"column:component" column:"component" json:"component"`      // 路由组件
	Path       string `gorm:"column:path" column:"path" json:"path"`                     // 路由路径
	Remark     string `gorm:"column:remark" column:"remark" json:"remark"`               // 备注
	Status     bool   `gorm:"column:status" column:"status" json:"status"`               //状态
	RouterType int    `gorm:"column:router_type" column:"router_type" json:"routerType"` // 路由类型
	Sort       int    `gorm:"column:sort" column:"sort" json:"sort"`                     // 排序
	ToolId     string `gorm:"column:tool_id" column:"tool_id" json:"toolId"`             // 工具id
	CreateTime string `gorm:"column:create_time" column:"create_time" json:"createTime"` // 成立时间
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

func (receiver *Router) TableName() string {
	return "app_router"
}
