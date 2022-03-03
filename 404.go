package apibuildr

import (
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"net/http"
	"strconv"
	"strings"
)

func FourZeroFour() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.RequestURI
		requestId := uuid.NewUUID().String()
		headers := make([]string, 0)
		for k, v := range r.Header {
			header := fmt.Sprintf("%s:%s", k, strings.Join(v, ","))
			headers = append(headers, header)
		}
		nf := ErrorResponse{
			Message:   fmt.Sprintf("requested path %s not present", path),
			Code:      "not found",
			Api:       r.RequestURI,
			RequestId: requestId,
			Headers:   headers,
		}

		response, _ := json.Marshal(&nf)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(response)))
		w.WriteHeader(404)
		_, _ = w.Write(response)
	})
}
