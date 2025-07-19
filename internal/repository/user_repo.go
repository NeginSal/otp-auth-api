package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/NeginSal/otp-auth-api/internal/model"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	collection := db.Database("otp_auth").Collection("users")
	return &UserRepository{Collection: collection}
}

func (r *UserRepository) FindByPhone(ctx context.Context, phone string) (*model.User, error) {
	var user model.User
	err := r.Collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	user.CreatedAt = time.Now()
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}
