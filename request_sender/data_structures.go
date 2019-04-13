package request_sender

type Response struct {
	httpStatus int
	message    string
	payload    string
}

func (r Response) GetPayload() string {
	return r.payload
}

func (r Response) GetMassage() string {
	return r.message
}

func (r Response) MessageEmpty() bool {
	return len(r.message) == 0
}

func (r Response) PayloadEmpty() bool {
	return len(r.payload) == 0
}

func (r Response) GetHttpStatus() int {
	return r.httpStatus
}
