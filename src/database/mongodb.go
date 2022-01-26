package database

import (
	context "context"
	fmt "fmt"
	log "log"
	time "time"

	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"

	Config "mLibAPI/src/config"
)

const (
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	username := Config.DB_Username
	password := Config.DB_Password
	clusterEndpoint := Config.DB_ClusterEndpoint

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil && Config.Mode == "DEV" {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil && Config.Mode == "DEV" {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil && Config.Mode == "DEV" {
		log.Printf("Failed to ping cluster: %v", err)
	}

	if Config.Mode == "DEV" {
		fmt.Println("Connected to MongoDB!")
	}

	return client, ctx, cancel
}

// var Client, _ = mongo.NewClient(options.Client().ApplyURI(Config.DBAdress))
// var Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
// var db = Client.Database("mo1394_lightning")
// var links = db.Collection("links")

// func Init() {
// 	var err = Client.Connect(Ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
