# 使用GIN开发后端API服务

## 启动

1. 在config目录创建config.yaml文件,按照以下代码块添加参数

```yaml
postgresql:
  user: [ ]
  password: [ ]
  host: [ ]
  port: [ ]
  name: [ ]

worker:
  workerId: # 机器id
  centerId: # 数据中心id
  sequence: # 序列号起始值
  epoch: # 开始时间戳，精确到毫秒

mail:
  user: # 发件人邮箱
  from: # 从哪里发送的邮件
  password: # 发件人邮箱授权码

mongo:
  user: [ ]
  password: [ ]
  host: [ ]
  port: [ ]

redis:
  host: [ ]
  port: [ ]
  password: [ ]

key:
  token: # 加密token使用的秘钥，越复杂越好

```

2. 在数据库MongoDB,创建数据库名`conf`,集合名`suffixes`,并导入文件`.\script\suffixes.json`,集合`suffixes`
   是用来校验注册时,邮箱后缀是否合法,后续想添加其他邮箱后缀,直接在集合里面添加,而不需要重新编译项目,字段名只能是`suffix`,字段名错误,添加的邮箱后缀无法生效

3. script存放项目需要的所有表,数据库使用的postgresql

## 依赖项

1. [gin](https://github.com/gin-gonic/gin)
2. [gorm](https://github.com/go-gorm/gorm)
3. [redis](https://github.com/go-redis/redis)
4. [mongo-driver](https://github.com/mongodb/mongo-go-driver)
5. [yaml_v3](https://github.com/go-yaml/yaml/tree/v3)
6. [jwt-go](https://github.com/dgrijalva/jwt-go)

