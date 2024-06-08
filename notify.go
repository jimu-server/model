package model

type AppNotify struct {
	Id         string `column:"id" json:"id"`
	PubId      string `column:"pub_id" json:"pubId"`       // 发布者 系统发布id为 system之外,其余的为用户id
	SubId      string `column:"sub_id" json:"subId"`       // 订阅者
	Title      string `column:"title" json:"title"`        // 通知主题
	MsgType    int    `column:"msg_type" json:"msgType"`   // 消息类型 1:info 2:success 3:warning 4:error
	Text       string `column:"text" json:"text"`          // 文本信息
	Type       int    `column:"type" json:"type"`          // 未使用
	Param      string `column:"param" json:"param"`        // 消息上下文参数
	Template   string `column:"template" json:"template"`  // 消息模板组件
	Status     int    `column:"status" json:"status"`      // 消息读取状态 0 未读 1 已读
	IsDelete   int    `column:"is_delete" json:"isDelete"` // 删除标识
	CreateTime string `column:"create_time" json:"createTime"`
	UpdateTime string `column:"update_time" json:"updateTime"`
}

func (notify *AppNotify) TextByte() []byte {
	return nil
}

func (notify *AppNotify) TemplateByte() []byte {
	return nil
}
