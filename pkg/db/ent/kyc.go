// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/kyc"
	"github.com/google/uuid"
)

// Kyc is the model entity for the Kyc schema.
type Kyc struct {
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
	// DocumentType holds the value of the "document_type" field.
	DocumentType string `json:"document_type,omitempty"`
	// IDNumber holds the value of the "id_number" field.
	IDNumber string `json:"id_number,omitempty"`
	// FrontImg holds the value of the "front_img" field.
	FrontImg string `json:"front_img,omitempty"`
	// BackImg holds the value of the "back_img" field.
	BackImg string `json:"back_img,omitempty"`
	// SelfieImg holds the value of the "selfie_img" field.
	SelfieImg string `json:"selfie_img,omitempty"`
	// EntityType holds the value of the "entity_type" field.
	EntityType string `json:"entity_type,omitempty"`
	// ReviewID holds the value of the "review_id" field.
	ReviewID uuid.UUID `json:"review_id,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Kyc) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case kyc.FieldID, kyc.FieldCreatedAt, kyc.FieldUpdatedAt, kyc.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case kyc.FieldDocumentType, kyc.FieldIDNumber, kyc.FieldFrontImg, kyc.FieldBackImg, kyc.FieldSelfieImg, kyc.FieldEntityType, kyc.FieldState:
			values[i] = new(sql.NullString)
		case kyc.FieldEntID, kyc.FieldAppID, kyc.FieldUserID, kyc.FieldReviewID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Kyc", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Kyc fields.
func (k *Kyc) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case kyc.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			k.ID = uint32(value.Int64)
		case kyc.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				k.CreatedAt = uint32(value.Int64)
			}
		case kyc.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				k.UpdatedAt = uint32(value.Int64)
			}
		case kyc.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				k.DeletedAt = uint32(value.Int64)
			}
		case kyc.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				k.EntID = *value
			}
		case kyc.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				k.AppID = *value
			}
		case kyc.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				k.UserID = *value
			}
		case kyc.FieldDocumentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field document_type", values[i])
			} else if value.Valid {
				k.DocumentType = value.String
			}
		case kyc.FieldIDNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id_number", values[i])
			} else if value.Valid {
				k.IDNumber = value.String
			}
		case kyc.FieldFrontImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field front_img", values[i])
			} else if value.Valid {
				k.FrontImg = value.String
			}
		case kyc.FieldBackImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field back_img", values[i])
			} else if value.Valid {
				k.BackImg = value.String
			}
		case kyc.FieldSelfieImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field selfie_img", values[i])
			} else if value.Valid {
				k.SelfieImg = value.String
			}
		case kyc.FieldEntityType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entity_type", values[i])
			} else if value.Valid {
				k.EntityType = value.String
			}
		case kyc.FieldReviewID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field review_id", values[i])
			} else if value != nil {
				k.ReviewID = *value
			}
		case kyc.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				k.State = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Kyc.
// Note that you need to call Kyc.Unwrap() before calling this method if this Kyc
// was returned from a transaction, and the transaction was committed or rolled back.
func (k *Kyc) Update() *KycUpdateOne {
	return (&KycClient{config: k.config}).UpdateOne(k)
}

// Unwrap unwraps the Kyc entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (k *Kyc) Unwrap() *Kyc {
	_tx, ok := k.config.driver.(*txDriver)
	if !ok {
		panic("ent: Kyc is not a transactional entity")
	}
	k.config.driver = _tx.drv
	return k
}

// String implements the fmt.Stringer.
func (k *Kyc) String() string {
	var builder strings.Builder
	builder.WriteString("Kyc(")
	builder.WriteString(fmt.Sprintf("id=%v, ", k.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", k.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", k.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", k.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", k.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", k.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", k.UserID))
	builder.WriteString(", ")
	builder.WriteString("document_type=")
	builder.WriteString(k.DocumentType)
	builder.WriteString(", ")
	builder.WriteString("id_number=")
	builder.WriteString(k.IDNumber)
	builder.WriteString(", ")
	builder.WriteString("front_img=")
	builder.WriteString(k.FrontImg)
	builder.WriteString(", ")
	builder.WriteString("back_img=")
	builder.WriteString(k.BackImg)
	builder.WriteString(", ")
	builder.WriteString("selfie_img=")
	builder.WriteString(k.SelfieImg)
	builder.WriteString(", ")
	builder.WriteString("entity_type=")
	builder.WriteString(k.EntityType)
	builder.WriteString(", ")
	builder.WriteString("review_id=")
	builder.WriteString(fmt.Sprintf("%v", k.ReviewID))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(k.State)
	builder.WriteByte(')')
	return builder.String()
}

// Kycs is a parsable slice of Kyc.
type Kycs []*Kyc

func (k Kycs) config(cfg config) {
	for _i := range k {
		k[_i].config = cfg
	}
}
