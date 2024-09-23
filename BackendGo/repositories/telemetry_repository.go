package repositories

import (
    "BackendGo/config"
    "BackendGo/models"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/primitive"
    "log"
)

var client *mongo.Client

func init() {
    var err error
    clientOptions := options.Client().ApplyURI(config.Config.DatabaseURL)
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
}

func InsertGyroscopeData(data models.GyroscopeData) error {
    collection := client.Database("telemetry").Collection("gyroscope_data")
    _, err := collection.InsertOne(context.Background(), data)
    return err
}

func InsertGpsData(data models.GpsData) error {
    collection := client.Database("telemetry").Collection("gps_data")
    _, err := collection.InsertOne(context.Background(), data)
    return err
}

func InsertPhotoData(data models.PhotoData) error {
    collection := client.Database("telemetry").Collection("photo_data")
    _, err := collection.InsertOne(context.Background(), data)
    return err
}
