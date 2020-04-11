package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type key string

const (
	protocol = "tcp"
	port     = ":50051"

	host     = "localhost"
	username = "dimzrio"
	password = "dimzrio123"
	database = "admin"
)

type dataItem struct {
	AppsName string `bson:"apps_name"`
	Language string `bson:"language"`
}

func insertDB(dbs *mongo.Database) (string, error) {
	data := &dataItem{
		AppsName: "deployer-tools",
		Language: "python3",
	}
	res, err := dbs.Collection("grpc").InsertOne(context.Background(), data)

	if err != nil {
		return "", err
	}

	log.Println("Insert to collection successfully...")
	oid := res.InsertedID.(primitive.ObjectID).Hex()
	return oid, nil
}

func configBD(ctx context.Context) *mongo.Database {
	url := fmt.Sprintf(`mongodb://%s:%s@%s:27017/%s`,
		ctx.Value(key("username")).(string),
		ctx.Value(key("password")).(string),
		ctx.Value(key("host")).(string),
		ctx.Value(key("database")).(string))

	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed connect to db: %v\n", err)
		os.Exit(1)
	}

	dbs := client.Database(database)
	return dbs
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, key("host"), host)
	ctx = context.WithValue(ctx, key("username"), username)
	ctx = context.WithValue(ctx, key("password"), password)
	ctx = context.WithValue(ctx, key("database"), database)

	db := configBD(ctx)
	log.Println("Connect to db is successfully...")

	resp, err := insertDB(db)

	if err != nil {
		log.Fatalf("Found error: %v", err)
		os.Exit(1)
	}

	log.Println(resp)
}
