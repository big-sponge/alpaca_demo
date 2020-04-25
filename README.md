# 开发说明

1、将代码下载到GOPATH的src目录下面，并且将项目目录名称改成 alpaca_demo

2、数据库脚步在 doc/sql下面 

3、环境相关的配置文件在dev下面，具体使用哪一个文件是由环境变量ENV_MODE决定，默认是使用development.json

4、运行或者调试cmd/alpaca下面的main函数

5、浏览器中打开http://127.0.0.1:9999/web/admin, 默认用户名admin，密码111111


# 目录说明

主要的几个目录的解释如下

---

-- **app** 目录放项目代码

-- **app/api** 放接口调用的入口函数

-- **app/api/admin** 是后台管理的模块的接口

-- **app/api/v1** 是前台模块的接口

-- **app/bootstrap** 目录放项目启动时候，初始化的一个函数，例如加载配置文件，配置log，启动gin的路由监听等

-- **app/common** 存放公共调用函数，例如接口的返回状态码，日志log，链接mysql，加密、解密函数等待

-- **app/config** 存放配置文件的目录，要注意这里和env目录的区别，env目录下是为了项目在不同的环境下加载不同的配置文件，一些固定不变的配置可以放到 app/config下


-- **app/models** 存放数据库访问的相关函数，一般是一个数据表 对应一个文件

-- **app/routers** 存放路由配置信息

-- **app/service** 存放编写复杂的业务逻辑

---

-- **cmd** 存放项目的入口文件

---

-- **doc** 存放一些文档资料，例如sql脚步也放到这里了

---

-- **env** 存放环境相关的配置，具体使用哪一个文件是由环境变量ENV_MODE决定，默认是使用development.json

---

-- **storage** 存放程序运行时的一些配置，例如日志，程序pid默认放到这里了

---

-- **vendor** 项目的依赖包在这里，用govendor管理

---

-- **web** 静态资源目录在这里

---

## 设计架构

**bootstrap** 包是启动相关的函数。main函数调用bootstrap包中的函数初始化配置信息，如加载配置文件，配置log，开启http监听等。

**common包** 里面是一些公用函数，任何其他包都可以引用common包，但是common包禁止引用除了vendor以外的其他一切包

**api接口** 函数可以引用common包、models包、services包等实现具体的业务逻辑

**services包** 函数可以引用common包，models包实现具体的业务逻辑

**models包** 可以引用common包，但是不可以引用services包


![img](https://oscimg.oschina.net/oscnet/up-75905893702a4524a3d38ac1e7be11c7f02.png)

