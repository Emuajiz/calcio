package group

import (
	"calcio/internal/helper"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type GroupDao struct {
	ID   bson.ObjectID `bson:"_id,omitempty"`
	Name string        `bson:"name"`
}

func (d *GroupDao) ToEntity() *Group {

	id := helper.ToStringID(d.ID)

	return &Group{
		ID:   id,
		Name: d.Name,
	}
}

func FromEntity(e *Group) *GroupDao {

	objId := helper.FromStringID(e.ID)

	return &GroupDao{
		ID:   objId,
		Name: e.Name,
	}
}

type GroupMongoRepo struct {
	db *mongo.Collection
}

func NewGroupMongoRepo(db *mongo.Database) *GroupMongoRepo {
	return &GroupMongoRepo{db: db.Collection("groups")}
}

func (r *GroupMongoRepo) Create(group *Group) error {
	_, err := r.db.InsertOne(context.Background(), group)
	return err
}
