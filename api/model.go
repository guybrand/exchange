/* This file is auto-generated and should not be modified */

package api

import (
	"github.com/guybrand/exchange/domain/entity"
)

type Trs entity.Trs

type NewPathRequest struct {
	Id string `json:"id" validate:"required"`
}

type NewPathResponse Trs

const (
	NewPathMethod = "newPath"
)
