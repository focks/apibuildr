package apibuildr

import (
	"context"

	"github.com/pborman/uuid"
	"go.uber.org/zap"
)

type ContextKey int

const (
	ApiName ContextKey = iota

	Token

	ClientID

	ClientSecret

	RequestID

	UserID

	Host
)

func GetApiName(ctx context.Context) string {
	name, ok := ctx.Value(ApiName).(string)
	if !ok {
		return ""
	}
	return name
}

func GetToken(ctx context.Context) string {
	token, ok := ctx.Value(Token).(string)
	if !ok {
		return ""
	}
	return token
}

func GetClientID(ctx context.Context) string {
	clientID, ok := ctx.Value(ClientID).(string)

	if !ok {
		return ""
	}
	return clientID
}

func GetClientSecret(ctx context.Context) string {
	clientSecret, ok := ctx.Value(ClientSecret).(string)
	if !ok {
		return ""
	}
	return clientSecret
}

func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(RequestID).(string)

	if !ok {
		return ""
	}

	return requestID
}

func GetUserID(ctx context.Context) string {
	userId, ok := ctx.Value(UserID).(string)
	if !ok {
		return ""
	}
	return userId
}

func GetHost(ctx context.Context) string {
	host, ok := ctx.Value(Host).(string)
	if !ok {
		return ""
	}
	return host
}

func Contextual(ctx context.Context, errs ...error) []zap.Field {
	fields := []zap.Field{}

	if requestId := GetRequestID(ctx); requestId != "" {
		fields = append(fields, zap.String("request_id", requestId))
	}

	if api := GetApiName(ctx); api != "" {
		fields = append(fields, zap.String("api_name", api))
	}

	if len(errs) > 0 {
		fields = append(fields, zap.Error(errs[0]))
	}

	return fields
}

func ApiRequestCtx(c context.Context, api string) context.Context {
	requestId := uuid.New()
	ctx := context.WithValue(c, RequestID, requestId)
	ctx = context.WithValue(ctx, ApiName, api)
	return ctx
}
