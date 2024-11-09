package converter

import (
	"github.com/li1553770945/sheepim-user-service/biz/constant"
	"github.com/li1553770945/sheepim-user-service/biz/internal/domain"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/base"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user"
)

func AssembleSuccessBaseResp() *base.BaseResp {
	return &base.BaseResp{
		Code: constant.Success,
	}
}

func UserInfoEntityToDTO(userEntity *domain.UserEntity) *user.UserInfoResp {
	return &user.UserInfoResp{
		BaseResp: AssembleSuccessBaseResp(),
		UserId:   &userEntity.ID,
		Username: &userEntity.Username,
		Nickname: &userEntity.Nickname,
		Role:     &userEntity.Role,
	}
}
