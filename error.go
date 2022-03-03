package apibuildr

type ErrorResponse struct {
	Message   string   `json:"message"`
	Code      string   `json:"error"`
	RequestId string   `json:"request-id"`
	Api       string   `json:"api"`
	Headers   []string `json:"headers"`
}
