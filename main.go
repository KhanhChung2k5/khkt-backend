package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"khkt-backend/controllers"
	"khkt-backend/services"
	"log"
	"os"
)

var (
	server       *gin.Engine
	cs           services.ChemistryService
	cc           controllers.ChemistryController
	ctx          context.Context
	chemistryCol *mongo.Collection
	refDocCol    *mongo.Collection
	mongoclient  *mongo.Client
	err          error
)

const (
	USERNAME = "username"
	PASSWORD = "password"
)

func init() {
	ctx = context.TODO()

	//uriConn := "mongodb://" + USER + ":" + PASSWORD + "@" + HOST + ":" + PORT + "/" + DB_NAME + "?authSource=admin"
	uriConn := "mongodb+srv://tuankiet:kietlu1712@bankaccount.lfuju.mongodb.net/?retryWrites=true&w=majority"
	option := options.Client().ApplyURI(uriConn)
	mongoclient, err = mongo.Connect(ctx, option)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	chemistryCol = mongoclient.Database("chemistry").Collection("chemistry")
	refDocCol = mongoclient.Database("chemistry").Collection("ref_document")

	cs = services.NewUserService(chemistryCol, refDocCol, ctx)
	cc = controllers.New(cs)
	server = gin.Default()
	server.Use(cors.Default())
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	cc.RegisterUserRoutes(basepath)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	server.Run(":" + port)
}
