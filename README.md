# sheepim用户服务

## 初始化项目
```bash
kitex -module "sheepim-user-service" -service sheepim-user-service idl/user.thrift
cd biz/infra/container
wire
```

## 开发环境

```bash
export ENV=development
```

## 生产环境

```bash
export ENV=production
```