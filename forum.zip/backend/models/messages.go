package models

type Message struct {
	Message  string `json:"message"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Created  string `json:"created"`
	Status   string `json:"status"`
}
type Chat struct {
	User1 string `json:"user1"`
	User2 string `json:"user2"`
}
