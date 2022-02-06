package fsserver

import "net/http"

type logResponseProxy struct {
	writer     http.ResponseWriter
	StatusCode int
}

func NewLogResponseProxy(writer http.ResponseWriter) *logResponseProxy {
	return &logResponseProxy{
		writer:     writer,
		StatusCode: http.StatusOK,
	}
}

func (p *logResponseProxy) Header() http.Header {
	return p.writer.Header()
}

func (p *logResponseProxy) Write(data []byte) (int, error) {
	return p.writer.Write(data)
}

func (p *logResponseProxy) WriteHeader(statusCode int) {
	p.writer.WriteHeader(statusCode)
	p.StatusCode = statusCode
}
