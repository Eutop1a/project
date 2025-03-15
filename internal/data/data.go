package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloworld/internal/conf"
	"helloworld/internal/data/dal/query"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewAdminRepo,
	NewRedis,
	NewMySQL,
)

// Data .
type Data struct {
	mysql *gorm.DB
	redis *redis.Client
}

// NewData .
func NewData(
	mysql *gorm.DB,
	redis *redis.Client,
	logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		mysql: mysql,
		redis: redis,
	}, cleanup, nil
}

func NewMySQL(c *conf.Data) *gorm.DB {
	host, port, username, password, dbname := c.Mysql.Host, c.Mysql.Port, c.Mysql.User, c.Mysql.Password, c.Mysql.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(int(c.Mysql.MaxIdle))
	sqlDB.SetMaxOpenConns(int(c.Mysql.MaxOpen))
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}
	query.SetDefault(db)
	return db
}

func NewRedis(c *conf.Data) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", c.Redis.Host, c.Redis.Port),
		Password:     c.Redis.Password,
		PoolSize:     int(c.Redis.PoolSize),
		MinIdleConns: int(c.Redis.MinIdleConns),
		MaxRetries:   int(c.Redis.MaxRetries),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		PoolTimeout:  c.Redis.PoolTimeout.AsDuration(),
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	return client
}
