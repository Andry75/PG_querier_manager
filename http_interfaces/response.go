package http_interfaces

type Response interface {
	GetPayload() string
	GetMassage() string
	MessageEmpty() bool
	GetHttpStatus() int
}
