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
	Picture        string `column:"picture" json:"picture"`
	UserId         string `column:"user_id" json:"userId"`
	MessageId      string `column:"message_id" json:"messageId"`
	ModelId        string `column:"model_id" json:"modelId"`
	Role           string `column:"role" json:"role"`
	Content        string `column:"content" json:"content"`
	CreateTime     string `column:"create_time" json:"createTime"`
}

type LLmModel struct {
	Id           string `column:"id" json:"id"`
	Pid          string `column:"pid" json:"pid"`
	UserId       string `column:"user_id" json:"userId"`
	Name         string `column:"name" json:"name"`
	Model        string `column:"model" json:"model"`
	Picture      string `column:"picture" json:"picture"`
	Size         string `column:"size" json:"size"`
	IsDownload   bool   `column:"is_download" json:"isDownload"`
	Digest       string `column:"digest" json:"digest"`
	ModelDetails string `column:"model_details" json:"modelDetails"`
	CreateTime   string `column:"create_time" json:"createTime"`
}

type AppChatKnowledgeFile struct {
	Id         string `column:"id" json:"id"`
	Pid        string `column:"pid" json:"pid"`
	FileName   string `column:"file_name" json:"fileName"`
	FilePath   string `column:"file_path" json:"filePath"`
	FileType   int    `column:"file_type" json:"fileType"`
	IsGen      string `column:"is_gen" json:"isGen"`
	CreateTime string `column:"create_time" json:"createTime"`
}
