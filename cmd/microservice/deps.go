package main

import (
	. "github.com/guybrand/currency/api/client/http"
	"github.com/guybrand/exchange/api/server/monolith"
	"github.com/guybrand/exchange/application/defaultApp"
)

func deps() []interface{} {
	return append(internalDeps(), externalDeps()...)
}

func internalDeps() []interface{} {
	return []interface{}{
		defaultApp.NewExchangeApp,
		monolith.NewExchangeInterface,
	}
}

func externalDeps() []interface{} {
	return []interface{}{
		NewCurrencyApiHTTPClient,
	}
}
