# NGB-SE API

## API

### 使用邮箱验证

#### 1、发送邮件 POST /user/get-login-code

- 请求体：

```json
{
    "Uid": "yourUid"
}
```

- 返回数据
- 返回体

成功

```json
{
    "status": "success",
    "data":""
}
```

失败

```json
    "status":"fail",
    "data":"$ERROR_MSG"
```
#### 2、发送收到的验证码 POST /user/send-login-code

- 请求体
```json
{
    "Uid": "yourUid",
    "Code": "YourAuthCode"
}
```