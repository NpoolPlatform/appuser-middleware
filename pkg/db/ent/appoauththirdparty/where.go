// Code generated by ent, DO NOT EDIT.

package appoauththirdparty

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// ThirdPartyID applies equality check predicate on the "third_party_id" field. It's identical to ThirdPartyIDEQ.
func ThirdPartyID(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyID), v))
	})
}

// ClientID applies equality check predicate on the "client_id" field. It's identical to ClientIDEQ.
func ClientID(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientID), v))
	})
}

// ClientSecret applies equality check predicate on the "client_secret" field. It's identical to ClientSecretEQ.
func ClientSecret(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientSecret), v))
	})
}

// CallbackURL applies equality check predicate on the "callback_url" field. It's identical to CallbackURLEQ.
func CallbackURL(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCallbackURL), v))
	})
}

// Salt applies equality check predicate on the "salt" field. It's identical to SaltEQ.
func Salt(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// ThirdPartyIDEQ applies the EQ predicate on the "third_party_id" field.
func ThirdPartyIDEQ(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDNEQ applies the NEQ predicate on the "third_party_id" field.
func ThirdPartyIDNEQ(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDIn applies the In predicate on the "third_party_id" field.
func ThirdPartyIDIn(vs ...uuid.UUID) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldThirdPartyID), v...))
	})
}

// ThirdPartyIDNotIn applies the NotIn predicate on the "third_party_id" field.
func ThirdPartyIDNotIn(vs ...uuid.UUID) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldThirdPartyID), v...))
	})
}

// ThirdPartyIDGT applies the GT predicate on the "third_party_id" field.
func ThirdPartyIDGT(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDGTE applies the GTE predicate on the "third_party_id" field.
func ThirdPartyIDGTE(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDLT applies the LT predicate on the "third_party_id" field.
func ThirdPartyIDLT(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDLTE applies the LTE predicate on the "third_party_id" field.
func ThirdPartyIDLTE(v uuid.UUID) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDIsNil applies the IsNil predicate on the "third_party_id" field.
func ThirdPartyIDIsNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThirdPartyID)))
	})
}

// ThirdPartyIDNotNil applies the NotNil predicate on the "third_party_id" field.
func ThirdPartyIDNotNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThirdPartyID)))
	})
}

// ClientIDEQ applies the EQ predicate on the "client_id" field.
func ClientIDEQ(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientID), v))
	})
}

// ClientIDNEQ applies the NEQ predicate on the "client_id" field.
func ClientIDNEQ(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientID), v))
	})
}

// ClientIDIn applies the In predicate on the "client_id" field.
func ClientIDIn(vs ...string) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientID), v...))
	})
}

// ClientIDNotIn applies the NotIn predicate on the "client_id" field.
func ClientIDNotIn(vs ...string) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientID), v...))
	})
}

// ClientIDGT applies the GT predicate on the "client_id" field.
func ClientIDGT(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientID), v))
	})
}

// ClientIDGTE applies the GTE predicate on the "client_id" field.
func ClientIDGTE(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientID), v))
	})
}

// ClientIDLT applies the LT predicate on the "client_id" field.
func ClientIDLT(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientID), v))
	})
}

// ClientIDLTE applies the LTE predicate on the "client_id" field.
func ClientIDLTE(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientID), v))
	})
}

// ClientIDContains applies the Contains predicate on the "client_id" field.
func ClientIDContains(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientID), v))
	})
}

// ClientIDHasPrefix applies the HasPrefix predicate on the "client_id" field.
func ClientIDHasPrefix(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientID), v))
	})
}

// ClientIDHasSuffix applies the HasSuffix predicate on the "client_id" field.
func ClientIDHasSuffix(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientID), v))
	})
}

// ClientIDIsNil applies the IsNil predicate on the "client_id" field.
func ClientIDIsNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldClientID)))
	})
}

// ClientIDNotNil applies the NotNil predicate on the "client_id" field.
func ClientIDNotNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldClientID)))
	})
}

// ClientIDEqualFold applies the EqualFold predicate on the "client_id" field.
func ClientIDEqualFold(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientID), v))
	})
}

// ClientIDContainsFold applies the ContainsFold predicate on the "client_id" field.
func ClientIDContainsFold(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientID), v))
	})
}

// ClientSecretEQ applies the EQ predicate on the "client_secret" field.
func ClientSecretEQ(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientSecret), v))
	})
}

// ClientSecretNEQ applies the NEQ predicate on the "client_secret" field.
func ClientSecretNEQ(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientSecret), v))
	})
}

// ClientSecretIn applies the In predicate on the "client_secret" field.
func ClientSecretIn(vs ...string) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientSecret), v...))
	})
}

// ClientSecretNotIn applies the NotIn predicate on the "client_secret" field.
func ClientSecretNotIn(vs ...string) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientSecret), v...))
	})
}

// ClientSecretGT applies the GT predicate on the "client_secret" field.
func ClientSecretGT(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientSecret), v))
	})
}

// ClientSecretGTE applies the GTE predicate on the "client_secret" field.
func ClientSecretGTE(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientSecret), v))
	})
}

// ClientSecretLT applies the LT predicate on the "client_secret" field.
func ClientSecretLT(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientSecret), v))
	})
}

// ClientSecretLTE applies the LTE predicate on the "client_secret" field.
func ClientSecretLTE(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientSecret), v))
	})
}

