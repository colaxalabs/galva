package models

import (
	"database/sql/driver"
	"math/big"
)

type BigInt big.Int

func (b *BigInt) Value() (driver.Value, error) {
	if b != nil {
		return b.String(), nil
	}
	return nil, nil
}

func (b *BigInt) Scan(value interface{}) error {
	if value == nil {
		b = nil
	}
	switch t := value.(type) {
	case int64:
		b = (*BigInt)(value.(int64))
	default:
		return fmt.Errorf("Could not scan type %T into bigint", t)
	}
}
