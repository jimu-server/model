package model

type AppSetting struct {
	Id         string `column:"id" json:"id"`
	Pid        string `column:"pid" json:"pid"`
	UserId     string `column:"user_id" json:"user_id"`
	Name       string `column:"name" json:"name"`
	Value      string `column:"value" json:"value"`
	ToolId     string `column:"tool_id" json:"tool_id"`
	Setting    string `column:"setting" json:"setting"`
	CreateTime string `column:"create_time" json:"create_time"`
}

func (receiver *AppSetting) GetId() string {
	return receiver.Id
}

func (receiver *AppSetting) GetPid() string {
	return receiver.Pid
}

func (receiver *AppSetting) GetName() string {
	return receiver.Name
}

type OllamaSetting struct {
	Host        string `json:"host"`
	OllamaModel string `json:"ollamaModel"`
}
