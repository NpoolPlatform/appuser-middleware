// Code generated by ent, DO NOT EDIT.

package appcontrol

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// RecaptchaMethod applies equality check predicate on the "recaptcha_method" field. It's identical to RecaptchaMethodEQ.
func RecaptchaMethod(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecaptchaMethod), v))
	})
}

// KycEnable applies equality check predicate on the "kyc_enable" field. It's identical to KycEnableEQ.
func KycEnable(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKycEnable), v))
	})
}

// SigninVerifyEnable applies equality check predicate on the "signin_verify_enable" field. It's identical to SigninVerifyEnableEQ.
func SigninVerifyEnable(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSigninVerifyEnable), v))
	})
}

// InvitationCodeMust applies equality check predicate on the "invitation_code_must" field. It's identical to InvitationCodeMustEQ.
func InvitationCodeMust(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvitationCodeMust), v))
	})
}

// CreateInvitationCodeWhen applies equality check predicate on the "create_invitation_code_when" field. It's identical to CreateInvitationCodeWhenEQ.
func CreateInvitationCodeWhen(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// MaxTypedCouponsPerOrder applies equality check predicate on the "max_typed_coupons_per_order" field. It's identical to MaxTypedCouponsPerOrderEQ.
func MaxTypedCouponsPerOrder(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// Maintaining applies equality check predicate on the "maintaining" field. It's identical to MaintainingEQ.
func Maintaining(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaintaining), v))
	})
}

// CouponWithdrawEnable applies equality check predicate on the "coupon_withdraw_enable" field. It's identical to CouponWithdrawEnableEQ.
func CouponWithdrawEnable(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCouponWithdrawEnable), v))
	})
}

// ResetUserMethod applies equality check predicate on the "reset_user_method" field. It's identical to ResetUserMethodEQ.
func ResetUserMethod(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResetUserMethod), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// SignupMethodsIsNil applies the IsNil predicate on the "signup_methods" field.
func SignupMethodsIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSignupMethods)))
	})
}

// SignupMethodsNotNil applies the NotNil predicate on the "signup_methods" field.
func SignupMethodsNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSignupMethods)))
	})
}

// ExternSigninMethodsIsNil applies the IsNil predicate on the "extern_signin_methods" field.
func ExternSigninMethodsIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExternSigninMethods)))
	})
}

// ExternSigninMethodsNotNil applies the NotNil predicate on the "extern_signin_methods" field.
func ExternSigninMethodsNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExternSigninMethods)))
	})
}

// RecaptchaMethodEQ applies the EQ predicate on the "recaptcha_method" field.
func RecaptchaMethodEQ(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecaptchaMethod), v))
	})
}

// RecaptchaMethodNEQ applies the NEQ predicate on the "recaptcha_method" field.
func RecaptchaMethodNEQ(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecaptchaMethod), v))
	})
}

// RecaptchaMethodIn applies the In predicate on the "recaptcha_method" field.
func RecaptchaMethodIn(vs ...string) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRecaptchaMethod), v...))
	})
}

// RecaptchaMethodNotIn applies the NotIn predicate on the "recaptcha_method" field.
func RecaptchaMethodNotIn(vs ...string) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRecaptchaMethod), v...))
	})
}

// RecaptchaMethodGT applies the GT predicate on the "recaptcha_method" field.
func RecaptchaMethodGT(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecaptchaMethod), v))
	})
}

// RecaptchaMethodGTE applies the GTE predicate on the "recaptcha_method" field.
func RecaptchaMethodGTE(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecaptchaMethod), v))
	})
}

// RecaptchaMethodLT applies the LT predicate on the "recaptcha_method" field.
func RecaptchaMethodLT(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecaptchaMethod), v))
	})
}

// RecaptchaMethodLTE applies the LTE predicate on the "recaptcha_method" field.
func RecaptchaMethodLTE(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecaptchaMethod), v))
	})
}

// RecaptchaMethodContains applies the Contains predicate on the "recaptcha_method" field.
func RecaptchaMethodContains(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRecaptchaMethod), v))
	})
}

// RecaptchaMethodHasPrefix applies the HasPrefix predicate on the "recaptcha_method" field.
func RecaptchaMethodHasPrefix(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRecaptchaMethod), v))
	})
}

// RecaptchaMethodHasSuffix applies the HasSuffix predicate on the "recaptcha_method" field.
func RecaptchaMethodHasSuffix(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRecaptchaMethod), v))
	})
}

// RecaptchaMethodIsNil applies the IsNil predicate on the "recaptcha_method" field.
func RecaptchaMethodIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRecaptchaMethod)))
	})
}

