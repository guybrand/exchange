/* This file is auto-generated and should not be modified */

package monolith

import (
	"context"
	"github.com/guybrand/exchange/api"
	"github.com/guybrand/exchange/api/server/monolith"
	. "github.com/orchestd/servicereply"
)

type ExchangeMonolithClient struct {
	MonolithInterface monolith.ExchangeInterface
}

func NewExchangeMonolithClient(monolithInterface monolith.ExchangeInterface) api.ExchangeApi {
	return ExchangeMonolithClient{MonolithInterface: monolithInterface}
}

func (p ExchangeMonolithClient) NewPath(c context.Context, req api.NewPathRequest) (api.NewPathResponse, ServiceReply) {
	return p.MonolithInterface.NewPath(c, req)
}
