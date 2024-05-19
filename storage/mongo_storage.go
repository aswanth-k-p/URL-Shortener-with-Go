package storage

import (
	"context"
	"log"
	"time"
	"url-shortener/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStorage(uri, dbName, collectionName string) *MongoStorage {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &MongoStorage{client: client, collection: collection}
}

func (s *MongoStorage) Save(url models.URL) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := s.collection.InsertOne(ctx, url)
	return err
}

func (s *MongoStorage) Get(short string) (models.URL, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var url models.URL
	err := s.collection.FindOne(ctx, bson.M{"short": short}).Decode(&url)
	if err != nil {
		return url, false
	}
	return url, true
}
