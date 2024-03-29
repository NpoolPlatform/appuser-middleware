package user

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	role "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role"
	user "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
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
	ret = npool.User{
		EntID:     uuid.NewString(),
		CreatedBy: uuid.NewString(),
		Role:      uuid.NewString(),
		AppID:     uuid.NewString(),
		UserID:    uuid.NewString(),
	}
)

func setupUser(t *testing.T) func(*testing.T) {
	ah, err := app.NewHandler(
		context.Background(),
		app.WithEntID(&ret.AppID, true),
		app.WithCreatedBy(&ret.UserID, true),
		app.WithName(&ret.AppID, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	app1, err := ah.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	ah, err = app.NewHandler(
		context.Background(),
		app.WithID(&app1.ID, true),
	)
	assert.Nil(t, err)

	emailAddress := fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+12000000) //nolint
	passwordHash := uuid.NewString()

	ret.AppName = ret.AppID

	uh1, err := user.NewHandler(
		context.Background(),
		user.WithEntID(&ret.CreatedBy, true),
		user.WithAppID(&ret.AppID, true),
		user.WithEmailAddress(&emailAddress, true),
		user.WithPasswordHash(&passwordHash, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, uh1)
	user2, err := uh1.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, user2)

	uh1, err = user.NewHandler(
		context.Background(),
		user.WithID(&user2.ID, true),
	)
	assert.Nil(t, err)

	rh, err := role.NewHandler(
		context.Background(),
		role.WithEntID(&ret.Role, true),
		role.WithAppID(&ret.AppID, true),
		role.WithCreatedBy(&ret.CreatedBy, true),
		role.WithRole(&ret.Role, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, rh)
	role1, err := rh.CreateRole(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, role1)

	rh, err = role.NewHandler(
		context.Background(),
		role.WithID(&role1.ID, true),
	)
	assert.Nil(t, err)

	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+1000000)            //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+11000000) //nolint
	ret.AppName = ret.AppID

	uh, err := user.NewHandler(
		context.Background(),
		user.WithEntID(&ret.UserID, true),
		user.WithAppID(&ret.AppID, true),
		user.WithPhoneNO(&ret.PhoneNO, true),
		user.WithEmailAddress(&ret.EmailAddress, true),
		user.WithPasswordHash(&passwordHash, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, uh)
	user1, err := uh.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, user1)

	uh, err = user.NewHandler(
		context.Background(),
		user.WithID(&user1.ID, true),
	)
	assert.Nil(t, err)

	ret.RoleID = ret.Role

	return func(*testing.T) {
		_, _ = uh.DeleteUser(context.Background())
		_, _ = uh1.DeleteUser(context.Background())
		_, _ = rh.DeleteRole(context.Background())
		_, _ = ah.DeleteApp(context.Background())
	}
}

func createUser(t *testing.T) {
	req := npool.UserReq{
		EntID:  &ret.EntID,
		AppID:  &ret.AppID,
		RoleID: &ret.Role,
		UserID: &ret.UserID,
	}
	info, err := CreateUser(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateUser(t *testing.T) {
	req := npool.UserReq{
		ID:     &ret.ID,
		RoleID: &ret.Role,
	}
	info, err := UpdateUser(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getUser(t *testing.T) {
	info, err := GetUser(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getUsers(t *testing.T) {
	_, total, err := GetUsers(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 0)
	}
}

func deleteUser(t *testing.T) {
	info, err := DeleteUser(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetUser(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
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

	teardown := setupUser(t)
	defer teardown(t)

	t.Run("createUser", createUser)
	t.Run("updateUser", updateUser)
	t.Run("getUser", getUser)
	t.Run("getUsers", getUsers)
	t.Run("deleteUser", deleteUser)
}
