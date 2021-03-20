package session

import (
	"context"
	"errors"
	"github.com/AlaaAttya/hatly/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoDBRepository struct {
	mongoClient *mongo.Client
}

type Repository interface {
	Save(session Session) (string, error)
	FindActiveSessionByCode(code string) (*Session, error)
	AddUserToSession(user user.User) error
}

func NewMongoDBRepository(mongoClient *mongo.Client) *MongoDBRepository {
	return &MongoDBRepository{
		mongoClient: mongoClient,
	}
}

func (r *MongoDBRepository) AddUserToSession(user user.User) error {
	return nil
}

func (r *MongoDBRepository) FindActiveSessionByCode(code string) (*Session, error) {
	collection := r.mongoClient.Database("hatly").Collection("sessions")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{
		{"code", code},
		{"isActive", true},
	}
	session := &Session{}
	err := collection.FindOne(ctx, filter).Decode(session)
	if err != nil {
		return nil, errors.New("failed to find session from db " + err.Error())
	}

	return session, nil
}

func (r *MongoDBRepository) Save(session Session) (string, error) {
	collection := r.mongoClient.Database("hatly").Collection("sessions")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{
		{"code", session.Code},
	})
	if err != nil {
		return "", err
	}
	id := res.InsertedID

	return id.(primitive.ObjectID).Hex(), nil
}
