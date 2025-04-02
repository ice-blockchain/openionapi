package litestorage

import (
	"github.com/ice-blockchain/openionapi/pkg/oas"
	"github.com/ice-blockchain/iongo/ton"
	"golang.org/x/net/context"
)

func (s *LiteStorage) GetMissedEvents(ctx context.Context, account ton.AccountID, lt uint64, limit int) ([]oas.AccountEvent, error) {
	return nil, nil
}
