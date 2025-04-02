package api

import (
	"context"
	"github.com/ice-blockchain/openionapi/pkg/references"

	"github.com/ice-blockchain/openionapi/pkg/oas"
)

func (h *Handler) GetExtraCurrencyInfo(ctx context.Context, params oas.GetExtraCurrencyInfoParams) (*oas.EcPreview, error) {
	meta := references.GetExtraCurrencyMeta(params.ID)
	return &oas.EcPreview{
		ID:       params.ID,
		Symbol:   meta.Symbol,
		Decimals: meta.Decimals,
		Image:    meta.Image,
	}, nil
}
