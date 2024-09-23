package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GyroscopeData struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    X         float64            `json:"x" bson:"x"`
    Y         float64            `json:"y" bson:"y"`
    Z         float64            `json:"z" bson:"z"`
    Timestamp int64              `json:"timestamp" bson:"timestamp"`
    DeviceID  string             `json:"deviceId" bson:"deviceId"`
}
