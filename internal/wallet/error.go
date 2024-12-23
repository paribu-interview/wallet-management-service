package wallet

import "github.com/pkg/errors"

var (
	ErrDuplicateWallet = errors.New("wallet already exists")
)
