// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AppQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppQueryRuleFunc func(context.Context, *ent.AppQuery) error

// EvalQuery return f(ctx, q).
func (f AppQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppQuery", q)
}

// The AppMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppMutationRuleFunc func(context.Context, *ent.AppMutation) error

// EvalMutation calls f(ctx, m).
func (f AppMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppMutation", m)
}

// The AppControlQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppControlQueryRuleFunc func(context.Context, *ent.AppControlQuery) error

// EvalQuery return f(ctx, q).
func (f AppControlQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppControlQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppControlQuery", q)
}

// The AppControlMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppControlMutationRuleFunc func(context.Context, *ent.AppControlMutation) error

// EvalMutation calls f(ctx, m).
func (f AppControlMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppControlMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppControlMutation", m)
}

// The AppOAuthThirdPartyQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppOAuthThirdPartyQueryRuleFunc func(context.Context, *ent.AppOAuthThirdPartyQuery) error

// EvalQuery return f(ctx, q).
func (f AppOAuthThirdPartyQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppOAuthThirdPartyQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppOAuthThirdPartyQuery", q)
}

// The AppOAuthThirdPartyMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppOAuthThirdPartyMutationRuleFunc func(context.Context, *ent.AppOAuthThirdPartyMutation) error

// EvalMutation calls f(ctx, m).
func (f AppOAuthThirdPartyMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppOAuthThirdPartyMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppOAuthThirdPartyMutation", m)
}

// The AppRoleQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppRoleQueryRuleFunc func(context.Context, *ent.AppRoleQuery) error

// EvalQuery return f(ctx, q).
func (f AppRoleQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppRoleQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppRoleQuery", q)
}

// The AppRoleMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppRoleMutationRuleFunc func(context.Context, *ent.AppRoleMutation) error

// EvalMutation calls f(ctx, m).
func (f AppRoleMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppRoleMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppRoleMutation", m)
}

// The AppRoleUserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppRoleUserQueryRuleFunc func(context.Context, *ent.AppRoleUserQuery) error

// EvalQuery return f(ctx, q).
func (f AppRoleUserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppRoleUserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppRoleUserQuery", q)
}

// The AppRoleUserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppRoleUserMutationRuleFunc func(context.Context, *ent.AppRoleUserMutation) error

// EvalMutation calls f(ctx, m).
func (f AppRoleUserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppRoleUserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppRoleUserMutation", m)
}

// The AppSubscribeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppSubscribeQueryRuleFunc func(context.Context, *ent.AppSubscribeQuery) error

// EvalQuery return f(ctx, q).
func (f AppSubscribeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppSubscribeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppSubscribeQuery", q)
}

// The AppSubscribeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppSubscribeMutationRuleFunc func(context.Context, *ent.AppSubscribeMutation) error

// EvalMutation calls f(ctx, m).
func (f AppSubscribeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppSubscribeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppSubscribeMutation", m)
}

// The AppUserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppUserQueryRuleFunc func(context.Context, *ent.AppUserQuery) error

// EvalQuery return f(ctx, q).
func (f AppUserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppUserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppUserQuery", q)
}

// The AppUserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppUserMutationRuleFunc func(context.Context, *ent.AppUserMutation) error

// EvalMutation calls f(ctx, m).
func (f AppUserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppUserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppUserMutation", m)
}

// The AppUserControlQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppUserControlQueryRuleFunc func(context.Context, *ent.AppUserControlQuery) error

// EvalQuery return f(ctx, q).
func (f AppUserControlQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppUserControlQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppUserControlQuery", q)
}

// The AppUserControlMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppUserControlMutationRuleFunc func(context.Context, *ent.AppUserControlMutation) error

// EvalMutation calls f(ctx, m).
func (f AppUserControlMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppUserControlMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppUserControlMutation", m)
}

// The AppUserExtraQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppUserExtraQueryRuleFunc func(context.Context, *ent.AppUserExtraQuery) error

// EvalQuery return f(ctx, q).
func (f AppUserExtraQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppUserExtraQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppUserExtraQuery", q)
}

// The AppUserExtraMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppUserExtraMutationRuleFunc func(context.Context, *ent.AppUserExtraMutation) error

// EvalMutation calls f(ctx, m).
func (f AppUserExtraMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppUserExtraMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppUserExtraMutation", m)
}

