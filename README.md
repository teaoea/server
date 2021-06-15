## Start

1. create the file `config.yaml` in the root directory

```yaml
postgresql:
  user: [ ]
  password: [ ]
  host: [ ]
  port: [ ]
  name: [ ]

worker:
  workerId: 1
  centerId: 1
  sequence: 0
  # timestamp
  epoch: 1609430400000

mail:
  user:
  from:
  password:

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
  # the key to encrypt the token
  token:

```

2. mongo
    * create the database `conf`and the collection `suffixes` in the database, and import the file `docs/suffixes.json`
    * create the collection `query` in the database `conf` and import the file `docs/query.json`
    * create the collection `ipaddr` in the database `conf` and import the file `docs/ip.json`

3. postgres
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
    * [go-simple-mail](https://github.com/xhit/go-simple-mail)
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
    * [jwt-go](https://github.com/dgrijalva/jwt-go)