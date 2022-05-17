package base

type Output struct {
	Header map[string]string
	Body   []byte

	StatusCode int

	Err error
}
