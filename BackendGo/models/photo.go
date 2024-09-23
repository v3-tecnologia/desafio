package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PhotoData struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    PhotoPath string             `json:"photoPath" bson:"photoPath"`
    Timestamp int64              `json:"timestamp" bson:"timestamp"`
    DeviceID  string             `json:"deviceId" bson:"deviceId"`
}
