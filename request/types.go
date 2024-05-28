package request

type Request struct {
	Method  string
	Target  string
	Headers map[string]string
	Body    string
}
