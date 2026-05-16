package api

// Chat is a normalized conversation (Graph or chatsvc).
type Chat struct {
	ID              string `json:"id"`
	Topic           string `json:"topic"`
	ChatType        string `json:"chatType"`
	CreatedDateTime string `json:"createdDateTime"`
}

// Message is a normalized chat message.
type Message struct {
	ID              string      `json:"id"`
	CreatedDateTime string      `json:"createdDateTime"`
	From            MessageFrom `json:"from"`
	Body            MessageBody `json:"body"`
}

type MessageFrom struct {
	User User `json:"user"`
}

type User struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}

type MessageBody struct {
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
}

// Messenger is implemented by Graph and chatsvc clients.
type Messenger interface {
	Doctor() (map[string]interface{}, error)
	ListChats() ([]Chat, error)
	GetMessages(chatID string, limit int) ([]Message, error)
	Backend() string
}
