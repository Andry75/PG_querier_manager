package web_server

type Query struct {
	endpointName string
	payload      string
}

func (q Query) GetEndpointName() string {
	return q.endpointName
}
func (q Query) GetPayload() string {
	return q.payload
}
