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

type AppChatMessage struct {
	Id             string `column:"id" json:"id"`
	ConversationId string `column:"conversation_id" json:"conversationId"`
	UserId         string `column:"user_id" json:"userId"`
	Role           string `column:"role" json:"role"`
	Content        string `column:"content" json:"content"`
	CreateTime     string `column:"create_time" json:"createTime"`
}

type LLmModel struct {
	Id         string `column:"id" json:"id"`
	Name       string `column:"name" json:"name"`
	Model      string `column:"model" json:"model"`
	Icon       string `column:"icon" json:"icon"`
	Parameters string `column:"parameters" json:"parameters"`
	Size       string `column:"size" json:"size"`
	Download   string `column:"download" json:"download"`
	CreateTime string `column:"create_time" json:"createTime"`
}
