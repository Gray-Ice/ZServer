package core

type CommonMessage struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	PluginName string `json:"call-back-name"`
}
