package cmd

type ApiFile struct {
	path     string
	template string
	err      error
	created  bool
}
