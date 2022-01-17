package main

import (
	"alexandria/activity"
	"alexandria/comment"
	"alexandria/handler"
	"alexandria/helper"
	"alexandria/note"
	"alexandria/user"
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

	//USER
	userRepository := user.NewRepository(devDatabase)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	//ACTIVITY
	activityRepository := activity.NewRepository(devDatabase)
	activityService := activity.NewService(activityRepository)
	activityHandler := handler.NewActivityHandler(activityService)

	//NOTE
	noteRepository := note.NewRepository(devDatabase)
	noteService := note.NewService(noteRepository)
	noteHandler := handler.NewNoteHandler(noteService, activityService)

	//COMMENT
	commentRepository := comment.NewRepository(devDatabase)
	commentService := comment.NewService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService, activityService)

	//ROUTE CONFIG
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("api/v1")

	//USER ROUTES
	api.POST("/users", userHandler.RegisterUser)
	api.PUT("/users", authMiddleware(), userHandler.UpdateUser)
	api.DELETE("/users/:uid", authMiddleware(), userHandler.DeleteUser)
	api.GET("/users/:uid", userHandler.GetUserByUID)

	//NOTES ROUTES
	api.POST("/notes", authMiddleware(), noteHandler.CreateNote)
	api.PUT("/notes/:id", authMiddleware(), noteHandler.UpdateNote)
	api.DELETE("/notes/:id", authMiddleware(), noteHandler.DeleteNote)
	api.GET("/notes", noteHandler.GetNotes)
	api.GET("/notes/:id", authMiddleware(), noteHandler.GetNoteByID)

	//COMMENTS
	api.POST("/comments", authMiddleware(), commentHandler.CreateComment)
	api.GET("/notes/:id/comments/", commentHandler.GetCommentsByNoteID)
	api.GET("/notes/:id/comments/:comment_id", commentHandler.GetCommentByID)
	api.DELETE("/notes/:id/comments/:comment_id", authMiddleware(), commentHandler.DeleteCommentByID)

	//ACTIVITIES
	api.GET("/activity/:uid", authMiddleware(), activityHandler.GetActivityByAffiliateID)
	api.PUT("/activity/:id", authMiddleware(), activityHandler.MarkActivityAsRead)

	router.Run()
}

func authMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		var uid string

		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			uid = tokenArray[1]
		}

		if !strings.Contains(authHeader, "Bearer") || uid == "" {
			response := helper.APIResponse(
				"Unauthorized request",
				http.StatusUnauthorized,
				"",
				nil,
			)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		context.Set("currentUID", uid)
	}
}
