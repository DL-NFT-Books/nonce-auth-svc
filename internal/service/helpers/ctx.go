package helpers

import (
	"context"
	"net/http"

	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/config"
	"gitlab.com/tokend/nft-books/nonce-auth-svc/internal/data"

	networkConnector "github.com/dl-nft-books/network-svc/connector"
	"gitlab.com/distributed_lab/logan/v3"
	doormanConnector "gitlab.com/tokend/nft-books/doorman/connector"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	dbCtxKey
	serviceConfigCtxKey
	doormanConnectorCtxKey
	networkConnectorCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}
func CtxDB(entry data.MasterQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCtxKey, entry)
	}
}
func DB(r *http.Request) data.MasterQ {
	return r.Context().Value(dbCtxKey).(data.MasterQ).New()
}
func CtxServiceConfig(entry *config.ServiceConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, serviceConfigCtxKey, entry)
	}
}
func ServiceConfig(r *http.Request) *config.ServiceConfig {
	return r.Context().Value(serviceConfigCtxKey).(*config.ServiceConfig)
}
func CtxDoormanConnector(entry doormanConnector.ConnectorI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, doormanConnectorCtxKey, entry)
	}
}
func DoormanConnector(r *http.Request) doormanConnector.ConnectorI {
	return r.Context().Value(doormanConnectorCtxKey).(doormanConnector.ConnectorI)
}

func CtxNetworkConnector(entry networkConnector.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, networkConnectorCtxKey, entry)
	}
}
func NetworkConnector(r *http.Request) networkConnector.Connector {
	return r.Context().Value(networkConnectorCtxKey).(networkConnector.Connector)
}
