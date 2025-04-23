package client

import (
	"context"

	grpcdef "github.com/sologenic/com-fs-aed-model"
	grpcclient "github.com/sologenic/com-fs-utils-lib/go/grpc-client"
)

const endpoint = "AED_STORE"

var client grpcdef.AEDServiceClient
var grpcClient *grpcclient.GRPCClient

/*
Initialize the client.
Depending on the parameter, the environment is determined to be either in cluster of local by:
localhost:port => local
localhost => No port is not local
*/
func initClient() {
	grpcClient = grpcclient.InitClient(endpoint)
	cl := grpcdef.NewAEDServiceClient(grpcClient.Conn)
	client = cl
}

func Client() grpcdef.AEDServiceClient {
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
