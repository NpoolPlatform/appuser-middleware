// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middware/pkg/db/ent/appuser"
	"github.com/google/uuid"
)

// AppUserCreate is the builder for creating a AppUser entity.
type AppUserCreate struct {
	config
	mutation *AppUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (auc *AppUserCreate) SetCreatedAt(u uint32) *AppUserCreate {
	auc.mutation.SetCreatedAt(u)
	return auc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auc *AppUserCreate) SetNillableCreatedAt(u *uint32) *AppUserCreate {
	if u != nil {
		auc.SetCreatedAt(*u)
	}
	return auc
}

// SetUpdatedAt sets the "updated_at" field.
func (auc *AppUserCreate) SetUpdatedAt(u uint32) *AppUserCreate {
	auc.mutation.SetUpdatedAt(u)
	return auc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auc *AppUserCreate) SetNillableUpdatedAt(u *uint32) *AppUserCreate {
	if u != nil {
		auc.SetUpdatedAt(*u)
	}
	return auc
}

// SetDeletedAt sets the "deleted_at" field.
func (auc *AppUserCreate) SetDeletedAt(u uint32) *AppUserCreate {
	auc.mutation.SetDeletedAt(u)
	return auc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auc *AppUserCreate) SetNillableDeletedAt(u *uint32) *AppUserCreate {
	if u != nil {
		auc.SetDeletedAt(*u)
	}
	return auc
}

// SetAppID sets the "app_id" field.
func (auc *AppUserCreate) SetAppID(u uuid.UUID) *AppUserCreate {
	auc.mutation.SetAppID(u)
	return auc
}

// SetEmailAddress sets the "email_address" field.
func (auc *AppUserCreate) SetEmailAddress(s string) *AppUserCreate {
	auc.mutation.SetEmailAddress(s)
	return auc
}

// SetPhoneNo sets the "phone_no" field.
func (auc *AppUserCreate) SetPhoneNo(s string) *AppUserCreate {
	auc.mutation.SetPhoneNo(s)
	return auc
}

// SetImportFromApp sets the "import_from_app" field.
func (auc *AppUserCreate) SetImportFromApp(u uuid.UUID) *AppUserCreate {
	auc.mutation.SetImportFromApp(u)
	return auc
}

// SetID sets the "id" field.
func (auc *AppUserCreate) SetID(u uuid.UUID) *AppUserCreate {
	auc.mutation.SetID(u)
	return auc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (auc *AppUserCreate) SetNillableID(u *uuid.UUID) *AppUserCreate {
	if u != nil {
		auc.SetID(*u)
	}
	return auc
}

// Mutation returns the AppUserMutation object of the builder.
func (auc *AppUserCreate) Mutation() *AppUserMutation {
	return auc.mutation
}

// Save creates the AppUser in the database.
func (auc *AppUserCreate) Save(ctx context.Context) (*AppUser, error) {
	var (
		err  error
		node *AppUser
	)
	if err := auc.defaults(); err != nil {
		return nil, err
	}
	if len(auc.hooks) == 0 {
		if err = auc.check(); err != nil {
			return nil, err
		}
		node, err = auc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auc.check(); err != nil {
				return nil, err
			}
			auc.mutation = mutation
			if node, err = auc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(auc.hooks) - 1; i >= 0; i-- {
			if auc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (auc *AppUserCreate) SaveX(ctx context.Context) *AppUser {
	v, err := auc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (auc *AppUserCreate) Exec(ctx context.Context) error {
	_, err := auc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auc *AppUserCreate) ExecX(ctx context.Context) {
	if err := auc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auc *AppUserCreate) defaults() error {
	if _, ok := auc.mutation.CreatedAt(); !ok {
		if appuser.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appuser.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appuser.DefaultCreatedAt()
		auc.mutation.SetCreatedAt(v)
	}
	if _, ok := auc.mutation.UpdatedAt(); !ok {
		if appuser.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appuser.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appuser.DefaultUpdatedAt()
		auc.mutation.SetUpdatedAt(v)
	}
	if _, ok := auc.mutation.DeletedAt(); !ok {
		if appuser.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appuser.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appuser.DefaultDeletedAt()
		auc.mutation.SetDeletedAt(v)
	}
	if _, ok := auc.mutation.ID(); !ok {
		if appuser.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized appuser.DefaultID (forgotten import ent/runtime?)")
		}
		v := appuser.DefaultID()
		auc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (auc *AppUserCreate) check() error {
	if _, ok := auc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppUser.created_at"`)}
	}
	if _, ok := auc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppUser.updated_at"`)}
	}
	if _, ok := auc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppUser.deleted_at"`)}
	}
	if _, ok := auc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppUser.app_id"`)}
	}
	if _, ok := auc.mutation.EmailAddress(); !ok {
		return &ValidationError{Name: "email_address", err: errors.New(`ent: missing required field "AppUser.email_address"`)}
	}
	if _, ok := auc.mutation.PhoneNo(); !ok {
		return &ValidationError{Name: "phone_no", err: errors.New(`ent: missing required field "AppUser.phone_no"`)}
	}
	if _, ok := auc.mutation.ImportFromApp(); !ok {
		return &ValidationError{Name: "import_from_app", err: errors.New(`ent: missing required field "AppUser.import_from_app"`)}
	}
	return nil
}

func (auc *AppUserCreate) sqlSave(ctx context.Context) (*AppUser, error) {
	_node, _spec := auc.createSpec()
	if err := sqlgraph.CreateNode(ctx, auc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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

func (auc *AppUserCreate) createSpec() (*AppUser, *sqlgraph.CreateSpec) {
	var (
		_node = &AppUser{config: auc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuser.FieldID,
			},
		}
	)
	_spec.OnConflict = auc.conflict
	if id, ok := auc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := auc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := auc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := auc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := auc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuser.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := auc.mutation.EmailAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuser.FieldEmailAddress,
		})
		_node.EmailAddress = value
	}
	if value, ok := auc.mutation.PhoneNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuser.FieldPhoneNo,
		})
		_node.PhoneNo = value
	}
	if value, ok := auc.mutation.ImportFromApp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuser.FieldImportFromApp,
		})
		_node.ImportFromApp = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUser.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (auc *AppUserCreate) OnConflict(opts ...sql.ConflictOption) *AppUserUpsertOne {
	auc.conflict = opts
	return &AppUserUpsertOne{
		create: auc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (auc *AppUserCreate) OnConflictColumns(columns ...string) *AppUserUpsertOne {
	auc.conflict = append(auc.conflict, sql.ConflictColumns(columns...))
	return &AppUserUpsertOne{
		create: auc,
	}
}

type (
	// AppUserUpsertOne is the builder for "upsert"-ing
	//  one AppUser node.
	AppUserUpsertOne struct {
		create *AppUserCreate
	}

	// AppUserUpsert is the "OnConflict" setter.
	AppUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppUserUpsert) SetCreatedAt(v uint32) *AppUserUpsert {
	u.Set(appuser.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUserUpsert) UpdateCreatedAt() *AppUserUpsert {
	u.SetExcluded(appuser.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppUserUpsert) AddCreatedAt(v uint32) *AppUserUpsert {
	u.Add(appuser.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUserUpsert) SetUpdatedAt(v uint32) *AppUserUpsert {
	u.Set(appuser.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUserUpsert) UpdateUpdatedAt() *AppUserUpsert {
	u.SetExcluded(appuser.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppUserUpsert) AddUpdatedAt(v uint32) *AppUserUpsert {
	u.Add(appuser.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUserUpsert) SetDeletedAt(v uint32) *AppUserUpsert {
	u.Set(appuser.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUserUpsert) UpdateDeletedAt() *AppUserUpsert {
	u.SetExcluded(appuser.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppUserUpsert) AddDeletedAt(v uint32) *AppUserUpsert {
	u.Add(appuser.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppUserUpsert) SetAppID(v uuid.UUID) *AppUserUpsert {
	u.Set(appuser.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserUpsert) UpdateAppID() *AppUserUpsert {
	u.SetExcluded(appuser.FieldAppID)
	return u
}

// SetEmailAddress sets the "email_address" field.
func (u *AppUserUpsert) SetEmailAddress(v string) *AppUserUpsert {
	u.Set(appuser.FieldEmailAddress, v)
	return u
}

// UpdateEmailAddress sets the "email_address" field to the value that was provided on create.
func (u *AppUserUpsert) UpdateEmailAddress() *AppUserUpsert {
	u.SetExcluded(appuser.FieldEmailAddress)
	return u
}

// SetPhoneNo sets the "phone_no" field.
func (u *AppUserUpsert) SetPhoneNo(v string) *AppUserUpsert {
	u.Set(appuser.FieldPhoneNo, v)
	return u
}

// UpdatePhoneNo sets the "phone_no" field to the value that was provided on create.
func (u *AppUserUpsert) UpdatePhoneNo() *AppUserUpsert {
	u.SetExcluded(appuser.FieldPhoneNo)
	return u
}

// SetImportFromApp sets the "import_from_app" field.
func (u *AppUserUpsert) SetImportFromApp(v uuid.UUID) *AppUserUpsert {
	u.Set(appuser.FieldImportFromApp, v)
	return u
}

// UpdateImportFromApp sets the "import_from_app" field to the value that was provided on create.
func (u *AppUserUpsert) UpdateImportFromApp() *AppUserUpsert {
	u.SetExcluded(appuser.FieldImportFromApp)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserUpsertOne) UpdateNewValues() *AppUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appuser.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppUser.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppUserUpsertOne) Ignore() *AppUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserUpsertOne) DoNothing() *AppUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserCreate.OnConflict
// documentation for more info.
func (u *AppUserUpsertOne) Update(set func(*AppUserUpsert)) *AppUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUserUpsertOne) SetCreatedAt(v uint32) *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppUserUpsertOne) AddCreatedAt(v uint32) *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUserUpsertOne) UpdateCreatedAt() *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUserUpsertOne) SetUpdatedAt(v uint32) *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppUserUpsertOne) AddUpdatedAt(v uint32) *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUserUpsertOne) UpdateUpdatedAt() *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUserUpsertOne) SetDeletedAt(v uint32) *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppUserUpsertOne) AddDeletedAt(v uint32) *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUserUpsertOne) UpdateDeletedAt() *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppUserUpsertOne) SetAppID(v uuid.UUID) *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserUpsertOne) UpdateAppID() *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateAppID()
	})
}