// The AppUserSecretQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppUserSecretQueryRuleFunc func(context.Context, *ent.AppUserSecretQuery) error

// EvalQuery return f(ctx, q).
func (f AppUserSecretQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppUserSecretQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppUserSecretQuery", q)
}

// The AppUserSecretMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppUserSecretMutationRuleFunc func(context.Context, *ent.AppUserSecretMutation) error

// EvalMutation calls f(ctx, m).
func (f AppUserSecretMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppUserSecretMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppUserSecretMutation", m)
}

// The AppUserThirdPartyQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppUserThirdPartyQueryRuleFunc func(context.Context, *ent.AppUserThirdPartyQuery) error

// EvalQuery return f(ctx, q).
func (f AppUserThirdPartyQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppUserThirdPartyQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppUserThirdPartyQuery", q)
}

// The AppUserThirdPartyMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppUserThirdPartyMutationRuleFunc func(context.Context, *ent.AppUserThirdPartyMutation) error

// EvalMutation calls f(ctx, m).
func (f AppUserThirdPartyMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppUserThirdPartyMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppUserThirdPartyMutation", m)
}

// The AuthQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AuthQueryRuleFunc func(context.Context, *ent.AuthQuery) error

// EvalQuery return f(ctx, q).
func (f AuthQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AuthQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AuthQuery", q)
}

// The AuthMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AuthMutationRuleFunc func(context.Context, *ent.AuthMutation) error

// EvalMutation calls f(ctx, m).
func (f AuthMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AuthMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AuthMutation", m)
}

// The AuthHistoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AuthHistoryQueryRuleFunc func(context.Context, *ent.AuthHistoryQuery) error

// EvalQuery return f(ctx, q).
func (f AuthHistoryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AuthHistoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AuthHistoryQuery", q)
}

// The AuthHistoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AuthHistoryMutationRuleFunc func(context.Context, *ent.AuthHistoryMutation) error

// EvalMutation calls f(ctx, m).
func (f AuthHistoryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AuthHistoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AuthHistoryMutation", m)
}

// The BanAppQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BanAppQueryRuleFunc func(context.Context, *ent.BanAppQuery) error

// EvalQuery return f(ctx, q).
func (f BanAppQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BanAppQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BanAppQuery", q)
}

// The BanAppMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BanAppMutationRuleFunc func(context.Context, *ent.BanAppMutation) error

// EvalMutation calls f(ctx, m).
func (f BanAppMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BanAppMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BanAppMutation", m)
}

// The BanAppUserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BanAppUserQueryRuleFunc func(context.Context, *ent.BanAppUserQuery) error

// EvalQuery return f(ctx, q).
func (f BanAppUserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BanAppUserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BanAppUserQuery", q)
}

// The BanAppUserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BanAppUserMutationRuleFunc func(context.Context, *ent.BanAppUserMutation) error

// EvalMutation calls f(ctx, m).
func (f BanAppUserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BanAppUserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BanAppUserMutation", m)
}

// The KycQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type KycQueryRuleFunc func(context.Context, *ent.KycQuery) error

// EvalQuery return f(ctx, q).
func (f KycQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.KycQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.KycQuery", q)
}

// The KycMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type KycMutationRuleFunc func(context.Context, *ent.KycMutation) error

// EvalMutation calls f(ctx, m).
func (f KycMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.KycMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.KycMutation", m)
}

// The LoginHistoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LoginHistoryQueryRuleFunc func(context.Context, *ent.LoginHistoryQuery) error

// EvalQuery return f(ctx, q).
func (f LoginHistoryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LoginHistoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LoginHistoryQuery", q)
}

// The LoginHistoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LoginHistoryMutationRuleFunc func(context.Context, *ent.LoginHistoryMutation) error

// EvalMutation calls f(ctx, m).
func (f LoginHistoryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LoginHistoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LoginHistoryMutation", m)
}

// The OAuthThirdPartyQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OAuthThirdPartyQueryRuleFunc func(context.Context, *ent.OAuthThirdPartyQuery) error

// EvalQuery return f(ctx, q).
func (f OAuthThirdPartyQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OAuthThirdPartyQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.OAuthThirdPartyQuery", q)
}

// The OAuthThirdPartyMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OAuthThirdPartyMutationRuleFunc func(context.Context, *ent.OAuthThirdPartyMutation) error

// EvalMutation calls f(ctx, m).
func (f OAuthThirdPartyMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.OAuthThirdPartyMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.OAuthThirdPartyMutation", m)
}

