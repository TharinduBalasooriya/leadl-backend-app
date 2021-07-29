package datamodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ProjectId   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProjectName string             `json:"projectname"`
	UserId      string             `json:"userid"`
	ExpireDate  string             `json:"expiredate"`
}
