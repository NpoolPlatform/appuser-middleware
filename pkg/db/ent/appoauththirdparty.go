// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appoauththirdparty"
	"github.com/google/uuid"
)

// AppOAuthThirdParty is the model entity for the AppOAuthThirdParty schema.
type AppOAuthThirdParty struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// ThirdPartyID holds the value of the "third_party_id" field.
	ThirdPartyID uuid.UUID `json:"third_party_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppOAuthThirdParty) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appoauththirdparty.FieldCreatedAt, appoauththirdparty.FieldUpdatedAt, appoauththirdparty.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case appoauththirdparty.FieldID, appoauththirdparty.FieldAppID, appoauththirdparty.FieldThirdPartyID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppOAuthThirdParty", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppOAuthThirdParty fields.
func (aotp *AppOAuthThirdParty) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appoauththirdparty.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				aotp.ID = *value
			}
		case appoauththirdparty.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				aotp.CreatedAt = uint32(value.Int64)
			}
		case appoauththirdparty.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				aotp.UpdatedAt = uint32(value.Int64)
			}
		case appoauththirdparty.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				aotp.DeletedAt = uint32(value.Int64)
			}
		case appoauththirdparty.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				aotp.AppID = *value
			}
		case appoauththirdparty.FieldThirdPartyID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field third_party_id", values[i])
			} else if value != nil {
				aotp.ThirdPartyID = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppOAuthThirdParty.
// Note that you need to call AppOAuthThirdParty.Unwrap() before calling this method if this AppOAuthThirdParty
// was returned from a transaction, and the transaction was committed or rolled back.
func (aotp *AppOAuthThirdParty) Update() *AppOAuthThirdPartyUpdateOne {
	return (&AppOAuthThirdPartyClient{config: aotp.config}).UpdateOne(aotp)
}

// Unwrap unwraps the AppOAuthThirdParty entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aotp *AppOAuthThirdParty) Unwrap() *AppOAuthThirdParty {
	_tx, ok := aotp.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppOAuthThirdParty is not a transactional entity")
	}
	aotp.config.driver = _tx.drv
	return aotp
}

// String implements the fmt.Stringer.
func (aotp *AppOAuthThirdParty) String() string {
	var builder strings.Builder
	builder.WriteString("AppOAuthThirdParty(")
	builder.WriteString(fmt.Sprintf("id=%v, ", aotp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", aotp.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", aotp.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", aotp.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", aotp.AppID))
	builder.WriteString(", ")
	builder.WriteString("third_party_id=")
	builder.WriteString(fmt.Sprintf("%v", aotp.ThirdPartyID))
	builder.WriteByte(')')
	return builder.String()
}

// AppOAuthThirdParties is a parsable slice of AppOAuthThirdParty.
type AppOAuthThirdParties []*AppOAuthThirdParty

func (aotp AppOAuthThirdParties) config(cfg config) {
	for _i := range aotp {
		aotp[_i].config = cfg
	}
}
