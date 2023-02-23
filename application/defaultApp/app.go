package defaultApp

import (
	currencyApi "github.com/guybrand/currency/api"
	"github.com/guybrand/exchange/domain/application"
)

type exchangeApp struct {
	currencyApi currencyApi.CurrencyApi
}

func NewExchangeApp(currencyApi currencyApi.CurrencyApi) application.ExchangeApp {
	return &exchangeApp{
		currencyApi: currencyApi,
	}
}
