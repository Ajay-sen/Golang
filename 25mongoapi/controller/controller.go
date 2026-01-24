package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ajay-sen/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString ="<put the mongodb connection string here>"
const dbName ="netflix"
const colName ="watchlist"

//MOST IMPORTANT
var collection *mongo.Collection

//connect with mongoDB
func init(){
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection =client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

//MONGODB helpers - file


//insert 1 record
func insertOneMovie(movie model.Netflix){
	inserted, err:=collection.InsertOne(context.Background(), movie)

	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in the db wih id: ", inserted.InsertedID)
}


//update 1 record
func updateOneMovie(movieId string){
	id, _:= primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}

	result, err:= collection.UpdateOne(context.Background(),filter,update)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println("modified  count:", result.ModifiedCount)
}


//delete 1 record
func deleteOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter :=bson.M{"_id":id}
	deleteCount,err := collection.DeleteOne(context.Background(),filter)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Movie got dleete with delete count: ", deleteCount)
}

//delete all recourds
func deleteAllMovie() int64{
	// filter := bson.D{{}}
	// collection.DeleteMany(context.Background())
	//instead you can write
	deleteResult, err := collection.DeleteMany(context.Background(),bson.D{{}},nil)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted are :",deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}


//get all movies from database
func getAllMovies() []primitive.M{
	cur, err := collection.Find(context.Background(),bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)

		if err != nil{
			log.Fatal(err)
		}
		movies = append(movies,movie)
	}

	defer cur.Close(context.Background())
	return movies
}


//actual controller - file
// the lower case function name s can not be expoorted to different files , hence need to pass in a different way 

func GetMyAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}


// function for creating movie
func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}


//marking movie watched , using update movie helper method
func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Methods","PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}


func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params :=mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// function to delete all the  movies using helper methods
func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}