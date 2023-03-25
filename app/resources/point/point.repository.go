package point

import (
	"context"
	"expl_app/app/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	pointRepoInstance *pointRepository
)

type PointRepository interface {
	GetUserPoints(userId primitive.ObjectID, query struct{}) ([]UserOrgPoints, error)
}

type pointRepository struct {
	db *mongo.Database
}

func NewPointRepository(db *mongo.Database) *pointRepository {
	pointRepoInstance = &pointRepository{db: db}
	return pointRepoInstance
}

// ==================================================
// Get instance
// ==================================================
func GetPointRepositoryInstance() *pointRepository {
	if pointRepoInstance == nil {
		pointRepoInstance = NewPointRepository(db.Database)
	}
	return pointRepoInstance
}

// ==================================================
// GetUserPoints
// ==================================================
func (r pointRepository) GetUserPoints(userId primitive.ObjectID, query struct{}) ([]UserOrgPoints, error) {
	aggr := bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "user", Value: userId}}}},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "organisations"},
					{Key: "localField", Value: "organisation"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "organisation"},
				},
			},
		},
		bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$organisation"}}}},
	}

	userOrgPointsCollection := db.GetCollection(db.Connection, "userorgpoints")

	cursor, err := userOrgPointsCollection.Aggregate(context.Background(), aggr)

	if err != nil {
		return []UserOrgPoints{}, err
	}

	defer cursor.Close(context.Background())

	var pointsArray []UserOrgPoints

	for cursor.Next(context.Background()) {
		var points UserOrgPoints
		cursor.Decode(&points)
		pointsArray = append(pointsArray, points)
	}

	return pointsArray, nil
}
