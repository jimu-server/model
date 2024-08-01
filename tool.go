package model

type Tool struct {
	Id         string `gorm:"column:id;primaryKey" column:"id" json:"id"`
	Name       string `gorm:"column:name" column:"name" json:"name"`                     // 工具名称
	RouterName string `gorm:"column:router_name" column:"router_name" json:"routerName"` // 路由名称,并且不能重复
	Icon       string `gorm:"column:icon" column:"icon" json:"icon"`                     // 图标
	IsOpen     string `gorm:"column:isOpen" column:"isOpen" json:"isOpen"`               //
	Type       int    `gorm:"column:type" column:"type" json:"type"`                     // 工具类型
	Component  string `gorm:"column:component" column:"component" json:"component"`      // 工具对应窗口组件
	Btn        string `gorm:"column:btn" column:"btn" json:"btn"`                        // 工具对应按钮
	Ws         string `gorm:"column:ws" column:"ws" json:"ws"`                           // 工具对应websocket
	Pull       string `gorm:"column:pull" column:"pull" json:"pull"`                     // 工具对应组件消息拉取
	Path       string `gorm:"column:path" column:"path" json:"path"`                     // 工具路径 /{basePath}/{name}
	Tip        string `gorm:"column:tip" column:"tip" json:"tip"`                        // 提示语,一般填写工具名称
	Position   int    `gorm:"column:position" column:"position" json:"position"`         // 工具按钮位置信息 1:左侧 2:右侧
	Status     bool   `gorm:"column:status" column:"status" json:"status"`               // 是否启用
	CreateTime string `gorm:"column:create_time" column:"create_time" json:"createTime"` // 成立时间
}

func (t *Tool) TableName() string {
	return "app_tool"
}
