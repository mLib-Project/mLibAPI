package database

import (
	context "context"
	log "log"
	time "time"

	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"

	Config "mLibAPI/src/config"
)

var Client, _ = mongo.NewClient(options.Client().ApplyURI(Config.DBAdress))
var Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
var db = Client.Database("mo1394_lightning")
var links = db.Collection("links")

func Init() {
	var err = Client.Connect(Ctx)
	if err != nil {
		log.Fatal(err)
	}
}
