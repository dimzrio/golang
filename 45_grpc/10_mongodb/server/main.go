package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dimzrio/grpc-mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	protocol = "tcp"
	port     = ":50051"

	host           = "localhost"
	username       = "dimzrio"
	password       = "dimzrio123"
	database       = "admin"
	collectionName = "grpc"
)

type dataItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Author   string             `bson:"apps_id"`
	AppsName string             `bson:"apps_name"`
	Language string             `bson:"language"`
}

type server struct {
}

var db *mongo.Database
var collections *mongo.Collection

// Create fuct to mongoDB
func (*server) Create(ctx context.Context, req *model.CreateReq) (*model.CreateResp, error) {
	author := req.GetDb().GetAuthor()
	appsName := req.GetDb().GetAppsname()
	language := req.GetDb().GetLanguage()

	log.Printf("Req => %v\n", req)

	data := dataItem{
		Author:   author,
		AppsName: appsName,
		Language: language,
	}

	res, err := collections.InsertOne(context.Background(), data)

	if err != nil {
		return nil, err
	}

	log.Println("Insert to collection successfully...")
	oid := res.InsertedID.(primitive.ObjectID)

	resp := &model.CreateResp{
		Db: &model.Database{
			Id:       oid.Hex(),
			Author:   author,
			Appsname: appsName,
			Language: language,
		},
	}

	return resp, nil
}

//Read Func from mongodb
func (*server) Read(ctx context.Context, req *model.ReadReq) (*model.ReadResp, error) {
	ID := req.GetId()
	blogID, err := primitive.ObjectIDFromHex(ID)

	log.Printf("Req => %v\n", req)

	if err != nil {
		errMsg := status.Errorf(codes.InvalidArgument, "Failed read ID")
		return nil, errMsg
	}

	data := &dataItem{}
	filterData := bson.M{"_id": blogID}

	res := collections.FindOne(context.Background(), filterData)

	if err := res.Decode(data); err != nil {
		errMsg := status.Errorf(codes.NotFound, "Can't find your requests")
		return nil, errMsg
	}

	resp := &model.ReadResp{
		Db: &model.Database{
			Id:       data.ID.Hex(),
			Author:   data.Author,
			Appsname: data.AppsName,
			Language: data.Language,
		},
	}

	return resp, nil
}

// Update Func from mongodb
func (*server) Update(ctx context.Context, req *model.UpdateReq) (*model.UpdateResp, error) {
	ID := req.GetDb().GetId()
	blogID, err := primitive.ObjectIDFromHex(ID)

	author := req.GetDb().GetAuthor()
	appsName := req.GetDb().GetAppsname()
	language := req.GetDb().GetLanguage()

	log.Printf("Req => %v\n", req)

	if err != nil {
		errMsg := status.Errorf(codes.InvalidArgument, "Failed read ID")
		return nil, errMsg
	}

	data := &dataItem{}
	filterData := bson.M{"_id": blogID}

	data.ID = blogID
	data.Author = author
	data.AppsName = appsName
	data.Language = language

	res := collections.FindOneAndReplace(context.Background(), filterData, data)

	if err := res.Decode(data); err != nil {
		errMsg := status.Errorf(codes.NotFound, "Can't find your requests")
		return nil, errMsg
	}

	resp := &model.UpdateResp{
		Db: &model.Database{
			Id:       ID,
			Author:   author,
			Appsname: appsName,
			Language: language,
		},
	}

	return resp, nil
}

// Delete func from mongodb
func (*server) Delete(ctx context.Context, req *model.DeleteReq) (*model.DeleteResp, error) {
	ID := req.GetId()
	blogID, err := primitive.ObjectIDFromHex(ID)

	log.Printf("Req => %v\n", req)

	if err != nil {
		errMsg := status.Errorf(codes.InvalidArgument, "Failed read ID")
		return nil, errMsg
	}

	data := &dataItem{}
	filterData := bson.M{"_id": blogID}

	res := collections.FindOne(context.Background(), filterData)

	if err := res.Decode(data); err != nil {
		errMsg := status.Errorf(codes.NotFound, "Can't find your requests")
		return nil, errMsg
	}

	delRecord, err := collections.DeleteOne(context.Background(), filterData)

	if err != nil {
		errMsg := status.Errorf(codes.NotFound, "Failed delete record")
		return nil, errMsg
	}

	resp := &model.DeleteResp{
		Db: &model.Database{
			Id:       data.ID.Hex(),
			Author:   data.Author,
			Appsname: data.AppsName,
			Language: data.Language,
		},
	}

	log.Println(delRecord.DeletedCount)
	return resp, nil
}

// List stream func from mongodb
func (*server) List(req *model.ListReq, stream model.DBService_ListServer) error {
	cursor, err := collections.Find(context.Background(), bson.D{})
	defer cursor.Close(context.Background())

	if err != nil {
		msgErr := status.Errorf(codes.Unavailable, "Failed connect to collections")
		return msgErr
	}

	for cursor.Next(context.Background()) {
		data := &dataItem{}
		err := cursor.Decode(data)

		if err != nil {
			errMsg := status.Errorf(codes.OutOfRange, "Failed read data from mongodb")
			return errMsg
		}

		resp := &model.ListResp{
			Db: &model.Database{
				Id:       data.ID.Hex(),
				Author:   data.Author,
				Appsname: data.AppsName,
				Language: data.Language,
			},
		}

		fmt.Println(resp)
		err = stream.Send(resp)

		time.Sleep(1 * time.Second)

		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func main() {
	log.Println("*** gRPC MongoDB ***")

	// Open connections to db
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := fmt.Sprintf(`mongodb://%s:%s@%s:27017/%s`, username, password, host, database)
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	err = client.Connect(ctx)

	if err != nil {
		log.Printf("Failed connect to db: %v\n", err)
		os.Exit(1)
	}
	db = client.Database(database)
	collections = db.Collection(collectionName)

	listen, err := net.Listen(protocol, port)

	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	model.RegisterDBServiceServer(gRPCServer, &server{})
	reflection.Register(gRPCServer)
	err = gRPCServer.Serve(listen)

	if err != nil {
		log.Fatal(err)
	}
}