// SetEmailAddress sets the "email_address" field.
func (u *AppUserUpsertOne) SetEmailAddress(v string) *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.SetEmailAddress(v)
	})
}

// UpdateEmailAddress sets the "email_address" field to the value that was provided on create.
func (u *AppUserUpsertOne) UpdateEmailAddress() *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateEmailAddress()
	})
}

// SetPhoneNo sets the "phone_no" field.
func (u *AppUserUpsertOne) SetPhoneNo(v string) *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.SetPhoneNo(v)
	})
}

// UpdatePhoneNo sets the "phone_no" field to the value that was provided on create.
func (u *AppUserUpsertOne) UpdatePhoneNo() *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdatePhoneNo()
	})
}

// SetImportFromApp sets the "import_from_app" field.
func (u *AppUserUpsertOne) SetImportFromApp(v uuid.UUID) *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.SetImportFromApp(v)
	})
}

// UpdateImportFromApp sets the "import_from_app" field to the value that was provided on create.
func (u *AppUserUpsertOne) UpdateImportFromApp() *AppUserUpsertOne {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateImportFromApp()
	})
}

// Exec executes the query.
func (u *AppUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppUserUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppUserUpsertOne.ID is not supported by MySQL driver. Use AppUserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppUserUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppUserCreateBulk is the builder for creating many AppUser entities in bulk.
type AppUserCreateBulk struct {
	config
	builders []*AppUserCreate
	conflict []sql.ConflictOption
}

// Save creates the AppUser entities in the database.
func (aucb *AppUserCreateBulk) Save(ctx context.Context) ([]*AppUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aucb.builders))
	nodes := make([]*AppUser, len(aucb.builders))
	mutators := make([]Mutator, len(aucb.builders))
	for i := range aucb.builders {
		func(i int, root context.Context) {
			builder := aucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppUserMutation)
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
					_, err = mutators[i+1].Mutate(root, aucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = aucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, aucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aucb *AppUserCreateBulk) SaveX(ctx context.Context) []*AppUser {
	v, err := aucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aucb *AppUserCreateBulk) Exec(ctx context.Context) error {
	_, err := aucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aucb *AppUserCreateBulk) ExecX(ctx context.Context) {
	if err := aucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (aucb *AppUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppUserUpsertBulk {
	aucb.conflict = opts
	return &AppUserUpsertBulk{
		create: aucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (aucb *AppUserCreateBulk) OnConflictColumns(columns ...string) *AppUserUpsertBulk {
	aucb.conflict = append(aucb.conflict, sql.ConflictColumns(columns...))
	return &AppUserUpsertBulk{
		create: aucb,
	}
}

// AppUserUpsertBulk is the builder for "upsert"-ing
// a bulk of AppUser nodes.
type AppUserUpsertBulk struct {
	create *AppUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserUpsertBulk) UpdateNewValues() *AppUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appuser.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppUserUpsertBulk) Ignore() *AppUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserUpsertBulk) DoNothing() *AppUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserCreateBulk.OnConflict
// documentation for more info.
func (u *AppUserUpsertBulk) Update(set func(*AppUserUpsert)) *AppUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUserUpsertBulk) SetCreatedAt(v uint32) *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppUserUpsertBulk) AddCreatedAt(v uint32) *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUserUpsertBulk) UpdateCreatedAt() *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUserUpsertBulk) SetUpdatedAt(v uint32) *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppUserUpsertBulk) AddUpdatedAt(v uint32) *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUserUpsertBulk) UpdateUpdatedAt() *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUserUpsertBulk) SetDeletedAt(v uint32) *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppUserUpsertBulk) AddDeletedAt(v uint32) *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUserUpsertBulk) UpdateDeletedAt() *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppUserUpsertBulk) SetAppID(v uuid.UUID) *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserUpsertBulk) UpdateAppID() *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateAppID()
	})
}

// SetEmailAddress sets the "email_address" field.
func (u *AppUserUpsertBulk) SetEmailAddress(v string) *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.SetEmailAddress(v)
	})
}

// UpdateEmailAddress sets the "email_address" field to the value that was provided on create.
func (u *AppUserUpsertBulk) UpdateEmailAddress() *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateEmailAddress()
	})
}

// SetPhoneNo sets the "phone_no" field.
func (u *AppUserUpsertBulk) SetPhoneNo(v string) *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.SetPhoneNo(v)
	})
}

// UpdatePhoneNo sets the "phone_no" field to the value that was provided on create.
func (u *AppUserUpsertBulk) UpdatePhoneNo() *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdatePhoneNo()
	})
}

// SetImportFromApp sets the "import_from_app" field.
func (u *AppUserUpsertBulk) SetImportFromApp(v uuid.UUID) *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.SetImportFromApp(v)
	})
}

// UpdateImportFromApp sets the "import_from_app" field to the value that was provided on create.
func (u *AppUserUpsertBulk) UpdateImportFromApp() *AppUserUpsertBulk {
	return u.Update(func(s *AppUserUpsert) {
		s.UpdateImportFromApp()
	})
}

// Exec executes the query.
func (u *AppUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
