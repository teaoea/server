## Start

1. create the file `config.yaml` in the root directory

```yaml
# start port :1234
addr: * # port

# website homepage
home: https://example.com

# table fields allowed to be queried
query: [ username,email,number ]

# allowed ip
ip: [ 127.0.0.1,172.18.0.1 ]

postgresql:
  user: [ ]
  password: [ ]
  host: [ ]
  port: [ ]
  name: [ ]

mongo:
  user: [ ]
  password: [ ]
  host: [ ]
  port: [ ]

redis:
  host: [ ]
  port: [ ]
  password: [ ]

worker:
  workerId: 1
  centerId: 1
  sequence: 0
  # timestamp
  epoch: 1609430400000

mail:
  user: *
  from: *
  password: *
  smtp: *
  post: *
  # service error, send mail to admin
  admin: [ ]
  # email suffixes allowed to sign up
  suffixes: [
      @gmail.com,
      @outlook.com,
      @yahoo.com,
      @googlemail.com,
      @live.com,
      @icould.com,
      @mail.com,
      @email.com,
      @ask.com,
      @msn.com,
      @163.com,
      @126.com,
      @qq.com,
      @hotmail.com
  ]

key:
  # the key to encrypt the token
  token:

```

2. postgres
    * the postgresql database table file is saved in the `script` directory

## Dependency

1. database
    * [postgresql](https://www.postgresql.org/)
    * [mongo](https://www.mongodb.com/)
    * [redis](https://redis.io/)

2. library
    * [gin](https://gin-gonic.com/)
    * [gorm](https://gorm.io/)
    * [go-redis](https://redis.uptrace.dev/)
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
    * [jwt-go](https://github.com/dgrijalva/jwt-go)