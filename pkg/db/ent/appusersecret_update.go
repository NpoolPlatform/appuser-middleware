// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appusersecret"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserSecretUpdate is the builder for updating AppUserSecret entities.
type AppUserSecretUpdate struct {
	config
	hooks     []Hook
	mutation  *AppUserSecretMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppUserSecretUpdate builder.
func (ausu *AppUserSecretUpdate) Where(ps ...predicate.AppUserSecret) *AppUserSecretUpdate {
	ausu.mutation.Where(ps...)
	return ausu
}

// SetCreatedAt sets the "created_at" field.
func (ausu *AppUserSecretUpdate) SetCreatedAt(u uint32) *AppUserSecretUpdate {
	ausu.mutation.ResetCreatedAt()
	ausu.mutation.SetCreatedAt(u)
	return ausu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ausu *AppUserSecretUpdate) SetNillableCreatedAt(u *uint32) *AppUserSecretUpdate {
	if u != nil {
		ausu.SetCreatedAt(*u)
	}
	return ausu
}

// AddCreatedAt adds u to the "created_at" field.
func (ausu *AppUserSecretUpdate) AddCreatedAt(u int32) *AppUserSecretUpdate {
	ausu.mutation.AddCreatedAt(u)
	return ausu
}

// SetUpdatedAt sets the "updated_at" field.
func (ausu *AppUserSecretUpdate) SetUpdatedAt(u uint32) *AppUserSecretUpdate {
	ausu.mutation.ResetUpdatedAt()
	ausu.mutation.SetUpdatedAt(u)
	return ausu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ausu *AppUserSecretUpdate) AddUpdatedAt(u int32) *AppUserSecretUpdate {
	ausu.mutation.AddUpdatedAt(u)
	return ausu
}

// SetDeletedAt sets the "deleted_at" field.
func (ausu *AppUserSecretUpdate) SetDeletedAt(u uint32) *AppUserSecretUpdate {
	ausu.mutation.ResetDeletedAt()
	ausu.mutation.SetDeletedAt(u)
	return ausu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ausu *AppUserSecretUpdate) SetNillableDeletedAt(u *uint32) *AppUserSecretUpdate {
	if u != nil {
		ausu.SetDeletedAt(*u)
	}
	return ausu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ausu *AppUserSecretUpdate) AddDeletedAt(u int32) *AppUserSecretUpdate {
	ausu.mutation.AddDeletedAt(u)
	return ausu
}

// SetEntID sets the "ent_id" field.
func (ausu *AppUserSecretUpdate) SetEntID(u uuid.UUID) *AppUserSecretUpdate {
	ausu.mutation.SetEntID(u)
	return ausu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ausu *AppUserSecretUpdate) SetNillableEntID(u *uuid.UUID) *AppUserSecretUpdate {
	if u != nil {
		ausu.SetEntID(*u)
	}
	return ausu
}

// SetAppID sets the "app_id" field.
func (ausu *AppUserSecretUpdate) SetAppID(u uuid.UUID) *AppUserSecretUpdate {
	ausu.mutation.SetAppID(u)
	return ausu
}

// SetUserID sets the "user_id" field.
func (ausu *AppUserSecretUpdate) SetUserID(u uuid.UUID) *AppUserSecretUpdate {
	ausu.mutation.SetUserID(u)
	return ausu
}

// SetPasswordHash sets the "password_hash" field.
func (ausu *AppUserSecretUpdate) SetPasswordHash(s string) *AppUserSecretUpdate {
	ausu.mutation.SetPasswordHash(s)
	return ausu
}

// SetSalt sets the "salt" field.
func (ausu *AppUserSecretUpdate) SetSalt(s string) *AppUserSecretUpdate {
	ausu.mutation.SetSalt(s)
	return ausu
}

// SetGoogleSecret sets the "google_secret" field.
func (ausu *AppUserSecretUpdate) SetGoogleSecret(s string) *AppUserSecretUpdate {
	ausu.mutation.SetGoogleSecret(s)
	return ausu
}

// SetNillableGoogleSecret sets the "google_secret" field if the given value is not nil.
func (ausu *AppUserSecretUpdate) SetNillableGoogleSecret(s *string) *AppUserSecretUpdate {
	if s != nil {
		ausu.SetGoogleSecret(*s)
	}
	return ausu
}

// Mutation returns the AppUserSecretMutation object of the builder.
func (ausu *AppUserSecretUpdate) Mutation() *AppUserSecretMutation {
	return ausu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ausu *AppUserSecretUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ausu.defaults(); err != nil {
		return 0, err
	}
	if len(ausu.hooks) == 0 {
		affected, err = ausu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserSecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ausu.mutation = mutation
			affected, err = ausu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ausu.hooks) - 1; i >= 0; i-- {
			if ausu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ausu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ausu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ausu *AppUserSecretUpdate) SaveX(ctx context.Context) int {
	affected, err := ausu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ausu *AppUserSecretUpdate) Exec(ctx context.Context) error {
	_, err := ausu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ausu *AppUserSecretUpdate) ExecX(ctx context.Context) {
	if err := ausu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ausu *AppUserSecretUpdate) defaults() error {
	if _, ok := ausu.mutation.UpdatedAt(); !ok {
		if appusersecret.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appusersecret.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appusersecret.UpdateDefaultUpdatedAt()
		ausu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ausu *AppUserSecretUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppUserSecretUpdate {
	ausu.modifiers = append(ausu.modifiers, modifiers...)
	return ausu
}

func (ausu *AppUserSecretUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appusersecret.Table,
			Columns: appusersecret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appusersecret.FieldID,
			},
		},
	}
	if ps := ausu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ausu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldCreatedAt,
		})
	}
	if value, ok := ausu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldCreatedAt,
		})
	}
	if value, ok := ausu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldUpdatedAt,
		})
	}
	if value, ok := ausu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldUpdatedAt,
		})
	}
	if value, ok := ausu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldDeletedAt,
		})
	}
	if value, ok := ausu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldDeletedAt,
		})
	}
	if value, ok := ausu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appusersecret.FieldEntID,
		})
	}
	if value, ok := ausu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appusersecret.FieldAppID,
		})
	}
	if value, ok := ausu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appusersecret.FieldUserID,
		})
	}
	if value, ok := ausu.mutation.PasswordHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appusersecret.FieldPasswordHash,
		})
	}
	if value, ok := ausu.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appusersecret.FieldSalt,
		})
	}
	if value, ok := ausu.mutation.GoogleSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appusersecret.FieldGoogleSecret,
		})
	}
	_spec.Modifiers = ausu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ausu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appusersecret.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppUserSecretUpdateOne is the builder for updating a single AppUserSecret entity.
type AppUserSecretUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppUserSecretMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ausuo *AppUserSecretUpdateOne) SetCreatedAt(u uint32) *AppUserSecretUpdateOne {
	ausuo.mutation.ResetCreatedAt()
	ausuo.mutation.SetCreatedAt(u)
	return ausuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ausuo *AppUserSecretUpdateOne) SetNillableCreatedAt(u *uint32) *AppUserSecretUpdateOne {
	if u != nil {
		ausuo.SetCreatedAt(*u)
	}
	return ausuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ausuo *AppUserSecretUpdateOne) AddCreatedAt(u int32) *AppUserSecretUpdateOne {
	ausuo.mutation.AddCreatedAt(u)
	return ausuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ausuo *AppUserSecretUpdateOne) SetUpdatedAt(u uint32) *AppUserSecretUpdateOne {
	ausuo.mutation.ResetUpdatedAt()
	ausuo.mutation.SetUpdatedAt(u)
	return ausuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ausuo *AppUserSecretUpdateOne) AddUpdatedAt(u int32) *AppUserSecretUpdateOne {
	ausuo.mutation.AddUpdatedAt(u)
	return ausuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ausuo *AppUserSecretUpdateOne) SetDeletedAt(u uint32) *AppUserSecretUpdateOne {
	ausuo.mutation.ResetDeletedAt()
	ausuo.mutation.SetDeletedAt(u)
	return ausuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ausuo *AppUserSecretUpdateOne) SetNillableDeletedAt(u *uint32) *AppUserSecretUpdateOne {
	if u != nil {
		ausuo.SetDeletedAt(*u)
	}
	return ausuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ausuo *AppUserSecretUpdateOne) AddDeletedAt(u int32) *AppUserSecretUpdateOne {
	ausuo.mutation.AddDeletedAt(u)
	return ausuo
}

// SetEntID sets the "ent_id" field.
func (ausuo *AppUserSecretUpdateOne) SetEntID(u uuid.UUID) *AppUserSecretUpdateOne {
	ausuo.mutation.SetEntID(u)
	return ausuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ausuo *AppUserSecretUpdateOne) SetNillableEntID(u *uuid.UUID) *AppUserSecretUpdateOne {
	if u != nil {
		ausuo.SetEntID(*u)
	}
	return ausuo
}

// SetAppID sets the "app_id" field.
func (ausuo *AppUserSecretUpdateOne) SetAppID(u uuid.UUID) *AppUserSecretUpdateOne {
	ausuo.mutation.SetAppID(u)
	return ausuo
}

