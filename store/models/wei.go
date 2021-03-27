package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/3dw1nM0535/galva/utils"
	"github.com/shopspring/decimal"
	"math/big"
)

// Wei returns the Wei and Size in Wei(blockchain unit)
type Wei big.Int

// NewWei returns a new instance of Wei in wei
func NewWei(c int64) *Wei {
	return (*Wei)(big.NewInt(c))
}

// NewWeiValue returns new cost value instance of Wei
func NewWeiValue(c int64) Wei {
	cost := NewWei(c)
	return *cost
}

// NewWeiValueS returns new cost value from a string cost value in wei
func NewWeiValueS(c string) (Wei, error) {
	s, err := decimal.NewFromString(c)
	if err != nil {
		return Wei{}, nil
	}
	w := s.Mul(decimal.RequireFromString("10").Pow(decimal.RequireFromString("18")))
	return *(*Wei)(w.BigInt()), nil
}

// Cmp mimic *big.Int.Cmp
func (b *Wei) Cmp(y *Wei) int {
	return b.ToInt().Cmp(y.ToInt())
}

// String returns Wei as string in wei
func (b *Wei) String() string {
	return format(b.ToInt(), 18)
}

// getDenominator returns 10**precision
func getDenominator(precision int) *big.Int {
	x := big.NewInt(10)
	return new(big.Int).Exp(x, big.NewInt(int64(precision)), nil)
}

func format(i *big.Int, precision int) string {
	r := big.NewRat(1, 1).SetFrac(i, getDenominator(precision))
	return fmt.Sprintf("%v", r.FloatString(precision))
}

// SetInt64 mimic *big.Int.SetInt64
func (b *Wei) SetInt64(c int64) *Wei {
	return (*Wei)(b.ToInt().SetInt64(c))
}

// SetString mimics *big.Int.SetString
func (b *Wei) SetString(s string, base int) (*Wei, bool) {
	c, ok := b.ToInt().SetString(s, base)
	return (*Wei)(c), ok
}

// MarshalText implements the encoding.TextMarshaler interface
func (b *Wei) MarshalText() ([]byte, error) {
	return b.ToInt().MarshalText()
}

// MarshalJSON implements the json.Marshaler interface
func (b Wei) MarshalJSON() ([]byte, error) {
	value, err := b.MarshalText()
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf(`"%s"`, value)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (b *Wei) UnmarshalJSON(data []byte) error {
	if utils.IsQuoted(data) {
		return b.UnmarshalText(utils.RemoveQuotes(data))
	}
	return errors.New("cannot unmarshal json.Number into currency")
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (b *Wei) UnmarshalText(text []byte) error {
	if _, ok := b.SetString(string(text), 10); !ok {
		return fmt.Errorf("cannot unmarshal %q into *Wei", text)
	}
	return nil
}

// IsZero return true if value is 0, otherwise false
func (b *Wei) IsZero() bool {
	zero := big.NewInt(0)
	return b.ToInt().Cmp(zero) == 0
}

// Symbol returns 'DAI'
func (b *Wei) Symbol() string {
	return "DAI"
}

// ToInt returns cost value as *big.Int
func (b *Wei) ToInt() *big.Int {
	return (*big.Int)(b)
}

// Scan reads the database value and returns an instance
func (b *Wei) Scan(value interface{}) error {
	return (*Biggy)(b).Scan(value)
}

// Value returns the cost value for serialization to database
func (b Wei) Value() (driver.Value, error) {
	return (Biggy)(b).Value()
}
