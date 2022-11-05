// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/MatthewBehnke/exampleGoApi/pkg/database/ent/user"
)

// User is the model domain for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// PasswordHash holds the value of the "password_hash" field.
	PasswordHash string `json:"password_hash,omitempty"`
	// AttemptCount holds the value of the "attempt_count" field.
	AttemptCount int `json:"attempt_count,omitempty"`
	// LastAttempt holds the value of the "last_attempt" field.
	LastAttempt time.Time `json:"last_attempt,omitempty"`
	// Locked holds the value of the "locked" field.
	Locked time.Time `json:"locked,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldAttemptCount:
			values[i] = new(sql.NullInt64)
		case user.FieldEmail, user.FieldPasswordHash, user.FieldRole:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldLastAttempt, user.FieldLocked:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPasswordHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_hash", values[i])
			} else if value.Valid {
				u.PasswordHash = value.String
			}
		case user.FieldAttemptCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attempt_count", values[i])
			} else if value.Valid {
				u.AttemptCount = int(value.Int64)
			}
		case user.FieldLastAttempt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_attempt", values[i])
			} else if value.Valid {
				u.LastAttempt = value.Time
			}
		case user.FieldLocked:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				u.Locked = value.Time
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User domain that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional domain")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password_hash=")
	builder.WriteString(u.PasswordHash)
	builder.WriteString(", ")
	builder.WriteString("attempt_count=")
	builder.WriteString(fmt.Sprintf("%v", u.AttemptCount))
	builder.WriteString(", ")
	builder.WriteString("last_attempt=")
	builder.WriteString(u.LastAttempt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(u.Locked.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(u.Role)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
