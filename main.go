package main

import (
	"alexandria/handler"
	"alexandria/helper"
	"alexandria/note"
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_NAME := os.Getenv("DB_USERNAME")
	DB_PASS := os.Getenv("DB_PASSWORD")

	//register connection
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + DB_NAME + ":" + DB_PASS + "@cluster0.6n5pt.mongodb.net/alexandria-development?retryWrites=true&w=majority")

	//connect
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//define collections
	devDatabase := client.Database("alexandria-development")

	//NOTE
	noteRepository := note.NewRepository(devDatabase)
	noteService := note.NewService(noteRepository)
	noteHandler := handler.NewNoteHandler(noteService)

	//ROUTE CONFIG
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("api/v1")

	//NOTES ROUTES
	api.POST("/notes", authMiddleware(), noteHandler.CreateNote)
	api.PUT("/notes/:id", authMiddleware(), noteHandler.UpdateNote)

	router.Run()

	//create note instance
	// notes, err := notesCollection.InsertOne(ctx, bson.D{
	// 	{"title", "Flutter Tweaks"},
	// 	{"author", "Andreas Notokusumo"},
	// 	{"tags", bson.A{
	// 		"flutter", "mobile", "coding",
	// 	}},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(notes.InsertedID)

	//update ONE note
	// id, _ := primitive.ObjectIDFromHex("61d93e56d9ce3ca9b6adf60e")
	// result, err := notesCollection.UpdateOne(
	// 	ctx,
	// 	bson.M{"_id": id},
	// 	bson.D{
	// 		{"$set", bson.D{
	// 			{"tags", bson.A{
	// 				"flutter", "tweaks", "mobile",
	// 			}},
	// 		}},
	// 	},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	//update MANY note
	// result, err := notesCollection.UpdateMany(
	// 	ctx,
	// 	bson.M{"title": "Flutter Tweaks"},
	// 	bson.D{
	// 		{"$set", bson.D{{"author", "Andrea"}}},
	// 	},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	//delete ONE note
	// id, _ := primitive.ObjectIDFromHex("61d94b103656bcf08c7cc898")
	// result, err := notesCollection.DeleteOne(
	// 	ctx,
	// 	bson.M{
	// 		"_id": id,
	// 	},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	//fetch all notes
	// cursor, err := notesCollection.Find(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var notes []bson.M
	// err = cursor.All(ctx, &notes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, note := range notes {
	// 	fmt.Println(note["title"])
	// }

	//fetch one note
	// var note bson.M
	// err = notesCollection.FindOne(ctx, bson.M{}).Decode(&note)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(note)

	//filter notes
	// filterCursor, err := notesCollection.Find(ctx, bson.M{"title": "Flutter Tweaks"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var filteredNotes []bson.M
	// err = filterCursor.All(ctx, &filteredNotes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(filteredNotes)

	//SORT AND FILTER BY STRING
	// opts := options.Find()
	// opts.SetSort(bson.D{
	// 	{"title", -1},
	// })
	// sortCursor, _ := notesCollection.Find(ctx, bson.D{
	// 	{"title", primitive.Regex{Pattern: "Flut", Options: ""}}}, opts)
	// var sortedNotes []bson.M
	// err = sortCursor.All(ctx, &sortedNotes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(sortedNotes)
}

func authMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse(
				"Unauthorized request",
				http.StatusUnauthorized,
				"",
				nil,
			)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var uid string
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			uid = tokenArray[1]
		}

		context.Set("currentUID", uid)
	}
}
