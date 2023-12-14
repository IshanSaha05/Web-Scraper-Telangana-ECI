package mongodb

import (
	"context"
	"fmt"

	"github.com/IshanSaha05/Web-Scrapper-Telegana/pkg/config"
	"github.com/IshanSaha05/Web-Scrapper-Telegana/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	context    context.Context
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func (object *MongoDB) GetMongoClient() error {

	// Declaring and connecting to a client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.DefaultMongoSite))

	// If there is no error, then we need to return a valid object filled with required details.
	if err == nil {
		object.context = context.Background()
		object.client = client
		object.database = nil
		object.collection = nil
	}

	return err
}

func (object *MongoDB) GetMongoDatabase(databaseName string) error {
	// Getting all the database names present in the client.
	allDBNames, err := object.client.ListDatabaseNames(object.context, bson.D{})

	if err != nil {
		fmt.Println("Error while fetching all database names present, to compare if the passed name already exists or not.")
		return err
	}

	// Parsing through the list of all the database names.
	for _, name := range allDBNames {

		// If the database with the passed name is already present, error is thrown.
		if name == databaseName {
			return fmt.Errorf(fmt.Sprintf("Database named \"%s\" already exists.", databaseName))
		}
	}

	// Database with passed name is not present, thus new database is created.
	object.database = object.client.Database(databaseName)

	return nil
}

func (object *MongoDB) GetMongoCollection(collectionName string) error {
	// Getting all the collection names present in the database.
	allCollectionNames, err := object.database.ListCollectionNames(object.context, bson.D{})

	if err != nil {
		fmt.Println("Error while fetching all collection names present, to compare if the passed name already exists or not.")
		return err
	}

	// Parsing through the list of all the collection names.
	for _, name := range allCollectionNames {

		// If the collection with the passed name is already present, just the collection is assigned.
		if name == collectionName {
			object.collection = object.database.Collection(collectionName)
			return nil
		}
	}

	// Collection with passed name is not present, thus new collection is created.
	err = object.database.CreateCollection(object.context, collectionName)

	if err != nil {
		return err
	}

	object.collection = object.database.Collection(collectionName)

	return nil
}

func (object *MongoDB) InsertIntoDB(results []models.Results) error {
	for _, result := range results {
		_, err := object.collection.InsertOne(object.context, result)

		if err != nil {
			return err
		}
	}

	return nil
}
