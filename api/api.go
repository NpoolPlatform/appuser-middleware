package api

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/api/app"
	"github.com/NpoolPlatform/appuser-middleware/api/authing/auth"
	"github.com/NpoolPlatform/appuser-middleware/api/authing/history"
	"github.com/NpoolPlatform/appuser-middleware/api/kyc"
	"github.com/NpoolPlatform/appuser-middleware/api/role"
	roleuser "github.com/NpoolPlatform/appuser-middleware/api/role/user"
	"github.com/NpoolPlatform/appuser-middleware/api/subscriber"
	"github.com/NpoolPlatform/appuser-middleware/api/user"

	appusermw "github.com/NpoolPlatform/message/npool/appuser/mw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Service struct {
	appusermw.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appusermw.RegisterMiddlewareServer(server, &Service{})
	app.Register(server)
	subscriber.Register(server)
	user.Register(server)
	role.Register(server)
	roleuser.Register(server)
	auth.Register(server)
	history.Register(server)
	kyc.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := appusermw.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
