package model

type AppSetting struct {
	Id         string `column:"id" json:"id"`
	Pid        string `column:"pid" json:"pid"`
	Name       string `column:"name" json:"name"`
	Value      string `column:"value" json:"value"`
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
