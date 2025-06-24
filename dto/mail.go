package dto

type MailNotificationMessage struct {
	Subject  string   `json:"subject"`
	Receiver []string `json:"receiver"`
	MsgType  string   `json:"message_type"`
	Message  string   `json:"message"`
}
