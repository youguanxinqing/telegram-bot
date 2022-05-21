package message

type Chat struct {
	FirstName string `json:"first_name"`
	Id        int64
	Type      string
}

type From struct {
	FirstName    string `json:"first_name"`
	Id           int64
	IsBot        bool
	LanguageCode string `json:"language_code"`
}

// HookMsg hook 回调消息结构.
type HookMsg struct {
	Message struct {
		Chat      Chat
		Date      int64
		From      From
		MessageId int `json:"message_id"`
		Text      string
	}
	UpdateId int64 `json:"update_id"`
}
