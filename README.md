## Start

1. create the file `config.json` in the root directory

```json5
{
  "support": {
    "addr": "",
    "home": "",
    "query": [
    ],
    "ip": [
    ],
    "suffixes": [
    ],
    "admin": []
  },
  "postgresql": {
    "user": [],
    "password": [],
    "host": [],
    "port": [],
    "name": []
  },
  "worker": {
    "workerId": 1,
    "centerId": 1,
    "sequence": 0,
    "epoch": 1609430400000
  },
  "mongo": {
    "user": [],
    "password": [],
    "host": [],
    "port": []
  },
  "redis": {
    "host": [],
    "port": [],
    "password": []
  },
  "mail": {
    "user": "",
    "from": "",
    "password": "",
    "smtp": "",
    "port": ""
  },
  "key": {
    "token": ""
  }
}
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