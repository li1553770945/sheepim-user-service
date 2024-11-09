package container

import (
	"github.com/li1553770945/sheepim-user-service/biz/infra/config"
	"github.com/li1553770945/sheepim-user-service/biz/internal/service"
	"sync"
)

type Container struct {
	Config      *config.Config
	UserService user.IUserService
}

var APP *Container
var once sync.Once

func GetGlobalContainer() *Container {
	if APP == nil {
		panic("APP在使用前未初始化")
	}
	return APP
}

func InitGlobalContainer(config *config.Config) {
	once.Do(func() {
		APP = GetContainer(config)
	})
}
func NewContainer(config *config.Config, userService user.IUserService,
) *Container {
	return &Container{
		Config:      config,
		UserService: userService,
	}

}
