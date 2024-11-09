// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	configInfra "github.com/li1553770945/sheepim-user-service/biz/infra/config"
	"github.com/li1553770945/sheepim-user-service/biz/infra/container"
	"github.com/li1553770945/sheepim-user-service/biz/infra/log"
	"github.com/li1553770945/sheepim-user-service/biz/infra/trace"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user/userservice"
	"net"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	config := configInfra.GetConfig(env)
	log.InitLog()
	p := trace.InitTrace(config)

	container.InitGlobalContainer(config)
	App := container.GetGlobalContainer()
	defer func(p provider.OtelProvider, ctx context.Context) {
		err := p.Shutdown(ctx)
		if err != nil {
			klog.Fatalf("server stopped with error:%s", err)
		}
	}(p, context.Background())

	addr, err := net.ResolveTCPAddr("tcp", App.Config.ServerConfig.ListenAddress)
	if err != nil {
		panic("设置监听地址出错")
	}

	r, err := etcd.NewEtcdRegistry(App.Config.EtcdConfig.Endpoint) // r should not be reused.
	if err != nil {
		panic(fmt.Sprintf("初始化注册中心失败:%v", err))
	}
	svr := userservice.NewServer(
		new(UserServiceImpl),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerConfig.ServiceName}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
	)
	if err := svr.Run(); err != nil {
		klog.Fatalf("服务启动失败:", err)
	}
}
