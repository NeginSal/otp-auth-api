package repository

import (
	"context"
	"time"

	"github.com/NeginSal/otp-auth-api/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OTPRepository struct {
	Collection *mongo.Collection
}

func NewOTPRepository(db *mongo.Client) *OTPRepository {
	collection := db.Database("otp_auth").Collection("otp_requests")
	return &OTPRepository{Collection: collection}
}

func (r *OTPRepository) SaveOTP(ctx context.Context, otp *model.OTPRequest) error {
	otp.CreatedAt = time.Now()
	_, err := r.Collection.InsertOne(ctx, otp)
	return err
}

func (r *OTPRepository) GetLatestByPhone(ctx context.Context, phone string) (*model.OTPRequest, error) {
	var otp model.OTPRequest
	filter := bson.M{"phone": phone}
	opts := options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}})
	err := r.Collection.FindOne(ctx, filter, opts).Decode(&otp)
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *OTPRepository) MarkVerified(ctx context.Context, otpID string) error {
	filter := bson.M{"_id": otpID}
	update := bson.M{"$set": bson.M{"verified": true}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *OTPRepository) CountRecentRequests(ctx context.Context, phone string, since time.Duration) (int64, error) {
	from := time.Now().Add(-since)
	filter := bson.M{
		"phone":      phone,
		"created_at": bson.M{"$gte": from},
	}
	return r.Collection.CountDocuments(ctx, filter)
}