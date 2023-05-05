# NGB-SE
这是（将会是）一个开箱即用的论坛系统，目前实现了基本的用户管理
```
Tree
├──config
|   ├──readConfig.go
├──controller
|   ├──admin.go
|   ├──commonUser.go
|   ├──nodeController.go
|   ├──responseStruct.go
|   ├──JWTMiddleWare.go
├──middleware
|   ├──JWTMiddleWare.go
|   ├──userVerification.go
├──model
|   ├──connectMysql.go
|   ├──message.go
|   ├──messageModel.go
|   ├──nodes.go
|   ├──nodesModel.go
|   ├──toolFunc.go
|   ├──userAccount.go
|   ├──userAccountModel.go
|   ├──userRelation.go
|   ├──userRelationModel.go
├──view
|   ├──router.go
├──API.yaml
├──configTemplate.yaml
├──server.go
├──go.mod
├──go.sum
├──README.go
```