# 1.接口文档

## (1)用户相关

### 用户注册

1. **接口名称：** 用户注册
2. **接口地址：**`/user/register`
3. **请求方式：**`POST` , `key` : `Content-Type` , `Value` : `application/json`
4. **请求参数：**

   | 请求参数           | 类型   | json中名称                 | 是否必需 |
   | -------------------- | -------- | ---------------------------- | ---------- |
   | 登录用户名         | string | username                   | yes      |
   | 密码               | string | password                   | yes      |
   | 权限               | string | role (包括"admin"和"user") | yes      |
   | 用户名(对外展示的) | string | nickname                   | yes      |


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
       "info": "OK"
   }
   ```

   用户名重复(用户名不可用):

   ```json
   {
       "status": 400,
       "info": "invalid username"
   }
   ```

   其他错误(例如操作数据库等等错误)会在info中展示:

   ```json
   {
   "status": 500,
   "info": "error..."
   }
   ```

### 用户登录

1. **接口名称：** 用户登录
2. **接口地址：**`/user/login`
3. **请求方式：**`POST`, `key` : `Content-Type` , `Value` : `application/json`
4. **请求参数：**

   | 请求参数   | 类型   | json中名称 | 是否必需 |
   | ------------ | -------- | ------------ | ---------- |
   | 登录用户名 | string | username   | yes      |
   | 密码       | string | password   | yes      |


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
       "status": 400,
       "info": "Wrong Password!"
   }
   ```

   其他错误(例如操作数据库等等错误)会在info中展示:

   ```json
   {
   "status": 500,
   "info": "error..."
   }
   ```

## (2)留言相关

### 发送留言

1. **接口名称：** 发送留言
2. **接口地址：**`/message/submit`
3. **请求方式：**`POST`, `key` : `Content-Type` , `Value` : `application/json`
4. **请求参数：**

   | 请求参数 | 类型   | json中名称 | 是否必需 |
   | ---------- | -------- | ------------ | ---------- |
   | 用户id   | int    | userid     | yes      |
   | 留言内容 | string | content    | yes      |


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
    "status": 400,
    "info": "invalid userid"
   }
   ```

   其他错误(例如操作数据库等等错误)会在info中展示:

   ```json
   {
   "status": 500,
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
   {
       "messages": [
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
       ],
       "respond code": {
           "status": 200,
           "info": "OK"
       }
   }
   ```

   其他错误(例如操作数据库等等错误)会在info中展示:

   ```json
   {
   "status": 500,
   "info": "error..."
   }
   ```

### 删除留言

1. **接口名称：** 删除留言
2. **接口地址：**`/admin/message/delete`
3. **请求方式：**`DELETE` , `key` : `Content-Type` , `Value` : `application/json`
4. **请求参数：**

   | 请求参数 | 类型 | json中名称 | 是否必需 |
   | ---------- | ------ | ------------ | ---------- |
   | 留言id   | int  | id         | yes      |


5. **请求示例：**

   ```json
   {
       "id":1
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
   "status": 500,
   "info": "error..."
   }
   ```
