package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
)

// Kyc holds the schema definition for the Kyc entity.
type Kyc struct {
	ent.Schema
}

func (Kyc) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Kyc.
func (Kyc) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("document_type").
			Optional().
			Default(basetypes.KycDocumentType_DefaultKycDocumentType.String()),
		field.
			String("id_number").
			Optional().
			Default(""),
		field.
			String("front_img").
			Optional().
			Default(""),
		field.
			String("back_img").
			Optional().
			Default(""),
		field.
			String("selfie_img").
			Optional().
			Default(""),
		field.
			String("entity_type").
			Optional().
			Default(basetypes.KycEntityType_Individual.String()),
		field.
			UUID("review_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("state").
			Optional().
			Default(basetypes.KycState_DefaultState.String()),
	}
}

// Edges of the Kyc.
func (Kyc) Edges() []ent.Edge {
	return nil
}
