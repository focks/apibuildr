package apibuildr

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleError(ctx context.Context, w http.ResponseWriter, err error) {
	foul, ok := err.(*ApiFoul)
	if !ok {
		HandleInternalError(ctx, w, err)
		return
	}
	handleApiError(ctx, w, foul)
}

func handleApiError(ctx context.Context, w http.ResponseWriter, err *ApiFoul) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)

	err.RequestId = GetRequestID(ctx)
	err.ApiName = GetApiName(ctx)
	payload, _ := json.Marshal(err)
	if _, er := w.Write(payload); er != nil {
		fmt.Println("something went wrong in the server")
	}
}

func HandleInternalError(ctx context.Context, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	e := ApiFoul{
		Message:    "something went wrong in the server.",
		RequestId:  GetRequestID(ctx),
		StatusCode: http.StatusInternalServerError,
		ApiName:    GetApiName(ctx),
	}
	payload, _ := json.Marshal(e)
	if _, er := w.Write(payload); er != nil {
		fmt.Println(er.Error())
	}
}
