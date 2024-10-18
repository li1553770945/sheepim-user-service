package container

import (
	"sheepim-user-service/biz/infra/config"
	"sheepim-user-service/biz/internal/service"
)

type Container struct {
	Config      *config.Config
	UserService user.IUserService
}

func NewContainer(config *config.Config, userService user.IUserService,
) *Container {
	return &Container{
		Config:      config,
		UserService: userService,
	}

}
