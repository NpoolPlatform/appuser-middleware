package appoauththirdparty

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/aes"
	appoauththirdpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/oauth/appoauththirdparty"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entappoauththirdparty "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appoauththirdparty"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/appoauththirdparty"
)

func (h *Handler) UpdateOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	info, err := h.GetOAuthThirdParty(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		oauthThirdParty, err := tx.AppOAuthThirdParty.
			Query().
			Where(
				entappoauththirdparty.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if oauthThirdParty == nil {
			return fmt.Errorf("invalid oauththirdparty")
		}
		if h.ClientSecret != nil && *h.ClientSecret != "" {
			salt, err := aes.NewAesKey(aes.AES256)
			if err != nil {
				return fmt.Errorf("get salt failed")
			}
			h.Salt = &salt
			clientSecret, err := aes.AesEncrypt([]byte(salt), []byte(*h.ClientSecret))
			if err != nil {
				return fmt.Errorf("encrypt clientSecret failed")
			}
			clientSecretStr := string(clientSecret)
			h.ClientSecret = &clientSecretStr
		}

		if _, err := appoauththirdpartycrud.UpdateSet(
			oauthThirdParty.Update(),
			&appoauththirdpartycrud.Req{
				ClientID:     h.ClientID,
				ClientSecret: h.ClientSecret,
				CallbackURL:  h.CallbackURL,
				Salt:         h.Salt,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}