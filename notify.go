package model

type AppNotify struct {
	Id         string `column:"id" json:"id"`
	PubId      string `column:"pub_id" json:"pubId"`
	SubId      string `column:"sub_id" json:"subId"`
	Title      string `column:"title" json:"title"`
	MsgType    int    `column:"msg_type" json:"msgType"`
	Text       string `column:"text" json:"text"`
	Status     int    `column:"status" json:"status"`
	Del        int    `column:"del" json:"del"`
	CreateTime string `column:"create_time" json:"createTime"`
	UpdateTime string `column:"update_time" json:"updateTime"`
}

func (notify *AppNotify) TextByte() []byte {
	return nil
}

func (notify *AppNotify) TemplateByte() []byte {
	return nil
}
