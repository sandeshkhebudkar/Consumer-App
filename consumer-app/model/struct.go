package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//DBLayer is dblayer structure of book
type DBLayer struct {
	//ID        *primitive.ObjectID `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

//Save use to save data
func Save(client *mongo.Client, dbBlog DBLayer) (interface{}, error) {

	collection := client.Database("eBook").Collection("book")
	doc, err := collection.InsertOne(context.TODO(), dbBlog)
	if err != nil {
		return "", err
	}
	return doc.InsertedID, nil
}
