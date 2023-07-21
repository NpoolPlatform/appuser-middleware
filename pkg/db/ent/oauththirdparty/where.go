// Code generated by ent, DO NOT EDIT.

package oauththirdparty

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// ClientName applies equality check predicate on the "client_name" field. It's identical to ClientNameEQ.
func ClientName(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientName), v))
	})
}

// ClientTag applies equality check predicate on the "client_tag" field. It's identical to ClientTagEQ.
func ClientTag(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientTag), v))
	})
}

// ClientLogoURL applies equality check predicate on the "client_logo_url" field. It's identical to ClientLogoURLEQ.
func ClientLogoURL(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientLogoURL), v))
	})
}

// ClientOauthURL applies equality check predicate on the "client_oauth_url" field. It's identical to ClientOauthURLEQ.
func ClientOauthURL(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientOauthURL), v))
	})
}

// ResponseType applies equality check predicate on the "response_type" field. It's identical to ResponseTypeEQ.
func ResponseType(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResponseType), v))
	})
}

// Scope applies equality check predicate on the "scope" field. It's identical to ScopeEQ.
func Scope(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScope), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// ClientNameEQ applies the EQ predicate on the "client_name" field.
func ClientNameEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientName), v))
	})
}

// ClientNameNEQ applies the NEQ predicate on the "client_name" field.
func ClientNameNEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientName), v))
	})
}

// ClientNameIn applies the In predicate on the "client_name" field.
func ClientNameIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientName), v...))
	})
}

// ClientNameNotIn applies the NotIn predicate on the "client_name" field.
func ClientNameNotIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientName), v...))
	})
}

// ClientNameGT applies the GT predicate on the "client_name" field.
func ClientNameGT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientName), v))
	})
}

// ClientNameGTE applies the GTE predicate on the "client_name" field.
func ClientNameGTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientName), v))
	})
}

// ClientNameLT applies the LT predicate on the "client_name" field.
func ClientNameLT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientName), v))
	})
}

// ClientNameLTE applies the LTE predicate on the "client_name" field.
func ClientNameLTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientName), v))
	})
}

// ClientNameContains applies the Contains predicate on the "client_name" field.
func ClientNameContains(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientName), v))
	})
}

// ClientNameHasPrefix applies the HasPrefix predicate on the "client_name" field.
func ClientNameHasPrefix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientName), v))
	})
}

// ClientNameHasSuffix applies the HasSuffix predicate on the "client_name" field.
func ClientNameHasSuffix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientName), v))
	})
}

// ClientNameIsNil applies the IsNil predicate on the "client_name" field.
func ClientNameIsNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldClientName)))
	})
}

// ClientNameNotNil applies the NotNil predicate on the "client_name" field.
func ClientNameNotNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldClientName)))
	})
}

// ClientNameEqualFold applies the EqualFold predicate on the "client_name" field.
func ClientNameEqualFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientName), v))
	})
}

// ClientNameContainsFold applies the ContainsFold predicate on the "client_name" field.
func ClientNameContainsFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientName), v))
	})
}

// ClientTagEQ applies the EQ predicate on the "client_tag" field.
func ClientTagEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientTag), v))
	})
}

// ClientTagNEQ applies the NEQ predicate on the "client_tag" field.
func ClientTagNEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientTag), v))
	})
}

// ClientTagIn applies the In predicate on the "client_tag" field.
func ClientTagIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientTag), v...))
	})
}

// ClientTagNotIn applies the NotIn predicate on the "client_tag" field.
func ClientTagNotIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientTag), v...))
	})
}

// ClientTagGT applies the GT predicate on the "client_tag" field.
func ClientTagGT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientTag), v))
	})
}

// ClientTagGTE applies the GTE predicate on the "client_tag" field.
func ClientTagGTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientTag), v))
	})
}

// ClientTagLT applies the LT predicate on the "client_tag" field.
func ClientTagLT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientTag), v))
	})
}

// ClientTagLTE applies the LTE predicate on the "client_tag" field.
func ClientTagLTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientTag), v))
	})
}

// ClientTagContains applies the Contains predicate on the "client_tag" field.
func ClientTagContains(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientTag), v))
	})
}

// ClientTagHasPrefix applies the HasPrefix predicate on the "client_tag" field.
func ClientTagHasPrefix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientTag), v))
	})
}

// ClientTagHasSuffix applies the HasSuffix predicate on the "client_tag" field.
func ClientTagHasSuffix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientTag), v))
	})
}

// ClientTagIsNil applies the IsNil predicate on the "client_tag" field.
func ClientTagIsNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldClientTag)))
	})
}

// ClientTagNotNil applies the NotNil predicate on the "client_tag" field.
func ClientTagNotNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldClientTag)))
	})
}