// RecaptchaMethodNotNil applies the NotNil predicate on the "recaptcha_method" field.
func RecaptchaMethodNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRecaptchaMethod)))
	})
}

// RecaptchaMethodEqualFold applies the EqualFold predicate on the "recaptcha_method" field.
func RecaptchaMethodEqualFold(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRecaptchaMethod), v))
	})
}

// RecaptchaMethodContainsFold applies the ContainsFold predicate on the "recaptcha_method" field.
func RecaptchaMethodContainsFold(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRecaptchaMethod), v))
	})
}

// KycEnableEQ applies the EQ predicate on the "kyc_enable" field.
func KycEnableEQ(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKycEnable), v))
	})
}

// KycEnableNEQ applies the NEQ predicate on the "kyc_enable" field.
func KycEnableNEQ(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKycEnable), v))
	})
}

// KycEnableIsNil applies the IsNil predicate on the "kyc_enable" field.
func KycEnableIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldKycEnable)))
	})
}

// KycEnableNotNil applies the NotNil predicate on the "kyc_enable" field.
func KycEnableNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldKycEnable)))
	})
}

// SigninVerifyEnableEQ applies the EQ predicate on the "signin_verify_enable" field.
func SigninVerifyEnableEQ(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSigninVerifyEnable), v))
	})
}

// SigninVerifyEnableNEQ applies the NEQ predicate on the "signin_verify_enable" field.
func SigninVerifyEnableNEQ(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSigninVerifyEnable), v))
	})
}

// SigninVerifyEnableIsNil applies the IsNil predicate on the "signin_verify_enable" field.
func SigninVerifyEnableIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSigninVerifyEnable)))
	})
}

// SigninVerifyEnableNotNil applies the NotNil predicate on the "signin_verify_enable" field.
func SigninVerifyEnableNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSigninVerifyEnable)))
	})
}

// InvitationCodeMustEQ applies the EQ predicate on the "invitation_code_must" field.
func InvitationCodeMustEQ(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvitationCodeMust), v))
	})
}

// InvitationCodeMustNEQ applies the NEQ predicate on the "invitation_code_must" field.
func InvitationCodeMustNEQ(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInvitationCodeMust), v))
	})
}

// InvitationCodeMustIsNil applies the IsNil predicate on the "invitation_code_must" field.
func InvitationCodeMustIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInvitationCodeMust)))
	})
}

// InvitationCodeMustNotNil applies the NotNil predicate on the "invitation_code_must" field.
func InvitationCodeMustNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInvitationCodeMust)))
	})
}

// CreateInvitationCodeWhenEQ applies the EQ predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenEQ(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// CreateInvitationCodeWhenNEQ applies the NEQ predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenNEQ(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// CreateInvitationCodeWhenIn applies the In predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenIn(vs ...string) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateInvitationCodeWhen), v...))
	})
}

// CreateInvitationCodeWhenNotIn applies the NotIn predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenNotIn(vs ...string) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateInvitationCodeWhen), v...))
	})
}

// CreateInvitationCodeWhenGT applies the GT predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenGT(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// CreateInvitationCodeWhenGTE applies the GTE predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenGTE(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// CreateInvitationCodeWhenLT applies the LT predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenLT(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// CreateInvitationCodeWhenLTE applies the LTE predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenLTE(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// CreateInvitationCodeWhenContains applies the Contains predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenContains(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// CreateInvitationCodeWhenHasPrefix applies the HasPrefix predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenHasPrefix(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// CreateInvitationCodeWhenHasSuffix applies the HasSuffix predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenHasSuffix(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// CreateInvitationCodeWhenIsNil applies the IsNil predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateInvitationCodeWhen)))
	})
}

// CreateInvitationCodeWhenNotNil applies the NotNil predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateInvitationCodeWhen)))
	})
}

// CreateInvitationCodeWhenEqualFold applies the EqualFold predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenEqualFold(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// CreateInvitationCodeWhenContainsFold applies the ContainsFold predicate on the "create_invitation_code_when" field.
func CreateInvitationCodeWhenContainsFold(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreateInvitationCodeWhen), v))
	})
}

// MaxTypedCouponsPerOrderEQ applies the EQ predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderEQ(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderNEQ applies the NEQ predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderNEQ(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderIn applies the In predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderIn(vs ...uint32) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMaxTypedCouponsPerOrder), v...))
	})
}

// MaxTypedCouponsPerOrderNotIn applies the NotIn predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderNotIn(vs ...uint32) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMaxTypedCouponsPerOrder), v...))
	})
}

// MaxTypedCouponsPerOrderGT applies the GT predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderGT(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderGTE applies the GTE predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderGTE(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderLT applies the LT predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderLT(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderLTE applies the LTE predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderLTE(v uint32) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderIsNil applies the IsNil predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMaxTypedCouponsPerOrder)))
	})
}

