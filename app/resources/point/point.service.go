package point

import (
	"expl_app/app/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserOrgPoints struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	User primitive.ObjectID `json:"user,omitempty"`
	// Organisation organisation.Organisation `json:"organisation,omitempty"` // omited in example
	Points    int                `json:"points,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updated_at"`
}

func GetUserPoints(userId primitive.ObjectID, query struct{}) ([]UserOrgPoints, error) {
	repo := NewPointRepository(db.Database)
	return repo.GetUserPoints(userId, query)
}

//
