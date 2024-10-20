package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"sheepim-user-service/biz/constant"
	"sheepim-user-service/biz/internal/assembler"
	"sheepim-user-service/kitex_gen/base"
	"sheepim-user-service/kitex_gen/user"
)

func (s *UserService) CheckUsernameAndPasswd(ctx context.Context, req *user.CheckUsernameAndPasswdReq) (resp *user.CheckUsernameAndPasswdResp, err error) {

	findUser, err := s.Repo.FindUserByUsername(req.Username)
	if err != nil {
		resp = &user.CheckUsernameAndPasswdResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: "系统错误",
			},
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(req.Password))
	if findUser.ID == 0 || err != nil {
		resp = &user.CheckUsernameAndPasswdResp{
			BaseResp: &base.BaseResp{
				Code:    constant.Unauthorized,
				Message: "用户名或密码错误",
			},
		}
		err = nil
		return
	}

	if findUser.CanUse == false {
		resp = &user.CheckUsernameAndPasswdResp{
			BaseResp: &base.BaseResp{
				Code:    constant.Unauthorized,
				Message: "抱歉，您的账户已被禁用",
			},
		}
		return
	}

	userId := findUser.ID
	resp = &user.CheckUsernameAndPasswdResp{
		BaseResp: &base.BaseResp{
			Code: constant.Success,
		},
		UserId: &userId,
	}
	return
}

func (s *UserService) GetUserInfo(ctx context.Context, req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	findUser, err := s.Repo.FindUserById(req.UserId)
	if err != nil {
		resp = &user.UserInfoResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: "系统错误",
			},
		}
		return
	}
	resp = assembler.UserInfoEntityToDTO(findUser)
	return
}

func (s *UserService) AddUser(ctx context.Context, req *user.AddUserReq) (resp *user.AddUserResp, err error) {

	findUser, err := s.Repo.FindUserByUsername(req.Username)
	if err != nil {
		resp = &user.AddUserResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: "系统错误",
			},
		}
		return
	}
	if findUser.ID == 0 {
		resp = &user.AddUserResp{
			BaseResp: &base.BaseResp{
				Code:    constant.Unauthorized,
				Message: "请先联系管理员获取激活码",
			},
		}
		return
	}

	if findUser.ActivateCode != req.ActiveCode {
		resp = &user.AddUserResp{
			BaseResp: &base.BaseResp{
				Code:    constant.Unauthorized,
				Message: "激活码错误",
			},
		}
		return
	}

	if findUser.IsActivate == true {
		resp = &user.AddUserResp{
			BaseResp: &base.BaseResp{
				Code:    constant.Unauthorized,
				Message: "用户名已被注册",
			},
		}
		return
	}

	findUser.IsActivate = true

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost) //加密处理
	encodePWD := string(hash)                                                          // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	findUser.Password = encodePWD
	findUser.Nickname = req.Nickname
	err = s.Repo.SaveUser(findUser)
	if err != nil {
		resp = &user.AddUserResp{
			BaseResp: &base.BaseResp{
				Code:    constant.SystemError,
				Message: "系统错误",
			},
		}
		return
	}

	resp = &user.AddUserResp{
		BaseResp: &base.BaseResp{
			Code: constant.Success,
		},
		UserId: &findUser.ID,
	}
	return
}
