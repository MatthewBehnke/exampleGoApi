// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"iseage/bank/pkg/database/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetPasswordHash sets the "password_hash" field.
func (uc *UserCreate) SetPasswordHash(s string) *UserCreate {
	uc.mutation.SetPasswordHash(s)
	return uc
}

// SetNillablePasswordHash sets the "password_hash" field if the given value is not nil.
func (uc *UserCreate) SetNillablePasswordHash(s *string) *UserCreate {
	if s != nil {
		uc.SetPasswordHash(*s)
	}
	return uc
}

// SetConfirmSelector sets the "confirm_selector" field.
func (uc *UserCreate) SetConfirmSelector(s string) *UserCreate {
	uc.mutation.SetConfirmSelector(s)
	return uc
}

// SetNillableConfirmSelector sets the "confirm_selector" field if the given value is not nil.
func (uc *UserCreate) SetNillableConfirmSelector(s *string) *UserCreate {
	if s != nil {
		uc.SetConfirmSelector(*s)
	}
	return uc
}

// SetConfirmVerifier sets the "confirm_verifier" field.
func (uc *UserCreate) SetConfirmVerifier(s string) *UserCreate {
	uc.mutation.SetConfirmVerifier(s)
	return uc
}

// SetNillableConfirmVerifier sets the "confirm_verifier" field if the given value is not nil.
func (uc *UserCreate) SetNillableConfirmVerifier(s *string) *UserCreate {
	if s != nil {
		uc.SetConfirmVerifier(*s)
	}
	return uc
}

// SetConfirmed sets the "confirmed" field.
func (uc *UserCreate) SetConfirmed(b bool) *UserCreate {
	uc.mutation.SetConfirmed(b)
	return uc
}

// SetNillableConfirmed sets the "confirmed" field if the given value is not nil.
func (uc *UserCreate) SetNillableConfirmed(b *bool) *UserCreate {
	if b != nil {
		uc.SetConfirmed(*b)
	}
	return uc
}

// SetAttemptCount sets the "attempt_count" field.
func (uc *UserCreate) SetAttemptCount(i int) *UserCreate {
	uc.mutation.SetAttemptCount(i)
	return uc
}

// SetNillableAttemptCount sets the "attempt_count" field if the given value is not nil.
func (uc *UserCreate) SetNillableAttemptCount(i *int) *UserCreate {
	if i != nil {
		uc.SetAttemptCount(*i)
	}
	return uc
}

// SetLastAttempt sets the "last_attempt" field.
func (uc *UserCreate) SetLastAttempt(t time.Time) *UserCreate {
	uc.mutation.SetLastAttempt(t)
	return uc
}

// SetNillableLastAttempt sets the "last_attempt" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastAttempt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLastAttempt(*t)
	}
	return uc
}

// SetLocked sets the "locked" field.
func (uc *UserCreate) SetLocked(t time.Time) *UserCreate {
	uc.mutation.SetLocked(t)
	return uc
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (uc *UserCreate) SetNillableLocked(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLocked(*t)
	}
	return uc
}

// SetRecoverSelector sets the "recover_selector" field.
func (uc *UserCreate) SetRecoverSelector(s string) *UserCreate {
	uc.mutation.SetRecoverSelector(s)
	return uc
}

// SetNillableRecoverSelector sets the "recover_selector" field if the given value is not nil.
func (uc *UserCreate) SetNillableRecoverSelector(s *string) *UserCreate {
	if s != nil {
		uc.SetRecoverSelector(*s)
	}
	return uc
}

// SetRecoverVerifier sets the "recover_verifier" field.
func (uc *UserCreate) SetRecoverVerifier(s string) *UserCreate {
	uc.mutation.SetRecoverVerifier(s)
	return uc
}

// SetNillableRecoverVerifier sets the "recover_verifier" field if the given value is not nil.
func (uc *UserCreate) SetNillableRecoverVerifier(s *string) *UserCreate {
	if s != nil {
		uc.SetRecoverVerifier(*s)
	}
	return uc
}

// SetRecoverTokenExpiry sets the "recover_token_expiry" field.
func (uc *UserCreate) SetRecoverTokenExpiry(t time.Time) *UserCreate {
	uc.mutation.SetRecoverTokenExpiry(t)
	return uc
}

// SetNillableRecoverTokenExpiry sets the "recover_token_expiry" field if the given value is not nil.
func (uc *UserCreate) SetNillableRecoverTokenExpiry(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetRecoverTokenExpiry(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt
		uc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if v, ok := uc.mutation.ID(); ok {
		if err := user.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "User.id": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := uc.mutation.PasswordHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPasswordHash,
		})
		_node.PasswordHash = value
	}
	if value, ok := uc.mutation.ConfirmSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldConfirmSelector,
		})
		_node.ConfirmSelector = value
	}
	if value, ok := uc.mutation.ConfirmVerifier(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldConfirmVerifier,
		})
		_node.ConfirmVerifier = value
	}
	if value, ok := uc.mutation.Confirmed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldConfirmed,
		})
		_node.Confirmed = value
	}
	if value, ok := uc.mutation.AttemptCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAttemptCount,
		})
		_node.AttemptCount = value
	}
	if value, ok := uc.mutation.LastAttempt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLastAttempt,
		})
		_node.LastAttempt = value
	}
	if value, ok := uc.mutation.Locked(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLocked,
		})
		_node.Locked = value
	}
	if value, ok := uc.mutation.RecoverSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRecoverSelector,
		})
		_node.RecoverSelector = value
	}
	if value, ok := uc.mutation.RecoverVerifier(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldRecoverVerifier,
		})
		_node.RecoverVerifier = value
	}
	if value, ok := uc.mutation.RecoverTokenExpiry(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldRecoverTokenExpiry,
		})
		_node.RecoverTokenExpiry = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
