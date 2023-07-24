# NGB-SE

这是（将会是）一个开箱即用的论坛系统，目前支持：用户的注册登录，帖子的发布以及管理，邮件订阅等

## Layout:

```
NGB-SE
├── README.md
├── build
│   ├── nginx-conf
│   │   └── proxy.conf
│   └── package
│       └── Dockerfile
├── cmd
│   └── server.go
├── config
│   ├── config-template.yaml
│   └── config.yaml
├── docs
│   ├── API-new.md
│   └── API.yaml
├── go.mod
├── go.sum
└── internal
    ├── conf
    │   └── read-config.go
    ├── controller
    │   ├── admin.go
    │   ├── common-user.go
    │   ├── message-controller.go
    │   ├── node-controller.go
    │   ├── param
    │   │   └── router.go
    │   ├── response-struct.go
    │   └── user-relations-controller.go
    ├── middleware
    │   ├── JWT-middleware.go
    │   ├── subscription.go
    │   └── user-verification.go
    ├── model
    │   ├── connect-redis.go
    │   ├── connet-mysql.go
    │   ├── message-model.go
    │   ├── message.go
    │   ├── nodes-model.go
    │   ├── nodes.go
    │   ├── tool-func.go
    │   ├── user-account-model.go
    │   ├── user-account.go
    │   ├── user-relation-model.go
    │   └── user-relation.go
    ├── util
    │   ├── email.go
    │   ├── log.go
    │   └── random-generate.go
    └── view
        └── router.go

```

# Deploy

前提：服务器有mysql,redis和docker服务

将 `./config/config-template.yaml`文件重命名为 `config.yaml`

在NGB-SE目录下执行以下命令

```bash
sudo docker build -t ngb -f ./build/package/Dockerfile .
```

要运行容器：

```bash
sudo docker run --network=host -v /etc/localtime:/etc/localtime ngb
```

其中 `-v /etc/localtime:/etc/localtime`是为了使得容器内的时间和系统时间一致

# API

所有的api由postman导出为json文档了，可以查看./docs/postman-api.json

后期会考虑写一个更加清楚的markdown文档
