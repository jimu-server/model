package model

type Tool struct {
	Id         string `column:"id" json:"id"`
	Name       string `column:"name" json:"name"`              // 工具名称
	RouterName string `column:"router_name" json:"routerName"` // 路由名称,并且不能重复
	Icon       string `column:"icon" json:"icon"`              // 图标
	IsOpen     string `column:"isOpen" json:"isOpen"`          //
	Type       int    `column:"type" json:"type"`              // 工具类型
	Component  string `column:"component" json:"component"`    // 工具对应窗口组件
	Btn        string `column:"btn" json:"btn"`                // 工具对应按钮
	Ws         string `column:"ws" json:"ws"`                  // 工具对应websocket
	Pull       string `column:"pull" json:"pull"`              // 工具对应组件消息拉取
	Path       string `column:"path" json:"path"`              // 工具路径 /{basePath}/{name}
	Tip        string `column:"tip" json:"tip"`                // 提示语,一般填写工具名称
	Position   int    `column:"position" json:"position"`      // 工具按钮位置信息 1:左侧 2:右侧
	CreateTime string `column:"create_time" json:"createTime"` // 成立时间s
}
