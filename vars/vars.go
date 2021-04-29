package vars

import "Server/config"

var (
	conf config.Config
	c    = conf.Yaml()

	PDB0 = config.PostgresqlClient(c.Postgresql.User[0], c.Postgresql.Password[0], c.Postgresql.Host[0], c.Postgresql.Port[0], c.Postgresql.Name[0])

	RDBCODE  = config.RedisClient(c.Redis.Host[0], c.Redis.Port[0], c.Redis.Password[0], 0)
	RDBTOKEN = config.RedisClient(c.Redis.Host[0], c.Redis.Port[0], c.Redis.Password[0], 1)

	MDBSUFFIXE = config.MongoClient(c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0], "conf", "suffixes")
	MDBHTTP    = config.MongoClient(c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0], "log", "http")
	MDBERROR   = config.MongoClient(c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0], "log", "error")
	MDBDRAFT   = config.MongoClient(c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0], "article", "draft")
	MDBPUBLISH = config.MongoClient(c.Mongo.User[0], c.Mongo.Password[0], c.Mongo.Host[0], c.Mongo.Port[0], "article", "publish")

	KEYTOKEN = c.Key.Token

	PROXYADDR = c.Proxy.Addr

	MailForm     = c.Mail.From
	MailUser     = c.Mail.User
	MailPassword = c.Mail.Password
	MailSmtp     = c.Mail.Smtp
	MailPort     = c.Mail.Port
)
