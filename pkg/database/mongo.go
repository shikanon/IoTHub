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

// 连接mongodb中的数据库，返回数据库连接
func MongoDbClient() (connect *mongo.Database) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	host := config.MongodbConfig.Host
	port := config.MongodbConfig.Port
	db := config.MongodbConfig.Db
	url := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println("ERR:", err)
	}
	db_client := client.Database(db)
	return db_client
}

// 插入一条数据
func MongoDbInsertOneData(collection_name string, data bson.M) (id_str string) {
	db_client := MongoDbClient()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := db_client.Collection(collection_name)
	res, _ := collection.InsertOne(ctx, &data)
	id := res.InsertedID.(primitive.ObjectID).Hex() // objectID转string
	return id
}

// 根据条件，获取一条数据
func MongoDbGetFilterData(collection_name string, filter bson.M) (result bson.M) {
	db_client := MongoDbClient()
	data := bson.M{}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := db_client.Collection(collection_name)
	err := collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

// 删除一条记录
func MongodbDeleteOneData(collection_name string, _id string){
	id, _ := primitive.ObjectIDFromHex(_id) // string转objectID
	db_client := MongoDbClient()
	collection := db_client.Collection(collection_name)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"_id": id}

	collection.DeleteOne(ctx, filter)
}