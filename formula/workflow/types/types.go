package types

type Message struct {
	ContentType string
	Payload     []byte
}
type MessageError struct {
	Err       error
	ErrorCode int
}