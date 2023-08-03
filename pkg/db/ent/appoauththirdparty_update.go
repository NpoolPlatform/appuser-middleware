// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appoauththirdparty"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppOAuthThirdPartyUpdate is the builder for updating AppOAuthThirdParty entities.
type AppOAuthThirdPartyUpdate struct {
	config
	hooks     []Hook
	mutation  *AppOAuthThirdPartyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppOAuthThirdPartyUpdate builder.
func (aotpu *AppOAuthThirdPartyUpdate) Where(ps ...predicate.AppOAuthThirdParty) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.Where(ps...)
	return aotpu
}

// SetCreatedAt sets the "created_at" field.
func (aotpu *AppOAuthThirdPartyUpdate) SetCreatedAt(u uint32) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.ResetCreatedAt()
	aotpu.mutation.SetCreatedAt(u)
	return aotpu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aotpu *AppOAuthThirdPartyUpdate) SetNillableCreatedAt(u *uint32) *AppOAuthThirdPartyUpdate {
	if u != nil {
		aotpu.SetCreatedAt(*u)
	}
	return aotpu
}

// AddCreatedAt adds u to the "created_at" field.
func (aotpu *AppOAuthThirdPartyUpdate) AddCreatedAt(u int32) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.AddCreatedAt(u)
	return aotpu
}

// SetUpdatedAt sets the "updated_at" field.
func (aotpu *AppOAuthThirdPartyUpdate) SetUpdatedAt(u uint32) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.ResetUpdatedAt()
	aotpu.mutation.SetUpdatedAt(u)
	return aotpu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (aotpu *AppOAuthThirdPartyUpdate) AddUpdatedAt(u int32) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.AddUpdatedAt(u)
	return aotpu
}

// SetDeletedAt sets the "deleted_at" field.
func (aotpu *AppOAuthThirdPartyUpdate) SetDeletedAt(u uint32) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.ResetDeletedAt()
	aotpu.mutation.SetDeletedAt(u)
	return aotpu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aotpu *AppOAuthThirdPartyUpdate) SetNillableDeletedAt(u *uint32) *AppOAuthThirdPartyUpdate {
	if u != nil {
		aotpu.SetDeletedAt(*u)
	}
	return aotpu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (aotpu *AppOAuthThirdPartyUpdate) AddDeletedAt(u int32) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.AddDeletedAt(u)
	return aotpu
}

// SetAppID sets the "app_id" field.
func (aotpu *AppOAuthThirdPartyUpdate) SetAppID(u uuid.UUID) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.SetAppID(u)
	return aotpu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (aotpu *AppOAuthThirdPartyUpdate) SetNillableAppID(u *uuid.UUID) *AppOAuthThirdPartyUpdate {
	if u != nil {
		aotpu.SetAppID(*u)
	}
	return aotpu
}

// ClearAppID clears the value of the "app_id" field.
func (aotpu *AppOAuthThirdPartyUpdate) ClearAppID() *AppOAuthThirdPartyUpdate {
	aotpu.mutation.ClearAppID()
	return aotpu
}

// SetThirdPartyID sets the "third_party_id" field.
func (aotpu *AppOAuthThirdPartyUpdate) SetThirdPartyID(u uuid.UUID) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.SetThirdPartyID(u)
	return aotpu
}

// SetNillableThirdPartyID sets the "third_party_id" field if the given value is not nil.
func (aotpu *AppOAuthThirdPartyUpdate) SetNillableThirdPartyID(u *uuid.UUID) *AppOAuthThirdPartyUpdate {
	if u != nil {
		aotpu.SetThirdPartyID(*u)
	}
	return aotpu
}

// ClearThirdPartyID clears the value of the "third_party_id" field.
func (aotpu *AppOAuthThirdPartyUpdate) ClearThirdPartyID() *AppOAuthThirdPartyUpdate {
	aotpu.mutation.ClearThirdPartyID()
	return aotpu
}

// SetClientID sets the "client_id" field.
func (aotpu *AppOAuthThirdPartyUpdate) SetClientID(s string) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.SetClientID(s)
	return aotpu
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (aotpu *AppOAuthThirdPartyUpdate) SetNillableClientID(s *string) *AppOAuthThirdPartyUpdate {
	if s != nil {
		aotpu.SetClientID(*s)
	}
	return aotpu
}

// ClearClientID clears the value of the "client_id" field.
func (aotpu *AppOAuthThirdPartyUpdate) ClearClientID() *AppOAuthThirdPartyUpdate {
	aotpu.mutation.ClearClientID()
	return aotpu
}

