// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"
	"github.com/google/uuid"
)

// AppRoleCreate is the builder for creating a AppRole entity.
type AppRoleCreate struct {
	config
	mutation *AppRoleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (arc *AppRoleCreate) SetCreatedAt(u uint32) *AppRoleCreate {
	arc.mutation.SetCreatedAt(u)
	return arc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableCreatedAt(u *uint32) *AppRoleCreate {
	if u != nil {
		arc.SetCreatedAt(*u)
	}
	return arc
}

// SetUpdatedAt sets the "updated_at" field.
func (arc *AppRoleCreate) SetUpdatedAt(u uint32) *AppRoleCreate {
	arc.mutation.SetUpdatedAt(u)
	return arc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableUpdatedAt(u *uint32) *AppRoleCreate {
	if u != nil {
		arc.SetUpdatedAt(*u)
	}
	return arc
}

// SetDeletedAt sets the "deleted_at" field.
func (arc *AppRoleCreate) SetDeletedAt(u uint32) *AppRoleCreate {
	arc.mutation.SetDeletedAt(u)
	return arc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableDeletedAt(u *uint32) *AppRoleCreate {
	if u != nil {
		arc.SetDeletedAt(*u)
	}
	return arc
}

// SetCreatedBy sets the "created_by" field.
func (arc *AppRoleCreate) SetCreatedBy(u uuid.UUID) *AppRoleCreate {
	arc.mutation.SetCreatedBy(u)
	return arc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableCreatedBy(u *uuid.UUID) *AppRoleCreate {
	if u != nil {
		arc.SetCreatedBy(*u)
	}
	return arc
}

// SetRole sets the "role" field.
func (arc *AppRoleCreate) SetRole(s string) *AppRoleCreate {
	arc.mutation.SetRole(s)
	return arc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableRole(s *string) *AppRoleCreate {
	if s != nil {
		arc.SetRole(*s)
	}
	return arc
}

// SetDescription sets the "description" field.
func (arc *AppRoleCreate) SetDescription(s string) *AppRoleCreate {
	arc.mutation.SetDescription(s)
	return arc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableDescription(s *string) *AppRoleCreate {
	if s != nil {
		arc.SetDescription(*s)
	}
	return arc
}

// SetAppID sets the "app_id" field.
func (arc *AppRoleCreate) SetAppID(u uuid.UUID) *AppRoleCreate {
	arc.mutation.SetAppID(u)
	return arc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableAppID(u *uuid.UUID) *AppRoleCreate {
	if u != nil {
		arc.SetAppID(*u)
	}
	return arc
}

// SetDefault sets the "default" field.
func (arc *AppRoleCreate) SetDefault(b bool) *AppRoleCreate {
	arc.mutation.SetDefault(b)
	return arc
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableDefault(b *bool) *AppRoleCreate {
	if b != nil {
		arc.SetDefault(*b)
	}
	return arc
}

// SetGenesis sets the "genesis" field.
func (arc *AppRoleCreate) SetGenesis(b bool) *AppRoleCreate {
	arc.mutation.SetGenesis(b)
	return arc
}

// SetNillableGenesis sets the "genesis" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableGenesis(b *bool) *AppRoleCreate {
	if b != nil {
		arc.SetGenesis(*b)
	}
	return arc
}

// SetID sets the "id" field.
func (arc *AppRoleCreate) SetID(u uuid.UUID) *AppRoleCreate {
	arc.mutation.SetID(u)
	return arc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableID(u *uuid.UUID) *AppRoleCreate {
	if u != nil {
		arc.SetID(*u)
	}
	return arc
}

// Mutation returns the AppRoleMutation object of the builder.
func (arc *AppRoleCreate) Mutation() *AppRoleMutation {
	return arc.mutation
}

// Save creates the AppRole in the database.
func (arc *AppRoleCreate) Save(ctx context.Context) (*AppRole, error) {
	var (
		err  error
		node *AppRole
	)
	if err := arc.defaults(); err != nil {
		return nil, err
	}
	if len(arc.hooks) == 0 {
		if err = arc.check(); err != nil {
			return nil, err
		}
		node, err = arc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = arc.check(); err != nil {
				return nil, err
			}
			arc.mutation = mutation
			if node, err = arc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(arc.hooks) - 1; i >= 0; i-- {
			if arc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = arc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, arc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppRole)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppRoleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AppRoleCreate) SaveX(ctx context.Context) *AppRole {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AppRoleCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AppRoleCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *AppRoleCreate) defaults() error {
	if _, ok := arc.mutation.CreatedAt(); !ok {
		if approle.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized approle.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := approle.DefaultCreatedAt()
		arc.mutation.SetCreatedAt(v)
	}
	if _, ok := arc.mutation.UpdatedAt(); !ok {
		if approle.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized approle.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := approle.DefaultUpdatedAt()
		arc.mutation.SetUpdatedAt(v)
	}
	if _, ok := arc.mutation.DeletedAt(); !ok {
		if approle.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized approle.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := approle.DefaultDeletedAt()
		arc.mutation.SetDeletedAt(v)
	}
	if _, ok := arc.mutation.CreatedBy(); !ok {
		if approle.DefaultCreatedBy == nil {
			return fmt.Errorf("ent: uninitialized approle.DefaultCreatedBy (forgotten import ent/runtime?)")
		}
		v := approle.DefaultCreatedBy()
		arc.mutation.SetCreatedBy(v)
	}
	if _, ok := arc.mutation.Role(); !ok {
		v := approle.DefaultRole
		arc.mutation.SetRole(v)
	}
	if _, ok := arc.mutation.Description(); !ok {
		v := approle.DefaultDescription
		arc.mutation.SetDescription(v)
	}
	if _, ok := arc.mutation.AppID(); !ok {
		if approle.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized approle.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := approle.DefaultAppID()
		arc.mutation.SetAppID(v)
	}
	if _, ok := arc.mutation.Default(); !ok {
		v := approle.DefaultDefault
		arc.mutation.SetDefault(v)
	}
	if _, ok := arc.mutation.Genesis(); !ok {
		v := approle.DefaultGenesis
		arc.mutation.SetGenesis(v)
	}
	if _, ok := arc.mutation.ID(); !ok {
		if approle.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized approle.DefaultID (forgotten import ent/runtime?)")
		}
		v := approle.DefaultID()
		arc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (arc *AppRoleCreate) check() error {
	if _, ok := arc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppRole.created_at"`)}
	}
	if _, ok := arc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppRole.updated_at"`)}
	}
	if _, ok := arc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppRole.deleted_at"`)}
	}
	return nil
}

func (arc *AppRoleCreate) sqlSave(ctx context.Context) (*AppRole, error) {
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (arc *AppRoleCreate) createSpec() (*AppRole, *sqlgraph.CreateSpec) {
	var (
		_node = &AppRole{config: arc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: approle.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: approle.FieldID,
			},
		}
	)
	_spec.OnConflict = arc.conflict
	if id, ok := arc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := arc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: approle.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := arc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: approle.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := arc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: approle.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := arc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: approle.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := arc.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: approle.FieldRole,
		})
		_node.Role = value
	}
	if value, ok := arc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: approle.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := arc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: approle.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := arc.mutation.Default(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: approle.FieldDefault,
		})
		_node.Default = value
	}
	if value, ok := arc.mutation.Genesis(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: approle.FieldGenesis,
		})
		_node.Genesis = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppRole.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppRoleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (arc *AppRoleCreate) OnConflict(opts ...sql.ConflictOption) *AppRoleUpsertOne {
	arc.conflict = opts
	return &AppRoleUpsertOne{
		create: arc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppRole.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (arc *AppRoleCreate) OnConflictColumns(columns ...string) *AppRoleUpsertOne {
	arc.conflict = append(arc.conflict, sql.ConflictColumns(columns...))
	return &AppRoleUpsertOne{
		create: arc,
	}
}

type (
	// AppRoleUpsertOne is the builder for "upsert"-ing
	//  one AppRole node.
	AppRoleUpsertOne struct {
		create *AppRoleCreate
	}

	// AppRoleUpsert is the "OnConflict" setter.
	AppRoleUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppRoleUpsert) SetCreatedAt(v uint32) *AppRoleUpsert {
	u.Set(approle.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateCreatedAt() *AppRoleUpsert {
	u.SetExcluded(approle.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppRoleUpsert) AddCreatedAt(v uint32) *AppRoleUpsert {
	u.Add(approle.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppRoleUpsert) SetUpdatedAt(v uint32) *AppRoleUpsert {
	u.Set(approle.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateUpdatedAt() *AppRoleUpsert {
	u.SetExcluded(approle.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppRoleUpsert) AddUpdatedAt(v uint32) *AppRoleUpsert {
	u.Add(approle.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppRoleUpsert) SetDeletedAt(v uint32) *AppRoleUpsert {
	u.Set(approle.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateDeletedAt() *AppRoleUpsert {
	u.SetExcluded(approle.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppRoleUpsert) AddDeletedAt(v uint32) *AppRoleUpsert {
	u.Add(approle.FieldDeletedAt, v)
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *AppRoleUpsert) SetCreatedBy(v uuid.UUID) *AppRoleUpsert {
	u.Set(approle.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateCreatedBy() *AppRoleUpsert {
	u.SetExcluded(approle.FieldCreatedBy)
	return u
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *AppRoleUpsert) ClearCreatedBy() *AppRoleUpsert {
	u.SetNull(approle.FieldCreatedBy)
	return u
}

// SetRole sets the "role" field.
func (u *AppRoleUpsert) SetRole(v string) *AppRoleUpsert {
	u.Set(approle.FieldRole, v)
	return u
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateRole() *AppRoleUpsert {
	u.SetExcluded(approle.FieldRole)
	return u
}

// ClearRole clears the value of the "role" field.
func (u *AppRoleUpsert) ClearRole() *AppRoleUpsert {
	u.SetNull(approle.FieldRole)
	return u
}

// SetDescription sets the "description" field.
func (u *AppRoleUpsert) SetDescription(v string) *AppRoleUpsert {
	u.Set(approle.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateDescription() *AppRoleUpsert {
	u.SetExcluded(approle.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *AppRoleUpsert) ClearDescription() *AppRoleUpsert {
	u.SetNull(approle.FieldDescription)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppRoleUpsert) SetAppID(v uuid.UUID) *AppRoleUpsert {
	u.Set(approle.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateAppID() *AppRoleUpsert {
	u.SetExcluded(approle.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppRoleUpsert) ClearAppID() *AppRoleUpsert {
	u.SetNull(approle.FieldAppID)
	return u
}

// SetDefault sets the "default" field.
func (u *AppRoleUpsert) SetDefault(v bool) *AppRoleUpsert {
	u.Set(approle.FieldDefault, v)
	return u
}

// UpdateDefault sets the "default" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateDefault() *AppRoleUpsert {
	u.SetExcluded(approle.FieldDefault)
	return u
}

// ClearDefault clears the value of the "default" field.
func (u *AppRoleUpsert) ClearDefault() *AppRoleUpsert {
	u.SetNull(approle.FieldDefault)
	return u
}

// SetGenesis sets the "genesis" field.
func (u *AppRoleUpsert) SetGenesis(v bool) *AppRoleUpsert {
	u.Set(approle.FieldGenesis, v)
	return u
}

// UpdateGenesis sets the "genesis" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateGenesis() *AppRoleUpsert {
	u.SetExcluded(approle.FieldGenesis)
	return u
}

// ClearGenesis clears the value of the "genesis" field.
func (u *AppRoleUpsert) ClearGenesis() *AppRoleUpsert {
	u.SetNull(approle.FieldGenesis)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppRole.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(approle.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppRoleUpsertOne) UpdateNewValues() *AppRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(approle.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppRole.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppRoleUpsertOne) Ignore() *AppRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppRoleUpsertOne) DoNothing() *AppRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppRoleCreate.OnConflict
// documentation for more info.
func (u *AppRoleUpsertOne) Update(set func(*AppRoleUpsert)) *AppRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppRoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppRoleUpsertOne) SetCreatedAt(v uint32) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppRoleUpsertOne) AddCreatedAt(v uint32) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateCreatedAt() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppRoleUpsertOne) SetUpdatedAt(v uint32) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppRoleUpsertOne) AddUpdatedAt(v uint32) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateUpdatedAt() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppRoleUpsertOne) SetDeletedAt(v uint32) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppRoleUpsertOne) AddDeletedAt(v uint32) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateDeletedAt() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *AppRoleUpsertOne) SetCreatedBy(v uuid.UUID) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateCreatedBy() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *AppRoleUpsertOne) ClearCreatedBy() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearCreatedBy()
	})
}

// SetRole sets the "role" field.
func (u *AppRoleUpsertOne) SetRole(v string) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetRole(v)
	})
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateRole() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateRole()
	})
}

// ClearRole clears the value of the "role" field.
func (u *AppRoleUpsertOne) ClearRole() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearRole()
	})
}

// SetDescription sets the "description" field.
func (u *AppRoleUpsertOne) SetDescription(v string) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateDescription() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *AppRoleUpsertOne) ClearDescription() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearDescription()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppRoleUpsertOne) SetAppID(v uuid.UUID) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateAppID() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppRoleUpsertOne) ClearAppID() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearAppID()
	})
}

