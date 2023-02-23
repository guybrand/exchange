/* This file is auto-generated and should not be modified */

package http

import (
	"context"
	"github.com/guybrand/exchange/api"
	"github.com/orchestd/dependencybundler/interfaces/transport"
	. "github.com/orchestd/servicereply"
)

const serviceName = "exchange"

type exchangeHTTPClient struct {
	client transport.HttpClient
}

func NewExchangeApiHTTPClient(client transport.HttpClient) api.ExchangeApi {
	return exchangeHTTPClient{client: client}
}

func (h exchangeHTTPClient) NewPath(c context.Context, req api.NewPathRequest) (api.NewPathResponse, ServiceReply) {
	var res api.NewPathResponse
	if reply := h.client.Call(c, req, serviceName, api.NewPathMethod, &res, nil); !reply.IsSuccess() {
		return res, reply
	}
	return res, nil
}