// ClientTagEqualFold applies the EqualFold predicate on the "client_tag" field.
func ClientTagEqualFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientTag), v))
	})
}

// ClientTagContainsFold applies the ContainsFold predicate on the "client_tag" field.
func ClientTagContainsFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientTag), v))
	})
}

// ClientLogoURLEQ applies the EQ predicate on the "client_logo_url" field.
func ClientLogoURLEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientLogoURL), v))
	})
}

// ClientLogoURLNEQ applies the NEQ predicate on the "client_logo_url" field.
func ClientLogoURLNEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientLogoURL), v))
	})
}

// ClientLogoURLIn applies the In predicate on the "client_logo_url" field.
func ClientLogoURLIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientLogoURL), v...))
	})
}

// ClientLogoURLNotIn applies the NotIn predicate on the "client_logo_url" field.
func ClientLogoURLNotIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientLogoURL), v...))
	})
}

// ClientLogoURLGT applies the GT predicate on the "client_logo_url" field.
func ClientLogoURLGT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientLogoURL), v))
	})
}

// ClientLogoURLGTE applies the GTE predicate on the "client_logo_url" field.
func ClientLogoURLGTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientLogoURL), v))
	})
}

// ClientLogoURLLT applies the LT predicate on the "client_logo_url" field.
func ClientLogoURLLT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientLogoURL), v))
	})
}

// ClientLogoURLLTE applies the LTE predicate on the "client_logo_url" field.
func ClientLogoURLLTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientLogoURL), v))
	})
}

// ClientLogoURLContains applies the Contains predicate on the "client_logo_url" field.
func ClientLogoURLContains(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientLogoURL), v))
	})
}

// ClientLogoURLHasPrefix applies the HasPrefix predicate on the "client_logo_url" field.
func ClientLogoURLHasPrefix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientLogoURL), v))
	})
}

// ClientLogoURLHasSuffix applies the HasSuffix predicate on the "client_logo_url" field.
func ClientLogoURLHasSuffix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientLogoURL), v))
	})
}

// ClientLogoURLIsNil applies the IsNil predicate on the "client_logo_url" field.
func ClientLogoURLIsNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldClientLogoURL)))
	})
}

// ClientLogoURLNotNil applies the NotNil predicate on the "client_logo_url" field.
func ClientLogoURLNotNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldClientLogoURL)))
	})
}

// ClientLogoURLEqualFold applies the EqualFold predicate on the "client_logo_url" field.
func ClientLogoURLEqualFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientLogoURL), v))
	})
}

// ClientLogoURLContainsFold applies the ContainsFold predicate on the "client_logo_url" field.
func ClientLogoURLContainsFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientLogoURL), v))
	})
}

// ClientOauthURLEQ applies the EQ predicate on the "client_oauth_url" field.
func ClientOauthURLEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientOauthURL), v))
	})
}

// ClientOauthURLNEQ applies the NEQ predicate on the "client_oauth_url" field.
func ClientOauthURLNEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientOauthURL), v))
	})
}

// ClientOauthURLIn applies the In predicate on the "client_oauth_url" field.
func ClientOauthURLIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientOauthURL), v...))
	})
}

// ClientOauthURLNotIn applies the NotIn predicate on the "client_oauth_url" field.
func ClientOauthURLNotIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientOauthURL), v...))
	})
}

// ClientOauthURLGT applies the GT predicate on the "client_oauth_url" field.
func ClientOauthURLGT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientOauthURL), v))
	})
}

// ClientOauthURLGTE applies the GTE predicate on the "client_oauth_url" field.
func ClientOauthURLGTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientOauthURL), v))
	})
}

// ClientOauthURLLT applies the LT predicate on the "client_oauth_url" field.
func ClientOauthURLLT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientOauthURL), v))
	})
}

// ClientOauthURLLTE applies the LTE predicate on the "client_oauth_url" field.
func ClientOauthURLLTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientOauthURL), v))
	})
}

// ClientOauthURLContains applies the Contains predicate on the "client_oauth_url" field.
func ClientOauthURLContains(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientOauthURL), v))
	})
}

// ClientOauthURLHasPrefix applies the HasPrefix predicate on the "client_oauth_url" field.
func ClientOauthURLHasPrefix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientOauthURL), v))
	})
}

// ClientOauthURLHasSuffix applies the HasSuffix predicate on the "client_oauth_url" field.
func ClientOauthURLHasSuffix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientOauthURL), v))
	})
}

// ClientOauthURLIsNil applies the IsNil predicate on the "client_oauth_url" field.
func ClientOauthURLIsNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldClientOauthURL)))
	})
}

