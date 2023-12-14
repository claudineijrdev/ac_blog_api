package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongodbComment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	PostID  string             `bson:"post_id" json:"postId" validate:"required"`
	UserID  string             `bson:"user_id" json:"userId" validate:"required"`
	Content string             `bson:"content" json:"content" validate:"required"`
}

type MongoDbPost struct {
	UserID   string           `bson:"user_id" json:"userId" validate:"required"`
	Title    string           `bson:"title" json:"title" validate:"required"`
	Content  string           `bson:"content" json:"content" validate:"required"`
	Comments []MongodbComment `bson:"comments" json:"comments"`
}
