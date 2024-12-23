package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

type CreateWalletRequest struct {
	Address string `json:"address"`
	Network string `json:"network"`
}

func (r CreateWalletRequest) Validate() error {
	fields := []*validation.FieldRules{
		validation.Field(&r.Network, validation.Required),
		validation.Field(&r.Address, validation.Required),
	}

	return errors.Wrap(validation.ValidateStruct(&r, fields...), "wallet create validation error")
}
