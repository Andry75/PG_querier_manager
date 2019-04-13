package http_interfaces

type Query interface {
	getEndpointName() string
	getParams() string
}
