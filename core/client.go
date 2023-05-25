package core

type MessageFromClient struct {
	Code               int    `json:"code"`
	Message            string `json:"message"`
	CallBackUrl        string `json:"call-back-url"`
	CallBackMethod     string `json:"call-back-method"`
	CallBackPluginName string `json:"call-back-plugin-name"`
}
