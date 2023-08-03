//nolint:dupl
package oauththirdparty

import (
	"context"

	oauththirdparty1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/oauth/oauththirdparty"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/oauththirdparty"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistOAuthThirdParty(ctx context.Context, in *npool.ExistOAuthThirdPartyRequest) (*npool.ExistOAuthThirdPartyResponse, error) {
	handler, err := oauththirdparty1.NewHandler(
		ctx,
		oauththirdparty1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOAuthThirdParty",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOAuthThirdPartyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	exist, err := handler.ExistOAuthThirdParty(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOAuthThirdParty",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOAuthThirdPartyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistOAuthThirdPartyResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistOAuthThirdPartyConds(ctx context.Context, in *npool.ExistOAuthThirdPartyCondsRequest) (*npool.ExistOAuthThirdPartyCondsResponse, error) {
	handler, err := oauththirdparty1.NewHandler(
		ctx,
		oauththirdparty1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOAuthThirdPartyConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOAuthThirdPartyCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	exist, err := handler.ExistOAuthThirdPartyConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOAuthThirdPartyConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOAuthThirdPartyCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistOAuthThirdPartyCondsResponse{
		Info: exist,
	}, nil
}
