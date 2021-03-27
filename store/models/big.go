package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/3dw1nM0535/galva/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// Biggy stores large integers and can deserialize variety of input
type Biggy big.Int

// NewBiggy returns an instance of Big from *big.Int
func NewBiggy(i *big.Int) *Biggy {
	if i != nil {
		b := Biggy(*i)
		return &b
	}
	return nil
}

// NewBiggyI return an instance of Big from int64
func NewBiggyI(i int64) *Biggy {
	return NewBiggy(big.NewInt(i))
}

// MarshalText marshals instance of Biggy to base 10 number as string
func (b *Biggy) MarshalText() ([]byte, error) {
	return []byte((*big.Int)(b).Text(10)), nil
}

// MarshalJSON marshal *Biggy instance to base 10 number as string
func (b *Biggy) MarshalJSON() ([]byte, error) {
	text, err := b.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

// UnmarshalText implements the encoding.TextUnmarshaler
func (b *Biggy) UnmarshalText(input []byte) error {
	input = utils.RemoveQuotes(input)
	str := string(input)
	if utils.HasHexPrefix(str) {
		decoded, err := hexutil.DecodeBig(str)
		if err != nil {
			return err
		}
		*b = Biggy(*decoded)
		return nil
	}

	_, ok := b.setString(str, 10)
	if !ok {
		return fmt.Errorf("unable to convert %s to Biggy", str)
	}
	return nil
}

func (b *Biggy) setString(s string, base int) (*Biggy, bool) {
	w, ok := (*big.Int)(b).SetString(s, base)
	return (*Biggy)(w), ok
}

// UnmarshalJSON implements encoding.JSONUnmarshaller
func (b *Biggy) UnmarshalJSON(input []byte) error {
	return b.UnmarshalText(input)
}

// Value returns *Biggy instance serialized for database
func (b Biggy) Value() (driver.Value, error) {
	return b.String(), nil
}

// Scan reads the database value and returns an instance of *Biggy
func (b *Biggy) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		decoded, ok := b.setString(v, 10)
		if !ok {
			return fmt.Errorf("unable to set string %v of %T to base 10 big.Int for Biggy", value, value)
		}
		*b = *decoded
	case []uint8:
		// SQL Query returns numeric() as []uint8 of string representation
		decoded, ok := b.setString(string(v), 10)
		if !ok {
			return fmt.Errorf("unable to set string %v of %T to base 10 big.Int for Biggy", value, value)
		}
		*b = *decoded
	default:
		return fmt.Errorf("unable to convert %v of %T to Biggy", value, value)
	}
	return nil
}

// ToInt converts b to big.Int
func (b *Biggy) ToInt() *big.Int {
	return (*big.Int)(b)
}

// String returns the base 10 encoding of b
func (b *Biggy) String() string {
	return b.ToInt().Text(10)
}

// Hex returns the hex encoding of b
func (b *Biggy) Hex() string {
	return hexutil.EncodeBig(b.ToInt())
}
