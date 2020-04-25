package db

import (
	"context"
	"time"

	"github.com/YaelJBS/Practicing_Golang/models"
)

func InsertPost(newPost models.Post) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("blog")
	collection := db.Collection("posts")

	_, err := collection.InsertOne(ctx, newPost)
	if err != nil {
		return "", false, err
	}

	return "", true, nil
}