// ClientOauthURLNotNil applies the NotNil predicate on the "client_oauth_url" field.
func ClientOauthURLNotNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldClientOauthURL)))
	})
}

// ClientOauthURLEqualFold applies the EqualFold predicate on the "client_oauth_url" field.
func ClientOauthURLEqualFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientOauthURL), v))
	})
}

// ClientOauthURLContainsFold applies the ContainsFold predicate on the "client_oauth_url" field.
func ClientOauthURLContainsFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientOauthURL), v))
	})
}

// ResponseTypeEQ applies the EQ predicate on the "response_type" field.
func ResponseTypeEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResponseType), v))
	})
}

// ResponseTypeNEQ applies the NEQ predicate on the "response_type" field.
func ResponseTypeNEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResponseType), v))
	})
}

// ResponseTypeIn applies the In predicate on the "response_type" field.
func ResponseTypeIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldResponseType), v...))
	})
}

// ResponseTypeNotIn applies the NotIn predicate on the "response_type" field.
func ResponseTypeNotIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldResponseType), v...))
	})
}

// ResponseTypeGT applies the GT predicate on the "response_type" field.
func ResponseTypeGT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResponseType), v))
	})
}

// ResponseTypeGTE applies the GTE predicate on the "response_type" field.
func ResponseTypeGTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResponseType), v))
	})
}

// ResponseTypeLT applies the LT predicate on the "response_type" field.
func ResponseTypeLT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResponseType), v))
	})
}

// ResponseTypeLTE applies the LTE predicate on the "response_type" field.
func ResponseTypeLTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResponseType), v))
	})
}

// ResponseTypeContains applies the Contains predicate on the "response_type" field.
func ResponseTypeContains(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResponseType), v))
	})
}

// ResponseTypeHasPrefix applies the HasPrefix predicate on the "response_type" field.
func ResponseTypeHasPrefix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResponseType), v))
	})
}

// ResponseTypeHasSuffix applies the HasSuffix predicate on the "response_type" field.
func ResponseTypeHasSuffix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResponseType), v))
	})
}

// ResponseTypeIsNil applies the IsNil predicate on the "response_type" field.
func ResponseTypeIsNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldResponseType)))
	})
}

// ResponseTypeNotNil applies the NotNil predicate on the "response_type" field.
func ResponseTypeNotNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldResponseType)))
	})
}

// ResponseTypeEqualFold applies the EqualFold predicate on the "response_type" field.
func ResponseTypeEqualFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResponseType), v))
	})
}

// ResponseTypeContainsFold applies the ContainsFold predicate on the "response_type" field.
func ResponseTypeContainsFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResponseType), v))
	})
}

// ScopeEQ applies the EQ predicate on the "scope" field.
func ScopeEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScope), v))
	})
}

// ScopeNEQ applies the NEQ predicate on the "scope" field.
func ScopeNEQ(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScope), v))
	})
}

// ScopeIn applies the In predicate on the "scope" field.
func ScopeIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldScope), v...))
	})
}

// ScopeNotIn applies the NotIn predicate on the "scope" field.
func ScopeNotIn(vs ...string) predicate.OAuthThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldScope), v...))
	})
}

// ScopeGT applies the GT predicate on the "scope" field.
func ScopeGT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScope), v))
	})
}

// ScopeGTE applies the GTE predicate on the "scope" field.
func ScopeGTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScope), v))
	})
}

// ScopeLT applies the LT predicate on the "scope" field.
func ScopeLT(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScope), v))
	})
}

// ScopeLTE applies the LTE predicate on the "scope" field.
func ScopeLTE(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScope), v))
	})
}

// ScopeContains applies the Contains predicate on the "scope" field.
func ScopeContains(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldScope), v))
	})
}

// ScopeHasPrefix applies the HasPrefix predicate on the "scope" field.
func ScopeHasPrefix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldScope), v))
	})
}

// ScopeHasSuffix applies the HasSuffix predicate on the "scope" field.
func ScopeHasSuffix(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldScope), v))
	})
}

// ScopeIsNil applies the IsNil predicate on the "scope" field.
func ScopeIsNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldScope)))
	})
}

// ScopeNotNil applies the NotNil predicate on the "scope" field.
func ScopeNotNil() predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldScope)))
	})
}

// ScopeEqualFold applies the EqualFold predicate on the "scope" field.
func ScopeEqualFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldScope), v))
	})
}

// ScopeContainsFold applies the ContainsFold predicate on the "scope" field.
func ScopeContainsFold(v string) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldScope), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OAuthThirdParty) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OAuthThirdParty) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
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
func Not(p predicate.OAuthThirdParty) predicate.OAuthThirdParty {
	return predicate.OAuthThirdParty(func(s *sql.Selector) {
		p(s.Not())
	})
}
