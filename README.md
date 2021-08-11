## Start

1. create the file `config.yaml` in the root directory

```yaml
support:
  addr: * # start port
  home: https://example.com # website homepage
  query: [ username,email,number ] # table fields allowed to be queried
  ip: [ 127.0.0.1,172.18.0.1 ] # allowed ip
  suffixes: [ ] # email suffixes allowed to sign up
  admin: [ ] # admin email address
  
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