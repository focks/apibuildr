package cmd

type ApiFile struct {
	path     string
	template []byte
	err      error
	created  bool
}