// SetDefault sets the "default" field.
func (u *AppRoleUpsertOne) SetDefault(v bool) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetDefault(v)
	})
}

// UpdateDefault sets the "default" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateDefault() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateDefault()
	})
}

// ClearDefault clears the value of the "default" field.
func (u *AppRoleUpsertOne) ClearDefault() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearDefault()
	})
}

// SetGenesis sets the "genesis" field.
func (u *AppRoleUpsertOne) SetGenesis(v bool) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetGenesis(v)
	})
}

// UpdateGenesis sets the "genesis" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateGenesis() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateGenesis()
	})
}

// ClearGenesis clears the value of the "genesis" field.
func (u *AppRoleUpsertOne) ClearGenesis() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearGenesis()
	})
}

// Exec executes the query.
func (u *AppRoleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppRoleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppRoleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppRoleUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppRoleUpsertOne.ID is not supported by MySQL driver. Use AppRoleUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppRoleUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppRoleCreateBulk is the builder for creating many AppRole entities in bulk.
type AppRoleCreateBulk struct {
	config
	builders []*AppRoleCreate
	conflict []sql.ConflictOption
}

// Save creates the AppRole entities in the database.
func (arcb *AppRoleCreateBulk) Save(ctx context.Context) ([]*AppRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AppRole, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = arcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AppRoleCreateBulk) SaveX(ctx context.Context) []*AppRole {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AppRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AppRoleCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppRole.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppRoleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (arcb *AppRoleCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppRoleUpsertBulk {
	arcb.conflict = opts
	return &AppRoleUpsertBulk{
		create: arcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppRole.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (arcb *AppRoleCreateBulk) OnConflictColumns(columns ...string) *AppRoleUpsertBulk {
	arcb.conflict = append(arcb.conflict, sql.ConflictColumns(columns...))
	return &AppRoleUpsertBulk{
		create: arcb,
	}
}

// AppRoleUpsertBulk is the builder for "upsert"-ing
// a bulk of AppRole nodes.
type AppRoleUpsertBulk struct {
	create *AppRoleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppRole.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(approle.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppRoleUpsertBulk) UpdateNewValues() *AppRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(approle.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppRole.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppRoleUpsertBulk) Ignore() *AppRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppRoleUpsertBulk) DoNothing() *AppRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppRoleCreateBulk.OnConflict
// documentation for more info.
func (u *AppRoleUpsertBulk) Update(set func(*AppRoleUpsert)) *AppRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppRoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppRoleUpsertBulk) SetCreatedAt(v uint32) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppRoleUpsertBulk) AddCreatedAt(v uint32) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateCreatedAt() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppRoleUpsertBulk) SetUpdatedAt(v uint32) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppRoleUpsertBulk) AddUpdatedAt(v uint32) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateUpdatedAt() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppRoleUpsertBulk) SetDeletedAt(v uint32) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppRoleUpsertBulk) AddDeletedAt(v uint32) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateDeletedAt() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *AppRoleUpsertBulk) SetCreatedBy(v uuid.UUID) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateCreatedBy() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *AppRoleUpsertBulk) ClearCreatedBy() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearCreatedBy()
	})
}

