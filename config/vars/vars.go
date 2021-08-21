package vars

import (
	"fmt"
	"time"

	"server/config"
)

var (
	conf config.Config
	c    = conf.Get()

	Addr = c.Support.Addr

	Home = c.Support.Home

	KeyToken = c.Key.Token

	Query = c.Support.Query

	EmailSuffixes = c.Support.Suffixes

	Admin = c.Support.Admin

	Ip = c.Support.Ip

	DB0 = config.PostgresqlClient(
		c.Postgresql.User[0], c.Postgresql.Password[0], c.Postgresql.Host[0],
		c.Postgresql.Port[0], c.Postgresql.Name[0],
	)

	RedisToken = config.RedisClient(
		c.Redis.Host[0], c.Redis.Port[0], c.Redis.Password[0], 0,
	)

	RedisLogoff = config.RedisClient(
		c.Redis.Host[0], c.Redis.Port[0], c.Redis.Password[0], 1,
	)

	RedisAuthCode = config.RedisClient(
		c.Redis.Host[0], c.Redis.Port[0], c.Redis.Password[0], 2,
	)

	RedisPasswordCode = config.RedisClient(
		c.Redis.Host[0], c.Redis.Port[0], c.Redis.Password[0], 3,
	)

	RedisEmailCode = config.RedisClient(
		c.Redis.Host[0], c.Redis.Port[0], c.Redis.Password[0], 4,
	)

	MongoHttp = config.MongoClient(
		c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0],
		"log", fmt.Sprintf("%d-%d-%d,%d:%d",
			time.Now().Year(), time.Now().Month(), time.Now().Day(),
			time.Now().Hour(), time.Now().Minute(),
		),
	)

	MongoError = config.MongoClient(
		c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0],
		c.Mongo.Port[0], "error", fmt.Sprintf("%d-%d-%d,%d:%d",
			time.Now().Year(), time.Now().Month(), time.Now().Day(),
			time.Now().Hour(), time.Now().Minute(),
		),
	)

	MongoAngularLogger = config.MongoClient(
		c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0],
		c.Mongo.Port[0], "angular-logger", fmt.Sprintf("%d-%d-%d,%d:%d",
			time.Now().Year(), time.Now().Month(), time.Now().Day(),
			time.Now().Hour(), time.Now().Minute(),
		),
	)

	MongoAngularError = config.MongoClient(
		c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0],
		c.Mongo.Port[0], "angular-error", fmt.Sprintf("%d-%d-%d,%d:%d",
			time.Now().Year(), time.Now().Month(), time.Now().Day(),
			time.Now().Hour(), time.Now().Minute(),
		),
	)
)
