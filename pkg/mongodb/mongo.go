package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Hosts    string `json:"hosts" yaml:"hosts"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	DBName   string `json:"db_name" yaml:"db_name"`
	ReplName string `json:"repl_name" yaml:"repl_name"`
}

type Mongo struct {
	*mongo.Client
	*mongo.Database
}

func NewMongo(cfg Config) *Mongo {
	uri := fmt.Sprintf("mongodb://%s", cfg.Hosts)
	opt := options.Client().ApplyURI(uri)
	if cfg.Username != "" && cfg.Password != "" {
		credential := options.Credential{
			Username:   cfg.Username,
			Password:   cfg.Password,
			AuthSource: cfg.DBName,
		}
		opt = opt.SetAuth(credential)
	}
	if cfg.ReplName != "" {
		opt = opt.SetReplicaSet(cfg.ReplName)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(err)
	}
	return &Mongo{Client: client, Database: client.Database(cfg.DBName)}
}

func (m *Mongo) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.Client.Disconnect(ctx)
}
