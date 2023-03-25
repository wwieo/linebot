package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type mongoDB struct {
	*mongo.Database
}

// initialize initializes the MySQL dao handle.
func (db *mongoDB) initialize(ctx context.Context, cfg dbConfig) {
	URI := fmt.Sprintf("%s:%d", cfg.Url, cfg.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := client.Ping(pingCtx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
		return
	}

	db.Database = client.Database(cfg.Database)
}

func (db *mongoDB) db() interface{} {
	return db.Database
}
