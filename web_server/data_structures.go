package web_server

type Query struct {
	endpointName string
	payload      string
	method       string
}

func (q Query) GetEndpointName() string {
	return q.endpointName
}
func (q Query) GetPayload() string {
	return q.payload
}

func (q Query) GetMethod() string {
	return q.method
}
