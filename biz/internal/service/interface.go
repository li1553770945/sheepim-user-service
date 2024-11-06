package user

import (
	"context"
	"github.com/li1553770945/sheepim-user-service/biz/internal/repo"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user"
)

type UserService struct {
	Repo repo.IRepository
}

type IUserService interface {
	CheckUsernameAndPasswd(ctx context.Context, req *user.CheckUsernameAndPasswdReq) (*user.CheckUsernameAndPasswdResp, error)
	GetUserInfo(ctx context.Context, req *user.UserInfoReq) (resp *user.UserInfoResp, err error)
	AddUser(ctx context.Context, req *user.AddUserReq) (resp *user.AddUserResp, err error)
}

func NewUserService(repo repo.IRepository) IUserService {
	return &UserService{
		Repo: repo,
	}
}
