package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GpsData struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Latitude  float64            `json:"latitude" bson:"latitude"`
    Longitude float64            `json:"longitude" bson:"longitude"`
    Timestamp int64              `json:"timestamp" bson:"timestamp"`
    DeviceID  string             `json:"deviceId" bson:"deviceId"`
}