// The PubsubMessageQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PubsubMessageQueryRuleFunc func(context.Context, *ent.PubsubMessageQuery) error

// EvalQuery return f(ctx, q).
func (f PubsubMessageQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PubsubMessageQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PubsubMessageQuery", q)
}

// The PubsubMessageMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PubsubMessageMutationRuleFunc func(context.Context, *ent.PubsubMessageMutation) error

// EvalMutation calls f(ctx, m).
func (f PubsubMessageMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PubsubMessageMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PubsubMessageMutation", m)
}

// The RecoveryCodeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RecoveryCodeQueryRuleFunc func(context.Context, *ent.RecoveryCodeQuery) error

// EvalQuery return f(ctx, q).
func (f RecoveryCodeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RecoveryCodeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RecoveryCodeQuery", q)
}

// The RecoveryCodeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RecoveryCodeMutationRuleFunc func(context.Context, *ent.RecoveryCodeMutation) error

// EvalMutation calls f(ctx, m).
func (f RecoveryCodeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RecoveryCodeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RecoveryCodeMutation", m)
}

// The SubscriberQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SubscriberQueryRuleFunc func(context.Context, *ent.SubscriberQuery) error

// EvalQuery return f(ctx, q).
func (f SubscriberQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SubscriberQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SubscriberQuery", q)
}

// The SubscriberMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SubscriberMutationRuleFunc func(context.Context, *ent.SubscriberMutation) error

// EvalMutation calls f(ctx, m).
func (f SubscriberMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SubscriberMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SubscriberMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.AppQuery:
		return q.Filter(), nil
	case *ent.AppControlQuery:
		return q.Filter(), nil
	case *ent.AppOAuthThirdPartyQuery:
		return q.Filter(), nil
	case *ent.AppRoleQuery:
		return q.Filter(), nil
	case *ent.AppRoleUserQuery:
		return q.Filter(), nil
	case *ent.AppSubscribeQuery:
		return q.Filter(), nil
	case *ent.AppUserQuery:
		return q.Filter(), nil
	case *ent.AppUserControlQuery:
		return q.Filter(), nil
	case *ent.AppUserExtraQuery:
		return q.Filter(), nil
	case *ent.AppUserSecretQuery:
		return q.Filter(), nil
	case *ent.AppUserThirdPartyQuery:
		return q.Filter(), nil
	case *ent.AuthQuery:
		return q.Filter(), nil
	case *ent.AuthHistoryQuery:
		return q.Filter(), nil
	case *ent.BanAppQuery:
		return q.Filter(), nil
	case *ent.BanAppUserQuery:
		return q.Filter(), nil
	case *ent.KycQuery:
		return q.Filter(), nil
	case *ent.LoginHistoryQuery:
		return q.Filter(), nil
	case *ent.OAuthThirdPartyQuery:
		return q.Filter(), nil
	case *ent.PubsubMessageQuery:
		return q.Filter(), nil
	case *ent.RecoveryCodeQuery:
		return q.Filter(), nil
	case *ent.SubscriberQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.AppMutation:
		return m.Filter(), nil
	case *ent.AppControlMutation:
		return m.Filter(), nil
	case *ent.AppOAuthThirdPartyMutation:
		return m.Filter(), nil
	case *ent.AppRoleMutation:
		return m.Filter(), nil
	case *ent.AppRoleUserMutation:
		return m.Filter(), nil
	case *ent.AppSubscribeMutation:
		return m.Filter(), nil
	case *ent.AppUserMutation:
		return m.Filter(), nil
	case *ent.AppUserControlMutation:
		return m.Filter(), nil
	case *ent.AppUserExtraMutation:
		return m.Filter(), nil
	case *ent.AppUserSecretMutation:
		return m.Filter(), nil
	case *ent.AppUserThirdPartyMutation:
		return m.Filter(), nil
	case *ent.AuthMutation:
		return m.Filter(), nil
	case *ent.AuthHistoryMutation:
		return m.Filter(), nil
	case *ent.BanAppMutation:
		return m.Filter(), nil
	case *ent.BanAppUserMutation:
		return m.Filter(), nil
	case *ent.KycMutation:
		return m.Filter(), nil
	case *ent.LoginHistoryMutation:
		return m.Filter(), nil
	case *ent.OAuthThirdPartyMutation:
		return m.Filter(), nil
	case *ent.PubsubMessageMutation:
		return m.Filter(), nil
	case *ent.RecoveryCodeMutation:
		return m.Filter(), nil
	case *ent.SubscriberMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
