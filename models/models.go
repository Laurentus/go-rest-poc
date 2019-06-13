package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Using golang tags to describe:
// - json (http response)
// - bson (mongo request)
// Marshalling follows this bidirectionally
type Campaign struct {
	MongoId primitive.ObjectID `json:"_id" bson:"_id"`
	Name    string             `json:"name,omitempty"`
	Id      string             `json:"id,omitempty"`
}

type Adgroup struct {
	MongoId    primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name,omitempty"`
	Id         string             `json:"id,omitempty"`
	CampaingId primitive.ObjectID `json:"campaign_id" bson:"campaign_id"`
}

type Ad struct {
	MongoId    primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name,omitempty"`
	Id         string             `json:"id,omitempty"`
	CampaingId primitive.ObjectID `json:"campaign_id" bson:"campaign_id"`
	AdgroupId  primitive.ObjectID `json:"adgroup_id" bson:"adgroup_id"`
}
