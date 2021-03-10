package models

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/3dw1nM0535/galva/utils"
	uuid "github.com/satori/go.uuid"
)

// ID is a UUID with custom display format
type ID uuid.UUID

// UUID converts it back to UUID
func (id ID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

// New returns a new ID
func NewID() *ID {
	uuid := uuid.NewV4()
	return (*ID)(&uuid)
}

// String satisfies the Stringer interface and remove all '-' from the string representation of the uuid
func (id *ID) String() string {
	return strings.Replace((*uuid.UUID)(id).String(), "-", "", -1)
}

// UnmarshalText implements encoding.TextMarshaler
func (id *ID) UnmarshalText(input []byte) error {
	input = utils.RemoveQuotes(input)
	return (*uuid.UUID)(id).UnmarshalText(input)
}

// UnmarshalString is a wrapper for UnmarshalText which takes a string
func (id *ID) UnmarshalString(input string) error {
	return id.UnmarshalText([]byte(input))
}

// Value returns this instance serialized for the store
func (id *ID) Value() (driver.Value, error) {
	if id == nil {
		return nil, nil
	}
	return id.String(), nil
}

// Scan reads the database value and returns an instance
func (id *ID) Scan(value interface{}) error {
	switch v := value.(type) {
	case []uint8:
		return id.UnmarshalText(v)
	case string:
		return id.UnmarshalString(v)
	default:
		return fmt.Errorf("unable to convert %v of %T to ID", value, value)
	}
}
