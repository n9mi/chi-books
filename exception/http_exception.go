package exception

type HTTPError struct {
	StatusCode int
	Message    string
}

// overload func
func (e *HTTPError) Error() string {
	return e.Message
}
