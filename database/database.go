package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Model[T any] struct {
 ID        primitive.ObjectID `bson:"_id,omitempty"`
 CreatedAt time.Time          `bson:"created_at"`
 UpdatedAt time.Time          `bson:"updated_at"`
}

var Client *mongo.Client

func Connect(uri string) (*mongo.Client, error) {
 clientOptions := options.Client().ApplyURI(uri)
 client, err := mongo.Connect(context.Background(), clientOptions)

 if err != nil {
  log.Fatal(err)
 }

 err = client.Ping(context.Background(), nil)
 if err != nil {
  log.Fatal(err)
 }

 fmt.Println("Successfully connected to MongoDB")
 Client = client
 return client, nil
}


func (m *Model[T]) Create(ctx context.Context, db *mongo.Database, collectionName string, model interface{}) error {
 collection := db.Collection(collectionName)

 m.CreatedAt = time.Now()
 m.UpdatedAt = time.Now()

 res, err := collection.InsertOne(ctx, model)
 if err != nil {
  return err
 }

 m.ID = res.InsertedID.(primitive.ObjectID)
 return nil
}

func (m *Model[T]) Read(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
 collection := db.Collection(collectionName)

 err := collection.FindOne(ctx, filter).Decode(result)
 if err != nil {
  return err
 }

 return nil
}

func (m *Model[T]) ReadMany(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result *[]T) error {
 collection := db.Collection(collectionName)

 cursor, err := collection.Find(ctx, filter)
 if err != nil {
  return err
 }
 defer cursor.Close(ctx)

 for cursor.Next(ctx) {
  var document T
  if err := cursor.Decode(&document); err != nil {
    return err
  }

  *result = append(*result, document)
 }

 return nil
}

func (m *Model[T]) Update(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, update interface{}) error {
 collection := db.Collection(collectionName)

 m.UpdatedAt = time.Now()

 _, err := collection.UpdateOne(ctx, filter, update)
 if err != nil {
 return err
}

return nil
}

func (m *Model[T]) Delete(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}) error {
  collection := db.Collection(collectionName)
  _, err := collection.DeleteOne(ctx, filter)
  if err != nil {
   return err
  }

return nil
}