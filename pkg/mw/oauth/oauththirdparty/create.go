package oauththirdparty

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	oauththirdpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/oauth/oauththirdparty"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/oauth/oauththirdparty"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

func (h *Handler) CreateOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	key := fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixCreateOAuthThirdParty, *h.ClientName)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	const limit = 2
	_handler, err := NewHandler(
		ctx,
		WithConds(&npool.Conds{
			ClientName: &basetypes.Int32Val{Op: cruder.EQ, Value: int32(*h.ClientName)},
		}),
		WithLimit(limit),
	)
	if err != nil {
		return nil, err
	}
	infos, _, err := _handler.GetOAuthThirdParties(ctx)
	if err != nil {
		return nil, err
	}
	if infos != nil || len(infos) > 0 {
		return infos[0], nil
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := oauththirdpartycrud.CreateSet(
			tx.OAuthThirdParty.Create(),
			&oauththirdpartycrud.Req{
				EntID:          h.EntID,
				ClientName:     h.ClientName,
				ClientTag:      h.ClientTag,
				ClientLogoURL:  h.ClientLogoURL,
				ClientOAuthURL: h.ClientOAuthURL,
				ResponseType:   h.ResponseType,
				Scope:          h.Scope,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOAuthThirdParty(ctx)
}
