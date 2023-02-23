/* This file is auto-generated and should not be modified */

package http

import (
	"github.com/guybrand/exchange/api"
	"github.com/guybrand/exchange/api/server/monolith"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/transport"
)

func InitHandlers(router transport.IRouter, m monolith.ExchangeInterface, c configuration.Config) {
	router.POST(api.NewPathMethod, transport.HandleFunc(m.NewPath))
}
