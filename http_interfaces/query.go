package http_interfaces

type Query interface {
	GetEndpointName() string
	GetPayload() string
	GetMethod() string
}
