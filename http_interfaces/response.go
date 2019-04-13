package http_interfaces

type Response interface {
	getData() string
	getMassage() string
	getHttpStatus() int
}
