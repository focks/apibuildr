package apibuildr

type ErrorResponse struct {
	Message   string   `json:"message"`
	ApiCode   string   `json:"error"`
	RequestId string   `json:"request-id"`
	ApiPath   string   `json:"api"`
	Headers   []string `json:"headers"`
}