// ClientSecretContains applies the Contains predicate on the "client_secret" field.
func ClientSecretContains(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientSecret), v))
	})
}

// ClientSecretHasPrefix applies the HasPrefix predicate on the "client_secret" field.
func ClientSecretHasPrefix(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientSecret), v))
	})
}

// ClientSecretHasSuffix applies the HasSuffix predicate on the "client_secret" field.
func ClientSecretHasSuffix(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientSecret), v))
	})
}

// ClientSecretIsNil applies the IsNil predicate on the "client_secret" field.
func ClientSecretIsNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldClientSecret)))
	})
}

// ClientSecretNotNil applies the NotNil predicate on the "client_secret" field.
func ClientSecretNotNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldClientSecret)))
	})
}

// ClientSecretEqualFold applies the EqualFold predicate on the "client_secret" field.
func ClientSecretEqualFold(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientSecret), v))
	})
}

// ClientSecretContainsFold applies the ContainsFold predicate on the "client_secret" field.
func ClientSecretContainsFold(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientSecret), v))
	})
}

// CallbackURLEQ applies the EQ predicate on the "callback_url" field.
func CallbackURLEQ(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCallbackURL), v))
	})
}

// CallbackURLNEQ applies the NEQ predicate on the "callback_url" field.
func CallbackURLNEQ(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCallbackURL), v))
	})
}

// CallbackURLIn applies the In predicate on the "callback_url" field.
func CallbackURLIn(vs ...string) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCallbackURL), v...))
	})
}

// CallbackURLNotIn applies the NotIn predicate on the "callback_url" field.
func CallbackURLNotIn(vs ...string) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCallbackURL), v...))
	})
}

// CallbackURLGT applies the GT predicate on the "callback_url" field.
func CallbackURLGT(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCallbackURL), v))
	})
}

// CallbackURLGTE applies the GTE predicate on the "callback_url" field.
func CallbackURLGTE(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCallbackURL), v))
	})
}

// CallbackURLLT applies the LT predicate on the "callback_url" field.
func CallbackURLLT(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCallbackURL), v))
	})
}

// CallbackURLLTE applies the LTE predicate on the "callback_url" field.
func CallbackURLLTE(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCallbackURL), v))
	})
}

// CallbackURLContains applies the Contains predicate on the "callback_url" field.
func CallbackURLContains(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCallbackURL), v))
	})
}

// CallbackURLHasPrefix applies the HasPrefix predicate on the "callback_url" field.
func CallbackURLHasPrefix(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCallbackURL), v))
	})
}

// CallbackURLHasSuffix applies the HasSuffix predicate on the "callback_url" field.
func CallbackURLHasSuffix(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCallbackURL), v))
	})
}

// CallbackURLIsNil applies the IsNil predicate on the "callback_url" field.
func CallbackURLIsNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCallbackURL)))
	})
}

// CallbackURLNotNil applies the NotNil predicate on the "callback_url" field.
func CallbackURLNotNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCallbackURL)))
	})
}

// CallbackURLEqualFold applies the EqualFold predicate on the "callback_url" field.
func CallbackURLEqualFold(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCallbackURL), v))
	})
}

// CallbackURLContainsFold applies the ContainsFold predicate on the "callback_url" field.
func CallbackURLContainsFold(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCallbackURL), v))
	})
}

// SaltEQ applies the EQ predicate on the "salt" field.
func SaltEQ(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalt), v))
	})
}

// SaltNEQ applies the NEQ predicate on the "salt" field.
func SaltNEQ(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSalt), v))
	})
}

// SaltIn applies the In predicate on the "salt" field.
func SaltIn(vs ...string) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSalt), v...))
	})
}

// SaltNotIn applies the NotIn predicate on the "salt" field.
func SaltNotIn(vs ...string) predicate.AppOAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSalt), v...))
	})
}

// SaltGT applies the GT predicate on the "salt" field.
func SaltGT(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSalt), v))
	})
}

// SaltGTE applies the GTE predicate on the "salt" field.
func SaltGTE(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSalt), v))
	})
}

// SaltLT applies the LT predicate on the "salt" field.
func SaltLT(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSalt), v))
	})
}

// SaltLTE applies the LTE predicate on the "salt" field.
func SaltLTE(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSalt), v))
	})
}

// SaltContains applies the Contains predicate on the "salt" field.
func SaltContains(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSalt), v))
	})
}

// SaltHasPrefix applies the HasPrefix predicate on the "salt" field.
func SaltHasPrefix(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSalt), v))
	})
}

// SaltHasSuffix applies the HasSuffix predicate on the "salt" field.
func SaltHasSuffix(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSalt), v))
	})
}

// SaltIsNil applies the IsNil predicate on the "salt" field.
func SaltIsNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSalt)))
	})
}

// SaltNotNil applies the NotNil predicate on the "salt" field.
func SaltNotNil() predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSalt)))
	})
}

// SaltEqualFold applies the EqualFold predicate on the "salt" field.
func SaltEqualFold(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSalt), v))
	})
}

// SaltContainsFold applies the ContainsFold predicate on the "salt" field.
func SaltContainsFold(v string) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSalt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppOAuthThirdParty) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppOAuthThirdParty) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppOAuthThirdParty) predicate.AppOAuthThirdParty {
	return predicate.AppOAuthThirdParty(func(s *sql.Selector) {
		p(s.Not())
	})
}
