package apibuildr

type ApiFoul struct {
	ApiName    string `json:"api_name"`
	Message    string `json:"message"`
	Cause      error  `json:"-"`
	RequestId  string `json:"request_id"`
	StatusCode int    `json:"status_code"`
	DomainCode string `json:"domain_codes"`
}

// Deprecated
func NewFoul(msg string) *ApiFoul {
	return &ApiFoul{
		Message: msg,
	}
}

func New(msg string) *ApiFoul {
	return &ApiFoul{
		Message: msg,
	}
}

func (e *ApiFoul) UnWrap() error {
	return e.Cause
}

func (f ApiFoul) Error() string {
	return f.Message
}

func (f *ApiFoul) WithCause(cause error) *ApiFoul {
	f.Cause = cause
	return f
}

func (f *ApiFoul) WithApiName(name string) *ApiFoul {
	f.ApiName = name
	return f
}

func (f *ApiFoul) WithStatusCode(status int) *ApiFoul {
	f.StatusCode = status
	return f
}

func (f *ApiFoul) WithDomainCode(code string) *ApiFoul {
	f.DomainCode = code
	return f
}

func (f *ApiFoul) WithMessage(msg string) *ApiFoul {
	f.Message = msg
	return f
}
