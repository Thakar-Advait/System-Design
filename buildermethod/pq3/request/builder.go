package request

type request struct {
	method  string
	url     string
	headers map[string]string
	body    string
}

type requestBuilder struct {
	method  string
	url     string
	headers map[string]string
	body    string
}

type Builder interface {
	SetMethod(method string) Builder
	SetURL(url string) Builder
	AddHeader(key, val string) Builder
	SetBody(body string) Builder
	Build() request
}

func NewRequestBuilder() *requestBuilder {
	return &requestBuilder{
		headers: make(map[string]string),
	}
}

func (r *requestBuilder) SetMethod(method string) Builder {
	r.method = method
	return r
	// panic("to be implemented")
}

func (r *requestBuilder) SetURL(url string) Builder {
	r.url = url
	return r
	// panic("to be implemented")
}

func (r *requestBuilder) AddHeader(key, val string) Builder {
	r.headers[key] = val
	return r
	// panic("to be implemented")
}

func (r *requestBuilder) SetBody(body string) Builder {
	r.body = body
	return r
	// panic("to be implemented")
}

func (r *requestBuilder) Build() request {
	return request{
		method:  r.method,
		url:     r.url,
		body:    r.body,
		headers: r.headers,
	}
	// panic("to be implemented")
}
