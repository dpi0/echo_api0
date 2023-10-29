package main

import (
	"fmt"
	"net/http"
	// "log"
	"os"
	"context"
	"time"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"

)

func init() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}


var client *mongo.Client // Declare client at the package level


func main(){
    mongoUri := os.Getenv("ATLAS_URI")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var err error
    client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

    if err != nil {
        panic(err)
    }

    // Defer the disconnection
    defer func() {
        if err = client.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()

    err = client.Ping(ctx, nil)

    if err != nil {
        fmt.Println("There was a problem connecting to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
        panic(err)
    }

    fmt.Println("Connected to MongoDB! ✅")

	UseClientFromHelper()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	fmt.Println("XXXXXXXXX", os.Getenv("CLUSTER_USERNAME"))

	e.Logger.Fatal(e.Start(":9001"))
}

func hello(c echo.Context) error {
	data := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	return c.JSONPretty(http.StatusOK, data, "  ")
}
