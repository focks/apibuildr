package apibuildr

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
)

func ReadRequestBody(ctx context.Context, r *http.Request, out interface{}) error {

	if reflect.TypeOf(out).Kind() != reflect.Pointer {
		return errors.New("invalid placeholder")
	}
	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bodyBytes, out); err != nil {
		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			apiError := ApiError{
				Message:    "bad data",
				ApiCode:    "",
				ApiName:    GetApiName(ctx),
				RequestId:  GetRequestID(ctx),
				Cause:      nil,
				StatusCode: http.StatusBadRequest,
				Errors: map[string]string{
					e.Field: err.Error(),
				},
			}

			return apiError
		case *json.SyntaxError:
			apiError := ApiError{
				Message:    "invalid json data",
				ApiCode:    "",
				ApiName:    GetApiName(ctx),
				RequestId:  GetRequestID(ctx),
				Cause:      nil,
				StatusCode: http.StatusBadRequest,
			}
			return apiError
		}
	}

	return nil
}
