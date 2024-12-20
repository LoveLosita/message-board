# 接口文档
现在改用apifox生成效果更好的接口文档。
- 链接：<a href="https://f3px86u2dx.apifox.cn/" target="_blank">api.md</a>
# 1.项目基本介绍
本项目为一个简单的留言板的后端部分，可以实现用户注册、登录、发表留言、获取所有留言、
删除留言、楼中楼回复留言、给留言点赞、获取用户信息、搜索留言等功能。
# 2.项目结构
我将功能分成几个大模块：主程序(cmd)、管理路由和接口(api)、
业务逻辑的处理(service)、 存放操作所需的模型(model)、操作数据库(dao)、
一些辅助函数(utils)、中间件(middleware)、认证服务(auth)八大主要模块， 以及本文件(介绍项目和接口文档)和go.mod文件。
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
## (7)中间件(middleware)
`jwt_checker.go`:<br>检查token是否有效的中间件。
## (8)认证服务(auth)
`checkPremission.go`:<br>检测用户是否有权限的服务。<br>
`jwt_generage.go`:<br>生成jwt的服务。
# 3.如何部署
请先确保本地有go环境并安装了git，创建了仓库。(如果你不想使用git，下载zip然后直接在IDE中打开也是可以的，这样的话，请你跳过第一步)
1. 克隆项目到本地:`git clone https://github.com/LoveLosita/message-board.git`
2. 整理所需依赖:`go mod tidy`
3. 运行:`go run cmd/main.go`