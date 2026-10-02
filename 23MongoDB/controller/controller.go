package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const connectionString = "mongodb+srv://tanmayrajpoot07_db_user:W5xkAbZpFiUIWjOH@cluster0.sovcjyl.mongodb.net/?appName=Cluster0&compressors=zlib"
const dbName = "Netflix"
const colName = "watchlist"

//msot imp

var collection *mongo.Collection

// connect with mongo db

func init() {

	//client option

	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOption) //todo is a contect  here to avoid keeping system alive after use
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo db connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance or reference

}
