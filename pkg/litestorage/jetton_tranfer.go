package litestorage

import (
	"context"
	"fmt"

	"github.com/ice-blockchain/openionapi/pkg/core"
	"github.com/ice-blockchain/iongo/ton"
)

func (s *LiteStorage) GetJettonTransferPayload(ctx context.Context, accountID, jettonMaster ton.AccountID) (*core.JettonTransferPayload, error) {
	return nil, fmt.Errorf("not implemented")
}
