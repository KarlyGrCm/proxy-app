package middleware

import "github.com/kataras/iris"

// struct to define the request
type Request struct {
	Domain   string
	Weigth   int
	Priority int
}

type Repository interface {
	Read() []*Request
}

var Queue []*Request

func MockQueue() []*Request {
	return []*Request{
		{
			Domain:   "alpha",
			Weigth:   5,
			Priority: 5,
		},
		{
			Domain:   "beta",
			Weigth:   4,
			Priority: 1,
		},
		{
			Domain:   "omega",
			Weigth:   4,
			Priority: 4,
		},
	}
}

func (r *Request) Read() []*Request {
	return MockQueue()
}

func InitQueue() {
	Queue = append(Queue, &Request{})
}

func ProxyMiddleware(context iris.Context) {
	domain := context.GetHeader("domain")
	repo := Repository{}
	repo = &Request{}

}