// MaxTypedCouponsPerOrderNotNil applies the NotNil predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMaxTypedCouponsPerOrder)))
	})
}

// MaintainingEQ applies the EQ predicate on the "maintaining" field.
func MaintainingEQ(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaintaining), v))
	})
}

// MaintainingNEQ applies the NEQ predicate on the "maintaining" field.
func MaintainingNEQ(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaintaining), v))
	})
}

// MaintainingIsNil applies the IsNil predicate on the "maintaining" field.
func MaintainingIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMaintaining)))
	})
}

// MaintainingNotNil applies the NotNil predicate on the "maintaining" field.
func MaintainingNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMaintaining)))
	})
}

// CouponWithdrawEnableEQ applies the EQ predicate on the "coupon_withdraw_enable" field.
func CouponWithdrawEnableEQ(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCouponWithdrawEnable), v))
	})
}

// CouponWithdrawEnableNEQ applies the NEQ predicate on the "coupon_withdraw_enable" field.
func CouponWithdrawEnableNEQ(v bool) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCouponWithdrawEnable), v))
	})
}

// CouponWithdrawEnableIsNil applies the IsNil predicate on the "coupon_withdraw_enable" field.
func CouponWithdrawEnableIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCouponWithdrawEnable)))
	})
}

// CouponWithdrawEnableNotNil applies the NotNil predicate on the "coupon_withdraw_enable" field.
func CouponWithdrawEnableNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCouponWithdrawEnable)))
	})
}

// CommitButtonTargetsIsNil applies the IsNil predicate on the "commit_button_targets" field.
func CommitButtonTargetsIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCommitButtonTargets)))
	})
}

// CommitButtonTargetsNotNil applies the NotNil predicate on the "commit_button_targets" field.
func CommitButtonTargetsNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCommitButtonTargets)))
	})
}

// ResetUserMethodEQ applies the EQ predicate on the "reset_user_method" field.
func ResetUserMethodEQ(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResetUserMethod), v))
	})
}

// ResetUserMethodNEQ applies the NEQ predicate on the "reset_user_method" field.
func ResetUserMethodNEQ(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResetUserMethod), v))
	})
}

// ResetUserMethodIn applies the In predicate on the "reset_user_method" field.
func ResetUserMethodIn(vs ...string) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldResetUserMethod), v...))
	})
}

// ResetUserMethodNotIn applies the NotIn predicate on the "reset_user_method" field.
func ResetUserMethodNotIn(vs ...string) predicate.AppControl {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldResetUserMethod), v...))
	})
}

// ResetUserMethodGT applies the GT predicate on the "reset_user_method" field.
func ResetUserMethodGT(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResetUserMethod), v))
	})
}

// ResetUserMethodGTE applies the GTE predicate on the "reset_user_method" field.
func ResetUserMethodGTE(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResetUserMethod), v))
	})
}

// ResetUserMethodLT applies the LT predicate on the "reset_user_method" field.
func ResetUserMethodLT(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResetUserMethod), v))
	})
}

// ResetUserMethodLTE applies the LTE predicate on the "reset_user_method" field.
func ResetUserMethodLTE(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResetUserMethod), v))
	})
}

// ResetUserMethodContains applies the Contains predicate on the "reset_user_method" field.
func ResetUserMethodContains(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResetUserMethod), v))
	})
}

// ResetUserMethodHasPrefix applies the HasPrefix predicate on the "reset_user_method" field.
func ResetUserMethodHasPrefix(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResetUserMethod), v))
	})
}

// ResetUserMethodHasSuffix applies the HasSuffix predicate on the "reset_user_method" field.
func ResetUserMethodHasSuffix(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResetUserMethod), v))
	})
}

// ResetUserMethodIsNil applies the IsNil predicate on the "reset_user_method" field.
func ResetUserMethodIsNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldResetUserMethod)))
	})
}

// ResetUserMethodNotNil applies the NotNil predicate on the "reset_user_method" field.
func ResetUserMethodNotNil() predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldResetUserMethod)))
	})
}

// ResetUserMethodEqualFold applies the EqualFold predicate on the "reset_user_method" field.
func ResetUserMethodEqualFold(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResetUserMethod), v))
	})
}

// ResetUserMethodContainsFold applies the ContainsFold predicate on the "reset_user_method" field.
func ResetUserMethodContainsFold(v string) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResetUserMethod), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppControl) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppControl) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
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
func Not(p predicate.AppControl) predicate.AppControl {
	return predicate.AppControl(func(s *sql.Selector) {
		p(s.Not())
	})
}
