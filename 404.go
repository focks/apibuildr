package apibuildr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FourZeroFour() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.RequestURI
		requestId := uuid.NewString()
		headers := make([]string, 0)
		for k, v := range r.Header {
			header := fmt.Sprintf("%s:%s", k, strings.Join(v, ","))
			headers = append(headers, header)
		}
		nf := ErrorResponse{
			Message:   fmt.Sprintf("requested path %s not present", path),
			ApiCode:   "not found",
			ApiPath:   r.RequestURI,
			RequestId: requestId,
			Headers:   headers,
		}

		response, _ := json.Marshal(&nf)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(response)))
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(response)
	})
}
