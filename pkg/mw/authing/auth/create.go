package auth

import (
	"context"

	authcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/auth"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateAuth(ctx context.Context) (*npool.Auth, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := authcrud.CreateSet(
			cli.Auth.Create(),
			&authcrud.Req{
				ID:       h.ID,
				AppID:    &h.AppID,
				RoleID:   h.RoleID,
				UserID:   h.UserID,
				Resource: h.Resource,
				Method:   h.Method,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAuth(ctx)
}

func (h *Handler) CreateAuths(ctx context.Context) ([]*npool.Auth, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.ID != nil {
				id = uuid.MustParse(*req.ID)
			}

			appID := uuid.MustParse(*req.AppID)
			roleID := uuid.MustParse(*req.RoleID)
			userID := uuid.MustParse(*req.UserID)

			if _, err := authcrud.CreateSet(
				tx.Auth.Create(),
				&authcrud.Req{
					ID:       &id,
					AppID:    &appID,
					RoleID:   &roleID,
					UserID:   &userID,
					Resource: req.Resource,
					Method:   req.Method,
				},
			).Save(_ctx); err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &authcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))
	infos, _, err := h.GetAuths(ctx)
	if err != nil {
		return nil, err
	}
	return infos, nil
}
