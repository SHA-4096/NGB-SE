# NGB-SE

这是（将会是）一个开箱即用的论坛系统，目前实现的功能详见api和router.go

## Layout:

```
NGB-SE
├─ .gitignore
├─ README.md
├─ cmd
│  └─ server.go
├─ config
│  ├─ config.yaml
├─ docs
│  └─ API.yaml
├─ go.mod
├─ go.sum
└─ internal
   ├─ conf
   │  └─ readConfig.go
   ├─ controller
   │  ├─ admin.go
   │  ├─ commonUser.go
   │  ├─ messageController.go
   │  ├─ nodeController.go
   │  ├─ responseStruct.go
   │  └─ userRelationsController.go
   ├─ middleware
   │  ├─ JWTMiddleWare.go
   │  └─ userVerification.go
   ├─ model
   │  ├─ connetMysql.go
   │  ├─ message.go
   │  ├─ messageModel.go
   │  ├─ nodes.go
   │  ├─ nodesModel.go
   │  ├─ toolFunc.go
   │  ├─ userAccount.go
   │  ├─ userAccountModel.go
   │  ├─ userRelation.go
   │  └─ userRelationModel.go
   └─ view
      └─ router.go

```
