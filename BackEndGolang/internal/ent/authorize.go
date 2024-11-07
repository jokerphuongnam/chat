// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-backend/internal/ent/authorize"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Authorize is the model entity for the Authorize schema.
type Authorize struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	Token        string `json:"token,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Authorize) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authorize.FieldToken:
			values[i] = new(sql.NullString)
		case authorize.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Authorize fields.
func (a *Authorize) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authorize.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case authorize.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				a.Token = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Authorize.
// This includes values selected through modifiers, order, etc.
func (a *Authorize) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Authorize.
// Note that you need to call Authorize.Unwrap() before calling this method if this Authorize
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Authorize) Update() *AuthorizeUpdateOne {
	return NewAuthorizeClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Authorize entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Authorize) Unwrap() *Authorize {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Authorize is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Authorize) String() string {
	var builder strings.Builder
	builder.WriteString("Authorize(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("token=")
	builder.WriteString(a.Token)
	builder.WriteByte(')')
	return builder.String()
}

// Authorizes is a parsable slice of Authorize.
type Authorizes []*Authorize
