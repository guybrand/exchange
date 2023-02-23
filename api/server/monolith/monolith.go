/* This file is auto-generated and should not be modified */

package monolith

import (
	"context"
	"github.com/guybrand/exchange/api"
	"github.com/guybrand/exchange/domain/application"
	"github.com/orchestd/dependencybundler/interfaces/validations"
	. "github.com/orchestd/servicereply"
)

type ExchangeInterface struct {
	exchangeApp application.ExchangeApp
	validation  validations.Validations
}

func NewExchangeInterface(exchangeApp application.ExchangeApp, validation validations.Validations) ExchangeInterface {
	return ExchangeInterface{exchangeApp: exchangeApp, validation: validation}
}

func (i ExchangeInterface) NewPath(c context.Context, req api.NewPathRequest) (api.NewPathResponse, ServiceReply) {
	if vErr := i.validation.Validate(req); vErr != nil && !vErr.IsSuccess() {
		return api.NewPathResponse{}, vErr
	}
	res, err := i.exchangeApp.NewPath(c, req.Id)
	return api.NewPathResponse(res), err
}
