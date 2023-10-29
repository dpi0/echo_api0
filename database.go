package main

import (
	"fmt"
	"context"
	"os"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)


func DBInit(){
	var mongoUri = os.Getenv("ATLAS_URI")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoUri,
	))
	
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
}
