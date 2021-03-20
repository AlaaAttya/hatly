package user

import "go.mongodb.org/mongo-driver/mongo"

type MongoDBRepository struct {
	mongoClient *mongo.Client
}

type Repository interface {
	Save(user User) error
}

func NewMongoDBRepository(mongoClient *mongo.Client) *MongoDBRepository {
	return &MongoDBRepository{
		mongoClient: mongoClient,
	}
}

func (r *MongoDBRepository) Save(user User) error {
	return nil
}
