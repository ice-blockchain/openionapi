package config

import (
	"github.com/ice-blockchain/iongo"
)

func ElectorAddress() tongo.AccountID {
	// TODO: read from the blockchain config
	return tongo.MustParseAccountID("Ef8zMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzMzM0vF")
}