// SetRole sets the "role" field.
func (u *AppRoleUpsertBulk) SetRole(v string) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetRole(v)
	})
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateRole() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateRole()
	})
}

// ClearRole clears the value of the "role" field.
func (u *AppRoleUpsertBulk) ClearRole() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearRole()
	})
}

// SetDescription sets the "description" field.
func (u *AppRoleUpsertBulk) SetDescription(v string) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateDescription() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *AppRoleUpsertBulk) ClearDescription() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearDescription()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppRoleUpsertBulk) SetAppID(v uuid.UUID) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateAppID() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppRoleUpsertBulk) ClearAppID() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearAppID()
	})
}

// SetDefault sets the "default" field.
func (u *AppRoleUpsertBulk) SetDefault(v bool) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetDefault(v)
	})
}

// UpdateDefault sets the "default" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateDefault() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateDefault()
	})
}

// ClearDefault clears the value of the "default" field.
func (u *AppRoleUpsertBulk) ClearDefault() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearDefault()
	})
}

// SetGenesis sets the "genesis" field.
func (u *AppRoleUpsertBulk) SetGenesis(v bool) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetGenesis(v)
	})
}

// UpdateGenesis sets the "genesis" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateGenesis() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateGenesis()
	})
}

// ClearGenesis clears the value of the "genesis" field.
func (u *AppRoleUpsertBulk) ClearGenesis() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearGenesis()
	})
}

// Exec executes the query.
func (u *AppRoleUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppRoleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppRoleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppRoleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}