// SetClientSecret sets the "client_secret" field.
func (aotpu *AppOAuthThirdPartyUpdate) SetClientSecret(s string) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.SetClientSecret(s)
	return aotpu
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (aotpu *AppOAuthThirdPartyUpdate) SetNillableClientSecret(s *string) *AppOAuthThirdPartyUpdate {
	if s != nil {
		aotpu.SetClientSecret(*s)
	}
	return aotpu
}

// ClearClientSecret clears the value of the "client_secret" field.
func (aotpu *AppOAuthThirdPartyUpdate) ClearClientSecret() *AppOAuthThirdPartyUpdate {
	aotpu.mutation.ClearClientSecret()
	return aotpu
}

// SetCallbackURL sets the "callback_url" field.
func (aotpu *AppOAuthThirdPartyUpdate) SetCallbackURL(s string) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.SetCallbackURL(s)
	return aotpu
}

// SetNillableCallbackURL sets the "callback_url" field if the given value is not nil.
func (aotpu *AppOAuthThirdPartyUpdate) SetNillableCallbackURL(s *string) *AppOAuthThirdPartyUpdate {
	if s != nil {
		aotpu.SetCallbackURL(*s)
	}
	return aotpu
}

// ClearCallbackURL clears the value of the "callback_url" field.
func (aotpu *AppOAuthThirdPartyUpdate) ClearCallbackURL() *AppOAuthThirdPartyUpdate {
	aotpu.mutation.ClearCallbackURL()
	return aotpu
}

// SetSalt sets the "salt" field.
func (aotpu *AppOAuthThirdPartyUpdate) SetSalt(s string) *AppOAuthThirdPartyUpdate {
	aotpu.mutation.SetSalt(s)
	return aotpu
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (aotpu *AppOAuthThirdPartyUpdate) SetNillableSalt(s *string) *AppOAuthThirdPartyUpdate {
	if s != nil {
		aotpu.SetSalt(*s)
	}
	return aotpu
}

// ClearSalt clears the value of the "salt" field.
func (aotpu *AppOAuthThirdPartyUpdate) ClearSalt() *AppOAuthThirdPartyUpdate {
	aotpu.mutation.ClearSalt()
	return aotpu
}

// Mutation returns the AppOAuthThirdPartyMutation object of the builder.
func (aotpu *AppOAuthThirdPartyUpdate) Mutation() *AppOAuthThirdPartyMutation {
	return aotpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aotpu *AppOAuthThirdPartyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := aotpu.defaults(); err != nil {
		return 0, err
	}
	if len(aotpu.hooks) == 0 {
		affected, err = aotpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppOAuthThirdPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aotpu.mutation = mutation
			affected, err = aotpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aotpu.hooks) - 1; i >= 0; i-- {
			if aotpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aotpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aotpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aotpu *AppOAuthThirdPartyUpdate) SaveX(ctx context.Context) int {
	affected, err := aotpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aotpu *AppOAuthThirdPartyUpdate) Exec(ctx context.Context) error {
	_, err := aotpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aotpu *AppOAuthThirdPartyUpdate) ExecX(ctx context.Context) {
	if err := aotpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aotpu *AppOAuthThirdPartyUpdate) defaults() error {
	if _, ok := aotpu.mutation.UpdatedAt(); !ok {
		if appoauththirdparty.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appoauththirdparty.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appoauththirdparty.UpdateDefaultUpdatedAt()
		aotpu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aotpu *AppOAuthThirdPartyUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppOAuthThirdPartyUpdate {
	aotpu.modifiers = append(aotpu.modifiers, modifiers...)
	return aotpu
}

func (aotpu *AppOAuthThirdPartyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appoauththirdparty.Table,
			Columns: appoauththirdparty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appoauththirdparty.FieldID,
			},
		},
	}
	if ps := aotpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aotpu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldCreatedAt,
		})
	}
	if value, ok := aotpu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldCreatedAt,
		})
	}
	if value, ok := aotpu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldUpdatedAt,
		})
	}
	if value, ok := aotpu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldUpdatedAt,
		})
	}
	if value, ok := aotpu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldDeletedAt,
		})
	}
	if value, ok := aotpu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldDeletedAt,
		})
	}
	if value, ok := aotpu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appoauththirdparty.FieldAppID,
		})
	}
	if aotpu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appoauththirdparty.FieldAppID,
		})
	}
	if value, ok := aotpu.mutation.ThirdPartyID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appoauththirdparty.FieldThirdPartyID,
		})
	}
	if aotpu.mutation.ThirdPartyIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appoauththirdparty.FieldThirdPartyID,
		})
	}
	if value, ok := aotpu.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoauththirdparty.FieldClientID,
		})
	}
	if aotpu.mutation.ClientIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appoauththirdparty.FieldClientID,
		})
	}
	if value, ok := aotpu.mutation.ClientSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoauththirdparty.FieldClientSecret,
		})
	}
	if aotpu.mutation.ClientSecretCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appoauththirdparty.FieldClientSecret,
		})
	}
	if value, ok := aotpu.mutation.CallbackURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoauththirdparty.FieldCallbackURL,
		})
	}
	if aotpu.mutation.CallbackURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appoauththirdparty.FieldCallbackURL,
		})
	}
	if value, ok := aotpu.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoauththirdparty.FieldSalt,
		})
	}
	if aotpu.mutation.SaltCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appoauththirdparty.FieldSalt,
		})
	}
	_spec.Modifiers = aotpu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, aotpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appoauththirdparty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppOAuthThirdPartyUpdateOne is the builder for updating a single AppOAuthThirdParty entity.
type AppOAuthThirdPartyUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppOAuthThirdPartyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetCreatedAt(u uint32) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.ResetCreatedAt()
	aotpuo.mutation.SetCreatedAt(u)
	return aotpuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetNillableCreatedAt(u *uint32) *AppOAuthThirdPartyUpdateOne {
	if u != nil {
		aotpuo.SetCreatedAt(*u)
	}
	return aotpuo
}

// AddCreatedAt adds u to the "created_at" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) AddCreatedAt(u int32) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.AddCreatedAt(u)
	return aotpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetUpdatedAt(u uint32) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.ResetUpdatedAt()
	aotpuo.mutation.SetUpdatedAt(u)
	return aotpuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) AddUpdatedAt(u int32) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.AddUpdatedAt(u)
	return aotpuo
}

// SetDeletedAt sets the "deleted_at" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetDeletedAt(u uint32) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.ResetDeletedAt()
	aotpuo.mutation.SetDeletedAt(u)
	return aotpuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetNillableDeletedAt(u *uint32) *AppOAuthThirdPartyUpdateOne {
	if u != nil {
		aotpuo.SetDeletedAt(*u)
	}
	return aotpuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) AddDeletedAt(u int32) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.AddDeletedAt(u)
	return aotpuo
}

// SetAppID sets the "app_id" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetAppID(u uuid.UUID) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.SetAppID(u)
	return aotpuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetNillableAppID(u *uuid.UUID) *AppOAuthThirdPartyUpdateOne {
	if u != nil {
		aotpuo.SetAppID(*u)
	}
	return aotpuo
}

// ClearAppID clears the value of the "app_id" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) ClearAppID() *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.ClearAppID()
	return aotpuo
}

// SetThirdPartyID sets the "third_party_id" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetThirdPartyID(u uuid.UUID) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.SetThirdPartyID(u)
	return aotpuo
}

// SetNillableThirdPartyID sets the "third_party_id" field if the given value is not nil.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetNillableThirdPartyID(u *uuid.UUID) *AppOAuthThirdPartyUpdateOne {
	if u != nil {
		aotpuo.SetThirdPartyID(*u)
	}
	return aotpuo
}

// ClearThirdPartyID clears the value of the "third_party_id" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) ClearThirdPartyID() *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.ClearThirdPartyID()
	return aotpuo
}

// SetClientID sets the "client_id" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetClientID(s string) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.SetClientID(s)
	return aotpuo
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetNillableClientID(s *string) *AppOAuthThirdPartyUpdateOne {
	if s != nil {
		aotpuo.SetClientID(*s)
	}
	return aotpuo
}

// ClearClientID clears the value of the "client_id" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) ClearClientID() *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.ClearClientID()
	return aotpuo
}

// SetClientSecret sets the "client_secret" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetClientSecret(s string) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.SetClientSecret(s)
	return aotpuo
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetNillableClientSecret(s *string) *AppOAuthThirdPartyUpdateOne {
	if s != nil {
		aotpuo.SetClientSecret(*s)
	}
	return aotpuo
}

// ClearClientSecret clears the value of the "client_secret" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) ClearClientSecret() *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.ClearClientSecret()
	return aotpuo
}

// SetCallbackURL sets the "callback_url" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetCallbackURL(s string) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.SetCallbackURL(s)
	return aotpuo
}

// SetNillableCallbackURL sets the "callback_url" field if the given value is not nil.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetNillableCallbackURL(s *string) *AppOAuthThirdPartyUpdateOne {
	if s != nil {
		aotpuo.SetCallbackURL(*s)
	}
	return aotpuo
}

// ClearCallbackURL clears the value of the "callback_url" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) ClearCallbackURL() *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.ClearCallbackURL()
	return aotpuo
}

// SetSalt sets the "salt" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetSalt(s string) *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.SetSalt(s)
	return aotpuo
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SetNillableSalt(s *string) *AppOAuthThirdPartyUpdateOne {
	if s != nil {
		aotpuo.SetSalt(*s)
	}
	return aotpuo
}

