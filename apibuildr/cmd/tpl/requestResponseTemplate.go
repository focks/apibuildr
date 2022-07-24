package tpl

var RequestResponseTemplate = `package api

type {{ .Name }}ApiRequest struct {
	Name string 
}

type {{ .Name }}ApiResponse struct {
	Status string 
}
`
