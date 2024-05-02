package model

import "os"

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
	UserId     string `column:"user_id" json:"userId"`
	Check      bool   `column:"check" json:"check"`
	FileName   string `column:"file_name" json:"fileName"`
	FilePath   string `column:"file_path" json:"filePath"`
	FileType   int    `column:"file_type" json:"fileType"`
	CreateTime string `column:"create_time" json:"createTime"`
}

func (receiver *AppChatKnowledgeFile) GetId() string {
	return receiver.Id
}

func (receiver *AppChatKnowledgeFile) GetPid() string {
	return receiver.Pid
}

func (receiver *AppChatKnowledgeFile) GetName() string {
	return receiver.FileName
}

type AppChatKnowledgeInstance struct {
	Id                   string `column:"id" json:"id"`
	UserId               string `column:"user_id" json:"userId"`
	KnowledgeName        string `column:"knowledge_name" json:"knowledgeName"`
	KnowledgeFiles       string `column:"knowledge_files" json:"knowledgeFiles"`
	KnowledgeDescription string `column:"knowledge_description" json:"knowledgeDescription"`
	KnowledgeType        int    `column:"knowledge_type" json:"knowledgeType"`
	CreateTime           string `column:"create_time" json:"createTime"`
}

type EmbeddingAnalysis struct {
	*AppChatKnowledgeFile
	os.FileInfo
	FileBody []byte `column:"file_body" json:"fileBody"`
	Lines    []string
}
