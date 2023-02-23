/* This file is auto-generated and should not be modified */

package application

import (
	"context"
	"github.com/guybrand/exchange/domain/entity"
	. "github.com/orchestd/servicereply"
)

type ExchangeApp interface {
	NewPath(c context.Context, id string) (entity.Trs, ServiceReply)
}
