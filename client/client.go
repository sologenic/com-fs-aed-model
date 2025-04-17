package client

import (
	"context"

	grpcdef "github.com/sologenic/com-fs-asset-time-series-model"
	grpcclient "github.com/sologenic/com-fs-utils-lib/go/grpc-client"
)

const endpoint = "OHLC_STORE"

var client *grpcdef.AssetTimeSeriesServiceClient
var grpcClient *grpcclient.GRPCClient

/*
Initialize the client.
Depending on the parameter, the environment is determined to be either in cluster of local by:
localhost:port => local
localhost => No port is not local
*/
func initClient() {
	grpcClient = grpcclient.InitClient(endpoint)

	cl := grpcdef.NewAssetTimeSeriesServiceClient(grpcClient.Conn)
	client = &cl
}

func Client() *grpcdef.AssetTimeSeriesServiceClient {
	if client == nil {
		initClient()
	}
	return client
}

func AuthCtx(ctx context.Context) context.Context {
	if grpcClient == nil {
		initClient()
	}
	return grpcClient.AuthCtx(ctx)
}
