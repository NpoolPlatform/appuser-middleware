package history

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	app1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	appusermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/login/history"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	app = appmwpb.App{
		EntID:     uuid.NewString(),
		Name:      uuid.NewString(),
		CreatedBy: uuid.NewString(),
	}

	user = appusermwpb.User{
		EntID:             uuid.NewString(),
		AppID:             app.EntID,
		EmailAddress:      "aaa@aaa.aaa",
		PhoneNO:           "+8613612203133",
		ImportedFromAppID: uuid.NewString(),
		Username:          "amwnrekadsf.are-",
	}

	passwordHash = "AAA"
	ret          = npool.History{
		EntID:        uuid.NewString(),
		AppID:        app.EntID,
		AppName:      app.Name,
		AppLogo:      app.Logo,
		UserID:       user.EntID,
		EmailAddress: user.EmailAddress,
		ClientIP:     "192.168.1.2",
		UserAgent:    uuid.NewString(),
		Location:     "Shanghai",
		LoginType:    basetypes.LoginType_FreshLogin,
		LoginTypeStr: basetypes.LoginType_FreshLogin.String(),
	}
)

func setupHistory(t *testing.T) func(*testing.T) {
	// app
	handler, err := app1.NewHandler(
		context.Background(),
		app1.WithEntID(&app.EntID, true),
		app1.WithCreatedBy(&app.CreatedBy, true),
		app1.WithName(&app.Name, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, handler)

	info, err := handler.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)

	handler, err = app1.NewHandler(
		context.Background(),
		app1.WithID(&info.ID, true),
	)
	assert.Nil(t, err)

	// user
	userHandler, err := user1.NewHandler(
		context.TODO(),
		user1.WithEntID(&user.EntID, true),
		user1.WithAppID(&user.AppID, true),
		user1.WithEmailAddress(&user.EmailAddress, true),
		user1.WithPasswordHash(&passwordHash, true),
	)

	assert.Nil(t, err)
	assert.NotNil(t, userHandler)

	_user, err := userHandler.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, _user)

	userHandler, err = user1.NewHandler(
		context.TODO(),
		user1.WithID(&_user.ID, true),
	)
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = handler.DeleteApp(context.Background())
		_, _ = userHandler.DeleteUser(context.Background())
	}
}

func createHistory(t *testing.T) {
	req := &npool.HistoryReq{
		EntID:     &ret.EntID,
		AppID:     &ret.AppID,
		UserID:    &ret.UserID,
		ClientIP:  &ret.ClientIP,
		UserAgent: &ret.UserAgent,
		Location:  &ret.Location,
		LoginType: &ret.LoginType,
	}

	info, err := CreateHistory(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.AppName = info.AppName
		ret.AppLogo = info.AppLogo
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func getHistory(t *testing.T) {
	info, err := GetHistory(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getHistories(t *testing.T) {
	infos, _, err := GetHistories(context.Background(), &npool.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		ClientIP:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.ClientIP},
		UserAgent: &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserAgent},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setupHistory(t)
	defer teardown(t)

	t.Run("createHistory", createHistory)
	t.Run("getHistory", getHistory)
	t.Run("getHistories", getHistories)
}
