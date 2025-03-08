package repository

import (
	"context"

	"user-service/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type userRepo struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(ctx context.Context, user entity.User) error {
	_, err := r.db.InsertOne(ctx, user)
	return err
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