// ClearSalt clears the value of the "salt" field.
func (aotpuo *AppOAuthThirdPartyUpdateOne) ClearSalt() *AppOAuthThirdPartyUpdateOne {
	aotpuo.mutation.ClearSalt()
	return aotpuo
}

// Mutation returns the AppOAuthThirdPartyMutation object of the builder.
func (aotpuo *AppOAuthThirdPartyUpdateOne) Mutation() *AppOAuthThirdPartyMutation {
	return aotpuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aotpuo *AppOAuthThirdPartyUpdateOne) Select(field string, fields ...string) *AppOAuthThirdPartyUpdateOne {
	aotpuo.fields = append([]string{field}, fields...)
	return aotpuo
}

// Save executes the query and returns the updated AppOAuthThirdParty entity.
func (aotpuo *AppOAuthThirdPartyUpdateOne) Save(ctx context.Context) (*AppOAuthThirdParty, error) {
	var (
		err  error
		node *AppOAuthThirdParty
	)
	if err := aotpuo.defaults(); err != nil {
		return nil, err
	}
	if len(aotpuo.hooks) == 0 {
		node, err = aotpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppOAuthThirdPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aotpuo.mutation = mutation
			node, err = aotpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aotpuo.hooks) - 1; i >= 0; i-- {
			if aotpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aotpuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, aotpuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppOAuthThirdParty)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppOAuthThirdPartyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aotpuo *AppOAuthThirdPartyUpdateOne) SaveX(ctx context.Context) *AppOAuthThirdParty {
	node, err := aotpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aotpuo *AppOAuthThirdPartyUpdateOne) Exec(ctx context.Context) error {
	_, err := aotpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aotpuo *AppOAuthThirdPartyUpdateOne) ExecX(ctx context.Context) {
	if err := aotpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aotpuo *AppOAuthThirdPartyUpdateOne) defaults() error {
	if _, ok := aotpuo.mutation.UpdatedAt(); !ok {
		if appoauththirdparty.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appoauththirdparty.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appoauththirdparty.UpdateDefaultUpdatedAt()
		aotpuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aotpuo *AppOAuthThirdPartyUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppOAuthThirdPartyUpdateOne {
	aotpuo.modifiers = append(aotpuo.modifiers, modifiers...)
	return aotpuo
}

func (aotpuo *AppOAuthThirdPartyUpdateOne) sqlSave(ctx context.Context) (_node *AppOAuthThirdParty, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appoauththirdparty.Table,
			Columns: appoauththirdparty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appoauththirdparty.FieldID,
			},
		},
	}
	id, ok := aotpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppOAuthThirdParty.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aotpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appoauththirdparty.FieldID)
		for _, f := range fields {
			if !appoauththirdparty.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appoauththirdparty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aotpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aotpuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldCreatedAt,
		})
	}
	if value, ok := aotpuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldCreatedAt,
		})
	}
	if value, ok := aotpuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldUpdatedAt,
		})
	}
	if value, ok := aotpuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldUpdatedAt,
		})
	}
	if value, ok := aotpuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldDeletedAt,
		})
	}
	if value, ok := aotpuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appoauththirdparty.FieldDeletedAt,
		})
	}
	if value, ok := aotpuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appoauththirdparty.FieldAppID,
		})
	}
	if aotpuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appoauththirdparty.FieldAppID,
		})
	}
	if value, ok := aotpuo.mutation.ThirdPartyID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appoauththirdparty.FieldThirdPartyID,
		})
	}
	if aotpuo.mutation.ThirdPartyIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appoauththirdparty.FieldThirdPartyID,
		})
	}
	if value, ok := aotpuo.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoauththirdparty.FieldClientID,
		})
	}
	if aotpuo.mutation.ClientIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appoauththirdparty.FieldClientID,
		})
	}
	if value, ok := aotpuo.mutation.ClientSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoauththirdparty.FieldClientSecret,
		})
	}
	if aotpuo.mutation.ClientSecretCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appoauththirdparty.FieldClientSecret,
		})
	}
	if value, ok := aotpuo.mutation.CallbackURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoauththirdparty.FieldCallbackURL,
		})
	}
	if aotpuo.mutation.CallbackURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appoauththirdparty.FieldCallbackURL,
		})
	}
	if value, ok := aotpuo.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoauththirdparty.FieldSalt,
		})
	}
	if aotpuo.mutation.SaltCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appoauththirdparty.FieldSalt,
		})
	}
	_spec.Modifiers = aotpuo.modifiers
	_node = &AppOAuthThirdParty{config: aotpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aotpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appoauththirdparty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
