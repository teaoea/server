<h1 style="text-align: center">使用GIN开发后端API服务</h1>

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

2. create the database `conf`and the collection `suffixes` in the database,
   and import the file `script/suffixes.json`
