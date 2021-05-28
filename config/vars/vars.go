package vars

import "server/config"

var (
	conf config.Config
	c    = conf.Yaml()

	DB0 = config.PostgresqlClient(c.Postgresql.User[0], c.Postgresql.Password[0], c.Postgresql.Host[0], c.Postgresql.Port[0], c.Postgresql.Name[0])

	RedisCode  = config.RedisClient(c.Redis.Host[0], c.Redis.Port[0], c.Redis.Password[0], 0)
	RedisToken = config.RedisClient(c.Redis.Host[0], c.Redis.Port[0], c.Redis.Password[0], 1)

	MongoSuffix       = config.MongoClient(c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0], "conf", "suffixes")
	MongoIpaddr       = config.MongoClient(c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0], "conf", "ipaddr")
	MongoHttp         = config.MongoClient(c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0], "log", "http")
	MongoError        = config.MongoClient(c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0], "log", "error")
	MongoAngularError = config.MongoClient(c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0], "angular", "error")

	KeyToken = c.Key.Token

	MailForm     = c.Mail.From
	MailUser     = c.Mail.User
	MailPassword = c.Mail.Password
	MailSmtp     = c.Mail.Smtp
	MailPort     = c.Mail.Port
)
