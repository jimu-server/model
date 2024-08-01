package model

type AppNotify struct {
	Id         string `gorm:"column:id;primaryKey" column:"id" json:"id"`
	PubId      string `gorm:"column:pub_id" column:"pub_id" json:"pubId"`          // 发布者 系统发布id为 system之外,其余的为用户id
	SubId      string `gorm:"column:sub_id" column:"sub_id" json:"subId"`          // 订阅者
	Title      string `gorm:"column:title" column:"title" json:"title"`            // 通知主题
	MsgType    int    `gorm:"column:msg_type" column:"msg_type" json:"msgType"`    // 消息类型 1:info 2:success 3:warning 4:error
	Text       string `gorm:"column:text" column:"text" json:"text"`               // 文本信息
	Type       int    `gorm:"column:type" column:"type" json:"type"`               // 未使用
	Param      string `gorm:"column:param" column:"param" json:"param"`            // 消息上下文参数
	Template   string `gorm:"column:template" column:"template" json:"template"`   // 消息模板组件
	Status     int    `gorm:"column:status" column:"status" json:"status"`         // 消息读取状态 0 未读 1 已读
	IsDelete   int    `gorm:"column:is_delete" column:"is_delete" json:"isDelete"` // 删除标识
	CreateTime string `gorm:"column:create_time" column:"create_time" json:"createTime"`
	UpdateTime string `gorm:"column:update_time" column:"update_time" json:"updateTime"`
}

func (notify *AppNotify) TextByte() []byte {
	return nil
}

func (notify *AppNotify) TemplateByte() []byte {
	return nil
}

func (notify *AppNotify) TableName() string {
	return "app_notify"
}
