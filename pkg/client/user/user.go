//nolint:nolintlint,dupl
package user

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.UserMwClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewUserMwClient(conn)

	return fn(_ctx, cli)
}

func CreateUser(ctx context.Context, in *npool.UserReq) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.UserMwClient) (cruder.Any, error) {
		resp, err := cli.CreateUser(ctx, &npool.CreateUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.User), nil
}

func UpdateUser(ctx context.Context, in *npool.UserReq) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.UserMwClient) (cruder.Any, error) {
		resp, err := cli.UpdateUser(ctx, &npool.UpdateUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.User), nil
}

func GetUser(ctx context.Context, appID, userID string) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.UserMwClient) (cruder.Any, error) {
		resp, err := cli.GetUser(ctx, &npool.GetUserRequest{
			AppID:  appID,
			UserID: userID,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.User), nil
}

func GetUsers(ctx context.Context, appID string, offset, limit int32) ([]*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.UserMwClient) (cruder.Any, error) {
		resp, err := cli.GetUsers(ctx, &npool.GetUsersRequest{
			AppID:  appID,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.([]*npool.User), nil
}

func GetManyUsers(ctx context.Context, ids []string) ([]*npool.User, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.UserMwClient) (cruder.Any, error) {
		resp, err := cli.GetManyUsers(ctx, &npool.GetManyUsersRequest{
			IDs: ids,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.User), nil
}
