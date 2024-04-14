package model

type AppChatConversationItem struct {
	Id         string `column:"id" json:"id"`
	Picture    string `column:"picture" json:"picture"`
	UserId     string `column:"user_id" json:"userId"`
	Title      string `column:"title" json:"title"`
	LastModel  string `column:"last_model" json:"lastModel"`
	LastMsg    string `column:"last_msg" json:"lastMsg"`
	LastTime   string `column:"last_time" json:"lastTime"`
	CreateTime string `column:"create_time" json:"createTime"`
}
