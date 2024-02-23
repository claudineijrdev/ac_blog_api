package db

import (
	"ac_blog_api/internal/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	entity.PostRepository
	mongoDb *mongo.Client
}

func NewMongoRepository(mongoDb *mongo.Client) *MongoRepository {
	return &MongoRepository{
		mongoDb: mongoDb,
	}
}

func (m *MongoRepository) CreatePost(post entity.Post) error {
	collection := GetCollection("posts", "blog")
	_, error := collection.InsertOne(context.TODO(), post)
	return error
}

func (m *MongoRepository) GetPostList() ([]entity.Post, error) {
	collection := GetCollection("posts", "blog")
	cursor, error := collection.Find(context.TODO(), bson.D{{}})

	if error != nil {
		return nil, error
	}

	var posts []entity.Post
	for cursor.Next(context.TODO()) {
		var post entity.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	return posts, nil

}

func (m *MongoRepository) GetPost(id string) (entity.Post, error) {
	var post entity.Post
	collection := GetCollection("posts", "blog")
	pID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Post{}, err
	}
	filter := bson.M{"_id": pID}
	error := collection.FindOne(context.TODO(), filter).Decode(&post)
	if error != nil {
		return entity.Post{}, error
	}

	return post, nil
}

func (m *MongoRepository) UpdatePost(id string, post entity.Post) error {
	pID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := GetCollection("posts", "blog")
	filter := bson.M{"_id": pID}

	update := bson.M{
		"$set": bson.M{
			"title":   post.Title,
			"content": post.Content,
		},
	}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	return err
}

func (m *MongoRepository) DeletePost(id string) error {
	pID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := GetCollection("posts", "blog")
	filter := bson.M{"_id": pID}

	_, err = collection.DeleteOne(context.Background(), filter)

	return err
}

func (m *MongoRepository) CreateComment(comment entity.Comment) error {
	collection := GetCollection("posts", "blog")
	pID, err := primitive.ObjectIDFromHex(comment.PostID)
	if err != nil {
		return err
	}

	comment.ID = primitive.NewObjectID()

	filter := bson.M{"_id": pID}

	update := bson.M{
		"$push": bson.M{
			"comments": comment,
		},
	}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	return err
}

func (m *MongoRepository) GetCommentList(postId string) ([]entity.Comment, error) {
	post, err := m.GetPost(postId)

	if err != nil {
		return nil, err
	}

	return post.Comments, nil

}

func (m *MongoRepository) DeletePostComment(postId string, commentID string) error {
	pID, err := primitive.ObjectIDFromHex(postId)

	if err != nil {
		return err
	}
	cID, err := primitive.ObjectIDFromHex(commentID)

	if err != nil {
		return err
	}

	collection := GetCollection("posts", "blog")

	filter := bson.M{"_id": pID}

	update := bson.M{
		"$pull": bson.M{
			"comments": bson.M{
				"_id": cID,
			},
		},
	}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	return err
}
