package main

import (
	"context"
	"sheepim-user-service/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CheckUsernameAndPasswd implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUsernameAndPasswd(ctx context.Context, req *user.CheckUsernameAndPasswdReq) (resp *user.CheckUsernameAndPasswdResp, err error) {
	resp, err = App.UserService.CheckUsernameAndPasswd(ctx, req)
	// TODO: Your code here...
	return
}

// AddUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddUser(ctx context.Context, req *user.AddUserReq) (resp *user.AddUserResp, err error) {
	resp, err = App.UserService.AddUser(ctx, req)
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	// TODO: Your code here...
	resp, err = App.UserService.GetUserInfo(ctx, req)
	return
}
