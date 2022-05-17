package base

type Input struct {
	Query map[string]string

	Header map[string]string
	Body   []byte
}
