// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuserextra"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppUserExtra is the model entity for the AppUserExtra schema.
type AppUserExtra struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// AddressFields holds the value of the "address_fields" field.
	AddressFields []string `json:"address_fields,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// PostalCode holds the value of the "postal_code" field.
	PostalCode string `json:"postal_code,omitempty"`
	// Age holds the value of the "age" field.
	Age uint32 `json:"age,omitempty"`
	// Birthday holds the value of the "birthday" field.
	Birthday uint32 `json:"birthday,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Organization holds the value of the "organization" field.
	Organization string `json:"organization,omitempty"`
	// IDNumber holds the value of the "id_number" field.
	IDNumber string `json:"id_number,omitempty"`
	// ActionCredits holds the value of the "action_credits" field.
	ActionCredits decimal.Decimal `json:"action_credits,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppUserExtra) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appuserextra.FieldAddressFields:
			values[i] = new([]byte)
		case appuserextra.FieldActionCredits:
			values[i] = new(decimal.Decimal)
		case appuserextra.FieldID, appuserextra.FieldCreatedAt, appuserextra.FieldUpdatedAt, appuserextra.FieldDeletedAt, appuserextra.FieldAge, appuserextra.FieldBirthday:
			values[i] = new(sql.NullInt64)
		case appuserextra.FieldUsername, appuserextra.FieldFirstName, appuserextra.FieldLastName, appuserextra.FieldGender, appuserextra.FieldPostalCode, appuserextra.FieldAvatar, appuserextra.FieldOrganization, appuserextra.FieldIDNumber:
			values[i] = new(sql.NullString)
		case appuserextra.FieldEntID, appuserextra.FieldAppID, appuserextra.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppUserExtra", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppUserExtra fields.
func (aue *AppUserExtra) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appuserextra.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aue.ID = uint32(value.Int64)
		case appuserextra.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				aue.CreatedAt = uint32(value.Int64)
			}
		case appuserextra.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				aue.UpdatedAt = uint32(value.Int64)
			}
		case appuserextra.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				aue.DeletedAt = uint32(value.Int64)
			}
		case appuserextra.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				aue.EntID = *value
			}
		case appuserextra.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				aue.AppID = *value
			}
		case appuserextra.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				aue.UserID = *value
			}
		case appuserextra.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				aue.Username = value.String
			}
		case appuserextra.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				aue.FirstName = value.String
			}
		case appuserextra.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				aue.LastName = value.String
			}
		case appuserextra.FieldAddressFields:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field address_fields", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &aue.AddressFields); err != nil {
					return fmt.Errorf("unmarshal field address_fields: %w", err)
				}
			}
		case appuserextra.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				aue.Gender = value.String
			}
		case appuserextra.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_code", values[i])
			} else if value.Valid {
				aue.PostalCode = value.String
			}
		case appuserextra.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				aue.Age = uint32(value.Int64)
			}
		case appuserextra.FieldBirthday:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				aue.Birthday = uint32(value.Int64)
			}
		case appuserextra.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				aue.Avatar = value.String
			}
		case appuserextra.FieldOrganization:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization", values[i])
			} else if value.Valid {
				aue.Organization = value.String
			}
		case appuserextra.FieldIDNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id_number", values[i])
			} else if value.Valid {
				aue.IDNumber = value.String
			}
		case appuserextra.FieldActionCredits:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field action_credits", values[i])
			} else if value != nil {
				aue.ActionCredits = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppUserExtra.
// Note that you need to call AppUserExtra.Unwrap() before calling this method if this AppUserExtra
// was returned from a transaction, and the transaction was committed or rolled back.
func (aue *AppUserExtra) Update() *AppUserExtraUpdateOne {
	return (&AppUserExtraClient{config: aue.config}).UpdateOne(aue)
}

// Unwrap unwraps the AppUserExtra entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aue *AppUserExtra) Unwrap() *AppUserExtra {
	_tx, ok := aue.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppUserExtra is not a transactional entity")
	}
	aue.config.driver = _tx.drv
	return aue
}

// String implements the fmt.Stringer.
func (aue *AppUserExtra) String() string {
	var builder strings.Builder
	builder.WriteString("AppUserExtra(")
	builder.WriteString(fmt.Sprintf("id=%v, ", aue.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", aue.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", aue.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", aue.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", aue.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", aue.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", aue.UserID))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(aue.Username)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(aue.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(aue.LastName)
	builder.WriteString(", ")
	builder.WriteString("address_fields=")
	builder.WriteString(fmt.Sprintf("%v", aue.AddressFields))
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(aue.Gender)
	builder.WriteString(", ")
	builder.WriteString("postal_code=")
	builder.WriteString(aue.PostalCode)
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", aue.Age))
	builder.WriteString(", ")
	builder.WriteString("birthday=")
	builder.WriteString(fmt.Sprintf("%v", aue.Birthday))
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(aue.Avatar)
	builder.WriteString(", ")
	builder.WriteString("organization=")
	builder.WriteString(aue.Organization)
	builder.WriteString(", ")
	builder.WriteString("id_number=")
	builder.WriteString(aue.IDNumber)
	builder.WriteString(", ")
	builder.WriteString("action_credits=")
	builder.WriteString(fmt.Sprintf("%v", aue.ActionCredits))
	builder.WriteByte(')')
	return builder.String()
}

// AppUserExtras is a parsable slice of AppUserExtra.
type AppUserExtras []*AppUserExtra

func (aue AppUserExtras) config(cfg config) {
	for _i := range aue {
		aue[_i].config = cfg
	}
}
