package config

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func PostgresqlClient(user, password, host, port, name string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, name, port)
	pgc, _ := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt: true,
	})
	sqlDB, _ := pgc.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	return pgc
}

func RedisClient(host, port, password string, DB int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       DB,
	})
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		panic("redis连接失败,检查config.yaml配置文件设置,节点redis是否设置正确")
	}
	return rdb
}

func MongoClient(user, password, host, port, database, collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin",
		user, password, host, port)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mgc, _ := mongo.Connect(ctx, options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(100).
		SetMinPoolSize(5))
	err := mgc.Ping(ctx, nil)
	if err != nil {
		panic("mongo连接失败,检查config.yaml配置文件是否设置正确,节点mongo是否设置正确")
	}
	coll := mgc.Database(database).Collection(collection)
	return coll
}
