package assembler

import (
	"sheepim-user-service/biz/constant"
	"sheepim-user-service/biz/internal/domain"
	"sheepim-user-service/kitex_gen/base"
	"sheepim-user-service/kitex_gen/user"
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
