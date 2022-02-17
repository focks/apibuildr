package apibuildr

type ErrorResponse struct {
	Message    string   `json:"message"`
	Code       string   `json:"error"`
	ResponseId string   `json:"response-id"`
	Api        string   `json:"api"`
	Headers    []string `json:"headers"`
}
