// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/auth"
	"github.com/google/uuid"
)

// AuthCreate is the builder for creating a Auth entity.
type AuthCreate struct {
	config
	mutation *AuthMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ac *AuthCreate) SetCreatedAt(u uint32) *AuthCreate {
	ac.mutation.SetCreatedAt(u)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AuthCreate) SetNillableCreatedAt(u *uint32) *AuthCreate {
	if u != nil {
		ac.SetCreatedAt(*u)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AuthCreate) SetUpdatedAt(u uint32) *AuthCreate {
	ac.mutation.SetUpdatedAt(u)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AuthCreate) SetNillableUpdatedAt(u *uint32) *AuthCreate {
	if u != nil {
		ac.SetUpdatedAt(*u)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AuthCreate) SetDeletedAt(u uint32) *AuthCreate {
	ac.mutation.SetDeletedAt(u)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AuthCreate) SetNillableDeletedAt(u *uint32) *AuthCreate {
	if u != nil {
		ac.SetDeletedAt(*u)
	}
	return ac
}

// SetEntID sets the "ent_id" field.
func (ac *AuthCreate) SetEntID(u uuid.UUID) *AuthCreate {
	ac.mutation.SetEntID(u)
	return ac
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ac *AuthCreate) SetNillableEntID(u *uuid.UUID) *AuthCreate {
	if u != nil {
		ac.SetEntID(*u)
	}
	return ac
}

// SetAppID sets the "app_id" field.
func (ac *AuthCreate) SetAppID(u uuid.UUID) *AuthCreate {
	ac.mutation.SetAppID(u)
	return ac
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ac *AuthCreate) SetNillableAppID(u *uuid.UUID) *AuthCreate {
	if u != nil {
		ac.SetAppID(*u)
	}
	return ac
}

// SetRoleID sets the "role_id" field.
func (ac *AuthCreate) SetRoleID(u uuid.UUID) *AuthCreate {
	ac.mutation.SetRoleID(u)
	return ac
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (ac *AuthCreate) SetNillableRoleID(u *uuid.UUID) *AuthCreate {
	if u != nil {
		ac.SetRoleID(*u)
	}
	return ac
}

// SetUserID sets the "user_id" field.
func (ac *AuthCreate) SetUserID(u uuid.UUID) *AuthCreate {
	ac.mutation.SetUserID(u)
	return ac
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ac *AuthCreate) SetNillableUserID(u *uuid.UUID) *AuthCreate {
	if u != nil {
		ac.SetUserID(*u)
	}
	return ac
}

// SetResource sets the "resource" field.
func (ac *AuthCreate) SetResource(s string) *AuthCreate {
	ac.mutation.SetResource(s)
	return ac
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (ac *AuthCreate) SetNillableResource(s *string) *AuthCreate {
	if s != nil {
		ac.SetResource(*s)
	}
	return ac
}

// SetMethod sets the "method" field.
func (ac *AuthCreate) SetMethod(s string) *AuthCreate {
	ac.mutation.SetMethod(s)
	return ac
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (ac *AuthCreate) SetNillableMethod(s *string) *AuthCreate {
	if s != nil {
		ac.SetMethod(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AuthCreate) SetID(u uint32) *AuthCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the AuthMutation object of the builder.
func (ac *AuthCreate) Mutation() *AuthMutation {
	return ac.mutation
}

// Save creates the Auth in the database.
func (ac *AuthCreate) Save(ctx context.Context) (*Auth, error) {
	var (
		err  error
		node *Auth
	)
	if err := ac.defaults(); err != nil {
		return nil, err
	}
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Auth)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AuthMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuthCreate) SaveX(ctx context.Context) *Auth {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AuthCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AuthCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AuthCreate) defaults() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		if auth.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized auth.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := auth.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		if auth.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized auth.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := auth.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		if auth.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized auth.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := auth.DefaultDeletedAt()
		ac.mutation.SetDeletedAt(v)
	}
	if _, ok := ac.mutation.EntID(); !ok {
		if auth.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized auth.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := auth.DefaultEntID()
		ac.mutation.SetEntID(v)
	}
	if _, ok := ac.mutation.AppID(); !ok {
		if auth.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized auth.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := auth.DefaultAppID()
		ac.mutation.SetAppID(v)
	}
	if _, ok := ac.mutation.RoleID(); !ok {
		if auth.DefaultRoleID == nil {
			return fmt.Errorf("ent: uninitialized auth.DefaultRoleID (forgotten import ent/runtime?)")
		}
		v := auth.DefaultRoleID()
		ac.mutation.SetRoleID(v)
	}
	if _, ok := ac.mutation.UserID(); !ok {
		if auth.DefaultUserID == nil {
			return fmt.Errorf("ent: uninitialized auth.DefaultUserID (forgotten import ent/runtime?)")
		}
		v := auth.DefaultUserID()
		ac.mutation.SetUserID(v)
	}
	if _, ok := ac.mutation.Resource(); !ok {
		v := auth.DefaultResource
		ac.mutation.SetResource(v)
	}
	if _, ok := ac.mutation.Method(); !ok {
		v := auth.DefaultMethod
		ac.mutation.SetMethod(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ac *AuthCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Auth.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Auth.updated_at"`)}
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Auth.deleted_at"`)}
	}
	if _, ok := ac.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "Auth.ent_id"`)}
	}
	return nil
}

func (ac *AuthCreate) sqlSave(ctx context.Context) (*Auth, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (ac *AuthCreate) createSpec() (*Auth, *sqlgraph.CreateSpec) {
	var (
		_node = &Auth{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: auth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: auth.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := ac.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: auth.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := ac.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: auth.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := ac.mutation.RoleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: auth.FieldRoleID,
		})
		_node.RoleID = value
	}
	if value, ok := ac.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: auth.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := ac.mutation.Resource(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: auth.FieldResource,
		})
		_node.Resource = value
	}
	if value, ok := ac.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: auth.FieldMethod,
		})
		_node.Method = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Auth.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ac *AuthCreate) OnConflict(opts ...sql.ConflictOption) *AuthUpsertOne {
	ac.conflict = opts
	return &AuthUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ac *AuthCreate) OnConflictColumns(columns ...string) *AuthUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AuthUpsertOne{
		create: ac,
	}
}

type (
	// AuthUpsertOne is the builder for "upsert"-ing
	//  one Auth node.
	AuthUpsertOne struct {
		create *AuthCreate
	}

	// AuthUpsert is the "OnConflict" setter.
	AuthUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AuthUpsert) SetCreatedAt(v uint32) *AuthUpsert {
	u.Set(auth.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AuthUpsert) UpdateCreatedAt() *AuthUpsert {
	u.SetExcluded(auth.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AuthUpsert) AddCreatedAt(v uint32) *AuthUpsert {
	u.Add(auth.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthUpsert) SetUpdatedAt(v uint32) *AuthUpsert {
	u.Set(auth.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthUpsert) UpdateUpdatedAt() *AuthUpsert {
	u.SetExcluded(auth.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AuthUpsert) AddUpdatedAt(v uint32) *AuthUpsert {
	u.Add(auth.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AuthUpsert) SetDeletedAt(v uint32) *AuthUpsert {
	u.Set(auth.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AuthUpsert) UpdateDeletedAt() *AuthUpsert {
	u.SetExcluded(auth.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AuthUpsert) AddDeletedAt(v uint32) *AuthUpsert {
	u.Add(auth.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *AuthUpsert) SetEntID(v uuid.UUID) *AuthUpsert {
	u.Set(auth.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AuthUpsert) UpdateEntID() *AuthUpsert {
	u.SetExcluded(auth.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AuthUpsert) SetAppID(v uuid.UUID) *AuthUpsert {
	u.Set(auth.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AuthUpsert) UpdateAppID() *AuthUpsert {
	u.SetExcluded(auth.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *AuthUpsert) ClearAppID() *AuthUpsert {
	u.SetNull(auth.FieldAppID)
	return u
}

// SetRoleID sets the "role_id" field.
func (u *AuthUpsert) SetRoleID(v uuid.UUID) *AuthUpsert {
	u.Set(auth.FieldRoleID, v)
	return u
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *AuthUpsert) UpdateRoleID() *AuthUpsert {
	u.SetExcluded(auth.FieldRoleID)
	return u
}

// ClearRoleID clears the value of the "role_id" field.
func (u *AuthUpsert) ClearRoleID() *AuthUpsert {
	u.SetNull(auth.FieldRoleID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *AuthUpsert) SetUserID(v uuid.UUID) *AuthUpsert {
	u.Set(auth.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthUpsert) UpdateUserID() *AuthUpsert {
	u.SetExcluded(auth.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *AuthUpsert) ClearUserID() *AuthUpsert {
	u.SetNull(auth.FieldUserID)
	return u
}

// SetResource sets the "resource" field.
func (u *AuthUpsert) SetResource(v string) *AuthUpsert {
	u.Set(auth.FieldResource, v)
	return u
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AuthUpsert) UpdateResource() *AuthUpsert {
	u.SetExcluded(auth.FieldResource)
	return u
}

// ClearResource clears the value of the "resource" field.
func (u *AuthUpsert) ClearResource() *AuthUpsert {
	u.SetNull(auth.FieldResource)
	return u
}

// SetMethod sets the "method" field.
func (u *AuthUpsert) SetMethod(v string) *AuthUpsert {
	u.Set(auth.FieldMethod, v)
	return u
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AuthUpsert) UpdateMethod() *AuthUpsert {
	u.SetExcluded(auth.FieldMethod)
	return u
}

// ClearMethod clears the value of the "method" field.
func (u *AuthUpsert) ClearMethod() *AuthUpsert {
	u.SetNull(auth.FieldMethod)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(auth.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AuthUpsertOne) UpdateNewValues() *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(auth.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Auth.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AuthUpsertOne) Ignore() *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthUpsertOne) DoNothing() *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthCreate.OnConflict
// documentation for more info.
func (u *AuthUpsertOne) Update(set func(*AuthUpsert)) *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AuthUpsertOne) SetCreatedAt(v uint32) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AuthUpsertOne) AddCreatedAt(v uint32) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateCreatedAt() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthUpsertOne) SetUpdatedAt(v uint32) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AuthUpsertOne) AddUpdatedAt(v uint32) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateUpdatedAt() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AuthUpsertOne) SetDeletedAt(v uint32) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AuthUpsertOne) AddDeletedAt(v uint32) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateDeletedAt() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AuthUpsertOne) SetEntID(v uuid.UUID) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateEntID() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AuthUpsertOne) SetAppID(v uuid.UUID) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateAppID() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AuthUpsertOne) ClearAppID() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.ClearAppID()
	})
}

// SetRoleID sets the "role_id" field.
func (u *AuthUpsertOne) SetRoleID(v uuid.UUID) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetRoleID(v)
	})
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateRoleID() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateRoleID()
	})
}

// ClearRoleID clears the value of the "role_id" field.
func (u *AuthUpsertOne) ClearRoleID() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.ClearRoleID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AuthUpsertOne) SetUserID(v uuid.UUID) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateUserID() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *AuthUpsertOne) ClearUserID() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.ClearUserID()
	})
}

// SetResource sets the "resource" field.
func (u *AuthUpsertOne) SetResource(v string) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetResource(v)
	})
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateResource() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateResource()
	})
}

// ClearResource clears the value of the "resource" field.
func (u *AuthUpsertOne) ClearResource() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.ClearResource()
	})
}

// SetMethod sets the "method" field.
func (u *AuthUpsertOne) SetMethod(v string) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateMethod() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateMethod()
	})
}

// ClearMethod clears the value of the "method" field.
func (u *AuthUpsertOne) ClearMethod() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.ClearMethod()
	})
}

// Exec executes the query.
func (u *AuthUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AuthUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AuthUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AuthCreateBulk is the builder for creating many Auth entities in bulk.
type AuthCreateBulk struct {
	config
	builders []*AuthCreate
	conflict []sql.ConflictOption
}

// Save creates the Auth entities in the database.
func (acb *AuthCreateBulk) Save(ctx context.Context) ([]*Auth, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Auth, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
					nodes[i].ID = uint32(id)
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AuthCreateBulk) SaveX(ctx context.Context) []*Auth {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AuthCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AuthCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Auth.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (acb *AuthCreateBulk) OnConflict(opts ...sql.ConflictOption) *AuthUpsertBulk {
	acb.conflict = opts
	return &AuthUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acb *AuthCreateBulk) OnConflictColumns(columns ...string) *AuthUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AuthUpsertBulk{
		create: acb,
	}
}

// AuthUpsertBulk is the builder for "upsert"-ing
// a bulk of Auth nodes.
type AuthUpsertBulk struct {
	create *AuthCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(auth.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AuthUpsertBulk) UpdateNewValues() *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(auth.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AuthUpsertBulk) Ignore() *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthUpsertBulk) DoNothing() *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthCreateBulk.OnConflict
// documentation for more info.
func (u *AuthUpsertBulk) Update(set func(*AuthUpsert)) *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AuthUpsertBulk) SetCreatedAt(v uint32) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AuthUpsertBulk) AddCreatedAt(v uint32) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateCreatedAt() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthUpsertBulk) SetUpdatedAt(v uint32) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AuthUpsertBulk) AddUpdatedAt(v uint32) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateUpdatedAt() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AuthUpsertBulk) SetDeletedAt(v uint32) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AuthUpsertBulk) AddDeletedAt(v uint32) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateDeletedAt() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AuthUpsertBulk) SetEntID(v uuid.UUID) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateEntID() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AuthUpsertBulk) SetAppID(v uuid.UUID) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateAppID() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AuthUpsertBulk) ClearAppID() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.ClearAppID()
	})
}

// SetRoleID sets the "role_id" field.
func (u *AuthUpsertBulk) SetRoleID(v uuid.UUID) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetRoleID(v)
	})
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateRoleID() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateRoleID()
	})
}

// ClearRoleID clears the value of the "role_id" field.
func (u *AuthUpsertBulk) ClearRoleID() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.ClearRoleID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AuthUpsertBulk) SetUserID(v uuid.UUID) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateUserID() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *AuthUpsertBulk) ClearUserID() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.ClearUserID()
	})
}

// SetResource sets the "resource" field.
func (u *AuthUpsertBulk) SetResource(v string) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetResource(v)
	})
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateResource() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateResource()
	})
}

// ClearResource clears the value of the "resource" field.
func (u *AuthUpsertBulk) ClearResource() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.ClearResource()
	})
}

// SetMethod sets the "method" field.
func (u *AuthUpsertBulk) SetMethod(v string) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateMethod() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateMethod()
	})
}

// ClearMethod clears the value of the "method" field.
func (u *AuthUpsertBulk) ClearMethod() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.ClearMethod()
	})
}

// Exec executes the query.
func (u *AuthUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AuthCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
