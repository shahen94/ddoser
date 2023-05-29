package domain

// RequestResult is a struct for request result
type RequestResult struct {
	StatusCode   int
	ResponseTime int64
	Error        bool
}

type NetworkInterface interface {
	Request(string, chan *RequestResult)
}
