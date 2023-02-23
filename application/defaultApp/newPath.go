package defaultApp

import (
	"context"
	currencyApi "github.com/guybrand/currency/api"
	"github.com/guybrand/exchange/domain/entity"
	. "github.com/orchestd/servicereply"
)

func (app exchangeApp) NewPath(c context.Context, id string) (entity.Trs, ServiceReply) {
	if q, err := app.currencyApi.NewPath(c, currencyApi.NewPathRequest{Id: id}); err != nil {
		return entity.Trs{}, nil
	} else {
		return entity.Trs{
			Id:              id,
			NewPathResponse: q,
		}, nil
	}
}
