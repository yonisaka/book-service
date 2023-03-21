package dto

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// WithCode setter response var name
func (r *Response) WithCode(c int) *Response {
	r.Code = c
	return r
}

// WithMessage setter custom message response
func (r *Response) WithMessage(v interface{}) *Response {
	if v != nil {
		r.Message = v
	}
	return r
}

// WithData setter data response
func (r *Response) WithData(v interface{}) *Response {
	r.Data = v
	return r
}
