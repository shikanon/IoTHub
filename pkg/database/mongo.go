package database

import (
	"context"
	"fmt"
	"github.com/shikanon/IoTOrbHub/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func MongoDbClient() (connect *mongo.Database, msg string) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	host := config.MongodbConfig.Host
	port := config.MongodbConfig.Port
	db := config.MongodbConfig.Db
	url := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, "mongodb数据库连接失败"
	}
	db_client := client.Database(db)
	return db_client, ""
}

func MongoDbInsertOneData(collection_name string, data bson.M) (id_str, mag string) {
	db_client, msg := MongoDbClient()
	if db_client == nil {
		return "", msg
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := db_client.Collection(collection_name)
	res, _ := collection.InsertOne(ctx, &data)
	id := res.InsertedID.(primitive.ObjectID).Hex() // objectID转string
	return id, ""
}

func MongoDbGetFilterData(collection_name string, filter bson.M) (result bson.M, msg string) {
	db_client, msg := MongoDbClient()
	if db_client == nil {
		return nil, ""
	}
	data := bson.M{}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := db_client.Collection(collection_name)
	err := collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, "数据查找失败"
	}
	return data, ""
}

func MongodbDeleteOneData(collection_name string, _id string) {
	id, _ := primitive.ObjectIDFromHex(_id) // string转objectID
	db_client, _ := MongoDbClient()
	collection := db_client.Collection(collection_name)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"_id": id}

	collection.DeleteOne(ctx, filter)
}
