// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldConfirmSelector holds the string denoting the confirm_selector field in the database.
	FieldConfirmSelector = "confirm_selector"
	// FieldConfirmVerifier holds the string denoting the confirm_verifier field in the database.
	FieldConfirmVerifier = "confirm_verifier"
	// FieldConfirmed holds the string denoting the confirmed field in the database.
	FieldConfirmed = "confirmed"
	// FieldAttemptCount holds the string denoting the attempt_count field in the database.
	FieldAttemptCount = "attempt_count"
	// FieldLastAttempt holds the string denoting the last_attempt field in the database.
	FieldLastAttempt = "last_attempt"
	// FieldLocked holds the string denoting the locked field in the database.
	FieldLocked = "locked"
	// FieldRecoverSelector holds the string denoting the recover_selector field in the database.
	FieldRecoverSelector = "recover_selector"
	// FieldRecoverVerifier holds the string denoting the recover_verifier field in the database.
	FieldRecoverVerifier = "recover_verifier"
	// FieldRecoverTokenExpiry holds the string denoting the recover_token_expiry field in the database.
	FieldRecoverTokenExpiry = "recover_token_expiry"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEmail,
	FieldPasswordHash,
	FieldConfirmSelector,
	FieldConfirmVerifier,
	FieldConfirmed,
	FieldAttemptCount,
	FieldLastAttempt,
	FieldLocked,
	FieldRecoverSelector,
	FieldRecoverVerifier,
	FieldRecoverTokenExpiry,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)
