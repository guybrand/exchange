/* This file is auto-generated and should not be modified */

package api

import (
	"context"
	. "github.com/orchestd/servicereply"
)

type ExchangeApi interface {
	NewPath(c context.Context, req NewPathRequest) (NewPathResponse, ServiceReply)
}
