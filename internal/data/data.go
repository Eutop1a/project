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
	query3 "helloworld/internal/data/knowledge/query"
	query2 "helloworld/internal/data/question/query"
	query1 "helloworld/internal/data/user/query"
)

// MySQLQuestionDB 表示问题相关的 MySQL 数据库连接
type MySQLQuestionDB struct {
	*gorm.DB
}

// MySQLUserDB 表示用户相关的 MySQL 数据库连接
type MySQLUserDB struct {
	*gorm.DB
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewAdminRepo,
	NewQuestionRepo,
	NewKnowledgeRepo,
	NewRedis,
	NewMySQLQuestion,
	NewMySQLUser,
)

// Data .
type Data struct {
	questionDB *MySQLQuestionDB
	userDB     *MySQLUserDB
	redis      *redis.Client
}

// NewData .
func NewData(
	questionDB *MySQLQuestionDB,
	userDB *MySQLUserDB,
	redis *redis.Client,
	logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		questionDB: questionDB,
		userDB:     userDB,
		redis:      redis,
	}, cleanup, nil
}

func NewMySQLQuestion(c *conf.Data) *MySQLQuestionDB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.DatabaseQuestion)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//DisableForeignKeyConstraintWhenMigrating: true,
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
	query2.SetDefault(db)
	query3.SetDefault(db)
	return &MySQLQuestionDB{db}
}

func NewMySQLUser(c *conf.Data) *MySQLUserDB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.DatabaseUser)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//DisableForeignKeyConstraintWhenMigrating: true,
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
	query1.SetDefault(db)
	return &MySQLUserDB{db}
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
