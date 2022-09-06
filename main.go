package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Handler struct {
	Mongo *mongo.Client
}

func (h *Handler) initMongo() *Handler {
	ctx := context.Background()

	connStr := fmt.Sprintf(
		"mongodb://%s:%s@%s/%s",
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASS"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_AUTH"),
	)

	replicaSet := os.Getenv("MONGO_REPLICA")
	if replicaSet != "" {
		connStr += fmt.Sprintf("?replicaSet=%s", replicaSet)
	}

	fmt.Println("connStr: ", connStr)

	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		panic(err)
	}

	if err := client.Connect(ctx); err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}

	h.Mongo = client
	return h
}

func main() {
	h := new(Handler)
	h.initMongo()
	log.Println("start")

	now := time.Now()
	after := now.Add(60 * time.Minute)

	for {
		if err := h.Mongo.Ping(context.Background(), readpref.Primary()); err != nil {
			log.Fatalf("%+v\n", err)
		} else {
			log.Println("success")
		}
		time.Sleep(3 * time.Second)

		now = time.Now()
		if now.After(after) {
			break
		}
	}

	log.Println("done")
}
