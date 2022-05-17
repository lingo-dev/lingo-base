package base

type Context struct {
	// database kv ...

	Envs    map[string]string
	Secrets map[string]string
}
