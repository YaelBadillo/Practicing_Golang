package models

type Post struct {
	//ID     primitive.ObjectID `bson: "_id" json: "id"`
	Name   string `bson: "name" json "name"`
	Text   string `bson: "text" json "text"`
	Writer string `bsonL: "writer" json: "writer"`
}
