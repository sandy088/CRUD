package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	model "saaster.tech/crud/internal/models"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, database string) *UserRepository {
	return &UserRepository{
		Collection: client.Database(database).Collection("users"),
	}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User

	cursor, err := r.Collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User

		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(ctx, user)

	return err
}
