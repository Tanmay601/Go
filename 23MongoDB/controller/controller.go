package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
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

	fmt.Println("Collection instance is ready")

}

//Mongo db helper

//insert 1 record

func insertOneMovie(movie model.Netlfix) { //helped method

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)

}

// update 1 record

func updateOneMovie(movie string) {

	id, _ := primitive.ObjectID(moviedId)
	filter := bson.M{"_id": id} //bson not json imp

	//update to set
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)

}

// delete 1 record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got delete with delete count: ", deleteCount)

}

//delete all record from collection

func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from collection

func getAllMovies() []primitive.M {

	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) { //return of cursor - a heap of data not direct values

		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movie = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies

}

// Actual controller - file which will be called by main.go

func GetAllMovies(w http.ResponseWriter, r *http.Request){
	"w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
	