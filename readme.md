# 1.项目基本介绍
本项目为一个简单的留言板的后端部分，可以实现用户注册、登录、发表留言、获取所有留言、
删除留言等功能。
# 2.项目结构
我将功能分成几个大模块：主程序(cmd)、管理路由和接口(api)、
业务逻辑的处理(service)、 存放操作所需的模型(model)、操作数据库(dao)、
一些辅助函数(utils)六大主要模块， 以及本文件(介绍项目和接口文档)和go.mod文件。
## (1)主程序(cmd)
`main.go`:
1. 初始化项目，加载配置。
2. 初始化路由（调用 api/router.go 的方法注册路由）。
3. 启动 HTTP 服务。
## (2)管理路由和接口(api)
1. `router.go`:定义所有的路由，并将路由与对应的接口函数关联。
2. `user.go`:定义用户相关接口。
3. `message.go`:定义留言相关接口。
## (3)业务逻辑的处理(service)
处理业务逻辑，是对` dao `层的进一步封装和逻辑控制。
1. `user.go`:用户相关逻辑(例如注册登录等)。
2. `message.go`:留言相关逻辑。
## (4)存放操作所需的模型(model)
1. `user.go`:用户数据结构。
2. `message.go`:留言数据结构。
## (5)操作数据库(dao)
负责直接与数据库交互，封装数据访问方法。
1. `user.go`:用户相关数据库操作。
2. `message.go`:留言相关数据库操作。
## (6)一些辅助函数(utils)
封装一些重复使用的变量和函数。  
`response.go`:<br>封装了一些重复使用的变量(例如成功，失败等)，还有一些包含自定义消息的成功和失败消息(例如自定义成功和失败消息)。
# 3.接口文档
## (1)用户相关
### 用户注册
1. **接口名称：** 用户注册
2. **接口地址：**`/user/register`
3. **请求方式：**`POST`
4. **请求参数：**
   | 参数名  | json中名称 |
   |--------|------|
   |用户名 | username  |
   | 密码   | password   |
   | 权限| role (包括"admin"和"user")   |
   |匿名|nickname|
5. **请求示例：**
    ```json
    {
       "username":"test",
       "password":"1234567",
       "role":"user",
       "nickname":"test"
    }
    ```
6. **响应示例：(返回的状态码和下面status相同)**
    <br>注册成功:
    ```json
   {
    "status": 200,
    "info": "成功插入数据"
    }
   ```
   用户名重复(用户名不可用):
   ```json
   {
   "status": 404,
   "info": "invalid-username"
   }
   ```
   其他错误(例如操作数据库等等错误)会在info中展示:
   ```json
   {
   "status": 404,
   "info": "error..."
   }
   ```
### 用户登录
1. **接口名称：** 用户登录
2. **接口地址：**`/user/login`
3. **请求方式：**`POST`
4. **请求参数：**
   | 参数名  | json中名称 |
   |--------|------|
   |用户名 | username  |
   | 密码   | password   |
5. **请求示例：**
    ```json
    {
       "username":"test",
       "password":"1234567"
    }
    ```
6. **响应示例：(返回的状态码和下面status相同)**
   <br>登录成功:
    ```json
   {
    "status": 200,
    "info": "OK"
    }
   ```
   密码错误:
   ```json
   {
    "status": 404,
    "info": "Wrong Password!"
   }
   ```
   其他错误(例如操作数据库等等错误)会在info中展示:
   ```json
   {
   "status": 404,
   "info": "error..."
   }
   ```
## (2)留言相关
### 发送留言
1. **接口名称：** 发送留言
2. **接口地址：**`/message/submit`
3. **请求方式：**`POST`
4. **请求参数：**
   | 参数名  | json中名称 |
   |--------|------|
   |用户名 | userid  |
   | 留言内容  | content   |
5. **请求示例：**
    ```json
    {
    "userid": 1,
    "content": "这还是一条测试消息~"
    }
    ```
6. **响应示例：(返回的状态码和下面status相同)**
   <br>发送成功:
    ```json
   {
    "status": 200,
    "info": "OK"
    }
   ```
   用户名无效:
   ```json
   {
    "status": 404,
    "info": "invalid userid"
   }
   ```
   其他错误(例如操作数据库等等错误)会在info中展示:
   ```json
   {
   "status": 404,
   "info": "error..."
   }
   ```
### 查看所有留言
1. **接口名称：** 查看所有留言
2. **接口地址：**`/admin/message/get-all`(管理员专用功能，为后期鉴权做准备)
3. **请求方式：**`GET`
4. **请求参数：** 无
5. **请求示例：** `http://127.0.0.1:8888/admin/message/get-all`
6. **响应示例：(返回的状态码和下面status相同)**
   <br>获取成功:
    ```json
   [
    {
        "id": 2,
        "userid": 1,
        "content": "这还是一条测试消息~",
        "created_at": "2024-12-03T22:27:09+08:00",
        "updated-at": "2024-12-03T22:27:09+08:00",
        "is_deleted": 0,
        "parent_id": null
    },
    {
        "id": 3,
        "userid": 1,
        "content": "这就就是一条测试消息~",
        "created_at": "2024-12-05T16:24:05+08:00",
        "updated-at": "2024-12-05T16:24:05+08:00",
        "is_deleted": 0,
        "parent_id": null
    },
    {
        "id": 4,
        "userid": 2,
        "content": "这就就是一条测试消息~",
        "created_at": "2024-12-05T16:25:16+08:00",
        "updated-at": "2024-12-05T16:29:07+08:00",
        "is_deleted": 0,
        "parent_id": null
    },
    {
        "id": 5,
        "userid": 2,
        "content": "这就就111是一条测试消息~",
        "created_at": "2024-12-05T16:31:15+08:00",
        "updated-at": "2024-12-05T16:31:15+08:00",
        "is_deleted": 0,
        "parent_id": null
    }
    ]
   ```
   其他错误(例如操作数据库等等错误)会在info中展示:
   ```json
   {
   "status": 404,
   "info": "error..."
   }
   ```
### 删除留言
1. **接口名称：** 删除留言
2. **接口地址：**`/admin/message/delete`
3. **请求方式：**`DELETE`
4. **请求参数：**
   | 参数名  | json中名称 |
   |--------|------|
   |留言id | id  |
5. **请求示例：**
    ```json
    {
    "userid": 1,
    "content": "这还是一条测试消息~"
    }
    ```
6. **响应示例：(返回的状态码和下面status相同)**
   <br>删除成功:
    ```json
   {
    "status": 200,
    "info": "message deleted successfully"
    }
   ```
   留言id无效(找不到留言或者留言已经删除):
   ```json
   {
    "status": 404,
    "info": "can't find this message"
   }
   ```
   其他错误(例如操作数据库等等错误)会在info中展示:
   ```json
   {
   "status": 404,
   "info": "error..."
   }
   ```
# 4.如何部署
请先确保本地有go环境并安装了git，创建了仓库。(如果你不想使用git，下载zip然后直接在IDE中打开也是可以的，这样的话，请你跳过第一步)
1. 克隆项目到本地:`git clone https://github.com/LoveLosita/message-board.git`
2. 整理所需依赖:`go mod tidy`
3. 运行:`go run cmd/main.go`