package errorHelper

type ErrResponse struct {
	Code Message
}

type Message struct {
	MessageFE string
	MessageBE string
	Code      string
}
