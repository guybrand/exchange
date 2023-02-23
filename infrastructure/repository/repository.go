package repository

import (
	"github.com/guybrand/exchange/domain/repository"
)

type exchangeRepository struct {
}

func NewExchangeRepository() repository.ExchangeRepository {
	return exchangeRepository{}
}