// SetUserID sets the "user_id" field.
func (ausuo *AppUserSecretUpdateOne) SetUserID(u uuid.UUID) *AppUserSecretUpdateOne {
	ausuo.mutation.SetUserID(u)
	return ausuo
}

// SetPasswordHash sets the "password_hash" field.
func (ausuo *AppUserSecretUpdateOne) SetPasswordHash(s string) *AppUserSecretUpdateOne {
	ausuo.mutation.SetPasswordHash(s)
	return ausuo
}

// SetSalt sets the "salt" field.
func (ausuo *AppUserSecretUpdateOne) SetSalt(s string) *AppUserSecretUpdateOne {
	ausuo.mutation.SetSalt(s)
	return ausuo
}

// SetGoogleSecret sets the "google_secret" field.
func (ausuo *AppUserSecretUpdateOne) SetGoogleSecret(s string) *AppUserSecretUpdateOne {
	ausuo.mutation.SetGoogleSecret(s)
	return ausuo
}

// SetNillableGoogleSecret sets the "google_secret" field if the given value is not nil.
func (ausuo *AppUserSecretUpdateOne) SetNillableGoogleSecret(s *string) *AppUserSecretUpdateOne {
	if s != nil {
		ausuo.SetGoogleSecret(*s)
	}
	return ausuo
}

// Mutation returns the AppUserSecretMutation object of the builder.
func (ausuo *AppUserSecretUpdateOne) Mutation() *AppUserSecretMutation {
	return ausuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ausuo *AppUserSecretUpdateOne) Select(field string, fields ...string) *AppUserSecretUpdateOne {
	ausuo.fields = append([]string{field}, fields...)
	return ausuo
}

// Save executes the query and returns the updated AppUserSecret entity.
func (ausuo *AppUserSecretUpdateOne) Save(ctx context.Context) (*AppUserSecret, error) {
	var (
		err  error
		node *AppUserSecret
	)
	if err := ausuo.defaults(); err != nil {
		return nil, err
	}
	if len(ausuo.hooks) == 0 {
		node, err = ausuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserSecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ausuo.mutation = mutation
			node, err = ausuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ausuo.hooks) - 1; i >= 0; i-- {
			if ausuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ausuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ausuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppUserSecret)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppUserSecretMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ausuo *AppUserSecretUpdateOne) SaveX(ctx context.Context) *AppUserSecret {
	node, err := ausuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ausuo *AppUserSecretUpdateOne) Exec(ctx context.Context) error {
	_, err := ausuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ausuo *AppUserSecretUpdateOne) ExecX(ctx context.Context) {
	if err := ausuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ausuo *AppUserSecretUpdateOne) defaults() error {
	if _, ok := ausuo.mutation.UpdatedAt(); !ok {
		if appusersecret.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appusersecret.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appusersecret.UpdateDefaultUpdatedAt()
		ausuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ausuo *AppUserSecretUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppUserSecretUpdateOne {
	ausuo.modifiers = append(ausuo.modifiers, modifiers...)
	return ausuo
}

func (ausuo *AppUserSecretUpdateOne) sqlSave(ctx context.Context) (_node *AppUserSecret, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appusersecret.Table,
			Columns: appusersecret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appusersecret.FieldID,
			},
		},
	}
	id, ok := ausuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppUserSecret.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ausuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appusersecret.FieldID)
		for _, f := range fields {
			if !appusersecret.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appusersecret.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ausuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ausuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldCreatedAt,
		})
	}
	if value, ok := ausuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldCreatedAt,
		})
	}
	if value, ok := ausuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldUpdatedAt,
		})
	}
	if value, ok := ausuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldUpdatedAt,
		})
	}
	if value, ok := ausuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldDeletedAt,
		})
	}
	if value, ok := ausuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusersecret.FieldDeletedAt,
		})
	}
	if value, ok := ausuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appusersecret.FieldEntID,
		})
	}
	if value, ok := ausuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appusersecret.FieldAppID,
		})
	}
	if value, ok := ausuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appusersecret.FieldUserID,
		})
	}
	if value, ok := ausuo.mutation.PasswordHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appusersecret.FieldPasswordHash,
		})
	}
	if value, ok := ausuo.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appusersecret.FieldSalt,
		})
	}
	if value, ok := ausuo.mutation.GoogleSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appusersecret.FieldGoogleSecret,
		})
	}
	_spec.Modifiers = ausuo.modifiers
	_node = &AppUserSecret{config: ausuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ausuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appusersecret.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
